# \SMContextsCollectionApi

All URIs are relative to *https://{apiRoot}/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSmContexts**](SMContextsCollectionApi.md#PostSmContexts) | **Post** /sm-contexts | Create SM Context


# **PostSmContexts**
> SmContextCreatedData PostSmContexts(ctx, body)
Create SM Context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Body**](Body.md)| representation of the SM context to be created in the SMF | 

### Return type

[**SmContextCreatedData**](SmContextCreatedData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: multipart/related
 - **Accept**: application/json, multipart/related

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

