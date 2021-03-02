/*
 * Nocf_EventExposure
 *
 * OCF Event Exposure Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type OcfEventState struct {
	Active         bool  `json:"active" bson:"active" `
	RemainReports  int32 `json:"remainReports,omitempty" bson:"remainReports" `
	RemainDuration int32 `json:"remainDuration,omitempty" bson:"remainDuration" `
}
