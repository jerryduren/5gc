# \IndividualPDUSessionHSMFApi

All URIs are relative to *https://{apiRoot}/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReleasePduSession**](IndividualPDUSessionHSMFApi.md#ReleasePduSession) | **Post** /pdu-sessions/{pduSessionRef}/release | Release
[**UpdatePduSession**](IndividualPDUSessionHSMFApi.md#UpdatePduSession) | **Post** /pdu-sessions/{pduSessionRef}/modify | Update (initiated by V-SMF)


# **ReleasePduSession**
> ReleasePduSession(ctx, pduSessionRef, optional)
Release

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pduSessionRef** | **string**| PDU session reference | 
 **optional** | ***ReleasePduSessionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReleasePduSessionOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **releaseData** | [**optional.Interface of ReleaseData**](ReleaseData.md)| representation of the data to be sent to H-SMF when releasing the PDU session | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePduSession**
> HsmfUpdatedData UpdatePduSession(ctx, hsmfUpdateData, pduSessionRef)
Update (initiated by V-SMF)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **hsmfUpdateData** | [**HsmfUpdateData**](HsmfUpdateData.md)| representation of the updates to apply to the PDU session | 
  **pduSessionRef** | **string**| PDU session reference | 

### Return type

[**HsmfUpdatedData**](HsmfUpdatedData.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, multipart/related
 - **Accept**: application/json, multipart/related

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

