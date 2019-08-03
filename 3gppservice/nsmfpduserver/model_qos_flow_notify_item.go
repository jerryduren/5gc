/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package nsmfpduserver

type QosFlowNotifyItem struct {

	Qfi *map[string]interface{} `json:"qfi"`

	NotificationCause *NotificationCause `json:"notificationCause"`
}