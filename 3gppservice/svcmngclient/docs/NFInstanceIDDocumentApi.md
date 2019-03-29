# \NFInstanceIDDocumentApi

All URIs are relative to *https://example.com/nnrf-nfm/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeregisterNFInstance**](NFInstanceIDDocumentApi.md#DeregisterNFInstance) | **Delete** /nf-instances/{nfInstanceID} | Deregisters a given NF Instance
[**GetNFInstance**](NFInstanceIDDocumentApi.md#GetNFInstance) | **Get** /nf-instances/{nfInstanceID} | Read the profile of a given NF Instance
[**RegisterNFInstance**](NFInstanceIDDocumentApi.md#RegisterNFInstance) | **Put** /nf-instances/{nfInstanceID} | Register a new NF Instance
[**UpdateNFInstance**](NFInstanceIDDocumentApi.md#UpdateNFInstance) | **Patch** /nf-instances/{nfInstanceID} | Update NF Instance profile


# **DeregisterNFInstance**
> DeregisterNFInstance(ctx, nfInstanceID)
Deregisters a given NF Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nfInstanceID** | [**string**](.md)| Unique ID of the NF Instance to deregister | 

### Return type

 (empty response body)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNFInstance**
> NfProfile GetNFInstance(ctx, nfInstanceID)
Read the profile of a given NF Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nfInstanceID** | [**string**](.md)| Unique ID of the NF Instance | 

### Return type

[**NfProfile**](NFProfile.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterNFInstance**
> NfProfile RegisterNFInstance(ctx, nfProfile, nfInstanceID)
Register a new NF Instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nfProfile** | [**NfProfile**](NfProfile.md)|  | 
  **nfInstanceID** | [**string**](.md)| Unique ID of the NF Instance to register | 

### Return type

[**NfProfile**](NFProfile.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNFInstance**
> NfProfile UpdateNFInstance(ctx, patchItem, nfInstanceID)
Update NF Instance profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **patchItem** | [**[]PatchItem**](array.md)|  | 
  **nfInstanceID** | [**string**](.md)| Unique ID of the NF Instance to update | 

### Return type

[**NfProfile**](NFProfile.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json-patch+json
 - **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

