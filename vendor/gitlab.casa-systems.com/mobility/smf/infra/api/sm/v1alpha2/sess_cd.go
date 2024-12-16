package v1alpha2

import "time"

type PduSessionInformation struct {
	// information of network slice serving the PDU session
	NetworkSlicingInfo *NetworkSlicingInfo `mapstructure:"networkSlicingInfo" json:"networkSlicingInfo,omitempty"`
	PduSessionID       int32               `mapstructure:"pduSessionID" json:"pduSessionID"`
	// type of the PDU session
	PduType *PduSessionType `mapstructure:"pduType" json:"pduType,omitempty"`
	// information of SSC Mode type.
	SscMode *SscMode `mapstructure:"sscMode" json:"sscMode,omitempty"`
	// PLMN identifier of the home network
	HPlmnId *PlmnId `mapstructure:"hPlmnId" json:"hPlmnId,omitempty"`
	// This field holds serving Network Function identifier.
	ServingNetworkFunctionID *SdServingNetworkFunctionId `mapstructure:"servingNetworkFunctionID" json:"servingNetworkFunctionID,omitempty"`
	// the RAT Type of the PDU session
	RatType *RatType `mapstructure:"ratType" json:"ratType,omitempty"`
	// a Data Network Name
	DnnId            string            `mapstructure:"dnnId" json:"dnnId"`
	DnnSelectionMode *DnnSelectionMode `json:"dnnSelectionMode,omitempty"`
	// the Charging Characteristics for this PDU session.
	ChargingCharacteristics string `mapstructure:"chargingCharacteristics" json:"chargingCharacteristics,omitempty"`
	// information about how the "Charging Characteristics" was selected.
	ChargingCharacteristicsSelectionMode *ChargingCharacteristicsSelectionMode `mapstructure:"chargingCharacteristicsSelectionMode" json:"chargingCharacteristicsSelectionMode,omitempty"`
	// the time in UTC format which represents the start of a PDU session at the SMF
	StartTime *time.Time `mapstructure:"startTime" json:"startTime,omitempty"`
	// the time in UTC format which represents the stop of a PDU session at the SMF
	StopTime *time.Time `mapstructure:"stopTime" json:"stopTime,omitempty"`
	// This field holds the 3GPP Data off Status when UE’s 3GPP Data Off status is Activated or Deactivated.
	Var3gppPSDataOffStatus *Model3GpppsDataOffStatus `mapstructure:"3gppPSDataOffStatus" json:"3gppPSDataOffStatus,omitempty"`
	// This field indicates to the CHF that the PDU session has been terminated.
	SessionStopIndicator bool `mapstructure:"sessionStopIndicator" json:"sessionStopIndicator,omitempty"`
	// Group of user ip address/prefix
	PduAddress *PduAddress `mapstructure:"pduAddress" json:"pduAddress,omitempty"`
	// provides a more detailed cause value from SMF.
	Diagnostics              *int32                `mapstructure:"diagnostics" json:"diagnostics,omitempty"`
	AuthorizedQoSInformation *AuthorizedDefaultQos `json:"authorizedQoSInformation,omitempty"`
	SubscribedQoSInformation *SubscribedDefaultQos `json:"subscribedQoSInformation,omitempty"`
	AuthorizedSessionAMBR    *Ambr                 `json:"authorizedSessionAMBR,omitempty"`
	SubscribedSessionAMBR    *Ambr                 `json:"subscribedSessionAMBR,omitempty"`
	// Serving Core Network Operator PLMN ID selected by the UE in shared networks.
	ServingCNPlmnId *PlmnId                `mapstructure:"servingCNPlmnId" json:"servingCNPlmnId,omitempty"`
	QoSInformation  *DefaultQosInformation `mapstructure:"qoSInformation" json:"qoSInformation,omitempty"`
}

type SubscribedDefaultQos struct {
	Var5qi        int32  `json:"5qi"`
	Arp           Arp    `json:"arp"`
	PriorityLevel *int32 `json:"priorityLevel,omitempty"`
}

type DnnSelectionMode = string

type UserInformation struct {
	ServedGPSI          string      `mapstructure:"servedGPSI" json:"servedGPSI"`
	ServedPEI           string      `mapstructure:"servedPEI" json:"servedPEI,omitempty"`
	UnauthenticatedFlag bool        `mapstructure:"unauthenticatedFlag" json:"unauthenticatedFlag,omitempty"`
	RoamerInOut         RoamerInOut `mapstructure:"roamerInOut" json:"roamerInOut,omitempty"`
}
type PduSessionChargingInformation struct {
	// Charging identifier for correlation between different records of a single PDU session
	ChargingId *uint32 `mapstructure:"chargingId" json:"chargingId,omitempty"`
	// including information of user and user equipment,
	UserInformation *UserInformation `mapstructure:"userInformation" json:"userInformation,omitempty"`
	// provides information on the location
	UserLocationinfo *UserLocation `mapstructure:"userLocationinfo" json:"userLocationinfo,omitempty"`
	UserLocationTime *time.Time    `mapstructure:"userLocationTime" json:"userLocationTime,omitempty"`
	// When the data type is present in response message, it includes the PRA information provisioned by the CHF, in which case the "presenceState" attribute within the PresenceInfo data type shall not be supplied. When the data type is present in request message, it’s used to report user presence reporting area status.
	// The praId attribute within the PresenceInfo data type shall be the key of the map.
	PresenceReportingAreaInformation map[string]PresenceInfo `mapstructure:"presenceReportingAreaInformation" json:"presenceReportingAreaInformation,omitempty"`
	// the UE Timezone the UE is currently located
	UetimeZone string `mapstructure:"uetimeZone" json:"uetimeZone,omitempty"`
	// PDU session level information, including PDU session ID, PDU type, SSC Mode, QoS, network slicing etc.
	PduSessionInformation PduSessionInformation `mapstructure:"pduSessionInformation" json:"pduSessionInformation"`
	UnusedQuotaTimer      *int32                `mapstructure:"unusedQuotaTimer" json:"unusedQuotaTimer,omitempty"`

	RANSecondaryRATUsageReport *RanSecondaryRatUsageReport `json:"rANSecondaryRATUsageReport,omitempty"`
	//version: 1.R15.0.0
	UnitCountInactivityTimer *int32 `json:"unitCountInactivityTimer,omitempty"`
}

type RanSecondaryRatUsageReport struct {
	RANSecondaryRATType  *RatType              `json:"rANSecondaryRATType,omitempty"`
	QosFlowsUsageReports []QosFlowsUsageReport `json:"qosFlowsUsageReports,omitempty"`
}
type QosFlowsUsageReport struct {
	QFI            *int32     `json:"qFI,omitempty"`
	StartTimestamp *time.Time `json:"startTimestamp,omitempty"`
	EndTimestamp   *time.Time `json:"endTimestamp,omitempty"`
	UplinkVolume   *int64     `json:"uplinkVolume,omitempty"`
	DownlinkVolume *int64     `json:"downlinkVolume,omitempty"`
}

type PduAddress struct {
	// the IPv4 address of the served SUPI allocated for the PDU session
	PduIPv4Address string `mapstructure:"pduIPv4Address" json:"pduIPv4Address,omitempty"`
	// the IPv6 address with prefix of the served SUPI allocated for the PDU session
	PduIPv6Address string `mapstructure:"pduIPv6Address" json:"pduIPv6Address,omitempty"`
	// PDU Address prefix length of an IPv6 typed Served PDU Address. The field needs not available for prefix length of 64 bits.
	PduAddressprefixlength *int32 `mapstructure:"pduAddressprefixlength" json:"pduAddressprefixlength,omitempty"`
	// This field indicates whether served IPv4 address is dynamically allocated. This field is missing if address is static.
	IPv4dynamicAddressFlag *bool `mapstructure:"IPv4dynamicAddressFlag" json:"iPv4dynamicAddressFlag,omitempty"`
	// This field indicates whether served IPv6 address prefix is dynamically allocated. This field is missing if address is static.
	IPv6dynamicAddressFlag *bool `mapstructure:"IPv6dynamicAddressFlag" json:"iPv6dynamicAddressFlag,omitempty"`
	IPv6dynamicPrefixFlag  *bool `json:"iPv6dynamicPrefixFlag,omitempty"`
	//version 1.R15.0.0
	PduIPv6AddresswithPrefix *string `json:"pduIPv6AddresswithPrefix,omitempty"`
}

type UrrIDMap struct {
	// URR ID
	UrrID uint32 `mapstructure:"urrID" json:"urrID,omitempty"`
	// RatingGroup
	RatingGroup uint32 `mapstructure:"ratingGroup" json:"ratingGroup,omitempty"`
	// Qfi
	Qfi uint32 `mapstructure:"qfi" json:"qfi,omitempty"`
	//reference counting, save the number how many flow is using it.
	RefCount      int32  `mapstructure:"refCount" json:"refCount,omitempty"`
	UpfInstanceId string `mapstructure:"upfInstanceId" json:"upfInstanceId,omitempty"`
	// Service id is omitempty, using point to judgement whether exist or not.
	ServiceID *uint32 `mapstructure:"serviceID" json:"serviceID,omitempty"`
}

type SdServingNetworkFunctionId struct {
	ServingNetworkFunctionInformation NfIdentification `json:"servingNetworkFunctionInformation"`
	AMFId                             string           `json:"aMFId,omitempty"`
	//version:1.R15.0.0
	// This field holds serving Network Function name
	ServingNetworkFunctionName *string `mapstructure:"servingNetworkFunctionName" json:"servingNetworkFunctionName,omitempty"`
	// This field holds serving Network Function instance identity
	ServingNetworkFunctionInstanceid *string `mapstructure:"servingNetworkFunctionInstanceid" json:"servingNetworkFunctionInstanceid,omitempty"`
	//gUAMI *Guami  `mapstructure:"gUAMI" json:"gUAMI,omitempty"`  //missing
}

type NfIdentification struct {
	NFName            string            `json:"nFName,omitempty"`
	NFIPv4Address     string            `json:"nFIPv4Address,omitempty"`
	NFIPv6Address     string            `json:"nFIPv6Address,omitempty"`
	NFPLMNID          *PlmnId           `json:"nFPLMNID,omitempty"`
	NodeFunctionality NodeFunctionality `json:"nodeFunctionality"`
	NFFqdn            string            `json:"nFFqdn,omitempty"`
}

type NodeFunctionality = string

// MultipleUnitUsage struct for MultipleUnitUsage
type MultipleUnitUsage struct {
	RatingGroup          uint32              `json:"ratingGroup"`
	RequestedUnit        *RequestedUnit      `json:"requestedUnit,omitempty"` //converged proprietary
	UsedUnitContainer    []UsedUnitContainer `json:"usedUnitContainer,omitempty"`
	UPFID                *string             `json:"uPFID,omitempty"`
	MultihomedPDUAddress *PduAddress         `json:"multihomedPDUAddress,omitempty"`
}

type RequestedUnit struct {
	// Report time, unit is second
	Time *int32 `mapstructure:"time" json:"time,omitempty"`
	// Total volume fix size qutoa, unit is byte
	TotalVolume *int64 `mapstructure:"totalVolume" json:"totalVolume,omitempty"`
	// Total volume fix size qutoa, unit is byte
	UplinkVolume *int64 `mapstructure:"uplinkVolume" json:"uplinkVolume,omitempty"`
	// Total volume fix size qutoa, unit is byte
	DownlinkVolume *int64 `mapstructure:"downlinkVolume" json:"downlinkVolume,omitempty"`
	// Event base charging unit, unit is event number
	ServiceSpecificUnits *int64 `mapstructure:"serviceSpecificUnits" json:"serviceSpecificUnits,omitempty"`
}

type UsedUnitContainer struct {
	ServiceId                *uint32                   `mapstructure:"serviceId" json:"serviceId,omitempty"`
	QuotaManagementIndicator *QuotaManagementIndicator `mapstructure:"quotaManagementIndicator" json:"quotaManagementIndicator,omitempty"`
	Triggers                 []SdTrigger               `mapstructure:"triggers" json:"triggers,omitempty"`
	TriggerTimestamp         *time.Time                `mapstructure:"triggerTimestamp" json:"triggerTimestamp,omitempty"`
	Time                     *int32                    `mapstructure:"time" json:"time,omitempty"`
	TotalVolume              *int64                    `mapstructure:"totalVolume" json:"totalVolume,omitempty"`
	UplinkVolume             *int64                    `mapstructure:"uplinkVolume" json:"uplinkVolume,omitempty"`
	DownlinkVolume           *int64                    `mapstructure:"downlinkVolume" json:"downlinkVolume,omitempty"`
	ServiceSpecificUnits     *int64                    `mapstructure:"serviceSpecificUnits" json:"serviceSpecificUnits,omitempty"`
	EventTimeStamps          *time.Time                `mapstructure:"eventTimeStamps" json:"eventTimeStamps,omitempty"`
	LocalSequenceNumber      int32                     `mapstructure:"localSequenceNumber" json:"localSequenceNumber"`
	PDUContainerInformation  *PduContainerInformation  `mapstructure:"pDUContainerInformation" json:"pDUContainerInformation,omitempty"`
}

type PduContainerInformation struct {
	TimeofFirstUsage                   *time.Time                   `mapstructure:"timeofFirstUsage" json:"timeofFirstUsage,omitempty"`
	TimeofLastUsage                    *time.Time                   `mapstructure:"timeofLastUsage" json:"timeofLastUsage,omitempty"`
	QoSInformation                     *SdQosData                   `mapstructure:"qoSInformation" json:"qoSInformation,omitempty"`
	AFCorrelationInformation           string                       `mapstructure:"aFCorrelationInformation" json:"aFCorrelationInformation,omitempty"`
	UserLocationInformation            *UserLocation                `mapstructure:"userLocationInformation" json:"userLocationInformation,omitempty"`
	UetimeZone                         string                       `mapstructure:"uetimeZone" json:"uetimeZone,omitempty"`
	RATType                            RatType                      `mapstructure:"rATType" json:"rATType,omitempty"`
	ServingNodeID                      []SdServingNetworkFunctionId `mapstructure:"servingNodeID" json:"servingNodeID,omitempty"` // R16
	PresenceReportingAreaInformation   map[string]PresenceInfo      `mapstructure:"presenceReportingAreaInformation" json:"presenceReportingAreaInformation,omitempty"`
	Var3gppPSDataOffStatus             Model3GpppsDataOffStatus     `mapstructure:"3gppPSDataOffStatus" json:"3gppPSDataOffStatus,omitempty"`
	SponsorIdentity                    string                       `mapstructure:"sponsorIdentity" json:"sponsorIdentity,omitempty"`
	ApplicationserviceProviderIdentity string                       `mapstructure:"applicationserviceProviderIdentity" json:"applicationserviceProviderIdentity,omitempty"`
	ChargingRuleBaseName               string                       `mapstructure:"chargingRuleBaseName" json:"chargingRuleBaseName,omitempty"`
}

// ChargingData struct for ChargingData
type SdChargingData struct {
	// Univocally identifies the charging control policy data within a PDU session.
	ChgId          string          `json:"chgId"`
	MeteringMethod *MeteringMethod `json:"meteringMethod,omitempty"`
	// Indicates the offline charging is applicable to the PCC rule.
	Offline *bool `json:"offline,omitempty"`
	// Indicates the online charging is applicable to the PCC rule.
	Online *bool `json:"online,omitempty"`
	// Indicates whether the service data flow is allowed to start while the SMF is waiting for the response to the credit request.
	SdfHandl       *bool          `json:"sdfHandl,omitempty"`
	RatingGroup    *uint32        `json:"ratingGroup,omitempty"`
	ReportingLevel ReportingLevel `json:"reportingLevel,omitempty"`
	ServiceId      *int32         `json:"serviceId,omitempty"`
	// Indicates the sponsor identity.
	SponsorId string `json:"sponsorId,omitempty"`
	// Indicates the application service provider identity.
	AppSvcProvId         string `json:"appSvcProvId,omitempty"`
	AfChargingIdentifier *int32 `json:"afChargingIdentifier,omitempty"`
}

type SdQosData struct {
	// Univocally identifies the QoS control policy data within a PDU session.
	QosId string `mapstructure:"qosId" json:"qosId"`
	// Identifier for the authorized QoS parameters for the service data flow. It shall be included when the QoS data decision is initially provisioned and "defQosFlowIndication" is not included or is included and set to false.
	Var5qi int32 `mapstructure:"5qi" json:"5qi"`
	// Indicates the max bandwidth in uplink.
	MaxbrUl string `mapstructure:"maxbrUl" json:"maxbrUl,omitempty"`
	// Indicates the max bandwidth in downlink.
	MaxbrDl string `mapstructure:"maxbrDl" json:"maxbrDl,omitempty"`
	// Indicates the guaranteed bandwidth in uplink.
	GbrUl string `mapstructure:"gbrUl" json:"gbrUl,omitempty"`
	// Indicates the guaranteed bandwidth in downlink.
	GbrDl string `mapstructure:"gbrDl" json:"gbrDl,omitempty"`
	// Indicates the allocation and retention priority. It shall be included when the QoS data decision is initially provisioned and "defQosFlowIndication" is not included or is included and set to false.
	Arp Arp `mapstructure:"arp" json:"arp"`
	// Indicates whether notifications are requested from 3GPP NG-RAN when the GFBR can no longer (or again) be guaranteed for a QoS Flow during the lifetime of the QoS Flow.
	Qnc bool `mapstructure:"qnc" json:"qnc,omitempty"`
	// Indicates a priority in scheduling resources among QoS Flows
	PriorityLevel *int32 `mapstructure:"priorityLevel" json:"priorityLevel,omitempty"`
	// Represents the duration over which the guaranteed and maximum bitrate shall be calculated
	AverWindow *int32 `mapstructure:"averWindow" json:"averWindow,omitempty"`
	// Denotes the largest amount of data that is required to be transferred within a period of 5G-AN PDB
	MaxDataBurstVol *int32 `mapstructure:"maxDataBurstVol" json:"maxDataBurstVol,omitempty"`
	// Indicates whether the QoS information is reflective for the corresponding service data flow.
	ReflectiveQos bool `mapstructure:"reflectiveQos" json:"reflectiveQos,omitempty"`
	// Indicates, by containing the same value, what PCC rules may share resource in downlink direction.
	SharingKeyDl string `mapstructure:"sharingKeyDl" json:"sharingKeyDl,omitempty"`
	// Indicates, by containing the same value, what PCC rules may share resource in uplink direction.
	SharingKeyUl string `mapstructure:"sharingKeyUl" json:"sharingKeyUl,omitempty"`
	// Indicates the downlink maximum rate for lost packets that can be tolerated for the service data flow.
	MaxPacketLossRateDl *int32 `mapstructure:"maxPacketLossRateDl" json:"maxPacketLossRateDl,omitempty"`
	// Indicates the uplink maximum rate for lost packets that can be tolerated for the service data flow.
	MaxPacketLossRateUl *int32 `mapstructure:"maxPacketLossRateUl" json:"maxPacketLossRateUl,omitempty"`
	// Indicates that the dynamic PCC rule shall always have its binding with the QoS Flow associated with the default QoS rule
	DefQosFlowIndication bool `mapstructure:"defQosFlowIndication" json:"defQosFlowIndication,omitempty"`
}

type ServingNetworkFunctionId struct {
	ServingNetworkFunctionInformation NfIdentification `json:"servingNetworkFunctionInformation"`
	AMFId                             string           `json:"aMFId,omitempty"`
}

type DefaultQosInformation struct {
	Var5qi        int32          `mapstructure:"var5qi" json:"var5qi"`
	NonDynamic5Qi *NonDynamic5Qi `mapstructure:"nonDynamic5Qi" json:"nonDynamic5Qi,omitempty"`
	Dynamic5Qi    *Dynamic5Qi    `mapstructure:"dynamic5Qi" json:"dynamic5Qi,omitempty"`
	Arp           *Arp           `mapstructure:"arp" json:"arp,omitempty"`
}
