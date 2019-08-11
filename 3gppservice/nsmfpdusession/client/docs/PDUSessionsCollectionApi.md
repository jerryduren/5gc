# \PDUSessionsCollectionApi

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPduSessions**](PDUSessionsCollectionApi.md#PostPduSessions) | **Post** /pdu-sessions | Create


# **PostPduSessions**
> PduSessionCreatedData PostPduSessions(ctx, pduSessionCreateData)
Create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pduSessionCreateData** | [**PduSessionCreateData**](PduSessionCreateData.md)| representation of the PDU session to be created in the H-SMF | 

### Return type

[**PduSessionCreatedData**](PduSessionCreatedData.md)

### Authorization

[oAuth2ClientCredentials](../README.md#oAuth2ClientCredentials)

### HTTP request headers

 - **Content-Type**: application/json, multipart/related
 - **Accept**: application/json, multipart/related

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
