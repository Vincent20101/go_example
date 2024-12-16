package v1alpha2

// AssumePositiveProfile
//
//	Purpose:
//	  This is assume positive profile configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured as a map under ApProfileMap.
type AssumePositiveProfile struct {
	// Indicate whether this profile can be used
	// Optional
	Enable bool `mapstructure:"enable" json:"enable,omitempty"`
	// The number of times tempQuota is allowed. If 0 is entered, there is no limit. If the value is greater than 0, terminate the session after the number of consecutive tempQuotas exceeds.
	// Optional
	AllowCount uint32 `mapstructure:"allowCount" json:"allowCount,omitempty"`
	// The temporary quota to assign when in Assume Positive.
	// unit is byte.
	// Default value 209715200.\n
	// Optional
	ApTempQuota *RequestedUnit `mapstructure:"apTempQuota" json:"apTempQuota,omitempty"`
	// The temporary Validity Time to assign when in Assume Positive.
	// unit is second.
	// Default value 36000.\n
	// Optional
	ApTempValidityTime *uint64 `mapstructure:"apTempValidityTime" json:"apTempValidityTime,omitempty"`
	// The temporary volume quota threshold to assign when in Assume Positive.
	// unit is byte.
	// Default value is nil
	// Optional
	ApTempVolumeQuotaThresholdByte *int64 `mapstructure:"apTempVolumeQuotaThresholdByte" json:"apTempVolumeQuotaThresholdByte,omitempty"`
	// Assume Positive status (Whether the error code and error information are allowed to be processed by AP).Mandatory, when enable is true.
	// Key is HTTP Status Code and problemDetails cause.
	// Default value false.\n
	// Optional
	ApMappingMap map[int]map[string]bool `mapstructure:"apMappingMap" json:"apMappingMap,omitempty"`
}
