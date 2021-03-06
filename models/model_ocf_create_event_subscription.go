/*
 * Nocf_EventExposure
 *
 * OCF Event Exposure Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type OcfCreateEventSubscription struct {
	Subscription      *OcfEventSubscription `json:"subscription" bson:"subscription" `
	SupportedFeatures string                `json:"supportedFeatures,omitempty" bson:"supportedFeatures" `
}
