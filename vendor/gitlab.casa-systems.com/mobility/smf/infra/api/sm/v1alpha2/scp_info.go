/* Adapted from: nnrf-api version December, 2018
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha2

// SCP information configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable SCP
//
//	 Data model:
//	   Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby SCP
type ScpConfig struct {
	// Provide the address information of scp
	// Mandatory
	Scps []ScpAddress `mapstructure:"scps" json:"scps"`
	// Indicate that SMF will not use SCP for some interfaces
	// Default value nil
	// Optional
	ExemptInterfaces []NfInterface `mapstructure:"exemptInterfaces" json:"exemptInterfaces,omitempty"`
	// Can switch between model C and model D
	// Default value C
	// Optional
	Model CommunicationModel `mapstructure:"model" json:"model,omitempty"`
	// SCP Health check info \n
	// Default value: nil \n
	// Optional \n
	HealthCheck *ScpHealthCheck `mapstructure:"healthCheck" json:"healthCheck,omitempty"`
}

// SCP address configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable SCP
//
//	 Data model:
//	   Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby SCP
type ScpAddress struct {
	// Provide the address information of scp
	// - FQDN
	// - IPV4 or IPV6
	// - SRV record, "IsSRV:true",  format: _Service._Proto.Name TTL Class SRV Priority Weight Port Target
	// Mandatory
	Fqdn string `mapstructure:"fqdn" json:"fqdn"`
	// Port in the range of 0-65535, The value shall indicate the port number for HTTP or HTTPS respectively
	// Default value 80
	// Optional
	Port int32 `mapstructure:"port" json:"port,omitempty"`
	// Priority in the range of 0-65535, lower values indicate a higher priority.
	// Default value 0
	// Optional
	Priority int32 `mapstructure:"priority" json:"priority,omitempty"`
	// Optional deployment specific string used to construct the apiRoot of the next hop SCP, as described in clause 6.10 of 3GPP TS 29.500 [4].
	// Default value ""
	// Optional
	ScpPrefix string `mapstructure:"scpPrefix" json:"scpPrefix,omitempty"`
	// Indicate if the fqdn is SRV fqdn
	// - "true" means SRV record
	// - "false" means not SRV record
	// Default value "false"
	// Optional
	IsSRV bool `mapstructure:"isSRV" json:"isSRV,omitempty"`
}

// ScpHealthCheck
//
//	Purpose:
//	  Configured SCP health check info
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under ScpConfig
type ScpHealthCheck struct {
	// Health check Interval \n
	// unit: millisecond \n
	// Range: 0-604800000 \n
	// Granularity: 100 \n
	// Default value: 30000 \n
	// Optional
	Interval uint32 `mapstructure:"interval" json:"interval,omitempty"`
	// Number of consecutive success responses from a SCP in blacklist, after which SMF adds it to whitelist \n
	// Range: 1-20 \n
	// Default value: 3 \n
	// Optional
	ConsecSucc uint8 `mapstructure:"consecSucc" json:"consecSucc,omitempty"`
	// Number of consecutive failure responses from SCP in whitelist, after which SMF adds it to blacklist \n
	// Range: 1-20 \n
	// Default value: 3 \n
	// Optional
	ConsecFail uint8 `mapstructure:"consecFail" json:"consecFail,omitempty"`
	// Path of resource URI (excluding ApiRoot) which will be used by SMF to send HTTP OPTIONS to a specific resource \n
	// Should start with a "/" character \n
	// Default value: /v1/status \n
	// Optional \n
	ResourcePath string `mapstructure:"resourcePath" json:"resourcePath,omitempty"`
	// Whether to enable health check \n
	// Default: false \n
	// Optional
	Enable bool `mapstructure:"enable" json:"enable,omitempty"`
}
