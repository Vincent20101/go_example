/* Adapted from: nnrf-api version December, 2018
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha2

// UDM information configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable UDM
//
//	 Data model:
//	 TS29510-g30 6.1.6.2.7 Type: UdmInfo, Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby UDM
type UdmInfo struct {
	// Identity of the UDM group that is served by the UDM instance.
	// If not provided, the UDM instance does not pertain to any UDM group
	// Optional
	GroupId string `mapstructure:"groupId" json:"groupId,omitempty"`
	// List of ranges of SUPIs whose profile data is available in the UDM instance. Either the start and end attributes, or the pattern attribute, shall be present.
	// Optional
	SupiRanges []SupiRange `mapstructure:"supiRanges" json:"supiRanges,omitempty"`
	// List of ranges of GPSIs whose profile data is available in the UDM instance. Either the start and end attributes, or the pattern attribute, shall be present.
	// Optional
	GpsiRanges []IdentityRange `mapstructure:"gpsiRanges" json:"gpsiRanges,omitempty"`
	// List of ranges of external groups whose profile data is available in the UDM instance. Either the start and end attributes, or the pattern attribute, shall be present.
	// Optional
	ExternalGroupIdentifiersRanges []IdentityRange `mapstructure:"externalGroupIdentifiersRanges" json:"externalGroupIdentifiersRanges,omitempty"`
	// List of Routing Indicator information that allows to route network signalling with SUCI (see 3GPP 23.003 [12]) to the UDM instance.
	// If not provided, the UDM can serve any Routing Indicator.
	// Pattern: '^[0-9]{1,4}$'
	// Optional
	RoutingIndicators []string `mapstructure:"routingIndicators" json:"routingIndicators,omitempty"`
}

// UDM nf profile from NRF or local static configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable UDM
//
//	 Data model:
//	   Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby UDM
type UdmNfProfile struct {
	// UDM's Profile
	// Mandatory
	NfProfile NfProfile `mapstructure:"nfProfile" json:"nfProfile"`
	// This IE is used to select UDM
	// Mandatory
	UdmInfo UdmInfo `mapstructure:"udmInfo" json:"udmInfo"`
}

// Deprecated
type UdmRuleDesc string

// Deprecated
const (
	UdmRule_Compare_SUPI       UdmRuleDesc = "SUPI"       //match SUPI range only
	UdmRule_Compare_DNN        UdmRuleDesc = "DNN"        //match DNN only
	UdmRule_Compare_SNSSAI     UdmRuleDesc = "SNSSAI"     //match SNSSAI only
	UdmRule_Compare_DNN_SNSSAI UdmRuleDesc = "DNN_SNSSAI" //match both DNN and SNSSAI
)

// Deprecated
type UdmSelectionPolicy struct {
	// List of PolicyMatchingRules
	SelectMode    PolicySelectMode     //policy selection mode
	MatchingRules []PolicyMatchingRule `mapstructure:"matchingRules" json:"matchingRules"`
}

// Sdm subscription static configuration
//
//	 Purpose:
//	   Is used to indicate when a subscribe information expires.
//
//	 Data model:
//	   Refer to the description for each attribute below.
//		  If neither ImplicitUnsubscribe nor Expires is set. Use smf hardcode ImplicitUnsubscribe = true.
//		  If ImplicitUnsubscribe and Expires are both set. Use ImplicitUnsubscribe = true.
//
//	 Usage:
//		  Smf global configuration
type SdmSubscriptionInfo struct {
	// Feature flag to enable  procedure SMF->(UDM-SDM) subscription in session setup call call flow
	// Default value "false"
	// Optional
	DisableSdmSubscription bool `mapstructure:"disableSdmSubscription" json:"disableSdmSubscription,omitempty"`
	// When set to true, this means that the subscription is terminated by UDM when the last PDU session of such SMF is deregistered for a given SUPI.
	// Default value: false, means doesn't include this IE to UDM in subscription
	// Optional
	ImplicitUnsubscribe bool `mapstructure:"implicitUnsubscribe" json:"implicitUnsubscribe,omitempty"`
	// Indicates the point in time at which the subscription expires.
	// Default value: 0, means doesn't include this IE to UDM in subscription
	// Unit: second
	// Optional
	Expires uint64 `mapstructure:"expires" json:"expires,omitempty"`
	// Indicates whether procedure SMF->(UDM-SDM) is asynchronous.\n
	// This field would work when DisableSdmSubscription is false.\n
	// Default value: false
	// Optional
	AsyncSubscription bool `mapstructure:"asyncSubscription" json:"asyncSubscription,omitempty"`
}

// Uecm registration static configuration
//
//	 Purpose:
//	   Is used to indicate whether to diable Uecm registration.
//
//	 Data model:
//	   Refer to the description for each attribute below.
//
//	 Usage:
//		  Smf global configuration.
type UecmRegistrationInfo struct {
	// Feature flag to enable  procedure SMF->(UDM-Uecm) registration in session establishment call flow.\n
	// Default value "false"
	// Optional
	DisableUecmRegistration bool `mapstructure:"disableUecmRegistration" json:"disableUecmRegistration,omitempty"`
}