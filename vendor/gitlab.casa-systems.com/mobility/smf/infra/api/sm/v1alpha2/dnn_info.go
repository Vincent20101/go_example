package v1alpha2

type PolicyFunctionSelection string

const (
	POLICY_IF_DIAMETER PolicyFunctionSelection = "GX_DIAMETER"
	POLICY_IF_SBA      PolicyFunctionSelection = "NPCF_SBA"
)

type ChargingFunctionSelection string

const (
	CHARGING_ONLINE_IF_GY              ChargingFunctionSelection = "GY_DIAMETER"
	CHARGING_OFFLINE_IF_RF             ChargingFunctionSelection = "RF_DIAMETER"
	CHARGING_OFFLINE_IF_CDR_ASN1       ChargingFunctionSelection = "CDR_ASN1"
	CHARGING_OFFLINE_IF_CDR_CSV        ChargingFunctionSelection = "CDR_CSV"
	CHARGING_OFFLINE_IF_RADIUS         ChargingFunctionSelection = "RADIUS"
	CHARGING_IF_SBA_CONVERGED_CHARGING ChargingFunctionSelection = "CONVERGED_CHARGING"
	CHARGING_IF_SBA_OFFLINE_ONLY       ChargingFunctionSelection = "OFFLINE_ONLY"
)

type PolicyIfMediationCfg struct {
	// Specify if 4G calls should use diameter-GX or 5G SBA
	// default is GX_DIAMETER
	// optional
	Policy4G PolicyFunctionSelection `mapstructure:"policy4G" json:"policy4G,omitempty"`
	// Specify if 5G calls should use diameter-GX or 5G SBA
	// default is NPCF_SBA
	// optional
	Policy5G PolicyFunctionSelection `mapstructure:"policy5G" json:"policy5G,omitempty"`
}

type ChargingIfMediationCfg struct {
	// Specify if 4G Online calls should use "GY_DIAMETER" / "CONVERGED_CHARGING"
	// default is "GY_DIAMETER"
	// optional
	Online4G ChargingFunctionSelection `mapstructure:"online4G" json:"online4G,omitempty"`
	// Specify if 4G Offline calls should use "CDR_ASN1"/ "CDR_CSV" / "RF_DIAMETER" / "CONVERGED_CHARGING" / "OFFLINE_ONLY" / "RADIUS"
	// default is "CDR_ASN1"
	// optional
	Offline4G ChargingFunctionSelection `mapstructure:"offline4G" json:"offline4G,omitempty"`
	// Specify if 5G Online calls should use "GY_DIAMETER" / "CONVERGED_CHARGING"
	// default is "CONVERGED_CHARGING"
	// optional
	Online5G ChargingFunctionSelection `mapstructure:"online5G" json:"online5G,omitempty"`
	// Specify if 5G Offline calls should use "CDR_ASN1"/ "CDR_CSV" / "RF_DIAMETER" / "CONVERGED_CHARGING" / "OFFLINE_ONLY" / "RADIUS"
	// default is "CONVERGED_CHARGING"
	// optional
	Offline5G ChargingFunctionSelection `mapstructure:"offline5G" json:"offline5G,omitempty"`
}

type PccIfMediationCfg struct {
	// Policy interface config for 4G and 5G
	// optional
	PolicyIfCfg *PolicyIfMediationCfg `mapstructure:"policyIfCfg" json:"policyIfCfg,omitempty"`
	// Charging interface config
	// optional
	ChargingIfCfg *ChargingIfMediationCfg `mapstructure:"chargingIfCfg" json:"chargingIfCfg,omitempty"`
}

// UserPlane Security Policy for default dnn configuration
//
//	Purpose:
//	  Defines User Plane Security policy to be applied during PDU session establishment
type UpSecurity struct {
	// This IE shall indicate whether UP integrity protection is required, preferred or not needed for all the traffic on the PDU Session.\n
	// -"REQUIRED"\n
	// -"PREFERRED"\n
	// -"NOT_NEEDED"\n
	// Default value: "NOT_NEEDED"\n
	// Mandatory
	UpIntegr string `mapstructure:"upIntegr" json:"upIntegr"`
	// This IE shall indicate whether UP confidentiality protection is required, preferred or not needed for all the traffic on the PDU Session.\n
	// -"REQUIRED"\n
	// -"PREFERRED"\n
	// -"NOT_NEEDED"\n
	// Default value: "NOT_NEEDED"\n
	// Mandatory
	UpConfid string `mapstructure:"upConfid" json:"upConfid"`
}

// Default PduSession Types
//
//	Purpose:
//	  Default/allowed session types for a data network
//
//	Data model:
//	 Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/29503-g50.zip
//	 section 6.1.6.2.11 Table 6.1.6.2.11-1: PduSessionTypes
type PduSessionTypes struct {
	// Default session type
	// Value can be "IPV4" "IPV6" "IPV4V6" "UNSTRUCTURED" "ETHERNET" \n
	// Mandatory
	DefaultSessionType string `mapstructure:"defaultSessionType" json:"defaultSessionType"`
	// Additional session types allowed for the data network
	// Value can be "IPV4" "IPV6" "IPV4V6" "UNSTRUCTURED" "ETHERNET" \n
	// The default value is nil
	// Optional
	AllowedSessionTypes []string `mapstructure:"allowedSessionTypes" json:"allowedSessionTypes,omitempty"`
}

// Session and Service Continuity configuration
//
//	Purpose:
//	  This is local SscModes configuration to be used by SMF when ssc mode not returned by UDM in Get
//	  Nudm_SubscriberDataManagement operation from UDM
//
//	Data model
//	 Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/29503-g50.zip
//	 section 6.1.6.2.12 Table 6.1.6.2.12-1: SscModes
type SscModes struct {
	// Default SSC mode \n
	// Value can be "SSC_MODE_1" "SSC_MODE_2" "SSC_MODE_3" \n
	// Mandatory
	DefaultSscMode string `mapstructure:"defaultSscMode" json:"defaultSscMode"`
	// Additional SSC modes allowed for the data network \n
	// Value can be "SSC_MODE_1" "SSC_MODE_2" "SSC_MODE_3" \n
	// default value is nil. \n
	// Optional
	AllowedSscModes []string `mapstructure:"allowedSscModes" json:"allowedSscModes,omitempty"`
}

// Default Dnn configuration
//
//	Purpose:
//	  This is default dnn configuration to be used by SMF when DnnConfiguration is not returned by UDM in Get
//	  Nudm_SubscriberDataManagement operation from UDM
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.503/29503-g50.zip
//	  section Table 6.1.6.2.9-1: Definition of type DnnConfiguration for details
//
//	Usage:
//	  Used as optional parameter "DefaultDnnConfigurations" in DnnProfile
type DnnConfiguration struct {
	// Default/Allowed session types
	PduSessionTypes PduSessionTypes `mapstructure:"pduSessionTypes" json:"pduSessionTypes"`
	// Default/Allowed SSC modes
	SscModes SscModes `mapstructure:"sscModes" json:"sscModes"`
	// Indicates whether interworking with EPS is subscribed. If this attribute is absent it means not subscribed.
	// True: Subscribed;
	// False: Not subscribed;
	IwkEpsInd bool `mapstructure:"iwkEpsInd" json:"iwkEpsInd,omitempty"`
	// This field has been deleted in the new version of the protocol
	LadnIndicator bool `mapstructure:"ladnIndicator" json:"ladnIndicator,omitempty"`
	// 5G QoS parameters associated to the session for a data network
	Var5gQosProfile *DefaultQos `mapstructure:"var5gQosProfile" json:"var5gQosProfile,omitempty"`
	// The maximum aggregated uplink and downlink bit rates to be shared across all Non-GBR QoS Flows in each PDU Session
	SessionAmbr *Ambr `mapstructure:"sessionAmbr" json:"sessionAmbr,omitempty"`
	// Subscribed charging characteristics data associated to the session for a data network.
	Var3gppChargingCharacteristics string `mapstructure:"var3gppChargingCharacteristics" json:"var3gppChargingCharacteristics,omitempty"`
	// Subscribed static IP address(es) of the IPv4 and/or IPv6 type
	StaticIpAddress []IpAddress `mapstructure:"staticIpAddress" json:"staticIpAddress,omitempty"`
	// When present, this IE shall indicate the security policy for integrity protection and encryption for the user plane.
	// Default value nil.\n
	// Optional
	UpSecurity *UpSecurity `mapstructure:"upSecurity" json:"upSecurity,omitempty"`
	// Default network slice for this DNN.\n
	// 4G session will use this slice for subsequent session establishment if uses static udm subscription data.\n
	// Either there is wildCardSlicePolicyRef or SlicePolicyRef map has an entry for this snssai
	// Default value nil.\n
	// optional
	DefaultSnssai *Snssai `mapstructure:"defaultSnssai" json:"defaultSnssai,omitempty"`
}

// Ladn policy configuration
//
//  Purpose:
//   Use LadnPolicy to determine whether SMF needs to release session when receiving AMF's LADN OUT event notification
//
//  Data model:
//    Please refer https://www.3gpp.org/ftp/Specs/archive/23_series/23.501/23501-g40.zip
//    section 5.6.5	Support for Local Area Data Network
//
//  Usage:
//    If SMF needs to release session when receiving AMF's Ladn out event notification, OutAreaDelete needs to be set to true, otherwise it is set to false
//

type LadnPolicy struct {
	// Indicates whether delete Out of LADN area or not.\n
	// If SMF needs to release session when receiving AMF's Ladn out event notification, OutAreaDelete needs to be set to true, otherwise it is set to false.\n
	// The value can be "true" or "false".\n
	// Default value "true".\n
	// Optional.
	OutAreaDelete bool `mapstructure:"outAreaDelete" json:"outAreaDelete,omitempty"`
}

// Emergency Dnn configuration Data
//
//	Purpose:
//	  Configure local dnn configuration information for emergency pdu sessions.
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  If udm returns that the dnn configuration is invalid and it is an emergency pdu session, the configuration will be used as the emergency dnn configuration.
type EmergencyDnnConfigData struct {
	// The UPF instance applied to the emergency session when the UPF discovery fails.\n
	// Default value "".\n
	// Mandatory.
	EmergencyUpf string `mapstructure:"emergencyUpf" json:"emergencyUpf"`
	// 5G QoS parameters associated to the emergency service session for a data network.\n
	// Default value {}.\n
	// Mandatory.
	EmergencyQosProfile DefaultQos `mapstructure:"emergencyQosProfile" json:"emergencyQosProfile"`
	// The maximum aggregated uplink and downlink bit rates for emergency service session.\n
	// Default value {}.\n
	// Mandatory.
	EmergencySessionAmbr Ambr `mapstructure:"emergencySessionAmbr" json:"emergencySessionAmbr"`
}

type InactivityActionType string

const (
	// For emergency session it should always be SESS_DELETE.\n
	// Release pdu session
	INACTIVITY_ACTION_SESS_DELETE InactivityActionType = "SESS_DELETE"
	// Set pdu session UP to DEACTIVATE,default can be UP_DEACTIVATE.
	INACTIVITY_ACTION_UP_DEACTIVATE InactivityActionType = "UP_DEACTIVATE"
)

// Indicate whether to modify or release PDU session when receiving sdm notification.
type UdmSubscriptionChangeAction string

const (
	SUBSCRIPTION_Change_CONTEXT_MODIFICATION UdmSubscriptionChangeAction = "SmContextModification"
	SUBSCRIPTION_Change_CONTEXT_RELEASE      UdmSubscriptionChangeAction = "SmContextRelease"
)

// TS 23.501 section 5.10.3	PDU Session User Plane Security
type SecurityEnforcementType string

const (
	// Indicates UP integrity/confidentiality protection is Required
	SECURITY_ENFORCEMENT_REQUIRED SecurityEnforcementType = "REQUIRED"
	// Indicates UP integrity/confidentiality protection is Preferred
	SECURITY_ENFORCEMENT_PREFERRED SecurityEnforcementType = "PREFERRED"
	// Indicates UP integrity/confidentiality protection is Not Needed
	SECURITY_ENFORCEMENT_NOT_NEEDED SecurityEnforcementType = "NOT_NEEDED"
)

// UserPlane Security Policy for slice policy configuration
//
//	Purpose:
//	  Defines User Plane Security policy to be applied during PDU session establishment
type UpSecurityPolicy struct {
	// Indicates whether UP integrity protection is Required, Preferred or Not Needed\n
	// -"REQUIRED"\n
	// -"PREFERRED"\n
	// -"NOT_NEEDED"\n
	// Default value: "NOT_NEEDED"\n
	// Mandatory
	IntegrityProtection SecurityEnforcementType `mapstructure:"integrityProtection" json:"integrityProtection"`
	// Indicates whether UP confidentiality protection is Required, Preferred or Not Needed\n
	// -"REQUIRED"\n
	// -"PREFERRED"\n
	// -"NOT_NEEDED"\n
	// Default value: "NOT_NEEDED"\n
	// Mandatory
	ConfidentialityProtection SecurityEnforcementType `mapstructure:"confidentialityProtection" json:"confidentialityProtection"`
	// In case the UP Security Policy for the PDU Session is determined to have Integrity Protection set to “Required”, The SMF may, based on local configuration, decide whether to accept or reject the PDU Session request based on the UE Integrity Protection Maximum Data Rate.\n
	// Default value: false\n
	// Mandatory
	MaxUeDataRateProtectReject bool `mapstructure:"maxUeDataRateProtectReject" json:"maxUeDataRateProtectReject"`
}

type DlDropThrhldFlagType string

const (
	// indicate the Downlink Packets field shall be present, default value is "".
	DOWNLINK_DATA_DROP_PACKETS DlDropThrhldFlagType = "DROP_PACKETS"
	// indicate the Number of Bytes of Downlink Data field shall be present, default value is "".
	DOWNLINK_DATA_DROP_BYTES DlDropThrhldFlagType = "DROP_BYTES"
)

// Dropped DL Traffic Threshold
//
//	Purpose:
//	  Defines dropped DL traffic volume thresholds.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.244/29244-g30.zip section 8.2.49 for details
//
//	Usage:
//	  Configured inside DnnProfile
type DlDropThrhld struct {
	// Indicate the Downlink Packets/Bytes field shall be present.\n
	// - "DROP_PACKETS": indicate the Downlink Packets field shall be present.\n
	// - "REJECT_QOS_PROFILE": indicate the Number of Bytes of Downlink Data field shall be present.\n
	// Default value "".\n
	// Mandatory
	Flag DlDropThrhldFlagType `mapstructure:"flag" json:"flag"`
	// The Number of Bytes of Downlink Data fields shall be Encoded as an Unsigned64 binary integer value. It shall contain the number of bytes of the downlink data or contain a number of downlink packets.\n
	// Mandatory
	Thrhld uint64 `mapstructure:"thrhld" json:"thrhld"`
}

type PathSwitchFailType string

const (
	//default can be SESS_DELETE.
	PATH_SWITCH_FAIL_SESS_DELETE   PathSwitchFailType = "SESS_DELETE"
	PATH_SWITCH_FAIL_UP_DEACTIVATE PathSwitchFailType = "UP_DEACTIVATE"
)

// Virtua Dnn Matching rule
//
//	Purpose:
//	  Matching rule for virtual DNN
//
//	Data model:
//	  Refer to the description for each attribute below
type VirtualDnnMatchingRule struct {
	// Range of SUPIs that can be served by virtual dnn \n
	// Optional
	SupiRange *Range `mapstructure:"supiRange" json:"supiRange,omitempty"`
	// The RAT Type that can be served by virtual dnn \n
	// the value can be "NR" "EUTRA" "WLAN" "VIRTUAL" \n
	// Optional
	RatType *RatType `mapstructure:"ratType" json:"ratType,omitempty"`
	// Serving core network operator PLMN identity \n
	// Optional
	ServingPlmn *PlmnId `mapstructure:"servingPlmn" json:"servingPlmn,omitempty"`
	// Reseller identity \n
	// Optional
	ResellerId string `mapstructure:"resellerId" json:"resellerId,omitempty"`
	// Is roaming or not \n
	// - true: roaming \n
	// - false: non-roaming \n
	// - nil: either \n
	// Default: nil means it can match either roaming or non-roaming case \n
	// Optional
	RoamingStatus *bool `mapstructure:"roamingStatus" json:"roamingStatus,omitempty"`
}

// Default VirtualDnn configuration
//
//	Purpose:
//	  This is default VirtualDnn configuration to be used by SMF to support virtual APN
type VirtualDnn struct {
	// Virtual Dnn name. \n
	// Mandatory
	Name string `mapstructure:"name" json:"name"`
	// Matching priority, \n
	// The value must Greater than 0. \n
	// The default value is 0. \n
	// Optional
	Priority uint32 `mapstructure:"priority" json:"priority,omitempty"`
	// Matching rules \n
	// Mandatory
	MatchingRule VirtualDnnMatchingRule `mapstructure:"matchingRule" json:"matchingRule"`
	// Allow original DNN or not. when configed true smf will allow the inclusion of the original DNN
	// (submitted by UE or selected based on the UDM subscription) in the procedures with other NFs\n
	// The value is true or false, default false \n
	// Optional
	AllowOriginalDnn bool `mapstructure:"allowOriginalDnn" json:"allowOriginalDnn,omitempty"`
}

// VsmfQosProfile
//
//	Purpose:
//	  Reject Or Override the higher level Qos Profile send by h-smf.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured as a map under DnnProfile.
type VsmfQosProfile struct {
	// When h-smf notify v-smf the qos change, v-smf acts based on the config defined.\n
	// - "OVERWRITE_QOS_PROFILE": Override the Qos Profile if received the higher level from SMF.\n
	// - "REJECT_QOS_PROFILE": Reject the Qos Profile if received the higher level from SMF.\n
	// Default value "".\n
	// Optional
	VsmfQosChangeAction VsmfQosChangeAction `mapstructure:"vsmfQosChangeAction" json:"vsmfQosChangeAction,omitempty"`
	// This IE shall contain the 5G QoS Identifier (5QI) of the QoS flow.\n
	// Mandatory
	Var5qi int32 `mapstructure:"var5qi" json:"var5qi"`
	// This IE shall indicate the QoS Characteristics for a standardized or pre-configured 5QI for downlink and uplink.\n
	// Optional
	NonDynamic5Qi *NonDynamic5Qi `mapstructure:"nonDynamic5Qi" json:"nonDynamic5Qi,omitempty"`
	// This IE shall indicate the QoS Characteristics for a Non-standardised or not pre-configured 5QI for downlink and uplink.\n
	// Optional
	Dynamic5Qi *Dynamic5Qi `mapstructure:"dynamic5Qi" json:"dynamic5Qi,omitempty"`
	// This IE shall contain the Allocation and Retention Priority (ARP) assigned to the QoS flow.\n
	// Optional
	Arp *Arp `mapstructure:"arp" json:"arp,omitempty"`
	// This IE shall contain the GBR QoS flow information.\n
	// Optional
	GbrQosFlowInfo *GbrQosFlowInformation `mapstructure:"gbrQosFlowInfo" json:"gbrQosFlowInfo,omitempty"`
}

// When h-smf notify v-smf the qos change, v-smf may overwrite or reject the qos profile.
type VsmfQosChangeAction string

const (
	// Override the Qos Profile, Default value is "".
	VSMF_QOS_CHANGE_ACTION_OVERWRITE VsmfQosChangeAction = "OVERWRITE_QOS_PROFILE"
	// Reject the Qos Profile, Default value is "".
	VSMF_QOS_CHANGE_ACTION_REJECT VsmfQosChangeAction = "REJECT_QOS_PROFILE"
)

// DnsServerAddress
//
//	Purpose:
//	  This is DNS Server Address configuration.
//
//	Data model:
//	  Refer to the description for each attribute below.
//	  Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g30.zip Table 5.2.2.
//
//	Usage:
//	  At least one attribute is set,the default network port 53.
type DnsServerAddress struct {
	// IPv4 address. \n
	// Pattern: '^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])$' \n
	// Example: '198.51.100.1' \n
	// The default network port 53,no need to specify port number.\n
	// Default value nil.\n
	// Optional.
	Ipv4Addresses []string `mapstructure:"ipv4Addresses" json:"ipv4Addresses,omitempty"`
	// IPv6 address. \n
	// Pattern: '^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))$' \n
	// And, \n
	// Pattern: '^((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))$' \n
	// Example: '2001:db8:85a3::8a2e:370:7334' \n
	// The default network port 53,no need to specify port number.\n
	// Default value nil.\n
	// Optional.
	Ipv6Addresses []string `mapstructure:"ipv6Addresses" json:"ipv6Addresses,omitempty"`
}

type MtuProfile struct {
	// Ipv4LinkMtu. \n
	// Maximum Transmission Unit for Ipv4 Pdu Session Type.\n
	// Default value 0.\n
	// Optional.
	Ipv4LinkMtu uint16 `mapstructure:"ipv4LinkMtu" json:"ipv4LinkMtu,omitempty"`
	// Ipv6LinkMtu. \n
	// Maximum Transmission Unit for Ipv4 Pdu Session Type.\n
	// Default value 0.\n
	// Optional.
	Ipv6LinkMtu uint16 `mapstructure:"ipv6LinkMtu" json:"ipv6LinkMtu,omitempty"`
	// NonIpLinkMtu. \n
	// Maximum Transmission Unit for Non-Ip Pdu Session Type.\n
	// Default value 0.\n
	// Optional.
	NonIpLinkMtu uint16 `mapstructure:"nonIpLinkMtu" json:"nonIpLinkMtu,omitempty"`
	// EthernetLinkMtu. \n
	// Maximum Transmission Unit for Ethernet Pdu Session Type.\n
	// Default value 0.\n
	// Optional.
	EthernetLinkMtu uint16 `mapstructure:"ethernetLinkMtu" json:"ethernetLinkMtu,omitempty"`
	// UnstructuredLinkMtu. \n
	// Maximum Transmission Unit for Unstructured Pdu Session Type.\n
	// Default value 0.\n
	// Optional.
	UnstructuredLinkMtu uint16 `mapstructure:"unstructuredLinkMtu" json:"unstructuredLinkMtu,omitempty"`
}

type AtsssCapability string

const (
	// allows MPTCP with any steering mode and ATSSS-LL with only Active-Standby steering mode.
	MPTCP_ATSSS_LL_WITH_ASMODE AtsssCapability = "MPTCP_ATSSS_LL_WITH_ASMODE"
	// allows ATSSS-LL with any steering mode.
	ATSSS_LL AtsssCapability = "ATSSS_LL"
	// allows both MPTCP and ATSSS-LL with any steering mode.
	MPTCP_ATSSS_LL AtsssCapability = "MPTCP_ATSSS_LL"
)

type AllowGbrFlowAccess string

const (
	// provides the QoS profile for the GBR QoS Flow to 3GPP access.
	ALLOW_GBR_FLOW_TO_3GPP = "3GPP_ACCESS"
	// provides the QoS profile for the GBR QoS Flow to Non-3GPP access.
	ALLOW_GBR_FLOW_TO_NON_3GPP = "NON_3GPP_ACCESS"
)

type ReleaseMaSessionAction string

const (
	// completely release the MA PDU Session.
	RELEASE_BOTH_ACCESS = "RELEASE_BOTH_ACCESS"
	// released the MA PDU Session over a single access.
	RELEASE_SINGLE_ACCESS = "RELEASE_SINGLE_ACCESS"
)

// AtsssConfiguration
//
//	Purpose:
//	  This is Atsss configuration.
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used as optional parameter "AtsssConfiguration" in DnnProfile
type AtsssConfiguration struct {
	// Smf support Atsss capability configuration. \n
	// Allow values: \n
	//  -MPTCP_ATSSS_LL_WITH_ASMODE: allows MPTCP with any steering mode and ATSSS-LL with only Active-Standby steering mode.\n
	//  -ATSSS_LL: allows ATSSS-LL with any steering mode.\n
	//  -MPTCP_ATSSS_LL: allows both MPTCP and ATSSS-LL with any steering mode.\n
	// Default value is "".\n
	// Optional.
	AtsssCapability AtsssCapability `mapstructure:"atsssCapability" json:"atsssCapability,omitempty"`
	// If the PCC rule allows a GBR QoS Flow in both accesses, the SMF decides to which access network to provide the QoS profile for the GBR QoS Flow. \n
	// Allow values: \n
	//  -3GPP_ACCESS: provides the QoS profile for the GBR QoS Flow to 3GPP access.\n
	//  -NON_3GPP_ACCESS: provides the QoS profile for the GBR QoS Flow to Non-3GPP access.\n
	// Default value is "".\n
	// Optional.
	AllowGbrFlowAccess AllowGbrFlowAccess `mapstructure:"allowGbrFlowAccess" json:"allowGbrFlowAccess,omitempty"`
	// defined how to release muti access pdu session. \n
	// if the AMF needs to release the MA PDU Session over a single access, SMF decides whether the MA PDU Session completely released or released over a single access based on its local policy.\n
	// Allow values: \n
	//  -RELEASE_BOTH_ACCESS: completely release the MA PDU Session.\n
	//  -RELEASE_SINGLE_ACCESS: released the MA PDU Session over a single access.\n
	// Default value is "".\n
	// Optional.
	ReleaseMaSessionAction ReleaseMaSessionAction `mapstructure:"releaseMaSessionAction" json:"releaseMaSessionAction,omitempty"`
}

// SaUwbPraConfig
//
//	Purpose:
//	  SMF VzW SA UWB PRA notification processing configuration
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as optional parameter "SaUwbPraConfig" in DnnProfile
type SaUwbPraConfig struct {
	// The PRA ID, e.g., "11000000".\n
	// Mandatory.
	PraId string `mapstructure:"praId" json:"praId"`
	// The segment number(n) to which element id belongs in ipv6 address, 1 <= n <=8.\n
	// Default value 6.\n
	// Optional.
	SegmentNum uint32 `mapstructure:"segmentNum" json:"segmentNum,omitempty"`
	// element id value for 4G\n
	// Default value: [406].\n
	// Optional.
	ElementIds []uint32 `mapstructure:"elementIds" json:"elementIds,omitempty"`
	// Stability Timer duration,value: 0, which means that the Stability Timer is not start\n
	// Unit: Seconds\n
	// Default value 0.\n
	// Optional.
	StabilityTimerDuration uint32 `mapstructure:"stabilityTimerDuration" json:"stabilityTimerDuration,omitempty"`
}

// SaUwbPraConfigInfo
//
//	Purpose:
//	  SMF VzW SA UWB Enhancement PRA notification processing configuration
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as optional parameter "SaUwbPraConfigInfo" in DnnProfile
type SaUwbPraConfigInfo struct {
	//The PRA Config map
	//PRA ID is the map key
	// Optional
	PraConfigs map[string]*PraConfig `mapstructure:"praConfigs" json:"praConfigs,omitempty"`
	// Stability Timer duration \n
	// Unit: Millisecond\n
	// Default value 0, means that the Stability Timer is disabled.\n
	// Optional.
	StabilityTimerDuration uint32 `mapstructure:"stabilityTimerDuration" json:"stabilityTimerDuration,omitempty"`
	// Flag to indicate if SMF uses any user data activity report by UPF to determine whether or not UE is moved to RRC active state
	// Default value false.\n
	// Optional.
	DisablePraReportRrcInactive bool `mapstructure:"disablePraReportRrcInactive" json:"disablePraReportRrcInactive,omitempty"`
	//Initial PRA status, SMF will suppress any initial PRA status if it matches this pre-configured default InitialPRAStatus value\n
	//Value can be "IN_AREA", "OUT_OF_AREA"
	//Default value "OUT_OF_AREA", means that SMF will suppress the report of inital PRA status "OUT_OF_AREA".\n
	// Optional.
	InitialPraStatus string `mapstructure:"initialPraStatus" json:"initialPraStatus,omitempty"`
	// Flag to indicate if SMF disable suppression of initial PRA report or not.
	// Default value false.\n
	// Optional.
	SuppressInitialPraReport bool `mapstructure:"suppressInitialPraReport" json:"suppressInitialPraReport,omitempty"`
}

type PraConfig struct {
	// The PRA ID, e.g., "11000000".\n
	// Mandatory.
	PraId string `mapstructure:"praId" json:"praId"`
	// The segment number(n) to which the network element id belongs in ipv6 address, 1 <= n <=8.\n
	// Default value 6.\n
	// Optional.
	SegmentNum uint32 `mapstructure:"segmentNum" json:"segmentNum,omitempty"`
	// Network element id value for 5G \n\
	// Value can be "406" "403" "405" \n
	// Default value: [406].\n
	// Optional.
	ElementIds []uint32 `mapstructure:"elementIds" json:"elementIds,omitempty"`
}

type AlwaysOnType string

const (
	// it is at dnn level,always tell ue to be always on
	ALWAYSON_TYPE_ALWAYS AlwaysOnType = "ALWAYS"
	//allow ue request alwaysOn
	ALWAYSON_TYPE_ALLOW AlwaysOnType = "ALLOW"
	//reject ue request alwaysOn
	ALWAYSON_TYPE_REJECT AlwaysOnType = "REJECT"
)

type ReanchorOption string

const (
	// re-anchoring is triggered by PRA change
	REANCHOR_OP_PRA ReanchorOption = "PRA"
	// re-anchoring is triggered by TAI change
	REANCHOR_OP_TAI ReanchorOption = "TAI"
)

// MecPraConfig
//
//	Purpose:
//	  define user plane re-anchoring trigger and/or a list of MEC PRA IDs.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used in DnnProfile to config MEC PRAs
type MecPraConfig struct {
	// List of MEC PRA IDs.\n
	// Optional
	PraIds []string `mapstructure:"praIds" json:"praIds,omitempty"`
	// Indicates whether MEC PRA state change or UE TAI change will trigger re-anchoring.\n
	// Allow values: \n
	//	-PRA: re-anchoring is triggered by PRA change.\n
	//	-TAI: re-anchoring is triggered by TAI change.\n
	// Default value: "TAI".\n
	// Optional
	ReanchorOption ReanchorOption `mapstructure:"reanchorOption" json:"reanchorOption,omitempty"`
}

type DnAaaListUsageStrategy string

const (
	// Override any UDM provided values with SMF configured values
	DNAAALIST_USAGE_STRATEGY_OVERRIDE DnAaaListUsageStrategy = "OVERRIDE"
	// Ignore DnAaaList attribute on N10 (for missed configuration) and do not use N40-Lite
	DNAAALIST_USAGE_STRATEGY_IGNORE DnAaaListUsageStrategy = "IGNORE"
	// Use UDM returned DN-AAA list (DnAaaList attribute on N10)
	DNAAALIST_USAGE_STRATEGY_USE DnAaaListUsageStrategy = "USE"
)

// EnterPrisePdnDnAaaConfig
//
//	Purpose:
//	  define Enterprise PDN DN-AAA configuration.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used to configure the NASREQ, S6b-DN-Lite and N40-Lite interfaces for communication between smf and DN-AAA
//	  in the enterprise PDN procedure
type EnterprisePdnDnAaaConfig struct {
	// S6b-Lite interface DnAaaList usage strategy. \n
	// Value: OVERRIDE, IGNORE, USE \n
	// Default value: "USE".\n
	// Optional
	S6bLiteStrategy DnAaaListUsageStrategy `mapstructure:"s6bLiteStrategy" json:"s6bLiteStrategy,omitempty"`
	// N40-Lite interface DnAaaList usage strategy. \n
	// Default value: "IGNORE".\n
	// Optional
	// Deprecated: Use N40LiteCfg instead.
	N40LiteStrategy DnAaaListUsageStrategy `mapstructure:"n40LiteStrategy" json:"n40LiteStrategy,omitempty"`
	// Deprecated: Use DnAaaServers instead
	// EnterPrise PDN DnAaaList Configuration
	// Default value nil.\n
	// optional
	DnAaaList []string `mapstructure:"dnAaaList" json:"dnAaaList,omitempty"`
	// the DN-AAA server realm
	// Optional : DnAaaRealm must be configured when EnterprisePdnDnAaaConfig is configured in DnnProfile.
	DnAaaRealm string `mapstructure:"dnAaaRealm" json:"dnAaaRealm,omitempty"`
	// DiamAppType to use for this EnterprisePdnDnAaaConfig.\n
	// - "S6bDnLite".\n
	// - "NASREQ".\n
	// Default value is "S6bDnLite".\n
	// Optional.
	DiamAppType string `mapstructure:"diamAppType" json:"diamAppType,omitempty"`
	// Default ip-pool name for this enterprise DNN, used when UE IP address is returned from DN-AAA, and ip-pool name is not returned.\n
	// Default value is "", which means not configured.\n
	// Optional.
	IpPoolName string `mapstructure:"ipPoolName" json:"ipPoolName,omitempty"`
	// N40Lite config for APN/DNN. \n
	// Optional
	N40LiteCfg *N40LiteCfg `mapstructure:"n40LiteCfg" json:"n40LiteCfg,omitempty"`
	// List of DN-AAA servers. \n
	// Optional
	DnAaaServers []DnAaaServer `mapstructure:"dnAaaServers" json:"dnAaaServers,omitempty"`
}

// DnAaaServer for List of DN-AAA servers.
type DnAaaServer struct {
	// The FQDN of EnterPrise PDN DnAaaServer \n
	// Mandatory.
	Fqdn string `mapstructure:"fqdn" json:"fqdn,omitempty"`
	// The Port of EnterPrise PDN DnAaaServer  \n
	// Only used for N40Lite \n
	// must config port value if n40LiteCfg.Enable is true \n
	// Optional.
	Port uint16 `mapstructure:"port" json:"port,omitempty"`
	// The Priority of EnterPrise PDN DnAaaServer \n
	// Range: 0 to 255. The smaller the value, the higher the priority \n
	// default value: 0
	// Optional.
	Priority uint8 `mapstructure:"priority" json:"priority,omitempty"`
}

type N40LiteCfg struct {
	// N40-Lite interface DnAaaList usage strategy.\n
	// Default value: "USE".\n
	// Allow value: "USE", "IGNORE", "OVERRIDE" \n
	// Optional
	DnAAASelectionStrategy DnAaaListUsageStrategy `mapstructure:"dnAAASelectionStrategy" json:"dnAAASelectionStrategy,omitempty"`
	// Allow smf interaction with mpn-proxy or not, if true means allow smf interaction with mpn-proxy via n40Lite interface \n
	// This value is true or false, Default false \n
	// Optinal
	EnableN40Lite bool `mapstructure:"enableN40Lite" json:"enableN40Lite,omitempty"`
	// The timeout value of the message transfer from smf to MPN-PROXY.\n
	// A configurable timeout value per APN/DNN for receiving HTTP/2 response on N40-Lite interface.\n
	// Unit: Millisecond \n
	// Allowed value: 300~15000 \n
	// Default: 7000 \n
	// Optional
	ReceiveTimeout uint32 `mapstructure:"receiveTimeout" json:"receiveTimeout,omitempty"`
	// The BatchInterval value to process next batch after a batch is processed.\n
	// Unit: Millisecond \n
	// Allowed value: 1000~300000 \n
	// Default: 1000 \n
	// Optional
	BatchInterval uint32 `mapstructure:"batchInterval" json:"batchInterval,omitempty"`
	// The BatchSize for sending the messages from buffer in batches with configurable number of messages.\n
	// Allowed value: 1~10000 \n
	// Default: 20 \n
	// Optional
	BatchSize uint32 `mapstructure:"batchSize" json:"batchSize,omitempty"`
}

// DnnProfile
//
//	Purpose:
//	  This is default DnnProfile configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured as a map under DnnProfileMap.
type DnnProfile struct {
	// LADN related policy.\n
	// If LadnPolicy is nil, the default is that smf needs to release session when it receives the LADN OUT event notification from AMF.\n
	// Default value nil.\n
	// optional
	LadnPolicy *LadnPolicy `mapstructure:"ladnPolicy" json:"ladnPolicy,omitempty"`
	// SA UWB PRA notification processing configuration.\n
	// Default value nil.\n
	// optional
	// Legacy configuration, deprecated
	SaUwbPraConfig *SaUwbPraConfig `mapstructure:"saUwbPraConfig" json:"saUwbPraConfig,omitempty"`
	// SA UWB PRA notification processing configuration for Pra enhancement.\n
	// default value is nil. \n means not support special pra
	// Optional
	SaUwbPraConfigInfo *SaUwbPraConfigInfo `mapstructure:"saUwbPraConfigInfo" json:"saUwbPraConfigInfo,omitempty"`
	// NSA UWB PRA ID.\n
	// Default value "" means SMF will not handle NSA PRA subscription and notification.\n
	// optional
	// Legacy configuration, deprecated
	NsaUwbPraId string `mapstructure:"nsaUwbPraId" json:"nsaUwbPraId,omitempty"`
	// NSA UWB PRA IDs.\n
	// Default value nil means SMF will not handle NSA PRA subscription and notification.\n
	// optional
	NsaUwbPraIds []string `mapstructure:"nsaUwbPraIds" json:"nsaUwbPraIds,omitempty"`
	// User plane re-anchoring processing configuration when PCF is not available.\n
	// Default value nil means SMF will not trigger user plane re-anchoring.\n
	// Optional
	MecPraConfig *MecPraConfig `mapstructure:"mecPraConfig" json:"mecPraConfig,omitempty"`
	// After this time expire, when N2 Ho complete, the resource of indirect tunnel will be release.\n
	// Unit: Seconds.\n
	// Default value 0.\n
	// Optional.
	N2HOIndirectTunnelRelTimeSeconds uint16 `mapstructure:"n2HOIndirectTunnelRelTimeSeconds" json:"n2HOIndirectTunnelRelTimeSeconds,omitempty"`
	// Indicates how long the SMF is willing to maintain the old PDU Session or keep the old ipv6 prefix in case of SSC Mode3.\n
	// Unit: Seconds
	// Value can be uint16 integer.\n
	// Default value 0.\n
	// Optional.
	OldPduSessOrPrefixLifetime uint16 `mapstructure:"oldPduSessOrPrefixLifetime" json:"oldPduSessOrPrefixLifetime,omitempty"`
	// When this timer is expired, smf release source Upf in N2/Xn handover relocation I-Upf procedure.\n
	// Unit: Seconds\n
	// Default value 0.\n
	// Optional.
	UpfRelTimerSeconds uint16 `mapstructure:"upfRelTimerSeconds" json:"upfRelTimerSeconds,omitempty"`
	// The T5to4ho timer value determines the number of seconds that the SMF waits for the Modify Bearer Request message during an EPC fallback.\n
	// Verizon proprietary timer: refer to form '5G_LTE_Interface_Timer_Inventory-phase 3-0.4.xlsm' T5to4ho timer.\n
	// Unit: Seconds.\n
	// Range: 1 s to 16 s .\n
	// Default value 5.\n
	// Optional.
	T5to4hoTimer uint16 `mapstructure:"t5to4hoTimer" json:"t5to4hoTimer,omitempty"`
	// Indicates if an immediate event report in the subscription response is requested.\n
	// Please refer: https://www.3gpp.org/ftp/Specs/archive/29_series/29.518/29518-g30.zip.\n
	// section Table 6.2.6.2.3-1: Definition of type AmfEvent.\n
	// The value can be "true" or "false".\n
	// Default value false.\n
	// Optional.
	ImmediateEventReport bool `mapstructure:"immediateEventReport" json:"immediateEventReport,omitempty"`
	// Indicates whether the SMF carries the User ID (IMSI, IMEI, MSISDN, NAI, GCI, GLI, EXTID, MAC, or EUI) and APN \n
	// information elements in the PFCP Session Establishment Request message, with the expected values of "IMSI", \n
	// "IMEI", "MSISDN", "NAI", "GCI", "GLI", "EXTID", "MAC", "EUI", and "APN". \n
	// For example, when it is required to carry the IMSI, IMEI, and APN information elements in a PFCP Session \n
	// Establishment Request message, the values should be "IMSI|IMEI|APN". \n
	// Optionally, the default value is "".
	UserIdApnMask string `mapstructure:"userIdApnMask" json:"userIdApnMask,omitempty"`
	// Enable/disable Pause Charging feature.\n
	// -true: enable\n
	// -false disable.\n
	// Default value false.\n
	// Optional.
	EnablePauseCharging bool `mapstructure:"enablePauseCharging" json:"enablePauseCharging,omitempty"`
	// Enable/disable sscmode2 delay feature.\n
	// -true: enable\n
	// -false: disable.\n
	// Default value false.\n
	// Optional.
	EnableSscmode2Delay bool `mapstructure:"enableSscmode2Delay" json:"enableSscmode2Delay,omitempty"`
	// Set DlDropThrhld to urr to Monitor the number of dropped packets/bytes in upf.
	// Optional.
	DlDropThrhld *DlDropThrhld `mapstructure:"dlDropThrhld" json:"dlDropThrhld,omitempty"`
	// Indicator for paging policy differentiation feature at dnn level
	// -"true" means enable
	// -"false" means disable
	// Default value "false"
	// optional
	PagingPolicyDiff bool `mapstructure:"pagingPolicyDifferentiation" json:"pagingPolicyDifferentiation,omitempty"`
	// Enable/disable 3GPP Ps Data Off feature at dnn level.\n
	// -true: enable\n
	// -false disable.\n
	// Default value false.\n
	// Optional.
	PsDataOffFeature bool `mapstructure:"psDataOffFeature" json:"psDataOffFeature,omitempty"`
	// To cleanup session after call done. Applicable to other Dnns for inactivity cleanup.\n
	// Value can be uint32 integer.\n
	// Default value 0.\n
	// Range: 0-4294967295.\n
	// The unit is seconds.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.244/29244-g30.zip Table 8.2.83.\n
	// Optional.
	InactivityTimer uint32 `mapstructure:"inactivityTimer" json:"inactivityTimer,omitempty"`
	// This config needed only for non-emergency session,For emergency session it should always be SESS_DELETE.\n
	// -"SESS_DELETE" indicated release pdu session.
	// -"UP_DEACTIVATE" indicated deactivate pdu session.
	// Default value "SESS_DELETE".\n
	// Optional.
	InactivityAction InactivityActionType `mapstructure:"inactivityAction" json:"inactivityAction,omitempty"`
	// Indicates the maximum life time of the session. A value of zero, means that this session has an unlimited number of seconds before termination.\n
	// Value can be uint32 integer.\n
	// Default value 0.\n
	// Range: 0-4294967295.\n
	// The unit is seconds.\n
	// Optional.
	SessionLifeTimeSec uint32 `mapstructure:"sessionLifeTimeSec" json:"sessionLifeTimeSec,omitempty"`
	//AlwaysOn with three values:  `ALWAYS`, `ALLOW`, `REJECT`. \n
	// - `ALWAYS`: mean it is at dnn level and to instruct UE to be always-on
	// - `ALLOW` : accept ue request to enable always-on
	// - `REJECT`: reject ue request to enable always-on
	// Default value is `ALLOW`.\n
	// Optional.
	AlwaysOn AlwaysOnType `mapstructure:"alwaysOn" json:"alwaysOn,omitempty"`
	// When udm notify smfsm the subscription data change, smf acts based on the config defined.\n
	// -"SmContextModification":modify the pdu session if received qos modification from udm.\n
	// -"SmContextRelease":release the pdu session if received qos modification from udm.\n
	// Default value "SmContextModification".\n
	// Optional.
	UdmSubscriptionChangeAction UdmSubscriptionChangeAction `mapstructure:"udmSubscriptionChangeAction" json:"udmSubscriptionChangeAction,omitempty"`
	// Applied to emergency pdu sessions.\n
	// If udm returns that the dnn configuration is invalid and it is an emergency pdu session, the configuration will be used as the emergency dnn configuration.\n
	// Default value nil.\n
	// Optional.
	EmergencyDnnConfig *EmergencyDnnConfigData `mapstructure:"emergencyDnnConfig"  json:"emergencyDnnConfig,omitempty"`
	// Wildcard slice policy, for default for all slice
	// Mandatory.
	WildCardSlicePolicyRef string `mapstructure:"wildCardSlicePolicyRef" json:"wildCardSlicePolicyRef"`
	// Sst-sd as key, as key to refer to SlicePolicyProfile which is configured in SlicePolicyMap
	// Optional.
	SlicePolicyRef map[string]string `mapstructure:"slicePolicyRef" json:"slicePolicyRef,omitempty"`
	// Refer to TS23502-g40 4.3.2.1: 4 The SMF may use DNN Selection Mode when deciding whether to retrieve the Session Management Subscription data \n
	// E.g. in case the (DNN, S-NSSAI) is not explicitly subscribed, the SMF may use local configuration instead of Session Management Subscription data. \n
	// if req.SelMode from AMF is UE_DNN_NOT_VERIFIED or NW_DNN_NOT_VERIFIED, get the data from DefaultDnnConfigurations rather than UDM \n
	// Mandatory.
	DefaultDnnConfigurations *DnnConfiguration `mapstructure:"defaultDnnConfigurations" json:"defaultDnnConfigurations,omitempty"`
	// Aaa profile reference by name
	// Optional.
	AaaProfileRef string `mapstructure:"aaaProfileRef" json:"aaaProfileRef,omitempty"`
	// Dhcp profile reference for this dnn. Serves as key for lookup in DhcpProfileMap in SmfSvcCfg
	// Optional.
	DhcpProfileRef string `mapstructure:"dhcpProfileRef" json:"dhcpProfileRef,omitempty"`
	// List of possible virutalDnn for this DNN
	// Optional.
	VirtualDnns []VirtualDnn `mapstructure:"virtualDnns" json:"virtualDnns,omitempty"`
	// If v-smf get a higher QoS level than local configuration, v-smf needs to be able to override/reject the QoS values.\n
	// The key is "Mcc-Mnc", different PLMN use different profile to overwrite.\n
	// Optional.
	VsmfQosProfiles map[string]VsmfQosProfile `mapstructure:"vsmfQosProfiles" json:"vsmfQosProfiles,omitempty"`
	// default qos profile reference, if there is not any "mcc-mnc" matched, use this as a key of VsmfQosProfiles.\n
	// Default value "".\n
	// Optional.
	VsmfDefQosProfileRef string `mapstructure:"vsmfDefQosProfileRef" json:"vsmfDefQosProfileRef,omitempty"`
	// The SMF can decide whether to release the PDU Session or to deactivate the UP connection of this PDU Session.\n
	// when a PDU Session is rejected by the Target NG-RAN,it will apply in the handover process and service request process.\n
	// -"SESS_DELETE" indicated release pdu session.
	// -"UP_DEACTIVATE" indicated deactivate pdu session.
	// Default value "SESS_DELETE".\n
	// Optional.
	PathSwitchFailAction PathSwitchFailType `mapstructure:"pathSwitchFailAction" json:"pathSwitchFailAction,omitempty"`
	// Indicates whether Extended Buffing applies in NIDD during SMF-NEF connection establishment.\n
	//  -"true" indicated enable Extended Buffing.
	//  -"false" indicated disable Extended Buffing.
	// Default value "false".\n
	// Optional.
	ExtBufSupport bool `mapstructure:"extendedBuffering" json:"extendedBuffering,omitempty"`
	// Indicates disable discover PCF, just use static Pcc.\n
	//  -"true" means enable just use static pcc, will not discover PCF.
	//  -"false" means disable just use static Pcc, will discover PCF.
	// Default value "false".\n
	// Optional.
	StaticPccOnly bool `mapstructure:"staticPccOnly" json:"staticPccOnly,omitempty"`
	// Indicates disable discover UDM, just use static local configuration .\n
	//  -"true" means enable just use static local configuration, will not discover UDM.
	//  -"false" means disable just use static local configuration, will discover UDM.
	// Default value "false".\n
	// Optional.
	StaticUdmOnly bool `mapstructure:"staticUdmOnly" json:"staticUdmOnly,omitempty"`
	// Indicates disable discover UPF, just use static UPFs .\n
	//  -"true" means just using static UPFs and will not discover UPF.
	//  -"false" means not just using static local configuration, will discover UPF.
	// Default value "false".\n
	// Optional.
	StaticUpfOnly bool `mapstructure:"staticUpfOnly" json:"staticUpfOnly,omitempty"`
	// This timer is used by source I-SMF during handover to release the SM Context of the PDU Session, if indirect forwarding tunnel(s) were previously established, otherwise, the Source I-SMF immediately releases the SM Context.\n
	// Unit: Seconds.\n
	// Default value 0.\n
	// Optional.
	SmContextRelTimerSeconds uint16 `mapstructure:"smContextRelTimerSeconds" json:"smContextRelTimerSeconds,omitempty"`
	// Indicates dns server address.\n
	// Supports multiple DNS server address.\n
	// Default value {}.\n
	// Optional.
	DnsServerAddress DnsServerAddress `mapstructure:"dnsServerAddress" json:"dnsServerAddress,omitempty"`
	// Indicates MTU configuration.\n
	// Supports pass MTU in nas in epco parameter during session establishment procedure.\n
	// Default value nil.\n
	// Optional.
	MtuProfile *MtuProfile `mapstructure:"mtuProfile" json:"mtuProfile,omitempty"`
	// CIoT(Cellular IoT) configuration. \n
	// Need to configure IoT related configuration If CIOT feature is enabled. \n
	// Default value is nil.\n
	// Optional.
	CIoTConfiguration *CIoTConfiguration `mapstructure:"cIoTConfiguration" json:"cIoTConfiguration,omitempty"`
	// Atsss configuration. \n
	// Need to configure Atsss related configuration If Atsss feature is enabled. \n
	// Default value is nil.\n
	// Optional.
	AtsssConfiguration *AtsssConfiguration `mapstructure:"atsssConfiguration" json:"atsssConfiguration,omitempty"`
	// Paging failure backoff timer in seconds. \n
	// Default value 5s \n
	// optional.
	PagingFailureBackoffTimer *uint8 `mapstructure:"pagingFailureBackoffTimer" json:"pagingFailureBackoffTimer,omitempty"`
	// Enterprise PDN DN-AAA configuration \n
	// Default value is nil.\n
	// Optional.
	EnterprisePdnDnAaaConfig *EnterprisePdnDnAaaConfig `mapstructure:"enterprisePdnDnAaaConfig" json:"enterprisePdnDnAaaConfig,omitempty"`
	// indicate Default assume positive profile
	// Default value default \n
	// Optional
	DefaultAssumePositiveProfile string `mapstructure:"defaultAssumePositiveProfile" json:"defaultAssumePositiveProfile,omitempty"`
	// Producer's locality used as node discover preferred-locality parameter which will overwrite SMF locality.\n
	// Key could be "amf"/"pcf"/"udm"/"chf". \n
	// Default value is nil.\n
	// Optional
	PreferredLocalityOverwrite map[string]string `mapstructure:"preferredLocalityOverwrite" json:"preferredLocalityOverwrite,omitempty"`
	// Policy and Charging Mediation Config
	// Optional
	PccIfMediationCfg *PccIfMediationCfg `mapstructure:"pccIfMediationCfg" json:"pccIfMediationCfg,omitempty"`
	// Indicates PGW behavior when PCRF don't included pcc rule under CCA-I.\n
	// -"true" means enable, PGW will reject ue attach if PCRF don't included pcc rule under CCA-I
	// -"false" means disable, PGW will allow ue attach if PCRF don't included pcc rule under CCA-I
	// Default value "false".\n
	// Optional.
	RejectNoRuleUeAttach bool `mapstructure:"rejectNoRuleUeAttach" json:"rejectNoRuleUeAttach,omitempty"`
	// If there is still no control plane signaling interaction after this duration, the session is considered stale and
	// a core network resource release procedure is triggered.
	// When configured to 0, this feature will be deactivated.
	// The unit is seconds, the range is 0 to 2147483647 and the default value is 0.
	// Optional
	SessionIdleTimeout uint32 `mapstructure:"sessionIdleTimeout" json:"sessionIdleTimeout,omitempty"`
	// Operator specific EPCO container at dnn level
	// Optional
	OperatorPCOContainerConfig *OperatorPCOContainerConfig `mapstructure:"operatorPCOContainerConfig" json:"operatorPCOContainerConfig,omitempty"`
	// Radius accounting profile reference by name
	// Optional.
	// Deprecated in the future
	RadiusProfileRef string `mapstructure:"radiusProfileRef" json:"radiusProfileRef,omitempty"`

	// Charging Data Report Profile reference by name
	// Value: the key of axyomServices.smfsmConfig.cdrProfileMap
	// Default value ""
	// Optional.
	CdrProfileRef string `mapstructure:"cdrProfileRef" json:"cdrProfileRef,omitempty"`

	// enable apn authorize with aaa server in s6b interface
	// Optional.
	AaaAuthorizeS2bRealm string `mapstructure:"aaaAuthorizeS2bRealm" json:"aaaAuthorizeS2bRealm,omitempty"`
	// Radius accounting profile reference by name
	// Optional.
	RadiusAcctProfileRef string `mapstructure:"radiusAcctProfileRef" json:"radiusAcctProfileRef,omitempty"`
	// Radius access profile reference by name
	// Optional.
	RadiusAuthProfileRef string `mapstructure:"radiusAuthProfileRef" json:"radiusAuthProfileRef,omitempty"`
	// Indicates whether SMF should continue the remaining procedures of Session Establishment\n
	// when procedure SMF->(UDM-SDM) subscription failed.\n
	// This field would work when DisableSdmSubscription and AsyncSubscription in smfsm/nudm are both false.\n
	// Default value "false".\n
	// Optional
	ContSessOnUdmSubFail bool `mapstructure:"contSessOnUdmSubFail" json:"contSessOnUdmSubFail,omitempty"`
	// Indicates whether SMF should continue the remaining procedures of Session Establishment\n
	// when procedure SMF->(UDM-UECM) registration failed.\n
	// This field would work when DisableUecmRegistration in smfsm/nudm is false.
	// Default value "false".\n
	// Optional
	ContSessOnUdmRegFail bool `mapstructure:"contSessOnUdmRegFail" json:"contSessOnUdmRegFail,omitempty"`
	//When present, this IE shall indicate Token Accumalation Interval for MBR(GBR type flow, or non-GBR flow when MBR is present).\n
	//Default value is 0, means do not support Token Accumalation Interval feature \n
	//Value Range: Interval: 1 - 1000 ms.
	//Optional.
	MbrTokenInterval uint32 `mapstructure:"mbrTokenInterval" json:"mbrTokenInterval,omitempty"`
	//When present, this IE shall indicate Token Accumalation Interval for AMBR(Non-GBR type flow).\n
	//Default value is 0, means do not support Token Accumalation Interval feature \n
	//Value Range: Interval: 1 - 1000 ms.
	//Optional.
	AmbrTokenInterval uint32 `mapstructure:"ambrTokenInterval" json:"ambrTokenInterval,omitempty"`
	// Configurable IP address stickiness time per DNN
	// unit: second, range 0-129600, default 0,0 means no ip address stickiness
	// Optional
	// Deprecated
	IpAddressStickiness uint32 `mapstructure:"ipAddressStickiness" json:"ipAddressStickiness,omitempty"`
	// To enable or disable N6 based forwarding for 5G VN Group
	// Default value false corresponds to N6 forwarding disabled for 5G VN Group
	// Optional
	VnGroupN6BasedForwarding bool `mapstructure:"vnGroupN6BasedForwarding" json:"vnGroupN6BasedForwarding,omitempty"`
	// Ethernet Pdu Session Configuration \n
	// Default value: nil \n
	// Optional
	EthernetPduConfig *EthernetPduConfig `mapstructure:"ethernetPduConfig" json:"ethernetPduConfig,omitempty"`
	// User-Name: If no username is available, a generic username, configurable on a per APN basis, shall be present.(refer to TS29061 16.4.1) \n
	// Default value: ""
	// Optional.
	RadiusDefaultUserName string `mapstructure:"radiusDefaultUserName" json:"radiusDefaultUserName,omitempty"`
}

// EthernetPduConfig
//
//	Purpose:
//	  This is Ethernet Pdu Session configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used to configure outermost S-TAG and/or C-TAGs for ethernet pdu sessions
type EthernetPduConfig struct {
	// Outermost S-TAG for ethernet pdu session \n
	// Service-VLAN tag containing the VID, PCP/DEI fields as defined in IEEE 802.1Q [17] and IETF RFC 7042 [18]. \n
	// S-TAG is encoded as a two-octet string in hexadecimal representation. Each character in the string shall take a value of "0" to "9" or "A" to "F" and shall represent 4 bits. The most significant character representing the PCP/DEI field shall appear first in the string, followed by character representing the 4 most significant bits of the VID field, and the character representing the 4 least significant bits of the VID field shall appear last in the string. \n
	// Default value "" corresponds to no STag configuration \n
	// Optional
	STag string `mapstructure:"sTag" json:"sTag,omitempty"`
}
