/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

type EpsBearerInfo struct {
	Ebi int32 `json:"ebi" xml:"ebi"`
	PgwS8uFteid string `json:"pgwS8uFteid" xml:"pgwS8uFteid"`
	BearerLevelQoS string `json:"bearerLevelQoS" xml:"bearerLevelQoS"`
}
