package v1alpha2

// 5GSM CAUSE defined
type Var5gSmCauseToUe uint8

const (
	SM_CAUSE_OPERATOR_DETERMINED_BARRING                      Var5gSmCauseToUe = 8
	SM_CAUSE_INSUFFICIENT_RESOURCES                           Var5gSmCauseToUe = 26
	SM_CAUSE_MISSING_OR_UNKNOWN_DNN                           Var5gSmCauseToUe = 27
	SM_CAUSE_UNKNOWN_PDUSESSION_TYPE                          Var5gSmCauseToUe = 28
	SM_CAUSE_USER_AUTH_FAILED                                 Var5gSmCauseToUe = 29
	SM_CAUSE_REQUEST_REJECTED                                 Var5gSmCauseToUe = 31
	SM_CAUSE_SERVICE_OPTION_NOT_SUPPORTED                     Var5gSmCauseToUe = 32
	SM_CAUSE_REQ_SRV_OPTION_NOT_SUBSCRIBED                    Var5gSmCauseToUe = 33
	SM_CAUSE_SRV_OPT_TMP_OUT_ORDER                            Var5gSmCauseToUe = 34
	SM_CAUSE_PTI_ALREADY_IN_USE                               Var5gSmCauseToUe = 35
	SM_CAUSE_REGULAR_DEACT                                    Var5gSmCauseToUe = 36
	SM_CAUSE_NETWORK_FAILED                                   Var5gSmCauseToUe = 38
	SM_CAUSE_REACT_RWQ                                        Var5gSmCauseToUe = 39
	SM_CAUSE_INVALID_PDUSESSION_ID                            Var5gSmCauseToUe = 43
	SM_CAUSE_SEMANTIC_ERRORS_IN_PACKET_FILTERS                Var5gSmCauseToUe = 44
	SM_CAUSE_SYNTACTICAL_ERROR_IN_PACKET_FILTERS              Var5gSmCauseToUe = 45
	SM_CAUSE_OUT_OF_LADN_SERVICE_AREA                         Var5gSmCauseToUe = 46
	SM_CAUSE_PTI_MISMATCH                                     Var5gSmCauseToUe = 47
	SM_CAUSE_PDUSESSION_TYPE_IPV4_ONLY                        Var5gSmCauseToUe = 50
	SM_CAUSE_PDUSESSION_TYPE_IPV6_ONLY                        Var5gSmCauseToUe = 51
	SM_CAUSE_PDUSESSION_DOSE_NOT_EXIST                        Var5gSmCauseToUe = 54
	SM_CAUSE_INSUFFICIENT_RSC_FOR_SEP_SLICE_DNN               Var5gSmCauseToUe = 67
	SM_CAUSE_NOT_SUPPORTED_SSC_MODE                           Var5gSmCauseToUe = 68
	SM_CAUSE_INSUFFICIENT_RSC_FOR_SEPCIFIC_SLICE              Var5gSmCauseToUe = 69
	SM_CAUSE_MISSING_UNKNOWN_DNN_IN_SLICE                     Var5gSmCauseToUe = 70
	SM_CAUSE_SERVING_NETWORK_NOT_AUTHORIZED                   Var5gSmCauseToUe = 73
	SM_CAUSE_INVALID_PTI_VALUE                                Var5gSmCauseToUe = 81
	SM_CAUSE_MAX_DATA_UE_FOR_UP_INTE_PROTECTION_LOW           Var5gSmCauseToUe = 82
	SM_CAUSE_SEMANTIC_ERROR_IN_QOS_OP                         Var5gSmCauseToUe = 83
	SM_CAUSE_SYNC_ERROR_IN_QOS_OP                             Var5gSmCauseToUe = 84
	SM_CAUSE_INVALID_MAPPED_EPS_BEARER_IDENTITY               Var5gSmCauseToUe = 85
	SM_CAUSE_DNN_NOT_SUPPORTED_OR_NOT_SUBSCRIBED_IN_THE_SLICE Var5gSmCauseToUe = 91
	SM_CAUSE_SEMANTICALLY_INCORRECT_MSG                       Var5gSmCauseToUe = 95
	SM_CAUSE_INVALID_MANDATORY_INFO                           Var5gSmCauseToUe = 96
	SM_CAUSE_MSG_TYPE_NO_EXIST                                Var5gSmCauseToUe = 97
	SM_CAUSE_MSG_TYPE_NOT_COMPATIBLE_WITH_PROTO_STATE         Var5gSmCauseToUe = 98
	SM_CAUSE_INFO_ELEMENT_NO_EXIST_IMPLE                      Var5gSmCauseToUe = 99
	SM_CAUSE_CONDITIONAL_IE_ERROR                             Var5gSmCauseToUe = 100
	SM_CAUSE_MSG_NOT_COMPATIBLE_WITH_PROTO_STATE              Var5gSmCauseToUe = 101
)

type AppErrorCode string

// sbi cause code, TS 29.500 Table 5.2.7.2-1: Protocol and application errors common to several 5GC SBI API specifications
const (
	SBI_CAUSE_INVALID_API                      AppErrorCode = "INVALID_API"
	SBI_CAUSE_INVALID_MSG_FORMAT               AppErrorCode = "INVALID_MSG_FORMAT"
	SBI_CAUSE_INVALID_QUERY_PARAM              AppErrorCode = "INVALID_QUERY_PARAM"
	SBI_CAUSE_MANDATORY_QUERY_PARAM_INCORRECT  AppErrorCode = "MANDATORY_QUERY_PARAM_INCORRECT"
	SBI_CAUSE_OPTIONAL_QUERY_PARAM_INCORRECT   AppErrorCode = "OPTIONAL_QUERY_PARAM_INCORRECT"
	SBI_CAUSE_MANDATORY_QUERY_PARAM_MISSING    AppErrorCode = "MANDATORY_QUERY_PARAM_MISSING"
	SBI_CAUSE_MANDATORY_IE_INCORRECT           AppErrorCode = "MANDATORY_IE_INCORRECT"
	SBI_CAUSE_OPTIONAL_IE_INCORRECT            AppErrorCode = "OPTIONAL_IE_INCORRECT"
	SBI_CAUSE_MANDATORY_IE_MISSING             AppErrorCode = "MANDATORY_IE_MISSING"
	SBI_CAUSE_UNSPECIFIED_MSG_FAILURE          AppErrorCode = "UNSPECIFIED_MSG_FAILURE"
	SBI_CAUSE_NF_DISCOVERY_FAILURE             AppErrorCode = "NF_DISCOVERY_FAILURE"
	SBI_CAUSE_INVALID_DISCOVERY_PARAM          AppErrorCode = "INVALID_DISCOVERY_PARAM"
	SBI_CAUSE_RESOURCE_CONTEXT_NOT_FOUND       AppErrorCode = "RESOURCE_CONTEXT_NOT_FOUND"
	SBI_CAUSE_MODIFICATION_NOT_ALLOWED         AppErrorCode = "MODIFICATION_NOT_ALLOWED"
	SBI_CAUSE_SUBSCRIPTION_NOT_FOUND           AppErrorCode = "SUBSCRIPTION_NOT_FOUND"
	SBI_CAUSE_RESOURCE_URI_STRUCTURE_NOT_FOUND AppErrorCode = "RESOURCE_URI_STRUCTURE_NOT_FOUND"
	SBI_CAUSE_INCORRECT_LENGTH                 AppErrorCode = "INCORRECT_LENGTH"
	SBI_CAUSE_NF_CONGESTION_RISK               AppErrorCode = "NF_CONGESTION_RISK"
	SBI_CAUSE_INSUFFICIENT_RESOURCES           AppErrorCode = "INSUFFICIENT_RESOURCES"
	SBI_CAUSE_UNSPECIFIED_NF_FAILURE           AppErrorCode = "UNSPECIFIED_NF_FAILURE"
	SBI_CAUSE_SYSTEM_FAILURE                   AppErrorCode = "SYSTEM_FAILURE"
	SBI_CAUSE_NF_FAILOVER                      AppErrorCode = "NF_FAILOVER"
	SBI_CAUSE_NF_SERVICE_FAILOVER              AppErrorCode = "NF_SERVICE_FAILOVER"
	SBI_CAUSE_NF_CONGESTION                    AppErrorCode = "NF_CONGESTION"
	SBI_CAUSE_TARGET_NF_NOT_REACHABLE          AppErrorCode = "TARGET_NF_NOT_REACHABLE"
	SBI_CAUSE_TIMED_OUT_REQUEST                AppErrorCode = "TIMED_OUT_REQUEST"
	SBI_CAUSE_SCP_REDIRECTION                  AppErrorCode = "SCP_REDIRECTION"
)

// PCF cause code, TS29512 Table 5.7.3-1: Application errors when PCF acts as a server
const (
	PCF_CAUSE_MAPPING_USER_UNKNOWN                        AppErrorCode = "USER_UNKNOWN"
	PCF_CAUSE_MAPPING_ERROR_INITIAL_PARAMETERS            AppErrorCode = "ERROR_INITIAL_PARAMETERS"
	PCF_CAUSE_MAPPING_ERROR_TRIGGER_EVENT                 AppErrorCode = "ERROR_TRIGGER_EVENT"
	PCF_CAUSE_MAPPING_TRAFFIC_MAPPING_INFO_REJECTED       AppErrorCode = "TRAFFIC_MAPPING_INFO_REJECTED"
	PCF_CAUSE_MAPPING_ERROR_TRAFFIC_MAPPING_INFO_REJECTED AppErrorCode = "ERROR_TRAFFIC_MAPPING_INFO_REJECTED"
	PCF_CAUSE_MAPPING_ERROR_CONFLICTING_REQUEST           AppErrorCode = "ERROR_CONFLICTING_REQUEST"
	PCF_CAUSE_MAPPING_LATE_OVERLAPPING_REQUEST            AppErrorCode = "LATE_OVERLAPPING_REQUEST"
	PCF_CAUSE_MAPPING_POLICY_CONTEXT_DENIED               AppErrorCode = "POLICY_CONTEXT_DENIED"
	PCF_CAUSE_MAPPING_VALIDATION_CONDITION_NOT_MET        AppErrorCode = "VALIDATION_CONDITION_NOT_MET"
	PCF_CAUSE_MAPPING_PENDING_TRANSACTION                 AppErrorCode = "PENDING_TRANSACTION"
	PCF_CAUSE_MAPPING_INVALID_BDT_POLICY                  AppErrorCode = "INVALID_BDT_POLICY"
)

// SDM cause code, TS29.503 Table 6.1.7.3-1: Application errors
const (
	SDM_CAUSE_MAPPING_NF_CONSUMER_REDIRECT_ONE_TXN AppErrorCode = "NF_CONSUMER_REDIRECT_ONE_TXN"
	SDM_CAUSE_MAPPING_CONTEXT_NOT_FOUND            AppErrorCode = "CONTEXT_NOT_FOUND"
	SDM_CAUSE_MAPPING_DATA_NOT_FOUND               AppErrorCode = "DATA_NOT_FOUND"
	SDM_CAUSE_MAPPING_USER_NOT_FOUND               AppErrorCode = "USER_NOT_FOUND"
	SDM_CAUSE_MAPPING_GROUP_IDENTIFIER_NOT_FOUND   AppErrorCode = "GROUP_IDENTIFIER_NOT_FOUND"
	SDM_CAUSE_MAPPING_SUBSCRIPTION_NOT_FOUND       AppErrorCode = "SUBSCRIPTION_NOT_FOUND"
	SDM_CAUSE_MAPPING_UNSUPPORTED_RESOURCE_URI     AppErrorCode = "UNSUPPORTED_RESOURCE_URI"
	SDM_CAUSE_NF_CONGESTION_RISK                   AppErrorCode = "NF_CONGESTION_RISK"
)

// UECM cause code, TS29503 Table 6.2.7.3-1: Application errors
const (
	UECM_CAUSE_MAPPING_NF_CONSUMER_REDIRECT_ONE_TXN              AppErrorCode = "NF_CONSUMER_REDIRECT_ONE_TXN"
	UECM_CAUSE_MAPPING_UNKNOWN_5GS_SUBSCRIPTION                  AppErrorCode = "UNKNOWN_5GS_SUBSCRIPTION"
	UECM_CAUSE_MAPPING_NO_PS_SUBSCRIPTION                        AppErrorCode = "NO_PS_SUBSCRIPTION"
	UECM_CAUSE_MAPPING_ROAMING_NOT_ALLOWED                       AppErrorCode = "ROAMING_NOT_ALLOWED"
	UECM_CAUSE_MAPPING_USER_NOT_FOUND                            AppErrorCode = "USER_NOT_FOUND"
	UECM_CAUSE_MAPPING_CONTEXT_NOT_FOUND                         AppErrorCode = "CONTEXT_NOT_FOUND"
	UECM_CAUSE_MAPPING_ACCESS_NOT_ALLOWED                        AppErrorCode = "ACCESS_NOT_ALLOWED"
	UECM_CAUSE_MAPPING_RAT_NOT_ALLOWED                           AppErrorCode = "RAT_NOT_ALLOWED"
	UECM_CAUSE_MAPPING_DNN_NOT_ALLOWED                           AppErrorCode = "DNN_NOT_ALLOWED"
	UECM_CAUSE_MAPPING_REAUTHENTICATION_REQUIRED                 AppErrorCode = "REAUTHENTICATION_REQUIRED"
	UECM_CAUSE_MAPPING_INVALID_GUAMI                             AppErrorCode = "INVALID_GUAMI"
	UECM_CAUSE_MAPPING_SERVICE_NOT_PROVISIONED                   AppErrorCode = "SERVICE_NOT_PROVISIONED"
	UECM_CAUSE_MAPPING_SERVICE_NOT_ALLOWED                       AppErrorCode = "SERVICE_NOT_ALLOWED"
	UECM_CAUSE_MAPPING_TEMPORARY_REJECT_REGISTRATION             AppErrorCode = "TEMPORARY_REJECT_REGISTRATION"
	UECM_CAUSE_MAPPING_TEMPORARY_REJECT_HANDOVER_ONGOING_ONGOING AppErrorCode = "TEMPORARY_REJECT_HANDOVER_ONGOING_ONGOING"
	UECM_CAUSE_MAPPING_UNPROCESSABLE_REQUEST                     AppErrorCode = "UNPROCESSABLE_REQUEST"
)

// CHF cause code, TS32291 Table 6.1.7.3-1: Application errors
const (
	CHF_CAUSE_MAPPING_CHARGING_FAILED           AppErrorCode = "CHARGING_FAILED"
	CHF_CAUSE_MAPPING_RE_AUTHORIZATION_FAILED   AppErrorCode = "RE_AUTHORIZATION_FAILED"
	CHF_CAUSE_MAPPING_CHARGING_NOT_APPLICABLE   AppErrorCode = "CHARGING_NOT_APPLICABLE"
	CHF_CAUSE_MAPPING_USER_UNKNOWN              AppErrorCode = "USER_UNKNOWN"
	CHF_CAUSE_MAPPING_END_USER_REQUEST_DENIED   AppErrorCode = "END_USER_REQUEST_DENIED"
	CHF_CAUSE_MAPPING_QUOTA_LIMIT_REACHED       AppErrorCode = "QUOTA_LIMIT_REACHED"
	CHF_CAUSE_MAPPING_END_USER_REQUEST_REJECTED AppErrorCode = "END_USER_REQUEST_REJECTED"
)

// GTPC cause code. 3GPP TS 29.274 V12.6 Section 8.4
type GtpcCauseCode uint8

const (
	GTP_CAUSE_RESERVED_0              = 0
	GTP_CAUSE_RESERVED                = 1
	GTP_CAUSE_LOCAL_DETACH            = 2
	GTP_CAUSE_COMPLETE_DETACH         = 3
	GTP_CAUSE_RAT_CHG_F3GPP_TN3GPP    = 4
	GTP_CAUSE_ISR_DEACTIVATION        = 5
	GTP_CAUSE_ERROR_INDI_RECVD        = 6
	GTP_CAUSE_IMSI_DETACH_ONLY        = 7
	GTP_CAUSE_REACT_REQUEST           = 8
	GTP_CAUSE_REACTIVATION_REQUESTED  = 8
	GTP_CAUSE_PDN_RECN_DISALLOWED     = 9
	GTP_CAUSE_ACCESS_CHG_FN3GPP_T3GPP = 10
	GTP_CAUSE_PDN_CON_INACT_TIMER_EXP = 11
	GTP_CAUSE_PGW_NO_RESPONDING       = 12
	GTP_CAUSE_NETWORK_FAILURE         = 13
	GTP_CAUSE_QOS_PARAMETER_MISMATCH  = 14
	// 15, Spare
	GTP_CAUSE_REQUEST_ACCEPTED                = 16
	GTP_CAUSE_REQUEST_ACCEPT_PARTIAL          = 17
	GTP_CAUSE_NEW_PDN_TYPE_NET_PREF           = 18
	GTP_CAUSE_NEW_PDN_TYPE_SINGLE_ADDR_BEARER = 19
	// 20 to 63 are reserved
	GTP_CAUSE_CONTEXT_NOT_FOUND                  = 64
	GTP_CAUSE_INVALID_MSG_FORMAT                 = 65
	GTP_CAUSE_VERSION_NOT_SUPPORTED_BY_NEXT_PEER = 66
	GTP_CAUSE_INVALID_LENGTH                     = 67
	GTP_CAUSE_SERVICE_NOT_SUPPORTED              = 68
	GTP_CAUSE_MANDATORY_IE_INCORRECT             = 69
	GTP_CAUSE_MANDATORY_IE_MISSING               = 70
	// 71 shall not be used
	GTP_CAUSE_SYSTEM_FAILURE        = 72
	GTP_CAUSE_NO_RESOURCE_AVAILABLE = 73
	GTP_CAUSE_SEM_ERROR_IN_TFT_OP   = 74
	GTP_CAUSE_SYN_ERROR_IN_TFT_OP   = 75
	GTP_CAUSE_SEMANTIC_ERROR_PKT    = 76
	GTP_CAUSE_SYNTACTIC_ERROR_PKT   = 77
	GTP_CAUSE_MISS_UNKNOWN_APN      = 78
	// 79 shall not be used
	GTP_CAUSE_GRE_KEY_NOT_FOUND                        = 80
	GTP_CAUSE_RELOCATION_FAILURE                       = 81
	GTP_CAUSE_DENIED_IN_RAT                            = 82
	GTP_CAUSE_PREFERRED_PDN_TYPE_UNSUP                 = 83
	GTP_CAUSE_ALL_DYNAMIC_ADDR_OCCUPIED                = 84
	GTP_CAUSE_UE_CONTEXT_WITHOUT_TFT_ALREADY_ACTIVATED = 85
	GTP_CAUSE_PROTOCOL_NOT_SUPPORTED                   = 86
	GTP_CAUSE_UE_NOT_RESPONDING                        = 87
	GTP_CAUSE_UE_REFUSES                               = 88
	GTP_CAUSE_SERVICE_DENIED                           = 89
	GTP_CAUSE_UNABLE_TO_PAGE_UE                        = 90
	GTP_CAUSE_NO_MEMORY_AVAILABLE                      = 91
	GTP_CAUSE_USER_AUTH_FAILED                         = 92
	GTP_CAUSE_APN_ACCESS_DENY_NO_SUB                   = 93
	GTP_CAUSE_REQUEST_REJECTED                         = 94
	GTP_CAUSE_SEMANTIC_ERROR_IN_THE_TAD_OPR            = 95
	GTP_CAUSE_IMSI_IMEI_NOT_KNOWN                      = 96
	GTP_CAUSE_SEMANTIC_ERROR_IN_TAD                    = 97
	GTP_CAUSE_SYNTACTIC_ERROR_IN_TAD                   = 98
	// 99 shall not be used
	GTP_CAUSE_REMOTE_PEER_NOT_RESPONDING              = 100
	GTP_CAUSE_COLLISION_NETWORK_INITIA_REQ            = 101
	GTP_CAUSE_UNABLE_TO_PAGE_UE_DUE_TO_SUSPENSION     = 102
	GTP_CAUSE_CONDITIONAL_IE_MISSING                  = 103
	GTP_CAUSE_APN_RESTRICTION_PDN_CON                 = 104
	GTP_CAUSE_INVALID_OVERALL_LENGTH                  = 105
	GTP_CAUSE_DATA_FORWARDING_NOT_SUPPORTED           = 106
	GTP_CAUSE_INVALID_REPLY_FROM_REMOTE_PEER          = 107
	GTP_CAUSE_FALLBACK_TO_GTPV1                       = 108
	GTP_CAUSE_INVALID_PEER                            = 109
	GTP_CAUSE_TEMPORARILY_REJECTED_DUE_TO_HANDOVER    = 110
	GTP_CAUSE_MODIFICATION_NOT_LIMITED_TO_S1U_BEARERS = 111
	GTP_CAUSE_REQ_REJECTED_FOR_PMIPV6                 = 112
	GTP_CAUSE_APN_CONGESTION                          = 113
	GTP_CAUSE_BEARER_HANDING_NOT_SUP                  = 114
	GTP_CAUSE_UE_ALREADY_REATTACHED                   = 115
	GTP_CAUSE_MULTI_PDN_CONNECITON_NOT_ALLOWED        = 116
	GTP_CAUSE_TARGET_ACCESS_RESTRICTED_FOR_SUBSCRIBER = 117
	// 118 shall not be used
	GTP_CAUSE_MME_SGSN_REFUSES_DUE_TO_VPLMN_POLICY       = 119
	GTP_CAUSE_GTPC_ENTITY_CONGESTION                     = 120
	GTP_CAUSE_MULTI_ACCESS_TO_PDN_CONNECTION_NOT_ALLOWED = 126
	GTPC_CAUSE_REQUEST_REJECTED_DUE_TO_UE_CAPABILITY     = 127
	// 128 to 239, Spare: Fot future use in a triggered/response message
	// 240 to 255, Spare: For future use in an initial/request message
)

// Diameter Result-Code, The Result-Code AVP values defined in Diameter base protocol, IETF RFC 6733
type DiameterErrorCode uint16

const (
	//Protocol Errors
	DIAMETER_COMMAND_UNSUPPORTED     = 3001
	DIAMETER_UNABLE_TO_DELIVER       = 3002
	DIAMETER_REALM_NOT_SERVED        = 3003
	DIAMETER_TOO_BUSY                = 3004
	DIAMETER_LOOP_DETECTED           = 3005
	DIAMETER_REDIRECT_INDICATION     = 3006
	DIAMETER_APPLICATION_UNSUPPORTED = 3007
	DIAMETER_INVALID_HDR_BITS        = 3008
	DIAMETER_INVALID_AVP_BITS        = 3009
	DIAMETER_UNKNOWN_PEER            = 3010

	//Transient Failures
	DIAMETER_AUTHENTICATION_REJECTED       = 4001
	DIAMETER_OUT_OF_SPACE                  = 4002
	ELECTION_LOST                          = 4003
	DIAMETER_END_USER_SERVICE_DENIED       = 4010
	DIAMETER_CREDIT_CONTROL_NOT_APPLICABLE = 4011
	DIAMETER_CREDIT_LIMIT_REACHED          = 4012
	//Permanent Failures
	DIAMETER_AVP_UNSUPPORTED           = 5001
	DIAMETER_UNKNOWN_SESSION_ID        = 5002
	DIAMETER_AUTHORIZATION_REJECTED    = 5003
	DIAMETER_INVALID_AVP_VALUE         = 5004
	DIAMETER_MISSING_AVP               = 5005
	DIAMETER_RESOURCES_EXCEEDED        = 5006
	DIAMETER_CONTRADICTING_AVPS        = 5007
	DIAMETER_AVP_NOT_ALLOWED           = 5008
	DIAMETER_AVP_OCCURS_TOO_MANY_TIMES = 5009
	DIAMETER_NO_COMMON_APPLICATION     = 5010
	DIAMETER_UNSUPPORTED_VERSION       = 5011
	DIAMETER_UNABLE_TO_COMPLY          = 5012
	DIAMETER_INVALID_BIT_IN_HEADER     = 5013
	DIAMETER_INVALID_AVP_LENGTH        = 5014
	DIAMETER_INVALID_MESSAGE_LENGTH    = 5015
	DIAMETER_INVALID_AVP_BIT_COMBO     = 5016
	DIAMETER_NO_COMMON_SECURITY        = 5017
	DIAMETER_USER_UNKNOWN              = 5030
	DIAMETER_RATING_FAILED             = 5031
)

// PFCP  TS29244
type N4ErrorValue uint8

const (
	PFCP_CAUSE_MAPPING_REQ_REJ         N4ErrorValue = 64
	PFCP_CAUSE_MAPPING_CTXT_NOT_FOUND  N4ErrorValue = 65
	PFCP_CAUSE_MAPPING_MIE_MISS        N4ErrorValue = 66
	PFCP_CAUSE_MAPPING_CIE_MISS        N4ErrorValue = 67
	PFCP_CAUSE_MAPPING_INVAL_LEN       N4ErrorValue = 68
	PFCP_CAUSE_MAPPING_MIE_WRONG       N4ErrorValue = 69
	PFCP_CAUSE_MAPPING_INVAL_FWDPLCY   N4ErrorValue = 70
	PFCP_CAUSE_MAPPING_INVAL_FTEID_OPT N4ErrorValue = 71
	PFCP_CAUSE_MAPPING_NO_ASSO         N4ErrorValue = 72
	PFCP_CAUSE_MAPPING_RULE_CM_FAIL    N4ErrorValue = 73
	PFCP_CAUSE_MAPPING_CONGESTION      N4ErrorValue = 74
	PFCP_CAUSE_MAPPING_NO_RESOURCE     N4ErrorValue = 75
	PFCP_CAUSE_MAPPING_SVC_NO_SUPPORT  N4ErrorValue = 76
	PFCP_CAUSE_MAPPING_SYS_FAIL        N4ErrorValue = 77
	PFCP_CAUSE_MAPPING_REDIR_REQ       N4ErrorValue = 78
)

// Application Error Code Mapping
//
//	Purpose:
//	  Mapping between application error codes and 5GSM cause codes by SMF.
//
//	Data model:
//	  Refer to the description for attribute below.
//
//	Usage:
//	  Used in CauseCodeMappingProfile as map value
type AppCauseMap struct {
	// HTTP application error code.\n
	//  Value for 5GC SBI API
	//  -"INVALID_API"
	//  -"INVALID_MSG_FORMAT"
	//  -"INVALID_QUERY_PARAM"
	//  -"MANDATORY_QUERY_PARAM_INCORRECT"
	//  -"OPTIONAL_QUERY_PARAM_INCORRECT"
	//  -"MANDATORY_QUERY_PARAM_MISSING"
	//  -"MANDATORY_IE_INCORRECT"
	//  -"OPTIONAL_IE_INCORRECT"
	//  -"MANDATORY_IE_MISSING"
	//  -"UNSPECIFIED_MSG_FAILURE"
	//  -"NF_DISCOVERY_FAILURE"
	//  -"INVALID_DISCOVERY_PARAM"
	//  -"RESOURCE_CONTEXT_NOT_FOUND"
	//  -"MODIFICATION_NOT_ALLOWED"
	//  -"SUBSCRIPTION_NOT_FOUND"
	//  -"RESOURCE_URI_STRUCTURE_NOT_FOUND"
	//  -"INCORRECT_LENGTH"
	//  -"NF_CONGESTION_RISK"
	//  -"INSUFFICIENT_RESOURCES"
	//  -"UNSPECIFIED_NF_FAILURE"
	//  -"SYSTEM_FAILURE"
	//  -"NF_FAILOVER"
	//  -"NF_SERVICE_FAILOVER"
	//  -"NF_CONGESTION"
	//  -"TARGET_NF_NOT_REACHABLE"
	//  -"TIMED_OUT_REQUEST"
	//  -"SCP_REDIRECTION"
	// Value for Npcf_SMPolicyControl Service:
	//  -"USER_UNKNOWN"
	//	-"ERROR_INITIAL_PARAMETERS"
	//	-"ERROR_TRIGGER_EVENT"
	//	-"TRAFFIC_MAPPING_INFO_REJECTED"
	//	-"ERROR_TRAFFIC_MAPPING_INFO_REJECTED"
	//	-"ERROR_CONFLICTING_REQUEST"
	//	-"LATE_OVERLAPPING_REQUEST"
	//	-"POLICY_CONTEXT_DENIED"
	//  -"VALIDATION_CONDITION_NOT_MET"
	//	-"PENDING_TRANSACTION"
	//	-"INVALID_BDT_POLICY"
	// Value for Nudm_SubscriberDataManagement service:
	//	-"NF_CONSUMER_REDIRECT_ONE_TXN"
	//  -"CONTEXT_NOT_FOUND"
	//	-"DATA_NOT_FOUND"
	//	-"USER_NOT_FOUND"
	//	-"GROUP_IDENTIFIER_NOT_FOUND"
	//	-"SUBSCRIPTION_NOT_FOUND"
	//	-"UNSUPPORTED_RESOURCE_URI"
	//	-"NF_CONGESTION_RISK"
	// Value for Nudm_UEContextManagement Service
	//	-"NF_CONSUMER_REDIRECT_ONE_TXN"
	//  -"UNKNOWN_5GS_SUBSCRIPTION"
	//	-"NO_PS_SUBSCRIPTION"
	//	-"ROAMING_NOT_ALLOWED"
	//	-"USER_NOT_FOUND"
	//	-"CONTEXT_NOT_FOUND"
	//	-"ACCESS_NOT_ALLOWED"
	//	-"RAT_NOT_ALLOWED"
	//	-"DNN_NOT_ALLOWED"
	//	-"REAUTHENTICATION_REQUIRED"
	//	-"INVALID_GUAMI"
	//	-"SERVICE_NOT_PROVISIONED"
	//	-"SERVICE_NOT_ALLOWED"
	//	-"TEMPORARY_REJECT_REGISTRATION"
	//	-"TEMPORARY_REJECT_HANDOVER_ONGOING_ONGOING"
	//	-"UNPROCESSABLE_REQUEST"
	// Default value "VZW_SPECIFIED_ERROR" means the 5GSM cause value can be applied to any error code associated to a HTTP status code
	// Optional
	AppErrCode AppErrorCode `mapstructure:"appErrCode" json:"appErrCode,omitempty"`
	// 5GSM cause values defined in https://www.3gpp.org/ftp/Specs/archive/24_series/24.501/24501-g40.zip Annex B (informative): Cause values for 5GS session management
	// Value can be:
	// 	-8: Operator Determined Barring
	// 	-26: Insufficient resources
	// 	-27: Missing or unknown DNN
	// 	-28: Unknown PDU session type
	// 	-29: User authentication or authorization failed
	// 	-31: Request rejected, unspecified
	// 	-32: Service option not supported
	// 	-33: Requested service option not subscribed
	// 	-35: PTI already in use
	// 	-36: Regular deactivation
	// 	-38: Network failure
	// 	-39: Reactivation requested
	// 	-43: Invalid PDU session identity
	// 	-44: Semantic errors in packet filter(s)
	// 	-45: Syntactical error in packet filter(s)
	// 	-46: Out of LADN service area
	// 	-47: PTI mismatch
	// 	-50: PDU session type IPv4 only allowed
	// 	-51: PDU session type IPv6 only allowed
	// 	-54: PDU session does not exist
	// 	-67: Insufficient resources for specific slice and DNN
	// 	-68: Not supported SSC mode
	// 	-69: Insufficient resources for specific slic
	// 	-70: Missing or unknown DNN in a slice
	//  -73: Serving network not authorized
	// 	-81: Invalid PTI value
	// 	-82: Maximum data rate per UE for user-plane integrity protection is too low
	// 	-83: Semantic error in the QoS operation
	// 	-84: Syntactical error in the QoS operation
	// 	-85: Invalid mapped EPS bearer identity
	//  -91: DNN not supported or not subscribed in the slice
	// 	-96: invalid mandatory information
	// 	-97: Message type non-existent or not implemented
	// 	-98: Message type not compatible with protocol state
	// 	-99: Information element non-existent or not implemented
	// 	-100: Conditional IE error
	// 	-101: Message not compatible with protocol state
	// At least one value above should be present.\n
	// optional.
	Var5GSmCause Var5gSmCauseToUe `mapstructure:"var5GSmCause" json:"var5GSmCause,omitempty"`
}

// Cause Code Mapping Configuration
//
//  Purpose:
//    Mapping between 5GC interfaces causes and 5GSM Cause Codes by SMF.
//    Overwrite the cause code mapping defined in back end
//
//  Data model:
//    Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.524/29524-g00.zip section 5
//    Refer to the description for attribute below.
//
//  Usage:
//    Use 5GC interfaces cause as the key and 5GSM Cause Code as the value to fill the map type field in the CauseCodeMappingProfile.
//    Save the mapping between N7 interfaces cause and 5GSM cause code in the NpcfCauseMapping field,
//    Save the mapping between N10 UEContextManagement service cause and 5GSM cause code in the NudmUecmCauseMapping field,
//    Save the mapping between N10 SubscriberDataManagement service cause and 5GSM cause code in the NudmSdmCauseMapping field,
//    Save the mapping between N4 interfaces cause and 5GSM cause code in the N4CauseMapping field.
//

type CauseCodeMappingProfile struct {
	// Cause code mapping profile identity.\n
	// If the field is not an empty string, it should be set to the DNN name, and the configuration of CauseCodeMappingProfile is only valid for this DNN.\n
	// Default value "".\n
	// optional.
	Name string `mapstructure:"name" json:"name,omitempty"`
	// List of ranges of SUPIs that can be served by this cause code mapping.\n
	// If the length of SupiRanges is not zero, the configuration of CauseCodeMappingProfile is only valid for the session of supi in the SupiRanges.\n
	// Default value nil.\n
	// optional.
	SupiRanges []SupiRange `mapstructure:"supiRanges" json:"supiRanges,omitempty"`
	// Mapping from PCF application error code to 5GSM cause values as listed in https://www.3gpp.org/ftp/Specs/archive/29_series/29.524/29524-g00.zip section 5.\n
	// The key defined in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip Table 5.7.3-1: Application errors when PCF acts as a server
	// The value defined in https://www.3gpp.org/ftp/Specs/archive/24_series/24.501/24501-g40.zip Annex B (informative): Cause values for 5GS session management
	// Key is HTTP Status Code
	// Default value nil.\n
	// optional.
	NpcfCauseMapping map[int][]AppCauseMap `mapstructure:"npcfCauseMapping" json:"npcfCauseMapping,omitempty"`
	// Mapping from UDM UECM application error code to 5GSM cause values as listed in https://www.3gpp.org/ftp/Specs/archive/29_series/29.524/29524-g00.zip section 5.\n
	// The application error defined in https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/29503-g30.zip Table 6.2.7.3-1: Application errors
	// The 5GSM cause value defined in https://www.3gpp.org/ftp/Specs/archive/24_series/24.501/24501-g40.zip Annex B (informative): Cause values for 5GS session management
	// Key is HTTP Status Code
	// Default value nil.\n
	// optional.
	NudmUecmCauseMapping map[int][]AppCauseMap `mapstructure:"nudmUecmCauseMapping" json:"nudmUecmCauseMapping,omitempty"`
	// Mapping from UDM SDM application error code to 5GSM cause values as listed in https://www.3gpp.org/ftp/Specs/archive/29_series/29.524/29524-g00.zip section 5.\n
	// The application error defined in https://www.3gpp.org/ftp/Specs/archive/29_series/29.503/29503-g30.zip Table 6.1.7.3-1: Application errors
	// The 5GSM cause value defined in https://www.3gpp.org/ftp/Specs/archive/24_series/24.501/24501-g40.zip Annex B (informative): Cause values for 5GS session management
	// Key is HTTP Status Code
	// Default value nil.\n
	// optional.
	NudmSdmCauseMapping map[int][]AppCauseMap `mapstructure:"nudmSdmCauseMapping" json:"nudmSdmCauseMapping,omitempty"`
	// Mapping from N4 error code to 5GSM cause values as listed in https://www.3gpp.org/ftp/Specs/archive/29_series/29.524/29524-g00.zip section 5.\n
	// The N4 error code defined in https://www.3gpp.org/ftp/Specs/archive/29_series/29.244/29244-g30.zip section 8.2.1 Cause
	// The 5GSM cause value defined in https://www.3gpp.org/ftp/Specs/archive/24_series/24.501/24501-g40.zip Annex B (informative): Cause values for 5GS session management
	// Key can be:
	// 	-64: Request rejected
	// 	-65: Session context not found
	// 	-66: Mandatory IE missing
	// 	-67: Conditional IE missing
	// 	-68: Invalid length
	// 	-69: Mandatory IE incorrec
	// 	-70: Invalid Forwarding Policy
	// 	-71: Invalid F-TEID allocation option
	// 	-72: No established PFCP Association
	// 	-73: Rule creation/modification Failure
	// 	-74: PFCP entity in congestion
	// 	-75: No resources available
	// 	-76: Service not supported
	// 	-77: System failure
	// 	-78: Redirection Requested
	// Value can be:
	// 	-8: Operator Determined Barring
	// 	-26: Insufficient resources
	// 	-27: Missing or unknown DNN
	// 	-28: Unknown PDU session type
	// 	-29: User authentication or authorization failed
	// 	-31: Request rejected, unspecified
	// 	-32: Service option not supported
	// 	-33: Requested service option not subscribed
	// 	-35: PTI already in use
	// 	-36: Regular deactivation
	// 	-38: Network failure
	// 	-39: Reactivation requested
	// 	-43: Invalid PDU session identity
	// 	-44: Semantic errors in packet filter(s)
	// 	-45: Syntactical error in packet filter(s)
	// 	-46: Out of LADN service area
	// 	-47: PTI mismatch
	// 	-50: PDU session type IPv4 only allowed
	// 	-51: PDU session type IPv6 only allowed
	// 	-54: PDU session does not exist
	// 	-67: Insufficient resources for specific slice and DNN
	// 	-68: Not supported SSC mode
	// 	-69: Insufficient resources for specific slic
	// 	-70: Missing or unknown DNN in a slice
	// 	-81: Invalid PTI value
	// 	-82: Maximum data rate per UE for user-plane integrity protection is too low
	// 	-83: Semantic error in the QoS operation
	// 	-84: Syntactical error in the QoS operation
	// 	-85: Invalid mapped EPS bearer identity
	//  -91: DNN not supported or not subscribed in the slice
	// 	-96: invalid mandatory information
	// 	-97: Message type non-existent or not implemented
	// 	-98: Message type not compatible with protocol state
	// 	-99: Information element non-existent or not implemented
	// 	-100: Conditional IE error
	// 	-101: Message not compatible with protocol state
	// Default value nil.\n
	// optional.
	N4CauseMapping map[N4ErrorValue]Var5gSmCauseToUe `mapstructure:"n4CauseMapping" json:"n4CauseMapping,omitempty"`
	// Normally, if the Gy call flow fails, the GTPC cause that SMF returns to the SGWC can be determined according to the standard protocol
	// However, if GTPC Cause is required to decide based on returning a specific Gy diameter result-code, it needs to be supported by configuration
	// e.g:if Gy diameter result-code is 4012(credit-limit-reached), GTPC Cause code shall be 92(User authentication failed)
	// Key is Gy diameter result-code
	// Value is Gtpc Cause
	// Default value 0.\n
	// optional.
	GyCauseMapping map[DiameterErrorCode]GtpcCauseCode `mapstructure:"gyCauseMapping" json:"gyCauseMapping,omitempty"`
}

type DiameterErrorHandleProfile struct {
	// Gx/Gy interface failure cause
	// The value can be timeout or result-code, if it is timeout, ResultCode shall be absent, vice versa.
	// if it is omited, the failure handling action will be apply for all cause.
	// Optional
	FailureType FailureType `mapstructure:"failureType" json:"failureType,omitempty"`
	// Gx/Gy/S6bLite diameter error code, it is a code range. i.e. "5001-5010".If Need to config singe code,keep start the same with end, i.e."5001-5001"
	// if it is omited, the failure handling action will be apply in any error code.
	// Default value nil.\n
	// Optional
	Code *ResultCodeRange `mapstructure:"code" json:"code,omitempty"`
	// The Gx/Gy Message type
	// The value can be CCR-I/CCR-U/CCR-T/CCR-I-U
	// Optional
	MsgType MsgType `mapstructure:"msgType" json:"msgType,omitempty"`
	// The failure handle action
	// The value for Gx/Gy interface such as CONTINUE/RETRY_AND_TERMINATE/TERMINATE/DROP_PACKET/DROP_PACKET_AND_RETRY/ALLOW
	// if Gx/Gy interface happen error. CONTINUE means the smf will continue the request and the service will be granted.
	// RETRY_AND_TERMINATE means the smf will re-send the request to an alternative server in the case of transport temporary failures,
	// 		and the session will be terminate if retry all failed.
	// TERMINATE  means the smf will terminate session and procedure if send request failed.
	// DROP means the smf will set FAR apply action: drop
	// DROP_AND_RETRY means smf will set FAR apply action: drop and retry the request
	// ALLOW means remove URR for RG if send request failed.
	// The value for S6bLite interface such as RETRY/NO_RETRY
	// if S6bLite interface happen error. RETRY means retry the request
	// NO_RETRY means no need to retry the request
	// Mandatory
	Action FailureAction `mapstructure:"action" json:"action,omitempty"`
	// The Retries mean needing to retry in error handle. Only for the Action is DROP_AND_RETRY.
	// The Retries do not include the original msg. It means to retry the request when the original request failed.
	// Range: 0 to 10. If value is 0, means no retry, it can be seen as "DROP" action at this scene.
	// Default value 0.
	// Optional
	Retries int `mapstructure:"retries" json:"retries,omitempty"`
	// the interval for retry Action, such as DROP_AND_RETRY, between DROP_AND_RETRY failed request retries
	// unit is millisecond.
	// Range: 0 - 600000 \n
	// Default value 0.\n
	// Optional
	Interval int `mapstructure:"interval" json:"interval,omitempty"`
}

type ResultCodeRange struct {
	// The start of the range
	// Mandatory
	Start DiameterErrorCode `mapstructure:"start" json:"start"`
	// The end of the range
	// Mandatory
	End DiameterErrorCode `mapstructure:"end" json:"end"`
}

type MsgType string

const (
	CCRequestType_UPDATE_REQUEST         MsgType = "CCR-U"
	CCRequestType_INITIAL_REQUEST        MsgType = "CCR-I"
	CCRequestType_INITIAL_UPDATE_REQUEST MsgType = "CCR-I-U"
	CCRequestType_TERMINATION_REQUEST    MsgType = "CCR-T"
)

type FailureType string

const (
	FailureType_TIMEOUT FailureType = "timeout"
	// Deprecated: Use FailureType_CMD_RESULT_CODE instead
	FailureType_RESULT_CODE      FailureType = "result-code"
	FailureType_CMD_RESULT_CODE  FailureType = "cmd-result-code"
	FailureType_MSCC_RESULT_CODE FailureType = "mscc-result-code"
)

type FailureAction string

const (
	// for Gx/Gy interface
	FailureAction_CONTINUE            FailureAction = "CONTINUE"
	FailureAction_RETRY_AND_TERMINATE FailureAction = "RETRY_AND_TERMINATE"
	FailureAction_TERMINATE           FailureAction = "TERMINATE"
	FailureAction_DROP                FailureAction = "DROP"
	FailureAction_DROP_AND_RETRY      FailureAction = "DROP_AND_RETRY"
	FailureAction_ALLOW               FailureAction = "ALLOW"
	FailureAction_AP_RETRY            FailureAction = "AP_RETRY"

	// for S6bLite interface
	FailureAction_RETRY    FailureAction = "RETRY"
	FailureAction_NO_RETRY FailureAction = "NO_RETRY"
)
