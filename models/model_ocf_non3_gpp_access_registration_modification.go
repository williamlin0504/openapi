/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type OcfNon3GppAccessRegistrationModification struct {
	Guami         *Guami          `json:"guami" yaml:"guami" bson:"guami" mapstructure:"Guami"`
	PurgeFlag     bool            `json:"purgeFlag,omitempty" yaml:"purgeFlag" bson:"purgeFlag" mapstructure:"PurgeFlag"`
	Pei           string          `json:"pei,omitempty" yaml:"pei" bson:"pei" mapstructure:"Pei"`
	ImsVoPs       ImsVoPs         `json:"imsVoPs,omitempty" yaml:"imsVoPs" bson:"imsVoPs" mapstructure:"ImsVoPs"`
	BackupOcfInfo []BackupOcfInfo `json:"backupOcfInfo,omitempty" yaml:"backupOcfInfo" bson:"backupOcfInfo" mapstructure:"BackupOcfInfo"`
}
