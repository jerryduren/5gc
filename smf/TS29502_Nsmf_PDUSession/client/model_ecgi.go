/*
 * SMF PDU Session
 *
 * SMF PDU Session Service
 *
 * API version: 1.PreR15.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

type Ecgi struct {
	PlmnId PlmnId `json:"plmnId" xml:"plmnId"`
	EutraCellId string `json:"eutraCellId" xml:"eutraCellId"`
}
