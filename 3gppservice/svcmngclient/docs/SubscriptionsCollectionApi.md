# \SubscriptionsCollectionApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubscription**](SubscriptionsCollectionApi.md#CreateSubscription) | **Post** /subscriptions | Create a new subscription


# **CreateSubscription**
> SubscriptionData CreateSubscription(ctx, subscriptionData)
Create a new subscription

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **subscriptionData** | [**SubscriptionData**](SubscriptionData.md)|  | 

### Return type

[**SubscriptionData**](SubscriptionData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

