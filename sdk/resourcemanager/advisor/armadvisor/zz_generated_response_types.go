//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

import "net/http"

// ConfigurationsCreateInResourceGroupResponse contains the response from method Configurations.CreateInResourceGroup.
type ConfigurationsCreateInResourceGroupResponse struct {
	ConfigurationsCreateInResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsCreateInResourceGroupResult contains the result from method Configurations.CreateInResourceGroup.
type ConfigurationsCreateInResourceGroupResult struct {
	ConfigData
}

// ConfigurationsCreateInSubscriptionResponse contains the response from method Configurations.CreateInSubscription.
type ConfigurationsCreateInSubscriptionResponse struct {
	ConfigurationsCreateInSubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsCreateInSubscriptionResult contains the result from method Configurations.CreateInSubscription.
type ConfigurationsCreateInSubscriptionResult struct {
	ConfigData
}

// ConfigurationsListByResourceGroupResponse contains the response from method Configurations.ListByResourceGroup.
type ConfigurationsListByResourceGroupResponse struct {
	ConfigurationsListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsListByResourceGroupResult contains the result from method Configurations.ListByResourceGroup.
type ConfigurationsListByResourceGroupResult struct {
	ConfigurationListResult
}

// ConfigurationsListBySubscriptionResponse contains the response from method Configurations.ListBySubscription.
type ConfigurationsListBySubscriptionResponse struct {
	ConfigurationsListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ConfigurationsListBySubscriptionResult contains the result from method Configurations.ListBySubscription.
type ConfigurationsListBySubscriptionResult struct {
	ConfigurationListResult
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationEntityListResult
}

// RecommendationMetadataGetResponse contains the response from method RecommendationMetadata.Get.
type RecommendationMetadataGetResponse struct {
	RecommendationMetadataGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecommendationMetadataGetResult contains the result from method RecommendationMetadata.Get.
type RecommendationMetadataGetResult struct {
	MetadataEntity
}

// RecommendationMetadataListResponse contains the response from method RecommendationMetadata.List.
type RecommendationMetadataListResponse struct {
	RecommendationMetadataListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecommendationMetadataListResult contains the result from method RecommendationMetadata.List.
type RecommendationMetadataListResult struct {
	MetadataEntityListResult
}

// RecommendationsGenerateResponse contains the response from method Recommendations.Generate.
type RecommendationsGenerateResponse struct {
	RecommendationsGenerateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecommendationsGenerateResult contains the result from method Recommendations.Generate.
type RecommendationsGenerateResult struct {
	// Location contains the information returned from the Location header response.
	Location *string

	// RetryAfter contains the information returned from the Retry-After header response.
	RetryAfter *string
}

// RecommendationsGetGenerateStatusResponse contains the response from method Recommendations.GetGenerateStatus.
type RecommendationsGetGenerateStatusResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecommendationsGetResponse contains the response from method Recommendations.Get.
type RecommendationsGetResponse struct {
	RecommendationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecommendationsGetResult contains the result from method Recommendations.Get.
type RecommendationsGetResult struct {
	ResourceRecommendationBase
}

// RecommendationsListResponse contains the response from method Recommendations.List.
type RecommendationsListResponse struct {
	RecommendationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RecommendationsListResult contains the result from method Recommendations.List.
type RecommendationsListResult struct {
	ResourceRecommendationBaseListResult
}

// SuppressionsCreateResponse contains the response from method Suppressions.Create.
type SuppressionsCreateResponse struct {
	SuppressionsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SuppressionsCreateResult contains the result from method Suppressions.Create.
type SuppressionsCreateResult struct {
	SuppressionContract
}

// SuppressionsDeleteResponse contains the response from method Suppressions.Delete.
type SuppressionsDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SuppressionsGetResponse contains the response from method Suppressions.Get.
type SuppressionsGetResponse struct {
	SuppressionsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SuppressionsGetResult contains the result from method Suppressions.Get.
type SuppressionsGetResult struct {
	SuppressionContract
}

// SuppressionsListResponse contains the response from method Suppressions.List.
type SuppressionsListResponse struct {
	SuppressionsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SuppressionsListResult contains the result from method Suppressions.List.
type SuppressionsListResult struct {
	SuppressionContractListResult
}
