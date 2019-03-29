# \SubscriptionIDDocumentApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RemoveSubscription**](SubscriptionIDDocumentApi.md#RemoveSubscription) | **Delete** /subscriptions/{subscriptionID} | Deletes a subscription
[**UpdateSubscription**](SubscriptionIDDocumentApi.md#UpdateSubscription) | **Patch** /subscriptions/{subscriptionID} | Updates a subscription


# **RemoveSubscription**
> RemoveSubscription(ctx, subscriptionID)
Deletes a subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionID** | **string**| Unique ID of the subscription to remove | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSubscription**
> SubscriptionData UpdateSubscription(ctx, patchItem, subscriptionID)
Updates a subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patchItem** | [**[]PatchItem**](array.md)|  | 
  **subscriptionID** | **string**| Unique ID of the subscription to update | 

### Return type

[**SubscriptionData**](SubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

