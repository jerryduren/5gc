/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.PreR15.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Amf3GppAccessRegistration struct {

	AmfId string `json:"amfId"`

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	PurgeFlag bool `json:"purgeFlag,omitempty"`

	Pei string `json:"pei,omitempty"`

	ImsVoPS *ImsVoPs `json:"imsVoPS,omitempty"`

	DeregCallbackUri string `json:"deregCallbackUri"`

	PcscfRestorationCallbackUri string `json:"pcscfRestorationCallbackUri,omitempty"`
}