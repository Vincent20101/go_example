package v1alpha2

import (
	"time"

	"gitlab.casa-systems.com/mobility/smf/infra/protocol/go-pfcp/pfcpie"
)

type SessDefaultQos struct {
	Var5qi          int32   `mapstructure:"5qi" json:"5qi"`
	Arp             Arp     `mapstructure:"arp" json:"arp"`
	PriorityLevel   *int32  `mapstructure:"priorityLevel" json:"priorityLevel,omitempty"`
	AverWindow      *int32  `mapstructure:"averWindow" json:"averWindow,omitempty"`
	MaxDataBurstVol *int32  `mapstructure:"maxDataBurstVol" json:"maxDataBurstVol,omitempty"`
	MaxbrUl         *string `json:"maxbrUl,omitempty"`
	MaxbrDl         *string `json:"maxbrDl,omitempty"`
	GbrUl           *string `json:"gbrUl,omitempty"`
	GbrDl           *string `json:"gbrDl,omitempty"`
	// Indicates whether notifications are requested from 3GPP NG-RAN when the GFBR can no longer (or again) be guaranteed for a QoS Flow during the lifetime of the QoS Flow.
	Qnc *bool `json:"qnc,omitempty"`
}

type SdSessionRule struct {
	// Authorized Session-AMBR
	AuthSessAmbr Ambr `mapstructure:"authSessAmbr" json:"authSessAmbr"`
	// Authorized default QoS information
	AuthDefQos AuthorizedDefaultQos `mapstructure:"authDefQos" json:"authDefQos"`
	// Univocally identifies the session rule within a PDU session.
	SessRuleId string `mapstructure:"sessRuleId" json:"sessRuleId"`
	// A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12.
	RefUmData string `mapstructure:"refUmData" json:"refUmData,omitempty"`
	// A reference to the condition data. It is the condId described in subclause 5.6.2.9.
	RefCondData string `mapstructure:"refCondData" json:"refCondData,omitempty"`
}

type SessRule struct {
	ActiveRuleName string                `mapstructure:"activeRuleName" json:"activeRuleName,omitempty"`
	ActiveCondName string                `mapstructure:"activeCondName" json:"activeCondName,omitempty"`
	QosRuleId      uint8                 `mapstructure:"qosRuleId" json:"qosRuleId"`
	AuthSessAmbr   *Ambr                 `mapstructure:"authSessAmbr" json:"authSessAmbr,omitempty"`
	AuthDefQos     *AuthorizedDefaultQos `mapstructure:"authDefQos" json:"authDefQos,omitempty"`
}

type AuthorizedDefaultQos struct {
	Var5qi          *int32 `mapstructure:"5qi" json:"5qi,omitempty"`
	Arp             *Arp   `mapstructure:"arp" json:"arp,omitempty"`
	PriorityLevel   *int32 `mapstructure:"priorityLevel" json:"priorityLevel,omitempty"`
	AverWindow      *int32 `mapstructure:"averWindow" json:"averWindow,omitempty"`
	MaxDataBurstVol *int32 `mapstructure:"maxDataBurstVol" json:"maxDataBurstVol,omitempty"`
	MaxbrUl         string `json:"maxbrUl,omitempty"`
	MaxbrDl         string `json:"maxbrDl,omitempty"`
	GbrUl           string `json:"gbrUl,omitempty"`
	GbrDl           string `json:"gbrDl,omitempty"`
	// Indicates whether notifications are requested from 3GPP NG-RAN when the GFBR can no longer (or again) be guaranteed for a QoS Flow during the lifetime of the QoS Flow.
	Qnc *bool `json:"qnc,omitempty"`
}

type SdPccRule struct {
	// An array of IP flow packet filter information.
	FlowInfos []SdFlowInformation `mapstructure:"flowInfos" json:"flowInfos,omitempty"`
	// A reference to the application detection filter configured at the UPF.
	AppId string `mapstructure:"appId" json:"appId,omitempty"`
	// Represents the content version of some content.
	ContVer *int32 `mapstructure:"contVer" json:"contVer,omitempty"`
	// Univocally identifies the PCC rule within a PDU session.
	PccRuleId string `mapstructure:"pccRuleId" json:"pccRuleId"`
	// Determines the order in which this PCC rule is applied relative to other PCC rules within the same PDU session.
	Precedence *int32 `mapstructure:"precedence" json:"precedence"`
	// Indicates the protocol used for signalling between the UE and the AF. The default value "NO_INFORMATION" shall apply, if the attribute is not present and has not been supplied previously.
	AfSigProtocol string `mapstructure:"afSigProtocol" json:"afSigProtocol,omitempty"`
	// Indication of application relocation possibility.
	AppReloc *bool `mapstructure:"appReloc" json:"appReloc,omitempty"`
	// A reference to the QoSData policy type decision type. It is the qosId described in subclause 5.6.2.8. (NOTE)
	RefQosData []string `mapstructure:"refQosData" json:"refQosData"`
	// A reference to the TrafficControlData policy decision type. It is the tcId described in subclause 5.6.2.10. (NOTE)
	RefTcData []string `mapstructure:"refTcData" json:"refTcData,omitempty"`
	// A reference to the ChargingData policy decision type. It is the chgId described in subclause 5.6.2.11. (NOTE)
	RefChgData []string `mapstructure:"refChgData" json:"refChgData,omitempty"`
	// A reference to UsageMonitoringData policy decision type. It is the umId described in subclause 5.6.2.12. (NOTE)
	RefUmData []string `mapstructure:"refUmData" json:"refUmData,omitempty"`
	// A reference to the condition data. It is the condId described in subclause 5.6.2.9.
	RefCondData string `mapstructure:"refCondData" json:"refCondData,omitempty"`
}

type SdFlowInformation struct {
	// Defines a packet filter for an IP flow.Refer to subclause 5.4.2 of 3GPP TS 29.212 [23] for encoding.
	FlowDescription string `mapstructure:"flowDescription" json:"flowDescription,omitempty"`
	// Defines a packet filter for an Ethernet flow. If the "fDir" attribute is included, it shall be set to "DOWNLINK". If the "fDir" attribute is never provided, the address information within the "ethFlowDescription" attribute shall be encoded in downlink direction.
	EthFlowDescription *EthFlowDescription `mapstructure:"ethFlowDescription" json:"ethFlowDescription,omitempty"`
	// An identifier of packet filter.
	PackFiltId string `mapstructure:"packFiltId" json:"packFiltId,omitempty"`
	// The packet shall be sent to the UE.
	PacketFilterUsage *bool `mapstructure:"packetFilterUsage" json:"packetFilterUsage,omitempty"`
	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass string `mapstructure:"tosTrafficClass" json:"tosTrafficClass,omitempty"`
	// The security parameter index of the IPSec packet.
	Spi string `mapstructure:"spi" json:"spi,omitempty"`
	// The Ipv6 flow label header field.
	FlowLabel string `mapstructure:"flowLabel" json:"flowLabel,omitempty"`
	// Indicates the direction/directions that a filter is applicable, downlink only, uplink only or both down- and uplink (bidirectional).
	FlowDirection string `mapstructure:"flowDirection" json:"flowDirection,omitempty"`
}

type SdTrafficControlData struct {
	// Univocally identifies the traffic control policy data within a PDU session.
	TcId string `mapstructure:"tcId" json:"tcId"`
	// Enum determining what action to perform on traffic. Possible values are: [enable, disable, enable_uplink, enable_downlink] . The default value "ENABLED" shall apply, if the attribute is not present and has not been supplied previously.
	FlowStatus string `mapstructure:"flowStatus" json:"flowStatus,omitempty"`
	// Each element indicates whether the detected application traffic should be redirected to another controlled address
	RedirectInfo *SdRedirectInformation `mapstructure:"redirectInfo" json:"redirectInfo,omitempty"`
	// Indicates whether applicat'on's start or stop notification is to be muted.
	MuteNotif bool `mapstructure:"muteNotif" json:"muteNotif,omitempty"`
	// Reference to a pre-configured traffic steering policy for downlink traffic at the SMF.
	TrafficSteeringPolIdDl string `mapstructure:"trafficSteeringPolIdDl" json:"trafficSteeringPolIdDl,omitempty"`
	// Reference to a pre-configured traffic steering policy for uplink traffic at the SMF.
	TrafficSteeringPolIdUl string `mapstructure:"trafficSteeringPolIdUl" json:"trafficSteeringPolIdUl,omitempty"`
	// A list of location which the traffic shall be routed to for the AF request
	RouteToLocs []RouteToLocation `mapstructure:"routeToLocs" json:"routeToLocs,omitempty"`
	// Contains the information about the AF subscriptions of the UP path change.
	UpPathChgEvent *UpPathChgEvent `mapstructure:"upPathChgEvent" json:"upPathChgEvent,omitempty"`
}

type SdRedirectInformation struct {
	// Indicates the redirect is enable.
	RedirectEnabled bool `mapstructure:"redirectEnabled" json:"redirectEnabled,omitempty"`
	// Indicates the type of redirect address.
	RedirectAddressType string `mapstructure:"redirectAddressType" json:"redirectAddressType,omitempty"`
	// Indicates the address of the redirect server.
	RedirectServerAddress string `mapstructure:"redirectServerAddress" json:"redirectServerAddress,omitempty"`
}

type RequestedRuleData struct {
	// An array of PCC rule id references to the PCC rules associated with the control data.
	RefPccRuleIds []string `mapstructure:"refPccRuleIds" json:"refPccRuleIds"`
	// Array of requested rule data type elements indicating what type of rule data is requested for the corresponding referenced PCC rules.
	ReqData []string `mapstructure:"reqData" json:"reqData"`
}

type RequestedUsageData struct {
	// An array of usage monitoring data id references to the usage monitoring data instances for which the PCF is requesting a usage report. This attribute shall only be provided when allUmIds is not set to true.
	RefUmIds []string `mapstructure:"refUmIds" json:"refUmIds,omitempty"`
	// Th ooleanean indicates whether requested usage data applies to all usage monitoring data instances. When it's not included, it means requested usage data shall only apply to the usage monitoring data instances referenced by the refUmIds attribute. default false
	AllUmIds bool `mapstructure:"allUmIds" json:"allUmIds,omitempty"`
}

type PendingPccReq struct {
	Pd                *SmPolicyDecision   `json:"pd,omitempty"`
	Op                *QosFlowBindOp      `json:"op,omitempty"`
	DetectedTriggers  DetectedTrigger     `json:"detectedTriggers,omitempty"` //store the current detected triggers that are necessary for the procedure handling
	N4_action         uint32              `json:"n4_action,omitempty"`        //pending n4 action flags.Might not need it! Need to see if we can recreate this from Op!
	PccRuleReport     []RuleReport        `json:"pccRuleReport,omitempty"`
	SessionRuleReport []SessionRuleReport `json:"sessionRuleReport,omitempty"`
}

type QosFlowBindOp struct {
	OpAction     QosOpAction  `json:"opAction,omitempty"`
	ReqOrigin    QosReqOrigin `json:"reqOrigin,omitempty"`
	QosFlows     []QosFlowOp  `json:"qosFlows,omitempty"`
	AuthSessAmbr *Ambr        `json:"authSessAmbr,omitempty"`
}

type QosFlowOp struct {
	Qfi           uint8       `json:"qfi"`
	Action        QosAction   `json:"action"`
	Failed        bool        `json:"failed"`
	PendingRules  []QosRuleOp `json:"pendingRules,omitempty"`
	PendingDefQos bool        `json:"pendingDefQos"`
}

type QosRuleOp struct {
	QosRuleId uint8     `json:"qosRuleId,omitempty"`
	PccRuleId string    `json:"pccRuleId,omitempty"`
	Action    QosAction `json:"action,omitempty"`
	SubAction uint32    `json:"subAction,omitempty"`
	UpfId     []string  `json:"upfId,omitempty"`
}

type PccRuleInfo struct {
	Rule    *PccRule             `json:"rule,omitempty"`
	QosData *SdQosData           `json:"qosData,omitempty"`
	ChgData *SessChargingData    `json:"chgData,omitempty"`
	TcData  *TrafficControlData  `json:"tcData,omitempty"`
	UmData  *UsageMonitoringData `json:"umData,omitempty"`
}

type QosAction int
type QosOpAction uint32
type QosReqOrigin int

type SmPolicyDecision struct {
	// A map of Sessionrules with the content being the SessionRule as described in subclause 5.6.2.7.
	SessRules map[string]SdSessionRule `mapstructure:"sessRules" json:"sessRules,omitempty"`
	// A map of PCC rules with the content being the PCCRule as described in subclause 5.6.2.6.
	PccRules map[string]SdPccRule `mapstructure:"pccRules" json:"pccRules,omitempty"`
	// If it is included and set to true, it indicates the P-CSCF Restoration is requested.
	PcscfRestIndication *bool `mapstructure:"pcscfRestIndication" json:"pcscfRestIndication,omitempty"`
	// Map of QoS data policy decisions.
	QosDecs map[string]SdQosData `mapstructure:"qosDecs" json:"qosDecs,omitempty"`
	// Map of Charging data policy decisions.
	ChgDecs map[string]SdChargingData `mapstructure:"chgDecs" json:"chgDecs,omitempty"`
	// Contains the CHF addresses of the PDU session.
	ChargingInfo *ChargingInformation `mapstructure:"chargingInfo" json:"chargingInfo,omitempty"`
	// Map of Traffic Control data policy decisions.
	TraffContDecs map[string]SdTrafficControlData `mapstructure:"traffContDecs" json:"traffContDecs,omitempty"`
	// Map of Usage Monitoring data policy decisions.
	UmDecs map[string]UsageMonitoringData `mapstructure:"umDecs" json:"umDecs,omitempty"`
	// Map of QoS characteristics for non standard 5QIs. This map uses the 5QI values as keys.
	QosChars map[string]QosCharacteristics `mapstructure:"qosChars" json:"qosChars,omitempty"`
	// Defines the lifetime of a UE derived QoS rule belonging to the PDU Session for reflective QoS.
	ReflectiveQoSTimer *int32 `mapstructure:"reflectiveQoSTimer" json:"reflectiveQoSTimer,omitempty"`
	// A map of condition data with the content being as described in subclause 5.6.2.9.
	Conds map[string]ConditionData `mapstructure:"conds" json:"conds,omitempty"`
	// Defines the time before which the SMF shall have to re-request PCC rules.
	RevalidationTime *time.Time `mapstructure:"revalidationTime" json:"revalidationTime,omitempty"`
	// Indicates the offline charging is applicable to the PDU session or PCC rule.
	Offline *bool `mapstructure:"offline" json:"offline,omitempty"`
	// Indicates the online charging is applicable to the PDU session or PCC rule.
	Online *bool `mapstructure:"online" json:"online,omitempty"`
	// Defines the policy control request triggers subscribed by the PCF.
	PolicyCtrlReqTriggers []string `mapstructure:"policyCtrlReqTriggers" json:"policyCtrlReqTriggers,omitempty"`
	// Defines the last list of rule control data requested by the PCF.
	LastReqRuleData []RequestedRuleData `mapstructure:"lastReqRuleData" json:"lastReqRuleData,omitempty"`
	// Defines the last requested usage data by the PCF.
	LastReqUsageData *RequestedUsageData `mapstructure:"lastReqUsageDataglobalNgenbId" json:"lastReqUsageData,omitempty"`
	// Map of PRA information.
	PraInfos map[string]PresenceInfoRm `mapstructure:"praInfos" json:"praInfos,omitempty"`
	// Information that identifies the IP address allocation method for IPv4 address allocation.
	Ipv4Index *int32 `mapstructure:"ipv4Index" json:"ipv4Index,omitempty"`
	// Information that identifies the IP address allocation method for IPv6 address allocation.
	Ipv6Index *int32 `mapstructure:"ipv6Index" json:"ipv6Index,omitempty"`
	// Indicates the required usage for default QoS flow
	QosFlowUsage string `mapstructure:"qosFlowUsage" json:"qosFlowUsage,omitempty"`
	// Indicates the list of negotiated supported features. This parameter shall be supplied by the PCF in the response to the POST request that requested the creation of an individual SM policy resource.
	SuppFeat string `mapstructure:"suppFeat" json:"suppFeat,omitempty"`
}

type SessFlow struct {
	DefFlow        FlowInfo      `json:"defFlow"`
	DedicatedFlows []*FlowInfo   `json:"dedicatedFlows,omitempty"`
	SessRule       SessRule      `json:"sessRule"`
	PccRules       []*PolicyRule `json:"pccRules,omitempty"`
	NextQfi        uint8         `json:"nextQfi"`
}

type FlowBinding struct {
	Var5qi          uint32 `json:"var5qi,omitempty"`
	Arp             Arp    `json:"arp,omitempty"`
	Qnc             bool   `json:"qnc,omitempty"`
	PriorityLevel   int32  `json:"priorityLevel,omitempty"`
	AverWindow      int32  `json:"averWindow,omitempty"`
	MaxDataBurstVol int32  `json:"maxDataBurstVol,omitempty"`
}

type QosFlowType uint8

const (
	QOS_FLOW_TYPE_NONE = iota
	QOS_FLOW_TYPE_GBR
	QOS_FLOW_TYPE_NON_GBR
	QOS_FLOW_TYPE_DELAY_CRITICAL_GBR
)

type AccessTypeNum uint8

const (
	ACCESS_TYPE_NONE = iota
	ACCESS_TYPE_3GPP
	ACCESS_TYPE_NON_3GPP
	ACCESS_TYPE_MULTI_ACCESS
)

type BearerTft struct {
	PktfilterIdMap map[uint8]string `json:"pktfilterIdMap,omitempty"` // key is filter ID, value is linked rule name
	NextFilterId   uint8            `json:"nextFilterId"`
}
type Dscp struct {
	Uplink   uint8 `json:"uplink"`   // key is filter ID, value is linked rule name
	Downlink uint8 `json:"downlink"` // key is filter ID, value is linked rule name
}

type FlowInfo struct {
	IsDefault       bool          `json:"isDefault"`
	Qfi             uint8         `json:"qfi"`
	GfbrUL          uint64        `json:"gfbrUL"` //guaranteed flow bitrate, unit kbps
	GfbrDL          uint64        `json:"gfbrDL"` //unit kbps
	MfbrUL          uint64        `json:"mfbrUL"` //max flow bitrate, unit kbps
	MfbrDL          uint64        `json:"mfbrDL"` //unit kbps
	MplrUl          int32         `json:"mplrUl"` // Max Packet Loss Rate (GBR only)
	MplrDl          int32         `json:"mplrDl"` // Max Packet Loss Rate (GBR only)
	Binding         FlowBinding   `json:"binding"`
	UpfId           []string      `json:"upfId,omitempty"`           //ip-address for UPFs (I-UPF and PSA UPFs)
	SdfRuleId       []uint8       `json:"sdfRuleId,omitempty"`       //pcc rules with sdf filters bound on this flow, use integer ID for efficiency
	AppRuleId       []uint8       `json:"appRuleId,omitempty"`       //pcc rules with app IDs bound on this flow, use integer ID for efficiency
	DataFwdAccepted bool          `json:"dataFwdAccepted,omitempty"` //For N2 Handover. We would need tp take forwarding decisions based on this
	Dscp            *Dscp         `json:"dscp,omitempty"`            // dscp value mapped from 5qi+arp_pl+ pl
	Ebi             int32         `json:"ebi,omitempty"`             //used for EPS interworking with N26,EBI allocation,TS 23502 Annex C (informative)
	QosFlowType     QosFlowType   `json:"qosFlowType,omitempty"`     // used for qos flow type, e.g gbr flow, non-gbr flow
	AccessType      AccessTypeNum `json:"accessType"`                //for MA PDU session, GBR flow can only have a single access at a given time, NON-GBR can have multiple access
	UPLocalTeid     *pfcpie.FTEID `json:"uPLocalTeid,omitempty"`     // user plane local teid for 4G NSA bearer, each has one tunnel
	UPRemoteTeid    *pfcpie.FTEID `json:"uPRemoteTeid,omitempty"`    // user plane peer sgw-u/epdg teid for 4G NSA bearer, each has one tunnel
	Tft             *BearerTft    `json:"tft,omitempty"`             // for 4G bearer
	ActivationTime  *time.Time    `json:"activationTime,omitempty"`  // Currently only applicable for calculating the duration of dedicated bearers
}

type PolicyRule struct {
	RefRuleID         string          `json:"refRuleID,omitempty"`
	QosRuleId         uint8           `json:"qosRuleId,omitempty"`
	Precedence        int32           `json:"precedence,omitempty"`
	Tp                RuleType        `json:"tp,omitempty"`
	PackFiltIds       []PackFiltIdMap `json:"packFiltIds,omitempty"`
	NasNextPackFiltId uint8           `json:"nasNextPackFiltId,omitempty"`
}

type PackFiltIdMap struct {
	NasId uint8  `json:"nasId"`
	PcfId string `json:"pcfId"`
}

type RuleType int

const (
	NONE_TYPE RuleType = iota
	SDF_RULE
	APP_RULE
)
