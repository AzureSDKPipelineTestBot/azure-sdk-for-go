//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// SystemTopicsClient contains the methods for the SystemTopics group.
// Don't use this type directly, use NewSystemTopicsClient() instead.
type SystemTopicsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSystemTopicsClient creates a new instance of SystemTopicsClient with the specified values.
func NewSystemTopicsClient(con *arm.Connection, subscriptionID string) *SystemTopicsClient {
	return &SystemTopicsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Asynchronously creates a new system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsBeginCreateOrUpdateOptions) (SystemTopicsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, systemTopicName, systemTopicInfo, options)
	if err != nil {
		return SystemTopicsCreateOrUpdatePollerResponse{}, err
	}
	result := SystemTopicsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SystemTopicsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return SystemTopicsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SystemTopicsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Asynchronously creates a new system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) createOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, systemTopicName, systemTopicInfo, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SystemTopicsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicInfo SystemTopic, options *SystemTopicsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, systemTopicInfo)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SystemTopicsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Delete existing system topic.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) BeginDelete(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsBeginDeleteOptions) (SystemTopicsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, systemTopicName, options)
	if err != nil {
		return SystemTopicsDeletePollerResponse{}, err
	}
	result := SystemTopicsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SystemTopicsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return SystemTopicsDeletePollerResponse{}, err
	}
	result.Poller = &SystemTopicsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete existing system topic.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) deleteOperation(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, systemTopicName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SystemTopicsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SystemTopicsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Get properties of a system topic.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) Get(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsGetOptions) (SystemTopicsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, systemTopicName, options)
	if err != nil {
		return SystemTopicsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SystemTopicsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SystemTopicsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SystemTopicsClient) getCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SystemTopicsClient) getHandleResponse(resp *http.Response) (SystemTopicsGetResponse, error) {
	result := SystemTopicsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemTopic); err != nil {
		return SystemTopicsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SystemTopicsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByResourceGroup - List all the system topics under a resource group.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) ListByResourceGroup(resourceGroupName string, options *SystemTopicsListByResourceGroupOptions) *SystemTopicsListByResourceGroupPager {
	return &SystemTopicsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SystemTopicsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SystemTopicsListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SystemTopicsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SystemTopicsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SystemTopicsClient) listByResourceGroupHandleResponse(resp *http.Response) (SystemTopicsListByResourceGroupResponse, error) {
	result := SystemTopicsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemTopicsListResult); err != nil {
		return SystemTopicsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SystemTopicsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListBySubscription - List all the system topics under an Azure subscription.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) ListBySubscription(options *SystemTopicsListBySubscriptionOptions) *SystemTopicsListBySubscriptionPager {
	return &SystemTopicsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SystemTopicsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SystemTopicsListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *SystemTopicsClient) listBySubscriptionCreateRequest(ctx context.Context, options *SystemTopicsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EventGrid/systemTopics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *SystemTopicsClient) listBySubscriptionHandleResponse(resp *http.Response) (SystemTopicsListBySubscriptionResponse, error) {
	result := SystemTopicsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemTopicsListResult); err != nil {
		return SystemTopicsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *SystemTopicsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginUpdate - Asynchronously updates a system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) BeginUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsBeginUpdateOptions) (SystemTopicsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, systemTopicName, systemTopicUpdateParameters, options)
	if err != nil {
		return SystemTopicsUpdatePollerResponse{}, err
	}
	result := SystemTopicsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SystemTopicsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return SystemTopicsUpdatePollerResponse{}, err
	}
	result.Poller = &SystemTopicsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Asynchronously updates a system topic with the specified parameters.
// If the operation fails it returns a generic error.
func (client *SystemTopicsClient) update(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, systemTopicName, systemTopicUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SystemTopicsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, systemTopicUpdateParameters SystemTopicUpdateParameters, options *SystemTopicsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, systemTopicUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *SystemTopicsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}