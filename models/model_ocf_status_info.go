/*
 * Nocf_Communication
 *
 * OCF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type OcfStatusInfo struct {
	GuamiList        []Guami      `json:"guamiList,omitempty"`
	StatusChange     StatusChange `json:"statusChange"`
	TargetOcfRemoval string       `json:"targetOcfRemoval,omitempty"`
	TargetOcfFailure string       `json:"targetOcfFailure,omitempty"`
}
