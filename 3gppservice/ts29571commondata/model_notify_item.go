/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ts29571commondata

type NotifyItem struct {
	ResourceId string `json:"resourceId,omitempty" xml:"resourceId"`
	Changes []ChangeItem `json:"changes,omitempty" xml:"changes"`
}
