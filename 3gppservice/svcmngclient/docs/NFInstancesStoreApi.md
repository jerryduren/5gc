# \NFInstancesStoreApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNFInstances**](NFInstancesStoreApi.md#GetNFInstances) | **Get** /nf-instances | Retrieves a collection of NF Instances


# **GetNFInstances**
> InlineResponse200 GetNFInstances(ctx, optional)
Retrieves a collection of NF Instances

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetNFInstancesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetNFInstancesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nfType** | [**optional.Interface of NfType**](.md)| Type of NF | 
 **limit** | **optional.Int32**| How many items to return at one time | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/3gppHal+json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

