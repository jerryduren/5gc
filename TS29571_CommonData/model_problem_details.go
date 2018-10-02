/*
 * Common Data Types
 *
 * Common Data Types
 *
 * API version: 1.R15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package TS29571_CommonData

type ProblemDetails struct {
	Type string `json:"type,omitempty" xml:"type"`
	Title string `json:"title,omitempty" xml:"title"`
	Status int32 `json:"status,omitempty" xml:"status"`
	Instance string `json:"instance,omitempty" xml:"instance"`
	Cause string `json:"cause,omitempty" xml:"cause"`
	InvalidParams []InvalidParam `json:"invalidParams,omitempty" xml:"invalidParams"`
}
