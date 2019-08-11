# \IndividualSMContextApi

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleaseSmContext**](IndividualSMContextApi.md#ReleaseSmContext) | **Post** /sm-contexts/{smContextRef}/release | Release SM Context
[**RetrieveSmContext**](IndividualSMContextApi.md#RetrieveSmContext) | **Post** /sm-contexts/{smContextRef}/retrieve | Retrieve SM Context
[**UpdateSmContext**](IndividualSMContextApi.md#UpdateSmContext) | **Post** /sm-contexts/{smContextRef}/modify | Update SM Context


# **ReleaseSmContext**
> ReleaseSmContext(ctx, smContextRef, optional)
Release SM Context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smContextRef** | **string**| SM context reference | 
 **optional** | ***ReleaseSmContextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReleaseSmContextOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smContextReleaseData** | [**optional.Interface of SmContextReleaseData**](SmContextReleaseData.md)| representation of the data to be sent to the SMF when releasing the SM context | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json, multipart/related
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveSmContext**
> SmContextRetrievedData RetrieveSmContext(ctx, smContextRef, optional)
Retrieve SM Context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smContextRef** | **string**| SM context reference | 
 **optional** | ***RetrieveSmContextOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RetrieveSmContextOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smContextRetrieveData** | [**optional.Interface of SmContextRetrieveData**](SmContextRetrieveData.md)| parameters used to retrieve the SM context | 

### Return type

[**SmContextRetrievedData**](SmContextRetrievedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSmContext**
> SmContextUpdatedData UpdateSmContext(ctx, smContextUpdateData, smContextRef)
Update SM Context

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **smContextUpdateData** | [**SmContextUpdateData**](SmContextUpdateData.md)| representation of the updates to apply to the SM context | 
  **smContextRef** | **string**| SM context reference | 

### Return type

[**SmContextUpdatedData**](SmContextUpdatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json, multipart/related
 - **Accept**: application/json, multipart/related

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

