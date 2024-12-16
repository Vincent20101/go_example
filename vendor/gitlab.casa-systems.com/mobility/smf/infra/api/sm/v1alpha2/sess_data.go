package v1alpha2

import (
	"time"
)

type SessionKey struct {
	Supi         string      `json:"supi,omitempty"`
	PduSessionId int32       `json:"pduSessionId,omitempty"`
	Gpsi         string      `json:"gpsi,omitempty"`
	SessionType  SessionType `json:"sessionType,omitempty"` // Indicate a 3G or 4G or 5G NSA session for SMF+GTP
}
type SessionType uint8

const (
	SessionType5G SessionType = iota
	SessionType4G
	SessionType3G
)

type UeInfo struct {
	Supi                string `json:"supi,omitempty"`
	UnauthenticatedSupi bool   `json:"unauthSupi,omitempty"`
	Pei                 string `json:"pei,omitempty"`
	Gpsi                string `json:"gpsi,omitempty"`
	SdmRefId            string `json:"sdmRefId,omitempty"`
	//4G UE parameters
	PeerRecovery uint32 `json:"peerRecovery,omitempty"`
}

type IpAddress struct {
	Ipv4Addr   string `json:"ipv4Addr,omitempty"`
	Ipv6Addr   string `json:"ipv6Addr,omitempty"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
}

type SessIpAddr struct {
	PduSessionType string    `json:"pduSessionType,omitempty"`
	Addr           IpAddress `json:"addr"`
	Static         bool      `json:"static"`
}

type PfcpFseid struct {
	Addr IpAddress `json:"addr"`
	Seid uint64    `json:"seid"`
}

type PfcpFteid struct {
	Addr IpAddress `json:"addr"`
	Teid uint64    `json:"teid"`
}

type OptAlwaysOnType uint8

const (
	OPT_ALWAYSON_NONE    OptAlwaysOnType = 0x0 //None operation.
	OPT_ALWAYSON_EST     OptAlwaysOnType = 0x1 //Operation Establish Always-on in PDU Session Establishment procedure.
	OPT_ALWAYSON_MOD     OptAlwaysOnType = 0x2 //Operation Need change to an Always-on in PDU Session Modification procedure.
	OPT_ALWAYSON_ENABLED OptAlwaysOnType = 0x3 //Marked as always-on dynamic enablement.
)

type SessionRsc struct {
	DefPsaUpf    UpfRsc   `json:"defPsaUpf,omitempty"`  //for default PSA
	IUpf         []UpfRsc `json:"iUpf,omitempty"`       //for ULCL/BP or other Iupf
	DnaiPsaUpf   []UpfRsc `json:"dnaiPsaUpf,omitempty"` //for additional PSA
	SmfTeid      uint32   `json:"smfTeid,omitempty"`    //smf gtpu tunnel teid between smf and upf for ipv6 SLAAC RS and RA
	UdmRef       string   `json:"udmRef,omitempty"`
	IsDynamicPcc bool     `json:"isDynamicPcc,omitempty"`
	PcfRef       string   `json:"pcfRef,omitempty"`
	PcfHostPort  string   `json:"pcfHostPort,omitempty"`
	PcfFromAmf   string   `json:"pcfFromAmf,omitempty"` // amf provided pcf instance id
	IsChargingOn bool     `json:"isChargingOn,omitempty"`
	ChfRef       string   `json:"chfRef,omitempty"`
	ChfHostPort  string   `json:"chfHostPort,omitempty"`
	//ChfApiPath              string
	ChfId                   uint32                     `json:"chfId,omitempty"`
	ChgType                 ConvergeChargingType       `json:"chgType,omitempty"`
	PcscfFqdn               string                     `json:"pcscfFqdn,omitempty"`
	UecmHostPort            string                     `json:"uecmHostPort,omitempty"`
	SdmHostPort             string                     `json:"sdmHostPort,omitempty"`
	ReleaseCause            string                     `json:"releaseCause,omitempty"`
	State                   string                     `json:"state,omitempty"`
	TxnId                   uint8                      `json:"txnId,omitempty"`
	PcfReqTriggers          uint32                     `json:"pcfReqTriggers,omitempty"` //pcf request trigger bit mask, see definition in sess_policy_utils.go
	EpsInterworkingInd      *EpsInterworkingIndication `json:"epsInterworkingInd,omitempty"`
	ChargingCharacteristics string                     `json:"chargingCharacteristics,omitempty"` //udm returned TODO: wuyubin choise another place
	ServiceAgent            []SdServiceAgent           `json:"serviceAgent,omitempty"`
	OptAlwaysOn             OptAlwaysOnType            `json:"optAlwaysOn,omitempty"` //Operation Always-on Pdu Session
	IsEmergency             bool                       `json:"isEmergency,omitempty"` //Emergency Session
	IsActive                bool                       `json:"isActive,omitempty"`
	NeedUnpauseCharging     bool                       `json:"needUnpauseCharging,omitempty"` //-true: need send update urr to upf to unpause Charging
	//virtual apn  must have a configurable option to allow the inclusion of the original DNN (submitted by UE or selected based on the UDM subscription) in the procedures with other NFs
	IsAllowOriginalDNN   bool        `json:"isAllowOriginalDNN,omitempty"`
	OriginalDNN          string      `json:"originalDNN,omitempty"` //save the original DNN
	SbiMessagePriority   string      `json:"sbiMessagePriority,omitempty"`
	DhcpData             *DhcpData   `json:"dhcpData,omitempty"`             //DHCP data, includes v4 and v6
	DefPccRuleFromConfig bool        `json:"defPccRuleFromConfig,omitempty"` //Indicates whether the pcc rule bound to default flow comes from the configuration
	TimeEvents           *TimeEvents `json:"timeEvents,omitempty"`           // save the session start/stop time, and indicate session has been terminated.
	PcfNfInstanceId      *string     `json:"pcfNfInstanceId,omitempty"`      // The PCF instance of the current Session connection
	UdmNfInstanceId      *string     `json:"udmNfInstanceId,omitempty"`      // The UDM(UECM UDM) instance of the current Session connection
	ChfNfInstanceId      *string     `json:"chfNfInstanceId,omitempty"`      // The CHF instance of the current Session connection
	AmfNfInstanceId      *string     `json:"amfNfInstanceId,omitempty"`      // The AMF instance of the current Session connection
	IdleTimeout          uint32      `json:"idleTimeout,omitempty"`          // If there is still no control plane signaling interaction after this duration, the session is considered stale and a core network resource release procedure is triggered. It is in seconds and is pre-configured locally or dynamically configured by the udm. When it is 0, this feature is deactivated.
}

type TimeEvents struct {
	// session start/stop time
	SessStartTime     time.Time  `json:"sessStartTime,omitempty"`
	SessStopTime      *time.Time `json:"sessStopTime,omitempty"`
	SessStopIndicator bool       `json:"sessStopIndicator,omitempty"`
}

type EpsInterworkingIndication = string

const (
	EpsInterworkingIndication_NONE        EpsInterworkingIndication = "NONE"
	EpsInterworkingIndication_WITH_N26    EpsInterworkingIndication = "WITH_N26"
	EpsInterworkingIndication_WITHOUT_N26 EpsInterworkingIndication = "WITHOUT_N26"
)

type ConvergeChargingType uint8

const (
	CHARGING_TYPE_NONE                    ConvergeChargingType = 0x0
	CHARGING_TYPE_ONLINE                  ConvergeChargingType = 0x1
	CHARGING_TYPE_OFFLINE                 ConvergeChargingType = 0x2
	CHARGING_TYPE_OFFLINE_ONLINE_CONVERGE ConvergeChargingType = 0x3
)

type UpfRsc struct {
	InstanceId           string         `json:"instanceId,omitempty"`
	UpfType              UpfRscType     `json:"upfType,omitempty"`
	SessIp               *SessIpAddr    `json:"sessIp,omitempty"` //mandatory for anchor PSA, optional for dnai-PSA(multi-homing)
	IpPoolV4             string         `json:"ipPoolV4,omitempty"`
	IpPoolV6             string         `json:"ipPoolV6,omitempty"`
	LFseid               *PfcpFseid     `json:"lFseid,omitempty"`
	RFseid               *PfcpFseid     `json:"rFseid,omitempty"`
	GnbFteid             *PfcpFteid     `json:"gnbFteid,omitempty"`             //GnB N3 tunnelId
	N3Fteid              *PfcpFteid     `json:"n3Fteid,omitempty"`              //Local N3 tunnelId
	N9LocalFteid         *PfcpFteid     `json:"n9LocalFteid,omitempty"`         //for I-UPF (UL-CL/BP UPF)
	N9RemoteFteid        *PfcpFteid     `json:"n9RemoteFteid,omitempty"`        //for I-UPF (UL-CL/BP UPF)
	IndTunnelLocalFteid  *PfcpFteid     `json:"indTunnelLocalFteid,omitempty"`  //for N2 HO indirect tunnel teid
	IndTunnelRemoteFteid *PfcpFteid     `json:"indTunnelRemoteFteid,omitempty"` //NGRAN indirect tunnel teid for both N2 and IRAT
	DefFlow              *UpfFlowInfo   `json:"defFlow,omitempty"`              //default flow
	DedicatedFlow        []*UpfFlowInfo `json:"dedicatedFlow,omitempty"`
	NextId               NextIds        `json:"nextId,omitempty"`   //for ID allocation
	NeedRemv             bool           `json:"needRemv,omitempty"` //for pcf notify flows terminated determine to remove iupf and dnaipsaupf
	N3IntfAddrType       string         `json:"n3IntfAddrType,omitempty"`
	N9IntfAddrType       string         `json:"n9IntfAddrType,omitempty"`
	PdrIndirect          uint32         `json:"pdrIndirect,omitempty"`       //For N2 HO. One indirect Pdr per session
	FarIndirect          uint32         `json:"farIndirect,omitempty"`       //For N2 HO. One indirect Far per session
	QerIndirect          uint32         `json:"qerIndirect,omitempty"`       //For N2 HO. One indirect Qer per session
	SupportedFeatures    uint64         `json:"supportedFeatures,omitempty"` //Features supported by this upf.
}

type ServiceAgentType int

const (
	SERVICE_TYPE_UDM_SDM ServiceAgentType = iota
	SERVICE_TYPE_UDM_UECM
	SERVICE_TYPE_PCF
)

type SdServiceAgent struct {
	ServiceType ServiceAgentType `json:"serviceType,omitempty"`
	TgppApiRoot string           `json:"tgppApiRoot,omitempty"` // maps to 3gpp-Sbi-Target-apiRoot query
	AxyomCbUri  string           `json:"axyomCbUri,omitempty"`  // maps to x-axyom-callback-ur header
	AxyomNfId   string           `json:"axyomNfId,omitempty"`   // maps to x-axyom-nfid header
}

type UpfRscType uint8

const (
	UPF_RSC_TYPE_UNDEFINED UpfRscType = 0x0
	UPF_RSC_TYPE_PSA       UpfRscType = 0x1
	UPF_RSC_TYPE_S_UPF     UpfRscType = 0x2
	UPF_RSC_TYPE_T_UPF     UpfRscType = 0x4
	UPF_RSC_TYPE_BP        UpfRscType = 0x8
	UPF_RSC_TYPE_ULCL      UpfRscType = 0x10
	UPF_RSC_TYPE_I_UPF     UpfRscType = 0x20
)

type IpAddr struct {
	Ipv4 uint32 `json:"ipv4,omitempty"`
	Ipv6 []byte `json:"ipv6,omitempty"`
}

type PfcpIeFteid struct {
	Flag     uint32  `json:"flag,omitempty"`
	ChooseId uint32  `json:"choose_id,omitempty"`
	Teid     uint32  `json:"teid,omitempty"`
	IpAddr   *IpAddr `json:"ip_addr,omitempty"`
	Port     uint32  `json:"port,omitempty"`
}

type PfcpIeFseid struct {
	Flag   uint32  `json:"flag,omitempty"`
	Seid   uint64  `json:"seid,omitempty"`
	IpAddr *IpAddr `json:"ip_addr,omitempty"`
	Port   uint32  `json:"port,omitempty"`
}

type AdcUrr struct {
	Id         uint32 `json:"id,omitempty"`
	MuteNotify bool   `json:"muteNotify,omitempty"`
}

type UpfFlowInfo struct {
	QosRuleId  uint8    `json:"qosRuleId,omitempty"` //only used for dedicated PccRule, use the integer qos rule ID for efficiency purpose
	PdrUL      uint32   `json:"pdrUL,omitempty"`
	PdrDL      uint32   `json:"pdrDL,omitempty"`
	FarUL      uint32   `json:"farUL,omitempty"`
	FarDL      uint32   `json:"farDL,omitempty"`
	QerUL      uint32   `json:"qerUL,omitempty"`
	QerDL      uint32   `json:"qerDL,omitempty"`
	QerULDef   uint32   `json:"qerULDef,omitempty"`
	QerDLDef   uint32   `json:"qerDLDef,omitempty"`
	ChgUrrs    []uint32 `json:"chgUrrs,omitempty"` // URRs for online/offline charging
	AdcUrr     *AdcUrr  `json:"adcUrr,omitempty"`  // URR for application detection control
	UmUrrs     []uint32 `json:"umUrrs,omitempty"`  // URRs for usage monitoring
	EnetUrr    uint32   `json:"enetUrr,omitempty"` // URR for Ethernet pdu session type report mac address
	Bar        uint32   `json:"bar,omitempty"`
	Rg         uint32   `json:"rg,omitempty"`         // rating group
	Precedence uint32   `json:"precedence,omitempty"` //RK: Do we need to add precedence separately for UL and DL?
	Qfi        uint8    `json:"qfi,omitempty"`        //used for 4G5G ho
	Ebi        uint8    `json:"ebi,omitempty"`        //used for 4G5G ho
	PccRuleId  string   `json:"pccRuleId,omitempty"`  //used for 4G5G ho. This contain sessRule for defFlow.
	NeedRemv   bool     `json:"needRemv,omitempty"`   //used for sscMode3 Multi_homed to remove old psa related flow
}

type NextIds struct {
	Pdr     uint32 `json:"pdr,omitempty"`
	Far     uint32 `json:"far,omitempty"`
	Qer     uint32 `json:"qer,omitempty"`
	AdcUrr  uint32 `json:"adcUrr,omitempty"`
	UmUrr   uint32 `json:"umUrr,omitempty"`
	EnetUrr uint32 `json:"enetUrr,omitempty"` // URR for Ethernet pdu session type report mac address
	Bar     uint32 `json:"bar,omitempty"`
}

type UeLocation struct {
	UeLocation     *UserLocation `json:"ueLocation,omitempty"`
	UeTimeZone     string        `json:"ueTimeZone,omitempty"`
	AddUeLocation  *UserLocation `json:"addUeLocation,omitempty"`
	Ts             int64         `json:"ts,omitempty"`
	PresenceInLadn PresenceState `json:"presenceInLadn,omitempty"`
}

type BackupAmfInfo struct {
	BackupAmf string  `json:"backupAmf,omitempty"`
	GuamiList []Guami `json:"guamiList,omitempty"`
}

type SessInfo struct {
	Dnn                string         `json:"dnn,omitempty"`
	SNssai             *Snssai        `json:"snssai,omitempty"`
	ServiceName        string         `json:"serviceName,omitempty"`
	ServingNetwork     *PlmnId        `json:"servingNetwork"`
	AnType             string         `json:"anType"`
	RatType            string         `json:"ratType,omitempty"`
	RequestType        string         `json:"requestType,omitempty"`
	SelMode            string         `json:"selMode,omitempty"`
	ServingNfId        string         `json:"servingNfId"`
	Guami              Guami          `json:"guami,omitempty"`
	SmContextStatusUri string         `json:"smContextStatusUri"`
	BackupAmfInfo      *BackupAmfInfo `json:"backupAmfInfo,omitempty"`
}

type SessRoamingInfo struct {
	// for V-SMF to use
	HSmfInstanceId    string   `json:"hSmfInstanceId,omitempty"`
	HSmfUri           string   `json:"hSmfUri,omitempty"`
	AdditionalHsmfUri []string `json:"additionalHsmfUri,omitempty"`
	HplmnSnssai       Snssai   `json:"hplmnSnssai,omitempty"`
	// for H-SMF to use
	VSmfInstanceId          string     `json:"vSmfInstanceId,omitempty"`
	VSmfPduSessionUri       string     `json:"vSmfPduSessionUri,omitempty"`
	OldPduSessionId         int32      `json:"oldPduSessionId,omitempty"`
	PduSessionsActivateList []int32    `json:"pduSessionsActivateList,omitempty"`
	ServiceName             string     `json:"serviceName,omitempty"`
	ServingNetwork          *PlmnId    `json:"servingNetwork,omitempty"`
	AnType                  AccessType `json:"anType,omitempty"`
	RatType                 RatType    `json:"ratType,omitempty"`
}

type PreemptionCapability = string

const (
	PreemptionCapability_NOT_PREEMPT PreemptionCapability = "NOT_PREEMPT"
	PreemptionCapability_MAY_PREEMPT PreemptionCapability = "MAY_PREEMPT"
)

type MultipleQuotaInformation struct {
	ResultCode           string       `json:"resultCode,omitempty"`
	RatingGroup          int32        `json:"ratingGroup"`
	GrantedUnit          *GrantedUnit `json:"grantedUnit,omitempty"`
	Triggers             []SdTrigger  `json:"triggers,omitempty"`
	ValidityTime         string       `json:"validityTime,omitempty"`
	QuotaHoldingTime     *int32       `json:"quotaHoldingTime,omitempty"`
	TimeQuotaThreshold   *int32       `json:"timeQuotaThreshold,omitempty"`
	VolumeQuotaThreshold *int32       `json:"volumeQuotaThreshold,omitempty"`
	UnitQuotaThreshold   *int32       `json:"unitQuotaThreshold,omitempty"`
	UPFID                *string      `json:"uPFID,omitempty"`
}

type ChargingDataResponse struct {
	InvocationTimeStamp           time.Time                      `json:"invocationTimeStamp,omitempty"`
	InvocationSequenceNumber      int32                          `json:"invocationSequenceNumber"`
	InvocationResult              *InvocationResult              `json:"invocationResult,omitempty"`
	SessionFailover               *SessionFailover               `json:"sessionFailover,omitempty"`
	MultipleUnitInformation       []MultipleUnitInformation      `json:"multipleUnitInformation,omitempty"`
	Triggers                      []SdTrigger                    `json:"triggers,omitempty"`
	PDUSessionChargingInformation *PduSessionChargingInformation `json:"pDUSessionChargingInformation,omitempty"`
	RoamingQBCInformation         *RoamingQbcInformation         `json:"roamingQBCInformation,omitempty"`
}

type RoamingQbcInformation struct {
	MultipleQFIcontainer   []MultipleQfIcontainer    `json:"multipleQFIcontainer,omitempty"`
	UPFID                  string                    `json:"uPFID,omitempty"`
	RoamingChargingProfile *SdRoamingChargingProfile `json:"roamingChargingProfile,omitempty"`
}

type SdRoamingChargingProfile struct {
	Triggers            []SdTrigger          `json:"triggers,omitempty"`
	PartialRecordMethod *PartialRecordMethod `json:"partialRecordMethod,omitempty"`
}

type MultipleQfIcontainer struct {
	Triggers                []SdTrigger              `json:"triggers,omitempty"`
	TriggerTimestamp        *time.Time               `json:"triggerTimestamp,omitempty"`
	Time                    *int32                   `json:"time,omitempty"`
	TotalVolume             *int64                   `json:"totalVolume,omitempty"`
	UplinkVolume            *int64                   `json:"uplinkVolume,omitempty"`
	LocalSequenceNumber     int32                    `json:"localSequenceNumber"`
	QFIContainerInformation *QfiContainerInformation `json:"qFIContainerInformation,omitempty"`
}

type QfiContainerInformation struct {
	QFI                              *int32                       `json:"qFI,omitempty"`
	ReportTime                       *time.Time                   `json:"reportTime,omitempty"`
	TimeofFirstUsage                 *time.Time                   `json:"timeofFirstUsage,omitempty"`
	TimeofLastUsage                  *time.Time                   `json:"timeofLastUsage,omitempty"`
	QoSInformation                   *SdQosData                   `json:"qoSInformation,omitempty"`
	UserLocationInformation          *UserLocation                `json:"userLocationInformation,omitempty"`
	UetimeZone                       string                       `json:"uetimeZone,omitempty"`
	PresenceReportingAreaInformation map[string]PresenceInfo      `json:"presenceReportingAreaInformation,omitempty"`
	RATType                          *RatType                     `json:"rATType,omitempty"`
	ServingNetworkFunctionID         []SdServingNetworkFunctionId `json:"servingNetworkFunctionID,omitempty"`
	Var3gppPSDataOffStatus           *Model3GpppsDataOffStatus    `json:"3gppPSDataOffStatus,omitempty"`
}

type SdTrigger struct {
	TriggerType     string          `json:"triggerType"`
	TriggerCategory TriggerCategory `json:"triggerCategory"`
	TimeLimit       *int32          `json:"timeLimit,omitempty"`
	VolumeLimit     *int32          `json:"volumeLimit,omitempty"`
	VolumeLimit64   *int64          `json:"volumeLimit64,omitempty"`
	MaxNumberOfccc  *int32          `json:"maxNumberOfccc,omitempty"`
	//version:1.R15.0.0
	Category *TriggerCategory `json:"category,omitempty"`
}

type MultipleUnitInformation struct {
	ResultCode           *ResultCode          `json:"resultCode,omitempty"`
	RatingGroup          uint32               `json:"ratingGroup"`
	GrantedUnit          *GrantedUnit         `json:"grantedUnit,omitempty"`
	Triggers             []SdTrigger          `json:"triggers,omitempty"`
	ValidityTime         interface{}          `json:"validityTime,omitempty"`
	QuotaHoldingTime     *int32               `json:"quotaHoldingTime,omitempty"`
	FinalUnitIndication  *FinalUnitIndication `json:"finalUnitIndication,omitempty"`
	TimeQuotaThreshold   *int32               `json:"timeQuotaThreshold,omitempty"`
	VolumeQuotaThreshold *int64               `json:"volumeQuotaThreshold,omitempty"`
	UnitQuotaThreshold   *int32               `json:"unitQuotaThreshold,omitempty"`
	UPFID                string               `json:"uPFID,omitempty"`
}

type GrantedUnit struct {
	TariffTimeChange     *time.Time `json:"tariffTimeChange,omitempty"`
	Time                 *int32     `json:"time,omitempty"`
	TotalVolume          *int64     `json:"totalVolume,omitempty"`
	UplinkVolume         *int64     `json:"uplinkVolume,omitempty"`
	DownlinkVolume       *int64     `json:"downlinkVolume,omitempty"`
	ServiceSpecificUnits *int64     `json:"serviceSpecificUnits,omitempty"`
}

type ResultCode = string

type SessionFailover = string

type InvocationResult struct {
	Error           *ProblemDetails  `json:"error,omitempty"`
	FailureHandling *FailureHandling `json:"failureHandling,omitempty"`
}

type ProblemDetails struct {
	Type              string         `json:"type,omitempty"`
	Title             string         `json:"title,omitempty"`
	Status            *int32         `json:"status,omitempty"`
	Detail            string         `json:"detail,omitempty"`
	Instance          string         `json:"instance,omitempty"`
	Cause             string         `json:"cause,omitempty"`
	InvalidParams     []InvalidParam `json:"invalidParams,omitempty"`
	SupportedFeatures string         `json:"supportedFeatures,omitempty"`
}

type InvalidParam struct {
	Param  string `json:"param"`
	Reason string `json:"reason,omitempty"`
}

type NextUrrID struct {
	OnlineUrrIncID  uint32 `json:"onlineUrrIncID,omitempty"`
	OfflineUrrIncID uint32 `json:"offlineUrrIncID,omitempty"`
}

type UdmRscInfo struct {
	SupportedFeatures           string            `json:"supportedFeatures,omitempty"`
	PgwFqdn                     string            `json:"pgwFqdn,omitempty"`
	SmSubsData                  *DnnConfiguration `json:"smSubsData,omitempty"`
	NfInstanceId                string            `json:"nfInstanceId"`
	AmfServiceName              string            `json:"amfServiceName,omitempty"`
	MonitoredResourceUris       []string          `json:"monitoredResourceUris,omitempty"`
	PcscfRestorationCallbackUri string            `json:"pcscfRestorationCallbackUri,omitempty"`
	SubscriptionId              string            `json:"subscriptionId,omitempty"`
}

type Nas5gsmCap struct {
	Mh6Pdu uint8 `json:"mh6Pdu"`
	Rqos   uint8 `json:"rqos"`
}

type SessAmbr struct {
	UnitDown uint8  `json:"uintDown"`
	AmbrDown uint16 `json:"ambrDown"`
	UnitUp   uint8  `json:"uintUp"`
	AmbrUp   uint8  `json:"ambrUp"`
}

type NasMaxDataRate struct {
	Uplink   uint8 `json:"uplink"`
	Downlink uint8 `json:"downlink"`
}

type NasData struct {
	PduSessType     uint8          `json:"pduSessType"`
	SscMode         uint8          `json:"sscMode"`
	SmCap           Nas5gsmCap     `json:"smCap"`
	MaxPktFilter    uint16         `json:"maxPft"`
	Pco             []uint16       `json:"pco"`
	MaxDataRate     NasMaxDataRate `json:"maxDataRate"`
	Ambr            SessAmbr       `json:"ambr"`
	Cause           uint8          `json:"cause"`
	Apsr            uint8          `json:"apsr"`
	NasNextId       NasNextId      `json:"nasNextId"`
	PsDataOffStatus bool           `json:"psDataOffStatus"`
}

type NasNextId struct {
	Pti       uint8 `json:"pti"`
	QosRuleId uint8 `json:"qosRuleId"`
}

type stType uint8

type FsmState struct {
	Current string `json:"current,omitempty"`
	Last    string `json:"last,omitempty"`
}

type RgDataList []*RgData

type CtfTriggerList []*CtfTrigger

type UrrIDList []*UrrIDMap

type CacheUsedUnitContainer struct {
	RatingGroup uint32
	CacheUUC    *UsedUnitContainer
}

type CacheUsedUnitContainerList []*CacheUsedUnitContainer

type SessChargingData struct {
	ChargingID                *uint32                        `json:"chargingID,omitempty"`
	HChargingID               *uint32                        `json:"hChargingID,omitempty"`
	IsReqFail                 bool                           `json:"isReqFail,omitempty"`
	IsQbc                     bool                           `json:"isQbc,omitempty"`
	PraInfo                   map[string]PresenceInfo        `json:"praInfo,omitempty"`
	FailureHandling           *FailureHandling               `json:"failureHandling,omitempty"`
	RoamingChargingProfile    *RoamingChargingProfile        `json:"roamingChargingProfile,omitempty"`
	RgTriggers                CtfTriggerList                 `json:"rgTriggers,omitempty"`
	SessTriggers              *CtfTrigger                    `json:"sesstriggers,omitempty"`
	QosFlowTriggers           *CtfTrigger                    `json:"qosFlowTriggers,omitempty"`
	CurCC                     uint32                         `json:"curCC"`
	RgMap                     RgDataList                     `json:"rgMap,omitempty"`
	DefaultRG                 uint32                         `json:"defaultRG"`
	QfiMap                    RgDataList                     `json:"qfiMap,omitempty"`
	SessLevelUrrID            *uint32                        `json:"sessLevelUrrID,omitempty"`
	UrrIDList                 UrrIDList                      `json:"urrIDList,omitempty"`
	RemovedUrrIDList          UrrIDList                      `json:"removedUrrIDList,omitempty"`
	NextID                    NextUrrID                      `json:"nextID"`
	LastReportTime            *time.Time                     `json:"lastReportTime,omitempty"`
	UnitUsageList             UnitUsageList                  `json:"unitUsageList,omitempty"`
	CacheUsedUnitContainer    CacheUsedUnitContainerList     `json:"cacheUsedUnitContainer"`
	MultipleQFIcontainer      []MultipleQfIcontainer         `json:"multipleQFIcontainer,omitempty"`
	UpdateData                *PduSessionChargingInformation `json:"updateData,omitempty"`
	UpfInstanceId             string                         `json:"upfInstanceId"`
	PccRuleBaseName           string                         `json:"pccRuleBaseName"`
	RatType                   string                         `json:"ratType"`
	InvocationSequenceNumber  uint32                         `json:"invocationSequenceNumber"`
	FairChargingDataReq       []*FailChargingReq             `json:"fairChargingDataReq,omitempty"`
	AddtionalPsaUpf           *UpfChargingData               `json:"addtionalPsaUpf,omitempty"`
	ApActivatedTrctn          ApNchfTrascations              `json:"apActivatedTrctn"`
	VzChargingOptions         *VzChargingOptions             `json:"vzChargingOptions,omitempty"`
	CdrLocalSeqNum            uint32                         `json:"cdrLocalSeqNum"`
	CcrCcRequestNumberPlusOne uint32                         `json:"ccrCcRequestNumberPlusOne"`
	OcsSessOn                 bool                           `json:"ocsSessOn"`
	LastUpdateChargingTime    *time.Time                     `json:"lastUpdateChargingTime,omitempty" cmp:"ignore"`
}

type UpfChargingData struct {
	UpfInstanceId string `json:"upfInstanceId"`
	//upf specific charging data, i.e. triggers, quotas etc.
	RgTriggers CtfTriggerList `json:"rgTriggers"` //RgTriggers under Charging data for PSA and this one is supposed to be mutually exclusive.
}

type VzChargingOptions struct {
	OfflineSessionTimePeriod *int       `json:"offlineSessionTimePeriod,omitempty"`
	BcrReportingTime         *time.Time `json:"bcrReportingTime,omitempty"`
	OfflineTariffTimeChange  *time.Time `json:"offlineTariffTimeChange,omitempty"`
}

type ApNchfTrascations string

const (
	NCHF_CREATE  = "Nchf_create"
	NCHF_UPDATE  = "Nchf_update"
	NCHF_RELEASE = "Nchf_release"
)

type ChargingReqType uint8

const (
	CHARGING_REQ_TYPE_CREATE  ChargingReqType = 1
	CHARGING_REQ_TYPE_UPDATE  ChargingReqType = 2
	CHARGING_REQ_TYPE_RELEASE ChargingReqType = 3
)

type FailChargingReq struct {
	// Req  *nchf.ChargingDataRequest `json:"req,omitempty"`
	Type ChargingReqType `json:"type"`
}

type UnitUsageList struct {
	Muus          []MultipleUnitUsage `json:"muus,omitempty"`
	ReportingKeys []ReportingKey      `json:"reportingKeys,omitempty"`
}

type ReportingKey struct {
	Level        ReportingLevel `json:"level"`
	RatingGroup  uint32         `json:"ratingGroup"`
	ServiceId    int32          `json:"serviceId"`
	SponsorId    string         `json:"sponsorId"`
	AppSvcProvId string         `json:"appSvcProvId"`
}

type OdbPacketServices = string

const (
	OdbPacketServices_ALL_PACKET_SERVICES    OdbPacketServices = "ALL_PACKET_SERVICES"
	OdbPacketServices_ROAMER_ACCESS_HPLMN_AP OdbPacketServices = "ROAMER_ACCESS_HPLMN_AP"
	OdbPacketServices_ROAMER_ACCESS_VPLMN_AP OdbPacketServices = "ROAMER_ACCESS_VPLMN_AP"
)

// FrameRouteInfo struct for FrameRouteInfo
type FrameRouteInfo struct {
	Ipv4Mask   string `json:"ipv4Mask,omitempty"`
	Ipv6Prefix string `json:"ipv6Prefix,omitempty"`
}

type SessDnnConfiguration struct {
	//Default/Allowed session types
	PduSessionTypes PduSessionTypes `mapstructure:"pduSessionTypes" json:"pduSessionTypes"`
	//Default/Allowed SSC modes
	SscModes SscModes `mapstructure:"sscModes" json:"sscModes"`
	//Indicates whether interworking with EPS is subscribed:
	// true: Subscribed;
	// false: Not subscribed;
	// If this attribute is absent it means not subscribed.
	IwkEpsInd *bool `mapstructure:"iwkEpsInd" json:"iwkEpsInd,omitempty"`
	//5G QoS parameters associated to the session for a data network
	Var5gQosProfile *SubscribedDefaultQos `mapstructure:"var5gQosProfile" json:"5gQosProfile,omitempty"`
	//The maximum aggregated uplink and downlink bit rates to be shared across all Non-GBR QoS Flows in each PDU Session
	SessionAmbr *Ambr `mapstructure:"sessionAmbr" json:"sessionAmbr,omitempty"`
	//Subscribed charging characteristics data associated to the session for a data network.
	Var3gppChargingCharacteristics string `mapstructure:"var3gppChargingCharacteristics" json:"3gppChargingCharacteristics,omitempty"`
	//Subscribed static IP address(es) of the IPv4 and/or IPv6 type
	StaticIpAddress []IpAddress `mapstructure:"staticIpAddress" json:"staticIpAddress,omitempty"`
	//When present, this IE shall indicate the security policy for integrity protection and encryption for the user plane.
	UpSecurity         *UpSecurity      `mapstructure:"upSecurity" json:"upSecurity,omitempty"`
	Ipv4FrameRouteList []FrameRouteInfo `json:"ipv4FrameRouteList,omitempty"`
	Ipv6FrameRouteList []FrameRouteInfo `json:"ipv6FrameRouteList,omitempty"`
}

type SessionManagementSubscriptionData struct {
	SingleNssai Snssai `json:"singleNssai"`
	// A map (list of key-value pairs where Dnn serves as key) of DnnConfigurations
	DnnConfigurations         map[string]SessDnnConfiguration `json:"dnnConfigurations,omitempty"`
	InternalGroupIds          []string                        `json:"internalGroupIds,omitempty"`
	SharedDnnConfigurationsId string                          `json:"sharedDnnConfigurationsIds,omitempty"`
	OdbPacketServices         *OdbPacketServices              `json:"odbPacketServices,omitempty"`
}

type SessUdmInfo struct {
	//SmfRegistration
	SmfRegistration   bool   `json:"smfRegistration,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
	PgwFqdn           string `json:"pgwFqdn,omitempty"`
	//SessionManagementSubscriptionData
	SmSubsData *SessionManagementSubscriptionData `json:"smSubsData,omitempty"`
	//SdmSubscription
	NfInstanceId                string       `json:"nfInstanceId,omitempty"`
	ImplicitUnsubscribe         bool         `json:"implicitUnsubscribe,omitempty"`
	Expires                     *time.Time   `json:"expires,omitempty"`
	AmfServiceName              *ServiceName `json:"amfServiceName,omitempty"`
	MonitoredResourceUris       []string     `json:"monitoredResourceUris,omitempty"`
	PcscfRestorationCallbackUri string       `json:"pcscfRestorationCallbackUri,omitempty"`
	SubscriptionId              string       `json:"subscriptionId,omitempty"`
}

type TimerInfo struct {
	TimerId   uint64
	TimerType uint8
}

type TimerInfoList struct {
	List []*TimerInfo
}
type TimerCache struct {
	NasMsg []byte
}

type SessionData struct {
	Key           *SessionKey       `json:"key,omitempty"`
	Nas           *NasData          `json:"nas,omitempty"`
	Rsc           *SessionRsc       `json:"rsc,omitempty"`
	State         *FsmState         `json:"state,omitempty"`
	Loc           *UeLocation       `json:"loc,omitempty"`
	AmfInfo       *SessAmfInfo      `json:"amfInfo,omitempty"`
	NwInfo        *SessNwInfo       `json:"nwInfo,omitempty"`
	UdmInfo       *SessUdmInfo      `json:"udmInfo,omitempty"`
	Roaming       *SessRoamingInfo  `json:"roaming,omitempty"`
	HoInfo        *SessHoInfo       `json:"hoInfo,omitempty"`
	EpcInfo       *SessEpcInfo      `json:"epcInfo,omitempty"`
	EventInfo     *EventInfo        `json:"eventInfo,omitempty"`
	TraceData     *TraceData        `json:"traceData,omitempty"`
	NrfUri        *NrfInfo          `json:"nrfUri,omitempty"`
	Flows         *SessFlow         `json:"flows,omitempty"`
	Pd            *SmPolicyDecision `json:"pd,omitempty"`
	PendingPccReq *PendingPccReq    `json:"pendingPccReq,omitempty"`
	ChargingData  *SessChargingData `json:"chargingData,omitempty"`
	AuthData      *AuthData         `json:"authData,omitempty"`
	TimerCache    *TimerCache       `json:"timerCache,omitempty"`
	TimerInfoList *TimerInfoList    `json:"timerInfoList,omitempty"`
}

// SMF detected trigger bit masks
type DetectedTrigger uint32

const (
	DETECT_TRIGGER_PLMN_CH        DetectedTrigger = 0x1
	DETECT_TRIGGER_RES_MO_RE      DetectedTrigger = 0x2
	DETECT_TRIGGER_AC_TY_CH       DetectedTrigger = 0x4
	DETECT_TRIGGER_UE_IP_CH       DetectedTrigger = 0x8
	DETECT_TRIGGER_UE_MAC_CH      DetectedTrigger = 0x10
	DETECT_TRIGGER_AN_CH_COR      DetectedTrigger = 0x20
	DETECT_TRIGGER_US_RE          DetectedTrigger = 0x40
	DETECT_TRIGGER_APP_STA        DetectedTrigger = 0x80
	DETECT_TRIGGER_APP_STO        DetectedTrigger = 0x100
	DETECT_TRIGGER_AN_INFO        DetectedTrigger = 0x200
	DETECT_TRIGGER_CM_SES_FAIL    DetectedTrigger = 0x400
	DETECT_TRIGGER_PS_DA_OFF      DetectedTrigger = 0x800
	DETECT_TRIGGER_DEF_QOS_CH     DetectedTrigger = 0x1000
	DETECT_TRIGGER_SE_AMBR_CH     DetectedTrigger = 0x2000
	DETECT_TRIGGER_QOS_NOTIF      DetectedTrigger = 0x4000
	DETECT_TRIGGER_NO_CREDIT      DetectedTrigger = 0x8000
	DETECT_TRIGGER_PRA_CH         DetectedTrigger = 0x10000
	DETECT_TRIGGER_SAREA_CH       DetectedTrigger = 0x20000
	DETECT_TRIGGER_SCNN_CH        DetectedTrigger = 0x40000
	DETECT_TRIGGER_RE_TIMEOUT     DetectedTrigger = 0x80000
	DETECT_TRIGGER_RES_RELEASE    DetectedTrigger = 0x100000
	DETECT_TRIGGER_SUCC_RES_ALLO  DetectedTrigger = 0x200000
	DETECT_TRIGGER_RAT_TY_CH      DetectedTrigger = 0x400000
	DETECT_TRIGGER_REF_QOS_IND_CH DetectedTrigger = 0x800000
	//TODO: add more triggers that are necessary
)

// Possible values are - ACTIVE: Indicates that the PCC rule(s) are successfully installed (for those provisioned from PCF) or activated (for those pre-defined in SMF), or the session rule(s) are successfully installed  - INACTIVE: Indicates that the PCC rule(s) are removed (for those provisioned from PCF) or inactive (for those pre-defined in SMF) or the session rule(s) are removed.
type RuleStatus = string

const (
	RuleStatus_ACTIVE   RuleStatus = "ACTIVE"
	RuleStatus_INACTIVE RuleStatus = "INACTIVE"
)

type FailureCode = string

const (
	FailureCode_UNK_RULE_ID            FailureCode = "UNK_RULE_ID"
	FailureCode_RA_GR_ERR              FailureCode = "RA_GR_ERR"
	FailureCode_SER_ID_ERR             FailureCode = "SER_ID_ERR"
	FailureCode_NF_MAL                 FailureCode = "NF_MAL"
	FailureCode_RES_LIM                FailureCode = "RES_LIM"
	FailureCode_MAX_NR_QoS_FLOW        FailureCode = "MAX_NR_QoS_FLOW"
	FailureCode_MISS_FLOW_INFO         FailureCode = "MISS_FLOW_INFO"
	FailureCode_RES_ALLO_FAIL          FailureCode = "RES_ALLO_FAIL"
	FailureCode_UNSUCC_QOS_VAL         FailureCode = "UNSUCC_QOS_VAL"
	FailureCode_INCOR_FLOW_INFO        FailureCode = "INCOR_FLOW_INFO"
	FailureCode_PS_TO_CS_HAN           FailureCode = "PS_TO_CS_HAN"
	FailureCode_APP_ID_ERR             FailureCode = "APP_ID_ERR"
	FailureCode_NO_QOS_FLOW_BOUND      FailureCode = "NO_QOS_FLOW_BOUND"
	FailureCode_FILTER_RES             FailureCode = "FILTER_RES"
	FailureCode_MISS_REDI_SER_ADDR     FailureCode = "MISS_REDI_SER_ADDR"
	FailureCode_CM_END_USER_SER_DENIED FailureCode = "CM_END_USER_SER_DENIED"
	FailureCode_CM_CREDIT_CON_NOT_APP  FailureCode = "CM_CREDIT_CON_NOT_APP"
	FailureCode_CM_AUTH_REJ            FailureCode = "CM_AUTH_REJ"
	FailureCode_CM_USER_UNK            FailureCode = "CM_USER_UNK"
	FailureCode_CM_RAT_FAILED          FailureCode = "CM_RAT_FAILED"
	FailureCode_UE_STA_SUSP            FailureCode = "UE_STA_SUSP"
	//FailureCode_SESS_AMBR_FAILURE       FailureCode = "SESS_AMBR_FAILURE"  //!!!!This structure/attribute was abandoned in 2019-09 version upgrade, DO NOT USE it any more!!!!
	//FailureCode_DEF_QOS_FAILURL         FailureCode = "DEF_QOS_FAILURL"  //!!!!This structure/attribute was abandoned in 2019-09 version upgrade, DO NOT USE it any more!!!!
)

type NgApCause struct {
	Group int32 `json:"group"`
	Value int32 `json:"value"`
}

type RanNasRelCause struct {
	NgApCause    *NgApCause `json:"ngApCause,omitempty"`
	Var5gMmCause *int32     `json:"5gMmCause,omitempty"`
	Var5gSmCause *int32     `json:"5gSmCause,omitempty"`
}

type RuleReport struct {
	// Contains the identifier of the affected PCC rule(s).
	PccRuleIds []string   `json:"pccRuleIds"`
	RuleStatus RuleStatus `json:"ruleStatus"`
	// Indicates the version of a PCC rule.
	ContVers    []int32          `json:"contVers,omitempty"`
	FailureCode *FailureCode     `json:"failureCode"`
	FinUnitAct  *FinalUnitAction `json:"finUnitAct,omitempty"`
	// indicates the RAN or NAS release cause code information.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty"`
}

// SessionRuleFailureCode Possible values are   - NF_MAL: Indicate that the PCC rule could not be successfully installed (for those provisioned from the PCF) or activated (for those pre-defined in SMF) or enforced (for those already successfully installed) due to SMF/UPF malfunction.   - RES_LIM: Indicate that the PCC rule could not be successfully installed (for those provisioned from PCF) or activated (for those pre-defined in SMF) or enforced (for those already successfully installed) due to a limitation of resources at the SMF/UPF.   - UNSUCC_QOS_VAL: indicate that the QoS validation has failed.   - UE_STA_SUSP: Indicates that the UE is in suspend state.
type SessionRuleFailureCode = string

const (
	SessionRuleFailureCode_NF_MAL         SessionRuleFailureCode = "NF_MAL"
	SessionRuleFailureCode_RES_LIM        SessionRuleFailureCode = "RES_LIM"
	SessionRuleFailureCode_UNSUCC_QOS_VAL SessionRuleFailureCode = "UNSUCC_QOS_VAL"
	SessionRuleFailureCode_UE_STA_SUSP    SessionRuleFailureCode = "UE_STA_SUSP"
)

// SessionRuleReport struct for SessionRuleReport
type SessionRuleReport struct {
	// Contains the identifier of the affected session rule(s).
	RuleIds             []string                `json:"ruleIds"`
	RuleStatus          RuleStatus              `json:"ruleStatus"`
	SessRuleFailureCode *SessionRuleFailureCode `json:"sessRuleFailureCode,omitempty"`
}

type NrfInfo struct {
	Uri string
}

type TraceData struct {
	TraceRef                 string     `json:"traceRef"`
	TraceDepth               TraceDepth `json:"traceDepth"`
	NeTypeList               string     `json:"neTypeList"`
	EventList                string     `json:"eventList"`
	CollectionEntityIpv4Addr string     `json:"collectionEntityIpv4Addr,omitempty"`
	CollectionEntityIpv6Addr string     `json:"collectionEntityIpv6Addr,omitempty"`
	InterfaceList            string     `json:"interfaceList,omitempty"`
}

type TraceDepth = string

const (
	TraceDepth_MINIMUM                     TraceDepth = "MINIMUM"
	TraceDepth_MEDIUM                      TraceDepth = "MEDIUM"
	TraceDepth_MAXIMUM                     TraceDepth = "MAXIMUM"
	TraceDepth_MINIMUM_WO_VENDOR_EXTENSION TraceDepth = "MINIMUM_WO_VENDOR_EXTENSION"
	TraceDepth_MEDIUM_WO_VENDOR_EXTENSION  TraceDepth = "MEDIUM_WO_VENDOR_EXTENSION"
	TraceDepth_MAXIMUM_WO_VENDOR_EXTENSION TraceDepth = "MAXIMUM_WO_VENDOR_EXTENSION"
)

type BearerBindingInfo struct {
	Qci uint32 `json:"qci,omitempty"`
	Pl  uint32 `json:"pl,omitempty"`
	Pvi uint32 `json:"pvi,omitempty"`
	Pci uint32 `json:"pci,omitempty"`
}

type BearerInfo struct {
	IsDefault     bool               `json:"is_default,omitempty"`
	Ebi           uint32             `json:"ebi,omitempty"`
	BearerBinding *BearerBindingInfo `json:"bearer_binding,omitempty"`
	//Qos of bearer should be the sum of qos of binding pcc rules
	MbrUl uint64 `json:"mbr_ul,omitempty"`
	MbrDl uint64 `json:"mbr_dl,omitempty"`
	GbrUl uint64 `json:"gbr_ul,omitempty"`
	GbrDl uint64 `json:"gbr_dl,omitempty"`
	Qfi   uint32 `json:"qfi,omitempty"`
}

type epc5gHoState uint8

const (
	Session5G      epc5gHoState = 0
	SessionFromEpc epc5gHoState = 1
	SessionToEpc   epc5gHoState = 2
)

type PfcpBindingInfo struct {
	PdrIdUl       uint32   `json:"pdr_id_ul,omitempty"`
	PdrIdDl       uint32   `json:"pdr_id_dl,omitempty"`
	PdrPriUl      uint32   `json:"pdr_pri_ul,omitempty"`
	PdrPriDl      uint32   `json:"pdr_pri_dl,omitempty"`
	FarIdUl       uint32   `json:"far_id_ul,omitempty"`
	FarIdDl       uint32   `json:"far_id_dl,omitempty"`
	QerIdUl       uint32   `json:"qer_id_ul,omitempty"`
	QerIdDl       uint32   `json:"qer_id_dl,omitempty"`
	UrrId         []uint32 `json:"urr_id,omitempty"`
	Qfi           uint32   `json:"qfi,omitempty"`
	Ebi           uint32   `json:"ebi,omitempty"`
	PccRuleId     string   `json:"pcc_rule_id,omitempty"`
	IsDefaultRule bool     `json:"is_default_rule,omitempty"`
}

type FteidInfo struct {
	Teid     uint32 `json:"teid,omitempty"`
	V4       bool   `json:"v4,omitempty"`
	V6       bool   `json:"v6,omitempty"`
	Ipv4Addr uint32 `json:"ipv4_addr,omitempty"`
	Ipv6Addr []byte `json:"ipv6_addr,omitempty"`
}

type EpsBearerSetupInfo struct {
	//upf_teid will be same for all ebi in 5G-4G HO.
	UpfTeid *FteidInfo `json:"upf_teid,omitempty"`
	//DL Data forwarding, NGRan->upf
	Ebi uint32 `json:"ebi,omitempty"`
}

type IndirectTunnelInfo struct {
	Ebi            uint8              `json:"ebi,omitempty"`
	PdrIndirect    uint32             `json:"pdrIndirect,omitempty"`
	FarIndirect    uint32             `json:"farIndirect,omitempty"`
	QerIndirect    uint32             `json:"qerIndirect,omitempty"`
	EpsBearerSetup EpsBearerSetupInfo `json:"epsBearerSetup,omitempty"`
}

type SessEpcInfo struct {
	UeEpsPdnConnection string                `json:"ueEpsPdnConnection,omitempty"`
	MappedBearers      []*BearerInfo         `json:"mappedBearers,omitempty"`
	PgwTeid            uint32                `json:"pgwTeid,omitempty"`
	ApnRestriction     uint8                 `json:"apnRestriction,omitempty"`
	Is4GSess           bool                  `json:"is4GSess,omitempty"`
	HoState            epc5gHoState          `json:"hoState,omitempty"` //when ho to 5G complete,set it to Session5G
	OldPfcpInfo        []*PfcpBindingInfo    `json:"oldPfcpInfo,omitempty"`
	BearerContext      [][]byte              `json:"bearerContext,omitempty"`   //can use gtpc project to decode it
	IndirectTunnels    []*IndirectTunnelInfo `json:"indirectTunnels,omitempty"` //For 4G-5G HO. One tunnel per EBI
	LocalTeid          uint32                `json:"localTeid,omitempty"`       // local pgwc gtpc teid
	PeerTeid           *Fteid                `json:"peerTeid,omitempty"`        // peer sgwc/epdg gtpc teid
}

type Fteid struct {
	V4           bool           `json:"v4,omitempty"`
	V6           bool           `json:"v6,omitempty"`
	IntfType     Fteid_IntfType `json:"intfType,omitempty"`
	TeidOrGreKey uint32         `json:"teidOrGreKey,omitempty"`
	Ipv4Addr     string         `json:"ipv4Addr,omitempty"`
	Ipv6Addr     string         `json:"ipv6Addr,omitempty"`
}

type Fteid_IntfType int32

const (
	Fteid_S1U_eNodeB_GTPU Fteid_IntfType = 0
	Fteid_S1U_SGW_GTPU    Fteid_IntfType = 1
	Fteid_S12_RNC_GTPU    Fteid_IntfType = 2
	Fteid_S12_SGW_GTPU    Fteid_IntfType = 3
	Fteid_SGW_GTPU        Fteid_IntfType = 4
	Fteid_S5S8_PGW_GTPU   Fteid_IntfType = 5
	Fteid_S5S8_SGW_GTPC   Fteid_IntfType = 6
	Fteid_S5S8_PGW_GTPC   Fteid_IntfType = 7
	Fteid_S5S8_SGW_PMIPv6 Fteid_IntfType = 8
	Fteid_S5S8_PGW_PMIPv6 Fteid_IntfType = 9
	Fteid_S11_MME_GTPC    Fteid_IntfType = 10
	Fteid_S11S4_SGW_GTPC  Fteid_IntfType = 11
	Fteid_S10_MME_GTPC    Fteid_IntfType = 12
	Fteid_S3_MME_GTPC     Fteid_IntfType = 13
	Fteid_S3_SGSN_GTPC    Fteid_IntfType = 14
	Fteid_S4_SGSN_GTPU    Fteid_IntfType = 15
	Fteid_S4_SGW_GTPU     Fteid_IntfType = 16
	Fteid_S4_SGSN_GTPC    Fteid_IntfType = 17
	Fteid_S16_SGSN_GTPC   Fteid_IntfType = 18
	Fteid_eNodeB_GTPU_DL  Fteid_IntfType = 19
	Fteid_eNodeB_GTPU_UL  Fteid_IntfType = 20
	Fteid_RNC_GTPU        Fteid_IntfType = 21
	Fteid_SGSN_GTPU       Fteid_IntfType = 22
	Fteid_SGWUPF_GTPU     Fteid_IntfType = 23
	Fteid_SmMBMS_GW_GTPC  Fteid_IntfType = 24
	Fteid_SnMBMS_GW_GTPC  Fteid_IntfType = 25
	Fteid_Sm_MME_GTPC     Fteid_IntfType = 26
	Fteid_Sn_SGSN_GTPC    Fteid_IntfType = 27
	Fteid_SGW_GTPU_UL     Fteid_IntfType = 28
	Fteid_Sn_SGSN_GTPU    Fteid_IntfType = 29
	Fteid_S2b_ePDG_GTPC   Fteid_IntfType = 30
	Fteid_S2b_ePDG_GTPU   Fteid_IntfType = 31
	Fteid_S2b_PGW_GTPC    Fteid_IntfType = 32
	Fteid_S2b_PGW_GTPU    Fteid_IntfType = 33
	Fteid_S2a_TWAN_GTPU   Fteid_IntfType = 34
	Fteid_S2a_TWAN_GTPC   Fteid_IntfType = 35
	Fteid_S2a_PGW_GTPC    Fteid_IntfType = 36
	Fteid_S2a_PGW_GTPU    Fteid_IntfType = 37
	Fteid_S11_MME_GTPU    Fteid_IntfType = 38
	Fteid_S11_SGW_GTPU    Fteid_IntfType = 39
)

type SessHoInfo struct {
	HoState                 string         `json:"hoState,omitempty"`
	HSmfUri                 string         `json:"hSmfUri,omitempty"`
	AdditionalHsmfUri       []string       `json:"additionalHsmfUri,omitempty"`
	OldPduSessionId         int32          `json:"oldPduSessionId,omitempty"`
	PduSessionsActivateList []int32        `json:"pduSessionsActivateList,omitempty"`
	HplmnSnssai             *Snssai        `json:"hplmnSnssai,omitempty"`
	TargetServingNfId       string         `json:"targetServingNfId,omitempty"`
	DirectPathAvailability  bool           `json:"directPathAvailability,omitempty"`
	TargetRANFteid          *PfcpIeFteid   `json:"targetRANFteid,omitempty"`
	TargetRANFwdFteid       *PfcpIeFteid   `json:"targetRANFwdFteid,omitempty"`
	TargetRANId             *NgRanTargetId `json:"targetRANId,omitempty"`
}

type NgRanTargetId struct {
	RanNodeId GlobalRanNodeId `json:"ranNodeId"`
	Tai       Tai             `json:"tai"`
}

type Time struct {
	wall uint64
	ext  int64
	loc  *Location
}

type Location struct {
	name       string
	zone       []zone
	tx         []zoneTrans
	cacheStart int64
	cacheEnd   int64
	cacheZone  *zone
}

type zoneTrans struct {
	when         int64 // transition time, in seconds since 1970 GMT
	index        uint8 // the index of the zone that goes into effect at that time
	isstd, isutc bool  // ignored - no idea what these mean
}

type zone struct {
	name   string // abbreviated name, "CET"
	offset int    // seconds east of UTC
	isDST  bool   // is this zone Daylight Savings Time?
}

type SessAmfInfo struct {
	ServingNfId        string          `json:"servingNfId,omitempty"`
	Guami              *Guami          `json:"guami,omitempty"`
	SmContextStatusUri string          `json:"smContextStatusUri,omitempty"`
	BackupAmfInfo      []BackupAmfInfo `json:"backupAmfInfo,omitempty"`
	SubscriptionUri    string          `json:"subscriptionUri,omitempty"`
	PraEventIndex      *int            `json:"praEventIndex,omitempty"`
}

type SessNwInfo struct {
	Dnn            string     `json:"dnn,omitempty"`
	SNssai         Snssai     `json:"sNssai,omitempty"`
	ServiceName    string     `json:"serviceName,omitempty"`
	ServingNetwork PlmnId     `json:"servingNetwork,omitempty"`
	AnType         AccessType `json:"anType,omitempty"`
	PduSessionType uint8      `json:"pduSessionType,omitempty"`
	RatType        RatType    `json:"ratType,omitempty"`
	RequestType    string     `json:"requestType,omitempty"`
	SelMode        string     `json:"selMode,omitempty"`
	RequestedDnn   string     `json:"requestedDnn,omitempty"`
}

type GetSessionRsp struct {
	UeInfo   UeInfo        `json:"ueInfo,omitempty"`
	Sessions []SessionData `json:"sessions,omitempty"`
}

type GetSessionListRsp struct {
	SessionKeys []string `json:"sessionKeys,omitempty"`
}

type FinalUnitAction = string

const (
	FinalUnitAction_TERMINATE       FinalUnitAction = "TERMINATE"
	FinalUnitAction_REDIRECT        FinalUnitAction = "REDIRECT"
	FinalUnitAction_RESTRICT_ACCESS FinalUnitAction = "RESTRICT_ACCESS"
)

type RedirectAddressType string

const (
	RedirectAddressType_IPV4_1 RedirectAddressType = "IPV4"
	RedirectAddressType_IPV6_1 RedirectAddressType = "IPV6"
	RedirectAddressType_URL    RedirectAddressType = "URL"
)

// RedirectServer struct for RedirectServer
type RedirectServer struct {
	RedirectAddressType   RedirectAddressType `json:"redirectAddressType"`
	RedirectServerAddress string              `json:"redirectServerAddress"`
}

type FinalUnitIndication struct {
	FinalUnitAction       FinalUnitAction `json:"finalUnitAction"`
	RestrictionFilterRule string          `json:"restrictionFilterRule,omitempty"`
	FilterId              string          `json:"filterId,omitempty"`
	RedirectServer        *RedirectServer `json:"redirectServer,omitempty"`
}

type GsuType uint32

const (
	GSU_TOTAL_VOLUME  GsuType = 0x01
	GSU_UPLINK_VOLUME GsuType = 0x02
	GSU_DLLINK_VOLUME GsuType = 0x04
	GSU_TIME          GsuType = 0x08
)

type GSU struct {
	Flags               GsuType `json:"flags"`
	UplinkVol           *uint64 `json:"uplinkVol,omitempty"`
	DnlinkVol           *uint64 `json:"dnlinkVol,omitempty"`
	TotalVol            *uint64 `json:"totalVol,omitempty"`
	Time                *uint32 `json:"time,omitempty"`
	ServiceSpecificUnit *uint64 `json:"serviceSpecificUnit,omitempty"`
}

type RgData struct {
	RatingGroup uint32          `json:"ratingGroup"`
	Qfi         uint32          `json:"Qfi"`
	Online      bool            `json:"online"`
	Offline     bool            `json:"offline"`
	LastSeqNum  int32           `json:"lastSeqNum"`
	ChgDesc     *SdChargingData `json:"chgDesc"`
	QosData     *SdQosData      `json:"qosData"`
	// comments from wuyubin:
	// service id is omitempty, I using point to judgement whether exist or not.
	ServiceID *uint32 `json:"serviceID,omitempty"`
}

type CtfTriggerType uint32

const (
	CHF_TRIGGER_QUOTA_THRESHOLD         CtfTriggerType = 0x1
	CHF_TRIGGER_QHT                     CtfTriggerType = 0x2
	CHF_TRIGGER_FINAL                   CtfTriggerType = 0x4
	CHF_TRIGGER_QUOTA_EXHAUSTED         CtfTriggerType = 0x8
	CHF_TRIGGER_VALIDITY_TIME           CtfTriggerType = 0x10
	CHF_TRIGGER_OTHER_QUOTA_TYPE        CtfTriggerType = 0x20
	CHF_TRIGGER_FORCED_REAUTHORISATION  CtfTriggerType = 0x40
	CHF_TRIGGER_UNUSED_QUOTA_TIMER      CtfTriggerType = 0x80
	CHF_TRIGGER_ABNORMAL_RELEASE        CtfTriggerType = 0x100
	CHF_TRIGGER_QOS_CHANGE              CtfTriggerType = 0x200
	CHF_TRIGGER_VOLUME_LIMIT            CtfTriggerType = 0x400
	CHF_TRIGGER_TIME_LIMIT              CtfTriggerType = 0x800
	CHF_TRIGGER_PLMN_CHANGE             CtfTriggerType = 0x1000
	CHF_TRIGGER_USER_LOCATION_CHANGE    CtfTriggerType = 0x2000
	CHF_TRIGGER_RAT_CHANGE              CtfTriggerType = 0x4000
	CHF_TRIGGER_UE_TIMEZONE_CHANGE      CtfTriggerType = 0x8000
	CHF_TRIGGER_TARIFF_TIME_CHANGE      CtfTriggerType = 0x10000
	CHF_TRIGGER_MAX_OF_CC               CtfTriggerType = 0x20000
	CHF_TRIGGER_MANAGEMENT_INTERVENTION CtfTriggerType = 0x40000
	CHF_TRIGGER_UE_REPORT_AREA_CH       CtfTriggerType = 0x80000
	CHF_TRIGGER_3GPP_PS_DATA_CH         CtfTriggerType = 0x100000
	CHF_TRIGGER_SERVING_NODE_CHANGE     CtfTriggerType = 0x200000
	CHF_TRIGGER_REMOVAL_OF_UPF          CtfTriggerType = 0x400000
	CHF_TRIGGER_ADDITION_OF_UPF         CtfTriggerType = 0x800000
)

type CtfTrigger struct {
	RatingGroup uint32         `mapstructure:"ratingGroup" json:"ratingGroup"`
	Triggers    CtfTriggerType `mapstructure:"triggers" json:"triggers"`
	//UnSetTrg       CtfTriggerType       `mapstructure:"unSetTrg" json:"unSetTrg"` // This flag using to handle online charging case: there is gsu but none releated trigger inside
	Category       CtfTriggerType       `mapstructure:"category" json:"category"` // if bits is 1: means category is immediately; if bits is 0: means category is deferred.
	TimeLimit      *uint32              `mapstructure:"timeLimit" json:"timeLimit,omitempty"`
	VolumeLimit    *uint64              `mapstructure:"volumeLimit" json:"volumeLimit,omitempty"`
	MaxNumberOfccc *uint32              `mapstructure:"naxNumberOfccc" json:"maxNumberOfccc,omitempty"`
	Gsu            *GSU                 `mapstructure:"gsu" json:"gsu,omitempty"`
	GsuThrld       *GSU                 `mapstructure:"gsuThrld" json:"gsuThrld,omitempty"`
	Qht            *uint32              `mapstructure:"qht" json:"qht,omitempty"`
	Fui            *FinalUnitIndication `mapstructure:"fui" json:"fui,omitempty"`
	ActFarId       *uint32              `mapstructure:"actFarId" json:"actFarId,omitempty"`
	VTRunning      bool                 `mapstructure:"vTRunning" json:"vTRunning"`
	VTVal          *int64               `mapstructure:"vTVal" json:"vTVal,omitempty"`
}

type FramedRouting uint32

const (
	FramedRouting_Value_None            FramedRouting = 0
	FramedRouting_Value_Broadcast       FramedRouting = 1
	FramedRouting_Value_Listen          FramedRouting = 2
	FramedRouting_Value_BroadcastListen FramedRouting = 3
)

type IP []byte

// An IP mask is an IP address.
type IPMask []byte

// An IPNet represents an IP network.
type IPNet struct {
	IP   IP     // network number
	Mask IPMask // network mask
}

// This data is stored in the Session DB and is accessed during DHCP exchanges
type DhcpV4SessData struct {
	TxId     []byte `json:"txId,omitempty"` //4 bytes for V4
	LeaseSec uint32 `json:"leaseSec,omitempty"`
	ServerId IP     `json:"serverId,omitempty"` // used for request/renew/rebind/release
	PoolId   string `json:"poolId,omitempty"`
}

// This data is stored in the Session DB and is accessed during DHCP exchanges
type DhcpV6SessData struct {
	TxId              []byte `json:"txId,omitempty"`              // 3Bytes for V6
	PreferredLeaseSec uint32 `json:"preferredLeaseSec,omitempty"` // overall lease time
	T1Sec             uint32 `json:"t1Sec,omitempty"`             // renew time
	T2Sec             uint32 `json:"t2Sec,omitempty"`             //rebind time
	ServerId          []byte `json:"serverId,omitempty"`          // server duid converted to []byte
	PoolId            string `json:"poolId,omitempty"`
}

type DhcpData struct {
	V4SessData *DhcpV4SessData `json:"v4SessData,omitempty"`
	V6SessData *DhcpV6SessData `json:"v6SessData,omitempty"`
}

type AuthData struct {
	AcctStart       bool          `json:"acctStart"`             // indicate that an account start record was sent
	RadiusState     []byte        `json:"radiusState,omitempty"` // need to keep state returned from server during eap transactions
	Eap             []byte        `json:"eap,omitempty"`
	Identity        []byte        `json:"identity,omitempty"`
	Class           []byte        `json:"class,omitempty"`
	AuthSessAmbr    *Ambr         `json:"authSessAmbr,omitempty"`
	Ipv4Addr        IP            `json:"ipv4Addr,omitempty"`
	Ipv6Prefix      *IPNet        `json:"ipv6Prefix,omitempty"`
	FramedRouting   FramedRouting `json:"framedRouting,omitempty"`
	FramedRoute     []byte        `json:"framedRoute,omitempty"`
	FramedIPv6Route []byte        `json:"framedIPv6Route,omitempty"`
}

type EventTypeOld = string
type SmfEvent = string

const (
	Access_Type_Event         SmfEvent = "AC_TY_CH"
	UP_Path_Change_Event      SmfEvent = "UP_PATH_CH"
	PDU_Session_Release_Event SmfEvent = "PDU_SES_REL"
	Plmn_Change_Event         SmfEvent = "PLMN_CH"
	UE_IP_Change_Event        SmfEvent = "UE_IP_CH"
)

// DnaiChangeType Possible values are - EARLY: Early notification of UP path reconfiguration. - EARLY_LATE: Early and late notification of UP path reconfiguration. This value shall only be present in the subscription to the DNAI change event. - LATE: Late notification of UP path reconfiguration.

type DnaiChangeType = string

const (
	Dnai_Change_Early_Late DnaiChangeType = "EARLY_LATE"
	Dnai_Change_Early      DnaiChangeType = "EARLY"
	Dnai_Change_Late       DnaiChangeType = "LATE"
)

// EventSubscription struct for EventSubscription
type EventSubscription struct {
	//201909 change to SmfEvent
	EventOld    EventTypeOld    `json:"eventTypeOld,omitempty"`
	Event       SmfEvent        `json:"event,omitempty"`
	DnaiChgType *DnaiChangeType `json:"dnaiChgType,omitempty"`
}

// NotificationMethod Possible values are - PERIODIC - ONE_TIME - ON_EVENT_DETECTION
type notification_method_type = string

const (
	On_Event_Detection notification_method_type = "ON_EVENT_DETECTION"
)

// NsmfEventExposure struct for NsmfEventExposure
type NsmfEventExposure struct {
	Supi string `json:"supi,omitempty"`
	Gpsi string `json:"gpsi,omitempty"`
	// Any UE indication. This IE shall be present if the event subscription is applicable to any UE. Default value \"FALSE\" is used, if not present.
	AnyUeInd *bool  `json:"anyUeInd,omitempty"`
	GroupId  string `json:"groupId,omitempty"`
	PduSeId  *int32 `json:"pduSeId,omitempty"`
	// Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501 [2]. In an OpenAPI [10] schema, the format shall be designated as \"SubId\".
	SubId string `json:"subId,omitempty"`
	// Notification Correlation ID assigned by the NF service consumer.
	NotifId  string `json:"notifId"`
	NotifUri string `json:"notifUri"`
	// Alternate or backup IPv4 Addess(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Addess(es) where to send Notifications.
	AltNotifIpv6Addrs []string `json:"altNotifIpv6Addrs,omitempty"`
	// Subscribed events
	EventSubs    []EventSubscription       `json:"eventSubs"`
	ImmeRep      *bool                     `json:"immeRep,omitempty"`
	NotifMethod  *notification_method_type `json:"notifMethod,omitempty"`
	MaxReportNbr *int32                    `json:"maxReportNbr,omitempty"`
	Expiry       *time.Time                `json:"expiry,omitempty"`
	RepPeriod    *int32                    `json:"repPeriod,omitempty"`
	Guami        *Guami                    `json:"guami,omitempty"`
	// If the NF service consumer is an AMF, it should provide the name of a service produced by the AMF that makes use of notifications about subscribed events.
	ServiveName       string `json:"serviveName,omitempty"`
	SupportedFeatures string `json:"supportedFeatures,omitempty"`
}

type EventInfo struct {
	EventMonitorsMap map[SmfEvent][]string        `json:"eventMonitorsMap,omitempty"` // Take the event type as key to save the mapping of the event to the subscriber
	SubscribersInfo  map[string]NsmfEventExposure `json:"subscribersInfo,omitempty"`  // Save subscription information with subId as key
}
