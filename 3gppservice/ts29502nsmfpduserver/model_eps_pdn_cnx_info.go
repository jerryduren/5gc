/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ts29502nsmfpduserver

type EpsPdnCnxInfo struct {

	PgwS8cFteid *map[string]interface{} `json:"pgwS8cFteid"`

	PgwNodeName *map[string]interface{} `json:"pgwNodeName,omitempty"`
}
