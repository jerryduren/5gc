# Go API client for ts29571commondata

Common Data Types for Service Based Interfaces 

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build date: 2019-03-23T14:48:57.020+08:00[Asia/Shanghai]
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./ts29571commondata"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------


## Documentation For Models

 - [AccessType](docs/AccessType.md)
 - [AccessTypeRm](docs/AccessTypeRm.md)
 - [AdditionalQosFlowInfo](docs/AdditionalQosFlowInfo.md)
 - [Ambr](docs/Ambr.md)
 - [AmbrRm](docs/AmbrRm.md)
 - [Area](docs/Area.md)
 - [Arp](docs/Arp.md)
 - [ArpRm](docs/ArpRm.md)
 - [Atom](docs/Atom.md)
 - [BackupAmfInfo](docs/BackupAmfInfo.md)
 - [ChangeItem](docs/ChangeItem.md)
 - [ChangeType](docs/ChangeType.md)
 - [ChargingId](docs/ChargingId.md)
 - [Cnf](docs/Cnf.md)
 - [CnfUnit](docs/CnfUnit.md)
 - [ComplexQuery](docs/ComplexQuery.md)
 - [CoreNetworkType](docs/CoreNetworkType.md)
 - [CoreNetworkTypeRm](docs/CoreNetworkTypeRm.md)
 - [DnaiChangeType](docs/DnaiChangeType.md)
 - [DnaiChangeTypeRm](docs/DnaiChangeTypeRm.md)
 - [Dnf](docs/Dnf.md)
 - [DnfUnit](docs/DnfUnit.md)
 - [Dynamic5Qi](docs/Dynamic5Qi.md)
 - [Ecgi](docs/Ecgi.md)
 - [EcgiRm](docs/EcgiRm.md)
 - [EutraLocation](docs/EutraLocation.md)
 - [EutraLocationRm](docs/EutraLocationRm.md)
 - [GNbId](docs/GNbId.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [Guami](docs/Guami.md)
 - [GuamiRm](docs/GuamiRm.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [Link](docs/Link.md)
 - [LinkRm](docs/LinkRm.md)
 - [LinksValueSchema](docs/LinksValueSchema.md)
 - [Model5GMmCause](docs/Model5GMmCause.md)
 - [N3gaLocation](docs/N3gaLocation.md)
 - [Ncgi](docs/Ncgi.md)
 - [NcgiRm](docs/NcgiRm.md)
 - [NetworkId](docs/NetworkId.md)
 - [NgApCause](docs/NgApCause.md)
 - [NonDynamic5Qi](docs/NonDynamic5Qi.md)
 - [NotificationControl](docs/NotificationControl.md)
 - [NotificationControlRm](docs/NotificationControlRm.md)
 - [NotifyItem](docs/NotifyItem.md)
 - [NrLocation](docs/NrLocation.md)
 - [NrLocationRm](docs/NrLocationRm.md)
 - [OdbData](docs/OdbData.md)
 - [OdbPacketServices](docs/OdbPacketServices.md)
 - [PatchItem](docs/PatchItem.md)
 - [PatchOperation](docs/PatchOperation.md)
 - [PduSessionType](docs/PduSessionType.md)
 - [PduSessionTypeRm](docs/PduSessionTypeRm.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdRm](docs/PlmnIdRm.md)
 - [PreemptionCapability](docs/PreemptionCapability.md)
 - [PreemptionCapabilityRm](docs/PreemptionCapabilityRm.md)
 - [PreemptionVulnerability](docs/PreemptionVulnerability.md)
 - [PreemptionVulnerabilityRm](docs/PreemptionVulnerabilityRm.md)
 - [PresenceInfo](docs/PresenceInfo.md)
 - [PresenceInfoRm](docs/PresenceInfoRm.md)
 - [PresenceState](docs/PresenceState.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [QosFlowUsageReport](docs/QosFlowUsageReport.md)
 - [QosResourceType](docs/QosResourceType.md)
 - [QosResourceTypeRm](docs/QosResourceTypeRm.md)
 - [RatType](docs/RatType.md)
 - [RatTypeRm](docs/RatTypeRm.md)
 - [RatingGroup](docs/RatingGroup.md)
 - [RefToBinaryData](docs/RefToBinaryData.md)
 - [RefToBinaryDataRm](docs/RefToBinaryDataRm.md)
 - [ReflectiveQoSAttribute](docs/ReflectiveQoSAttribute.md)
 - [ReflectiveQoSAttributeRm](docs/ReflectiveQoSAttributeRm.md)
 - [RestrictionType](docs/RestrictionType.md)
 - [RestrictionTypeRm](docs/RestrictionTypeRm.md)
 - [RoamingOdb](docs/RoamingOdb.md)
 - [RouteInformation](docs/RouteInformation.md)
 - [RouteToLocation](docs/RouteToLocation.md)
 - [SecondaryRatUsageReport](docs/SecondaryRatUsageReport.md)
 - [SelfLink](docs/SelfLink.md)
 - [ServiceAreaRestriction](docs/ServiceAreaRestriction.md)
 - [ServiceId](docs/ServiceId.md)
 - [Snssai](docs/Snssai.md)
 - [SscMode](docs/SscMode.md)
 - [SscModeRm](docs/SscModeRm.md)
 - [SubscribedDefaultQos](docs/SubscribedDefaultQos.md)
 - [Tai](docs/Tai.md)
 - [TaiRm](docs/TaiRm.md)
 - [TraceData](docs/TraceData.md)
 - [TraceDepth](docs/TraceDepth.md)
 - [TraceDepthRm](docs/TraceDepthRm.md)
 - [UpConfidentiality](docs/UpConfidentiality.md)
 - [UpConfidentialityRm](docs/UpConfidentialityRm.md)
 - [UpIntegrity](docs/UpIntegrity.md)
 - [UpIntegrityRm](docs/UpIntegrityRm.md)
 - [UpSecurity](docs/UpSecurity.md)
 - [UpSecurityRm](docs/UpSecurityRm.md)
 - [UriScheme](docs/UriScheme.md)
 - [UserLocation](docs/UserLocation.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

