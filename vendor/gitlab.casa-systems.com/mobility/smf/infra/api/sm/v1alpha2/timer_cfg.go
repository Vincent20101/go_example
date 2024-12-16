/* timer config
 */

package v1alpha2

type timertype string

const (
	// TS24501 10.3.2
	TIMER_TYPE_T3590 timertype = "T3590"
	TIMER_TYPE_T3591           = "T3591"
	TIMER_TYPE_T3592           = "T3592"
	TIMER_TYPE_T3593           = "T3593"

	//dhcp timer
	TIMER_TYPE_DISC_REQ_TIMEOUT = "DISC_REQ_TIMEOUT"
	TIMER_TYPE_RENEW_V4         = "RENEW_V4"
	TIMER_TYPE_REBIND_V4        = "REBIND_V4"
	TIMER_TYPE_RELEASE_V4       = "RELEASE_V4"

	TIMER_TYPE_SOLICIT_REQ_TIMEOUT   = "SOLICIT_REQ_TIMEOUT"
	TIMER_TYPE_RENEW_V6              = "RENEW_V6"
	TIMER_TYPE_RENEW_REQ_V6_TIMEOUT  = "RENEW_REQ_V6_TIMEOUT"
	TIMER_TYPE_REBIND_V6             = "REBIND_V6"
	TIMER_TYPE_REBIND_REQ_V6_TIMEOUT = "REBIND_REQ_V6_TIMEOUT"
	TIMER_TYPE_RELEASE_V6            = "RELEASE_V6"

	//remove IUPF timer
	TIMER_TYPE_RELEASE_UPF             = "RELEASE_UPF"
	TIMER_TYPE_RELEASE_INDIRECT_TUNNEL = "RELEASE_INDIRECT_TUNNEL"

	// release sm context only in V-SMF/I-SMF, not to invoke resource release in H-SMF/SMF.
	TIMER_TYPE_RELEASE_SM_CONTEXT_VSMFISMF_ONLY = "RELEASE_SM_CONTEXT_VSMFISMF_ONLY"

	// If the CreateSmContext process fails (UpdateSmContext is not received), then initiate a release request
	//Deprecated //fix GCS-10659 the same time may cause a conflict,during PDU Session Establishment
	TIMER_TYPE_CREATE_SM_CONTEXT = "CREATE_SM_CONTEXT"

	// For various SM context update procedure failures(pending in wating state), shall revert back session state to active or idle
	// The active state revert triggered by a timer has been deprecated and enhanced to a passive state revert on state
	// update implemented by tryRevertStateAfterUpdateContextTimeout.
	TIMER_TYPE_UPDATE_SM_CONTEXT = "UPDATE_SM_CONTEXT"

	//If the first paging fails, the backoff timer will be started, and paging will be tried again if the timer expires
	TIMER_TYPE_RETRY_N1N2_PAGING = "RETRY_N1N2_PAGING"

	// Pra Stability Timer
	TIMER_TYPE_PRA_STABILITY = "PRA_STABILITY"

	// The SMF shall start the timer based on the revalidation time and shall send the PCC rule request before the indicated revalidation time.
	TIMER_TYPE_PCC_RULE_REVALIDATION = "PCC_RULE_REVALIDATION"

	// smf active the condition pcc rule
	TIMER_TYPE_ACTIVE_CONDITION_PCC_RULE = "ACTIVE_CONDITION_PCC_RULE"

	// smf deactive the condition pcc rule
	TIMER_TYPE_DEACTIVE_CONDITION_PCC_RULE = "DEACTIVE_CONDITION_PCC_RULE"

	// smf actice the condition session rule
	TIMER_TYPE_ACTIVE_CONDITION_SESSION_RULE = "ACTIVE_CONDITION_SESSION_RULE"

	// start the timer to modfify sdm subscription before it expires
	TIMER_TYPE_MODIFY_SDMSUBSCRIPTION = "MODIFY_SDMSUBSCRIPTION"

	// timer to retry charging request for gy error handling: DROP_AND_RETRY
	TIMER_TYPE_ERROR_HANDLE_CHARGING_RETRY = "ERROR_HANDLE_CHARGING_RETRY"

	// Verizon proprietary timer: sending offlline records to CHF.
	// refer to clause 'Verizon N40 spec Phase II - v2.02.doc' annex V2
	TIMER_TYPE_VZ_OFFLINE_SESSION_TIME = "VZ_OFFLINE_SESSION_TIME"

	// Verizon proprietary timer: sending all offline and all online reports to CHF.
	// refer to clause 'Verizon N40 spec Phase II - v2.02.doc' annex V2
	TIMER_TYPE_VZ_BCR_REPORTING_TIME = "VZ_BCR_REPORTING_TIME"

	//N2 TIMER
	TIMER_TYPE_N2_RES_SETUP_REQ = "N2_RES_SETUP_REQ"
	TIMER_TYPE_N2_RES_MOD_REQ   = "N2_RES_MOD_REQ"
	TIMER_TYPE_N2_RES_REL_COM   = "N2_RES_REL_COM"
	// Verizon proprietary timer: determines the number of seconds that the SMF waits for the Modify Bearer Request message during an EPC fallback.
	// refer to form '5G_LTE_Interface_Timer_Inventory-phase 3-0.4.xlsm' T5to4ho timer
	TIMER_TYPE_VZ_T5TO4HO = "VZ_T5TO4HO"

	// 4G to 5G HO timer for error scenario handling
	TIMER_TYPE_T4TO5HO = "T4TO5HO"

	// guard timer for NSA async create/update/delete bearer request
	// once timeout shall recover session state to ACTIVE
	// The active state revert triggered by a timer has been deprecated and enhanced to a passive state revert on state
	// update implemented by tryRevertStateAfterUpdateContextTimeout.
	TIMER_TYPE_BEARER_REQUEST_PENDING = "BEARER_REQUEST_PENDING"

	// SMF to Core release timer
	TIMER_TYPE_SMF_TO_CORE_RELEASE_TIME = "SMF_TO_CORE_RELEASE_TIME"

	// Upon reception of a rejection for an EPS bearer(s) PDN GW initiated procedure with an indication that the request
	// has been temporarily rejected due to mobility procedure in progress, the PDN GW start a locally configured guard
	// timer. The PDN GW shall re-attempt, up to a pre-configured number of times, when either it detects that the
	// Tracking Area Update procedure is completed or has failed using message reception or at expiry of the guard
	// timer.
	TIMER_TYPE_RE_ATTEMPT_BEARER_CREATION_MOD = "RE_ATTEMPT_BEARER_CREATION_MOD"

	// ePDG to 5G HO timer for error scenario handling
	TIMER_TYPE_EPDG_TO_5G = "EPDG_TO_5G"

	// The Session-Timeout attribute contains the maximum number of seconds of service to be provided to the user
	// before termination of the session.
	TIMER_TYPE_SESSION_TIMEOUT = "SESSION_TIMEOUT"

	// timer to retry ccr request for gx ap error handling: continue
	TIMER_TYPE_GX_ERROR_HANDLE_RETRY = "GX_ERROR_HANDLE_RETRY"
)

// Smf Timer configuration
//
//	Purpose:
//	  This is timer configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of smfSvcCfg
type TimerCfg struct {
	// Enable local timer strategy configuration.\n
	// By default, all timers use timer-svc to provide timing services.\n
	// Default value nil.
	// Optional
	LocalTimerCfg *LocalTimerCfg `mapstructure:"localTimerCfg"    json:"localTimerCfg,omitempty"`
}

// Smf local timer configuration
//
//	Purpose:
//	  This is local timer configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of TimerCfg
type LocalTimerCfg struct {
	// timer configuration using local timer.\n
	// Default value nil.\n
	// Optional
	Timers []timertype `mapstructure:"timers"    json:"timers,omitempty"`
}
