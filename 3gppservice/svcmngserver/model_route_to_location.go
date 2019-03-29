/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package svcmngserver

type RouteToLocation struct {

	Dnai string `json:"dnai"`

	RouteInfo *RouteInformation `json:"routeInfo,omitempty"`

	RouteProfId string `json:"routeProfId,omitempty"`
}
