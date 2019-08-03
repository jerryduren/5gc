# \SMContextsCollectionApi

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSmContexts**](SMContextsCollectionApi.md#PostSmContexts) | **Post** /sm-contexts | Create SM Context


# **PostSmContexts**
> SmContextCreatedData PostSmContexts(ctx, optional)
Create SM Context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostSmContextsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostSmContextsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject** | [**optional.Interface of InlineObject**](InlineObject.md)|  | 

### Return type

[**SmContextCreatedData**](SmContextCreatedData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/related
 - **Accept**: application/json, multipart/related, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

