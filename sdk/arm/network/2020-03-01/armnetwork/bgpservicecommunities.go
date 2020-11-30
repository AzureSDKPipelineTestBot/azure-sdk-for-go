// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// BgpServiceCommunitiesOperations contains the methods for the BgpServiceCommunities group.
type BgpServiceCommunitiesOperations interface {
	// List - Gets all the available bgp service communities.
	List() (BgpServiceCommunityListResultPager, error)
}

// bgpServiceCommunitiesOperations implements the BgpServiceCommunitiesOperations interface.
type bgpServiceCommunitiesOperations struct {
	*Client
	subscriptionID string
}

// List - Gets all the available bgp service communities.
func (client *bgpServiceCommunitiesOperations) List() (BgpServiceCommunityListResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &bgpServiceCommunityListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *BgpServiceCommunityListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.BgpServiceCommunityListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.BgpServiceCommunityListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *bgpServiceCommunitiesOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/bgpServiceCommunities"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *bgpServiceCommunitiesOperations) listHandleResponse(resp *azcore.Response) (*BgpServiceCommunityListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := BgpServiceCommunityListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.BgpServiceCommunityListResult)
}

// listHandleError handles the List error response.
func (client *bgpServiceCommunitiesOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}