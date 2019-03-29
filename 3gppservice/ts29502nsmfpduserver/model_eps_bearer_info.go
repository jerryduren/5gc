/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ts29502nsmfpduserver

type EpsBearerInfo struct {

	Ebi int32 `json:"ebi"`

	PgwS8uFteid *map[string]interface{} `json:"pgwS8uFteid"`

	BearerLevelQoS *map[string]interface{} `json:"bearerLevelQoS"`
}
