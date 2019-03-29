# NfService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstanceId** | **string** |  | 
**ServiceName** | [**ServiceName**](ServiceName.md) |  | 
**Versions** | [**[]NfServiceVersion**](NFServiceVersion.md) |  | 
**Scheme** | [**UriScheme**](UriScheme.md) |  | 
**NfServiceStatus** | [**NfServiceStatus**](NFServiceStatus.md) |  | 
**Fqdn** | **string** |  | [optional] 
**InterPlmnFqdn** | **string** |  | [optional] 
**IpEndPoints** | [**[]IpEndPoint**](IpEndPoint.md) |  | [optional] 
**ApiPrefix** | **string** |  | [optional] 
**DefaultNotificationSubscriptions** | [**[]DefaultNotificationSubscription**](DefaultNotificationSubscription.md) |  | [optional] 
**AllowedPlmns** | [**[]PlmnId**](PlmnId.md) |  | [optional] 
**AllowedNfTypes** | [**[]NfType**](NFType.md) |  | [optional] 
**AllowedNfDomains** | **[]string** |  | [optional] 
**AllowedNssais** | [**[]Snssai**](Snssai.md) |  | [optional] 
**Priority** | **int32** |  | [optional] 
**Capacity** | **int32** |  | [optional] 
**Load** | **int32** |  | [optional] 
**RecoveryTime** | [**time.Time**](time.Time.md) |  | [optional] 
**ChfServiceInfo** | [**ChfServiceInfo**](ChfServiceInfo.md) |  | [optional] 
**SupportedFeatures** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


