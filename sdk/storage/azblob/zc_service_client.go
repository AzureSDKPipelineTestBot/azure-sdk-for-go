// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azblob

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/uuid"
	"net/url"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

//nolint
const (
	// ContainerNameRoot is the special Azure Storage name used to identify a storage account's root container.
	ContainerNameRoot = "$root"

	// ContainerNameLogs is the special Azure Storage name used to identify a storage account's logs container.
	ContainerNameLogs = "$logs"
)

// A ServiceClient represents a URL to the Azure Storage Blob service allowing you to manipulate blob containers.
type ServiceClient struct {
	client *serviceClient
	u      url.URL
	cred   StorageAccountCredential
}

// URL returns the URL endpoint used by the ServiceClient object.
func (s ServiceClient) URL() string {
	return s.client.con.u
}

// NewServiceClient creates a ServiceClient object using the specified URL, credential, and options.
// Example of serviceURL: https://<your_storage_account>.blob.core.windows.net
func NewServiceClient(serviceURL string, cred azcore.Credential, options *ClientOptions) (ServiceClient, error) {
	u, err := url.Parse(serviceURL)
	if err != nil {
		return ServiceClient{}, err
	}

	c, _ := cred.(*SharedKeyCredential)

	return ServiceClient{client: &serviceClient{
		con: newConnection(serviceURL, cred, options.getConnectionOptions()),
	}, u: *u, cred: c}, nil
}

// NewServiceClientFromConnectionString creates a service client from the given connection string.
//nolint
func NewServiceClientFromConnectionString(connectionString string, options *ClientOptions) (ServiceClient, error) {
	primaryURL, _, cred, err := ParseConnectionString(connectionString, "")
	if err != nil {
		return ServiceClient{}, nil
	}

	svcClient, err := NewServiceClient(primaryURL, cred, options)
	return svcClient, err
}

// NewContainerClient creates a new ContainerClient object by concatenating containerName to the end of
// ServiceClient's URL. The new ContainerClient uses the same request policy pipeline as the ServiceClient.
// To change the pipeline, create the ContainerClient and then call its WithPipeline method passing in the
// desired pipeline object. Or, call this package's NewContainerClient instead of calling this object's
// NewContainerClient method.
func (s ServiceClient) NewContainerClient(containerName string) ContainerClient {
	containerURL := appendToURLPath(s.client.con.u, containerName)
	containerConnection := &connection{containerURL, s.client.con.p}
	return ContainerClient{
		client: &containerClient{
			con: containerConnection,
		},
		cred: s.cred,
	}
}

// CreateContainer is a lifecycle method to creates a new container under the specified account.
// If the container with the same name already exists, a ResourceExistsError will
// be raised. This method returns a client with which to interact with the newly
// created container.
func (s ServiceClient) CreateContainer(ctx context.Context, containerName string, options *CreateContainerOptions) (ContainerCreateResponse, error) {
	containerClient := s.NewContainerClient(containerName)
	containerCreateResp, err := containerClient.Create(ctx, options)
	return containerCreateResp, err
}

// DeleteContainer is a lifecycle method that marks the specified container for deletion.
// The container and any blobs contained within it are later deleted during garbage collection.
// If the container is not found, a ResourceNotFoundError will be raised.
func (s ServiceClient) DeleteContainer(ctx context.Context, containerName string, options *DeleteContainerOptions) (ContainerDeleteResponse, error) {
	containerClient := s.NewContainerClient(containerName)
	containerDeleteResp, err := containerClient.Delete(ctx, options)
	return containerDeleteResp, err
}

func (s ServiceClient) NewContainerLeaseClient(containerName string, leaseID *string) ContainerLeaseClient {
	containerURL := appendToURLPath(s.client.con.u, containerName)
	containerConnection := &connection{containerURL, s.client.con.p}

	if leaseID == nil {
		leaseID = to.StringPtr(uuid.New().String())
	}

	return ContainerLeaseClient{
		ContainerClient: ContainerClient{
			client: &containerClient{
				con: containerConnection,
			},
		},
		LeaseID: leaseID,
	}
}

// appendToURLPath appends a string to the end of a URL's path (prefixing the string with a '/' if required)
func appendToURLPath(u string, name string) string {
	// e.g. "https://ms.com/a/b/?k1=v1&k2=v2#f"
	// When you call url.Parse() this is what you'll get:
	//     Scheme: "https"
	//     Opaque: ""
	//       User: nil
	//       Host: "ms.com"
	//       Path: "/a/b/"	This should start with a / and it might or might not have a trailing slash
	//    RawPath: ""
	// ForceQuery: false
	//   RawQuery: "k1=v1&k2=v2"
	//   Fragment: "f"
	uri, _ := url.Parse(u)

	if len(uri.Path) == 0 || uri.Path[len(uri.Path)-1] != '/' {
		uri.Path += "/" // Append "/" to end before appending name
	}
	uri.Path += name
	return uri.String()
}

func (s ServiceClient) GetAccountInfo(ctx context.Context) (ServiceGetAccountInfoResponse, error) {
	resp, err := s.client.GetAccountInfo(ctx, nil)

	return resp, handleError(err)
}

// GetUserDelegationCredential obtains a UserDelegationKey object using the base ServiceClient object.
// OAuth is required for this call, as well as any role that can delegate access to the storage account.
// Strings in KeyInfo should be formatted with SASTimeFormat.
func (s ServiceClient) GetUserDelegationCredential(ctx context.Context, startTime, expiryTime *time.Time) (UserDelegationCredential, error) {
	if startTime == nil {
		startTime = to.TimePtr(time.Now().UTC())
	}

	udk, err := s.client.GetUserDelegationKey(ctx, KeyInfo{
		Start:  to.StringPtr(startTime.UTC().Format(SASTimeFormat)),
		Expiry: to.StringPtr(expiryTime.UTC().Format(SASTimeFormat)),
	}, nil)
	if err != nil {
		return UserDelegationCredential{}, handleError(err)
	}
	urlParts := NewBlobURLParts(s.URL())
	return NewUserDelegationCredential(strings.Split(urlParts.Host, ".")[0], udk.UserDelegationKey), nil
}

// The ListContainersSegment operation returns a pager of the containers under the specified account.
// Use an empty Marker to start enumeration from the beginning. Container names are returned in lexicographic order.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/list-containers2.
func (s ServiceClient) ListContainersSegment(o *ListContainersSegmentOptions) *ServiceListContainersSegmentPager {
	listOptions := o.pointers()
	pager := s.client.ListContainersSegment(listOptions)
	// override the generated advancer, which is incorrect
	if pager.Err() != nil {
		return pager
	}

	pager.advancer = func(cxt context.Context, response ServiceListContainersSegmentResponse) (*azcore.Request, error) {
		if response.ListContainersSegmentResponse.NextMarker == nil {
			return nil, handleError(errors.New("unexpected missing NextMarker"))
		}
		req, err := s.client.listContainersSegmentCreateRequest(cxt, listOptions)
		if err != nil {
			return nil, handleError(err)
		}
		queryValues, _ := url.ParseQuery(req.URL.RawQuery)
		queryValues.Set("marker", *response.ServiceListContainersSegmentResult.NextMarker)

		req.URL.RawQuery = queryValues.Encode()
		return req, nil
	}

	return pager
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics
// and CORS (Cross-Origin Resource Sharing) rules.
func (s ServiceClient) GetProperties(ctx context.Context) (ServiceGetPropertiesResponse, error) {
	resp, err := s.client.GetProperties(ctx, nil)

	return resp, handleError(err)
}

// SetProperties Sets the properties of a storage account's Blob service, including Azure Storage Analytics.
// If an element (e.g. analytics_logging) is left as None, the existing settings on the service for that functionality are preserved.
func (s ServiceClient) SetProperties(ctx context.Context, properties StorageServiceProperties) (ServiceSetPropertiesResponse, error) {
	resp, err := s.client.SetProperties(ctx, properties, nil)

	return resp, handleError(err)
}

// GetStatistics Retrieves statistics related to replication for the Blob service.
// It is only available when read-access geo-redundant replication is enabled for  the storage account.
// With geo-redundant replication, Azure Storage maintains your data durable
// in two locations. In both locations, Azure Storage constantly maintains
// multiple healthy replicas of your data. The location where you read,
// create, update, or delete data is the primary storage account location.
// The primary location exists in the region you choose at the time you
// create an account via the Azure Management Azure classic portal, for
// example, North Central US. The location to which your data is replicated
// is the secondary location. The secondary location is automatically
// determined based on the location of the primary; it is in a second data
// center that resides in the same region as the primary location. Read-only
// access is available from the secondary location, if read-access geo-redundant
// replication is enabled for your storage account.
func (s ServiceClient) GetStatistics(ctx context.Context) (ServiceGetStatisticsResponse, error) {
	resp, err := s.client.GetStatistics(ctx, nil)

	return resp, handleError(err)
}

func (s ServiceClient) CanGetAccountSASToken() bool {
	return s.cred != nil
}

// GetAccountSASToken is a convenience method for generating a SAS token for the currently pointed at account.
// It can only be used if the supplied azcore.Credential during creation was a SharedKeyCredential.
// This validity can be checked with CanGetAccountSASToken().
func (s ServiceClient) GetAccountSASToken(services AccountSASServices, resources AccountSASResourceTypes, permissions AccountSASPermissions, validityTime time.Duration) (SASQueryParameters, error) {
	return AccountSASSignatureValues{
		Version: SASVersion,

		Permissions:   permissions.String(),
		Services:      services.String(),
		ResourceTypes: resources.String(),

		StartTime:  time.Now().UTC(),
		ExpiryTime: time.Now().UTC().Add(validityTime),
	}.NewSASQueryParameters(s.cred.(*SharedKeyCredential))
}

// FindBlobsByTags operation finds all blobs in the storage account whose tags match a given search expression.
// Filter blobs searches across all containers within a storage account but can be scoped within the expression to a single container.
// https://docs.microsoft.com/en-us/rest/api/storageservices/find-blobs-by-tags
// eg. "dog='germanshepherd' and penguin='emperorpenguin'"
// To specify a container, eg. "@container=’containerName’ and Name = ‘C’"
func (s ServiceClient) FindBlobsByTags(ctx context.Context, options ServiceFilterBlobsByTagsOptions) (ServiceFilterBlobsResponse, error) {
	// TODO: Use pager here? Missing support from zz_generated_pagera.go
	serviceFilterBlobsOptions := options.pointer()
	return s.client.FilterBlobs(ctx, serviceFilterBlobsOptions)
}