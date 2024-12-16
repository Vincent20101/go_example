package pfcpie

// Table 7.4.3.1-2: Application ID's PFDs
type ApplicationIDsPFDs struct {
	ApplicationID *ApplicationID `tlv:"24" pres:"mandatory" json:"applicationID"`
	// FlowInactivityTimer is a self-defined IE, not shown in reference.  The tlv tag is
	// from OID(Object Identifier Descriptors) of Casa Systems Inc. which is
	// 1.3.6.1.4.1.20858.
	FlowInactivityTimer *FlowInactivityTimer `tlv:"32794" json:"flowInactivityTimer,omitempty"`
	PFD                 []*PFDContext        `tlv:"59" json:"pFD,omitempty"`
}

// Table 7.4.3.1-3: PFD context
type PFDContext struct {
	PFDContents []*PFDContents `tlv:"61" json:"pFDContents,omitempty"`
}

// 7.4.5.1.2 User Plane Path Failure Report IE within PFCP Node Report Request
type UserPlanePathFailure struct {
	RemoteGTPUPeer []*RemoteGTPUPeer `tlv:"103" pres:"mandatory" json:"remoteGTPUPeer"`
}

// 7.5.2.2 Create PDR IE within PFCP Session Establishment Request
type CreatePDR struct {
	PDRID                   *PacketDetectionRuleID     `tlv:"56" pres:"mandatory" json:"pDRID"`
	Precedence              *Precedence                `tlv:"29" json:"precedence,omitempty"`
	PDI                     *PDI                       `tlv:"2" pres:"mandatory" json:"pDI"`
	OuterHeaderRemoval      *OuterHeaderRemoval        `tlv:"95" json:"outerHeaderRemoval,omitempty"`
	FARID                   *FARID                     `tlv:"108" json:"fARID,omitempty"`
	URRID                   []*URRID                   `tlv:"81" json:"uRRID,omitempty"`
	QERID                   []*QERID                   `tlv:"109" json:"qERID,omitempty"`
	ActivatePredefinedRules *ActivatePredefinedRules   `tlv:"106" json:"activatePredefinedRules,omitempty"`
	Enterprise              *Enterprise                `tlv:"32781" json:"enterprise,omitempty"`
	MARID                   *MARID                     `tlv:"170" json:"mARID,omitempty"`
	MAI                     *MPTCPApplicableIndication `tlv:"265" json:"mAI,omitempty"`
	ActivationTime          *TimeStamp                 `tlv:"163" json:"activationTime,omitempty" cmp:"ignore"`
	DeactivationTime        *TimeStamp                 `tlv:"164" json:"deactivationTime,omitempty" cmp:"ignore"`
}

// Table 7.5.2.2-2: PDI IE within PFCP Session Establishment Request
type PDI struct {
	SourceInterface               *SourceInterface               `tlv:"20" pres:"mandatory" json:"sourceInterface"`
	LocalFTEID                    *FTEID                         `tlv:"21" json:"localFTEID,omitempty"`
	NetworkInstance               *Dnn                           `tlv:"22" json:"networkInstance,omitempty"`
	RTDetectionParameters         *RTDetectionParameters         `tlv:"255" json:"rTDetectionParameters,omitempty"`
	UEIPAddress                   *UEIPAddress                   `tlv:"93" json:"uEIPAddress,omitempty"`
	TrafficEndpointID             *TrafficEndpointID             `tlv:"131" json:"trafficEndpointID,omitempty"`
	SDFFilter                     []*SDFFilter                   `tlv:"23" json:"sDFFilter,omitempty"`
	ApplicationID                 *ApplicationID                 `tlv:"24" json:"applicationID,omitempty"`
	EthernetPDUSessionInformation *EthernetPDUSessionInformation `tlv:"142" json:"ethernetPDUSessionInformation,omitempty"`
	EthernetPacketFilter          []*EthernetPacketFilter        `tlv:"132" json:"ethernetPacketFilter,omitempty"`
	QFI                           []*QFI                         `tlv:"124" json:"qFI,omitempty"`
	FramedRoute                   *FramedRoute                   `tlv:"153" json:"framedRoute,omitempty"`
	FramedRouting                 *FramedRouting                 `tlv:"154" json:"framedRouting,omitempty"`
	FramedIPv6Route               *FramedIPv6Route               `tlv:"155" json:"framedIPv6Route,omitempty"`
}

// Table 7.5.2.2-3: Ethernet Packet Filter IE within PFCP Session Establishment Request
type EthernetPacketFilter struct {
	EthernetFilterID         *EthernetFilterID         `tlv:"138" json:"ethernetFilterID,omitempty"`
	EthernetFilterProperties *EthernetFilterProperties `tlv:"139" json:"ethernetFilterProperties,omitempty"`
	MACAddress               *MACAddress               `tlv:"133" json:"mACAddress,omitempty"`
	Ethertype                *Ethertype                `tlv:"136" json:"ethertype,omitempty"`
	CTAG                     *CTAG                     `tlv:"134" json:"cTAG,omitempty"`
	STAG                     *STAG                     `tlv:"135" json:"sTAG,omitempty"`
	SDFFilter                []*SDFFilter              `tlv:"23" json:"sDFFilter,omitempty"`
	STAG1                    *STAG1                    `tlv:"32793" json:"sTAG1,omitempty"`
	STAG2                    *STAG2                    `tlv:"32794" json:"sTAG2,omitempty"`
}

// Table 7.5.2.2-5: Redundant Transmission Detection Parameters IE in PDI
type RTDetectionParameters struct {
	LocalFTEIDForRT      *FTEID `tlv:"21" pres:"mandatory" json:"localFTEIDForRT"`
	NetworkInstanceForRT *Dnn   `tlv:"22" json:"networkInstanceForRT,omitempty"`
}

// 7.5.2.3 Create FAR IE within PFCP Session Establishment Request
type CreateFAR struct {
	FARID                  *FARID                         `tlv:"108" pres:"mandatory" json:"fARID"`
	ApplyAction            *ApplyAction                   `tlv:"44" pres:"mandatory" json:"applyAction"`
	ForwardingParameters   *ForwardingParametersIEInFAR   `tlv:"4" json:"forwardingParameters,omitempty"`
	DuplicatingParameters  *DuplicatingParameters         `tlv:"5" json:"duplicatingParameters,omitempty"`
	RTForwardingParameters *RTForwardingParametersIEInFar `tlv:"270" json:"rTForwardingParameters,omitempty"`
	BARID                  *BARID                         `tlv:"88" json:"bARID,omitempty"`
	Enterprise             *Enterprise                    `tlv:"32781" json:"enterprise,omitempty"`
	// Enterprise is a self-defined IE, not shown in reference.  The tlv tag is
	// from OID(Object Identifier Descriptors) of Casa Systems Inc. which is
	// 1.3.6.1.4.1.20858.
}

// Table 7.5.2.3-2: Forwarding Parameters IE in FAR
type ForwardingParametersIEInFAR struct {
	DestinationInterface  *DestinationInterface  `tlv:"42" pres:"mandatory" json:"destinationInterface"`
	NetworkInstance       *Dnn                   `tlv:"22" json:"networkInstance,omitempty"`
	RedirectInformation   *RedirectInformation   `tlv:"38" json:"redirectInformation,omitempty"`
	OuterHeaderCreation   *OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking *TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy      *ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
	// Enrichment of UL traffic. When present, it shall contain information for
	// header enrichment.
	HeaderEnrichment              []HeaderEnrichment             `tlv:"98" json:"headerEnrichment,omitempty"`
	LinkedTrafficEndpointID       *TrafficEndpointID             `tlv:"131" json:"linkedTrafficEndpointID,omitempty"`
	Proxying                      *Proxying                      `tlv:"137" json:"proxying,omitempty"`
	Enterprise                    *Enterprise                    `tlv:"32776" json:"enterprise,omitempty"`
	AdditionalOuterHeaderCreation *AdditionalOuterHeaderCreation `tlv:"32795" json:"vzWOuterHeaderCreation,omitempty"`
}

// Table 7.5.2.3-3: Duplicating Parameters IE in FAR
type DuplicatingParametersIEInFAR struct {
	DestinationInterface  *DestinationInterface  `tlv:"42" pres:"mandatory" json:"destinationInterface"`
	OuterHeaderCreation   *OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking *TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy      *ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
}

// Table 7.5.2.3-4: Redundant Transmission Forwarding Parameters IE in FAR
type RTForwardingParametersIEInFar struct {
	OuterHeaderCreation  *OuterHeaderCreation `tlv:"84" pres:"mandatory" json:"outerHeaderCreation"`
	NetworkInstanceForRT *Dnn                 `tlv:"22" json:"networkInstanceForRT,omitempty"`
}

// 7.5.2.4 Create URR IE within PFCP Session Establishment Request
type CreateURR struct {
	URRID                               *URRID                      `tlv:"81" pres:"mandatory" json:"uRRID"`
	MeasurementMethod                   *MeasurementMethod          `tlv:"62" pres:"mandatory,encnullable,decnullable" json:"measurementMethod"`
	MeasurementPeriod                   *MeasurementPeriod          `tlv:"64" json:"measurementPeriod,omitempty"`
	ReportingTriggers                   *ReportingTriggers          `tlv:"37" pres:"mandatory,encnullable,decnullable" json:"reportingTriggers"`
	VolumeThreshold                     *VolumeThreshold            `tlv:"31" json:"volumeThreshold,omitempty"`
	VolumeQuota                         *VolumeQuota                `tlv:"73" json:"volumeQuota,omitempty"`
	TimeThreshold                       *TimeThreshold              `tlv:"32" json:"timeThreshold,omitempty"`
	TimeQuota                           *TimeQuota                  `tlv:"74" json:"timeQuota,omitempty"`
	QuotaHoldingTime                    *QuotaHoldingTime           `tlv:"71" json:"quotaHoldingTime,omitempty"`
	DroppedDLTrafficThreshold           *DroppedDLTrafficThreshold  `tlv:"72" json:"droppedDLTrafficThreshold,omitempty"`
	QuotaValidityTime                   *QuotaValidityTime          `tlv:"181" json:"quotaValidityTime,omitempty"`
	MonitoringTime                      *MonitoringTime             `tlv:"33" json:"monitoringTime,omitempty"`
	EventQuota                          *EventQuota                 `tlv:"148" json:"eventQuota,omitempty"`
	EventThreshold                      *EventThreshold             `tlv:"149" json:"eventThreshold,omitempty"`
	SubsequentVolumeThreshold           *SubsequentVolumeThreshold  `tlv:"34" json:"subsequentVolumeThreshold,omitempty"`
	SubsequentTimeThreshold             *SubsequentTimeThreshold    `tlv:"35" json:"subsequentTimeThreshold,omitempty"`
	SubsequentVolumeQuota               *SubsequentVolumeQuota      `tlv:"121" json:"subsequentVolumeQuota,omitempty"`
	SubsequentTimeQuota                 *SubsequentTimeQuota        `tlv:"122" json:"subsequentTimeQuota,omitempty"`
	InactivityDetectionTime             *InactivityDetectionTime    `tlv:"36" json:"inactivityDetectionTime,omitempty"`
	LinkedURRID                         []*LinkedURRID              `tlv:"82" json:"linkedURRID,omitempty"`
	MeasurementInformation              *MeasurementInformation     `tlv:"100" json:"measurementInformation,omitempty"`
	TimeQuotaMechanism                  *TimeQuotaMechanism         `tlv:"115" json:"timeQuotaMechanism,omitempty"`
	AggregatedURRs                      []*AggregatedURRs           `tlv:"118" json:"aggregatedURRs,omitempty"`
	FARIDForQuotaAction                 *FARID                      `tlv:"108" json:"fARIDForQuotaAction,omitempty"`
	EthernetInactivityTimer             *EthernetInactivityTimer    `tlv:"146" json:"ethernetInactivityTimer,omitempty"`
	AdditionalMonitoringTime            []*AdditionalMonitoringTime `tlv:"147" json:"additionalMonitoringTime,omitempty"`
	NumberOfReport                      *NumberOfReport             `tlv:"182" json:"numberOfReport,omitempty"`
	Enterprise                          *Enterprise                 `tlv:"32781" json:"enterprise,omitempty"`
	ExemptedApplicationIDForQuotaAction *ApplicationID              `tlv:"24" json:"exemptedApplicationIDForQuotaAction,omitempty"`
}

// Table 7.5.2.4-2: Aggregated URRs
type AggregatedURRs struct {
	AggregatedURRID *AggregatedURRID `tlv:"120" pres:"mandatory" json:"aggregatedURRID"`
	Multiplier      *Multiplier      `tlv:"119" pres:"mandatory" json:"multiplier"`
}

// Table 7.5.2.4-3: Additional Monitoring Time
type AdditionalMonitoringTime struct {
	MonitoringTime            *MonitoringTime            `tlv:"33" pres:"mandatory" json:"monitoringTime"`
	SubsequentVolumeThreshold *SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold,omitempty"`
	SubsequentTimeThreshold   *SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold,omitempty"`
	SubsequentVolumeQuota     *SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota,omitempty"`
	SubsequentTimeQuota       *SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota,omitempty"`
}

// TODO: Chenglong: Not found in TS29244, don't know which field is mandatory.
type EventInformation struct {
	EventID        *EventID        `tlv:"150" json:"eventID,omitempty"`
	EventThreshold *EventThreshold `tlv:"151" json:"eventThreshold,omitempty"`
}

// 7.5.2.5 Create QER IE within PFCP Session Establishment Request
type CreateQER struct {
	QERID              *QERID              `tlv:"109" pres:"mandatory" json:"qERID"`
	QERCorrelationID   *QERCorrelationID   `tlv:"28" json:"qERCorrelationID,omitempty"`
	GateStatus         *GateStatus         `tlv:"25" pres:"mandatory" json:"gateStatus"`
	MaximumBitrate     *MBR                `tlv:"26" json:"maximumBitrate,omitempty"`
	GuaranteedBitrate  *GBR                `tlv:"27" json:"guaranteedBitrate,omitempty"`
	PacketRate         *PacketRate         `tlv:"94" json:"packetRate,omitempty"`
	DLFlowLevelMarking *DLFlowLevelMarking `tlv:"97" json:"dLFlowLevelMarking,omitempty"`
	QoSFlowIdentifier  *QFI                `tlv:"124" json:"qoSFlowIdentifier,omitempty"`
	ReflectiveQoS      *RQI                `tlv:"123" json:"reflectiveQoS,omitempty"`
	Enterprise         *Enterprise         `tlv:"32781" json:"enterprise,omitempty"`
	AveragingWindow    *AveragingWindow    `tlv:"157" json:"averagingWindow,omitempty"`
	// This IE shall be present if Token Accumalation Interval is configured.
	// Optional
	RefillInterval *RefillInterval `tlv:"32798" json:"refillInterval,omitempty"`
}

// 7.5.2.6 Create BAR IE within PFCP Session Establishment Request
type CreateBAR struct {
	BARID                          *BARID                          `tlv:"88" pres:"mandatory" json:"bARID"`
	DownlinkDataNotificationDelay  *DownlinkDataNotificationDelay  `tlv:"46" json:"downlinkDataNotificationDelay,omitempty"`
	SuggestedBufferingPacketsCount *SuggestedBufferingPacketsCount `tlv:"140" json:"suggestedBufferingPacketsCount,omitempty"`
	Enterprise                     *Enterprise                     `tlv:"32781" json:"enterprise,omitempty"`
}

// 7.5.2.7 Create Traffic Endpoint IE within PFCP Session Establishment Request
type CreateTrafficEndpoint struct {
	TrafficEndpointID             *TrafficEndpointID             `tlv:"131" pres:"mandatory" json:"trafficEndpointID"`
	LocalFTEID                    *FTEID                         `tlv:"21" json:"localFTEID,omitempty"`
	NetworkInstance               *Dnn                           `tlv:"22" json:"networkInstance,omitempty"`
	UEIPAddress                   *UEIPAddress                   `tlv:"93" json:"uEIPAddress,omitempty"`
	EthernetPDUSessionInformation *EthernetPDUSessionInformation `tlv:"142" json:"ethernetPDUSessionInformation,omitempty"`
	FramedRoute                   *FramedRoute                   `tlv:"153" json:"framedRoute,omitempty"`
	FramedRouting                 *FramedRouting                 `tlv:"154" json:"framedRouting,omitempty"`
	FramedIPv6Route               *FramedIPv6Route               `tlv:"155" json:"framedIPv6Route,omitempty"`
}

// 7.5.3.2 Created PDR IE within PFCP Session Establishment Response
type CreatedPDR struct {
	PDRID      *PacketDetectionRuleID `tlv:"56" pres:"mandatory" json:"pDRID"`
	LocalFTEID []*FTEID               `tlv:"21" json:"localFTEID,omitempty"`
	Enterprise *Enterprise            `tlv:"32781" json:"enterprise,omitempty"`
}

// 7.5.3.3 Load Control Information IE within PFCP Session Establishment Response
type LoadControlInformation struct {
	LoadControlSequenceNumber *SequenceNumber `tlv:"52" pres:"mandatory" json:"loadControlSequenceNumber"`
	LoadMetric                *Metric         `tlv:"53" pres:"mandatory" json:"loadMetric"`
}

// 7.5.3.4 Overload Control Information IE within PFCP Session Establishment Response
type OverloadControlInformation struct {
	OverloadControlSequenceNumber   *SequenceNumber `tlv:"52" pres:"mandatory" json:"overloadControlSequenceNumber"`
	OverloadReductionMetric         *Metric         `tlv:"53" pres:"mandatory" json:"overloadReductionMetric"`
	PeriodOfValidity                *Timer          `tlv:"55" pres:"mandatory" json:"periodOfValidity"`
	OverloadControlInformationFlags *OCIFlags       `tlv:"110" json:"overloadControlInformationFlags,omitempty"`
}

// 7.5.3.5 Created Traffic Endpoint IE within PFCP Session Establishment Response
type CreatedTrafficEndpoint struct {
	TrafficEndpointID *TrafficEndpointID `tlv:"131" pres:"mandatory" json:"trafficEndpointID"`
	LocalFTEID        *FTEID             `tlv:"21" json:"localFTEID,omitempty"`
}

// 7.5.4.2 Update PDR IE within PFCP Session Modification Request
type UpdatePDR struct {
	PDRID                     *PacketDetectionRuleID     `tlv:"56" pres:"mandatory" json:"pDRID"`
	OuterHeaderRemoval        *OuterHeaderRemoval        `tlv:"95" json:"outerHeaderRemoval,omitempty"`
	Precedence                *Precedence                `tlv:"29" json:"precedence,omitempty"`
	PDI                       *PDI                       `tlv:"2" json:"pDI,omitempty"`
	FARID                     *FARID                     `tlv:"108" json:"fARID,omitempty"`
	URRID                     []*URRID                   `tlv:"81" json:"uRRID,omitempty"`
	QERID                     []*QERID                   `tlv:"109" json:"qERID,omitempty"`
	ActivatePredefinedRules   *ActivatePredefinedRules   `tlv:"106" json:"activatePredefinedRules,omitempty"`
	DeactivatePredefinedRules *DeactivatePredefinedRules `tlv:"107" json:"deactivatePredefinedRules,omitempty"`
	ActivationTime            *TimeStamp                 `tlv:"163" json:"activationTime,omitempty" cmp:"ignore"`
	DeactivationTime          *TimeStamp                 `tlv:"164" json:"deactivationTime,omitempty" cmp:"ignore"`
}

// 7.5.4.3 Update FAR IE within PFCP Session Modification Request
// Table 7.5.4.3-1: Update FAR IE within PFCP Session Modification Request
type UpdateFAR struct {
	FARID                       *FARID                             `tlv:"108" pres:"mandatory" json:"fARID"`
	ApplyAction                 *ApplyAction                       `tlv:"44" json:"applyAction,omitempty"`
	UpdateForwardingParameters  *UpdateForwardingParametersIEInFAR `tlv:"11" json:"updateForwardingParameters,omitempty"`
	UpdateDuplicatingParameters *UpdateDuplicatingParameters       `tlv:"105" json:"updateDuplicatingParameters,omitempty"`
	RTForwardingParameters      *RTForwardingParametersIEInFar     `tlv:"270" json:"rTForwardingParameters,omitempty"`
	BARID                       *BARID                             `tlv:"88" json:"bARID,omitempty"`
}

// Table 7.5.4.3-2: Update Forwarding Parameters IE in FAR
type UpdateForwardingParametersIEInFAR struct {
	DestinationInterface    *DestinationInterface  `tlv:"42" json:"destinationInterface,omitempty"`
	NetworkInstance         *Dnn                   `tlv:"22" json:"networkInstance,omitempty"`
	RedirectInformation     *RedirectInformation   `tlv:"38" json:"redirectInformation,omitempty"`
	OuterHeaderCreation     *OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking   *TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy        *ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
	HeaderEnrichment        []HeaderEnrichment     `tlv:"98" json:"headerEnrichment,omitempty"`
	PFCPSMReqFlags          *PFCPSMReqFlags        `tlv:"49" json:"pFCPSMReqFlags,omitempty"`
	LinkedTrafficEndpointID *TrafficEndpointID     `tlv:"131" json:"linkedTrafficEndpointID,omitempty"`
}

// Table 7.5.4.3-3: Update Duplicating Parameters IE in FAR
type UpdateDuplicatingParametersIEInFAR struct {
	DestinationInterface  *DestinationInterface  `tlv:"42" json:"destinationInterface,omitempty"`
	OuterHeaderCreation   *OuterHeaderCreation   `tlv:"84" json:"outerHeaderCreation,omitempty"`
	TransportLevelMarking *TransportLevelMarking `tlv:"30" json:"transportLevelMarking,omitempty"`
	ForwardingPolicy      *ForwardingPolicy      `tlv:"41" json:"forwardingPolicy,omitempty"`
}

// 7.5.4.4 Update URR IE within PFCP Session Modification Request
type UpdateURR struct {
	URRID                     *URRID                     `tlv:"81" pres:"mandatory" json:"uRRID"`
	MeasurementMethod         *MeasurementMethod         `tlv:"62" json:"measurementMethod,omitempty"`
	ReportingTriggers         *ReportingTriggers         `tlv:"37" json:"reportingTriggers,omitempty"`
	MeasurementPeriod         *MeasurementPeriod         `tlv:"64" json:"measurementPeriod,omitempty"`
	VolumeThreshold           *VolumeThreshold           `tlv:"31" json:"volumeThreshold,omitempty"`
	VolumeQuota               *VolumeQuota               `tlv:"73" json:"volumeQuota,omitempty"`
	TimeThreshold             *TimeThreshold             `tlv:"32" json:"timeThreshold,omitempty"`
	TimeQuota                 *TimeQuota                 `tlv:"74" json:"timeQuota,omitempty"`
	QuotaHoldingTime          *QuotaHoldingTime          `tlv:"71" json:"quotaHoldingTime,omitempty"`
	DroppedDLTrafficThreshold *DroppedDLTrafficThreshold `tlv:"72" json:"droppedDLTrafficThreshold,omitempty"`
	QuotaValidityTime         *QuotaValidityTime         `tlv:"181" json:"quotaValidityTime,omitempty"`
	MonitoringTime            *MonitoringTime            `tlv:"33" json:"monitoringTime,omitempty"`
	EventQuota                *EventQuota                `tlv:"148" json:"eventQuota,omitempty"`
	EventThreshold            *EventThreshold            `tlv:"149" json:"eventThreshold,omitempty"`
	SubsequentVolumeThreshold *SubsequentVolumeThreshold `tlv:"34" json:"subsequentVolumeThreshold,omitempty"`
	SubsequentTimeThreshold   *SubsequentTimeThreshold   `tlv:"35" json:"subsequentTimeThreshold,omitempty"`
	SubsequentVolumeQuota     *SubsequentVolumeQuota     `tlv:"121" json:"subsequentVolumeQuota,omitempty"`
	SubsequentTimeQuota       *SubsequentTimeQuota       `tlv:"122" json:"subsequentTimeQuota,omitempty"`
	InactivityDetectionTime   *InactivityDetectionTime   `tlv:"36" json:"inactivityDetectionTime,omitempty"`
	LinkedURRID               *LinkedURRID               `tlv:"82" json:"linkedURRID,omitempty"`
	MeasurementInformation    *MeasurementInformation    `tlv:"100" json:"measurementInformation,omitempty"`
	TimeQuotaMechanism        *TimeQuotaMechanism        `tlv:"115" json:"timeQuotaMechanism,omitempty"`
	AggregatedURRs            *AggregatedURRs            `tlv:"118" json:"aggregatedURRs,omitempty"`
	FARIDForQuotaAction       *FARID                     `tlv:"108" json:"fARIDForQuotaAction,omitempty"`
	EthernetInactivityTimer   *EthernetInactivityTimer   `tlv:"146" json:"ethernetInactivityTimer,omitempty"`
	AdditionalMonitoringTime  *AdditionalMonitoringTime  `tlv:"147" json:"additionalMonitoringTime,omitempty"`
}

// 7.5.4.5 Update QER IE within PFCP Session Modification Request
type UpdateQER struct {
	QERID              *QERID              `tlv:"109" pres:"mandatory" json:"qERID"`
	QERCorrelationID   *QERCorrelationID   `tlv:"28" json:"qERCorrelationID,omitempty"`
	GateStatus         *GateStatus         `tlv:"25" json:"gateStatus,omitempty"`
	MaximumBitrate     *MBR                `tlv:"26" json:"maximumBitrate,omitempty"`
	GuaranteedBitrate  *GBR                `tlv:"27" json:"guaranteedBitrate,omitempty"`
	PacketRate         *PacketRate         `tlv:"94" json:"packetRate,omitempty"`
	DLFlowLevelMarking *DLFlowLevelMarking `tlv:"97" json:"dLFlowLevelMarking,omitempty"`
	QoSFlowIdentifier  *QFI                `tlv:"124" json:"qoSFlowIdentifier,omitempty"`
	ReflectiveQoS      *RQI                `tlv:"123" json:"reflectiveQoS,omitempty"`
}

// 7.5.4.6 Remove PDR IE within PFCP Session Modification Request
type RemovePDR struct {
	PDRID *PacketDetectionRuleID `tlv:"56" pres:"mandatory" json:"pDRID"`
}

// 7.5.4.7 Remove FAR IE within PFCP Session Modification Request
type RemoveFAR struct {
	FARID *FARID `tlv:"108" pres:"mandatory" json:"fARID"`
}

// 7.5.4.8 Remove URR IE within PFCP Session Modification Request
type RemoveURR struct {
	URRID *URRID `tlv:"81" pres:"mandatory" json:"uRRID"`
}

// 7.5.4.9 Remove QER IE PFCP Session Modification Request
type RemoveQERIEPFCPSessionModificationRequest struct {
	QERID *QERID `tlv:"109" pres:"mandatory" json:"qERID"`
}

// 7.5.4.10 Query URR IE within PFCP Session Modification Request
type QueryURR struct {
	URRID *URRID `tlv:"81" pres:"mandatory" json:"uRRID"`
}

// 7.5.4.12 Remove BAR IE within PFCP Session Modification Request
type RemoveBAR struct {
	BARID *BARID `tlv:"88" pres:"mandatory" json:"bARID"`
}

// 7.5.4.13 Update Traffic Endpoint IE within PFCP Session Modification Request
type UpdateTrafficEndpoint struct {
	TrafficEndpointID *TrafficEndpointID `tlv:"131" pres:"mandatory" json:"trafficEndpointID"`
	LocalFTEID        *FTEID             `tlv:"21" json:"localFTEID,omitempty"`
	NetworkInstance   *Dnn               `tlv:"22" json:"networkInstance,omitempty"`
	UEIPAddress       *UEIPAddress       `tlv:"93" json:"uEIPAddress,omitempty"`
	FramedRoute       *FramedRoute       `tlv:"153" json:"framedRoute,omitempty"`
	FramedRouting     *FramedRouting     `tlv:"154" json:"framedRouting,omitempty"`
	FramedIPv6Route   *FramedIPv6Route   `tlv:"155" json:"framedIPv6Route,omitempty"`
}

// 7.5.4.14 Remove Traffic Endpoint IE within PFCP Session Modification Request
type RemoveTrafficEndpoint struct {
	TrafficEndpointID *TrafficEndpointID `tlv:"131" pres:"mandatory" json:"trafficEndpointID"`
}

// 7.5.8.2 Downlink Data Report IE within PFCP Session Report Request
type DownlinkDataReport struct {
	PDRID                          *PacketDetectionRuleID          `tlv:"56" pres:"mandatory" json:"pDRID"`
	DownlinkDataServiceInformation *DownlinkDataServiceInformation `tlv:"45" json:"downlinkDataServiceInformation,omitempty"`
}

// TODO: Chenglong: Not found in TS29244, don't know which field is mandatory.
type EventReporting struct {
	EventID *EventID `tlv:"150" json:"eventID,omitempty"`
}

// Table 7.5.8.3-3: Ethernet Traffic Information IE within Usage Report IE
type EthernetTrafficInformation struct {
	MACAddressesDetected *MACAddressesDetected `tlv:"144" json:"mACAddressesDetected,omitempty"`
	MACAddressesRemoved  *MACAddressesRemoved  `tlv:"145" json:"mACAddressesRemoved,omitempty"`
}

// 7.5.8.4 Error Indication Report IE within PFCP Session Report Request
type ErrorIndicationReport struct {
	RemoteFTEID *FTEID `tlv:"21" pres:"mandatory" json:"remoteFTEID"`
}

// Table 7.5.8.3-2: Application Detection Information IE within Usage Report IE
type ApplicationDetectionInformation struct {
	ApplicationID         *ApplicationID         `tlv:"24" pres:"mandatory" json:"applicationID"`
	ApplicationInstanceID *ApplicationInstanceID `tlv:"91" json:"applicationInstanceID,omitempty"`
	FlowInformation       *FlowInformation       `tlv:"92" json:"flowInformation,omitempty"`
}

// 7.5.4.11 Update BAR IE within PFCP Session Modification Request
type UpdateBARPFCPSessionModificationRequest struct {
	BARID                          *BARID                          `tlv:"88" pres:"mandatory" json:"bARID"`
	DownlinkDataNotificationDelay  *DownlinkDataNotificationDelay  `tlv:"46" json:"downlinkDataNotificationDelay,omitempty"`
	SuggestedBufferingPacketsCount *SuggestedBufferingPacketsCount `tlv:"140" json:"suggestedBufferingPacketsCount,omitempty"`
}

// 7.5.5.2 Usage Report IE within PFCP Session Modification Response
type UsageReport struct {
	URRID                           *URRID                           `tlv:"81" pres:"mandatory" json:"uRRID"`
	URSEQN                          *URSEQN                          `tlv:"104" pres:"mandatory" json:"uRSEQN"`
	UsageReportTrigger              *UsageReportTrigger              `tlv:"63" pres:"mandatory" json:"usageReportTrigger"`
	StartTime                       *StartTime                       `tlv:"75" json:"startTime,omitempty"`
	EndTime                         *EndTime                         `tlv:"76" json:"endTime,omitempty"`
	VolumeMeasurement               *VolumeMeasurement               `tlv:"66" json:"volumeMeasurement,omitempty"`
	DurationMeasurement             *DurationMeasurement             `tlv:"67" json:"durationMeasurement,omitempty"`
	ApplicationDetectionInformation *ApplicationDetectionInformation `tlv:"68" json:"applicationDetectionInformation,omitempty"`
	UEIPAddress                     *UEIPAddress                     `tlv:"93" json:"uEIPAddress,omitempty"`
	NetworkInstance                 *Dnn                             `tlv:"22" json:"networkInstance,omitempty"`
	TimeOfFirstPacket               *TimeOfFirstPacket               `tlv:"69" json:"timeOfFirstPacket,omitempty"`
	TimeOfLastPacket                *TimeOfLastPacket                `tlv:"70" json:"timeOfLastPacket,omitempty"`
	UsageInformation                *UsageInformation                `tlv:"90" json:"usageInformation,omitempty"`
	QueryURRReference               *QueryURRReference               `tlv:"125" json:"queryURRReference,omitempty"`
	EventTimeStamp                  []*TimeStamp                     `tlv:"156" json:"eventTimeStamp,omitempty"`
	EthernetTrafficInformation      *EthernetTrafficInformation      `tlv:"143" json:"ethernetTrafficInformation,omitempty"`
}

// 7.5.5.5 Updated PDR IE within PFCP Session Modification Response
type UpdatedPDR struct {
	PDRID           *PacketDetectionRuleID `tlv:"56" pres:"mandatory" json:"pDRID"`
	LocalFTEIDForRT *FTEID                 `tlv:"21" json:"localFTEIDForRT,omitempty"`
}

// 7.5.2.8 Create MAR IE within PFCP Session Establishment Request
type CreateMAR struct {
	MARID                 *MARID                 `tlv:"170" pres:"mandatory" json:"mARID"`
	SteeringFunctionality *SteeringFunctionality `tlv:"171" pres:"mandatory" json:"steeringFunctionality"`
	SteeringMode          *SteeringMode          `tlv:"172" pres:"mandatory" json:"steeringMode"`
	Afai3GPP              *Afai3GPP              `tlv:"166" json:"afai3GPP,omitempty"`
	AfaiNon3GPP           *AfaiNon3GPP           `tlv:"167" json:"afaiNon3GPP,omitempty"`
}

// Table 7.5.4.16-2: Update 3GPP Access Forwarding Action Information IE in Update MAR IE
type Afai3GPP struct {
	FARID    *FARID    `tlv:"108" json:"fARID,omitempty"`
	Weight   *Weight   `tlv:"173" json:"weight,omitempty"`
	Priority *Priority `tlv:"174" json:"priority,omitempty"`
	URRID    []*URRID  `tlv:"81" json:"uRRID,omitempty"`
}

type AfaiNon3GPP Afai3GPP

// 7.5.2.10 Provide ATSSS Control Information IE within PFCP Session Establishment Request
type AtsssControlInfo struct {
	MptcpControlInfo   *MptcpControlInfo   `tlv:"222" json:"mptcpControlInfo,omitempty"`
	AtsssLLControlInfo *AtsssLLControlInfo `tlv:"223" json:"atsssLLControlInfo,omitempty"`
	PmfControlInfo     *PmfControlInfo     `tlv:"224" json:"pmfControlInfo,omitempty"`
}

// 7.5.3.7 ATSSS Control Parameters IE within PFCP Session Establishment Response
type AtsssControlParameters struct {
	MptcpParameters   *MptcpParameters   `tlv:"225" json:"mptcpParameters,omitempty"`
	AtsssLLParameters *AtsssLLParameters `tlv:"226" json:"atsssLLParameters,omitempty"`
	PmfParameters     *PmfParameters     `tlv:"227" json:"pmfParameters,omitempty"`
}

// Table 7.5.3.7-2: MPTCP Parameters IE within PFCP Session Establishment Response
type MptcpParameters struct {
	MptcpAddrInfo    *MptcpAddrInfo    `tlv:"228" pres:"mandatory" json:"mptcpAddrInfo"`
	UeLinkSpecIPAddr *UeLinkSpecIPAddr `tlv:"229" pres:"mandatory" json:"ueLinkSpecIPAddr"`
}

// Table 7.5.3.7-3: ATSSS-LL Parameters IE within PFCP Session Establishment Response
type AtsssLLParameters struct {
	AtsssLLInfo *AtsssLLInfo `tlv:"231" pres:"mandatory" json:"atsssLLInfo"`
}

// Table 7.5.3.7-4: PMF Parameters IE within PFCP Session Establishment Response
type PmfParameters struct {
	PmfAddrInfo *PmfAddrInfo `tlv:"230" pres:"mandatory" json:"pmfAddrInfo"`
}

// 7.5.4.15 Remove MAR IE within PFCP Session Modification Request
type RemoveMAR struct {
	MARID *MARID `tlv:"170" pres:"mandatory" json:"mARID"`
}

type UpdateAfai3GPP Afai3GPP
type UpdateAfaiNon3GPP Afai3GPP

// 7.5.4.16 Update MAR IE within PFCP Session Modification Request
type UpdateMAR struct {
	MARID                 *MARID                 `tlv:"170" pres:"mandatory" json:"mARID"`
	SteeringFunctionality *SteeringFunctionality `tlv:"171" json:"steeringFunctionality,omitempty"`
	SteeringMode          *SteeringMode          `tlv:"172" json:"steeringMode,omitempty"`
	UpdateAfai3GPP        *UpdateAfai3GPP        `tlv:"175" json:"updateAfai3GPP,omitempty"`
	UpdateAfaiNon3GPP     *UpdateAfaiNon3GPP     `tlv:"176" json:"updateAfaiNon3GPP,omitempty"`
	Afai3GPP              *Afai3GPP              `tlv:"166" json:"afai3GPP,omitempty"`
	AfaiNon3GPP           *AfaiNon3GPP           `tlv:"167" json:"afaiNon3GPP,omitempty"`
}

// 7.5.2.9	Create SRR IE within PFCP Session Establishment Request
type CreateSRR struct {
	SRRID                 *SRRID                      `tlv:"215" pres:"mandatory" json:"sRRID"`
	QosMonPerQosFlowInfos []*QosMonPerQosFlowCtrlInfo `tlv:"242" json:"qosMonPerQosFlowInfos,omitempty"`
}

// Table 7.5.2.9-3: QoS Monitoring per QoS flow Control Information
type QosMonPerQosFlowCtrlInfo struct {
	QFIs              []*QFI                  `tlv:"124" pres:"mandatory" json:"qFIs"`
	ReqQosMon         *RequestedQosMonitoring `tlv:"243" pres:"mandatory" json:"reqQosMon"`
	ReportingFreq     *ReportingFrequency     `tlv:"244" pres:"mandatory" json:"reportingFreq"`
	PktDelayThresh    *PacketDelayThreshold   `tlv:"245" json:"pktDelayThresh,omitempty"`
	MinWaitTime       *MinimumWaitTime        `tlv:"246" json:"minWaitTime,omitempty"`
	MeasurementPeriod *MeasurementPeriod      `tlv:"64" json:"measurementPeriod,omitempty"`
}

// 7.5.4.20	Update SRR IE within PFCP Session Modification Request
type UpdateSRR CreateSRR

// 7.5.4.19	Remove SRR IE within PFCP Session Modification Request
type RemoveSRR struct {
	SRRID *SRRID `tlv:"215" pres:"mandatory" json:"sRRID"`
}

// 7.5.8.6	Session Report IE within PFCP Session Report Request
type SessionReport struct {
	SRRID        *SRRID                 `tlv:"215" pres:"mandatory" json:"sRRID"`
	QosMonReport []*QosMonitoringReport `tlv:"247" json:"qosMonReport,omitempty"`
}

// Table 7.5.8.6-3: QoS Monitoring Report IE
type QosMonitoringReport struct {
	QFI               *QFI                      `tlv:"124" pres:"mandatory" json:"qFI"`
	QosMonMeasurement *QosMonitoringMeasurement `tlv:"248" pres:"mandatory" json:"qosMonMeasurement"`
	TimeStamp         *TimeStamp                `tlv:"156" pres:"mandatory" json:"timeStamp"`
	StartTime         *StartTime                `tlv:"75" json:"startTime,omitempty"`
}

// type UsageReportPFCPSessionModificationResponse struct {
//	URRID                      *URRID                      `tlv:"81" json:"//,omitempty"`
//	URSEQN                     *URSEQN                     `tlv:"104" json:"//,omitempty"`
//	UsageReportTrigger         *UsageReportTrigger         `tlv:"63" json:"//,omitempty"`
//	StartTime                  *StartTime                  `tlv:"75" json:"//,omitempty"`
//	EndTime                    *EndTime                    `tlv:"76" json:"//,omitempty"`
//	VolumeMeasurement          *VolumeMeasurement          `tlv:"66" json:"//,omitempty"`
//	DurationMeasurement        *DurationMeasurement        `tlv:"67" json:"//,omitempty"`
//	TimeOfFirstPacket          *TimeOfFirstPacket          `tlv:"69" json:"//,omitempty"`
//	TimeOfLastPacket           *TimeOfLastPacket           `tlv:"70" json:"//,omitempty"`
//	UsageInformation           *UsageInformation           `tlv:"90" json:"//,omitempty"`
//	QueryURRReference          *QueryURRReference          `tlv:"125" json:"//,omitempty"`
//	EthernetTrafficInformation *EthernetTrafficInformation `tlv:"143" json:"//,omitempty"`
// }

// type UsageReportPFCPSessionDeletionResponse struct {
//	URRID                      *URRID                      `tlv:"81" json:"//,omitempty"`
//	URSEQN                     *URSEQN                     `tlv:"104" json:"//,omitempty"`
//	UsageReportTrigger         *UsageReportTrigger         `tlv:"63" json:"//,omitempty"`
//	StartTime                  *StartTime                  `tlv:"75" json:"//,omitempty"`
//	EndTime                    *EndTime                    `tlv:"76" json:"//,omitempty"`
//	VolumeMeasurement          *VolumeMeasurement          `tlv:"66" json:"//,omitempty"`
//	DurationMeasurement        *DurationMeasurement        `tlv:"67" json:"//,omitempty"`
//	TimeOfFirstPacket          *TimeOfFirstPacket          `tlv:"69" json:"//,omitempty"`
//	TimeOfLastPacket           *TimeOfLastPacket           `tlv:"70" json:"//,omitempty"`
//	UsageInformation           *UsageInformation           `tlv:"90" json:"//,omitempty"`
//	EthernetTrafficInformation *EthernetTrafficInformation `tlv:"143" json:"//,omitempty"`
// }

// type UsageReportPFCPSessionReportRequest struct {
//	URRID                           *URRID                           `tlv:"81" json:"//,omitempty"`
//	URSEQN                          *URSEQN                          `tlv:"104" json:"//,omitempty"`
//	UsageReportTrigger              *UsageReportTrigger              `tlv:"63" json:"//,omitempty"`
//	StartTime                       *StartTime                       `tlv:"75" json:"//,omitempty"`
//	EndTime                         *EndTime                         `tlv:"76" json:"//,omitempty"`
//	VolumeMeasurement               *VolumeMeasurement               `tlv:"66" json:"//,omitempty"`
//	DurationMeasurement             *DurationMeasurement             `tlv:"67" json:"//,omitempty"`
//	ApplicationDetectionInformation *ApplicationDetectionInformation `tlv:"68" json:"//,omitempty"`
//	UEIPAddress                     *UEIPAddress                     `tlv:"93" json:"//,omitempty"`
//	NetworkInstance                 *Dnn                             `tlv:"22" json:"//,omitempty"`
//	TimeOfFirstPacket               *TimeOfFirstPacket               `tlv:"69" json:"//,omitempty"`
//	TimeOfLastPacket                *TimeOfLastPacket                `tlv:"70" json:"//,omitempty"`
//	UsageInformation                *UsageInformation                `tlv:"90" json:"//,omitempty"`
//	QueryURRReference               *QueryURRReference               `tlv:"125" json:"//,omitempty"`
//	EventTimeStamp                  []*TimeStamp                      `tlv:"156" json:"//,omitempty"`
//	EthernetTrafficInformation      *EthernetTrafficInformation      `tlv:"143" json:"//,omitempty"`
// }

// 7.4.5.1.4 Clock Drift Report IE within PFCP Node Report Request
type ClockDriftReport struct {
	TSNTimeDomainNumber            *TSNTimeDomainNumber            `tlv:"206" json:"tSNTimeDomainNumber,omitempty"`
	TimeOffsetMeasurement          *TimeOffsetMeasurement          `tlv:"209" json:"timeOffsetMeasurement,omitempty"`
	CumulativeRateRatioMeasurement *CumulativeRateRatioMeasurement `tlv:"210" json:"cumulativeRateRatioMeasurement,omitempty"`
	TimeStamp                      *TimeStamp                      `tlv:"156" json:"timeStamp,omitempty"`
}

// 7.4.4.1.2 Clock Drift Control Information IE within PFCP Association Setup Request
type ClockDriftControlInformation struct {
	TSNTimeDomainNumber            *TSNTimeDomainNumber            `tlv:"206" json:"tSNTimeDomainNumber,omitempty"`
	RequestedClockDriftInformation *RequestedClockDriftInformation `tlv:"204" json:"requestedClockDriftInformation,omitempty"`
	TimeOffsetThreshold            *TimeOffsetThreshold            `tlv:"207" json:"timeOffsetThreshold,omitempty"`
	CumulativeRateRatioThreshold   *CumulativeRateRatioThreshold   `tlv:"208" json:"cumulativeRateRatioThreshold,omitempty"`
}

// 7.5.3.6 Created Bridge Info for TSC IE within PFCP Session Establishment Response
type CreatedBridgeInfoForTSC struct {
	DSTTPortNumber *DSTTPortNumber `tlv:"196" json:"dSTTPortNumber,omitempty"`
	TSNBridgeID    *TSNBridgeID    `tlv:"198" json:"tSNBridgeID,omitempty"`
}

// 7.5.4.18 TSC Management Information IE within PFCP Session Modification Request
type TSCManagementInformation struct {
	PortManagementInformationContainer   *PortManagementInformationContainer   `tlv:"202" json:"portManagementInformationContainer,omitempty"`
	BridgeManagementInformationContainer *BridgeManagementInformationContainer `tlv:"266" json:"bridgeManagementInformationContainer,omitempty"`
	NWTTPortNumber                       *NWTTPortNumber                       `tlv:"197" json:"nWTTPortNumber,omitempty"`
}
