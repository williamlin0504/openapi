/*
 * Nocf_Communication
 *
 * OCF Communication Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type OcfEventArea struct {
	PresenceInfo *PresenceInfo `json:"presenceInfo,omitempty"`
	LadnInfo     *LadnInfo     `json:"ladnInfo,omitempty"`
}
