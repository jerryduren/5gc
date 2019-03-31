/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package svcmngserver

import (
	"time"
)

type NfService struct {

	ServiceInstanceId string `json:"serviceInstanceId"`

	ServiceName *ServiceName `json:"serviceName"`

	Versions []NfServiceVersion `json:"versions"`

	Scheme *UriScheme `json:"scheme"`

	NfServiceStatus *NfServiceStatus `json:"nfServiceStatus"`

	Fqdn string `json:"fqdn,omitempty"`

	InterPlmnFqdn string `json:"interPlmnFqdn,omitempty"`

	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`

	ApiPrefix string `json:"apiPrefix,omitempty"`

	DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`

	AllowedPlmns []PlmnId `json:"allowedPlmns,omitempty"`

	AllowedNfTypes []NfType `json:"allowedNfTypes,omitempty"`

	AllowedNfDomains []string `json:"allowedNfDomains,omitempty"`

	AllowedNssais []Snssai `json:"allowedNssais,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	Capacity int32 `json:"capacity,omitempty"`

	Load int32 `json:"load,omitempty"`

	RecoveryTime time.Time `json:"recoveryTime,omitempty"`

	ChfServiceInfo *ChfServiceInfo `json:"chfServiceInfo,omitempty"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}