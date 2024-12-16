package v1alpha2

// Session Rule
//
//	Purpose:
//	  Defines Session Rule to be applied for a session.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.7 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and referenced by DefaultPccProfile under StaticPcc.
type SessionRule struct {
	// Authorized Session-AMBR.\n
	// Mandatory
	AuthSessAmbr Ambr `mapstructure:"authSessAmbr" json:"authSessAmbr"`
	// Authorized default QoS information.\n
	// Mandatory
	AuthDefQos DefaultQos `mapstructure:"authDefQos" json:"authDefQos"`
	// Identifier for this session rule.\n
	// Mandatory
	SessRuleId string `mapstructure:"sessRuleId" json:"sessRuleId"`
	// A reference to UsageMonitoringData configured under static pcc.\n
	// Optional
	RefUmData string `mapstructure:"refUmData" json:"refUmData,omitempty"`
	// A reference to the condition data configured under static pcc.\n
	// Optional
	RefCondData string `mapstructure:"refCondData" json:"refCondData,omitempty"`
}

// Ethernet Flow Description
//
//	Purpose:
//	  Defines a packet filter for an Ethernet flow.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/29514-g40.zip section 5.6.2.17 for details
//
//	Usage:
//	  Configured inside FlowInformation for PccRule
type EthFlowDescription struct {
	// Destination MAC address.
	DestMacAddr string `mapstructure:"destMacAddr" json:"destMacAddr,omitempty"`
	// A two-octet string that represents the Ethertype, as described in IEEE 802.3 [16] and IETF RFC 7042 [18] in hexadecimal representation. Each character in the string shall take a value of "0" to "9" or "A" to "F" and shall represent 4 bits. The most significant character representing the 4 most significant bits of the ethType shall appear first in the string, and the character representing the 4 least significant bits of the ethType shall appear last in the string.
	EthType string `mapstructure:"ethType" json:"ethType"`
	// Contains the flow description for the Uplink or Downlink IP flow. It shall be present when the Ethertype is IP.
	FDesc string `mapstructure:"fDesc" json:"fDesc,omitempty"`
	// Contains the packet filter direction. Only the "DOWNLINK" or "UPLINK" value is applicable.
	FDir string `mapstructure:"fDir" json:"fDir,omitempty"`
	// Source MAC address.
	SourceMacAddr string `mapstructure:"sourceMacAddr" json:"sourceMacAddr,omitempty"`
	// Customer-VLAN and/or Service-VLAN tags containing the VID, PCP/DEI fields as defined in IEEE 802.1Q [17] and IETF RFC 7042 [18]. The first/lower instance in the array stands for the Customer-VLAN tag and the second/higher instance in the array stands for the Service-VLAN tag.
	// Each field is encoded as a two-octet string in hexadecimal representation. Each character in the string shall take a value of "0" to "9" or "A" to "F" and shall represent 4 bits. The most significant character representing the PCP/DEI field shall appear first in the string, followed by character representing the 4 most significant bits of the VID field, and the character representing the 4 least significant bits of the VID field shall appear last in the string.
	// If only Service-VLAN tag is provided, empty string for Customer-VLAN tag shall be provided.
	VlanTags []string `mapstructure:"vlanTags" json:"vlanTags,omitempty"`
}

// Flow Information
//
//	Purpose:
//	  An array of IP flow packet filter information.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.14 for details
//
//	Usage:
//	  Used to configure packet filters under pcc rule.
type FlowInformation struct {
	// Defines a packet filter for an IP flow.Refer to subclause 5.4.2 of 3GPP TS 29.212 [23] for encoding.
	FlowDescription string `mapstructure:"flowDescription" json:"flowDescription,omitempty"`
	// Defines a packet filter for an Ethernet flow. If the "fDir" attribute is included, it shall be set to "DOWNLINK". If the "fDir" attribute is never provided, the address information within the "ethFlowDescription" attribute shall be encoded in downlink direction.
	EthFlowDescription *EthFlowDescription `mapstructure:"ethFlowDescription" json:"ethFlowDescription,omitempty"`
	// An identifier of packet filter.
	PackFiltId string `mapstructure:"packFiltId" json:"packFiltId,omitempty"`
	// The packet shall be sent to the UE.
	PacketFilterUsage bool `mapstructure:"packetFilterUsage" json:"packetFilterUsage,omitempty"`
	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and mask field.
	TosTrafficClass string `mapstructure:"tosTrafficClass" json:"tosTrafficClass,omitempty"`
	// The security parameter index of the IPSec packet.
	Spi string `mapstructure:"spi" json:"spi,omitempty"`
	// The Ipv6 flow label header field.
	FlowLabel string `mapstructure:"flowLabel" json:"flowLabel,omitempty"`
	// Indicates the direction/directions that a filter is applicable, downlink only, uplink only or both down- and uplink (bidirectional).
	FlowDirection string `mapstructure:"flowDirection" json:"flowDirection,omitempty"`
}

// Pcc Rule
//
//	Purpose:
//	  Defines Pcc Rule to be applied for qos flow in a session.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.6 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and referenced by Rulebase and DefaultPccRule.
type PccRule struct {
	// An array of IP flow packet filter information.\n
	// Optional
	FlowInfos []FlowInformation `mapstructure:"flowInfos" json:"flowInfos,omitempty"`
	// A reference to the application detection filter configured at the UPF.\n
	// Optional
	AppId string `mapstructure:"appId" json:"appId,omitempty"`
	// Represents the content version of pcc rule.\n
	// Optional
	ContVer *int32 `mapstructure:"contVer" json:"contVer,omitempty"`
	// Univocally identifies the PCC rule within a PDU session.\n
	// Mandatory
	PccRuleId string `mapstructure:"pccRuleId" json:"pccRuleId"`
	// Determines the order in which this PCC rule is applied relative to other PCC rules within the same PDU session.\n
	// Optional
	Precedence *int32 `mapstructure:"precedence" json:"precedence"`
	// Indicates the protocol used for signalling between the UE and the AF. The default value "NO_INFORMATION" shall apply, if the attribute is not present and has not been supplied previously.\n
	// Optional
	AfSigProtocol string `mapstructure:"afSigProtocol" json:"afSigProtocol,omitempty"`
	// Indication of application relocation possibility.\n
	// Optional
	AppReloc *bool `mapstructure:"appReloc" json:"appReloc,omitempty"`
	// A reference the QosData policy decision type(QosId)  configured in QosDecs map under static pcc to which this Pcc Rule applies.\n
	// Optional
	RefQosData []string `mapstructure:"refQosData" json:"refQosData"`
	// A reference to the TrafficControlData policy decision type(TcId) configured in TraffContDecs map under static pcc.\n
	// Optional
	RefTcData []string `mapstructure:"refTcData" json:"refTcData,omitempty"`
	// A reference to the ChargingData policy decision type(ChgId) configured in ChgDecs map under static pcc.\n
	// Optional
	RefChgData []string `mapstructure:"refChgData" json:"refChgData,omitempty"`
	// A reference to UsageMonitoringData policy decision type(umId) configured in UmDecs map under static pcc.\n
	// Optional
	RefUmData []string `mapstructure:"refUmData" json:"refUmData,omitempty"`
	// A reference to the condition data(condId) configured in conds map under static pcc.\n
	// Optional
	RefCondData string `mapstructure:"refCondData" json:"refCondData,omitempty"`
	// References to the HeaderEnrichmentCfg configured in StaticPccMap to support multiple header enrichment
	// Optional
	// To be deprecated
	HeaderEnrichmentRefs []string `mapstructure:"headerEnrichmentRefs" json:"headerEnrichmentRefs,omitempty"`
	// References to the HeaderFieldCfg configured in StaticPccMap to support multiple header enrichment for UL
	// Optional
	HeaderEnrichmentRefsUL []string `mapstructure:"headerEnrichmentRefsUL" json:"headerEnrichmentRefsUL,omitempty"`
	// References to the HeaderFieldCfg configured in StaticPccMap to support multiple header enrichment for DL
	// Optional
	HeaderEnrichmentRefsDL []string `mapstructure:"headerEnrichmentRefsDL" json:"headerEnrichmentRefsDL,omitempty"`
	// A reference to the XHeaderCfg configured in StaticPccMap
	// Optional
	// Deprecated effective rel 7.4
	XHeaderRef string `mapstructure:"xHeaderRef" json:"xHeaderRef,omitempty"`
	// A reference to the TlsExtensionCfg configured in StaticPccMap
	// Optional
	// Deprecated effective rel 7.4
	TlsExtensionRef string `mapstructure:"tlsExtensionRef" json:"tlsExtensionRef,omitempty"`
	// SMF doesnt add encrypted user IDs. Only send usedID mask to UPF in Forwarding Parameter IE of FAR. \n
	// Value can be "IMSI"/"IMEI"/"MSISDN". \n
	// Can set one value or multiple values. \n
	// - 1)just set one value. E.g:  RedirectUserIdMask:"IMSI"
	// - 2) To set multiple values, need to use "|" to separate. E.g: RedirectUserIdMask:"IMSI|MSISDN"
	// Optional.
	RedirectUserIdMask string `mapstructure:"redirectUserIdMask" json:"redirectUserIdMask,omitempty"`
}

// Arp (Allocation And Retention Priority)
//
//	Purpose:
//	  Defines allocation and retention priority for the qos flow.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.4.1 for details
//
//	Usage:
//	  Configured under Default Qos and QosData.
type Arp struct {
	// Defines the relative importance of a resource request.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/23_series/23.501/23501-g40.zip section 5.7.2.2.\n
	// Allowed value: 1-15.\n
	// Mandatory
	PriorityLevel int32 `mapstructure:"priorityLevel" json:"priorityLevel"`
	// Defines whether a service data flow may get resources that were already assigned to another service data flow with a lower priority level.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.3.1 Table 5.5.3.1-1.\n
	// -"NOT_PREEMPT":Shall not trigger pre-emption.\n
	// -"MAY_PREEMPT":May trigger pre-emption.\n
	// Mandatory
	PreemptCap string `mapstructure:"preemptCap" json:"preemptCap,omitempty"`
	// Defines whether a service data flow may lose the resources assigned to it in order to admit a service data flow with higher priority level.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.3.2 Table 5.5.3.2-1.\n
	// -"NOT_PREEMPTABLE":Shall not be pre-empted.\n
	// -"PREEMPTABLE":May be pre-empted.\n
	// Mandatory
	PreemptVuln string `mapstructure:"preemptVuln" json:"preemptVuln,omitempty"`
}

type ViolateActionType string

// This IE shall be present if an action need to be taken if enforcement requirements (MBR, GBR, Burst Size) are violated
const (
	ViolateActionNone              ViolateActionType = "none"
	ViolateActionDiscard           ViolateActionType = "discard"
	ViolateActionLowerIpPrecedence ViolateActionType = "lower-ip-precedence"
)

// Qos Data
//
//	Purpose:
//	  Describes Qos information associated to a flow.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.8 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and referenced by PccRule.
type QosData struct {
	// Univocally identifies the QoS control policy data within a PDU session.\n
	// Mandatory
	QosId string `mapstructure:"qosId" json:"qosId"`
	// Identifier for the authorized QoS parameters for the service data flow. It shall be included when the QoS data decision is initially provisioned and "defQosFlowIndication" is not included or is included and set to false.\n
	// Optional
	Var5qi int32 `mapstructure:"var5qi" json:"var5qi"`
	// Indicates the max bandwidth in uplink.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// Pattern: '^\d+(\.\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$'\n
	// Examples:"125 Mbps", "0.125 Gbps", "125000 Kbps"\n
	// Optional
	MaxbrUl string `mapstructure:"maxbrUl" json:"maxbrUl,omitempty"`
	// Indicates the max bandwidth in downlink.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// Pattern: '^\d+(\.\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$'\n
	// Examples:"125 Mbps", "0.125 Gbps", "125000 Kbps"\n
	// Optional
	MaxbrDl string `mapstructure:"maxbrDl" json:"maxbrDl,omitempty"`
	// Indicates the guaranteed bandwidth in uplink.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// Pattern: '^\d+(\.\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$'\n
	// Examples:"125 Mbps", "0.125 Gbps", "125000 Kbps"\n
	// Optional
	GbrUl string `mapstructure:"gbrUl" json:"gbrUl,omitempty"`
	// Indicates the guaranteed bandwidth in downlink.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// Pattern: '^\d+(\.\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$'\n
	// Examples:"125 Mbps", "0.125 Gbps", "125000 Kbps"\n
	// Optional
	GbrDl string `mapstructure:"gbrDl" json:"gbrDl,omitempty"`
	// Indicates the allocation and retention priority. It shall be included when the QoS data decision is initially provisioned and "defQosFlowIndication" is not included or is included and set to false.\n
	// Optional
	Arp Arp `mapstructure:"arp" json:"arp"`
	// Indicates whether notifications are requested from 3GPP NG-RAN when the GFBR can no longer (or again) be guaranteed for a QoS Flow during the lifetime of the QoS Flow.\n
	// Optional
	Qnc bool `mapstructure:"qnc" json:"qnc,omitempty"`
	// Indicates a priority in scheduling resources among QoS Flows.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// Allowed value:1 to 127.\n
	// Optional
	PriorityLevel *int32 `mapstructure:"priorityLevel" json:"priorityLevel,omitempty"`
	// Represents the duration over which the guaranteed and maximum bitrate shall be calculated\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// expressed in milliseconds.\n
	// Minimum = 1. Maximum = 4095.\n
	// Optional
	AverWindow *int32 `mapstructure:"averWindow" json:"averWindow,omitempty"`
	// Denotes the largest amount of data that is required to be transferred within a period of 5G-AN PDB.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// expressed in Bytes.\n
	// Minimum = 1. Maximum = 4095.\n
	// Optional
	MaxDataBurstVol *int32 `mapstructure:"maxDataBurstVol" json:"maxDataBurstVol,omitempty"`
	// Indicates whether the QoS information is reflective for the corresponding service data flow.\n
	// Optional
	ReflectiveQos bool `mapstructure:"reflectiveQos" json:"reflectiveQos,omitempty"`
	// Indicates, by containing the same value, what PCC rules may share resource in downlink direction.\n
	// Optional
	SharingKeyDl string `mapstructure:"sharingKeyDl" json:"sharingKeyDl,omitempty"`
	// Indicates, by containing the same value, what PCC rules may share resource in uplink direction.\n
	// Optional
	SharingKeyUl string `mapstructure:"sharingKeyUl" json:"sharingKeyUl,omitempty"`
	// Indicates the downlink maximum rate for lost packets that can be tolerated for the service data flow.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// expressed in tenth of percent.\n
	// Minimum = 0. Maximum = 1000.\n
	// Optional
	MaxPacketLossRateDl *int32 `mapstructure:"maxPacketLossRateDl" json:"maxPacketLossRateDl,omitempty"`
	// Indicates the uplink maximum rate for lost packets that can be tolerated for the service data flow.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// expressed in tenth of percent.\n
	// Minimum = 0. Maximum = 1000.\n
	// Optional
	MaxPacketLossRateUl *int32 `mapstructure:"maxPacketLossRateUl" json:"maxPacketLossRateUl,omitempty"`
	// Indicates that the dynamic PCC rule shall always have its binding with the QoS Flow associated with the default QoS rule.\n
	// Optional
	DefQosFlowIndication bool `mapstructure:"defQosFlowIndication" json:"defQosFlowIndication,omitempty"`
	// Indicate the uplink peak burst size to be enforced for packets matching the PDR. \n
	// Unit is byte. \n
	// Default value is nil.\n
	// Optional
	UlBurstSize *int64 `mapstructure:"ulBurstSize" json:"ulBurstSize,omitempty"`
	// Indicate the downlink peak burst size to be enforced for packets matching the PDR. \n
	// Unit is byte. \n
	// Default value is nil.\n
	// Optional
	DlBurstSize *int64 `mapstructure:"dlBurstSize" json:"dlBurstSize,omitempty"`
	// An action need to be taken if enforcement requirements (MBR, GBR, Burst Size) are violated. \n
	// Default value is nil. \n
	// Optional
	ViolateAction *ViolateActionType `mapstructure:"violateAction" json:"violateAction,omitempty"`
}

// Charging Information configuration
//
//  Purpose:
//    This field specifies primary and secondary CHF addresses to be used by SMF for sending charging request to.
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Used as mandatory parameter "ChargingInfo" in ChargingCharacteristics and optional parameter "ChargingInfo" in SmPolicyDecision
//

type ChargingInformation struct {
	// Contains the primary CHF address in format of ip:port.\n
	// Port 80 will be used as default port when port is absent.\n
	// Default value "".\n
	// Mandatory.
	PrimaryChfAddress string `mapstructure:"primaryChfAddress" json:"primaryChfAddress"`
	// Contains the secondary CHF address in format of ip:port.\n.
	// Port 80 will be used as default port when port is absent.\n
	// Default value "".\n
	// Mandatory.
	SecondaryChfAddress string `mapstructure:"secondaryChfAddress" json:"secondaryChfAddress"`
}

type IdentityEmbeddingCfg struct {
	// Indicates whether some user identities are attached after the redirect URL, used for UPF to find the identities info
	// Allows the user to set any tag, e.g: tag value is "encryptedInfo" or "encToken"
	// Default value "encryptedInfo"
	// Optional
	RedirectRuleTag string `mapstructure:"redirectRuleTag" json:"redirectRuleTag,omitempty"`
	// Indicates a list of user identities to be embedded, separated by "|" if more than one included
	// Currently the follow 4 kinds of identities are supported: IMSI, IMEI, MSISDN, NAI
	// Default value ""
	// Optional
	RedirectUserIdMask string `mapstructure:"redirectUserIdMask" json:"redirectUserIdMask,omitempty"`
	// Indicates the delimiter between fields in the redirect info
	// Default value ";"
	// Optional
	RedirectFieldDelimiter string `mapstructure:"redirectFieldDelimiter" json:"redirectFieldDelimiter,omitempty"`
}

// Redirect Information
//
//	Purpose:
//	  It indicates whether the detected application traffic should be redirected to another controlled address.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.13 for details
//
//	Usage:
//	  Configured under TrafficControlData for PccRule
type RedirectInformation struct {
	// Indicates the redirect is enable.\n
	// Optional
	RedirectEnabled bool `mapstructure:"redirectEnabled" json:"redirectEnabled,omitempty"`
	// Indicates the type of redirect address.\n
	// Possible values are(refer TS 29512-5.6.3.12)
	// - IPV4_ADDR: Indicates that the address type is in the form of \"dotted-decimal\" IPv4 address.
	// - IPV6_ADDR: Indicates that the address type is in the form of IPv6 address.
	// - URL: Indicates that the address type is in the form of Uniform Resource Locator.
	// - SIP_URI: Indicates that the address type is in the form of SIP Uniform Resource Identifier.
	// Optional
	RedirectAddressType string `mapstructure:"redirectAddressType" json:"redirectAddressType,omitempty"`
	// Indicates the address of the redirect server.\n
	// Optional
	RedirectServerAddress string `mapstructure:"redirectServerAddress" json:"redirectServerAddress,omitempty"`
	// Indicates whether/how to embed identities (refer to IdentityEmbeddingCfg struct above for details)
	// Default value nil (indicating no identity to be embedded)
	// Optional
	IdentityEmbeddingCfg *IdentityEmbeddingCfg `mapstructure:"identityEmbeddingCfg" json:"identityEmbeddingCfg,omitempty"`
}

// Route Information
//
//	Purpose:
//	  Includes traffic routing information.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.4.4.16 for details
//
//	Usage:
//	  Configured under RouteToLocation in TrafficControlData for PccRule
type RouteInformation struct {
	// Ipv4address of the tunnel end point in the data network.
	Ipv4Addr string `mapstructure:"ipv4Addr" json:"ipv4Addr,omitempty"`
	// Ipv6 address of the tunnel end point in the data network.
	Ipv6Addr string `mapstructure:"ipv6Addr" json:"ipv6Addr,omitempty"`
	// UDP port number of the tunnel end point in the data network.
	PortNumber int32 `mapstructure:"portNumber" json:"portNumber"`
}

// RouteToLocation
//
//	Purpose:
//	  Location to which the traffic shall be routed to for the AF request.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.4.4.15 for details
//
//	Usage:
//	  Configured under TrafficControlData for PccRule
type RouteToLocation struct {
	// Identifies the location of the application.\n
	// Mandatory
	Dnai string `mapstructure:"dnai" json:"dnai"`
	// Includes the traffic routing information.\n
	// Optional
	RouteInfo *RouteInformation `mapstructure:"routeInfo" json:"routeInfo,omitempty"`
	// Identifies the routing profile Id.\n
	// Optional
	RouteProfId string `mapstructure:"routeProfId" json:"routeProfId,omitempty"`
}

// UpPathChgEvent
//
//	Purpose:
//	  Contains the information about the AF subscriptions of the UP path change.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.20 for details
//
//	Usage:
//	  Configured under TrafficControlData for PccRule
type UpPathChgEvent struct {
	// Notification address of AF receiving the event notification.\n
	// Mandatory
	NotificationUri string `mapstructure:"notificationUri" json:"notificationUri"`
	// It is used to set the value of Notification Correlation ID in the notification sent by the SMF.\n
	// Mandatory
	NotifCorreId string `mapstructure:"notifCorreId" json:"notifCorreId"`
	// Indicates the type of DNAI change.\n
	// Mandatory
	DnaiChgType string `mapstructure:"dnaiChgType" json:"dnaiChgType"`
}

// Traffic Control Data
//
//	Purpose:
//	  Describes traffic control information like flow status, traffic steering policy etc associated to a flow.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.10 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and referenced by PccRule.
type TrafficControlData struct {
	// Univocally identifies the traffic control policy data within a PDU session.\n
	// Mandatory
	TcId string `mapstructure:"tcId" json:"tcId"`
	// Enum determining what action to perform on traffic.\n
	// Allowed values:[enable, disable, enable_uplink, enable_downlink，removed].\n
	// Optional
	FlowStatus string `mapstructure:"flowStatus" json:"flowStatus,omitempty"`
	// Each element indicates whether the detected application traffic should be redirected to another controlled address.\n
	// Optional
	RedirectInfo *RedirectInformation `mapstructure:"redirectInfo" json:"redirectInfo,omitempty"`
	// Indicates whether applicat'on's start or stop notification is to be muted.\n
	// Optional
	MuteNotif bool `mapstructure:"muteNotif" json:"muteNotif,omitempty"`
	// Reference to a pre-configured traffic steering policy for downlink traffic at the SMF.\n
	// Optional
	TrafficSteeringPolIdDl string `mapstructure:"trafficSteeringPolIdDl" json:"trafficSteeringPolIdDl,omitempty"`
	// Reference to a pre-configured traffic steering policy for uplink traffic at the SMF.\n
	// Optional
	TrafficSteeringPolIdUl string `mapstructure:"trafficSteeringPolIdUl" json:"trafficSteeringPolIdUl,omitempty"`
	// A list of location which the traffic shall be routed to for the AF request.\n
	// Optional
	RouteToLocs []RouteToLocation `mapstructure:"routeToLocs" json:"routeToLocs,omitempty"`
	// Contains the information about the AF subscriptions of the UP path change.\n
	// Optional
	UpPathChgEvent *UpPathChgEvent `mapstructure:"upPathChgEvent" json:"upPathChgEvent,omitempty"`
}

// Usage Monitoring Data
//
//	Purpose:
//	  Describes the usage monitoring to be enabled for a flow.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.12 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and referenced by PccRule.
type UsageMonitoringData struct {
	// Univocally identifies the usage monitoring policy data within a PDU session.\n
	// Mandatory
	UmId string `mapstructure:"umId" json:"umId"`
	// Indicates the total volume threshold.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/29122-g50.zip section 5.2.1.3.2 Table 5.2.1.3.2-2.\n
	// units of bytes.\n
	// Optional
	VolumeThreshold *int32 `mapstructure:"volumeThreshold" json:"volumeThreshold,omitempty"`
	// Indicates a volume threshold in uplink.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/29122-g50.zip section 5.2.1.3.2 Table 5.2.1.3.2-2.\n
	// units of bytes.\n
	// Optional
	VolumeThresholdUplink *int32 `mapstructure:"volumeThresholdUplink" json:"volumeThresholdUplink,omitempty"`
	// Indicates a volume threshold in downlink.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/29122-g50.zip section 5.2.1.3.2 Table 5.2.1.3.2-2.\n
	// units of bytes.\n
	// Optional
	VolumeThresholdDownlink *int32 `mapstructure:"volumeThresholdDownlink" json:"volumeThresholdDownlink,omitempty"`
	// Indicates a time threshold.\n
	// Optional
	TimeThreshold *int32 `mapstructure:"timeThreshold" json:"timeThreshold,omitempty"`
	// Indicates the time at which the UP function is expected to reapply the next thresholds (e.g. nextVolThreshold).\n
	// Optional
	MonitoringTime string `mapstructure:"monitoringTime" json:"monitoringTime,omitempty"`
	// Indicates a volume threshold after the Monitoring Time.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/29122-g50.zip section 5.2.1.3.2 Table 5.2.1.3.2-2.\n
	// units of bytes.\n
	// Optional
	NextVolThreshold *int32 `mapstructure:"nextVolThreshold" json:"nextVolThreshold,omitempty"`
	// Indicates a volume threshold in uplink after the Monitoring Time.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/29122-g50.zip section 5.2.1.3.2 Table 5.2.1.3.2-2.\n
	// units of bytes.\n
	// Optional
	NextVolThresholdUplink *int32 `mapstructure:"nextVolThresholdUplink" json:"nextVolThresholdUplink,omitempty"`
	// Indicates al volume threshold in downlink after the Monitoring Time.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/29122-g50.zip section 5.2.1.3.2 Table 5.2.1.3.2-2.\n
	// units of bytes.\n
	// Optional
	NextVolThresholdDownlink *int32 `mapstructure:"nextVolThresholdDownlink" json:"nextVolThresholdDownlink,omitempty"`
	// Indicates a time threshold after the Monitoring.\n
	// Optional
	NextTimeThreshold *int32 `mapstructure:"nextTimeThreshold" json:"nextTimeThreshold,omitempty"`
	// Defines the period of time after which the time measurement shall stop, if no packets are received.\n
	// Optional
	InactivityTime *int32 `mapstructure:"inactivityTime" json:"inactivityTime,omitempty"`
	// Contains the PCC rule identifier(s) which corresponding service data flow(s) shall be excluded from PDU Session usage monitoring. It is only included in the UsageMonitoringData instance for session level usage monitoring.\n
	// Optional
	ExUsagePccRuleIds []string `mapstructure:"exUsagePccRuleIds" json:"exUsagePccRuleIds,omitempty"`
}

// Condition Data
//
//	Purpose:
//	  Describes the condition data for enabling session rule.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.9 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and referenced by SessRule. Whenever described condition is met, session rule associated to condition data is activated if available. Otherwise default session rule without condition data would be activated.
type ConditionData struct {
	// Uniquely identifies the condition data within a PDU session.\n
	// Mandatory
	CondId string `mapstructure:"condId" json:"condId"`
	// The time when the decision data shall be activated.\n
	// Optional
	ActivationTime string `mapstructure:"activationTime" json:"activationTime,omitempty"`
	// The time when the decision data shall be deactivated.\n
	// Optional
	DeactivationTime string `mapstructure:"deactivationTime" json:"deactivationTime,omitempty"`
	// Access type condition data for rule activation.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.4.3.1 Table 5.4.3.1-1.\n
	// -"3GPP_ACCESS":3GPP access.\n
	// -"NON_3GPP_ACCESS":Non-3GPP access.\n
	// Optional
	AccessType AccessType `mapstructure:"accessType" json:"accessType,omitempty"`
	// RAT type condition data for rule activation.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.4.3.2 Table 5.4.3.2-1.\n
	// Optional
	RatType RatType `mapstructure:"ratType" json:"ratType,omitempty"`
}

type TriggerCategory = string

const (
	TriggerCategory_IMMEDIATE_REPORT TriggerCategory = "IMMEDIATE_REPORT"
	TriggerCategory_DEFERRED_REPORT  TriggerCategory = "DEFERRED_REPORT"
)

// Trigger configuration
//
//	Purpose:
//	  This field specifies the reason for usage reporting for one or more types of unit.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.1.7-1.
//
//	Usage:
//	  Used as optioanl parameter "Triggers" in ChargingCharacteristics
type Trigger struct {
	// The events whose occurrence lead to charging event is issued towards the CHF
	Type TriggerType `mapstructure:"type" json:"type"`
	// This field indicates whether the charging data generated by the NF consumer for the trigger lead to a Charging Event towards the CHF immediately or not.
	Category TriggerCategory `mapstructure:"category" json:"category,omitempty"`
	// Time limit if trigger type is "Expiry of data time limit"
	TimeLimit int32 `mapstructure:"timeLimit" json:"timeLimit,omitempty"`
	// Volume limit if trigger type is "Expiry of data volume limit". This attribute is not valid from Nchf_ConvergedCharging API version v2.0.0
	VolumeLimit int32 `mapstructure:"volumeLimit" json:"volumeLimit,omitempty"`
	// Maximum number if trigger type is "Max nb of number of charging condition changes"
	MaxNumberOfccc int32 `mapstructure:"maxNumberOfccc" json:"maxNumberOfccc,omitempty"`
}

type TrafficTime struct {
	// TBD
	StartTime uint32 `mapstructure:"startTime" json:"startTime"`
	// TBD
	EndTime uint32 `mapstructure:"endTime" json:"endTime"`
}

type TriggerType string

const (
	TriggerType_QUOTA_THRESHOLD                                  TriggerType = "QUOTA_THRESHOLD"
	TriggerType_QHT                                              TriggerType = "QHT"
	TriggerType_FINAL                                            TriggerType = "FINAL"
	TriggerType_QUOTA_EXHAUSTED                                  TriggerType = "QUOTA_EXHAUSTED"
	TriggerType_VALIDITY_TIME                                    TriggerType = "VALIDITY_TIME"
	TriggerType_OTHER_QUOTA_TYPE                                 TriggerType = "OTHER_QUOTA_TYPE"
	TriggerType_FORCED_REAUTHORISATION                           TriggerType = "FORCED_REAUTHORISATION"
	TriggerType_UNUSED_QUOTA_TIMER                               TriggerType = "UNUSED_QUOTA_TIMER"
	TriggerType_ABNORMAL_RELEASE                                 TriggerType = "ABNORMAL_RELEASE"
	TriggerType_QOS_CHANGE                                       TriggerType = "QOS_CHANGE"
	TriggerType_VOLUME_LIMIT                                     TriggerType = "VOLUME_LIMIT"
	TriggerType_TIME_LIMIT                                       TriggerType = "TIME_LIMIT"
	TriggerType_PLMN_CHANGE                                      TriggerType = "PLMN_CHANGE"
	TriggerType_USER_LOCATION_CHANGE                             TriggerType = "USER_LOCATION_CHANGE"
	TriggerType_RAT_CHANGE                                       TriggerType = "RAT_CHANGE"
	TriggerType_UE_TIMEZONE_CHANGE                               TriggerType = "UE_TIMEZONE_CHANGE"
	TriggerType_TARIFF_TIME_CHANGE                               TriggerType = "TARIFF_TIME_CHANGE"
	TriggerType_MAX_NUMBER_OF_CHANGES_IN_CHARGING_CONDITIONS     TriggerType = "MAX_NUMBER_OF_CHANGES_IN_CHARGING_CONDITIONS"
	TriggerType_MANAGEMENT_INTERVENTION                          TriggerType = "MANAGEMENT_INTERVENTION"
	TriggerType_CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA TriggerType = "CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA"
	TriggerType_CHANGE_OF_3GPP_PS_DATA_OFF_STATUS                TriggerType = "CHANGE_OF_3GPP_PS_DATA_OFF_STATUS"
	TriggerType_SERVING_NODE_CHANGE                              TriggerType = "SERVING_NODE_CHANGE"
	TriggerType_REMOVAL_OF_UPF                                   TriggerType = "REMOVAL_OF_UPF"
	TriggerType_ADDITION_OF_UPF                                  TriggerType = "ADDITION_OF_UPF"
	TriggerType_ECGI_CHANGE                                      TriggerType = "ECGI_CHANGE"
	TriggerType_TAI_CHANGE                                       TriggerType = "TAI_CHANGE"
	TriggerType_SESSION_AMBR_CHANGE                              TriggerType = "SESSION_AMBR_CHANGE"
	TriggerType_TRAFFIC_FLOW_TEMPLATE_CHANGE                     TriggerType = "TRAFFIC_FLOW_TEMPLATE_CHANGE"
	TriggerType_UNIT_COUNT_INACTIVITY_TIMER                      TriggerType = "UNIT_COUNT_INACTIVITY_TIMER"
	TriggerType_CGI_SAI_CHANGE                                   TriggerType = "CGI_SAI_CHANGE"
)

type NetworkSlicingInfo struct {
	//Single network slice selection assistance information
	SNSSAI Snssai `mapstructure:"sNSSAI" json:"sNSSAI"`
}

type Model3GpppsDataOffStatus = string

const (
	Model3GpppsDataOffStatus_ACTIVE   Model3GpppsDataOffStatus = "ACTIVE"
	Model3GpppsDataOffStatus_INACTIVE Model3GpppsDataOffStatus = "INACTIVE"
)

type ChargingCharacteristicsSelectionMode = string

const (
	ChargingCharacteristicsSelectionMode_HOME_DEFAULT     ChargingCharacteristicsSelectionMode = "HOME_DEFAULT"
	ChargingCharacteristicsSelectionMode_ROAMING_DEFAULT  ChargingCharacteristicsSelectionMode = "ROAMING_DEFAULT"
	ChargingCharacteristicsSelectionMode_VISITING_DEFAULT ChargingCharacteristicsSelectionMode = "VISITING_DEFAULT"
)

type SscMode = string

const (
	SscMode_SSC_MODE_1 SscMode = "SSC_MODE_1"
	SscMode_SSC_MODE_2 SscMode = "SSC_MODE_2"
	SscMode_SSC_MODE_3 SscMode = "SSC_MODE_3"
)

// NonDynamic5Qi
//
//	Purpose:
//	  Defines the QoS Characteristics for a standardized or pre-configured 5QI for downlink and uplink.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.4.4 for details
//
//	Usage:
//	  Configured inside VsmfQosProfile and DefaultQosInformation
type NonDynamic5Qi struct {
	// Defines the 5QI Priority Level.When present, it contains the 5QI Priority Level value that overrides the standardized or pre-configured value.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// Allowed Value: 1-127.\n
	// Optional
	PriorityLevel *int32 `mapstructure:"priorityLevel" json:"priorityLevel,omitempty"`
	// Defines the averaging window. This IE may be present for a GBR QoS flow or a Delay Critical GBR QoS flow. When present, it contains the Averaging Window that overrides the standardized or pre-configured value.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in milliseconds.\n
	// Allowed Value:1-4095.\n
	// Optional
	AverWindow *int32 `mapstructure:"averWindow" json:"averWindow,omitempty"`
	// Defines the maximum data burst volume.This IE may be present for a Delay Critical GBR QoS flow. When present, it contains the Maximum Data Burst Volume value that overrides the standardized or pre-configured value.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in Bytes.\n
	// Allowed Value:1-4095.\n
	// Optional
	MaxDataBurstVol *int32 `mapstructure:"maxDataBurstVol" json:"maxDataBurstVol,omitempty"`
}

// Dynamic5Qi
//
//	Purpose:
//	  Defines the QoS Characteristics for a Non-standardised or not pre-configured 5QI for downlink and uplink.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.4.3 for details
//
//	Usage:
//	  Configured inside VsmfQosProfile and DefaultQosInformation
type Dynamic5Qi struct {
	// Defines the 5QI resource type.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.3.6 Table 5.5.3.6-1.\n
	// -"NON_GBR":Non-GBR QoS Flow.\n
	// -"NON_CRITICAL_GBR":Non-delay critical GBR QoS flow.\n
	// -"CRITICAL_GBR":Delay critical GBR QoS flow.\n
	// Mandatory
	ResourceType QosResourceType `mapstructure:"resourceType" json:"resourceType"`
	// Defines the 5QI Priority Level.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// Allowed Value: 1-127.\n
	// Mandatory
	PriorityLevel int32 `mapstructure:"priorityLevel" json:"priorityLevel"`
	// Defines the packet delay budget.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in milliseconds.\n
	// Minimum = 1.\n
	// Mandatory
	PacketDelayBudget int32 `mapstructure:"packetDelayBudget" json:"packetDelayBudget"`
	// Defines the packet error rate.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// Pattern: '^([0-9]E-[0-9])$'.\n
	// Example: Packer Error Rate 4x10-6 shall be encoded as "4E-6".\n
	// Mandatory
	PacketErrRate string `mapstructure:"packetErrRate" json:"packetErrRate"`
	// Defines the averaging window. This IE shall be present only for a GBR QoS flow or a Delay Critical GBR QoS flow.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in milliseconds.\n
	// Minimum = 1. Maximum = 4095.\n
	// Optional
	AverWindow *int32 `mapstructure:"averWindow" json:"averWindow,omitempty"`
	// Defines the maximum data burst volume. This IE shall be present for a Delay Critical GBR QoS flow.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in Bytes.\n
	// Minimum = 1. Maximum = 4095.\n
	// Optional
	MaxDataBurstVol *int32 `mapstructure:"maxDataBurstVol" json:"maxDataBurstVol,omitempty"`
}

// GbrQosFlowInformation
//
//	Purpose:
//	  Defines the GBR QoS flow information.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.502/29502-g30.zip section 6.1.6.2.23 for details.
//
//	Usage:
//	  Configured inside VsmfQosProfile.
type GbrQosFlowInformation struct {
	// This IE shall contain the Maximum Bit Rate in Downlink.\n
	// Mandatory
	MaxFbrDl string `json:"maxFbrDl"`
	// This IE shall contain the Maximum Bit Rate in Uplink.\n
	// Mandatory
	MaxFbrUl string `json:"maxFbrUl"`
	// This IE shall contain the Guaranteed Bit Rate in Downlink.\n
	// Mandatory
	GuaFbrDl string `json:"guaFbrDl"`
	// This IE shall contain the Guaranteed Bit Rate in Uplink.\n
	// Mandatory
	GuaFbrUl string `json:"guaFbrUl"`
	// This IE indicate whether notifications are requested from the RAN when the GFBR can no longer (or again) be fulfilled for a QoS flow during the lifetime of the QoS flow.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.3.5 Table 5.5.3.5-1.\n
	// -"REQUESTED":Notifications are requested from the RAN.\n
	// -"NOT_REQUESTED":Notifications are not requested from the RAN.\n
	// Default value "".\n
	// Optional
	NotifControl NotificationControl `json:"notifControl,omitempty"`
	// This IE indicate the maximum rate for lost packets that can be tolerated in the downlink direction.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in tenth of percent.\n
	// Minimum = 0. Maximum = 1000.\n
	// Optional
	MaxPacketLossRateDl *int32 `json:"maxPacketLossRateDl,omitempty"`
	// This IE indicate the maximum rate for lost packets that can be tolerated in the Uplink direction.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in tenth of percent.\n
	// Minimum = 0. Maximum = 1000.\n
	// Optional
	MaxPacketLossRateUl *int32 `json:"maxPacketLossRateUl,omitempty"`
}

type NotificationControl = string

const (
	NotificationControl_REQUESTED     NotificationControl = "REQUESTED"
	NotificationControl_NOT_REQUESTED NotificationControl = "NOT_REQUESTED"
)

type QosResourceType = string

const (
	QosResourceType_NON_GBR          QosResourceType = "NON_GBR"
	QosResourceType_NON_CRITICAL_GBR QosResourceType = "NON_CRITICAL_GBR"
	QosResourceType_CRITICAL_GBR     QosResourceType = "CRITICAL_GBR"
)

type QuotaManagementIndicator = string

const (
	QuotaManagementIndicator_ONLINE_CHARGING  QuotaManagementIndicator = "ONLINE_CHARGING"
	QuotaManagementIndicator_OFFLINE_CHARGING QuotaManagementIndicator = "OFFLINE_CHARGING"
)

type RoamerInOut = string

const (
	RoamerInOut_IN_BOUND  RoamerInOut = "IN_BOUND"
	RoamerInOut_OUT_BOUND RoamerInOut = "OUT_BOUND"
)

// Ambr (Aggregate Maximum Bit Rate)
//
//	Purpose:
//	  Defines Ambr for a session.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.4.2 for details
//
//	Usage:
//	  Configured under SessRule
type Ambr struct {
	// AMBR for uplink. String in the form "<value> <unit>". Ex. "125 Mbps"
	Uplink string `mapstructure:"uplink" json:"uplink"`
	// AMBR for downlink. String in the form "<value> <unit>". Ex. "2 Gbps"
	Downlink string `mapstructure:"downlink" json:"downlink"`
}

// Default Qos
//
//	Purpose:
//	  Defines qos to be applied for the default flow of the session.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SessRule.
type DefaultQos struct {
	// 5G QoS Identifier.
	// Allowed Values: Refer https://www.3gpp.org/ftp/Specs/archive/24_series/24.501/24501-g40.zip section 9.11.4.12
	// Mandatory
	Var5qi int32 `mapstructure:"var5qi" json:"var5qi"`

	// Indicates the allocation and retention priority.
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.4.1 for details
	// Mandatory
	Arp Arp `mapstructure:"arp" json:"arp"`

	// Indicates the 5QI Priority Level
	// Allowed Value: value within a range of 1 to 127.
	// Optional
	PriorityLevel *int32 `mapstructure:"priorityLevel" json:"priorityLevel,omitempty"`

	// Represents the duration over which the guaranteed and maximum bitrate shall be calculated\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2.\n
	// expressed in milliseconds.\n
	// Minimum = 1. Maximum = 4095.\n
	// Default: nil.\n
	// Optional
	AverWindow *int32 `mapstructure:"averWindow" json:"averWindow,omitempty"`
}

// Qos Characteristics
//
//	Purpose:
//	  QoS characteristics for non-standard 5QIs and non-preconfigured 5QIs.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.16 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and used when non-standard and non-preconfigured 5QIs are used for flows in a session.
type QosCharacteristics struct {
	// Identifier for the authorized QoS parameters for the service data flow. Applies to PCC rule and PDU session level.\n
	// Mandatory
	Var5qi int32 `mapstructure:"var5qi" json:"var5qi"`
	// Indicates whether the resource type is GBR, delay critical GBR, or non-GBR.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.3.6 Table 5.5.3.6-1
	// -"NON_GBR":Non-GBR QoS Flow.\n
	// -"NON_CRITICAL_GBR":Non-delay critical GBR QoS flow.\n
	// -"CRITICAL_GBR":Delay critical GBR QoS flow.\n
	// Mandatory
	ResourceType string `mapstructure:"resourceType" json:"resourceType"`
	// Unsigned integer indicating the 5QI Priority Level.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// Allowed values: 1-127.\n
	// Mandatory
	PriorityLevel int32 `mapstructure:"priorityLevel" json:"priorityLevel"`
	// Unsigned integer indicates the packet delay budget.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in milliseconds.\n
	// Minimum = 1.\n
	// Mandatory
	PacketDelayBudget int32 `mapstructure:"packetDelayBudget" json:"packetDelayBudget"`
	// String indicating the packet error rate.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// Pattern: '^([0-9]E-[0-9])$'\n
	// Examples: Packer Error Rate 4x10-6 shall be encoded as "4E-6".\n
	// Mandatory
	PacketErrorRate string `mapstructure:"packetErrorRate" json:"packetErrorRate"`
	// Indicates the averaging window. This IE shall be present only for a GBR QoS flow or a Delay Critical GBR QoS flow.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// expressed in milliseconds.\n
	// Minimum = 1. Maximum = 4095.\n
	// Optional
	AveragingWindow *int32 `mapstructure:"averagingWindow" json:"averagingWindow,omitempty"`
	// Unsigned Integer.Indicates the maximum data burst volume. This IE shall be present for a Delay Critical GBR QoS flow.\n
	// expressed in Bytes.\n
	// Minimum = 1. Maximum = 4095.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip section 5.5.2 Table 5.5.2-1.\n
	// Optional
	MaxDataBurstVol *int32 `mapstructure:"maxDataBurstVol" json:"maxDataBurstVol,omitempty"`
}

// Chf 32291 6.1.6.3.11
type FailureHandling string

const (
	FailureHandling_TERMINATE           FailureHandling = "TERMINATE"
	FailureHandling_CONTINUE            FailureHandling = "CONTINUE"
	FailureHandling_RETRY_AND_TERMINATE FailureHandling = "RETRY_AND_TERMINATE"
)

// Quotamanage 32255 5.2.1.5
type QuotaManageType string

const (
	QuotaManage_SHARE_QUOTA   QuotaManageType = "SHARE_QUOTA"
	QuotaManage_PER_UPF_QUOTA QuotaManageType = "PER_UPF_QUOTA"
)

// Quotamanage 32255 5.2.1.5
type TrafficCountType string

const (
	TrafficCount_MULTIPLE_UPF TrafficCountType = "MULTIPLE_UPF"
	TrafficCount_SINGLE_UPF   TrafficCountType = "SINGLE_UPF"
)

type CourtesyExhaustHandle string

const (
	CourtesyExhaustHandle_Terminate CourtesyExhaustHandle = "TERMINATE"
	CourtesyExhaustHandle_Convert   CourtesyExhaustHandle = "CONVERT"
)

type ChargingService string

const (
	ChargingService_Converged    ChargingService = "CONVERGED_CHARGING"
	ChargingService_Offline_Only ChargingService = "OFFLINE_ONLY_CHARGING"
)

type PartialRecordMethod string

const (
	PartialRecordMethod_DEFAULT    PartialRecordMethod = "DEFAULT"
	PartialRecordMethod_INDIVIDUAL PartialRecordMethod = "INDIVIDUAL"
)

type ShareQuotaFixSize struct {
	// Total volume fix size qutoa, unit is byte
	TotalVolume *int64 `mapstructure:"totalVolume" json:"totalVolume,omitempty"`
	// Uplink volume fix size qutoa, unit is byte
	UplinkVolume *int64 `mapstructure:"uplinkVolume" json:"uplinkVolume,omitempty"`
	// Downlink volume fix size qutoa, unit is byte
	DownlinkVolume *int64 `mapstructure:"downlinkVolume" json:"downlinkVolume,omitempty"`
}

// Rating Group Triggers configuration
//
//	Purpose:
//	  This field specifies the reason for usage reporting for one or more types of unit for rating group.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value in RgTriggerMap
type RgTriggers struct {
	// Default rate group triggers.\n
	// This field specifies the reason for usage reporting for one or more types of unit associated to the rating group.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.1.7-1.\n
	// Optional.
	Triggers []Trigger `mapstructure:"triggers" json:"triggers,omitempty"`
	// Default quota
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.1.9-1.\n
	// Optional.
	RequestedUnit *RequestedUnit `mapstructure:"requestedUnit" json:"requestedUnit,omitempty"`
	// Default threshold
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.1.9-1.\n
	// Optional.
	DefThrhld *RequestedUnit `mapstructure:"defThrhld" json:"defThrhld,omitempty"`
	// Courtesy quota, using when can't get rsp from chf
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.1.9-1.\n
	// Optional.
	CourtesyQuota *RequestedUnit `mapstructure:"courtesyQuota" json:"courtesyQuota,omitempty"`
	// Courtesy threshold, using when can't get rsp from chf
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.1.9-1.\n
	// Optional.
	CourtesyThrhld *RequestedUnit `mapstructure:"courtesyThrhld" json:"courtesyThrhld,omitempty"`
	// Local quota hold time in seconds.\n
	// The NF Consumer shall deem a quota to have expired when no traffic associated with the quota is observed for the value indicated by this attribute.
	// Default value nil.\n
	// Optional.
	QuotaHoldingTime *int32 `mapstructure:"quotaHoldingTime" json:"quotaHoldingTime,omitempty"`
	// Local Validity Time in seconds. \n
	// Value of zero indicates that this mechanism shall not be used. \n
	// Default value nil.\n
	// Optional
	ValidityTime *int64 `mapstructure:"validityTime" json:"validityTime,omitempty"`
	// Fix size quota in byte
	// Only used when quota manage type is share quota
	// Default value nil.\n
	// Optional.
	FixSizeQuota *ShareQuotaFixSize `mapstructure:"fixSizeQuota" json:"fixSizeQuota,omitempty"`
}

// Roaming Charging Profile configuration
//
//	Purpose:
//	  This is default roaming charging profile associated to the PDU session for roaming QBC.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.15-1\n
//
//	Usage:
//	  Used as optional parameter "DefRoamingChargingProfile" in ChargingCharacteristics
type RoamingChargingProfile struct {
	// Trigger for roaming QBC
	Triggers []Trigger `json:"triggers,omitempty"`
	// Method uses for partial record closure
	PartialRecordMethod *PartialRecordMethod `json:"partialRecordMethod,omitempty"`
}

// Charging Characteristics configuration
//
//	Purpose:
//	  This is default charging configuration to be used by SMF when interacting with CHF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of CcMap
type ChargingCharacteristics struct {
	// CHF primary address and secondary address in format of ip:port.\n
	// Port 80 will be used as default port when port is absent.\n
	// Mandatory.
	ChargingInfo ChargingInformation `mapstructure:"chargingInfo" json:"chargingInfo"`
	// Indicator for session level online charging.
	//  -"true" means enable online charging.
	//  -"false" means disable online charging.
	// Default value "false".\n
	// Optional.
	Online bool `mapstructure:"online" json:"online,omitempty"`
	// Indicator for session level offline charging
	//  -"true" means enable offline charging.
	//  -"false" means disable offline charging.
	// Default value "false".\n
	// Optional.
	Offline bool `mapstructure:"offline" json:"offline,omitempty"`
	// Type of charging service for session.
	//  -"CONVERGED_CHARGING" represents Nchf_ConvergedCharging service (please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip section 5.2).
	//  -"OFFLINE_ONLY_CHARGING" represents Nchf_OfflineOnlyCharging service (please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip section 5.3).
	// Default value "".\n
	// Optional.
	ChargingService ChargingService `mapstructure:"chargingService" json:"chargingService,omitempty"`
	// The failure handling to be performed by the NF consumer when charging service invocation is temporarily prevented.\n
	// Value can be "TERMINATE", "CONTINUE" or "RETRY_AND_TERMINATE".
	// Default value "".\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.3.11-1.\n
	// Mandatory.
	FailureHandling FailureHandling `mapstructure:"failureHandling" json:"failureHandling"`
	// Retry count for sending charging data request while send request fail.\n
	// Value can be uint32 integer.\n
	// Default value 0 means not to retry sending the request.\n
	// Mandatory.
	Retries uint32 `mapstructure:"retries" json:"retries"`
	// Type of Traffic count.\n
	// Value can be "MULTIPLE_UPF" or "SINGLE_UPF".\n
	// Default value "".\n
	// Mandatory.
	TrafficCountType TrafficCountType `mapstructure:"trafficCountType" json:"trafficCountType"`
	// Type of Quota manage.\n
	// Value can be "SHARE_QUOTA" or "PER_UPF_QUOTA".\n
	// Default value "".\n
	// Mandatory.
	QuotaManageType QuotaManageType `mapstructure:"quotaManageType" json:"quotaManageType"`
	// Identifier of default rating group.\n
	// Value can be uint32 integer.\n
	// Default value 0.\n
	// Mandatory.
	RatingGroup uint32 `mapstructure:"ratingGroup" json:"ratingGroup"`
	// Session level default trigger.\n
	// This field specifies the reason for usage reporting for one or more types of unit associated to the rating group.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.1.7-1.\n
	// Default value nil.\n
	// Optional.
	Triggers []Trigger `mapstructure:"triggers" json:"triggers,omitempty"`
	// Default traffic time intervals.\n
	// Value can be a group of start and end time.\n
	// Default value nil.\n
	// Optional.
	TrafficTimes []TrafficTime `mapstructure:"trafficTimes" json:"trafficTimes,omitempty"`
	// Default triggers for rating group.\n
	// Rating group as key and rating group triggers as value\n
	// Default value nil.\n
	// Optional.
	RgTriggerMap map[uint32]RgTriggers `mapstructure:"rgTriggerMap" json:"rgTriggerMap,omitempty"`
	// The handling to be performed by the NF consumer when courtesy quota is exhausted.\n
	// Value can be "TERMINATE" or "CONVERT".\n
	// Default value ""\n
	// Optional.
	CourtesyExhaustHandle CourtesyExhaustHandle `mapstructure:"courtesyExhaustHandle" json:"courtesyExhaustHandle,omitempty"`
	// Default roaming charging profile.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/32_series/32.291/32291-g30.zip Table 6.1.6.2.15-1\n
	// Default value nil.\n
	// Optional.
	DefRoamingChargingProfile *RoamingChargingProfile `mapstructure:"defRoamingChargingProfile" json:"defRoamingChargingProfile,omitempty"`
	// Configurable options for charging enhancement.\n
	// Default value nil.\n
	// Optional.
	ChargingOptions *ChargingOptions `mapstructure:"chargingOptions" json:"chargingOptions,omitempty"`
}

// PCF selection and failure handling configuration
//
//	 Purpose:
//	   SMF can support applying the following failure handling operations on the interactions with PCF with local configurable policy
//
//	 Data model:
//	   Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator policy at slice level for reliable policy operation on HTTP timeout and on each HTTP error code.
type PcfClientPolicy struct {
	// Local configure fail handle while pcf data request fails because of timeout
	// - TERMINATE
	// - CONTINUE
	// - RETRY_AND_TERMINATE
	// Mandatory
	TimeoutHandling FailureHandling `mapstructure:"timeoutHandling" json:"timeoutHandling"`
	// Local configure fail handle while pcf data request fails with error code
	// - "400": TERMINATE
	// - "500": TERMINATE
	// Optional
	ErrorCodeHandling map[int]FailureHandling `mapstructure:"errorCodeHandling" json:"errorCodeHandling,omitempty"`
	// Indicates the pcf data request retransmission times while sending request timeout.
	// When smfsm sends msg to a PCF and gets timeout, smfsm will retry several times (number defined in retries) to make sure the timeout is not caused by e.g., network congestion in a short time.
	// If still can't reach PCF after retries, smfsm knows it is indeed a timeout, then smfsm will check timeoutHandling to determine if will send same msg to secondary PCF or directly terminate the service.
	// It is best to set to 0 (no retry) if TimeoutHandling is "TERMINATE"
	// Mandatory
	Retries uint32 `mapstructure:"retries" json:"retries"`
}

// Rule Base
//
//	Purpose:
//	  Reference to set of pcc rules to be applied for a session. SmPolicyDecision from PCF can refer to any predefined rule base under static pcc to apply set of pcc rule to a session
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Configured under static pcc as a map and used in Pcc Profile to refer to rule base for the session.
type RuleBase struct {
	// Reference to one or more Pcc rules configured under static pcc.
	// Mandatory
	RefPccRules []string `mapstructure:"refPccRules"   json:"refPccRules"`
	// Identifier for this rulebase.
	// Mandatory
	RuleBaseId string `mapstructure:"ruleBaseId"   json:"ruleBaseId"`
}

// Pcc Profile
//
//	Purpose:
//	  Pcc Profile has the reference to SessRule, Rulebase and DefaultPccRule to be used for a particular Dnn and Slice
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used in Static Pcc to config SessRule, Rulebase and DefaultPccRule for the session.
type PccProfile struct {
	// Reference to SessRule
	// Allowed Value: One of the session rules in SessRules map under static pcc.
	// Mandatory
	RefSessRule string `mapstructure:"refSessRule" json:"refSessRule"`

	// Reference to Rulebase which is a set of pcc rules
	// Allowed Value: One of the rulebases in Rulebases map under static pcc.
	// Optional
	RefRuleBase string `mapstructure:"refRuleBase"  json:"refRuleBase,omitempty"`

	// Reference to PccRule to be applied to default flow.
	// Allowed Value: Pcc rule configured in the PccRules map under static pcc for default flow.
	// Mandatory
	RefDefPccRule string `mapstructure:"refDefPccRule"  json:"refDefPccRule"`
}

// Pcc 29512 5.6.3.5
type MeteringMethod string

const (
	MeteringMethod_DURATION        MeteringMethod = "DURATION"
	MeteringMethod_VOLUME          MeteringMethod = "VOLUME"
	MeteringMethod_DURATION_VOLUME MeteringMethod = "DURATION_VOLUME"
	MeteringMethod_EVENT           MeteringMethod = "EVENT"
)

type ReportingLevel string

const (
	ReportingLevel_SER_ID_LEVEL   ReportingLevel = "SER_ID_LEVEL"
	ReportingLevel_RAT_GR_LEVEL   ReportingLevel = "RAT_GR_LEVEL"
	ReportingLevel_SPON_CON_LEVEL ReportingLevel = "SPON_CON_LEVEL"
)

// Charging Data
//
//	Purpose:
//	  Describes charging related information associated to a flow.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.2.11 for details
//
//	Usage:
//	  Configured under StaticPcc as a map and referenced by PccRule.
type ChargingData struct {
	// Univocally identifies the charging control policy data within a PDU session.\n
	// Mandatory.
	ChgId string `mapstructure:"chgId"           json:"chgId"`
	// Defines what parameters shall be metered for offline charging.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.3.5 Table 5.6.3.5-1.\n
	// -"DURATION":Indicates that the duration of the service data flow traffic shall be metered.\n
	// -"VOLUME":Indicates that volume of the service data flow traffic shall be metered.\n
	// -"DURATION_VOLUME":Indicates that the duration and the volume of the service data flow traffic shall be metered.\n
	// -"EVENT":Indicates that events of the service data flow traffic shall be metered.\n
	// Optional
	MeteringMethod *MeteringMethod `mapstructure:"meteringMethod"  json:"meteringMethod,omitempty"`
	// Indicates the offline charging is applicable to the PDU session or PCC rule.\n
	// Optional
	Offline *bool `mapstructure:"offline"                          json:"offline,omitempty"`
	// Indicates the online charging is applicable to the PDU session or PCC rule.\n
	// Optional
	Online *bool `mapstructure:"online"          json:"online,omitempty"`
	// Indicates whether the service data flow is allowed to start while the SMF is waiting for the response to the credit request.\n
	// Optional
	SdfHandl *bool `mapstructure:"sdfHandl" json:"sdfHandl,omitempty"`
	// The charging key for the PCC rule used for rating purposes.\n
	// Optional
	RatingGroup *uint32 `mapstructure:"ratingGroup"     json:"ratingGroup,omitempty"`
	// Defines on what level the SMF reports the usage for the related PCC rule.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip section 5.6.3.4 Table 5.6.3.4-1.\n
	// -"SER_ID_LEVEL":Indicates that the usage shall be reported on service id and rating group combination level.\n
	// -"RAT_GR_LEVEL":Indicates that the usage shall be reported on rating group level.\n
	// -"SPON_CON_LEVEL":Indicates that the usage shall be reported on sponsor identity and rating group combination level.\n
	// Optional
	ReportingLevel *ReportingLevel `mapstructure:"reportingLevel"  json:"reportingLevel,omitempty"`
	// Indicates the identifier of the service or service component the service data flow in a PCC rule relates to.\n
	// Optional
	ServiceId *int32 `mapstructure:"serviceId"       json:"serviceId,omitempty"`
	// Indicates the sponsor identity.\n
	// Optional
	SponsorId *string `mapstructure:"sponsorId"                    json:"sponsorId,omitempty"`
	// Indicates the application service provider identity.\n
	// Optional
	AppSvcProvId *string `mapstructure:"appSvcProvId"    json:"appSvcProvId,omitempty"`
	// An identifier, provided from the AF, correlating the measurement for the Charging key/Service identifier values in this PCC rule with application level reports.\n
	// Optional
	AfChargingIdentifier *int32 `mapstructure:"afChargingIdentifier"  json:"afChargingIdentifier,omitempty"`
}

// LocationTcRule
//
//	Purpose:
//	  Location based traffic steering rules.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured as a slice under StaticPcc.
type LocationTcRule struct {
	// Use TaiList for now, in the future to support Locality defined in location_info.go
	// Optional
	TaiList []Tai `mapstructure:"taiList" json:"taiList,omitempty"`
	// List of Traffic steering rules reference to data in StaticPcc.PccRules
	// Optional
	RefRuleBase string `mapstructure:"refRuleBase"  json:"refRuleBase,omitempty"`
}

type DscpDirection string

const (
	DscpDirection_UPLINK        DscpDirection = "UPLINK"
	DscpDirection_DOWNLINK      DscpDirection = "DOWNLINK"
	DscpDirection_BIDIRECTIONAL DscpDirection = "BIDIRECTIONAL"
)

// MessageMarkingMapping
//
//	Purpose:
//	  use 5qi, ArpPriorityLevel(option) and PriorityLevel(option) to map dscp.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured as a slice under SlicePolicyProfile.
type MessageMarkingMapping struct {
	// Transport level marking shall be controlled by the CP function by providing the DSCP in the ToS or Traffic Class within the Transport Level Marking IE in the FAR that is associated to the PDR matching the traffic to be marked.\n
	// Allowed Value:0-63.\n
	// Default Value 0.\n
	// Mandatory
	Dscp uint8 `mapstructure:"dscp" json:"dscp"`
	// This IE shall contain the 5G QoS Identifier (5QI) of the QoS flow.\n
	// Mandatory
	Var5qi uint32 `mapstructure:"var5qi" json:"var5qi"`
	// The priority in ARP.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/23_series/23.501/23501-g40.zip section 5.7.2.2.\n
	// Allowed Value:1-15\n
	// Default value nil.\n
	// Optional
	ArpPriorityLevel *int32 `mapstructure:"arpPriorityLevel" json:"arpPriorityLevel,omitempty"`
	// Indicates a priority in scheduling resources among QoS Flows. The PriorityLevel.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/23_series/23.501/23501-g40.zip section 5.7.3.3.\n
	// Optional
	PriorityLevel *int32 `mapstructure:"priorityLevel" json:"priorityLevel,omitempty"`
	// Indicates the direction of dscp.\n
	// Allowed Value: UPLINK, DOWNLINK, BIDIRECTIONAL.\n
	// Default Value: BIDIRECTIONAL.\n
	// Optional
	Direction DscpDirection `mapstructure:"direction" json:"direction,omitempty"`
}

// Static Pcc
//
//	Purpose:
//	  Static Pcc configuration is used in the following scenarios:
//	    1) if StaticPccOnly attribute is set under DnnProfile.
//	    2) if SMF fail to discover PCF from either NRF or config.
//	    3) if Npcf_SMPolicyControl_Create Service Operation fails.
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used to define elements of StaticPccMap in SmfConfig.
type StaticPcc struct {
	// Default PccProfile which has reference to SessRule, RuleBase and DefaultPccRule used for Dnn and Slice
	// Mandatory
	DefPccProfile PccProfile `mapstructure:"defPccProfile" json:"defPccProfile"`

	// Map of Sessionrules with key as sessRuleId attribute of the corresponding SessionRule and value as SessionRule described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.7.
	// Mandatory
	SessRules map[string]SessionRule `mapstructure:"sessRules" json:"sessRules"`

	// Map of Rulebase with key as RuleBase name and value as Rulebase which is defined as a set of PccRules.
	// Conditional. Has to be present if referenced by DefPccProfile
	RuleBases map[string]RuleBase `mapstructure:"ruleBases"    json:"ruleBases,omitempty"`

	// Map of Pcc rules with key as pccRuleId and value as pcc rule described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.6.
	// Mandatory : Should have atleast default pcc rule.
	PccRules map[string]PccRule `mapstructure:"pccRules" json:"pccRules"`

	// Map of QoS data policy decisions with key as qosId and value as QosData described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.8.
	// Mandatory : Should have atleast QosData for default flow
	QosDecs map[string]QosData `mapstructure:"qosDecs" json:"qosDecs"`

	// Map of Charging data policy decisions with key as chgId and value as ChargingData described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.11
	// Optional
	ChgDecs map[string]ChargingData `mapstructure:"chgDecs" json:"chgDecs,omitempty"`

	// Map of Traffic Control data policy decisions with key as tcId attribute of the corresponding TrafficControlData and value as TrafficControlData described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.10
	// Optional
	TraffContDecs map[string]TrafficControlData `mapstructure:"traffContDecs" json:"traffContDecs,omitempty"`

	// Map of Usage Monitoring data policy decisions with key as umId attribute of the corresponding UsageMonitoringData and value as UsageMonitoringData described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.12
	// Optional
	UmDecs map[string]UsageMonitoringData `mapstructure:"umDecs" json:"umDecs,omitempty"`

	// Map of QoS characteristics for non standard 5QIs and non-preconfigured 5QIs with key as 5QI and value as QosCharacteristics described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.16
	// Optional
	QosChars map[string]QosCharacteristics `mapstructure:"qosChars" json:"qosChars,omitempty"`
	// Map of condition data with key as condId attribute of the corresponding ConditionData and value as ConditionData described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/29512-g40.zip subclause 5.6.2.9
	// Optional
	Conds map[string]ConditionData `mapstructure:"conds" json:"conds,omitempty"`

	// Map of Presence Reporting Area information provisioned by the PCF with key as "praId"and value as PresenceInfoRm described in https://www.3gpp.org/ftp/Specs/archive/29_series/29.571/29571-g30.zip subclause 5.4.4.30. The "presenceState" attribute within the PresenceInfo data type shall not be supplied.
	// Optional
	PraInfos map[string]PresenceInfoRm `mapstructure:"praInfos" json:"praInfos,omitempty"`

	// Location based traffic steering rules
	// Optional
	LocationTcRules []LocationTcRule `mapstructure:"locationTcRules" json:"locationTcRules,omitempty"`

	// Rule base for 3GPP PS Datat Off Exempt Services configured in the rulebase map under static pcc.
	// All services belonging to the list(s) of 3GPP PS Data Off Exempt Services can be represented by PCC rule(s) which allows the traffic to pass while in all other PCC rules (not belonging to the list(s) of 3GPP PS Data Off Exempt Services) the gate status can be "closed" for downlink and optionally uplink direction.
	// Optional
	PsDataOffExemptServicesRefRuleBase string `mapstructure:"psDataOffExemptServicesRefRuleBase" json:"psDataOffExemptServicesRefRuleBase,omitempty"`

	// Time in seconds which defines the lifetime of a UE derived QoS rule belonging to the PDU Session for reflective QoS.
	// Optional
	ReflectiveQoSTimer int32 `mapstructure:"reflectiveQoSTimer" json:"reflectiveQoSTimer,omitempty"`

	// Indicates whether the PDU Session is a redundant PDU session.
	// - true: End to end redundant PDU session.
	// - false: Not end to end redundant PDU session.
	// Default value false
	// Optional
	RedSessIndication bool `mapstructure:"redSessIndication" json:"redSessIndication,omitempty"`

	// XHeader config map with XHeader name as key. This is referenced from PccRule using XHeaderRef field.
	// Optional
	// Deprecated effective rel 7.4
	XHeaderCfg map[string]XHeaderInfo `mapstructure:"xHeaderCfg" json:"xHeaderCfg,omitempty"`

	// TlsExtention config map with TlsExtension name as key. This is referenced from PccRule using TlsExtensionRef field.
	// Optional
	// Deprecated effective rel 7.4
	TlsExtensionCfg map[string]TlsExtension `mapstructure:"tlsExtensionCfg" json:"tlsExtensionCfg,omitempty"`

	// Map HeaderEnrichmentRef to HttpHeaderInfo
	// Optional
	// To be deprecated
	HeaderEnrichmentCfg HeaderEnrichmentConfig `mapstructure:"headerEnrichmentCfg" json:"headerEnrichmentCfg,omitempty"`

	// Map HeaderEnrichmentRef to HeaderField
	// Optional
	HeaderFieldCfg HeaderFieldConfig `mapstructure:"headerFieldCfg" json:"headerFieldCfg,omitempty"`

	// OperatorPCORules configure different rule for pco/epco
	// Optional
	OperatorPCORules []OperatorPCORule `mapstructure:"operatorPCORules" json:"operatorPCORules,omitempty"`
}

// OperatorPCORule
//
//	Purpose:
//	  OperatorPCORule configure pco/epco rule
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used to configure under StaticPcc.
type OperatorPCORule struct {
	// Config operator pco rule name, will be active by PCF returned pcc rule name Or CHF returned filterId.
	// Mandatory
	Name string `mapstructure:"name" json:"name"`
	// operator specific epco use container id, a hexadecimal character string, e.g "FF00"
	// Mandatory
	Container string `mapstructure:"container" json:"container"`
	// operator specific epco use container value, a decimal string, e.g "5"
	// Optional
	Value string `mapstructure:"value" json:"value,omitempty"`
}

// Verizon Charging Options
//
//	Purpose:
//	  Configurable options for charging enhancement.
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used to configure under ChargingCharacteristics.
type ChargingOptions struct {
	// Period in seconds to run an additional Offline-Session-Timer for all offline-only charging rules.
	// The goal for the Offline-Session-Timer to expire periodically is to force report the usage of
	// offline-only charging rules to CHF.
	// Optional
	OfflineSessionTimePeriod *int `mapstructure:"offlineSessionTimePeriod" json:"offlineSessionTimePeriod,omitempty"`

	// Indicates whether empty containers for offline-only charging rules should be suppressed.
	// When a trigger occurs that would normally cause usage to be reported, if there is
	// no usage on that Rating Group, there is no need to send an empty container that Verizon
	// will filter out downstream.
	// Default value false
	// Optional
	SuppressEmptyOfflineOnlyContainers bool `mapstructure:"suppressEmptyOfflineOnlyContainers" json:"suppressEmptyOfflineOnlyContainers,omitempty"`

	// Indicates whether empty containers for online charging rules should be suppressed.
	// When a trigger occurs that would normally cause usage to be reported, if there is
	// no usage on that Rating Group, there is no need to send an empty container that Verizon
	// will filter out downstream.
	// Default value false
	// Optional
	SuppressEmptyOnlineOnlyContainers bool `mapstructure:"suppressEmptyOnlineOnlyContainers" json:"suppressEmptyOnlineOnlyContainers,omitempty"`

	// Configurable options for wait no wait control.\n
	// Default value nil.\n
	// Optional.
	WaitNoWaitControl *ChargingWaitNoWaitControlOptions `mapstructure:"waitNoWaitControl" json:"waitNoWaitControl,omitempty"`
}

// Charging Wait/No-wait Control Options
//
//	Purpose:
//	  Configurable options for wait/no-wait control in charging scenarios.
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used to configure under ChargingOptions.
type ChargingWaitNoWaitControlOptions struct {
	// Indicates whether no wait CHF for online charging rules.
	// Default value false (false: wait, true: no wait)
	// Optional
	NoWaitForOnlineCharging bool `mapstructure:"noWaitForOnlineCharging" json:"noWaitForOnlineCharging,omitempty"`

	// Indicates whether no wait CHF for offline charging rules.
	// Default value false (false: wait, true: no wait)
	// Optional
	NoWaitForOfflineCharging bool `mapstructure:"noWaitForOfflineCharging" json:"noWaitForOfflineCharging,omitempty"`

	// When don't wait CHF, and the traffic is blocked due to the sdfhandl parameter set to 'block',
	// whether to buffer (otherwise drop) packets for online charging rules.
	// Default value nil (nil: none, false: drop packets, true: buffer packets)
	// Optional
	BufferPacketsForSdfBlocking *bool `mapstructure:"bufferPacketsForSdfBlocking" json:"bufferPacketsForSdfBlocking,omitempty"`
}
