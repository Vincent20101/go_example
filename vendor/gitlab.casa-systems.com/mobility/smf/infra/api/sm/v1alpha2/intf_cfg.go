package v1alpha2

type TimerType string

const (
	// Default timer timeout
	TIMER_DEFAULT TimerType = "default"

	//namf timer type
	TIMER_N1N2MESSAGE_TRANSFER     TimerType = "n1n2message_transfer"
	TIMER_EBI_ASSIGNMENT           TimerType = "ebi_assignment"
	TIMER_EE_SUBSCRIBE             TimerType = "ee_subscribe"
	TIMER_EE_SUBSCRIBE_MODIFY      TimerType = "ee_subscribe_modify"
	TIMER_EE_UNSUBSCRIBE           TimerType = "ee_unsubscribe"
	TIMER_SM_CONTEXT_STATUS_NOTIFY TimerType = "sm_context_status_notify"
	TIMER_STATUS_CHANGE_SUBSCRIBE  TimerType = "status_change_subscribe"

	//npcf timer type
	TIMER_POLICY_CREATE TimerType = "policy_create"
	TIMER_POLICY_UPDATE TimerType = "policy_update"
	TIMER_POLICY_DELETE TimerType = "policy_delete"

	//nchf timer type
	TIMER_CHARGING_CREATE  TimerType = "charging_create"
	TIMER_CHARGING_UPDATE  TimerType = "charging_update"
	TIMER_CHARGING_RELEASE TimerType = "charging_release"

	//nnrf timer type
	TIMER_NF_STATUS_SUBSCRIBE        TimerType = "nf_status_subscribe"
	TIMER_NF_STATUS_SUBSCRIBE_UPDATE TimerType = "nf_status_subscribe_update"
	TIMER_NF_STATUS_UNSUBSCRIBE      TimerType = "nf_status_unsubscribe"
	TIMER_NF_PROFILE_RETRIEVAL       TimerType = "nf_profile_retrieval"
	TIMER_NF_DISCOVERY               TimerType = "nf_discovery"
	TIMER_ACCESS_TOKEN               TimerType = "access_token"

	//nudm timer type
	TIMER_SDM_GET         TimerType = "sdm_get"
	TIMER_SDM_SUBSCRIBE   TimerType = "sdm_subscribe"
	TIMER_SDM_UNSUBSCRIBE TimerType = "sdm_unsubscribe"
	TIMER_REGISTRATION    TimerType = "registration"
	TIMER_DEREGISTRATION  TimerType = "deregistration"

	// n1 timer type
	TIMER_T3591 TimerType = "t3591"
	TIMER_T3592 TimerType = "t3592"

	// n2 timer type
	TIMER_RES_SETUP_REQ TimerType = "n2_res_setup_req"
	TIMER_RES_MOD_REQ   TimerType = "n2_res_mod_req"
	TIMER_RES_REL_CMD   TimerType = "n2_res_rel_cmd"
	TIMER_5TO4_HO       TimerType = "n2_5to4_ho"
)

// Custom api version Configuration
//
//	Purpose:
//	  Configure the api full version of the NF service supported by smf
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under each interface configuration
type ApiVersion struct {
	// Defines the highest api full version supported by smf when selecting a service.\n
	// Default: "".\n
	// Optional.
	PreferredApiVersion string `mapstructure:"preferredApiVersion" json:"preferredApiVersion,omitempty"`
	// Defines all api full versions supported by smf when selecting a service.\n
	// Default: nil.\n
	// Optional.
	SupportedApiVersions []string `mapstructure:"supportedApiVersions" json:"supportedApiVersions,omitempty"`
}

// Timer Configuration
//
//	Purpose:
//	  Configured application timer parameters of smfsm
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under each interface configuration
type TimerConfig struct {
	// Interval or Timeout value that is used by smf when sending out a msg.\n
	// Unit: millisecond.\n
	// A default value will be set as following when it is 0 \n
	// - 16000ms for n1 and n2 interfaces \n
	// - 7000ms for gtpc/aaa-client grpc interfaces \n
	// - 3000ms for sbi and other grpc interfaces \n
	// - 10000ms for pcscf health check interval  \n
	// Optional.
	Interval uint16 `mapstructure:"interval" json:"interval,omitempty"`
	// Number of retries, excluding the original msg.\n
	// Default: 0. means no need to resend the msg.\n
	// Optional.
	Retries uint16 `mapstructure:"retries" json:"retries,omitempty"`
}

// Namf Interface Related Configuration
//
//	Purpose:
//	  Configured Namf interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type NamfCfg struct {
	// Timer configuration per Namf service operation.\n
	// Key can be:
	//  - "n1n2message_transfer"
	//	- "ebi_assignment"
	//	- "ee_subscribe"
	//	- "ee_subscribe_modify"
	//	- "ee_unsubscribe"
	//	- "sm_context_status_notify"
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: nil.\n
	// Optional.
	Timer map[TimerType]TimerConfig `mapstructure:"timer" json:"timer,omitempty"`
	// Default timer configuration for all Namf service operations.\n
	// Value configured in Timer takes higher precedence over this.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
	// Namf service version \n
	// Default value: nil \n
	// Optional.
	NamfSvcApiFullVersion *ApiVersion `mapstructure:"namfSvcApiFullVersion"    json:"namfSvcApiFullVersion,omitempty"`
	// SupportedFeatures configuration of Namf_Communication service api, referring to 3GPP TS 29.518\n
	// Default: nil \n
	// Optional
	CommSupportedFeatures *NamfCommSuppFeat `mapstructure:"commSupportedFeatures"    json:"commSupportedFeatures,omitempty"`
	// SupportedFeatures configuration of Namf_EventExposure service api, referring to 3GPP TS 29.518\n
	// Default: nil \n
	// Optional
	EventExposureSupportedFeatures *NamfEventExposureSuppFeat `mapstructure:"eventExposureSupportedFeatures"    json:"eventExposureSupportedFeatures,omitempty"`
	// Whether smf uses nfset information about this interface from nrf.\n
	// Default: false.\n
	// Optional.
	EnableNfSet bool `mapstructure:"enableNfSet" json:"enableNfSet,omitempty"`
}

// Namf Communication Service SupportedFeatures Related Configuration
//
//	Purpose:
//	  Configured Namf Communication service SupportedFeatures
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under Namf
type NamfCommSuppFeat struct {
	// Deployments Topologies with specific SMF Service Areas \n
	// Default: false \n
	// Optional \n
	DTSSA bool `mapstructure:"dTSSA"    json:"dTSSA,omitempty"`
	// Cellular IoT \n
	// Default: false \n
	// Optional \n
	CIOT bool `mapstructure:"cIOT"    json:"cIOT,omitempty"`
	// Enhanced LCS \n
	// Default: false \n
	// Optional \n
	ELCS bool `mapstructure:"eLCS"    json:"eLCS,omitempty"`
}

// Namf EventExposure Service SupportedFeatures Related Configuration
//
//	Purpose:
//	  Configured Namf EventExposure service SupportedFeatures
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under Namf
type NamfEventExposureSuppFeat struct {
	// Immediate Event Report in Subscription Creation Response \n
	// Default: false \n
	// Optional \n
	IERSR bool `mapstructure:"iERSR"    json:"iERSR,omitempty"`
}

// Npcf Interface Related Configuration
//
//	Purpose:
//	  Configured Npcf interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type NpcfCfg struct {
	// Timer configuration per Npcf service operation.\n
	// Key can be:
	//  - "policy_create"
	//	- "policy_update"
	//	- "policy_delete"
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries .\n
	// Optional.
	Timer map[TimerType]TimerConfig `mapstructure:"timer" json:"timer,omitempty"`
	// Default timer configuration for all Npcf service operations.\n
	// Value configured in Timer takes higher precedence over this.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: nil \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
	// Npcf service version \n
	// Default value: nil \n
	// Optional.
	NpcfSvcApiFullVersion *ApiVersion `mapstructure:"npcfSvcApiFullVersion"    json:"npcfSvcApiFullVersion,omitempty"`
	// Decide whether to include the DNN carrier identifier
	// DnnNIOnly is true: SMF send NI to PCF
	// DnnNIOnly is false: SMF send NI+OI to PCF
	// Default: false
	// Note: Pcms.Gx interface has similar config, which in GxConfig.GxApnNIOnly
	// Optional
	DnnNIOnly bool `mapstructure:"dnnNIOnly" json:"dnnNIOnly,omitempty"`
	// SupportedFeatures configuration of Npcf_SMPolicyControl service api, referring to 3GPP TS 29.512\n
	// Default: nil \n
	// Optional
	SmPolicySupportedFeatures *NpcfSmPolicySuppFeat `mapstructure:"smPolicySupportedFeatures"    json:"smPolicySupportedFeatures,omitempty"`
	// Whether smf uses nfset information about this interface from nrf.\n
	// Default: false.\n
	// Optional.
	EnableNfSet bool `mapstructure:"enableNfSet" json:"enableNfSet,omitempty"`
	// Timeout configuration for Npcf service operation in NSA call flows.\n
	// Value is automatically deduced by operator.\n
	// Unit: Millisecond \n
	// Value Range: 1300 - 8300 \n
	// Default: 3300 \n
	// Optional.
	InternalTimeoutToPcms uint16 `mapstructure:"internalTimeoutToPcms" json:"internalTimeoutToPcms,omitempty"`
	// Whether to enable sending UDR Group ID to PCF in httpheader
	// Default：false
	// Optional
	AddSbiHeaderUdrGroupId bool `mapstructure:"addSbiHeaderUdrGroupId" json:"addSbiHeaderUdrGroupId,omitempty"`
}

// Npcf SMPolicyControl Service SupportedFeatures Related Configuration
//
//	Purpose:
//	  Configured Npcf SMPolicyControl service SupportedFeatures
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under Npcf
type NpcfSmPolicySuppFeat struct {
	// Traffic Steering Control \n
	// Default: false \n
	// Optional \n
	TSC bool `mapstructure:"tSC"    json:"tSC,omitempty"`
	// 3GPP PS Data Off Status Change Reporting \n
	// Default: false \n
	// Optional \n
	PSDataOff bool `mapstructure:"pSDataOff"    json:"pSDataOff,omitempty"`
	// Application Detection and Control \n
	// Default: false \n
	// Optional \n
	ADC bool `mapstructure:"aDC"    json:"aDC,omitempty"`
	// Usage Monitoring Control \n
	// Default: false \n
	// Optional \n
	UMC bool `mapstructure:"uMC"    json:"uMC,omitempty"`
	// Access Network Information Reporting for 5GS \n
	// Default: false \n
	// Optional \n
	NetLoc bool `mapstructure:"netLoc"    json:"netLoc,omitempty"`
	// Detailed Release Cause Code Information from the Access Network \n
	// Default: false \n
	// Optional \n
	RANNASCause bool `mapstructure:"rANNASCause"    json:"rANNASCause,omitempty"`
	// Presence Reporting Area Change Reporting.\n
	// Default: false \n
	// Optional \n
	PRA bool `mapstructure:"pRA"    json:"pRA,omitempty"`
	// PCC Rule Versioning \n
	// Default: false \n
	// Optional \n
	RuleVersioning bool `mapstructure:"ruleVersioning"    json:"ruleVersioning,omitempty"`
	// Maximum Packet Loss Rate Value(s) for Uplink and/or Downlink Voice Service Data Flow(s) \n
	// Default: false \n
	// Optional \n
	RANSupportInfo bool `mapstructure:"rANSupportInfo"    json:"rANSupportInfo,omitempty"`
	// Access Type Conditioned Authorized Session AMBR \n
	// Default: false \n
	// Optional \n
	AccessTypeCondition bool `mapstructure:"accessTypeCondition"    json:"accessTypeCondition,omitempty"`
	// Multiple Ipv6 Address Prefixes Reporting \n
	// Default: false \n
	// Optional \n
	MultiIpv6AddrPrefix bool `mapstructure:"multiIpv6AddrPrefix"    json:"multiIpv6AddrPrefix,omitempty"`
	// Session Rule Error Handling \n
	// Default: false \n
	// Optional \n
	SessionRuleErrorHandling bool `mapstructure:"sessionRuleErrorHandling"    json:"sessionRuleErrorHandling,omitempty"`
	// Long Character Strings as Charging Identifiers \n
	// Default: false \n
	// Optional \n
	AFChargingIdentifier bool `mapstructure:"aFChargingIdentifier"    json:"aFChargingIdentifier,omitempty"`
	// Race Condition Handling \n
	// Default: false \n
	// Optional \n
	PendingTransaction bool `mapstructure:"pendingTransaction"    json:"pendingTransaction,omitempty"`
	// MAC addresses with a Specific Range in the Traffic Filter \n
	// Default: false \n
	// Optional \n
	MacAddressRange bool `mapstructure:"macAddressRange"    json:"macAddressRange,omitempty"`
	// QoS Monitoring \n
	// Default: false \n
	// Optional \n
	QosMonitoring bool `mapstructure:"qosMonitoring"    json:"qosMonitoring,omitempty"`
	// DN-AAA Authorization Data for Policy Control \n
	// Default: false \n
	// Optional \n
	DNAuthorization bool `mapstructure:"dNAuthorization"    json:"dNAuthorization,omitempty"`
	// PDU Session Release Cause \n
	// Default: false \n
	// Optional \n
	PDUSessionRelCause bool `mapstructure:"pDUSessionRelCause"    json:"pDUSessionRelCause,omitempty"`
	// Handling PDU Session Termination Functionality \n
	// Default: false \n
	// Optional \n
	RespBasedSessionRel bool `mapstructure:"respBasedSessionRel"    json:"respBasedSessionRel,omitempty"`
	// DNN Selection Mode\n
	// Default: false \n
	// Optional \n
	DNNSelectionMode bool `mapstructure:"dNNSelectionMode"    json:"dNNSelectionMode,omitempty"`
	// Error Report Handling of the Policy Decision and/or Condition Data which is not Referred by any PCC Rule or Session Rule \n
	// Default: false \n
	// Optional \n
	PolicyDecisionErrorHandling bool `mapstructure:"policyDecisionErrorHandling"    json:"policyDecisionErrorHandling,omitempty"`
}

// Nchf Interface Related Configuration
//
//	Purpose:
//	  Configured Nchf interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type NchfCfg struct {
	// Timer configuration per Nchf service operation.\n
	// Key can be:
	//  - "charging_create"
	//	- "charging_update"
	//	- "charging_release"
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: nil.\n
	// Optional.
	Timer map[TimerType]TimerConfig `mapstructure:"timer" json:"timer,omitempty"`
	// Default timer configuration for all Nchf service operations.\n
	// Value configured in Timer takes higher precedence over this.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries  \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
	// Nchf interface version \n
	//  -"v1"
	//  -"v2"
	//  -"v3"
	// Default: "v3"
	// Optional
	NchfApiVersion string `mapstructure:"nchfApiVersion"    json:"nchfApiVersion,omitempty"`
	// SMF shall support User Denied Time in FUA implementation.\n
	// This is used to wait for requesting quota when CHF denies a quota request.\n
	// Unit: Seconds \n
	// Default: 43200 \n
	// Optional
	UserDeniedTimeSec uint32 `mapstructure:"userDeniedTimeSec" json:"userDeniedTimeSec,omitempty"`
	// send to same chf fail req count \n
	// Optional
	ChfReqFailCount uint `mapstructure:"chfReqFailCount" json:"chfReqFailCount,omitempty"`
	// how long to retry to send req to chf \n
	// Uint: Seconds
	// Optional
	ChfReqRetryInterval uint `mapstructure:"chfReqRetryInterval" json:"chfReqRetryInterval,omitempty"`
	// The value indicates the number of UsedUnitContainers that may be included in a single Nchf message. \n
	// If this field is not configured or is zero, it would indicate that there is no limit.
	// Optional
	MaxUUCinChfMsg uint `mapstructure:"maxUUCinChfMsg" json:"maxUUCinChfMsg,omitempty"`
	// The value indicates the max number of UsedUnitContainers can be stored in a session, when enter AP state. \n
	// If this field is not configured or is zero, it would indicate that there is no limit.
	// Optional
	MaxUUCForSess uint `mapstructure:"maxUUCForSess" json:"maxUUCForSess,omitempty"`
	// The value indicates the max number of UsedUnitContainers can be stored in SMF, when enter AP state. \n
	// If this field is not configured or is zero, it would indicate that there is no limit.
	// Optional
	MaxUUCForAllSess uint `mapstructure:"maxUUCForAllSess" json:"maxUUCForAllSess,omitempty"`
	// whether include default triggers in charging create req. \n
	// Default value: false
	// Optional
	IncludeDefTrigInInitReq bool `mapstructure:"includeDefTrigInInitReq" json:"includeDefTrigInInitReq,omitempty"`
	// Whether to include OrigUserLocationinfo in charging req. \n
	// Default value: false \n
	// Optional
	IncludeOrigUserLocationInfo bool `mapstructure:"includeOrigUserLocationInfo" json:"includeOrigUserLocationInfo,omitempty"`
	// Nchf service version \n
	// Default value: nil \n
	// Optional.
	NchfSvcApiFullVersion *ApiVersion `mapstructure:"nchfSvcApiFullVersion" json:"nchfSvcApiFullVersion,omitempty"`
	// Decide whether to include the DNN carrier identifier
	// DnnNIOnly is true: SMF send NI to CHF
	// DnnNIOnly is false: SMF send NI+OI to CHF
	// Default: false
	// Optional
	DnnNIOnly bool `mapstructure:"dnnNIOnly" json:"dnnNIOnly,omitempty"`
	// Whether smf uses nfset information about this interface from nrf.\n
	// Default: false.\n
	// Optional.
	EnableNfSet bool `mapstructure:"enableNfSet" json:"enableNfSet,omitempty"`
	// Timeout configuration for Nchf service operation in NSA call flows.\n
	// Value is automatically deduced by operator.\n
	// Unit: Millisecond \n
	// Value Range: 1300 - 8300 \n
	// Default: 3300 \n
	// Optional.
	InternalTimeoutToPcms uint16 `mapstructure:"internalTimeoutToPcms" json:"internalTimeoutToPcms,omitempty"`
	// SupportedFeatures configuration of Nchf_ConvergedCharging service api, referring to 3GPP TS 32.291\n
	// Default: nil \n
	// Optional
	// ConvChargingSupportedFeatures    *NchfConvChargingSuppFeat    `mapstructure:"convChargingSupportedFeatures"    json:"convChargingSupportedFeatures,omitempty"`
	// SupportedFeatures configuration of Nchf_OfflineOnlyCharging service api, referring to 3GPP TS 32.291\n
	// Default: nil \n
	// Optional
	// OfflineOnlyChargingSupportedFeatures *NchfOfflineOnlyChargingSuppFeat `mapstructure:"offlineOnlyChargingSupportedFeatures"    json:"offlineOnlyChargingSupportedFeatures,omitempty"`
}

// Nchf ConvergedCharging Service SupportedFeatures Related Configuration
//
//  Purpose:
//    Configured Nchf ConvergedCharging service SupportedFeatures
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Configured under Nchf
//	Note: This Service has required supported feature for VzW Phase 3
//type NchfConvChargingSuppFeat struct {
//}

// Nchf OfflineOnlyCharging Service SupportedFeatures Related Configuration
//
//  Purpose:
//    Configured Nchf OfflineOnlyCharging service SupportedFeatures
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Configured under Nchf
//type NchfOfflineOnlyChargingSuppFeat struct {
//}

// Nnef Interface Related Configuration
//
//	Purpose:
//	  Configured Nnef interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type NnefCfg struct {
	// Default timer configuration for all N service operations.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries  \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
}

// Nnrf Interface Related Configuration
//
//	Purpose:
//	  Configured Nnrf interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type NnrfCfg struct {
	// Timer configuration per Nnrf service operation.\n
	// Key can be:
	//  - "nf_status_subscribe"
	//	- "nf_status_subscribe_update"
	//	- "nf_status_unsubscribe"
	//  - "nf_profile_retrieval"
	//	- "nf_discovery"
	//	- "access_token"
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: nil.\n
	// Optional.
	Timer map[TimerType]TimerConfig `mapstructure:"timer" json:"timer,omitempty"`
	// Default timer configuration for all Nnrf service operations.\n
	// Value configured in Timer takes higher precedence over this.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries  \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
	// Nnrf service version \n
	// Default value: nil \n
	// Optional.
	NnrfSvcApiFullVersion *ApiVersion `mapstructure:"nnrfSvcApiFullVersion" json:"nnrfSvcApiFullVersion,omitempty"`
	// SupportedFeatures configuration of Nnrf_NfManagement service api, referring to 3GPP TS 29.510\n
	// Default: nil \n
	// Optional
	// NfMgmtSupportedFeatures *NnrfNfmgmtSuppFeat `mapstructure:"nfMgmtSupportedFeatures"    json:"nfMgmtSupportedFeatures,omitempty"`
	// SupportedFeatures configuration of Nnrf_NfDiscovery service api, referring to 3GPP TS 29.510\n
	// Default: nil \n
	// Optional
	NfDiscSupportedFeatures *NnrfNfDiscSuppFeat `mapstructure:"nfDiscSupportedFeatures"    json:"nfDiscSupportedFeatures,omitempty"`
	// Health check configuration for each NRF service operation \n
	// Value Range: \n
	//  - Interval: 10000 - 300000 ms \n
	//  - ConsecSuccOrFail: 1 - 255 \n
	//  - Disable: false / true
	// Default: 90000 for interval /  2 for consecSuccOrFail / false for Disable \n
	// Optional
	HealthCheck *NrfHealthCheck `mapstructure:"healthCheck"    json:"healthCheck,omitempty"`
	// NF discovery optional query parameters that will be excluded from the request towards NRF \n
	// The map key could be "amf"/"pcf"/"udm"/"chf"/"pcscf"/"nef" \n
	// The val would be a group of discovery query parameters \n
	// Default: nil \n
	// optional
	ExcludedDiscParams map[string]*QueryParams `mapstructure:"excludedDiscParams" json:"excludedDiscParams,omitempty"`
	// NF discovery caching can be turned off \n
	// The map key could be "amf"/"chf"/"nef"/"pcf"/"udm"/"upf"/"pcscf" \n
	// Defaule: nil \n
	// Optional
	DisableNfCache map[string]bool `mapstructure:"disableNfCache" json:"disableNfCache,omitempty"`
	// Feature flag to enable SMF pass NI only part of DNN when sending discovery query for discovering PCF
	// DnnNIOnlyForPcfDiscovery is true: SMF pass NI only part of DNN when sending discovery query for discovering PCF
	// DnnNIOnlyForPcfDiscovery is false: SMF pass NI+OI when sending discovery query for discovering PCF
	// Default: false
	// Optional
	DnnNIOnlyForPcfDiscovery bool `mapstructure:"dnnNIOnlyForPcfDiscovery" json:"dnnNIOnlyForPcfDiscovery,omitempty"`
}

// NrfHealthCheck
//
//	Purpose:
//	  Configured Nnrf health check info
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under Nnrf
type NrfHealthCheck struct {
	// Check Interval
	// unit: millisecond
	// Default: 90000,range 10000-300000
	// Optional
	Interval uint32 `mapstructure:"interval" json:"interval,omitempty"`
	// Check Retries,in case of N consecutive failure responses then SMF performs failover/failback to the highest priority available NRF
	// Default:2,range 1-255
	// Optional
	ConsecSuccOrFail uint8 `mapstructure:"consecSuccOrFail" json:"consecSuccOrFail,omitempty"`
	// Health check switch, default switch on
	// Default:false
	// Optional
	Disable bool `mapstructure:"disable" json:"disable,omitempty"`
}

// Nnrf NfManagement Service SupportedFeatures Related Configuration
//
//  Purpose:
//    Configured Nnrf NfManagement service SupportedFeatures
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Configured under Nnrf
//type NnrfNfmgmtSuppFeat struct {
//}

// Nnrf NfDiscovery Service SupportedFeatures Related Configuration
//
//	Purpose:
//	  Configured Nnrf NfDiscovery service SupportedFeatures
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under Nnrf
type NnrfNfDiscSuppFeat struct {
	// Support of some additional query parameters, please refer to 3GPP TS 29.510 \n
	// Default: false \n
	// Optional
	QueryParamsExt1 bool `mapstructure:"queryParamsExt1"    json:"queryParamsExt1,omitempty"`
	// Support of some additional query parameters, please refer to 3GPP TS 29.510 \n
	// Default: false \n
	// Optional
	QueryParamsExt2 bool `mapstructure:"queryParamsExt2"    json:"queryParamsExt2,omitempty"`
}

// Nsmf Interface Related Configuration
//
//	Purpose:
//	  Configured Nsmf interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type NsmfCfg struct {
	// Default timer configuration for all Nsmf service operations.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries  \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
	// SupportedFeatures configuration of Nsmf_PDUSession service api, referring to 3GPP TS 29.502\n
	// Default: nil \n
	// Optional
	PduSessSupportedFeatures *NsmfPduSessSuppFeat `mapstructure:"pduSessSupportedFeatures" json:"pduSessSupportedFeatures,omitempty"`
	// SupportedFeatures configuration of Nsmf_EventExposure service api, referring to 3GPP TS 29.508\n
	// Default: nil \n
	// Optional
	// EventExposureSupportedFeatures *NsmfEventExposureSuppFeat `mapstructure:"eventExpoSupportedFeatures" json:"eventExpoSupportedFeatures,omitempty"`
}

// Nsmf PDUSession Service SupportedFeatures Related Configuration
//
//	Purpose:
//	  Configured Nsmf PDUSession service SupportedFeatures
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under Nsmf
type NsmfPduSessSuppFeat struct {
	// Cellular IoT \n
	// Default: false \n
	// Optional
	CIOT bool `mapstructure:"cIOT"    json:"cIOT,omitempty"`
	// Deployments Topologies with specific SMF Service Areas \n
	// Default: false \n
	// Optional
	DTSSA bool `mapstructure:"dTSSA"    json:"dTSSA,omitempty"`
}

// Nsmf EventExposure Service SupportedFeatures Related Configuration
//
//  Purpose:
//    Configured Nsmf EventExposure service SupportedFeatures
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Configured under Nsmf
//	Note: This Service has required supported feature for VzW Phase 3
//type NsmfEventExposureSuppFeat struct {
//}

// Nudm Interface Related Configuration
//
//	Purpose:
//	  Configured Nudm interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type NudmCfg struct {
	// Timer configuration per Nudm service operation.\n
	// Key can be:
	//  - "sdm_get"
	//	- "sdm_subscribe"
	//	- "sdm_unsubscribe"
	//  - "registration"
	//	- "deregistration"
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: nil.\n
	// Optional.
	Timer map[TimerType]TimerConfig `mapstructure:"timer" json:"timer,omitempty"`
	// Default timer configuration for all Nudm service operations.\n
	// Value configured in Timer takes higher precedence over this.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries  \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
	// Nudm service version \n
	// Default value: nil \n
	// Optional.
	NudmSvcApiFullVersion *ApiVersion `mapstructure:"nudmSvcApiFullVersion" json:"nudmSvcApiFullVersion,omitempty"`
	// Sdm subscribe related configuration
	// Optional
	SubscriptionCfg *SdmSubscriptionInfo `mapstructure:"subscriptionCfg" json:"subscriptionCfg,omitempty"`
	// SupportedFeatures configuration of Nudm_SDM service api, referring to 3GPP TS 29.503\n
	// Default: nil \n
	// Optional
	// SdmSupportedFeature  *NudmSdmSuppFeat  `mapstructure:"sdmSupportedFeature"    json:"sdmSupportedFeature,omitempty"`
	// SupportedFeatures configuration of Nudm_UECM service api, referring to 3GPP TS 29.503\n
	// Default: nil \n
	// Optional
	// UecmSupportedFeature *NudmUecmSuppFeat `mapstructure:"uecmSupportedFeature"    json:"uecmSupportedFeature,omitempty"`
	// Uecm registration related configuration
	// Optional
	RegistrationCfg *UecmRegistrationInfo `mapstructure:"registrationCfg" json:"registrationCfg,omitempty"`
	// Whether smf uses nfset information about this interface from nrf.\n
	// Default: false.\n
	// Optional.
	EnableNfSet bool `mapstructure:"enableNfSet" json:"enableNfSet,omitempty"`
	// GetSmData optional query parameters that will be excluded from the request towards UDM \n
	// Default: nil \n
	// optional
	ExcludedGetSmDataParams *GetSmDataQueryParams `mapstructure:"excludedGetSmDataParams" json:"excludedGetSmDataParams,omitempty"`
	// Decide whether to include the DNN operator identifier.\n
	// DnnNIAndOI is false: SMF send NI to UDM.\n
	// DnnNIAndOI is true: SMF send NI+OI to UDM.\n
	// Default: false.\n
	// Optional
	DnnNIAndOI bool `mapstructure:"dnnNIAndOI" json:"dnnNIAndOI,omitempty"`
	// Whether smf consolidates NF discovery requests for UECM and SDM
	// When true, SMF includes both UECM and SDM services in UDM NF Discovery request to NRF
	// When false or not configured, SMF includes only the required service (either SDM or UECM) in UDM NF Discovery request to NRF
	// Default: false
	// Optional
	DiscCombineSvc bool `mapstructure:"discCombineSvc" json:"discCombineSvc,omitempty"`
}

// Nudm SDM Service SupportedFeatures Related Configuration
//
//  Purpose:
//    Configured Nudm SDM  service SupportedFeatures
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Configured under Nudm
//type NudmSdmSuppFeat struct {
//}

// Nudm UECM Service SupportedFeatures Related Configuration
//
//  Purpose:
//    Configured Nudm UECM service SupportedFeatures
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Configured under Nudm
//type NudmUecmSuppFeat struct {
//}

// N1 Interface Related Configuration
//
//	Purpose:
//	  Configured N1 interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type N1Cfg struct {
	// Timer configuration per N1 msg.\n
	// Key can be:
	//  - "t3591"
	//	- "t3592"
	// Value Range:
	//  - Interval: 100 - 24000 ms
	//  - Retries: 0 - 5
	// Default: nil.\n
	// Optional.
	Timer map[TimerType]TimerConfig `mapstructure:"timer" json:"timer,omitempty"`
	// Default timer configuration for all N1 service operations.\n
	// Value configured in Timer takes higher precedence over this.\n
	// Value Range:
	//  - Interval: 100 - 24000 ms
	//  - Retries: 0 - 5
	// Default: 16000 for interval and 0 for retries  \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
}

// N2 Interface Related Configuration
//
//	Purpose:
//	  Configured N2 interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type N2Cfg struct {
	// Timer configuration per N2 msg.\n
	// Key can be:
	//  - "n2_res_setup_req"
	//	- "n2_res_mod_req"
	//  - "n2_res_rel_cmd"
	//	- "n2_5to4_ho"
	// Value Range:
	//  - Interval: 100 - 16000 ms
	//  - Retries: 0 - 5
	// Default: nil.\n
	// Optional.
	Timer map[TimerType]TimerConfig `mapstructure:"timer" json:"timer,omitempty"`
	// Default timer configuration for all N2 service operations.\n
	// Value configured in Timer takes higher precedence over this.\n
	// Value Range:
	//  - Interval: 100 - 16000 ms
	//  - Retries: 0 - 5
	// Default: 16000 for interval and 0 for retries  \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
}

// GtpcGrpc Interface Related Configuration
//
//	Purpose:
//	  Configured Gtpc Grpc interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type GtpcGrpcCfg struct {
	// Default timer configuration for all Gtpc service Grpc interface operations.\n
	// Value Range
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 7000 for interval and 0 for retries \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
}

// PfcpGrpc Interface Related Configuration
//
//	Purpose:
//	  Configured Pfcp Grpc interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type PfcpGrpcCfg struct {
	// Default timer configuration for Pfcp service Grpc interface operations.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
}

// AAA Client Grpc Interface Related Configuration
//
//	Purpose:
//	  Configured  AAA Client Grpc interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type AaaClientGrpcCfg struct {
	// Default timer configuration for all AAA Client service Grpc interface operations.\n
	// Value Range:
	//  - Interval: 100 - 10000 ms
	//  - Retries: 0 - 5
	// Default: 7000 for interval and 0 for retries \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
}

// CdrMgr Grpc Interface Related Configuration
//
//	Purpose:
//	  Configured  CdrMgr Grpc interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type CdrMgrGrpcCfg struct {
	// Default timer configuration for all CdrMgr service Grpc interface operations.\n
	// Value Range:
	//  - Interval: 100 - 8000 ms
	//  - Retries: 0 - 5
	// Default: 3000 for interval and 0 for retries \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
}

// PcscfCfg pcscf Interface Related Configuration
//
//	Purpose:
//	  Configured  pcscf interface parameters of smfsm
//	  e.g., application timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type PcscfCfg struct {
	// Default timer configuration for all Pcscf check ip healthy operations.\n
	// Value Range:
	//  - Interval: 1000 - 100000 ms
	//  - Retries: 0 - 5
	// Default: 10000 for interval and 0 for retries \n
	// Optional.
	DefaultTimer *TimerConfig `mapstructure:"defaultTimer" json:"defaultTimer,omitempty"`
	// The value indicates max number of pcscf Ip to cache.\n
	// Value Range: 0 - 1000
	// Default: 100 \n
	// Optional.
	MaxCachedIp uint32 `mapstructure:"maxCachedIp" json:"maxCachedIp,omitempty"`
	// Whether allows cached Ip to be used to continue session setup, this is only applicable when DNS resolution failed.\n
	// Default: false, means no allow cached Ip to be used to continue session setup.\n
	// Optional.
	AllowUsingCachedIp bool `mapstructure:"allowUsingCachedIp" json:"allowUsingCachedIp,omitempty"`
}

// S5S8 Interface Related Configuration
//
//	Purpose:
//	  Configured S5S8 interface parameters of smfsm
//	  e.g., gtpc sleep timer etc
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfSvcCfg
type S5S8Cfg struct {
	// The time is used to delay sending the triggered msg, when there is a triggered create/update/delete bearer request after CSR, MBR.
	// Uint: Millisecond
	// Default: 200
	// Value in range 50~2000
	// Optional
	GtpcReqSleepTimer uint `mapstructure:"gtpcReqSleepTimer" json:"gtpcReqSleepTimer,omitempty"`
}

// Diameter Interface Related Configuration
//
//	Purpose:
//	  Configured some interface parameters of diameter
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  a common Configured under AaaSvcCfg and smfsm
type DiameterIntfCfg struct {
	// The timeout of send diameter message, unit is million second
	// Default value 1000.\n
	// S6b lite Default value 3000
	// S6b Lite Range 100 - 60000 ms
	// Optional.
	ReqTimeoutMilliSec int `mapstructure:"reqTimeoutMilliSec" json:"reqTimeoutMilliSec,omitempty"`
	// The retry times if aaa-client need to do retry in error handle.The total retry times for one diameter server.
	// Default value 0 .\n
	// Optional
	RetryTimes int `mapstructure:"retryTimes" json:"retryTimes,omitempty"`
	// Use to describe Gx interface error handle. It is an error handle name reference to error handle template
	// if it is omited, use terminate action for all kind of condition in Gx interface.
	// Optional
	ErrorHandleRef string `mapstructure:"errorHandleRef" json:"errorHandleRef,omitempty"`
}

// TS 23501
// Annex E (informative):
// Communication models for NF/NF services interaction
// E.1	General
// This annex provides a high level description of the different communication models that NF and NF services can use to interact which each other. Table E.1-1 summarizes the communication models, their usage and how they relate to the usage of an SCP.
//
// Model A - Direct communication without NRF interaction: Neither NRF nor SCP are used. Consumers are configured with producers' "NF profiles" and directly communicate with a producer of their choice.
//
// Model B - Direct communication with NRF interaction: Consumers do discovery by querying the NRF. Based on the discovery result, the consumer does the selection. The consumer sends the request to the selected producer.
//
// Model C - Indirect communication without delegated discovery: Consumers do discovery by querying the NRF. Based on discovery result, the consumer does the selection of an NF Set or a specific NF instance of NF set. The consumer sends the request to the SCP containing the address of the selected service producer pointing to a NF service instance or a set of NF service instances. In the latter case, the SCP selects an NF Service instance. If possible, the SCP interacts with NRF to get selection parameters such as location, capacity, etc. The SCP routes the request to the selected NF service producer instance.
//
// Model D - Indirect communication with delegated discovery: Consumers do not do any discovery or selection. The consumer adds any necessary discovery and selection parameters required to find a suitable producer to the service request. The SCP uses the request address and the discovery and selection parameters in the request message to route the request to a suitable producer instance. The SCP can perform discovery with an NRF and obtain a discovery result.
type CommunicationModel string

const (
	MODEL_A CommunicationModel = "A"
	MODEL_B CommunicationModel = "B"
	MODEL_C CommunicationModel = "C"
	MODEL_D CommunicationModel = "D"
)

type NfInterface string

const (
	NfInterface_AMF NfInterface = "namf"
	NfInterface_UDM NfInterface = "nudm"
	NfInterface_PCF NfInterface = "npcf"
	NfInterface_CHF NfInterface = "nchf"
)

// NF discovery optional query parameters that SMF is using \n
//
//	Purpose:\n
//	  Configured as value of ExcludedDiscQueryParams map \n
//	  Indicate that SMF will not include those in discovery query towards NRF if a parameter is set as true\n
//
//	Data model:\n
//	  Refer to TS 29510 V16.6.0 (2020-12) Table 6.2.3.2.3.1-1 \n
//
//	Usage:\n
//	  Configured under NnrfCfg\n
type QueryParams struct {
	// requester-nf-instance-id query parameter \n
	// Default: false \n
	// Optional
	RequesterNfInstanceId bool `mapstructure:"requesterNfInstanceId" json:"requesterNfInstanceId,omitempty"`
	// requester-nf-instance-fqdn query parameter \n
	// Default: false \n
	// Optional
	RequesterNfInstanceFqdn bool `mapstructure:"requesterNfInstanceFqdn" json:"requesterNfInstanceFqdn,omitempty"`
	// service-names query parameter \n
	// Default: false \n
	// Optional
	ServiceNames bool `mapstructure:"serviceNames" json:"serviceNames,omitempty"`
	// target-plmn-list query parameter \n
	// Default: false \n
	// Optional
	TargetPlmnList bool `mapstructure:"targetPlmnList" json:"targetPlmnList,omitempty"`
	// target-nf-instance-id query parameter \n
	// Default: false \n
	// Optional
	TargetNfInstanceId bool `mapstructure:"targetNfInstanceId" json:"targetNfInstanceId,omitempty"`
	// snssais query parameter \n
	// Default: false \n
	// Optional
	Snssais bool `mapstructure:"snssais" json:"snssais,omitempty"`
	// dnn query parameter \n
	// Default: false \n
	// Optional
	Dnn bool `mapstructure:"dnn" json:"dnn,omitempty"`
	// tai query parameter \n
	// Default: false \n
	// Optional
	Tai bool `mapstructure:"tai" json:"tai,omitempty"`
	// amf-region-id query parameter \n
	// Default: false \n
	// Optional
	AmfRegionId bool `mapstructure:"amfRegionId" json:"amfRegionId,omitempty"`
	// amf-set-id query parameter \n
	// Default: false \n
	// Optional
	AmfSetId bool `mapstructure:"amfSetId" json:"amfSetId,omitempty"`
	// guami query parameter \n
	// Default: false \n
	// Optional
	Guami bool `mapstructure:"guami" json:"guami,omitempty"`
	// supi query parameter \n
	// Default: false \n
	// Optional
	Supi bool `mapstructure:"supi" json:"supi,omitempty"`
	// gpsi query parameter \n
	// Default: false \n
	// Optional
	Gpsi bool `mapstructure:"gpsi" json:"gpsi,omitempty"`
	// routing-indicator query parameter \n
	// Default: false \n
	// Optional
	RoutingIndicator bool `mapstructure:"routingIndicator" json:"routingIndicator,omitempty"`
	// group-id-list query parameter \n
	// Default: false \n
	// Optional
	GroupIdList bool `mapstructure:"groupIdList" json:"groupIdList,omitempty"`
	// preferred-locality query parameter \n
	// Default: false \n
	// Optional
	PreferredLocality bool `mapstructure:"preferredLocality" json:"preferredLocality,omitempty"`
	// access-type query parameter \n
	// Default: false \n
	// Optional
	AccessType bool `mapstructure:"accessType" json:"accessType,omitempty"`
	// target-nf-set-id query parameter \n
	// Default: false \n
	// Optional
	TargetNfSetId bool `mapstructure:"targetNfSetId" json:"targetNfSetId,omitempty"`
	// nef-id query parameter \n
	// Default: false \n
	// Optional
	NefId bool `mapstructure:"nefId" json:"nefId,omitempty"`
}

// UDM Get SM Data optional query parameters that SMF is using \n
//
//	Purpose:\n
//	  Configured as value of ExcludedGetSmDataParams \n
//	  Indicate that SMF will not include those in GetSmData query towards UDM if a parameter is set as true\n
//
//	Data model:\n
//	  Refer to TS 29503 V16.6.0 (2020-12) 5.2.2.2.5 \n
//
//	Usage:\n
//	  Configured under NudmCfg\n
type GetSmDataQueryParams struct {
	// snssai query parameter \n
	// Default: false \n
	// Optional
	Snssai bool `mapstructure:"snssai" json:"snssai,omitempty"`
	// dnn query parameter \n
	// Default: false \n
	// Optional
	Dnn bool `mapstructure:"dnn" json:"dnn,omitempty"`
}
