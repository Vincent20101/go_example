package base

import (
	"gitlab.casa-systems.com/mobility/smf/infra/protocol/go-diameter/diam/datatype"
)

// Diameter vendor identifier
const DIAM_VENDOR_ID_3GPP uint32 = 20858

// Diameter application identifier
const (
	DIAM_APP_ID_BASE     uint32 = 0        // Diameter Common Messages - Diameter protocol association establishment/teardown/maintenance
	DIAM_APP_ID_NASREQ   uint32 = 1        // NASREQ - As defined in RFC7155
	DIAM_APP_ID_BA_RF    uint32 = 3        // Diameter Base Accounting and 3GPP Rf - IMS CTF to CDF interface
	DIAM_APP_ID_CC_GY_RO uint32 = 4        // Diameter Credit Control and 3GPP Gy/Ro - interface between PCEF and OCS
	DIAM_APP_ID_CX_DX    uint32 = 16777216 // 3GPP Cx/Dx - IMS I/S-CSCF to HSS interface
	DIAM_APP_ID_SH       uint32 = 16777217 // 3GPP Sh - VoIP/IMS SIP Application Server to HSS interface
	DIAM_APP_ID_RX       uint32 = 16777236 // 3GPP Rx - Policy and charging control - Interface between CRF and AF
	DIAM_APP_ID_GX       uint32 = 16777238 // 3GPP Gx - Policy and charging control - interface between PCEF and PCRF
	DIAM_APP_ID_S6A_S6D  uint32 = 16777251 // 3GPP S6a/S6d - LTE Roaming signaling
	DIAM_APP_ID_S13      uint32 = 16777252 // 3GPP S13 - Interface between EIR and MME
	DIAM_APP_ID_SLG      uint32 = 16777255 // 3GPP LCS - SLg Location services
	DIAM_APP_ID_S6T      uint32 = 16777345 // 3GPP S6t - Interface between SCEF and HSS
	DIAM_APP_ID_S6B      uint32 = 16777272 // 3GPP S6b - Interface between PDN-GW and 3GPP AAA Server
	DIAM_APP_ID_S6B_LITE uint32 = 16777995 // VzW  S6b-DN-lite - Verizon Proprietary Interface between SMF and MPN-Proxy(DN-AAA)
)

type AcctRecordType datatype.Enumerated

const (
	AcctRecordType_EVENT   AcctRecordType = 1
	AcctRecordType_START   AcctRecordType = 2
	AcctRecordType_INTERIM AcctRecordType = 3
	AcctRecordType_STOP    AcctRecordType = 4
)

type DRMP datatype.Enumerated

const (
	DRMP_PRIORITY_0  DRMP = 0
	DRMP_PRIORITY_1  DRMP = 1
	DRMP_PRIORITY_2  DRMP = 2
	DRMP_PRIORITY_3  DRMP = 3
	DRMP_PRIORITY_4  DRMP = 4
	DRMP_PRIORITY_5  DRMP = 5
	DRMP_PRIORITY_6  DRMP = 6
	DRMP_PRIORITY_7  DRMP = 7
	DRMP_PRIORITY_8  DRMP = 8
	DRMP_PRIORITY_9  DRMP = 9
	DRMP_PRIORITY_10 DRMP = 10
	DRMP_PRIORITY_11 DRMP = 11
	DRMP_PRIORITY_12 DRMP = 12
	DRMP_PRIORITY_13 DRMP = 13
	DRMP_PRIORITY_14 DRMP = 14
	DRMP_PRIORITY_15 DRMP = 15
)

type AuthRequestType datatype.Enumerated

const (
	AuthRequestType_AUTHENTICATE_ONLY      AuthRequestType = 1
	AuthRequestType_AUTHORIZE_ONLY         AuthRequestType = 2
	AuthRequestType_AUTHORIZE_AUTHENTICATE AuthRequestType = 3
)

type ReAuthRequestType datatype.Enumerated

const (
	ReAuthRequestType_AUTHORIZE_ONLY         ReAuthRequestType = 0
	ReAuthRequestType_AUTHORIZE_AUTHENTICATE ReAuthRequestType = 1
)

type SessionReleaseCause datatype.Enumerated

const (
	SessionReleaseCause_UNSPECIFIED_REASON            SessionReleaseCause = 0
	SessionReleaseCause_UE_SUBSCRIPTION_REASON        SessionReleaseCause = 1
	SessionReleaseCause_INSUFFICIENT_SERVER_RESOURCES SessionReleaseCause = 2
	SessionReleaseCause_IP_CAN_SESSION_TERMINATION    SessionReleaseCause = 3
	SessionReleaseCause_UE_IP_ADDRESS_RELEASE         SessionReleaseCause = 4
)

type EventTrigger datatype.Enumerated

const (
	EventTrigger_SGSN_CHANGE                                             EventTrigger = 0
	EventTrigger_QOS_CHANGE                                              EventTrigger = 1
	EventTrigger_RAT_CHANGE                                              EventTrigger = 2
	EventTrigger_TFT_CHANGE                                              EventTrigger = 3
	EventTrigger_PLMN_CHANGE                                             EventTrigger = 4
	EventTrigger_LOSS_OF_BEARER                                          EventTrigger = 5
	EventTrigger_RECOVERY_OF_BEARER                                      EventTrigger = 6
	EventTrigger_IP_CAN_CHANGE                                           EventTrigger = 7
	EventTrigger_QOS_CHANGE_EXCEEDING_AUTHORIZATION                      EventTrigger = 11
	EventTrigger_RAI_CHANGE                                              EventTrigger = 12
	EventTrigger_USER_LOCATION_CHANGE                                    EventTrigger = 13
	EventTrigger_NO_EVENT_TRIGGERS                                       EventTrigger = 14
	EventTrigger_OUT_OF_CREDIT                                           EventTrigger = 15
	EventTrigger_REALLOCATION_OF_CREDIT                                  EventTrigger = 16
	EventTrigger_REVALIDATION_TIMEOUT                                    EventTrigger = 17
	EventTrigger_UE_IP_ADDRESS_ALLOCATE                                  EventTrigger = 18
	EventTrigger_UE_IP_ADDRESS_RELEASE                                   EventTrigger = 19
	EventTrigger_DEFAULT_EPS_BEARER_QOS_CHANGE                           EventTrigger = 20
	EventTrigger_AN_GW_CHANGE                                            EventTrigger = 21
	EventTrigger_SUCCESSFUL_RESOURCE_ALLOCATION                          EventTrigger = 22
	EventTrigger_RESOURCE_MODIFICATION_REQUEST                           EventTrigger = 23
	EventTrigger_PGW_TRACE_CONTROL                                       EventTrigger = 24
	EventTrigger_UE_TIME_ZONE_CHANGE                                     EventTrigger = 25
	EventTrigger_TAI_CHANGE                                              EventTrigger = 26
	EventTrigger_ECGI_CHANGE                                             EventTrigger = 27
	EventTrigger_CHARGING_CORRELATION_EXCHANGE                           EventTrigger = 28
	EventTrigger_APN_AMBR_MODIFICATION_FAILURE                           EventTrigger = 29
	EventTrigger_USER_CSG_INFORMATION_CHANGE                             EventTrigger = 30
	EventTrigger_USAGE_REPORT                                            EventTrigger = 33
	EventTrigger_DEFAULT_EPS_BEARER_QOS_MODIFICATION_FAILURE             EventTrigger = 34
	EventTrigger_USER_CSG_HYBRID_SUBSCRIBED_INFORMATION_CHANGE           EventTrigger = 35
	EventTrigger_USER_CSG_HYBRID_UNSUBSCRIBED_INFORMATION_CHANGE         EventTrigger = 36
	EventTrigger_ROUTING_RULE_CHANGE                                     EventTrigger = 37
	EventTrigger_APPLICATION_START                                       EventTrigger = 39
	EventTrigger_APPLICATION_STOP                                        EventTrigger = 40
	EventTrigger_CS_TO_PS_HANDOVER                                       EventTrigger = 42
	EventTrigger_UE_LOCAL_IP_ADDRESS_CHANGE                              EventTrigger = 43
	EventTrigger_HNB_LOCAL_IP_ADDRESS_CHANGE                             EventTrigger = 44
	EventTrigger_ACCESS_NETWORK_INFO_REPORT                              EventTrigger = 45
	EventTrigger_CREDIT_MANAGEMENT_SESSION_FAILURE                       EventTrigger = 46
	EventTrigger_DEFAULT_QOS_CHANGE                                      EventTrigger = 47
	EventTrigger_CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA_REPORT EventTrigger = 48
	EventTrigger_ADDITION_OF_ACCESS                                      EventTrigger = 49
	EventTrigger_REMOVAL_OF_ACCESS                                       EventTrigger = 50
	EventTrigger_UNAVAILABLITY_OF_ACCESS                                 EventTrigger = 51
	EventTrigger_AVAILABLITY_OF_ACCESS                                   EventTrigger = 52
	EventTrigger_RESOURCE_RELEASE                                        EventTrigger = 53
	EventTrigger_ENODEB_CHANGE                                           EventTrigger = 54
	EventTrigger_3GPP_PS_DATA_OFF_CHANGE                                 EventTrigger = 55
	EventTrigger_UE_STATUS_RESUME                                        EventTrigger = 56
)

/*
+-------------------------------------+-----------+---------+----+----+-----+----------+
| AVP                                 | Type      |  Layers | R8 | R9 | R10 | mutative |
+-------------------------------------+-----------+---------+----+----+-----+----------+
| AN-Trusted                          | enumerate |         |    |    |     |          |
| Event-Trigger                       | enumerate |         | 1  | 1  |     |          |
| User-CSG-Information                | struct    | 1       |    | 1  |     |          |
| IP-CAN-Type                         | enumerate |         |    |    |     |          |
| AN-GW-Address                       | string    |         |    |    |     |          |
| 3GPP-SGSN-Address                   | string    |         | 1  | 1  |     |          |
| 3GPP-SGSN-IPv6-Address              | string    |         | 1  | 1  |     |          |
| 3GPP-SGSN-MCC-MNC                   | string    |         |    |    |     |          |
| Framed-IP-Address                   | string    |         |    |    |     |          |
| RAT-Type                            | enumerate |         | 1  | 1  |     |          |
| QoS-Information                     | struct    | 2       |    |    |     | yes      |
| RAI                                 | enumerate |         | 1  | 1  |     |          |
| 3GPP-User-Location-Info             | string    |         | 1  | 1  |     |          |
| Trace-Data                          | struct    | 1       | 1  | 1  |     |          |
| Trace-Reference                     | string    |         | 1  | 1  |     |          |
| 3GPP2-BSID                          | string    |         | 1  | 1  |     |          |
| 3GPP-MS-TimeZone                    | string    |         | 1  | 1  |     |          |
| Routing-IP-Address                  | byte      |         |    |    |     |          |
| UE-Local-IP-Address                 | byte      |         |    |    |     |          |
| HeNB-Local-IP-Address               | byte      |         |    |    |     |          |
| UDP-Source-Port                     | string    |         |    |    |     |          |
| Presence-Reporting-Area-Information | struct    | 1       |    |    |     |          |
+-------------------------------------+-----------+---------+----+----+-----+----------+
*/
type EventReportIndication struct {
	AnTrusted                        *AnTrusted                        `avp:"AN-Trusted" json:"anTrusted,omitempty"`
	EventTrigger                     []EventTrigger                    `avp:"Event-Trigger" json:"eventTrigger,omitempty"`
	UserCsgInformation               *UserCsgInformation               `avp:"User-CSG-Information" json:"userCsgInformation,omitempty"`
	IpCanType                        *IpCanType                        `avp:"IP-CAN-Type" json:"ipCanType,omitempty"`
	AnGwAddress                      []datatype.Address                `avp:"AN-GW-Address" json:"anGwAddress,omitempty"`
	TgppSgsnAddress                  *datatype.OctetString             `avp:"3GPP-SGSN-Address" json:"tgppSgsnAddress,omitempty"`
	TgppSgsnIpv6Address              *datatype.OctetString             `avp:"3GPP-SGSN-IPv6-Address" json:"tgppSgsnIpv6Address,omitempty"`
	TgppSgsnMccMnc                   *datatype.UTF8String              `avp:"3GPP-SGSN-MCC-MNC" json:"tgppSgsnMccMnc,omitempty"`
	FramedIpAddress                  *datatype.OctetString             `avp:"Framed-IP-Address" json:"framedIpAddress,omitempty"`
	RatType                          *RatType                          `avp:"RAT-Type" json:"ratType,omitempty"`
	QosInformation                   []QoSInformation                  `avp:"QoS-Information" json:"qosInformation,omitempty"`
	Rai                              *datatype.UTF8String              `avp:"RAI" json:"rai,omitempty"`
	TgppUserLocationInfo             *datatype.OctetString             `avp:"3GPP-User-Location-Info" json:"tgppUserLocationInfo,omitempty"`
	TraceData                        *TraceData                        `avp:"Trace-Data" json:"traceData,omitempty"`
	TraceReference                   *datatype.OctetString             `avp:"Trace-Reference" json:"TraceReference,omitempty"`
	Tgpp2BSId                        *datatype.OctetString             `avp:"3GPP2-BSID" json:"tgpp2BSId,omitempty"`
	TgppMsTimeZone                   *datatype.OctetString             `avp:"3GPP-MS-TimeZone" json:"tgppMsTimeZone,omitempty"`
	RoutingIPAddress                 *datatype.Address                 `avp:"Routing-IP-Address" json:"routingIPAddress,omitempty"`
	UeLocalIpAddress                 *datatype.Address                 `avp:"UE-Local-IP-Address" json:"ueLocalIpAddress,omitempty"`
	HeNBLocalIpAddress               *datatype.Address                 `avp:"HeNB-Local-IP-Address" json:"heNBLocalIpAddress,omitempty"`
	UpfSourcePort                    *datatype.Unsigned32              `avp:"UDP-Source-Port" json:"upfSourcePort,omitempty"`
	PresenceReportingAreaInformation *PresenceReportingAreaInformation `avp:"Presence-Reporting-Area-Information" json:"presenceReportingAreaInformation,omitempty"`
}

type AnTrusted datatype.Enumerated

const (
	AnTrusted_TRUSTED   AnTrusted = 0
	AnTrusted_UNTRUSTED AnTrusted = 1
)

type UserCsgInformation struct {
	CsgId                   datatype.Unsigned32      `avp:"CSG-Id" json:"csgId"`
	CsgAccessMode           CsgAccessMode            `avp:"CSG-Access-Mode" json:"csgAccessMode"`
	CsgMembershipIndication *CsgMembershipIndication `avp:"CSG-Membership-Indication" json:"csgMembershipIndication,omitempty"`
}

type CsgAccessMode datatype.Enumerated

const (
	CsgAccessMode_Closed_mode CsgAccessMode = 0
	CsgAccessMode_Hybrid_mode CsgAccessMode = 1
)

type CsgMembershipIndication datatype.Enumerated

const (
	CsgMembershipIndication_Not_CSG_member CsgMembershipIndication = 0
	CsgMembershipIndication_CSG_Member     CsgMembershipIndication = 1
)

type IpCanType datatype.Enumerated

const (
	IpCanType_3GPP_GPRS    IpCanType = 0
	IpCanType_DOCSIS       IpCanType = 1
	IpCanType_xDSL         IpCanType = 2
	IpCanType_WiMAX        IpCanType = 3
	IpCanType_3GPP2        IpCanType = 4
	IpCanType_3GPP_EPS     IpCanType = 5
	IpCanType_Non_3GPP_EPS IpCanType = 6
	IpCanType_FBA          IpCanType = 7
	IpCanType_3GPP_5GS     IpCanType = 8
	IpCanType_Non_3GPP_5GS IpCanType = 9
)

// Ts 29212 5.3.31	RAT-Type AVP
//Value       Description
//0           WLAN
//1           VIRTUAL
//2           TRUSTED-N3GA
//3           WIRELINE
//4           WIRELINE-CABLE
//5           WIRELINE-BBF
//1000        UTRAN
//1001        GERAN
//1002        GAN
//1003        HSPA_EVOLUTION
//1004        EUTRAN
//1005        EUTRAN-NB-IoT
//1006        NR
//1007        LTE-M
//1008        NR-U
//2000        CDMA2000_1X
//2001        HRPD
//2002        UMB
//2003        EHRPD

type RatType datatype.Enumerated

const (
	RatType_WLAN           RatType = 0
	RatType_VIRTUAL        RatType = 1
	RatType_TRUSTED_N3GA   RatType = 2
	RatType_WIRELINE       RatType = 3
	RatType_WIRELINE_CABLE RatType = 4
	RatType_WIRELINE_BBF   RatType = 5
	RatType_UTRAN          RatType = 1000
	RatType_GERAN          RatType = 1001
	RatType_GAN            RatType = 1002
	RatType_HSPA_EVOLUTION RatType = 1003
	RatType_EUTRAN         RatType = 1004
	RatType_EUTRAN_NBIOT   RatType = 1005
	RatType_NR             RatType = 1006
	RatType_LTE_M          RatType = 1007
	RatType_NR_U           RatType = 1008
	RatType_CDMA2000_1X    RatType = 2000
	RatType_HRPD           RatType = 2001
	RatType_UMB            RatType = 2002
	RatType_EHRPD          RatType = 2003
)

type TraceData struct {
	TraceReference        datatype.OctetString  `avp:"Trace-Reference" json:"traceReference"`
	TraceDepth            TraceDepth            `avp:"Trace-Depth" json:"traceDepth"`
	TraceNeTypeList       datatype.OctetString  `avp:"Trace-NE-Type-List" json:"traceNeTypeList"`
	TraceInterfaceList    *datatype.OctetString `avp:"Trace-Interface-List" json:"traceInterfaceList,omitempty"`
	TraceEventList        datatype.OctetString  `avp:"Trace-Event-List" json:"traceEventList,"`
	OmcId                 *datatype.OctetString `avp:"OMC-Id" json:"omcId,omitempty"`
	TraceCollectionEntity datatype.Address      `avp:"Trace-Collection-Entity" json:"traceCollectionEntity"`
	MdtConfiguration      *MdtConfiguration     `avp:"MDT-Configuration" json:"mdtConfiguration,omitempty"`
}

type TraceDepth datatype.Enumerated

const (
	TraceDepth_Minimum                               TraceDepth = 0
	TraceDepth_Medium                                TraceDepth = 1
	TraceDepth_Maximum                               TraceDepth = 2
	TraceDepth_MinimumWithoutVendorSpecificExtension TraceDepth = 3
	TraceDepth_MediumWithoutVendorSpecificExtension  TraceDepth = 4
	TraceDepth_MaximumWithoutVendorSpecificExtension TraceDepth = 5
)

type ReservationPriority datatype.Enumerated

const (
	ReservationPriority_DEFAULT ReservationPriority = 0
	ReservationPriority_ONE     ReservationPriority = 1
	ReservationPriority_TWO     ReservationPriority = 2
	ReservationPriority_THREE   ReservationPriority = 3
	ReservationPriority_FOUR    ReservationPriority = 4
	ReservationPriority_FIVE    ReservationPriority = 5
	ReservationPriority_SIX     ReservationPriority = 6
	ReservationPriority_SEVEN   ReservationPriority = 7
)

type MdtConfiguration struct {
	JobType                 JobType                  `avp:"Job-Type" json:"jobType"`
	AreaScope               *AreaScope               `avp:"Area-Scope" json:"areaScope,omitempty"`
	ListOfMeasurements      *datatype.Unsigned32     `avp:"List-Of-Measurements" json:"listOfMeasurements,omitempty"`
	ReportingTrigger        *datatype.Unsigned32     `avp:"Reporting-Trigger" json:"reportingTrigger,omitempty"`
	ReportInterval          *ReportInterval          `avp:"Report-Interval" json:"reportInterval,omitempty"`
	ReportAmount            *ReportAmount            `avp:"Report-Amount" json:"reportAmount,omitempty"`
	EventThresholdRsrp      *datatype.Unsigned32     `avp:"Event-Threshold-RSRP" json:"eventThresholdRsrp,omitempty"`
	EventThresholdRsrq      *datatype.Unsigned32     `avp:"Event-Threshold-RSRQ" json:"eventThresholdRsrq,omitempty"`
	LoggingInterval         *LoggingInterval         `avp:"Logging-Interval" json:"loggingInterval,omitempty"`
	LoggingDuration         *LoggingDuration         `avp:"Logging-Duration" json:"loggingDuration,omitempty"`
	MeasurementPeriodLte    *MeasurementPeriodLte    `avp:"Measurement-Period-LTE" json:"measurementPeriodLte,omitempty"`
	MeasurementPeriodUmts   *MeasurementPeriodUmts   `avp:"Measurement-Period-UMTS" json:"measurementPeriodUmts,omitempty"`
	CollectionPeriodRrmLte  *CollectionPeriodRrmLte  `avp:"Collection-Period-RRM-LTE" json:"collectionPeriodRrmLte,omitempty"`
	CollectionPeriodRrmUmts *CollectionPeriodRrmUmts `avp:"Collection-Period-RRM-UMTS" json:"collectionPeriodRrmUmts,omitempty"`
	PositioningMethod       *datatype.OctetString    `avp:"Positioning-Method" json:"positioningMethod,omitempty"`
	MeasurementQuantity     *datatype.OctetString    `avp:"Measurement-Quantity" json:"measurementQuantity,omitempty"`
	EventThresholdEvent1F   *datatype.Integer32      `avp:"Event-Threshold-Event-1F" json:"eventThresholdEvent1F,omitempty"`
	EventThresholdEvent1I   *datatype.Integer32      `avp:"Event-Threshold-Event-1I" json:"eventThresholdEvent1I,omitempty"`
	MdtAllowedPlmnId        []datatype.OctetString   `avp:"MDT-Allowed-PLMN-Id" json:"mdtAllowedPlmnId,omitempty"`
	MbsfnArea               []MbsfnArea              `avp:"MBSFN-Area" json:"MbsfnArea,omitempty"`
}

type JobType datatype.Enumerated

const (
	JobType_Immediate_MDT_only      JobType = 0
	JobType_Logged_MDT_only         JobType = 1
	JobType_Trace_only              JobType = 2
	JobType_Immediate_MDT_and_Trace JobType = 3
	JobType_RLF_reports_only        JobType = 4
	JobType_RCEF_reports_only       JobType = 5
	JobType_Logged_MBSFN_MDT        JobType = 6
)

type AreaScope struct {
	CellGlobalIdentity       []datatype.OctetString `avp:"Cell-Global-Identity" json:"cellGlobalIdentity,omitempty"`
	EUtranCellGlobalIdentity []datatype.OctetString `avp:"E-UTRAN-Cell-Global-Identity" json:"eUtranCellGlobalIdentity,omitempty"`
	RoutingAreaIdentity      []datatype.OctetString `avp:"Routing-Area-Identity" json:"routingAreaIdentity,omitempty"`
	LocationAreaIdentity     []datatype.OctetString `avp:"Location-Area-Identity" json:"locationAreaIdentity,omitempty"`
	TrackingAreaIdentity     []datatype.OctetString `avp:"Tracking-Area-Identity" json:"trackingAreaIdentity,omitempty"`
}

type ReportInterval datatype.Enumerated

const (
	ReportInterval_UTMTS_250   ReportInterval = 0
	ReportInterval_UTMTS_500   ReportInterval = 1
	ReportInterval_UTMTS_1000  ReportInterval = 2
	ReportInterval_UTMTS_2000  ReportInterval = 3
	ReportInterval_UTMTS_3000  ReportInterval = 4
	ReportInterval_UTMTS_4000  ReportInterval = 5
	ReportInterval_UTMTS_6000  ReportInterval = 6
	ReportInterval_UTMTS_8000  ReportInterval = 7
	ReportInterval_UTMTS_12000 ReportInterval = 8
	ReportInterval_UTMTS_16000 ReportInterval = 9
	ReportInterval_UTMTS_20000 ReportInterval = 10
	ReportInterval_UTMTS_24000 ReportInterval = 11
	ReportInterval_UTMTS_28000 ReportInterval = 12
	ReportInterval_UTMTS_32000 ReportInterval = 13
	ReportInterval_UTMTS_64000 ReportInterval = 14
	ReportInterval_LTE_120     ReportInterval = 15
	ReportInterval_LTE_240     ReportInterval = 16
	ReportInterval_LTE_480     ReportInterval = 17
	ReportInterval_LTE_640     ReportInterval = 18
	ReportInterval_LTE_1024    ReportInterval = 19
	ReportInterval_LTE_2048    ReportInterval = 20
	ReportInterval_LTE_5120    ReportInterval = 21
	ReportInterval_LTE_10240   ReportInterval = 22
	ReportInterval_LTE_60000   ReportInterval = 23
	ReportInterval_LTE_360000  ReportInterval = 24
	ReportInterval_LTE_720000  ReportInterval = 25
	ReportInterval_LTE_1800000 ReportInterval = 26
	ReportInterval_LTE_3600000 ReportInterval = 27
	ReportInterval_NR_120      ReportInterval = 28
	ReportInterval_NR_240      ReportInterval = 29
	ReportInterval_NR_480      ReportInterval = 30
	ReportInterval_NR_640      ReportInterval = 31
	ReportInterval_NR_1024     ReportInterval = 32
	ReportInterval_NR_2048     ReportInterval = 33
	ReportInterval_NR_5120     ReportInterval = 34
	ReportInterval_NR_10240    ReportInterval = 35
	ReportInterval_NR_20480    ReportInterval = 36
	ReportInterval_NR_40960    ReportInterval = 37
	ReportInterval_NR_60000    ReportInterval = 38
	ReportInterval_NR_360000   ReportInterval = 39
	ReportInterval_NR_720000   ReportInterval = 40
	ReportInterval_NR_1800000  ReportInterval = 41
	ReportInterval_NR_3600000  ReportInterval = 42
)

type ReportAmount datatype.Enumerated

const (
	ReportAmount_1        ReportAmount = 0
	ReportAmount_2        ReportAmount = 1
	ReportAmount_4        ReportAmount = 2
	ReportAmount_8        ReportAmount = 3
	ReportAmount_16       ReportAmount = 4
	ReportAmount_32       ReportAmount = 5
	ReportAmount_64       ReportAmount = 6
	ReportAmount_Infinity ReportAmount = 7
)

type LoggingInterval datatype.Enumerated

const (
	loggingInterval_UMTS_OR_LTE_OR_NR_1280  LoggingInterval = 0
	loggingInterval_UMTS_OR_LTE_OR_NR_2560  LoggingInterval = 1
	loggingInterval_UMTS_OR_LTE_OR_NR_5120  LoggingInterval = 2
	loggingInterval_UMTS_OR_LTE_OR_NR_10240 LoggingInterval = 3
	loggingInterval_UMTS_OR_LTE_OR_NR_20480 LoggingInterval = 4
	loggingInterval_UMTS_OR_LTE_OR_NR_30720 LoggingInterval = 5
	loggingInterval_UMTS_OR_LTE_OR_NR_40960 LoggingInterval = 6
	loggingInterval_UMTS_OR_LTE_OR_NR_61440 LoggingInterval = 7
	loggingInterval_ONLY_NR_320             LoggingInterval = 8
	loggingInterval_ONLY_NR_640             LoggingInterval = 9
	loggingInterval_ONLY_NR_Infinity        LoggingInterval = 10
)

type LoggingDuration datatype.Enumerated

const (
	LoggingDuration_UTMS_OR_LTE_600  LoggingInterval = 0
	LoggingDuration_UTMS_OR_LTE_1200 LoggingInterval = 1
	LoggingDuration_UTMS_OR_LTE_2400 LoggingInterval = 2
	LoggingDuration_UTMS_OR_LTE_3600 LoggingInterval = 3
	LoggingDuration_UTMS_OR_LTE_5400 LoggingInterval = 4
	LoggingDuration_UTMS_OR_LTE_7200 LoggingInterval = 5
	LoggingDuration_NR_600           LoggingInterval = 6
	LoggingDuration_NR_1200          LoggingInterval = 7
	LoggingDuration_NR_2400          LoggingInterval = 8
	LoggingDuration_NR_3600          LoggingInterval = 9
	LoggingDuration_NR_5400          LoggingInterval = 10
	LoggingDuration_NR_7200          LoggingInterval = 11
)

type MeasurementPeriodLte datatype.Enumerated

const (
	MeasurementPeriodLte_1024  MeasurementPeriodLte = 0
	MeasurementPeriodLte_1280  MeasurementPeriodLte = 1
	MeasurementPeriodLte_2048  MeasurementPeriodLte = 2
	MeasurementPeriodLte_2560  MeasurementPeriodLte = 3
	MeasurementPeriodLte_5120  MeasurementPeriodLte = 4
	MeasurementPeriodLte_10240 MeasurementPeriodLte = 5
	MeasurementPeriodLte_60000 MeasurementPeriodLte = 6
)

type MeasurementPeriodUmts datatype.Enumerated

const (
	MeasurementPeriodUmts_250   MeasurementPeriodUmts = 0
	MeasurementPeriodUmts_500   MeasurementPeriodUmts = 1
	MeasurementPeriodUmts_1000  MeasurementPeriodUmts = 2
	MeasurementPeriodUmts_2000  MeasurementPeriodUmts = 3
	MeasurementPeriodUmts_3000  MeasurementPeriodUmts = 4
	MeasurementPeriodUmts_4000  MeasurementPeriodUmts = 5
	MeasurementPeriodUmts_6000  MeasurementPeriodUmts = 6
	MeasurementPeriodUmts_8000  MeasurementPeriodUmts = 7
	MeasurementPeriodUmts_12000 MeasurementPeriodUmts = 8
	MeasurementPeriodUmts_16000 MeasurementPeriodUmts = 9
	MeasurementPeriodUmts_20000 MeasurementPeriodUmts = 10
	MeasurementPeriodUmts_24000 MeasurementPeriodUmts = 11
	MeasurementPeriodUmts_28000 MeasurementPeriodUmts = 12
	MeasurementPeriodUmts_32000 MeasurementPeriodUmts = 13
	MeasurementPeriodUmts_64000 MeasurementPeriodUmts = 14
)

type CollectionPeriodRrmLte datatype.Enumerated

const (
	CollectionPeriodRrmLte_1024  CollectionPeriodRrmLte = 0
	CollectionPeriodRrmLte_1280  CollectionPeriodRrmLte = 1
	CollectionPeriodRrmLte_2048  CollectionPeriodRrmLte = 2
	CollectionPeriodRrmLte_2560  CollectionPeriodRrmLte = 3
	CollectionPeriodRrmLte_5120  CollectionPeriodRrmLte = 4
	CollectionPeriodRrmLte_10240 CollectionPeriodRrmLte = 5
	CollectionPeriodRrmLte_60000 CollectionPeriodRrmLte = 6
)

type CollectionPeriodRrmUmts datatype.Enumerated

const (
	CollectionPeriodRrmUmts_250   CollectionPeriodRrmUmts = 0
	CollectionPeriodRrmUmts_500   CollectionPeriodRrmUmts = 1
	CollectionPeriodRrmUmts_1000  CollectionPeriodRrmUmts = 2
	CollectionPeriodRrmUmts_2000  CollectionPeriodRrmUmts = 3
	CollectionPeriodRrmUmts_3000  CollectionPeriodRrmUmts = 4
	CollectionPeriodRrmUmts_4000  CollectionPeriodRrmUmts = 5
	CollectionPeriodRrmUmts_6000  CollectionPeriodRrmUmts = 6
	CollectionPeriodRrmUmts_8000  CollectionPeriodRrmUmts = 7
	CollectionPeriodRrmUmts_12000 CollectionPeriodRrmUmts = 8
	CollectionPeriodRrmUmts_16000 CollectionPeriodRrmUmts = 9
	CollectionPeriodRrmUmts_20000 CollectionPeriodRrmUmts = 10
	CollectionPeriodRrmUmts_24000 CollectionPeriodRrmUmts = 11
	CollectionPeriodRrmUmts_28000 CollectionPeriodRrmUmts = 12
	CollectionPeriodRrmUmts_32000 CollectionPeriodRrmUmts = 13
	CollectionPeriodRrmUmts_64000 CollectionPeriodRrmUmts = 14
)

type MbsfnArea struct {
	MbsfnAreaId      *datatype.Unsigned32 `avp:"MBSFN-Area-ID" json:"mbsfnAreaId,omitempty"`
	CarrierFrequency *datatype.Unsigned32 `avp:"Carrier-Frequency" json:"carrierFrequency,omitempty"`
}

type ExperimentalResult struct {
	VendorId               datatype.Unsigned32 `avp:"Vendor-Id" json:"vendorId"`
	ExperimentalResultCode datatype.Unsigned32 `avp:"Experimental-Result-Code" json:"experimentalResultCode"`
}

// PermanentFailure see 29212 5.5.3 Permanent Failures
type PermanentFailure = datatype.Unsigned32

const (
	DIAMETER_USER_UNKNOWN                        PermanentFailure = 5030
	DIAMETER_ERROR_LATE_OVERLAPPING_REOUEST      PermanentFailure = 5453
	DIAMETER_ERROR_TIMED_OUT_REQUEST             PermanentFailure = 5454
	DIAMETER_ERROR_INITIAL_PARAMETERS            PermanentFailure = 5140
	DIAMETER_ERROR_TRIGGER_EVENT                 PermanentFailure = 5141
	DIAMETER_PCC_RULE_EVENT                      PermanentFailure = 5142
	DIAMETER_ERROR_BEARER_NOT_AUTHORIZED         PermanentFailure = 5143
	DIAMETER_ERROR_TRAFFIC_MAPPING_INFO_REJECTED PermanentFailure = 5144
	DIAMETER_ERROR_CONFLICTING_REQUEST           PermanentFailure = 5147
	DIAMETER_ADC_RULE_EVENT                      PermanentFailure = 5148
	DIAMETER_ERROR_NBIFOM_NOT_AUTHORIZED         PermanentFailure = 5149
)

// TransientFailure see 29212 5.5.4 Transient Failures
type TransientFailure = datatype.Unsigned32

const (
	DIAMETER_PCC_BEARER_EVENT    TransientFailure = 4141
	DIAMETER_AN_GW_FAILED        TransientFailure = 4143
	DIAMETER_PENDING_TRANSACTION TransientFailure = 4144
	DIAMETER_UE_STATUS_SUSPEND   TransientFailure = 4145
)

type OCOLR struct {
	OCSequenceNumber      datatype.Unsigned64  `avp:"OC-Sequence-Number" json:"oCSequenceNumber"`
	OCReportType          OCReportType         `avp:"OC-Report-Type" json:"oCReportType"`
	OCReductionPercentage *datatype.Unsigned32 `avp:"OC-Reduction-Percentage" json:"oCReductionPercentage,omitempty"`
	OCValidityDuration    *datatype.Unsigned32 `avp:"OC-Validity-Duration" json:"oCValidityDuration,omitempty"`
}

type OCReportType datatype.Enumerated

const (
	OCReportType_HOST_REPORT  OCReportType = 0
	OCReportType_REALM_REPORT OCReportType = 1
)

type BearerControlMode datatype.Enumerated

const (
	BearerControlMode_UE_ONLY  BearerControlMode = 0
	BearerControlMode_RESERVED BearerControlMode = 1
	BearerControlMode_UE_NW    BearerControlMode = 2
)

type RedirectHostUsage datatype.Enumerated

const (
	RedirectHostUsage_DONT_CACHE            RedirectHostUsage = 0
	RedirectHostUsage_ALL_SESSION           RedirectHostUsage = 1
	RedirectHostUsage_ALL_REALM             RedirectHostUsage = 2
	RedirectHostUsage_REALM_AND_APPLICATION RedirectHostUsage = 3
	RedirectHostUsage_ALL_APPLICATION       RedirectHostUsage = 4
	RedirectHostUsage_ALL_HOST              RedirectHostUsage = 5
	RedirectHostUsage_ALL_USER              RedirectHostUsage = 6
)

/*
+-------------------------------+-----------+---------+----+----+-----+----------+
| AVP                           | Type      |  Layers | R8 | R9 | R10 | mutative |
+-------------------------------+-----------+---------+----+----+-----+----------+
| Charging-Rule-Name            | string    |         | 1  | 1  |     |          |
| Charging-Rule-Base-Name       | string    |         | 1  | 1  |     |          |
| Required-Access-Info          | enumerate |         |    |    |     |          |
| Resource-Release-Notification | enumerate |         |    |    |     |          |
+-------------------------------+-----------+---------+----+----+-----+----------+
*/
type ChargingRuleRemove struct {
	ChargingRuleName            []datatype.OctetString       `avp:"Charging-Rule-Name" json:"chargingRuleName,omitempty"`
	ChargingRuleBaseName        []datatype.UTF8String        `avp:"Charging-Rule-Base-Name" json:"chargingRuleBaseName,omitempty"`
	RequiredAccessInfo          []RequiredAccessInfo         `avp:"Required-Access-Info" json:"requiredAccessInfo,omitempty"`
	ResourceReleaseNotification *ResourceReleaseNotification `avp:"Resource-Release-Notification" json:"resourceReleaseNotification,omitempty"`
}

/*
+----------------------------------+-----------+---------+----+----+-----+----------+
| AVP                              | Type      |  Layers | R8 | R9 | R10 | mutative |
+----------------------------------+-----------+---------+----+----+-----+----------+
| Charging-Rule-Definition         | struct    | 3       | 1  | 1  |     | yes      |
| Charging-Rule-Name               | string    |         | 1  | 1  |     |          |
| Charging-Rule-Base-Name          | string    |         | 1  | 1  |     |          |
| Bearer-Identifier                | string    |         | 1  | 1  |     |          |
| Monitoring-Flags                 | uint32    |         |    |    |     |          |
| Rule-Activation-Time             | time      |         | 1  | 1  |     |          |
| Rule-Deactivation-Time           | time      |         | 1  | 1  |     |          |
| Resource-Allocation-Notification | enumerate |         | 1  | 1  |     |          |
| Charging-Correlation-Indicator   | enumerate |         | 1  | 1  |     |          |
| IP-CAN-Type                      | enumerate |         |    |    |     |          |
+----------------------------------+-----------+---------+----+----+-----+----------+
*/
type ChargingRuleInstall struct {
	ChargingRuleDefinition         []ChargingRuleDefinition        `avp:"Charging-Rule-Definition" json:"ChargingRuleDefinition,omitempty"`
	ChargingRuleName               []datatype.OctetString          `avp:"Charging-Rule-Name" json:"chargingRuleName,omitempty"`
	ChargingRuleBaseName           []datatype.UTF8String           `avp:"Charging-Rule-Base-Name" json:"chargingRuleBaseName,omitempty"`
	BearerIdentifier               *datatype.OctetString           `avp:"Bearer-Identifier" json:"bearerIdentifier,omitempty"`
	MonitoringFlags                *datatype.Unsigned32            `avp:"Monitoring-Flags" json:"monitoringFlags,omitempty"`
	RuleActivationTime             *datatype.Time                  `avp:"Rule-Activation-Time" json:"ruleActivationTime,omitempty"`
	RuleDeactivationTime           *datatype.Time                  `avp:"Rule-Deactivation-Time" json:"ruleDeactivationTime,omitempty"`
	ResourceAllocationNotification *ResourceAllocationNotification `avp:"Resource-Allocation-Notification" json:"resourceAllocationNotification,omitempty"`
	ChargingCorrelationIndicator   *ChargingCorrelationIndicator   `avp:"Charging-Correlation-Indicator" json:"chargingCorrelationIndicator,omitempty"`
	IPCANType                      *IpCanType                      `avp:"IP-CAN-Type" json:"iPCANType,omitempty"`
}

type ResourceAllocationNotification int

const (
	ResourceAllocationNotification_ENABLE_NOTIFICATION ResourceAllocationNotification = 0
)

type ChargingCorrelationIndicator int

const (
	ChargingCorrelationIndicator_CHARGING_IDENTIFIER_REQUIRED ChargingCorrelationIndicator = 0
)

/*
+---------------------------------------+-----------+---------+----+----+-----+----------+
| AVP                                   | Type      |  Layers | R8 | R9 | R10 | mutative |
+---------------------------------------+-----------+---------+----+----+-----+----------+
| Charging-Rule-Name                    | string    |         | 1  | 1  |     |          |
| Rating-Group                          | uint32    |         | 1  | 1  |     |          |
| Service-Identifier                    | uint32    |         | 1  | 1  |     |          |
| Flow-Information                      | struct    | 1       | 1  | 1  |     | yes      |
| Flow-Description                      | string    |         | 1  |    |     |          |
| Default-Bearer-Indication             | enumerate |         |    |    |     |          |
| TDF-Application-Identifier            | string    |         |    |    |     |          |
| Flow-Status                           | enumerate |         | 1  | 1  |     |          |
| QoS-Information                       | struct    | 2       | 1  | 1  |     | yes      |
| PS-to-CS-Session-Continuity           | enumerate |         |    |    |     |          |
| Reporting-Level                       | enumerate |         | 1  | 1  |     |          |
| Online                                | enumerate |         | 1  | 1  |     |          |
| Offline                               | enumerate |         | 1  | 1  |     |          |
| Max-PLR-DL                            | float32   |         |    |    |     |          |
| Max-PLR-UL                            | float32   |         |    |    |     |          |
| Metering-Method                       | enumerate |         | 1  | 1  |     |          |
| Precedence                            | uint32    |         | 1  | 1  |     |          |
| AF-Charging-Identifier                | string    |         | 1  | 1  |     |          |
| Flows                                 | struct    | 1       | 1  | 1  |     | yes      |
| Monitoring-Key                        | string    |         |    | 1  |     |          |
| Redirect-Information                  | struct    | 1       |    |    |     |          |
| Mute-Notification                     | enumerate |         |    |    |     |          |
| AF-Signalling-Protocol                | enumerate |         |    | 1  |     |          |
| Sponsor-Identity                      | string    |         |    |    |     |          |
| Application-Service-Provider-Identity | string    |         |    |    |     |          |
| Required-Access-Info                  | enumerate |         |    |    |     |          |
| Sharing-Key-DL                        | uint32    |         |    |    |     |          |
| Sharing-Key-UL                        | uint32    |         |    |    |     |          |
| Traffic-Steering-Policy-Identifier-DL | string    |         |    |    |     |          |
| Traffic-Steering-Policy-Identifier-UL | string    |         |    |    |     |          |
| Content-Version                       | uint64    |         |    |    |     |          |
| Calling-Party-Address                 | string    |         |    |    |     |          |
| Callee-Information                    | struct    | 1       |    |    |     |          |
+---------------------------------------+-----------+---------+----+----+-----+----------+
*/
type ChargingRuleDefinition struct {
	ChargingRuleName                   datatype.OctetString     `avp:"Charging-Rule-Name" json:"chargingRuleName"`
	RatingGroup                        *datatype.Unsigned32     `avp:"Rating-Group" json:"ratingGroup,omitempty"`
	ServiceIdentifier                  *datatype.Unsigned32     `avp:"Service-Identifier" json:"serviceIdentifier,omitempty"`
	FlowInformation                    []FlowInformation        `avp:"Flow-Information" json:"flowInformation,omitempty"`
	FlowDescription                    *datatype.IPFilterRule   `avp:"Flow-Description" json:"flowDescription,omitempty"`
	DefaultBearerIndication            *DefaultBearerIndication `avp:"Default-Bearer-Indication" json:"defaultBearerIndication,omitempty"`
	TDFApplicationIdentifier           *datatype.OctetString    `avp:"TDF-Application-Identifier" json:"tDFApplicationIdentifier,omitempty"`
	FlowStatus                         *FlowStatus              `avp:"Flow-Status" json:"flowStatus,omitempty"`
	QoSInformation                     *QoSInformation          `avp:"QoS-Information" json:"qoSInformation,omitempty"`
	PStoCSSessionContinuity            *PStoCSSessionContinuity `avp:"PS-to-CS-Session-Continuity" json:"pStoCSSessionContinuity,omitempty"`
	ReportingLevel                     *ReportingLevel          `avp:"Reporting-Level" json:"reportingLevel,omitempty"`
	Online                             *Online                  `avp:"Online" json:"online,omitempty"`
	Offline                            *Offline                 `avp:"Offline" json:"offline,omitempty"`
	MaxPLRDL                           *datatype.Float32        `avp:"Max-PLR-DL" json:"MaxPLRDL,omitempty"`
	MaxPLRUL                           *datatype.Float32        `avp:"Max-PLR-UL" json:"maxPLRUL,omitempty"`
	MeteringMethod                     *MeteringMethod          `avp:"Metering-Method" json:"meteringMethod,omitempty"`
	Precedence                         *datatype.Unsigned32     `avp:"Precedence" json:"precedence,omitempty" cmp:"ignore"`
	AFChargingIdentifier               *datatype.OctetString    `avp:"AF-Charging-Identifier" json:"aFChargingIdentifier,omitempty"`
	Flows                              []Flows                  `avp:"Flows" json:"flows,omitempty"`
	MonitoringKey                      *datatype.OctetString    `avp:"Monitoring-Key" json:"monitoringKey,omitempty"`
	RedirectInformation                *RedirectInformation     `avp:"Redirect-Information" json:"redirectInformation,omitempty"`
	MuteNotification                   *MuteNotification        `avp:"Mute-Notification" json:"muteNotification,omitempty"`
	AFSignallingProtocol               *AFSignallingProtocol    `avp:"AF-Signalling-Protocol" json:"aFSignallingProtocol,omitempty"`
	SponsorIdentity                    *datatype.UTF8String     `avp:"Sponsor-Identity" json:"sponsorIdentity,omitempty"`
	ApplicationServiceProviderIdentity *datatype.UTF8String     `avp:"Application-Service-Provider-Identity" json:"applicationServiceProviderIdentity,omitempty"`
	RequiredAccessInfo                 []RequiredAccessInfo     `avp:"Required-Access-Info" json:"requiredAccessInfo,omitempty"`
	SharingKeyDL                       *datatype.Unsigned32     `avp:"Sharing-Key-DL" json:"sharingKeyDL,omitempty"`
	SharingKeyUL                       *datatype.Unsigned32     `avp:"Sharing-Key-UL" json:"sharingKeyUL,omitempty"`
	TrafficSteeringPolicyIdentifierDL  *datatype.OctetString    `avp:"Traffic-Steering-Policy-Identifier-DL" json:"trafficSteeringPolicyIdentifierDL,omitempty"`
	TrafficSteeringPolicyIdentifierUL  *datatype.OctetString    `avp:"Traffic-Steering-Policy-Identifier-UL" json:"trafficSteeringPolicyIdentifierUL,omitempty"`
	ContentVersion                     *datatype.Unsigned64     `avp:"Content-Version" json:"contentVersion,omitempty"`
	CallingPartyAddress                []datatype.UTF8String    `avp:"Calling-Party-Address" json:"callingPartyAddress,omitempty"`
	CalleeInformation                  *CalleeInformation       `avp:"Callee-Information" json:"calleeInformation,omitempty"`
}

/*
+--------------------------+-----------+---------+----+----+-----+----------+
| AVP                      | Type      |  Layers | R8 | R9 | R10 | mutative |
+--------------------------+-----------+---------+----+----+-----+----------+
| Flow-Description         | string    |         | 1  | 1  |     |          |
| Packet-Filter-Identifier | string    |         | 1  | 1  |     |          |
| Packet-Filter-Usage      | int       |         |    | 1  |     |          |
| ToS-Traffic-Class        | string    |         | 1  | 1  |     |          |
| Security-Parameter-Index | string    |         | 1  | 1  |     |          |
| Flow-Label               | string    |         | 1  | 1  |     |          |
| Flow-Direction           | enumerate |         |    | 1  |     |          |
| Routing-Rule-Identifier  | string    |         |    |    |     |          |
+--------------------------+-----------+---------+----+----+-----+----------+
*/
type FlowInformation struct {
	FlowDescription        *datatype.IPFilterRule `avp:"Flow-Description" json:"flowDescription,omitempty"`
	PacketFilterIdentifier *datatype.OctetString  `avp:"Packet-Filter-Identifier" json:"packetFilterIdentifier,omitempty"`
	PacketFilterUsage      *PacketFilterUsage     `avp:"Packet-Filter-Usage" json:"packetFilterUsage,omitempty"`
	ToSTrafficClass        *datatype.OctetString  `avp:"ToS-Traffic-Class" json:"toSTrafficClass,omitempty"`
	SecurityParameterIndex *datatype.OctetString  `avp:"Security-Parameter-Index" json:"securityParameterIndex,omitempty"`
	FlowLabel              *datatype.OctetString  `avp:"Flow-Label" json:"flowLabel,omitempty"`
	FlowDirection          *FlowDirection         `avp:"Flow-Direction" json:"flowDirection,omitempty"`
	RoutingRuleIdentifier  *datatype.OctetString  `avp:"Routing-Rule-Identifier" json:"RoutingRuleIdentifier,omitempty"`
}

type DefaultBearerIndication int

const (
	DefaultBearerIndication_BIND_TO_DEF_BEARER        DefaultBearerIndication = 0
	DefaultBearerIndication_BIND_TO_APPLICABLE_BEARER DefaultBearerIndication = 1
)

type FlowStatus int

const (
	FlowStatus_ENABLED_UPLINK   FlowStatus = 0
	FlowStatus_ENABLED_DOWNLINK FlowStatus = 1
	FlowStatus_ENABLED          FlowStatus = 2
	FlowStatus_DISABLED         FlowStatus = 3
	FlowStatus_REMOVED          FlowStatus = 4
)

type PacketFilterUsage int

const (
	PacketFilterUsage_SEND_TO_UE PacketFilterUsage = 1
)

type PStoCSSessionContinuity int

const (
	PStoCSSessionContinuity_VIDEO_PS2CS_CONT_CANDIDATE PacketFilterUsage = 0
)

type ReportingLevel int

const (
	ReportingLevel_SERVICE_IDENTIFIER_LEVEL     ReportingLevel = 0
	ReportingLevel_RATING_GROUP_LEVEL           ReportingLevel = 1
	ReportingLevel_SPONSORED_CONNECTIVITY_LEVEL ReportingLevel = 2
)

type MeteringMethod int

const (
	MeteringMethod_DURATION        MeteringMethod = 0
	MeteringMethod_VOLUME          MeteringMethod = 1
	MeteringMethod_DURATION_VOLUME MeteringMethod = 2
	MeteringMethod_EVENT           MeteringMethod = 3
)

type MuteNotification int

const (
	MuteNotification_MUTE_REQUIRED MuteNotification = 0
)

type AFSignallingProtocol int

const (
	AFSignallingProtocol_NO_INFORMATION AFSignallingProtocol = 0
)

type RequiredAccessInfo int

const (
	RequiredAccessInfo_USER_LOCATION RequiredAccessInfo = 0
	RequiredAccessInfo_MS_TIME_ZONE  RequiredAccessInfo = 1
)

// TS29.214
/*
+------------------------+-----------+---------+----+----+-----+----------+
| AVP                    | Type      |  Layers | R8 | R9 | R10 | mutative |
+------------------------+-----------+---------+----+----+-----+----------+
| Media-Component-Number | uint32    |         | 1  | 1  |     |          |
| Flow-Number            | uint32    |         | 1  | 1  |     |          |
| Content-Version        | uint64    |         |    |    |     |          |
| Final-Unit-Action      | enumerate |         | 1  | 1  |     |          |
| Media-Component-Status | uint32    |         |    |    |     |          |
+------------------------+-----------+---------+----+----+-----+----------+
*/
type Flows struct {
	MediaComponentNumber datatype.Unsigned32   `avp:"Media-Component-Number" json:"mediaComponentNumber"`
	FlowNumber           []datatype.Unsigned32 `avp:"Flow-Number" json:"flowNumber,omitempty"`
	ContentVersion       []datatype.Unsigned64 `avp:"Content-Version" json:"contentVersion,omitempty"`
	FinalUnitAction      *FinalUnitAction      `avp:"Final-Unit-Action" json:"finalUnitAction,omitempty"`
	MediaComponentStatus *datatype.Unsigned32  `avp:"Media-Component-Status" json:"mediaComponentStatus,omitempty"`
}

type CalleeInformation struct {
	CalledPartyAddress     *datatype.UTF8String  `avp:"Called-Party-Address" json:"calledPartyAddress,omitempty"`
	RequestedPartyAddress  []datatype.UTF8String `avp:"Requested-Party-Address" json:"requestedPartyAddress,omitempty"`
	CalledAssertedIdentity []datatype.UTF8String `avp:"Called-Asserted-Identity" json:"calledAssertedIdentity,omitempty"`
}

type RedirectInformation struct {
	RedirectSupport       *RedirectSupport     `avp:"Redirect-Support" json:"redirectSupport,omitempty"`
	RedirectAddressType   *RedirectAddressType `avp:"Redirect-Address-Type" json:"redirectAddressType,omitempty"`
	RedirectServerAddress *datatype.UTF8String `avp:"Redirect-Server-Address" json:"redirectServerAddress,omitempty"`
}

type RedirectSupport datatype.Enumerated

const (
	RedirectSupport_REDIRECTION_DISABLED RedirectSupport = 0
	RedirectSupport_REDIRECTION_ENABLED  RedirectSupport = 1
)

type RedirectAddressType datatype.Enumerated

const (
	RedirectAddressType_IPv4_Address RedirectAddressType = 0
	RedirectAddressType_IPv6_Address RedirectAddressType = 1
	RedirectAddressType_URL          RedirectAddressType = 2
	RedirectAddressType_SIP_URI      RedirectAddressType = 3
)

type DefaultEpsBearerQos struct {
	QoSClassIdentifier          *QoSClassIdentifier          `avp:"QoS-Class-Identifier" json:"qosClassIdentifier,omitempty"`
	AllocationRetentionPriority *AllocationRetentionPriority `avp:"Allocation-Retention-Priority" json:"allocationRetentionPriority,omitempty"`
}

type AllocationRetentionPriority struct {
	PriorityLevel           datatype.Unsigned32      `avp:"Priority-Level" json:"priorityLevel"`
	PreEmptionCapability    *PreEmptionCapability    `avp:"Pre-emption-Capability" json:"preEmptionCapability,omitempty"`
	PreEmptionVulnerability *PreEmptionVulnerability `avp:"Pre-emption-Vulnerability" json:"preEmptionVulnerability,omitempty"`
}

type PreEmptionCapability datatype.Enumerated

const (
	PRE_EMPTION_CAPABILITY_ENABLED  PreEmptionCapability = 0
	PRE_EMPTION_CAPABILITY_DISABLED PreEmptionCapability = 1
)

type PreEmptionVulnerability datatype.Enumerated

const (
	PRE_EMPTION_VULNERABILITY_ENABLED  PreEmptionVulnerability = 0
	PRE_EMPTION_VULNERABILITY_DISABLED PreEmptionVulnerability = 1
)

type RemovalOfAccess datatype.Enumerated

const (
	REMOVAL_OF_ACCESS RemovalOfAccess = 0
)

type CsgInformationReporting datatype.Enumerated

const (
	CsgInformationReporting_CHANGE_CSG_CELL                     CsgInformationReporting = 0
	CsgInformationReporting_CHANGE_CSG_SUBSCRIBED_HYBRID_CELL   CsgInformationReporting = 1
	CsgInformationReporting_CHANGE_CSG_UNSUBSCRIBED_HYBRID_CELL CsgInformationReporting = 2
)

type PraInstall struct {
	PresenceReportingAreaInformation []PresenceReportingAreaInformation `avp:"Presence-Reporting-Area-Information" json:"presenceReportingAreaInformation,omitempty"`
}

type PresenceReportingAreaStatus datatype.Enumerated

const (
	PresenceReportingAreaStatus_In_Area     PresenceReportingAreaStatus = 0
	PresenceReportingAreaStatus_Out_of_Area PresenceReportingAreaStatus = 1
	PresenceReportingAreaStatus_Inactive    PresenceReportingAreaStatus = 2
)

type PraRemove struct {
	PresenceReportingAreaIdentifier []datatype.OctetString `avp:"Presence-Reporting-Area-Identifier" json:"presenceReportingAreaIdentifier,omitempty"`
}

type ProxyInfo struct {
	ProxyHost  datatype.DiameterIdentity `avp:"Proxy-Host" json:"proxyHost"`
	ProxyState datatype.OctetString      `avp:"Proxy-State" json:"proxyState"`
}

type GrantedServiceUnit struct {
	TariffTimeChange       *datatype.Time       `avp:"Tariff-Time-Change" json:"tariffTimeChange,omitempty"`
	CcTime                 *datatype.Unsigned32 `avp:"CC-Time" json:"ccTime,omitempty"`
	CcMoney                *CcMoney             `avp:"CC-Money" json:"ccMoney,omitempty"`
	CcTotalOctets          *datatype.Unsigned64 `avp:"CC-Total-Octets" json:"ccTotalOctets,omitempty"`
	CcInputOctets          *datatype.Unsigned64 `avp:"CC-Input-Octets" json:"ccInputOctets,omitempty"`
	CcOutputOctets         *datatype.Unsigned64 `avp:"CC-Output-Octets" json:"ccOutputOctets,omitempty"`
	CcServiceSpecificUnits *datatype.Unsigned64 `avp:"CC-Service-Specific-Units" json:"ccServiceSpecificUnits,omitempty"`
}

type CcMoney struct {
	UnitValue    UnitValue            `avp:"Unit-Value" json:"unitValue"`
	CurrencyCode *datatype.Unsigned32 `avp:"Currency-Code" json:"currencyCode,omitempty"`
}

type UnitValue struct {
	ValueDigits *datatype.Integer64 `avp:"Value-Digits" json:"valueDigits,omitempty"`
	Exponent    *datatype.Integer32 `avp:"Exponent" json:"Exponent,omitempty"`
}

type ConditionalPolicyInformation struct {
	ExecutionTime                     *datatype.Time                      `avp:"Execution-Time" json:"executionTime,omitempty"`
	DefaultEpsBearerQoS               *DefaultEpsBearerQos                `avp:"Default-EPS-Bearer-QoS" json:"defaultEpsBearerQoS,omitempty"`
	ApnAggregateMaxBitrateUL          *datatype.Unsigned32                `avp:"APN-Aggregate-Max-Bitrate-UL" json:"apnAggregateMaxBitrateUL,omitempty"`
	ApnAggregateMaxBitrateDL          *datatype.Unsigned32                `avp:"APN-Aggregate-Max-Bitrate-DL" json:"apnAggregateMaxBitrateDL,omitempty"`
	ExtendedApnAMbrUL                 *datatype.Unsigned32                `avp:"Extended-APN-AMBR-UL" json:"extendedApnAMbrUL,omitempty"`
	ExtendedApnAMbrDL                 *datatype.Unsigned32                `avp:"Extended-APN-AMBR-DL" json:"extendedApnAMbrDL,omitempty"`
	ConditionalApnAggregateMaxBitrate []ConditionalAPNAggregateMaxBitrate `avp:"Conditional-APN-Aggregate-Max-Bitrate" json:"conditionalApnAggregateMaxBitrate,omitempty"`
}

type SubscriptionId struct {
	SubscriptionIdType SubscriptionIdType  `avp:"Subscription-Id-Type" json:"subscriptionIdType"`
	SubscriptionIdData datatype.UTF8String `avp:"Subscription-Id-Data" json:"subscriptionIdData"`
}

type SubscriptionIdType datatype.Enumerated

const (
	SubscriptionIdType_END_USER_E164    SubscriptionIdType = 0
	SubscriptionIdType_END_USER_IMSI    SubscriptionIdType = 1
	SubscriptionIdType_END_USER_SIP_URI SubscriptionIdType = 2
	SubscriptionIdType_END_USER_NAI     SubscriptionIdType = 3
	SubscriptionIdType_END_USER_PRIVATE SubscriptionIdType = 4
)

type SupportedFeatures struct {
	VendorId      datatype.Unsigned32 `avp:"Vendor-Id" json:"vendorId"`
	FeatureListID datatype.Unsigned32 `avp:"Feature-List-ID" json:"featureListID"`
	FeatureList   datatype.Unsigned32 `avp:"Feature-List" json:"featureList"`
}

type TDFInformation struct {
	TDFDestinationRealm *datatype.DiameterIdentity `avp:"TDF-Destination-Realm" json:"tdfDestinationRealm,omitempty"`
	TDFDestinationHost  *datatype.DiameterIdentity `avp:"TDF-Destination-Host" json:"tdfDestinationHost,omitempty"`
	TDFIPAddress        *datatype.Address          `avp:"TDF-IP-Address" json:"tdfIPAddress,omitempty"`
}

type NetworkRequestSupport datatype.Enumerated

const (
	NetworkRequestSupport_NETWORK_REQUEST_NOT_SUPPORTED NetworkRequestSupport = 0
	NetworkRequestSupport_NETWORK_REQUEST_SUPPORTED     NetworkRequestSupport = 1
)

type PacketFilterOperation datatype.Enumerated

const (
	PacketFilterOperation_DELETION     PacketFilterOperation = 0
	PacketFilterOperation_ADDITION     PacketFilterOperation = 1
	PacketFilterOperation_MODIFICATION PacketFilterOperation = 2
)

/*
+--------------------------+-----------+--------+----+----+-----+
| AVP                      | Type      | Layers | R8 | R9 | R10 |
+--------------------------+-----------+--------+----+----+-----+
| Packet-Filter-Identifier | string    |        | 1  | 1  |     |
| Precedence               | uint32    |        | 1  | 1  |     |
| Packet-Filter-Content    | string    |        | 1  | 1  |     |
| ToS-Traffic-Class        | string    |        | 1  | 1  |     |
| Security-Parameter-Index | string    |        | 1  | 1  |     |
| Flow-Label               | string    |        | 1  | 1  |     |
| Flow-Direction           | enumerate |        |    | 1  |     |
+--------------------------+-----------+--------+----+----+-----+
*/

type PacketFilterInformation struct {
	PacketFilterIdentifier *datatype.OctetString  `avp:"Packet-Filter-Identifier" json:"packetFilterIdentifier,omitempty"`
	Precedence             *datatype.Unsigned32   `avp:"Precedence" json:"precedence,omitempty"`
	PacketFilterContent    *datatype.IPFilterRule `avp:"Packet-Filter-Content" json:"packetFilterContent,omitempty"`
	ToSTrafficClass        *datatype.OctetString  `avp:"ToS-Traffic-Class" json:"toSTrafficClass,omitempty"`
	SecurityParameterIndex *datatype.OctetString  `avp:"Security-Parameter-Index" json:"securityParameterIndex,omitempty"`
	FlowLabel              *datatype.OctetString  `avp:"Flow-Label" json:"flowLabel,omitempty"`
	FlowDirection          *FlowDirection         `avp:"Flow-Direction" json:"flowDirection,omitempty"`
}

type FlowDirection datatype.Enumerated

const (
	FlowDirection_UNSPECIFIED   FlowDirection = 0
	FlowDirection_DOWNLINK      FlowDirection = 1
	FlowDirection_UPLINK        FlowDirection = 2
	FlowDirection_BIDIRECTIONAL FlowDirection = 3
)

type BearerOperation datatype.Enumerated

const (
	BearerOperation_TERMINATION   BearerOperation = 0
	BearerOperation_ESTABLISHMENT BearerOperation = 1
	BearerOperation_MODIFICATION  BearerOperation = 2
)

type DynamicAddressFlag datatype.Enumerated

const (
	DynamicAddressFlag_Static  DynamicAddressFlag = 0
	DynamicAddressFlag_Dynamic DynamicAddressFlag = 1
)

type DynamicAddressFlagExtension datatype.Enumerated

const (
	DynamicAddressFlagExtension_Static  DynamicAddressFlagExtension = 0
	DynamicAddressFlagExtension_Dynamic DynamicAddressFlagExtension = 1
)

type ANTrusted datatype.Enumerated

const (
	ANTrusted_TRUSTED   ANTrusted = 0
	ANTrusted_UNTRUSTED ANTrusted = 1
)

type TerminationCause datatype.Enumerated

const (
	TerminationCause_DIAMETER_LOGOUT               TerminationCause = 1
	TerminationCause_DIAMETER_SERVICE_NOT_PROVIDED TerminationCause = 2
	TerminationCause_DIAMETER_BAD_ANSWER           TerminationCause = 3
	TerminationCause_DIAMETER_ADMINISTRATIVE       TerminationCause = 4
	TerminationCause_DIAMETER_LINK_BROKEN          TerminationCause = 5
	TerminationCause_DIAMETER_AUTH_EXPIRED         TerminationCause = 6
	TerminationCause_DIAMETER_USER_MOVED           TerminationCause = 7
	TerminationCause_DIAMETER_SESSION_TIMEOUT      TerminationCause = 8
)

type UserEquipmentInfo struct {
	UserEquipmentInfoType  UserEquipmentInfoType `avp:"User-Equipment-Info-Type" json:"userEquipmentInfoType"`
	UserEquipmentInfoValue datatype.OctetString  `avp:"User-Equipment-Info-Value" json:"userEquipmentInfoValue"`
}

type UserEquipmentInfoType datatype.Enumerated

const (
	UserEquipmentInfoType_IMEISV         UserEquipmentInfoType = 0
	UserEquipmentInfoType_MAC            UserEquipmentInfoType = 1
	UserEquipmentInfoType_EUI64          UserEquipmentInfoType = 2
	UserEquipmentInfoType_MODIFIED_EUI64 UserEquipmentInfoType = 3
)

/*
+---------------------------------------+-----------+--------+----+----+-----+----------+
| AVP                                   | Type      | Layers | R8 | R9 | R10 | mutative |
+---------------------------------------+-----------+--------+----+----+-----+----------+
| QoS-Class-Identifier                  | enumerate |        | 1  | 1  |     |          |
| Max-Requested-Bandwidth-UL            | uint32    |        | 1  | 1  |     |          |
| Max-Requested-Bandwidth-DL            | uint32    |        | 1  | 1  |     |          |
| Extended-Max-Requested-BW-UL          | uint32    |        |    |    |     |          |
| Extended-Max-Requested-BW-DL          | uint32    |        |    |    |     |          |
| Guaranteed-Bitrate-UL                 | uint32    |        | 1  | 1  |     |          |
| Guaranteed-Bitrate-DL                 | uint32    |        | 1  | 1  |     |          |
| Extended-GBR-UL                       | uint32    |        |    |    |     |          |
| Extended-GBR-DL                       | uint32    |        |    |    |     |          |
| Bearer-Identifier                     | string    |        | 1  | 1  |     |          |
| Allocation-Retention-Priority         | struct    | 1      | 1  | 1  |     | no       |
| APN-Aggregate-Max-Bitrate-UL          | uint32    |        | 1  | 1  |     |          |
| APN-Aggregate-Max-Bitrate-DL          | uint32    |        | 1  | 1  |     |          |
| Extended-APN-AMBR-UL                  | uint32    |        |    |    |     |          |
| Extended-APN-AMBR-DL                  | uint32    |        |    |    |     |          |
| Conditional-APN-Aggregate-Max-Bitrate | struct    | 1      |    |    |     |          |
+---------------------------------------+-----------+--------+----+----+-----+----------+
*/
type QoSInformation struct {
	QoSClassIdentifier                *QoSClassIdentifier                 `avp:"QoS-Class-Identifier" json:"qoSClassIdentifier,omitempty"`
	MaxRequestedBandwidthUL           *datatype.Unsigned32                `avp:"Max-Requested-Bandwidth-UL" json:"maxRequestedBandwidthUL,omitempty"`
	MaxRequestedBandwidthDL           *datatype.Unsigned32                `avp:"Max-Requested-Bandwidth-DL" json:"maxRequestedBandwidthDL,omitempty"`
	ExtendedMaxRequestedBWUL          *datatype.Unsigned32                `avp:"Extended-Max-Requested-BW-UL" json:"extendedMaxRequestedBWUL,omitempty"`
	ExtendedMaxRequestedBWDL          *datatype.Unsigned32                `avp:"Extended-Max-Requested-BW-DL" json:"extendedMaxRequestedBWDL,omitempty"`
	GuaranteedBitrateUL               *datatype.Unsigned32                `avp:"Guaranteed-Bitrate-UL" json:"guaranteedBitrateUL,omitempty"`
	GuaranteedBitrateDL               *datatype.Unsigned32                `avp:"Guaranteed-Bitrate-DL" json:"guaranteedBitrateDL,omitempty"`
	ExtendedGBRUL                     *datatype.Unsigned32                `avp:"Extended-GBR-UL" json:"extendedGBRUL,omitempty"`
	ExtendedGBRDL                     *datatype.Unsigned32                `avp:"Extended-GBR-DL" json:"extendedGBRDL,omitempty"`
	BearerIdentifier                  *datatype.OctetString               `avp:"Bearer-Identifier" json:"bearerIdentifier,omitempty"`
	AllocationRetentionPriority       *AllocationRetentionPriority        `avp:"Allocation-Retention-Priority" json:"allocationRetentionPriority,omitempty"`
	APNAggregateMaxBitrateUL          *datatype.Unsigned32                `avp:"APN-Aggregate-Max-Bitrate-UL" json:"aPNAggregateMaxBitrateUL,omitempty"`
	APNAggregateMaxBitrateDL          *datatype.Unsigned32                `avp:"APN-Aggregate-Max-Bitrate-DL" json:"aPNAggregateMaxBitrateDL,omitempty"`
	ExtendedAPNAMBRUL                 *datatype.Unsigned32                `avp:"Extended-APN-AMBR-UL" json:"extendedAPNAMBRUL,omitempty"`
	ExtendedAPNAMBRDL                 *datatype.Unsigned32                `avp:"Extended-APN-AMBR-DL" json:"extendedAPNAMBRDL,omitempty"`
	ConditionalAPNAggregateMaxBitrate []ConditionalAPNAggregateMaxBitrate `avp:"Conditional-APN-Aggregate-Max-Bitrate" json:"conditionalAPNAggregateMaxBitrate,omitempty"`
}

type QoSClassIdentifier datatype.Enumerated

const (
	Reserved                  QoSClassIdentifier = 0
	QoSClassIdentifier_QCI_1  QoSClassIdentifier = 1
	QoSClassIdentifier_QCI_2  QoSClassIdentifier = 2
	QoSClassIdentifier_QCI_3  QoSClassIdentifier = 3
	QoSClassIdentifier_QCI_4  QoSClassIdentifier = 4
	QoSClassIdentifier_QCI_5  QoSClassIdentifier = 5
	QoSClassIdentifier_QCI_6  QoSClassIdentifier = 6
	QoSClassIdentifier_QCI_7  QoSClassIdentifier = 7
	QoSClassIdentifier_QCI_8  QoSClassIdentifier = 8
	QoSClassIdentifier_QCI_9  QoSClassIdentifier = 9
	QoSClassIdentifier_QCI_65 QoSClassIdentifier = 65
	QoSClassIdentifier_QCI_66 QoSClassIdentifier = 66
	QoSClassIdentifier_QCI_67 QoSClassIdentifier = 67
	QoSClassIdentifier_QCI_69 QoSClassIdentifier = 69
	QoSClassIdentifier_QCI_70 QoSClassIdentifier = 70
	QoSClassIdentifier_QCI_71 QoSClassIdentifier = 71
	QoSClassIdentifier_QCI_72 QoSClassIdentifier = 72
	QoSClassIdentifier_QCI_73 QoSClassIdentifier = 73
	QoSClassIdentifier_QCI_74 QoSClassIdentifier = 74
	QoSClassIdentifier_QCI_75 QoSClassIdentifier = 75
	QoSClassIdentifier_QCI_76 QoSClassIdentifier = 76
	QoSClassIdentifier_QCI_77 QoSClassIdentifier = 77
	QoSClassIdentifier_QCI_79 QoSClassIdentifier = 79
	QoSClassIdentifier_QCI_80 QoSClassIdentifier = 80
	QoSClassIdentifier_QCI_82 QoSClassIdentifier = 82
	QoSClassIdentifier_QCI_83 QoSClassIdentifier = 83
	QoSClassIdentifier_QCI_84 QoSClassIdentifier = 84
	QoSClassIdentifier_QCI_85 QoSClassIdentifier = 85
)

type PreemptionCapability datatype.Enumerated

const (
	PreemptionCapability_PRE_EMPTION_CAPABILITY_ENABLED  PreemptionCapability = 0
	PreemptionCapability_PRE_EMPTION_CAPABILITY_DISABLED PreemptionCapability = 1
)

type PreemptionVulnerability datatype.Enumerated

const (
	PreemptionVulnerability_PRE_EMPTION_VULNERABILITY_ENABLED  PreemptionVulnerability = 0
	PreemptionVulnerability_PRE_EMPTION_VULNERABILITY_DISABLED PreemptionVulnerability = 1
)

type ConditionalAPNAggregateMaxBitrate struct {
	APNAggregateMaxBitrateUL *datatype.Unsigned32 `avp:"APN-Aggregate-Max-Bitrate-UL" json:"apnAggregateMaxBitrateUL,omitempty"`
	APNAggregateMaxBitrateDL *datatype.Unsigned32 `avp:"APN-Aggregate-Max-Bitrate-DL" json:"apnAggregateMaxBitrateDL,omitempty"`
	ExtendedAPNAMBRUL        *datatype.Unsigned32 `avp:"Extended-APN-AMBR-UL" json:"extendedAPNAMBRUL,omitempty"`
	ExtendedAPNAMBRDL        *datatype.Unsigned32 `avp:"Extended-APN-AMBR-DL" json:"extendedAPNAMBRDL,omitempty"`
	IPCANType                []IpCanType          `avp:"IP-CAN-Type" json:"ipCANType,omitempty"`
	RATType                  []RatType            `avp:"RAT-Type" json:"ratType,omitempty"`
}

type QoSNegotiation datatype.Enumerated

const (
	QoSNegotiation_NO_QoS_NEGOTIATION        QoSNegotiation = 0
	QoSNegotiation_QoS_NEGOTIATION_SUPPORTED QoSNegotiation = 1
)

type QoSUpgrade datatype.Enumerated

const (
	QoSUpgrade_QoS_UPGRADE_NOT_SUPPORTED QoSUpgrade = 0
	QoSUpgrade_QoS_QoS_UPGRADE_SUPPORTED QoSUpgrade = 1
)

type DefaultQoSInformation struct {
	QoSClassIdentifier      *QoSClassIdentifier  `avp:"QoS-Class-Identifier" json:"qoSClassIdentifier,omitempty"`
	MaxRequestedBandwidthUL *datatype.Unsigned32 `avp:"Max-Requested-Bandwidth-UL" json:"maxRequestedBandwidthUL,omitempty"`
	MaxRequestedBandwidthDL *datatype.Unsigned32 `avp:"Max-Requested-Bandwidth-DL" json:"maxRequestedBandwidthDL,omitempty"`
	DefaultQoSName          *datatype.UTF8String `avp:"Default-QoS-Name" json:"defaultQoSName,omitempty"`
}

type ANGWStatus datatype.Enumerated

const (
	ANGWStatus_AN_GW_FAILED ANGWStatus = 0
)

type FixedUserLocationInfo struct {
	SSID             *datatype.UTF8String  `avp:"SSID" json:"SSID,omitempty"`
	BSSID            *datatype.UTF8String  `avp:"BSSID" json:"BSSID,omitempty"`
	LogicalAccessID  *datatype.OctetString `avp:"Logical-Access-ID" json:"LogicalAccessID,omitempty"`
	PhysicalAccessID *datatype.UTF8String  `avp:"Physical-Access-ID" json:"PhysicalAccessID,omitempty"`
}

type BearerUsage datatype.Enumerated

const (
	BearerUsage_GENERAL        BearerUsage = 0
	BearerUsage_IMS_SIGNALLING BearerUsage = 1
)

type Online datatype.Enumerated

const (
	Online_DISABLE_ONLINE Online = 0
	Online_ENABLE_ONLINE  Online = 1
)

type Offline datatype.Enumerated

const (
	Offline_DISABLE_OFFLINE Offline = 0
	Offline_ENABLE_OFFLINE  Offline = 1
)

/*
+--------------------------+-----------+--------+----+----+-----+
| AVP                      | Type      | Layers | R8 | R9 | R10 |
+--------------------------+-----------+--------+----+----+-----+
| Precedence               | uint32    |        | 1  | 1  |     |
| TFT-Filter               | string    |        | 1  | 1  |     |
| ToS-Traffic-Class        | string    |        | 1  | 1  |     |
| Security-Parameter-Index | string    |        | 1  | 1  |     |
| Flow-Label               | string    |        | 1  | 1  |     |
| Flow-Direction           | enumerate |        |    | 1  |     |
+--------------------------+-----------+--------+----+----+-----+
*/
type TFTPacketFilterInformation struct {
	Precedence             *datatype.Unsigned32   `avp:"Precedence" json:"precedence,omitempty"`
	TFTFilter              *datatype.IPFilterRule `avp:"TFT-Filter" json:"tftFilter,omitempty"`
	ToSTrafficClass        *datatype.OctetString  `avp:"ToS-Traffic-Class" json:"toSTrafficClass,omitempty"`
	SecurityParameterIndex *datatype.OctetString  `avp:"Security-Parameter-Index" json:"securityParameterIndex,omitempty"`
	FlowLabel              *datatype.OctetString  `avp:"Flow-Label" json:"flowLabel,omitempty"`
	FlowDirection          *FlowDirection         `avp:"Flow-Direction" json:"flowDirection,omitempty"`
}

/*
+-------------------------+-----------+---------+----+----+-----+----------+
| AVP                     | Type      |  Layers | R8 | R9 | R10 | mutative |
+-------------------------+-----------+---------+----+----+-----+----------+
| Charging-Rule-Name      | string    |         | 1  | 1  |     |          |
| Charging-Rule-Base-Name | string    |         | 1  | 1  |     |          |
| Bearer-Identifier       | string    |         | 1  | 1  |     |          |
| PCC-Rule-Status         | enumerate |         | 1  | 1  |     |          |
| Rule-Failure-Code       | enumerate |         | 1  | 1  |     |          |
| Final-Unit-Indication   | struct    | 2       | 1  | 1  |     | no       |
| RAN-NAS-Release-Cause   | string    |         |    |    |     |          |
| Content-Version         | uint64    |         |    |    |     |          |
+-------------------------+-----------+---------+----+----+-----+----------+
*/
type ChargingRuleReport struct {
	ChargingRuleName     []datatype.OctetString `avp:"Charging-Rule-Name" json:"chargingRuleName,omitempty"`
	ChargingRuleBaseName []datatype.UTF8String  `avp:"Charging-Rule-Base-Name" json:"chargingRuleBaseName,omitempty"`
	BearerIdentifier     *datatype.OctetString  `avp:"Bearer-Identifier" json:"bearerIdentifier,omitempty"`
	PCCRuleStatus        *PCCRuleStatus         `avp:"PCC-Rule-Status" json:"pccRuleStatus,omitempty"`
	RuleFailureCode      *RuleFailureCode       `avp:"Rule-Failure-Code" json:"ruleFailureCode,omitempty"`
	FinalUnitIndication  *FinalUnitIndication   `avp:"Final-Unit-Indication" json:"finalUnitIndication,omitempty"`
	RANNASReleaseCause   []datatype.OctetString `avp:"RAN-NAS-Release-Cause" json:"ranNASReleaseCause,omitempty"`
	ContentVersion       []datatype.Unsigned64  `avp:"Content-Version" json:"contentVersion,omitempty"`
}

type PCCRuleStatus datatype.Enumerated

const (
	PCCRuleStatus_ACTIVE               PCCRuleStatus = 0
	PCCRuleStatus_INACTIVE             PCCRuleStatus = 1
	PCCRuleStatus_TEMPORARILY_INACTIVE PCCRuleStatus = 2
)

type RuleFailureCode datatype.Enumerated

const (
	RuleFailureCode_UNKNOWN_RULE_NAME                  RuleFailureCode = 1
	RuleFailureCode_RATING_GROUP_ERROR                 RuleFailureCode = 2
	RuleFailureCode_SERVICE_IDENTIFIER_ERROR           RuleFailureCode = 3
	RuleFailureCode_GW_PCEF_MALFUNCTION                RuleFailureCode = 4
	RuleFailureCode_RESOURCES_LIMITATION               RuleFailureCode = 5
	RuleFailureCode_MAX_NR_BEARERS_REACHED             RuleFailureCode = 6
	RuleFailureCode_UNKNOWN_BEARER_ID                  RuleFailureCode = 7
	RuleFailureCode_MISSING_BEARER_ID                  RuleFailureCode = 8
	RuleFailureCode_MISSING_FLOW_INFORMATION           RuleFailureCode = 9
	RuleFailureCode_RESOURCE_ALLOCATION_FAILURE        RuleFailureCode = 10
	RuleFailureCode_UNSUCCESSFUL_QOS_VALIDATION        RuleFailureCode = 11
	RuleFailureCode_INCORRECT_FLOW_INFORMATION         RuleFailureCode = 12
	RuleFailureCode_PS_TO_CS_HANDOVER                  RuleFailureCode = 13
	RuleFailureCode_TDF_APPLICATION_IDENTIFIER_ERROR   RuleFailureCode = 14
	RuleFailureCode_NO_BEARER_BOUND                    RuleFailureCode = 15
	RuleFailureCode_FILTER_RESTRICTIONS                RuleFailureCode = 16
	RuleFailureCode_AN_GW_FAILED                       RuleFailureCode = 17
	RuleFailureCode_MISSING_REDIRECT_SERVER_ADDRESS    RuleFailureCode = 18
	RuleFailureCode_CM_END_USER_SERVICE_DENIED         RuleFailureCode = 19
	RuleFailureCode_CM_CREDIT_CONTROL_NOT_APPLICABLE   RuleFailureCode = 20
	RuleFailureCode_CM_AUTHORIZATION_REJECTED          RuleFailureCode = 21
	RuleFailureCode_CM_USER_UNKNOWN                    RuleFailureCode = 22
	RuleFailureCode_CM_RATING_FAILED                   RuleFailureCode = 23
	RuleFailureCode_ROUTING_RULE_REJECTION             RuleFailureCode = 24
	RuleFailureCode_UNKNOWN_ROUTING_ACCESS_INFORMATION RuleFailureCode = 25
	RuleFailureCode_NO_NBIFOM_SUPPORT                  RuleFailureCode = 26
	RuleFailureCode_UE_STATE_SUSPEND                   RuleFailureCode = 27
)

type ApplicationDetectionInformation struct {
	TDFApplicationIdentifier         datatype.OctetString  `avp:"TDF-Application-Identifier" json:"tdfApplicationIdentifier"`
	TDFApplicationInstanceIdentifier *datatype.OctetString `avp:"TDF-Application-Instance-Identifier" json:"tdfApplicationInstanceIdentifier,omitempty"`
	FlowInformation                  []FlowInformation     `avp:"Flow-Information" json:"flowInformation,omitempty"`
}

/*
+------------------------------------------+-----------+---------+----+----+-----+----------+
| AVP                                      | Type      |  Layers | R8 | R9 | R10 | mutative |
+------------------------------------------+-----------+---------+----+----+-----+----------+
| Access-Network-Charging-Identifier-Value | string    |         | 1  | 1  |     |          |
| Charging-Rule-Base-Name                  | string    |         | 1  | 1  |     |          |
| Charging-Rule-Name                       | string    |         | 1  | 1  |     |          |
| IP-CAN-Session-Charging-Scope            | enumerate |         |    |    |     |          |
+------------------------------------------+-----------+---------+----+----+-----+----------+
*/
type AccessNetworkChargingIdentifierGx struct {
	AccessNetworkChargingIdentifierValue datatype.OctetString       `avp:"Access-Network-Charging-Identifier-Value" json:"accessNetworkChargingIdentifierValue"`
	ChargingRuleBaseName                 []datatype.UTF8String      `avp:"Charging-Rule-Base-Name" json:"chargingRuleBaseName,omitempty"`
	ChargingRuleName                     []datatype.OctetString     `avp:"Charging-Rule-Name" json:"chargingRuleName,omitempty"`
	IPCANSessionChargingScope            *IPCANSessionChargingScope `avp:"IP-CAN-Session-Charging-Scope" json:"ipCANSessionChargingScope,omitempty"`
}

type IPCANSessionChargingScope datatype.Enumerated

const (
	IPCANSessionChargingScope_IP_CAN_SESSION_SCOPE IPCANSessionChargingScope = 0
)

type CoAInformation struct {
	TunnelInformation TunnelInformation `avp:"Tunnel-Information" json:"tunnelInformation"`
	CoAIPAddress      datatype.Address  `avp:"CoA-IP-Address" json:"coAIPAddress"`
}

type TunnelInformation struct {
	TunnelHeaderLength *datatype.Unsigned32    `avp:"Tunnel-Header-Length" json:"tunnelHeaderLength,omitempty"`
	TunnelHeaderFilter []datatype.IPFilterRule `avp:"Tunnel-Header-Filter" json:"tunnelHeaderFilter,omitempty"`
}

type UsageMonitoringInformation struct {
	MonitoringKey          *datatype.OctetString   `avp:"Monitoring-Key" json:"monitoringKey,omitempty"`
	GrantedServiceUnit     []GrantedServiceUnit    `avp:"Granted-Service-Unit" json:"grantedServiceUnit,omitempty"`
	UsedServiceUnit        []UsedServiceUnit       `avp:"Used-Service-Unit" json:"usedServiceUnit,omitempty"`
	QuotaConsumptionTime   *datatype.Unsigned32    `avp:"Quota-Consumption-Time" json:"quotaConsumptionTime,omitempty"`
	UsageMonitoringLevel   *UsageMonitoringLevel   `avp:"Usage-Monitoring-Level" json:"usageMonitoringLevel,omitempty"`
	UsageMonitoringReport  *UsageMonitoringReport  `avp:"Usage-Monitoring-Report" json:"usageMonitoringReport,omitempty"`
	UsageMonitoringSupport *UsageMonitoringSupport `avp:"Usage-Monitoring-Support" json:"usageMonitoringSupport,omitempty"`
}

type UsedServiceUnit struct {
	TariffChangeUsage      *TariffChangeUsage   `avp:"Tariff-Change-Usage" json:"tariffChangeUsage,omitempty"`
	CCTime                 *datatype.Unsigned32 `avp:"CC-Time" json:"ccTime,omitempty"`
	CcMoney                *CcMoney             `avp:"CC-Money" json:"ccMoney,omitempty"`
	CCTotalOctets          *datatype.Unsigned64 `avp:"CC-Total-Octets" json:"ccTotalOctets,omitempty"`
	CCInputOctets          *datatype.Unsigned64 `avp:"CC-Input-Octets" json:"ccInputOctets,omitempty"`
	CCOutputOctets         *datatype.Unsigned64 `avp:"CC-Output-Octets" json:"ccOutputOctets,omitempty"`
	CCServiceSpecificUnits *datatype.Unsigned64 `avp:"CC-Service-Specific-Units" json:"ccServiceSpecificUnits,omitempty"`
}

type TariffChangeUsage datatype.Enumerated

const (
	TariffChangeUsage_UNIT_BEFORE_TARIFF_CHANGE TariffChangeUsage = 0

	TariffChangeUsage_UNIT_AFTER_TARIFF_CHANGE TariffChangeUsage = 1
	TariffChangeUsage_UNIT_INDETERMINATE       TariffChangeUsage = 2
)

type UsageMonitoringLevel datatype.Enumerated

const (
	UsageMonitoringLevel_SESSION_LEVEL  UsageMonitoringLevel = 0
	UsageMonitoringLevel_PCC_RULE_LEVEL UsageMonitoringLevel = 1
	UsageMonitoringLevel_ADC_RULE_LEVEL UsageMonitoringLevel = 2
)

type UsageMonitoringReport datatype.Enumerated

const (
	UsageMonitoringReport_USAGE_MONITORING_REPORT_REQUIRED UsageMonitoringReport = 0
)

type UsageMonitoringSupport datatype.Enumerated

const (
	USAGE_MONITORING_DISABLED UsageMonitoringSupport = 0
)

type NBIFOMSupport datatype.Enumerated

const (
	NBIFOMSupport_NBIFOM_NOT_SUPPORTED NBIFOMSupport = 0
	NBIFOMSupport_NBIFOM_SUPPORTED     NBIFOMSupport = 1
)

type NBIFOMMode datatype.Enumerated

const (
	NBIFOMMode_UE_INITIATED      NBIFOMMode = 0
	NBIFOMMode_NETWORK_INITIATED NBIFOMMode = 1
)

type DefaultAccess datatype.Enumerated

const (
	DefaultAccess_TGPPGPRS   DefaultAccess = 0
	DefaultAccess_DOCSIS     DefaultAccess = 1
	DefaultAccess_xDSL       DefaultAccess = 2
	DefaultAccess_WiMAX      DefaultAccess = 3
	DefaultAccess_TGPP2      DefaultAccess = 4
	DefaultAccess_TGPPEPS    DefaultAccess = 5
	DefaultAccess_Non3GPPEPS DefaultAccess = 6
	DefaultAccess_FBA        DefaultAccess = 7
	DefaultAccess_TGPP5GS    DefaultAccess = 8
	DefaultAccess_Non3GPP5GS DefaultAccess = 9
)

type RoutingRuleInstall struct {
	RoutingRuleDefinition []RoutingRuleDefinition `avp:"Routing-Rule-Definition" json:"routingRuleDefinition,omitempty"`
}

type RoutingRuleDefinition struct {
	RoutingRuleIdentifier datatype.OctetString `avp:"Routing-Rule-Identifier" json:"routingRuleIdentifier"`
	RoutingFilter         []RoutingFilter      `avp:"Routing-Filter" json:"routingFilter,omitempty"`
	Precedence            *datatype.Unsigned32 `avp:"Precedence" json:"precedence,omitempty"`
	RoutingIPAddress      *datatype.Address    `avp:"Routing-IP-Address" json:"routingIPAddress,omitempty"`
	IPCANType             *IpCanType           `avp:"IP-CAN-Type" json:"ipCANType,omitempty"`
}

type RoutingFilter struct {
	FlowDescription        datatype.IPFilterRule `avp:"Flow-Description" json:"flowDescription"`
	FlowDirection          FlowDirection         `avp:"Flow-Direction" json:"fowDirection"`
	ToSTrafficClass        *datatype.OctetString `avp:"ToS-Traffic-Class" json:"toSTrafficClass,omitempty"`
	SecurityParameterIndex *datatype.OctetString `avp:"Security-Parameter-Index" json:"securityParameterIndex,omitempty"`
	FlowLabel              *datatype.OctetString `avp:"Flow-Label" json:"flowLabel,omitempty"`
}

type RoutingRuleRemove struct {
	RoutingRuleIdentifier []datatype.OctetString `avp:"Routing-Rule-Identifier" json:"routingRuleIdentifier,omitempty"`
}

type PresenceReportingAreaInformation struct {
	PresenceReportingAreaIdentifier   *datatype.OctetString `avp:"Presence-Reporting-Area-Identifier" json:"presenceReportingAreaIdentifier,omitempty"`
	PresenceReportingAreaStatus       *datatype.Unsigned32  `avp:"Presence-Reporting-Area-Status" json:"presenceReportingAreaStatus,omitempty"`
	PresenceReportingAreaElementsList *datatype.OctetString `avp:"Presence-Reporting-Area-Elements-List" json:"presenceReportingAreaElementsList,omitempty"`
	PresenceReportingAreaNode         *datatype.Unsigned32  `avp:"Presence-Reporting-Area-Node" json:"presenceReportingAreaNode,omitempty"`
}

type TGPPPSDataOffStatus datatype.Enumerated

const (
	TGPPPSDataOffStatus_ACTIVE   TGPPPSDataOffStatus = 0
	TGPPPSDataOffStatus_INACTIVE TGPPPSDataOffStatus = 1
)

type CCRequestType datatype.Enumerated

const (
	CCRequestType_INITIAL_REQUEST     CCRequestType = 1
	CCRequestType_UPDATE_REQUEST      CCRequestType = 2
	CCRequestType_TERMINATION_REQUEST CCRequestType = 3
	CCRequestType_EVENT_REQUEST       CCRequestType = 4
)

type OCSupportedFeatures struct {
	OCFeatureVector *datatype.Unsigned64 `avp:"OC-Feature-Vector" json:"ocFeatureVector,omitempty"`
}

type LoadType datatype.Enumerated

const (
	LoadType_HOST LoadType = 0
	LoadType_PEER LoadType = 1
)

type RoutingRuleFailureCode datatype.Enumerated

const (
	RoutingRuleFailureCode_SUBSCRIPTION_LIMITATION                RoutingRuleFailureCode = 0
	RoutingRuleFailureCode_OPERAOR_POLICY                         RoutingRuleFailureCode = 1
	RoutingRuleFailureCode_RESOURCE_LIMITATION                    RoutingRuleFailureCode = 2
	RoutingRuleFailureCode_ROUTING_ACCESS_INFORMATION_NOT_ALLOWED RoutingRuleFailureCode = 3
	RoutingRuleFailureCode_UNSPECIFIED_ERROR                      RoutingRuleFailureCode = 4
)

type FinalUnitAction datatype.Enumerated

const (
	FinalUnitAction_TERMINATE       FinalUnitAction = 0
	FinalUnitAction_REDIRECT        FinalUnitAction = 1
	FinalUnitAction_RESTRICT_ACCESS FinalUnitAction = 2
)

type Load struct {
	LoadType  *LoadType                  `avp:"Load-Type" json:"LoadType,omitempty"`
	LoadValue *datatype.Unsigned64       `avp:"Load-Value" json:"LoadValue,omitempty"`
	SourceId  *datatype.DiameterIdentity `avp:"SourceID" json:"SourceId,omitempty"`
}

type RoutingRuleReport struct {
	RoutingRuleIdentifier  []datatype.OctetString  `avp:"Routing-Rule-Identifier" json:"RoutingRuleIdentifier,omitempty"`
	PccRuleStatus          *PCCRuleStatus          `avp:"PCC-Rule-Status" json:"PccRuleStatus,omitempty"`
	RoutingRuleFailireCode *RoutingRuleFailureCode `avp:"Routing-Rule-Failure-Code" json:"RoutingRuleFailireCode,omitempty"`
}

// TS29.229
type ChargingInformation struct {
	PrimaryEventChargingFunctionName        *datatype.DiameterURI `avp:"Primary-Event-Charging-Function-Name" json:"primaryEventChargingFunctionName,omitempty"`
	SecondaryEventChargingFunctionName      *datatype.DiameterURI `avp:"Secondary-Event-Charging-Function-Name" json:"secondaryEventChargingFunctionName,omitempty"`
	PrimaryChargingCollectionFunctionName   *datatype.DiameterURI `avp:"Primary-Charging-Collection-Function-Name" json:"primaryChargingCollectionFunctionName,omitempty"`
	SecondaryChargingCollectionFunctionName *datatype.DiameterURI `avp:"Secondary-Charging-Collection-Function-Name" json:"secondaryChargingCollectionFunctionName,omitempty"`
}

type PSFurnishChargingInformation struct {
	TGppChargingId         *datatype.OctetString   `avp:"3GPP-Charging-Id" json:"tGppChargingId,omitempty"`
	PSFreeFormatData       *datatype.OctetString   `avp:"PS-Free-Format-Data" json:"pSFreeFormatData,omitempty"`
	PSAppendFreeFormatData *PSAppendFreeFormatData `avp:"PS-Append-Free-Format-Data" json:"pSAppendFreeFormatData,omitempty"`
}
type FinalUnitIndication struct {
	FinalUnitAction       FinalUnitAction         `avp:"Final-Unit-Action" json:"finalUnitAction"`
	RestrictionFilterRule []datatype.IPFilterRule `avp:"Restriction-Filter-Rule" json:"restrictionFilterRule,omitempty"`
	FilterId              []datatype.UTF8String   `avp:"Filter-Id" json:"filterId,omitempty"`
	RedirectServer        *RedirectServer         `avp:"Redirect-Server" json:"redirectServer,omitempty"`
}

type RedirectServer struct {
	RedirectAddressType   RedirectAddressType `avp:"Redirect-Address-Type" json:"redirectAddressType"`
	RedirectServerAddress datatype.UTF8String `avp:"Redirect-Server-Address" json:"redirectServerAddress"`
}

type PSAppendFreeFormatData int

const (
	PSAppendFreeFormatData_Append    PSAppendFreeFormatData = 0
	PSAppendFreeFormatData_Overwrite PSAppendFreeFormatData = 0
)

type ResourceReleaseNotification int

const (
	ResourceReleaseNotification_ENABLE_NOTIFICATION ResourceReleaseNotification = 0
)

const (
	RESULT_CODE_SUCCESS_START   datatype.Unsigned32 = 2000
	RESULT_CODE_SUCCESS         datatype.Unsigned32 = 2001
	RESULT_CODE_LIMITED_SUCCESS datatype.Unsigned32 = 2002
	RESULT_CODE_SUCCESS_END     datatype.Unsigned32 = 2999
)

// RX specific AVPs
type AuthSessionState datatype.Enumerated

const (
	AuthSessionState_STATE_MAINTAINED    AuthSessionState = 0
	AuthSessionState_NO_STATE_MAINTAINED AuthSessionState = 1
)

type AbortCause datatype.Enumerated

const (
	AbortCause_BEARER_RELEASED                        AbortCause = 0
	AbortCause_INSUFFICIENT_SERVER_RESOURCES          AbortCause = 1
	AbortCause_INSUFFICIENT_BEARER_RESOURCES          AbortCause = 2
	AbortCause_PS_TO_CS_HANDOVER                      AbortCause = 3
	AbortCause_SPONSORED_DATA_CONNECTIVITY_DISALLOWED AbortCause = 4
)

type FlowUsage datatype.Enumerated

const (
	FlowUsage_NO_INFORMATION FlowUsage = 0
	FlowUsage_RTCP           FlowUsage = 1
	FlowUsage_AF_SIGNALLING  FlowUsage = 2
)

type IMSContentType datatype.Enumerated

const (
	IMSContentType_NO_CONTENT_DETAIL IMSContentType = 0
	IMSContentType_CAT               IMSContentType = 1
	IMSContentType_CONFERENCE        IMSContentType = 2
)

type MediaType datatype.Enumerated

const (
	MediaType_AUDIO       MediaType = 0
	MediaType_VIDEO       MediaType = 1
	MediaType_DATA        MediaType = 2
	MediaType_APPLICATION MediaType = 3
	MediaType_CONTROL     MediaType = 4
	MediaType_TEXT        MediaType = 5
	MediaType_MESSAGE     MediaType = 6
	MediaType_OTHER       MediaType = 0xFFFFFFF
)

type PrioritySharingIndicator datatype.Enumerated

const (
	PrioritySharingIndicator_PRIORITY_SHARING_ENABLED  PrioritySharingIndicator = 0
	PrioritySharingIndicator_PRIORITY_SHARING_DISABLED PrioritySharingIndicator = 1
)

type RxRequestType datatype.Enumerated

const (
	RxRequestType_INITIAL_REQUEST   RxRequestType = 0
	RxRequestType_UPDATE_REQUEST    RxRequestType = 1
	RxRequestType_PCSCF_RESTORATION RxRequestType = 2
)

type ServiceInfoStatus datatype.Enumerated

const (
	ServiceInfoStatus_FINAL_SERVICE_INFORMATION       ServiceInfoStatus = 0
	ServiceInfoStatus_PRELIMINARY_SERVICE_INFORMATION ServiceInfoStatus = 1
)

type SIPForkingIndication datatype.Enumerated

const (
	SIPForkingIndication_SINGLE_DIALOGUE   SIPForkingIndication = 0
	SIPForkingIndication_SEVERAL_DIALOGUES SIPForkingIndication = 1
)

type SponsoringAction datatype.Enumerated

const (
	SponsoringAction_DISABLE_SPONSORING SponsoringAction = 0
	SponsoringAction_ENABLE_SPONSORING  SponsoringAction = 1
)

type SpecificAction datatype.Enumerated

const (
	SpecificAction_CHARGING_CORRELATION_EXCHANGE                       SpecificAction = 1
	SpecificAction_INDICATION_OF_LOSS_OF_BEARER                        SpecificAction = 2
	SpecificAction_INDICATION_OF_RECOVERY_OF_BEARER                    SpecificAction = 3
	SpecificAction_INDICATION_OF_RELEASE_OF_BEARER                     SpecificAction = 4
	SpecificAction_IP_CAN_CHANGE                                       SpecificAction = 6
	SpecificAction_INDICATION_OF_OUT_OF_CREDIT                         SpecificAction = 7
	SpecificAction_INDICATION_OF_SUCCESSFUL_RESOURCES_ALLOCATION       SpecificAction = 8
	SpecificAction_INDICATION_OF_FAILED_RESOURCES_ALLOCATION           SpecificAction = 9
	SpecificAction_INDICATION_OF_LIMITED_PCC_DEPLOYMENT                SpecificAction = 10
	SpecificAction_USAGE_REPORT                                        SpecificAction = 11
	SpecificAction_ACCESS_NETWORK_INFO_REPORT                          SpecificAction = 12
	SpecificAction_INDICATION_OF_RECOVERY_FROM_LIMITED_PCC_DEPLOYMENT  SpecificAction = 13
	SpecificAction_INDICATION_OF_ACCESS_NETWORK_INFO_REPORTING_FAILURE SpecificAction = 14
	SpecificAction_INDICATION_OF_TRANSFER_POLICY_EXPIRED               SpecificAction = 15
	SpecificAction_PLMN_CHANGE                                         SpecificAction = 16
	SpecificAction_EPS_FALLBACK                                        SpecificAction = 17
	SpecificAction_INDICATION_OF_REALLOCATION_OF_CREDIT                SpecificAction = 18
)

type FGSRANNASReleaseCause struct {
	FGMMCause *datatype.Unsigned32 `avp:"5GMM-Cause" json:"fGMMCause,omitempty"`
	FGSMCause *datatype.Unsigned32 `avp:"5GSM-Cause" json:"fGSMCause,omitempty"`
	// TODO NGAPCause *UnresolvedTypeXXX   `avp:"NGAP-Cause" json:"nGAPCause,omitempty"`
}

type AccessNetworkChargingIdentifier struct {
	AccessNetworkChargingIdentifierValue datatype.OctetString `avp:"Access-Network-Charging-Identifier-Value" json:"accessNetworkChargingIdentifierValue"`
	Flows                                []Flows              `avp:"Flows" json:"flows,omitempty"`
}

type AcceptableServiceInfo struct {
	MediaComponentDescription []MediaComponentDescription `avp:"Media-Component-Description" json:"mediaComponentDescription,omitempty"`
	MaxRequestedBandwidthDL   *datatype.Unsigned32        `avp:"Max-Requested-Bandwidth-DL" json:"maxRequestedBandwidthDL,omitempty"`
	MaxRequestedBandwidthUL   *datatype.Unsigned32        `avp:"Max-Requested-Bandwidth-UL" json:"maxRequestedBandwidthUL,omitempty"`
	ExtendedMaxRequestedBWDL  *datatype.Unsigned32        `avp:"Extended-Max-Requested-BW-DL" json:"extendedMaxRequestedBWDL,omitempty"`
	ExtendedMaxRequestedBWUL  *datatype.Unsigned32        `avp:"Extended-Max-Requested-BW-UL" json:"extendedMaxRequestedBWUL,omitempty"`
}

type MAInformation struct {
	IPCANType           *IpCanType     `avp:"IP-CAN-Type" json:"iPCANType,omitempty"`
	RATType             *RatType       `avp:"RAT-Type" json:"rATType,omitempty"`
	MAInformationAction *MAInformation `avp:"MA-Information-Action" json:"mAInformationAction,omitempty"`
}

type MediaComponentDescription struct {
	MediaComponentNumber     datatype.Unsigned32       `avp:"Media-Component-Number" json:"mediaComponentNumber"`
	MediaSubComponent        []MediaSubComponent       `avp:"Media-Sub-Component" json:"mediaSubComponent,omitempty"`
	AFApplicationIdentifier  *datatype.OctetString     `avp:"AF-Application-Identifier" json:"aFApplicationIdentifier,omitempty"`
	FLUSIdentifier           *datatype.OctetString     `avp:"FLUS-Identifier" json:"fLUSIdentifier,omitempty"`
	MediaType                *MediaType                `avp:"Media-Type" json:"mediaType,omitempty"`
	MaxRequestedBandwidthUL  *datatype.Unsigned32      `avp:"Max-Requested-Bandwidth-UL" json:"maxRequestedBandwidthUL,omitempty"`
	MaxRequestedBandwidthDL  *datatype.Unsigned32      `avp:"Max-Requested-Bandwidth-DL" json:"maxRequestedBandwidthDL,omitempty"`
	MaxSupportedBandwidthUL  *datatype.Unsigned32      `avp:"Max-Supported-Bandwidth-UL" json:"maxSupportedBandwidthUL,omitempty"`
	MaxSupportedBandwidthDL  *datatype.Unsigned32      `avp:"Max-Supported-Bandwidth-DL" json:"maxSupportedBandwidthDL,omitempty"`
	MinDesiredBandwidthUL    *datatype.Unsigned32      `avp:"Min-Desired-Bandwidth-UL" json:"minDesiredBandwidthUL,omitempty"`
	MinDesiredBandwidthDL    *datatype.Unsigned32      `avp:"Min-Desired-Bandwidth-DL" json:"minDesiredBandwidthDL,omitempty"`
	MinRequestedBandwidthUL  *datatype.Unsigned32      `avp:"Min-Requested-Bandwidth-UL" json:"minRequestedBandwidthUL,omitempty"`
	MinRequestedBandwidthDL  *datatype.Unsigned32      `avp:"Min-Requested-Bandwidth-DL" json:"minRequestedBandwidthDL,omitempty"`
	ExtendedMaxRequestedBWUL *datatype.Unsigned32      `avp:"Extended-Max-Requested-BW-UL" json:"extendedMaxRequestedBWUL,omitempty"`
	ExtendedMaxRequestedBWDL *datatype.Unsigned32      `avp:"Extended-Max-Requested-BW-DL" json:"extendedMaxRequestedBWDL,omitempty"`
	ExtendedMaxSupportedBWUL *datatype.Unsigned32      `avp:"Extended-Max-Supported-BW-UL" json:"extendedMaxSupportedBWUL,omitempty"`
	ExtendedMaxSupportedBWDL *datatype.Unsigned32      `avp:"Extended-Max-Supported-BW-DL" json:"extendedMaxSupportedBWDL,omitempty"`
	ExtendedMinDesiredBWUL   *datatype.Unsigned32      `avp:"Extended-Min-Desired-BW-UL" json:"extendedMinDesiredBWUL,omitempty"`
	ExtendedMinDesiredBWDL   *datatype.Unsigned32      `avp:"Extended-Min-Desired-BW-DL" json:"extendedMinDesiredBWDL,omitempty"`
	ExtendedMinRequestedBWUL *datatype.Unsigned32      `avp:"Extended-Min-Requested-BW-UL" json:"extendedMinRequestedBWUL,omitempty"`
	ExtendedMinRequestedBWDL *datatype.Unsigned32      `avp:"Extended-Min-Requested-BW-DL" json:"extendedMinRequestedBWDL,omitempty"`
	FlowStatus               *FlowStatus               `avp:"Flow-Status" json:"flowStatus,omitempty"`
	PrioritySharingIndicator *PrioritySharingIndicator `avp:"Priority-Sharing-Indicator" json:"prioritySharingIndicator,omitempty"`
	PreemptionCapability     *PreemptionCapability     `avp:"Pre-emption-Capability" json:"preemptionCapability,omitempty"`
	PreemptionVulnerability  *PreemptionVulnerability  `avp:"Pre-emption-Vulnerability" json:"preemptionVulnerability,omitempty"`
	ReservationPriority      ReservationPriority       `avp:"Reservation-Priority" json:"reservationPriority,omitempty"`
	RSBandwidth              *datatype.Unsigned32      `avp:"RS-Bandwidth" json:"rSBandwidth,omitempty"`
	RRBandwidth              *datatype.Unsigned32      `avp:"RR-Bandwidth" json:"rRBandwidth,omitempty"`
	CodecData                []datatype.OctetString    `avp:"Codec-Data" json:"codecData,omitempty"`
	SharingKeyDL             *datatype.Unsigned32      `avp:"Sharing-Key-DL" json:"sharingKeyDL,omitempty"`
	SharingKeyUL             *datatype.Unsigned32      `avp:"Sharing-Key-UL" json:"sharingKeyUL,omitempty"`
	ContentVersion           *datatype.Unsigned64      `avp:"Content-Version" json:"contentVersion,omitempty"`
	MaxPLRDL                 *datatype.Float32         `avp:"Max-PLR-DL" json:"maxPLRDL,omitempty"`
	MaxPLRUL                 *datatype.Float32         `avp:"Max-PLR-UL" json:"maxPLRUL,omitempty"`
	DesiredMaxLatency        *datatype.Float32         `avp:"Desired-Max-Latency" json:"desiredMaxLatency,omitempty"`
	DesiredMaxLoss           *datatype.Float32         `avp:"Desired-Max-Loss" json:"desiredMaxLoss,omitempty"`
}

type MediaSubComponent struct {
	FlowNumber               datatype.Unsigned32     `avp:"Flow-Number" json:"flowNumber"`
	FlowDescription          []datatype.IPFilterRule `avp:"Flow-Description" json:"flowDescription,omitempty"`
	FlowStatus               *FlowStatus             `avp:"Flow-Status" json:"flowStatus,omitempty"`
	FlowUsage                *FlowUsage              `avp:"Flow-Usage" json:"flowUsage,omitempty"`
	MaxRequestedBandwidthUL  *datatype.Unsigned32    `avp:"Max-Requested-Bandwidth-UL" json:"maxRequestedBandwidthUL,omitempty"`
	MaxRequestedBandwidthDL  *datatype.Unsigned32    `avp:"Max-Requested-Bandwidth-DL" json:"maxRequestedBandwidthDL,omitempty"`
	ExtendedMaxRequestedBWDL *datatype.Unsigned32    `avp:"Extended-Max-Requested-BW-DL" json:"extendedMaxRequestedBWDL,omitempty"`
	ExtendedMaxRequestedBWUL *datatype.Unsigned32    `avp:"Extended-Max-Requested-BW-UL" json:"extendedMaxRequestedBWUL,omitempty"`
	AFSignallingProtocol     *AFSignallingProtocol   `avp:"AF-Signalling-Protocol" json:"aFSignallingProtocol,omitempty"`
	ToSTrafficClass          *datatype.OctetString   `avp:"ToS-Traffic-Class" json:"toSTrafficClass,omitempty"`
}

type SponsoredConnectivityData struct {
	SponsorIdentity                    *datatype.UTF8String `avp:"Sponsor-Identity" json:"sponsorIdentity,omitempty"`
	ApplicationServiceProviderIdentity *datatype.UTF8String `avp:"Application-Service-Provider-Identity" json:"applicationServiceProviderIdentity,omitempty"`
	GrantedServiceUnit                 *GrantedServiceUnit  `avp:"Granted-Service-Unit" json:"grantedServiceUnit,omitempty"`
	UsedServiceUnit                    *UsedServiceUnit     `avp:"Used-Service-Unit" json:"usedServiceUnit,omitempty"`
	SponsoringAction                   *SponsoringAction    `avp:"Sponsoring-Action" json:"sponsoringAction,omitempty"`
}

type WirelineUserLocationInfo struct {
	HFCNodeIdentifier *datatype.OctetString `avp:"HFC-Node-Identifier" json:"hFCNodeIdentifier,omitempty"`
	GLIIdentifier     *datatype.OctetString `avp:"GLI-Identifier" json:"gLIIdentifier,omitempty"`
	LineType          *datatype.Unsigned32  `avp:"Line-Type" json:"lineType,omitempty"`
}

type TracingConfig struct {
	TraceRef     datatype.OctetString `avp:"Trace-Ref" json:"traceRef"`
	TraceSessRef datatype.OctetString `avp:"Trace-Sess-Ref" json:"traceSessRef"`
	TraceSpanId  datatype.Unsigned64  `avp:"Trace-Span-Id" json:"traceSpanId"`
	TraceIdLow   datatype.Unsigned64  `avp:"Trace-Id-Low" json:"traceIdLow"`
	TraceIdHigh  datatype.Unsigned64  `avp:"Trace-Id-High" json:"traceIdHigh"`
	TraceDepth   TraceDepth           `avp:"Trace-Depth" json:"traceDepth"`
	TraceUEId    datatype.OctetString `avp:"Trace-UE-Id" json:"traceUEId"`
}

type RequestedServiceUnit struct {
	CCTime                 *datatype.Unsigned32 `avp:"CC-Time" json:"cCTime,omitempty"`
	CCMoney                *CCMoney             `avp:"CC-Money" json:"cCMoney,omitempty"`
	CCTotalOctets          *datatype.Unsigned64 `avp:"CC-Total-Octets" json:"cCTotalOctets,omitempty"`
	CCInputOctets          *datatype.Unsigned64 `avp:"CC-Input-Octets" json:"cCInputOctets,omitempty"`
	CCOutputOctets         *datatype.Unsigned64 `avp:"CC-Output-Octets" json:"cCOutputOctets,omitempty"`
	CCServiceSpecificUnits *datatype.Unsigned64 `avp:"CC-Service-Specific-Units" json:"cCServiceSpecificUnits,omitempty"`
}

type CCMoney struct {
	UnitValue    UnitValue            `avp:"Unit-Value" json:"unitValue"`
	CurrencyCode *datatype.Unsigned32 `avp:"Currency-Code" json:"currencyCode,omitempty"`
}

type MultipleServicesCreditControl struct {
	GrantedServiceUnit           *GrantedServiceUnit           `avp:"Granted-Service-Unit" json:"grantedServiceUnit,omitempty"`
	RequestedServiceUnit         *RequestedServiceUnit         `avp:"Requested-Service-Unit" json:"requestedServiceUnit,omitempty"`
	UsedServiceUnit              []UsedServiceUnit             `avp:"Used-Service-Unit" json:"usedServiceUnit,omitempty"`
	TariffChangeUsage            *TariffChangeUsage            `avp:"Tariff-Change-Usage" json:"tariffChangeUsage,omitempty"`
	ServiceIdentifier            []datatype.Unsigned32         `avp:"Service-Identifier" json:"serviceIdentifier,omitempty"`
	RatingGroup                  *datatype.Unsigned32          `avp:"Rating-Group" json:"ratingGroup,omitempty"`
	GSUPoolReference             []GSUPoolReference            `avp:"G-S-U-Pool-Reference" json:"gSUPoolReference,omitempty"`
	ValidityTime                 *datatype.Unsigned32          `avp:"Validity-Time" json:"validityTime,omitempty"`
	ResultCode                   *datatype.Unsigned32          `avp:"Result-Code" json:"resultCode,omitempty"`
	FinalUnitIndication          *FinalUnitIndication          `avp:"Final-Unit-Indication" json:"finalUnitIndication,omitempty"`
	TimeQuotaThreshold           *datatype.Unsigned32          `avp:"Time-Quota-Threshold" json:"timeQuotaThreshold,omitempty"`
	VolumeQuotaThreshold         *datatype.Unsigned32          `avp:"Volume-Quota-Threshold" json:"volumeQuotaThreshold,omitempty"`
	UnitQuotaThreshold           *datatype.Unsigned32          `avp:"Unit-Quota-Threshold" json:"unitQuotaThreshold,omitempty"`
	QuotaHoldingTime             *datatype.Unsigned32          `avp:"Quota-Holding-Time" json:"quotaHoldingTime,omitempty"`
	QuotaConsumptionTime         *datatype.Unsigned32          `avp:"Quota-Consumption-Time" json:"quotaConsumptionTime,omitempty"`
	ReportingReason              []ReportingReason             `avp:"Reporting-Reason" json:"reportingReason,omitempty"`
	Trigger                      *Trigger                      `avp:"Trigger" json:"trigger,omitempty"`
	PSFurnishChargingInformation *PSFurnishChargingInformation `avp:"PS-Furnish-Charging-Information" json:"pSFurnishChargingInformation,omitempty"`
	RefundInformation            *datatype.OctetString         `avp:"Refund-Information" json:"refundInformation,omitempty"`
	AFCorrelationInformation     []AFCorrelationInformation    `avp:"AF-Correlation-Information" json:"aFCorrelationInformation,omitempty"`
	Envelope                     []Envelope                    `avp:"Envelope" json:"envelope,omitempty"`
	EnvelopeReporting            *EnvelopeReporting            `avp:"Envelope-Reporting" json:"envelopeReporting,omitempty"`
	TimeQuotaMechanism           *TimeQuotaMechanism           `avp:"Time-Quota-Mechanism" json:"timeQuotaMechanism,omitempty"`
	ServiceSpecificInfo          []ServiceSpecificInfo         `avp:"Service-Specific-Info" json:"serviceSpecificInfo,omitempty"`
	QoSInformation               *QoSInformation               `avp:"QoS-Information" json:"qoSInformation,omitempty"`
	AnnouncementInformation      []AnnouncementInformation     `avp:"Announcement-Information" json:"announcementInformation,omitempty"`
	TGPPRATType                  *datatype.OctetString         `avp:"3GPP-RAT-Type" json:"tGPPRATType,omitempty"`
	RelatedTrigger               *RelatedTrigger               `avp:"Related-Trigger" json:"relatedTrigger,omitempty"`
}

type GSUPoolReference struct {
	GSUPoolIdentifier datatype.Unsigned32 `avp:"G-S-U-Pool-Identifier" json:"gSUPoolIdentifier"`
	CCUnitType        CCUnitType          `avp:"CC-Unit-Type" json:"cCUnitType"`
	UnitValue         UnitValue           `avp:"Unit-Value" json:"unitValue"`
}

type Trigger struct {
	TriggerType []TriggerType `avp:"Trigger-Type" json:"triggerType,omitempty"`
}

type AFCorrelationInformation struct {
	AFChargingIdentifier datatype.OctetString `avp:"AF-Charging-Identifier" json:"aFChargingIdentifier"`
	Flows                []Flows              `avp:"Flows" json:"flows,omitempty"`
}

type Envelope struct {
	EnvelopeStartTime      datatype.Time        `avp:"Envelope-Start-Time" json:"envelopeStartTime"`
	EnvelopeEndTime        *datatype.Time       `avp:"Envelope-End-Time" json:"envelopeEndTime,omitempty"`
	CCTotalOctets          *datatype.Unsigned64 `avp:"CC-Total-Octets" json:"cCTotalOctets,omitempty"`
	CCInputOctets          *datatype.Unsigned64 `avp:"CC-Input-Octets" json:"cCInputOctets,omitempty"`
	CCOutputOctets         *datatype.Unsigned64 `avp:"CC-Output-Octets" json:"cCOutputOctets,omitempty"`
	CCServiceSpecificUnits *datatype.Unsigned64 `avp:"CC-Service-Specific-Units" json:"cCServiceSpecificUnits,omitempty"`
}

type TimeQuotaMechanism struct {
	TimeQuotaType    TimeQuotaType       `avp:"Time-Quota-Type" json:"timeQuotaType"`
	BaseTimeInterval datatype.Unsigned32 `avp:"Base-Time-Interval" json:"baseTimeInterval"`
}

type ServiceSpecificInfo struct {
	ServiceSpecificData *datatype.UTF8String `avp:"Service-Specific-Data" json:"serviceSpecificData,omitempty"`
	ServiceSpecificType *datatype.Unsigned32 `avp:"Service-Specific-Type" json:"serviceSpecificType,omitempty"`
}

type AnnouncementInformation struct {
	AnnouncementIdentifier datatype.Unsigned32  `avp:"Announcement-Identifier" json:"announcementIdentifier"`
	VariablePart           []VariablePart       `avp:"Variable-Part" json:"variablePart,omitempty"`
	TimeIndicator          *datatype.Unsigned32 `avp:"Time-Indicator" json:"timeIndicator,omitempty"`
	QuotaIndicator         *QuotaIndicator      `avp:"Quota-Indicator" json:"quotaIndicator,omitempty"`
	AnnouncementOrder      *datatype.Unsigned32 `avp:"Announcement-Order" json:"announcementOrder,omitempty"`
	PlayAlternative        *PlayAlternative     `avp:"Play-Alternative" json:"playAlternative,omitempty"`
	PrivacyIndicator       *PrivacyIndicator    `avp:"Privacy-Indicator" json:"privacyIndicator,omitempty"`
	Language               *datatype.UTF8String `avp:"Language" json:"language,omitempty"`
}

type VariablePart struct {
	VariablePartOrder *datatype.Unsigned32 `avp:"Variable-Part-Order" json:"variablePartOrder,omitempty"`
	VariablePartType  datatype.Unsigned32  `avp:"Variable-Part-Type" json:"variablePartType"`
	VariablePartValue datatype.UTF8String  `avp:"Variable-Part-Value" json:"variablePartValue"`
}

type RelatedTrigger struct {
	TriggerType []TriggerType `avp:"Trigger-Type" json:"triggerType,omitempty"`
}

type ServiceParameterInfo struct {
	ServiceParameterType  datatype.Unsigned32  `avp:"Service-Parameter-Type" json:"serviceParameterType"`
	ServiceParameterValue datatype.OctetString `avp:"Service-Parameter-Value" json:"serviceParameterValue"`
}

type ServiceInformation struct {
	SubscriptionId            []SubscriptionId           `avp:"Subscription-Id" json:"subscriptionId,omitempty"`
	AoCInformation            *AoCInformation            `avp:"AoC-Information" json:"aoCInformation,omitempty"`
	PSInformation             *PSInformation             `avp:"PS-Information" json:"pSInformation,omitempty"`
	IMSInformation            *IMSInformation            `avp:"IMS-Information" json:"iMSInformation,omitempty"`
	MMSInformation            *MMSInformation            `avp:"MMS-Information" json:"mMSInformation,omitempty"`
	LCSInformation            *LCSInformation            `avp:"LCS-Information" json:"lCSInformation,omitempty"`
	PoCInformation            *PoCInformation            `avp:"PoC-Information" json:"poCInformation,omitempty"`
	MBMSInformation           *MBMSInformation           `avp:"MBMS-Information" json:"mBMSInformation,omitempty"`
	SMSInformation            *SMSInformation            `avp:"SMS-Information" json:"sMSInformation,omitempty"`
	VCSInformation            *VCSInformation            `avp:"VCS-Information" json:"vCSInformation,omitempty"`
	MMTelInformation          *MMTelInformation          `avp:"MMTel-Information" json:"mMTelInformation,omitempty"`
	ProSeInformation          *ProSeInformation          `avp:"ProSe-Information" json:"proSeInformation,omitempty"`
	ServiceGenericInformation *ServiceGenericInformation `avp:"Service-Generic-Information" json:"serviceGenericInformation,omitempty"`
	IMInformation             *IMInformation             `avp:"IM-Information" json:"iMInformation,omitempty"`
	DCDInformation            *DCDInformation            `avp:"DCD-Information" json:"dCDInformation,omitempty"`
	M2MInformation            *M2MInformation            `avp:"M2M-Information" json:"m2MInformation,omitempty"`
	CPDTInformation           *CPDTInformation           `avp:"CPDT-Information" json:"cPDTInformation,omitempty"`
}

type AoCInformation struct {
	AoCCostInformation         *AoCCostInformation         `avp:"AoC-Cost-Information" json:"aoCCostInformation,omitempty"`
	TariffInformation          *TariffInformation          `avp:"Tariff-Information" json:"tariffInformation,omitempty"`
	AoCSubscriptionInformation *AoCSubscriptionInformation `avp:"AoC-Subscription-Information" json:"aoCSubscriptionInformation,omitempty"`
}

type AoCCostInformation struct {
	AccumulatedCost *AccumulatedCost     `avp:"Accumulated-Cost" json:"accumulatedCost,omitempty"`
	IncrementalCost []IncrementalCost    `avp:"Incremental-Cost" json:"incrementalCost,omitempty"`
	CurrencyCode    *datatype.Unsigned32 `avp:"Currency-Code" json:"currencyCode,omitempty"`
}

type AccumulatedCost struct {
	ValueDigits datatype.Integer64  `avp:"Value-Digits" json:"valueDigits"`
	Exponent    *datatype.Integer32 `avp:"Exponent" json:"exponent,omitempty"`
}

type IncrementalCost struct {
	ValueDigits datatype.Integer64  `avp:"Value-Digits" json:"valueDigits"`
	Exponent    *datatype.Integer32 `avp:"Exponent" json:"exponent,omitempty"`
}

type TariffInformation struct {
	CurrentTariff    CurrentTariff  `avp:"Current-Tariff" json:"currentTariff"`
	TariffTimeChange *datatype.Time `avp:"Tariff-Time-Change" json:"tariffTimeChange,omitempty"`
	NextTariff       *NextTariff    `avp:"Next-Tariff" json:"nextTariff,omitempty"`
}

type CurrentTariff struct {
	CurrencyCode *datatype.Unsigned32 `avp:"Currency-Code" json:"currencyCode,omitempty"`
	ScaleFactor  *ScaleFactor         `avp:"Scale-Factor" json:"scaleFactor,omitempty"`
	RateElement  []RateElement        `avp:"Rate-Element" json:"rateElement,omitempty"`
}

type ScaleFactor struct {
	ValueDigits datatype.Integer64  `avp:"Value-Digits" json:"valueDigits"`
	Exponent    *datatype.Integer32 `avp:"Exponent" json:"exponent,omitempty"`
}

type RateElement struct {
	CCUnitType         CCUnitType           `avp:"CC-Unit-Type" json:"cCUnitType"`
	ChargeReasonCode   *ChargeReasonCode    `avp:"Charge-Reason-Code" json:"chargeReasonCode,omitempty"`
	UnitValue          *UnitValue           `avp:"Unit-Value" json:"unitValue,omitempty"`
	UnitCost           *UnitCost            `avp:"Unit-Cost" json:"unitCost,omitempty"`
	UnitQuotaThreshold *datatype.Unsigned32 `avp:"Unit-Quota-Threshold" json:"unitQuotaThreshold,omitempty"`
}

type UnitCost struct {
	ValueDigits datatype.Integer64  `avp:"Value-Digits" json:"valueDigits"`
	Exponent    *datatype.Integer32 `avp:"Exponent" json:"exponent,omitempty"`
}

type NextTariff struct {
	CurrencyCode *datatype.Unsigned32 `avp:"Currency-Code" json:"currencyCode,omitempty"`
	ScaleFactor  *ScaleFactor         `avp:"Scale-Factor" json:"scaleFactor,omitempty"`
	RateElement  []RateElement        `avp:"Rate-Element" json:"rateElement,omitempty"`
}

type AoCSubscriptionInformation struct {
	AoCService           []AoCService         `avp:"AoC-Service" json:"aoCService,omitempty"`
	AoCFormat            *AoCFormat           `avp:"AoC-Format" json:"aoCFormat,omitempty"`
	PreferredAoCCurrency *datatype.Unsigned32 `avp:"Preferred-AoC-Currency" json:"preferredAoCCurrency,omitempty"`
}

type AoCService struct {
	AoCServiceObligatoryType *AoCServiceObligatoryType `avp:"AoC-Service-Obligatory-Type" json:"aoCServiceObligatoryType,omitempty"`
	AoCServiceType           *AoCServiceType           `avp:"AoC-Service-Type" json:"aoCServiceType,omitempty"`
}

type PSInformation struct {
	SupportedFeatures                    []SupportedFeatures                   `avp:"Supported-Features" json:"supportedFeatures,omitempty"`
	TGPPChargingId                       *datatype.OctetString                 `avp:"3GPP-Charging-Id" json:"tGPPChargingId,omitempty" cmp:"ignore"`
	PDNConnectionChargingID              *datatype.Unsigned32                  `avp:"PDN-Connection-Charging-ID" json:"pDNConnectionChargingID,omitempty" cmp:"ignore"`
	NodeId                               *datatype.UTF8String                  `avp:"Node-Id" json:"nodeId,omitempty" cmp:"ignore"`
	TGPPPDPType                          *TGPPPDPType                          `avp:"3GPP-PDP-Type" json:"tGPPPDPType,omitempty"`
	TGPPGPRSNegotiatedQoSProfile         *datatype.UTF8String                  `avp:"3GPP-GPRS-Negotiated-QoS-Profile" json:"tGPPGPRSNegotiatedQoSProfile,omitempty" cmp:"ignore"`
	PDPAddress                           []datatype.Address                    `avp:"PDP-Address" json:"pDPAddress,omitempty" cmp:"ignore"`
	PDPAddressPrefixLength               *datatype.Unsigned32                  `avp:"PDP-Address-Prefix-Length" json:"pDPAddressPrefixLength,omitempty" cmp:"ignore"`
	DynamicAddressFlag                   *DynamicAddressFlag                   `avp:"Dynamic-Address-Flag" json:"dynamicAddressFlag,omitempty" cmp:"ignore"`
	DynamicAddressFlagExtension          *DynamicAddressFlagExtension          `avp:"Dynamic-Address-Flag-Extension" json:"dynamicAddressFlagExtension,omitempty" cmp:"ignore"`
	QoSInformation                       *QoSInformation                       `avp:"QoS-Information" json:"qoSInformation,omitempty"`
	SGSNAddress                          []datatype.Address                    `avp:"SGSN-Address" json:"sGSNAddress,omitempty" cmp:"ignore"`
	GGSNAddress                          []datatype.Address                    `avp:"GGSN-Address" json:"gGSNAddress,omitempty" cmp:"ignore"`
	TDFIPAddress                         []datatype.Address                    `avp:"TDF-IP-Address" json:"tDFIPAddress,omitempty" cmp:"ignore"`
	SGWAddress                           []datatype.Address                    `avp:"SGW-Address" json:"sGWAddress,omitempty" cmp:"ignore"`
	EPDGAddress                          []datatype.Address                    `avp:"EPDG-Address" json:"ePDGAddress,omitempty" cmp:"ignore"`
	TWAGAddress                          []datatype.Address                    `avp:"TWAG-Address" json:"tWAGAddress,omitempty" cmp:"ignore"`
	CGAddress                            *datatype.Address                     `avp:"CG-Address" json:"cGAddress,omitempty" cmp:"ignore"`
	ServingNodeType                      *ServingNodeType                      `avp:"Serving-Node-Type" json:"servingNodeType,omitempty" cmp:"ignore"`
	SGWChange                            *SGWChange                            `avp:"SGW-Change" json:"sGWChange,omitempty" cmp:"ignore"`
	TGPPIMSIMCCMNC                       *datatype.UTF8String                  `avp:"3GPP-IMSI-MCC-MNC" json:"tGPPIMSIMCCMNC,omitempty" cmp:"ignore"`
	IMSIUnauthenticatedFlag              *IMSIUnauthenticatedFlag              `avp:"IMSI-Unauthenticated-Flag" json:"iMSIUnauthenticatedFlag,omitempty"`
	TGPPGGSNMCCMNC                       *datatype.UTF8String                  `avp:"3GPP-GGSN-MCC-MNC" json:"tGPPGGSNMCCMNC,omitempty" cmp:"ignore"`
	TGPPNSAPI                            *datatype.OctetString                 `avp:"3GPP-NSAPI" json:"tGPPNSAPI,omitempty" cmp:"ignore"`
	CalledStationId                      *datatype.UTF8String                  `avp:"Called-Station-Id" json:"calledStationId,omitempty" cmp:"ignore"`
	TGPPSessionStopIndicator             *datatype.UTF8String                  `avp:"3GPP-Session-Stop-Indicator" json:"tGPPSessionStopIndicator,omitempty" cmp:"ignore"`
	TGPPSelectionMode                    *datatype.UTF8String                  `avp:"3GPP-Selection-Mode" json:"tGPPSelectionMode,omitempty"`
	TGPPChargingCharacteristics          *datatype.UTF8String                  `avp:"3GPP-Charging-Characteristics" json:"tGPPChargingCharacteristics,omitempty"`
	ChargingCharacteristicsSelectionMode *ChargingCharacteristicsSelectionMode `avp:"Charging-Characteristics-Selection-Mode" json:"chargingCharacteristicsSelectionMode,omitempty"`
	TGPPSGSNMCCMNC                       *datatype.UTF8String                  `avp:"3GPP-SGSN-MCC-MNC" json:"tGPPSGSNMCCMNC,omitempty" cmp:"ignore"`
	TGPPMSTimeZone                       *datatype.OctetString                 `avp:"3GPP-MS-TimeZone" json:"tGPPMSTimeZone,omitempty" cmp:"ignore"`
	ChargingRuleBaseName                 *datatype.UTF8String                  `avp:"Charging-Rule-Base-Name" json:"chargingRuleBaseName,omitempty"`
	ADCRuleBaseName                      *datatype.UTF8String                  `avp:"ADC-Rule-Base-Name" json:"aDCRuleBaseName,omitempty"`
	TGPPUserLocationInfo                 *datatype.OctetString                 `avp:"3GPP-User-Location-Info" json:"tGPPUserLocationInfo,omitempty" cmp:"ignore"`
	UserLocationInfoTime                 *datatype.Time                        `avp:"User-Location-Info-Time" json:"userLocationInfoTime,omitempty" cmp:"ignore"`
	UserCSGInformation                   *UserCSGInformation                   `avp:"User-CSG-Information" json:"userCSGInformation,omitempty" cmp:"ignore"`
	PresenceReportingAreaInformation     []PresenceReportingAreaInformation    `avp:"Presence-Reporting-Area-Information" json:"presenceReportingAreaInformation,omitempty"`
	TGPP2BSID                            *datatype.UTF8String                  `avp:"3GPP2-BSID" json:"tGPP2BSID,omitempty" cmp:"ignore"`
	TWANUserLocationInfo                 *TWANUserLocationInfo                 `avp:"TWAN-User-Location-Info" json:"tWANUserLocationInfo,omitempty" cmp:"ignore"`
	UWANUserLocationInfo                 *UWANUserLocationInfo                 `avp:"UWAN-User-Location-Info" json:"uWANUserLocationInfo,omitempty" cmp:"ignore"`
	TGPPRATType                          *datatype.OctetString                 `avp:"3GPP-RAT-Type" json:"tGPPRATType,omitempty"`
	PSFurnishChargingInformation         *PSFurnishChargingInformation         `avp:"PS-Furnish-Charging-Information" json:"pSFurnishChargingInformation,omitempty" cmp:"ignore"`
	PDPContextType                       *PDPContextType                       `avp:"PDP-Context-Type" json:"pDPContextType,omitempty" cmp:"ignore"`
	OfflineCharging                      *OfflineCharging                      `avp:"Offline-Charging" json:"offlineCharging,omitempty"`
	TrafficDataVolumes                   []TrafficDataVolumes                  `avp:"Traffic-Data-Volumes" json:"trafficDataVolumes,omitempty" cmp:"ignore"`
	ServiceDataContainer                 []ServiceDataContainer                `avp:"Service-Data-Container" json:"serviceDataContainer,omitempty"`
	UserEquipmentInfo                    *UserEquipmentInfo                    `avp:"User-Equipment-Info" json:"userEquipmentInfo,omitempty" cmp:"ignore"`
	TerminalInformation                  *TerminalInformation                  `avp:"Terminal-Information" json:"terminalInformation,omitempty" cmp:"ignore"`
	StartTime                            *datatype.Time                        `avp:"Start-Time" json:"startTime,omitempty" cmp:"ignore"`
	StopTime                             *datatype.Time                        `avp:"Stop-Time" json:"stopTime,omitempty" cmp:"ignore"`
	ChangeCondition                      *datatype.Integer32                   `avp:"Change-Condition" json:"changeCondition,omitempty"`
	Diagnostics                          *datatype.Integer32                   `avp:"Diagnostics" json:"diagnostics,omitempty" cmp:"ignore"`
	LowPriorityIndicator                 *LowPriorityIndicator                 `avp:"Low-Priority-Indicator" json:"lowPriorityIndicator,omitempty" cmp:"ignore"`
	NBIFOMMode                           *NBIFOMMode                           `avp:"NBIFOM-Mode" json:"nBIFOMMode,omitempty" cmp:"ignore"`
	NBIFOMSupport                        *NBIFOMSupport                        `avp:"NBIFOM-Support" json:"nBIFOMSupport,omitempty" cmp:"ignore"`
	MMENumberforMTSMS                    *datatype.OctetString                 `avp:"MME-Number-for-MT-SMS" json:"mMENumberforMTSMS,omitempty" cmp:"ignore"`
	MMEName                              *datatype.DiameterIdentity            `avp:"MME-Name" json:"mMEName,omitempty" cmp:"ignore"`
	MMERealm                             *datatype.DiameterIdentity            `avp:"MME-Realm" json:"mMERealm,omitempty" cmp:"ignore"`
	LogicalAccessID                      *datatype.OctetString                 `avp:"Logical-Access-ID" json:"logicalAccessID,omitempty" cmp:"ignore"`
	PhysicalAccessID                     *datatype.UTF8String                  `avp:"Physical-Access-ID" json:"physicalAccessID,omitempty" cmp:"ignore"`
	FixedUserLocationInfo                *FixedUserLocationInfo                `avp:"Fixed-User-Location-Info" json:"fixedUserLocationInfo,omitempty" cmp:"ignore"`
	CNOperatorSelectionEntity            *CNOperatorSelectionEntity            `avp:"CN-Operator-Selection-Entity" json:"cNOperatorSelectionEntity,omitempty" cmp:"ignore"`
	EnhancedDiagnostics                  *EnhancedDiagnostics                  `avp:"Enhanced-Diagnostics" json:"enhancedDiagnostics,omitempty" cmp:"ignore"`
	SGiPtPTunnellingMethod               *SGiPtPTunnellingMethod               `avp:"SGi-PtP-Tunnelling-Method" json:"sGiPtPTunnellingMethod,omitempty" cmp:"ignore"`
	CPCIoTEPSOptimisationIndicator       *CPCIoTEPSOptimisationIndicator       `avp:"CP-CIoT-EPS-Optimisation-Indicator" json:"cPCIoTEPSOptimisationIndicator,omitempty" cmp:"ignore"`
	UNIPDUCPOnlyFlag                     *UNIPDUCPOnlyFlag                     `avp:"UNI-PDU-CP-Only-Flag" json:"uNIPDUCPOnlyFlag,omitempty" cmp:"ignore"`
	ServingPLMNRateControl               *ServingPLMNRateControl               `avp:"Serving-PLMN-Rate-Control" json:"servingPLMNRateControl,omitempty" cmp:"ignore"`
	APNRateControl                       *APNRateControl                       `avp:"APN-Rate-Control" json:"aPNRateControl,omitempty" cmp:"ignore"`
	ChargingPerIPCANSessionIndicator     *ChargingPerIPCANSessionIndicator     `avp:"Charging-Per-IP-CAN-Session-Indicator" json:"chargingPerIPCANSessionIndicator,omitempty" cmp:"ignore"`
	RRCCauseCounter                      *RRCCauseCounter                      `avp:"RRC-Cause-Counter" json:"rRCCauseCounter,omitempty" cmp:"ignore"`
	TGPPPSDataOffStatus                  *TGPPPSDataOffStatus                  `avp:"3GPP-PS-Data-Off-Status" json:"tGPPPSDataOffStatus,omitempty" cmp:"ignore"`
	SCSASAddress                         *SCSASAddress                         `avp:"SCS-AS-Address" json:"sCSASAddress,omitempty" cmp:"ignore"`
	UnusedQuotaTimer                     *datatype.Unsigned32                  `avp:"Unused-Quota-Timer" json:"unusedQuotaTimer,omitempty" cmp:"ignore"`
	RANSecondaryRATUsageReport           []RANSecondaryRATUsageReport          `avp:"RAN-Secondary-RAT-Usage-Report" json:"rANSecondaryRATUsageReport,omitempty"`
}

type UserCSGInformation struct {
	CSGId                   datatype.Unsigned32      `avp:"CSG-Id" json:"cSGId"`
	CSGAccessMode           CSGAccessMode            `avp:"CSG-Access-Mode" json:"cSGAccessMode"`
	CSGMembershipIndication *CSGMembershipIndication `avp:"CSG-Membership-Indication" json:"cSGMembershipIndication,omitempty"`
}

type TWANUserLocationInfo struct {
	SSID                    datatype.UTF8String   `avp:"SSID" json:"sSID"`
	BSSID                   *datatype.UTF8String  `avp:"BSSID" json:"bSSID,omitempty"`
	CivicAddressInformation *datatype.UTF8String  `avp:"Civic-Address-Information" json:"civicAddressInformation,omitempty"`
	WLANOperatorId          *WLANOperatorId       `avp:"WLAN-Operator-Id" json:"wLANOperatorId,omitempty"`
	LogicalAccessID         *datatype.OctetString `avp:"Logical-Access-ID" json:"logicalAccessID,omitempty"`
}

type WLANOperatorId struct {
	WLANPLMNId       *datatype.UTF8String `avp:"WLAN-PLMN-Id" json:"wLANPLMNId,omitempty"`
	WLANOperatorName *datatype.UTF8String `avp:"WLAN-Operator-Name" json:"wLANOperatorName,omitempty"`
}

type UWANUserLocationInfo struct {
	UELocalIPAddress        datatype.Address      `avp:"UE-Local-IP-Address" json:"uELocalIPAddress"`
	UDPSourcePort           *datatype.Unsigned32  `avp:"UDP-Source-Port" json:"uDPSourcePort,omitempty"`
	SSID                    *datatype.UTF8String  `avp:"SSID" json:"sSID,omitempty"`
	BSSID                   *datatype.UTF8String  `avp:"BSSID" json:"bSSID,omitempty"`
	TCPSourcePort           *datatype.Unsigned32  `avp:"TCP-Source-Port" json:"tCPSourcePort,omitempty"`
	CivicAddressInformation *datatype.UTF8String  `avp:"Civic-Address-Information" json:"civicAddressInformation,omitempty"`
	WLANOperatorId          *WLANOperatorId       `avp:"WLAN-Operator-Id" json:"wLANOperatorId,omitempty"`
	LogicalAccessID         *datatype.OctetString `avp:"Logical-Access-ID" json:"logicalAccessID,omitempty"`
}

type OfflineCharging struct {
	QuotaConsumptionTime          *datatype.Unsigned32            `avp:"Quota-Consumption-Time" json:"quotaConsumptionTime,omitempty"`
	TimeQuotaMechanism            *TimeQuotaMechanism             `avp:"Time-Quota-Mechanism" json:"timeQuotaMechanism,omitempty"`
	EnvelopeReporting             *EnvelopeReporting              `avp:"Envelope-Reporting" json:"envelopeReporting,omitempty"`
	MultipleServicesCreditControl []MultipleServicesCreditControl `avp:"Multiple-Services-Credit-Control" json:"multipleServicesCreditControl,omitempty"`
}

type TrafficDataVolumes struct {
	QoSInformation                    *QoSInformation                    `avp:"QoS-Information" json:"qoSInformation,omitempty"`
	AccountingInputOctets             *datatype.Unsigned64               `avp:"Accounting-Input-Octets" json:"accountingInputOctets,omitempty"`
	AccountingOutputOctets            *datatype.Unsigned64               `avp:"Accounting-Output-Octets" json:"accountingOutputOctets,omitempty"`
	ChangeCondition                   *datatype.Integer32                `avp:"Change-Condition" json:"changeCondition,omitempty"`
	ChangeTime                        *datatype.Time                     `avp:"Change-Time" json:"changeTime,omitempty"`
	TGPPUserLocationInfo              *datatype.OctetString              `avp:"3GPP-User-Location-Info" json:"tGPPUserLocationInfo,omitempty"`
	UWANUserLocationInfo              *UWANUserLocationInfo              `avp:"UWAN-User-Location-Info" json:"uWANUserLocationInfo,omitempty"`
	TGPPChargingId                    *datatype.OctetString              `avp:"3GPP-Charging-Id" json:"tGPPChargingId,omitempty"`
	PresenceReportingAreaStatus       *datatype.Unsigned32               `avp:"Presence-Reporting-Area-Status" json:"presenceReportingAreaStatus,omitempty"`
	PresenceReportingAreaInformation  []PresenceReportingAreaInformation `avp:"Presence-Reporting-Area-Information" json:"presenceReportingAreaInformation,omitempty"`
	UserCSGInformation                *UserCSGInformation                `avp:"User-CSG-Information" json:"userCSGInformation,omitempty"`
	TGPPRATType                       *datatype.OctetString              `avp:"3GPP-RAT-Type" json:"tGPPRATType,omitempty"`
	AccessAvailabilityChangeReason    *datatype.Unsigned32               `avp:"Access-Availability-Change-Reason" json:"accessAvailabilityChangeReason,omitempty"`
	RelatedChangeConditionInformation *RelatedChangeConditionInformation `avp:"Related-Change-Condition-Information" json:"relatedChangeConditionInformation,omitempty"`
	Diagnostics                       *datatype.Integer32                `avp:"Diagnostics" json:"diagnostics,omitempty"`
	EnhancedDiagnostics               *EnhancedDiagnostics               `avp:"Enhanced-Diagnostics" json:"enhancedDiagnostics,omitempty"`
	CPCIoTEPSOptimisationIndicator    *CPCIoTEPSOptimisationIndicator    `avp:"CP-CIoT-EPS-Optimisation-Indicator" json:"cPCIoTEPSOptimisationIndicator,omitempty"`
	ServingPLMNRateControl            *ServingPLMNRateControl            `avp:"Serving-PLMN-Rate-Control" json:"servingPLMNRateControl,omitempty"`
	APNRateControl                    *APNRateControl                    `avp:"APN-Rate-Control" json:"aPNRateControl,omitempty"`
}

type RelatedChangeConditionInformation struct {
	SGSNAddress                 *datatype.Address     `avp:"SGSN-Address" json:"sGSNAddress,omitempty"`
	ChangeCondition             []datatype.Integer32  `avp:"Change-Condition" json:"changeCondition,omitempty"`
	TGPPUserLocationInfo        *datatype.OctetString `avp:"3GPP-User-Location-Info" json:"tGPPUserLocationInfo,omitempty"`
	TGPP2BSID                   *datatype.UTF8String  `avp:"3GPP2-BSID" json:"tGPP2BSID,omitempty"`
	UWANUserLocationInfo        *UWANUserLocationInfo `avp:"UWAN-User-Location-Info" json:"uWANUserLocationInfo,omitempty"`
	PresenceReportingAreaStatus *datatype.Unsigned32  `avp:"Presence-Reporting-Area-Status" json:"presenceReportingAreaStatus,omitempty"`
	UserCSGInformation          *UserCSGInformation   `avp:"User-CSG-Information" json:"userCSGInformation,omitempty"`
	TGPPRATType                 *datatype.OctetString `avp:"3GPP-RAT-Type" json:"tGPPRATType,omitempty"`
}

type EnhancedDiagnostics struct {
	RANNASReleaseCause []datatype.OctetString `avp:"RAN-NAS-Release-Cause" json:"rANNASReleaseCause,omitempty"`
}

type ServingPLMNRateControl struct {
	UplinkRateLimit   *datatype.Unsigned32 `avp:"Uplink-Rate-Limit" json:"uplinkRateLimit,omitempty"`
	DownlinkRateLimit *datatype.Unsigned32 `avp:"Downlink-Rate-Limit" json:"downlinkRateLimit,omitempty"`
}

type APNRateControl struct {
	APNRateControlUplink   *APNRateControlUplink   `avp:"APN-Rate-Control-Uplink" json:"aPNRateControlUplink,omitempty"`
	APNRateControlDownlink *APNRateControlDownlink `avp:"APN-Rate-Control-Downlink" json:"aPNRateControlDownlink,omitempty"`
}

type APNRateControlUplink struct {
	AdditionalExceptionReports *AdditionalExceptionReports `avp:"Additional-Exception-Reports" json:"additionalExceptionReports,omitempty"`
	RateControlTimeUnit        *datatype.Unsigned32        `avp:"Rate-Control-Time-Unit" json:"rateControlTimeUnit,omitempty"`
	RateControlMaxRate         *datatype.Unsigned32        `avp:"Rate-Control-Max-Rate" json:"rateControlMaxRate,omitempty"`
}

type APNRateControlDownlink struct {
	RateControlTimeUnit       *datatype.Unsigned32 `avp:"Rate-Control-Time-Unit" json:"rateControlTimeUnit,omitempty"`
	RateControlMaxRate        *datatype.Unsigned32 `avp:"Rate-Control-Max-Rate" json:"rateControlMaxRate,omitempty"`
	RateControlMaxMessageSize *datatype.Unsigned32 `avp:"Rate-Control-Max-Message-Size" json:"rateControlMaxMessageSize,omitempty"`
}

type ServiceDataContainer struct {
	AFCorrelationInformation           *AFCorrelationInformation          `avp:"AF-Correlation-Information" json:"aFCorrelationInformation,omitempty"`
	ChargingRuleBaseName               *datatype.UTF8String               `avp:"Charging-Rule-Base-Name" json:"chargingRuleBaseName,omitempty"`
	AccountingInputOctets              *datatype.Unsigned64               `avp:"Accounting-Input-Octets" json:"accountingInputOctets,omitempty"`
	AccountingOutputOctets             *datatype.Unsigned64               `avp:"Accounting-Output-Octets" json:"accountingOutputOctets,omitempty"`
	LocalSequenceNumber                *datatype.Unsigned32               `avp:"Local-Sequence-Number" json:"localSequenceNumber,omitempty"`
	QoSInformation                     *QoSInformation                    `avp:"QoS-Information" json:"qoSInformation,omitempty"`
	RatingGroup                        *datatype.Unsigned32               `avp:"Rating-Group" json:"ratingGroup,omitempty"`
	ChangeTime                         *datatype.Time                     `avp:"Change-Time" json:"changeTime,omitempty" cmp:"ignore"`
	ServiceIdentifier                  *datatype.Unsigned32               `avp:"Service-Identifier" json:"serviceIdentifier,omitempty"`
	ServiceSpecificInfo                *ServiceSpecificInfo               `avp:"Service-Specific-Info" json:"serviceSpecificInfo,omitempty"`
	ADCRuleBaseName                    *datatype.UTF8String               `avp:"ADC-Rule-Base-Name" json:"aDCRuleBaseName,omitempty"`
	SGSNAddress                        *datatype.Address                  `avp:"SGSN-Address" json:"sGSNAddress,omitempty" cmp:"ignore"`
	TimeFirstUsage                     *datatype.Time                     `avp:"Time-First-Usage" json:"timeFirstUsage,omitempty" cmp:"ignore"`
	TimeLastUsage                      *datatype.Time                     `avp:"Time-Last-Usage" json:"timeLastUsage,omitempty" cmp:"ignore"`
	TimeUsage                          *datatype.Unsigned32               `avp:"Time-Usage" json:"timeUsage,omitempty"`
	ChangeCondition                    []datatype.Integer32               `avp:"Change-Condition" json:"changeCondition,omitempty"`
	TGPPUserLocationInfo               *datatype.OctetString              `avp:"3GPP-User-Location-Info" json:"tGPPUserLocationInfo,omitempty" cmp:"ignore"`
	TGPP2BSID                          *datatype.UTF8String               `avp:"3GPP2-BSID" json:"tGPP2BSID,omitempty"`
	UWANUserLocationInfo               *UWANUserLocationInfo              `avp:"UWAN-User-Location-Info" json:"uWANUserLocationInfo,omitempty"`
	TWANUserLocationInfo               *TWANUserLocationInfo              `avp:"TWAN-User-Location-Info" json:"tWANUserLocationInfo,omitempty"`
	SponsorIdentity                    *datatype.UTF8String               `avp:"Sponsor-Identity" json:"sponsorIdentity,omitempty"`
	ApplicationServiceProviderIdentity *datatype.UTF8String               `avp:"Application-Service-Provider-Identity" json:"applicationServiceProviderIdentity,omitempty"`
	PresenceReportingAreaInformation   []PresenceReportingAreaInformation `avp:"Presence-Reporting-Area-Information" json:"presenceReportingAreaInformation,omitempty"`
	PresenceReportingAreaStatus        *datatype.Unsigned32               `avp:"Presence-Reporting-Area-Status" json:"presenceReportingAreaStatus,omitempty"`
	UserCSGInformation                 *UserCSGInformation                `avp:"User-CSG-Information" json:"userCSGInformation,omitempty"`
	TGPPRATType                        *datatype.OctetString              `avp:"3GPP-RAT-Type" json:"tGPPRATType,omitempty"`
	RelatedChangeConditionInformation  *RelatedChangeConditionInformation `avp:"Related-Change-Condition-Information" json:"relatedChangeConditionInformation,omitempty"`
	ServingPLMNRateControl             *ServingPLMNRateControl            `avp:"Serving-PLMN-Rate-Control" json:"servingPLMNRateControl,omitempty"`
	APNRateControl                     *APNRateControl                    `avp:"APN-Rate-Control" json:"aPNRateControl,omitempty"`
	TGPPPSDataOffStatus                *TGPPPSDataOffStatus               `avp:"3GPP-PS-Data-Off-Status" json:"tGPPPSDataOffStatus,omitempty"`
	TrafficSteeringPolicyIdentifierDL  *datatype.OctetString              `avp:"Traffic-Steering-Policy-Identifier-DL" json:"trafficSteeringPolicyIdentifierDL,omitempty"`
	TrafficSteeringPolicyIdentifierUL  *datatype.OctetString              `avp:"Traffic-Steering-Policy-Identifier-UL" json:"trafficSteeringPolicyIdentifierUL,omitempty"`
	VoLTEInformation                   *VoLTEInformation                  `avp:"VoLTE-Information" json:"voLTEInformation,omitempty"`
}

type VoLTEInformation struct {
	CallingPartyAddress *datatype.UTF8String `avp:"Calling-Party-Address" json:"callingPartyAddress,omitempty"`
	CalleeInformation   *CalleeInformation   `avp:"Callee-Information" json:"calleeInformation,omitempty"`
}

type TerminalInformation struct {
	IMEI            *datatype.UTF8String `avp:"IMEI" json:"iMEI,omitempty"`
	SoftwareVersion *datatype.UTF8String `avp:"Software-Version" json:"softwareVersion,omitempty"`
}

type RRCCauseCounter struct {
	CounterValue        *datatype.Unsigned32 `avp:"Counter-Value" json:"counterValue,omitempty"`
	RRCCounterTimestamp *datatype.Time       `avp:"RRC-Counter-Timestamp" json:"rRCCounterTimestamp,omitempty"`
}

type SCSASAddress struct {
	SCSRealm   *datatype.DiameterIdentity `avp:"SCS-Realm" json:"sCSRealm,omitempty"`
	SCSAddress *datatype.Address          `avp:"SCS-Address" json:"sCSAddress,omitempty"`
}

type RANSecondaryRATUsageReport struct {
	SecondaryRATType       *datatype.OctetString `avp:"Secondary-RAT-Type" json:"secondaryRATType,omitempty"`
	RANStartTimestamp      *datatype.Time        `avp:"RAN-Start-Timestamp" json:"rANStartTimestamp,omitempty" cmp:"ignore"`
	RANEndTimestamp        *datatype.Time        `avp:"RAN-End-Timestamp" json:"rANEndTimestamp,omitempty" cmp:"ignore"`
	AccountingInputOctets  *datatype.Unsigned64  `avp:"Accounting-Input-Octets" json:"accountingInputOctets,omitempty"`
	AccountingOutputOctets *datatype.Unsigned64  `avp:"Accounting-Output-Octets" json:"accountingOutputOctets,omitempty"`
	TGPPChargingId         *datatype.OctetString `avp:"3GPP-Charging-Id" json:"tGPPChargingId,omitempty" cmp:"ignore"`
}

type IMSInformation struct {
	EventType                           *EventType                     `avp:"Event-Type" json:"eventType,omitempty"`
	RoleOfNode                          *RoleOfNode                    `avp:"Role-Of-Node" json:"roleOfNode,omitempty"`
	NodeFunctionality                   NodeFunctionality              `avp:"Node-Functionality" json:"nodeFunctionality"`
	UserSessionId                       *datatype.UTF8String           `avp:"User-Session-Id" json:"userSessionId,omitempty"`
	OutgoingSessionId                   *datatype.UTF8String           `avp:"Outgoing-Session-Id" json:"outgoingSessionId,omitempty"`
	SessionPriority                     *SessionPriority               `avp:"Session-Priority" json:"sessionPriority,omitempty"`
	CallingPartyAddress                 []datatype.UTF8String          `avp:"Calling-Party-Address" json:"callingPartyAddress,omitempty"`
	CalledPartyAddress                  *datatype.UTF8String           `avp:"Called-Party-Address" json:"calledPartyAddress,omitempty"`
	CalledAssertedIdentity              []datatype.UTF8String          `avp:"Called-Asserted-Identity" json:"calledAssertedIdentity,omitempty"`
	CalledIdentityChange                *CalledIdentityChange          `avp:"Called-Identity-Change" json:"calledIdentityChange,omitempty"`
	NumberPortabilityRoutingInformation *datatype.UTF8String           `avp:"Number-Portability-Routing-Information" json:"numberPortabilityRoutingInformation,omitempty"`
	CarrierSelectRoutingInformation     *datatype.UTF8String           `avp:"Carrier-Select-Routing-Information" json:"carrierSelectRoutingInformation,omitempty"`
	AlternateChargedPartyAddress        *datatype.UTF8String           `avp:"Alternate-Charged-Party-Address" json:"alternateChargedPartyAddress,omitempty"`
	RequestedPartyAddress               []datatype.UTF8String          `avp:"Requested-Party-Address" json:"requestedPartyAddress,omitempty"`
	AssociatedURI                       []datatype.UTF8String          `avp:"Associated-URI" json:"associatedURI,omitempty"`
	TimeStamps                          *TimeStamps                    `avp:"Time-Stamps" json:"timeStamps,omitempty"`
	ApplicationServerInformation        []ApplicationServerInformation `avp:"Application-Server-Information" json:"applicationServerInformation,omitempty"`
	InterOperatorIdentifier             []InterOperatorIdentifier      `avp:"Inter-Operator-Identifier" json:"interOperatorIdentifier,omitempty"`
	TransitIOIList                      []datatype.UTF8String          `avp:"Transit-IOI-List" json:"transitIOIList,omitempty"`
	IMSChargingIdentifier               *datatype.UTF8String           `avp:"IMS-Charging-Identifier" json:"iMSChargingIdentifier,omitempty"`
	SDPSessionDescription               []datatype.UTF8String          `avp:"SDP-Session-Description" json:"sDPSessionDescription,omitempty"`
	SDPMediaComponent                   []SDPMediaComponent            `avp:"SDP-Media-Component" json:"sDPMediaComponent,omitempty"`
	ServedPartyIPAddress                *datatype.Address              `avp:"Served-Party-IP-Address" json:"servedPartyIPAddress,omitempty"`
	ServerCapabilities                  *ServerCapabilities            `avp:"Server-Capabilities" json:"serverCapabilities,omitempty"`
	TrunkGroupID                        *TrunkGroupID                  `avp:"Trunk-Group-ID" json:"trunkGroupID,omitempty"`
	BearerService                       *datatype.OctetString          `avp:"Bearer-Service" json:"bearerService,omitempty"`
	ServiceId                           *datatype.UTF8String           `avp:"Service-Id" json:"serviceId,omitempty"`
	ServiceSpecificInfo                 []ServiceSpecificInfo          `avp:"Service-Specific-Info" json:"serviceSpecificInfo,omitempty"`
	MessageBody                         []MessageBody                  `avp:"Message-Body" json:"messageBody,omitempty"`
	CauseCode                           *datatype.Integer32            `avp:"Cause-Code" json:"causeCode,omitempty"`
	ReasonHeader                        []datatype.UTF8String          `avp:"Reason-Header" json:"reasonHeader,omitempty"`
	AccessNetworkInformation            []datatype.OctetString         `avp:"Access-Network-Information" json:"accessNetworkInformation,omitempty"`
	CellularNetworkInformation          *datatype.OctetString          `avp:"Cellular-Network-Information" json:"cellularNetworkInformation,omitempty"`
	EarlyMediaDescription               []EarlyMediaDescription        `avp:"Early-Media-Description" json:"earlyMediaDescription,omitempty"`
	IMSCommunicationServiceIdentifier   *datatype.UTF8String           `avp:"IMS-Communication-Service-Identifier" json:"iMSCommunicationServiceIdentifier,omitempty"`
	IMSApplicationReferenceIdentifier   *datatype.UTF8String           `avp:"IMS-Application-Reference-Identifier" json:"iMSApplicationReferenceIdentifier,omitempty"`
	OnlineChargingFlag                  *OnlineChargingFlag            `avp:"Online-Charging-Flag" json:"onlineChargingFlag,omitempty"`
	RealTimeTariffInformation           *RealTimeTariffInformation     `avp:"Real-Time-Tariff-Information" json:"realTimeTariffInformation,omitempty"`
	AccountExpiration                   *datatype.Time                 `avp:"Account-Expiration" json:"accountExpiration,omitempty"`
	InitialIMSChargingIdentifier        *datatype.UTF8String           `avp:"Initial-IMS-Charging-Identifier" json:"initialIMSChargingIdentifier,omitempty"`
	NNIInformation                      []NNIInformation               `avp:"NNI-Information" json:"nNIInformation,omitempty"`
	FromAddress                         *datatype.UTF8String           `avp:"From-Address" json:"fromAddress,omitempty"`
	IMSEmergencyIndicator               *IMSEmergencyIndicator         `avp:"IMS-Emergency-Indicator" json:"iMSEmergencyIndicator,omitempty"`
	IMSVisitedNetworkIdentifier         *datatype.UTF8String           `avp:"IMS-Visited-Network-Identifier" json:"iMSVisitedNetworkIdentifier,omitempty"`
	AccessNetworkInfoChange             []AccessNetworkInfoChange      `avp:"Access-Network-Info-Change" json:"accessNetworkInfoChange,omitempty"`
	AccessTransferInformation           []AccessTransferInformation    `avp:"Access-Transfer-Information" json:"accessTransferInformation,omitempty"`
	RelatedIMSChargingIdentifier        *datatype.UTF8String           `avp:"Related-IMS-Charging-Identifier" json:"relatedIMSChargingIdentifier,omitempty"`
	RelatedIMSChargingIdentifierNode    *datatype.Address              `avp:"Related-IMS-Charging-Identifier-Node" json:"relatedIMSChargingIdentifierNode,omitempty"`
	RouteHeaderReceived                 *datatype.UTF8String           `avp:"Route-Header-Received" json:"routeHeaderReceived,omitempty"`
	RouteHeaderTransmitted              *datatype.UTF8String           `avp:"Route-Header-Transmitted" json:"routeHeaderTransmitted,omitempty"`
	InstanceId                          *datatype.UTF8String           `avp:"Instance-Id" json:"instanceId,omitempty"`
	TADIdentifier                       *TADIdentifier                 `avp:"TAD-Identifier" json:"tADIdentifier,omitempty"`
	FEIdentifierList                    *datatype.UTF8String           `avp:"FE-Identifier-List" json:"fEIdentifierList,omitempty"`
}

type EventType struct {
	SIPMethod *datatype.UTF8String `avp:"SIP-Method" json:"sIPMethod,omitempty"`
	Event     *datatype.UTF8String `avp:"Event" json:"event,omitempty"`
	Expires   *datatype.Unsigned32 `avp:"Expires" json:"expires,omitempty"`
}

type CalledIdentityChange struct {
	CalledIdentity *datatype.UTF8String `avp:"Called-Identity" json:"calledIdentity,omitempty"`
	ChangeTime     *datatype.Time       `avp:"Change-Time" json:"changeTime,omitempty"`
}

type TimeStamps struct {
	SIPRequestTimestamp          *datatype.Time       `avp:"SIP-Request-Timestamp" json:"sIPRequestTimestamp,omitempty"`
	SIPResponseTimestamp         *datatype.Time       `avp:"SIP-Response-Timestamp" json:"sIPResponseTimestamp,omitempty"`
	SIPRequestTimestampFraction  *datatype.Unsigned32 `avp:"SIP-Request-Timestamp-Fraction" json:"sIPRequestTimestampFraction,omitempty"`
	SIPResponseTimestampFraction *datatype.Unsigned32 `avp:"SIP-Response-Timestamp-Fraction" json:"sIPResponseTimestampFraction,omitempty"`
}

type ApplicationServerInformation struct {
	ApplicationServer                     *datatype.UTF8String  `avp:"Application-Server" json:"applicationServer,omitempty"`
	ApplicationProvidedCalledPartyAddress []datatype.UTF8String `avp:"Application-Provided-Called-Party-Address" json:"applicationProvidedCalledPartyAddress,omitempty"`
	StatusASCode                          *StatusASCode         `avp:"Status-AS-Code" json:"statusASCode,omitempty"`
}

type InterOperatorIdentifier struct {
	OriginatingIOI *datatype.UTF8String `avp:"Originating-IOI" json:"originatingIOI,omitempty"`
	TerminatingIOI *datatype.UTF8String `avp:"Terminating-IOI" json:"terminatingIOI,omitempty"`
}

type SDPMediaComponent struct {
	SDPMediaName                         *datatype.UTF8String          `avp:"SDP-Media-Name" json:"sDPMediaName,omitempty"`
	SDPMediaDescription                  []datatype.UTF8String         `avp:"SDP-Media-Description" json:"sDPMediaDescription,omitempty"`
	LocalGWInsertedIndication            *LocalGWInsertedIndication    `avp:"Local-GW-Inserted-Indication" json:"localGWInsertedIndication,omitempty"`
	IPRealmDefaultIndication             *IPRealmDefaultIndication     `avp:"IP-Realm-Default-Indication" json:"iPRealmDefaultIndication,omitempty"`
	TranscoderInsertedIndication         *TranscoderInsertedIndication `avp:"Transcoder-Inserted-Indication" json:"transcoderInsertedIndication,omitempty"`
	MediaInitiatorFlag                   *MediaInitiatorFlag           `avp:"Media-Initiator-Flag" json:"mediaInitiatorFlag,omitempty"`
	MediaInitiatorParty                  *datatype.UTF8String          `avp:"Media-Initiator-Party" json:"mediaInitiatorParty,omitempty"`
	TGPPChargingId                       *datatype.OctetString         `avp:"3GPP-Charging-Id" json:"tGPPChargingId,omitempty"`
	AccessNetworkChargingIdentifierValue *datatype.OctetString         `avp:"Access-Network-Charging-Identifier-Value" json:"accessNetworkChargingIdentifierValue,omitempty"`
	SDPType                              *SDPType                      `avp:"SDP-Type" json:"sDPType,omitempty"`
}

type ServerCapabilities struct {
	MandatoryCapability []datatype.Unsigned32 `avp:"Mandatory-Capability" json:"mandatoryCapability,omitempty"`
	OptionalCapability  []datatype.Unsigned32 `avp:"Optional-Capability" json:"optionalCapability,omitempty"`
	ServerName          []datatype.UTF8String `avp:"Server-Name" json:"serverName,omitempty"`
}

type TrunkGroupID struct {
	IncomingTrunkGroupID *datatype.UTF8String `avp:"Incoming-Trunk-Group-ID" json:"incomingTrunkGroupID,omitempty"`
	OutgoingTrunkGroupID *datatype.UTF8String `avp:"Outgoing-Trunk-Group-ID" json:"outgoingTrunkGroupID,omitempty"`
}

type MessageBody struct {
	ContentType        datatype.UTF8String  `avp:"Content-Type" json:"contentType"`
	ContentLength      datatype.Unsigned32  `avp:"Content-Length" json:"contentLength"`
	ContentDisposition *datatype.UTF8String `avp:"Content-Disposition" json:"contentDisposition,omitempty"`
	Originator         *Originator          `avp:"Originator" json:"originator,omitempty"`
}

type EarlyMediaDescription struct {
	SDPTimeStamps         *SDPTimeStamps        `avp:"SDP-TimeStamps" json:"sDPTimeStamps,omitempty"`
	SDPMediaComponent     []SDPMediaComponent   `avp:"SDP-Media-Component" json:"sDPMediaComponent,omitempty"`
	SDPSessionDescription []datatype.UTF8String `avp:"SDP-Session-Description" json:"sDPSessionDescription,omitempty"`
}

type SDPTimeStamps struct {
	SDPOfferTimestamp  *datatype.Time `avp:"SDP-Offer-Timestamp" json:"sDPOfferTimestamp,omitempty"`
	SDPAnswerTimestamp *datatype.Time `avp:"SDP-Answer-Timestamp" json:"sDPAnswerTimestamp,omitempty"`
}

type RealTimeTariffInformation struct {
	TariffInformation *TariffInformation   `avp:"Tariff-Information" json:"tariffInformation,omitempty"`
	TariffXML         *datatype.UTF8String `avp:"Tariff-XML" json:"tariffXML,omitempty"`
}

type NNIInformation struct {
	SessionDirection     *SessionDirection `avp:"Session-Direction" json:"sessionDirection,omitempty"`
	NNIType              *NNIType          `avp:"NNI-Type" json:"nNIType,omitempty"`
	RelationshipMode     *RelationshipMode `avp:"Relationship-Mode" json:"relationshipMode,omitempty"`
	NeighbourNodeAddress *datatype.Address `avp:"Neighbour-Node-Address" json:"neighbourNodeAddress,omitempty"`
}

type AccessNetworkInfoChange struct {
	AccessNetworkInformation   []datatype.OctetString `avp:"Access-Network-Information" json:"accessNetworkInformation,omitempty"`
	CellularNetworkInformation *datatype.OctetString  `avp:"Cellular-Network-Information" json:"cellularNetworkInformation,omitempty"`
	ChangeTime                 *datatype.Time         `avp:"Change-Time" json:"changeTime,omitempty"`
}

type AccessTransferInformation struct {
	AccessTransferType               *AccessTransferType    `avp:"Access-Transfer-Type" json:"accessTransferType,omitempty"`
	AccessNetworkInformation         []datatype.OctetString `avp:"Access-Network-Information" json:"accessNetworkInformation,omitempty"`
	CellularNetworkInformation       *datatype.OctetString  `avp:"Cellular-Network-Information" json:"cellularNetworkInformation,omitempty"`
	InterUETransfer                  *InterUETransfer       `avp:"Inter-UE-Transfer" json:"interUETransfer,omitempty"`
	UserEquipmentInfo                *UserEquipmentInfo     `avp:"User-Equipment-Info" json:"userEquipmentInfo,omitempty"`
	InstanceId                       *datatype.UTF8String   `avp:"Instance-Id" json:"instanceId,omitempty"`
	RelatedIMSChargingIdentifier     *datatype.UTF8String   `avp:"Related-IMS-Charging-Identifier" json:"relatedIMSChargingIdentifier,omitempty"`
	RelatedIMSChargingIdentifierNode *datatype.Address      `avp:"Related-IMS-Charging-Identifier-Node" json:"relatedIMSChargingIdentifierNode,omitempty"`
	ChangeTime                       *datatype.Time         `avp:"Change-Time" json:"changeTime,omitempty"`
}

type MMSInformation struct {
	OriginatorAddress        *OriginatorAddress        `avp:"Originator-Address" json:"originatorAddress,omitempty"`
	RecipientAddress         []RecipientAddress        `avp:"Recipient-Address" json:"recipientAddress,omitempty"`
	SubmissionTime           *datatype.Time            `avp:"Submission-Time" json:"submissionTime,omitempty"`
	MMContentType            *MMContentType            `avp:"MM-Content-Type" json:"mMContentType,omitempty"`
	Priority                 *Priority                 `avp:"Priority" json:"priority,omitempty"`
	MessageID                *datatype.UTF8String      `avp:"Message-ID" json:"messageID,omitempty"`
	MessageType              *MessageType              `avp:"Message-Type" json:"messageType,omitempty"`
	MessageSize              *datatype.Unsigned32      `avp:"Message-Size" json:"messageSize,omitempty"`
	MessageClass             *MessageClass             `avp:"Message-Class" json:"messageClass,omitempty"`
	DeliveryReportRequested  *DeliveryReportRequested  `avp:"Delivery-Report-Requested" json:"deliveryReportRequested,omitempty"`
	ReadReplyReportRequested *ReadReplyReportRequested `avp:"Read-Reply-Report-Requested" json:"readReplyReportRequested,omitempty"`
	MMBoxStorageRequested    *MMBoxStorageRequested    `avp:"MMBox-Storage-Requested" json:"mMBoxStorageRequested,omitempty"`
	ApplicID                 *datatype.UTF8String      `avp:"Applic-ID" json:"applicID,omitempty"`
	ReplyApplicID            *datatype.UTF8String      `avp:"Reply-Applic-ID" json:"replyApplicID,omitempty"`
	AuxApplicInfo            *datatype.UTF8String      `avp:"Aux-Applic-Info" json:"auxApplicInfo,omitempty"`
	ContentClass             *ContentClass             `avp:"Content-Class" json:"contentClass,omitempty"`
	DRMContent               *DRMContent               `avp:"DRM-Content" json:"dRMContent,omitempty"`
	Adaptations              *Adaptations              `avp:"Adaptations" json:"adaptations,omitempty"`
	VASPId                   *datatype.UTF8String      `avp:"VASP-Id" json:"vASPId,omitempty"`
	VASId                    *datatype.UTF8String      `avp:"VAS-Id" json:"vASId,omitempty"`
}

type OriginatorAddress struct {
	AddressType   *AddressType         `avp:"Address-Type" json:"addressType,omitempty"`
	AddressData   *datatype.UTF8String `avp:"Address-Data" json:"addressData,omitempty"`
	AddressDomain *AddressDomain       `avp:"Address-Domain" json:"addressDomain,omitempty"`
}

type AddressDomain struct {
	DomainName     *datatype.UTF8String `avp:"Domain-Name" json:"domainName,omitempty"`
	TGPPIMSIMCCMNC *datatype.UTF8String `avp:"3GPP-IMSI-MCC-MNC" json:"tGPPIMSIMCCMNC,omitempty"`
}

type RecipientAddress struct {
	AddressType   *AddressType         `avp:"Address-Type" json:"addressType,omitempty"`
	AddressData   *datatype.UTF8String `avp:"Address-Data" json:"addressData,omitempty"`
	AddressDomain *AddressDomain       `avp:"Address-Domain" json:"addressDomain,omitempty"`
	AddresseeType *AddresseeType       `avp:"Addressee-Type" json:"addresseeType,omitempty"`
}

type MMContentType struct {
	TypeNumber                   *TypeNumber                    `avp:"Type-Number" json:"typeNumber,omitempty"`
	AdditionalTypeInformation    *datatype.UTF8String           `avp:"Additional-Type-Information" json:"additionalTypeInformation,omitempty"`
	ContentSize                  *datatype.Unsigned32           `avp:"Content-Size" json:"contentSize,omitempty"`
	AdditionalContentInformation []AdditionalContentInformation `avp:"Additional-Content-Information" json:"additionalContentInformation,omitempty"`
}

type AdditionalContentInformation struct {
	TypeNumber                *TypeNumber          `avp:"Type-Number" json:"typeNumber,omitempty"`
	AdditionalTypeInformation *datatype.UTF8String `avp:"Additional-Type-Information" json:"additionalTypeInformation,omitempty"`
	ContentSize               *datatype.Unsigned32 `avp:"Content-Size" json:"contentSize,omitempty"`
}

type MessageClass struct {
	ClassIdentifier *ClassIdentifier     `avp:"Class-Identifier" json:"classIdentifier,omitempty"`
	TokenText       *datatype.UTF8String `avp:"Token-Text" json:"tokenText,omitempty"`
}

type LCSInformation struct {
	LCSClientID      *LCSClientID          `avp:"LCS-Client-ID" json:"lCSClientID,omitempty"`
	LocationType     *LocationType         `avp:"Location-Type" json:"locationType,omitempty"`
	LocationEstimate *datatype.OctetString `avp:"Location-Estimate" json:"locationEstimate,omitempty"`
	PositioningData  *datatype.UTF8String  `avp:"Positioning-Data" json:"positioningData,omitempty"`
	TGPPIMSI         *datatype.UTF8String  `avp:"3GPP-IMSI" json:"tGPPIMSI,omitempty"`
	MSISDN           *datatype.OctetString `avp:"MSISDN" json:"mSISDN,omitempty"`
}

type LCSClientID struct {
	LCSClientType       *LCSClientType       `avp:"LCS-Client-Type" json:"lCSClientType,omitempty"`
	LCSClientExternalID *datatype.UTF8String `avp:"LCS-Client-External-ID" json:"lCSClientExternalID,omitempty"`
	LCSClientDialedByMS *datatype.UTF8String `avp:"LCS-Client-Dialed-By-MS" json:"lCSClientDialedByMS,omitempty"`
	LCSClientName       *LCSClientName       `avp:"LCS-Client-Name" json:"lCSClientName,omitempty"`
	LCSAPN              *datatype.UTF8String `avp:"LCS-APN" json:"lCSAPN,omitempty"`
	LCSRequestorID      *LCSRequestorID      `avp:"LCS-Requestor-ID" json:"lCSRequestorID,omitempty"`
}

type LCSClientName struct {
	LCSDataCodingScheme *datatype.UTF8String `avp:"LCS-Data-Coding-Scheme" json:"lCSDataCodingScheme,omitempty"`
	LCSNameString       *datatype.UTF8String `avp:"LCS-Name-String" json:"lCSNameString,omitempty"`
	LCSFormatIndicator  *LCSFormatIndicator  `avp:"LCS-Format-Indicator" json:"lCSFormatIndicator,omitempty"`
}

type LCSRequestorID struct {
	LCSDataCodingScheme  *datatype.UTF8String `avp:"LCS-Data-Coding-Scheme" json:"lCSDataCodingScheme,omitempty"`
	LCSRequestorIDString *datatype.UTF8String `avp:"LCS-Requestor-ID-String" json:"lCSRequestorIDString,omitempty"`
}

type LocationType struct {
	LocationEstimateType      *LocationEstimateType `avp:"Location-Estimate-Type" json:"locationEstimateType,omitempty"`
	DeferredLocationEventType *datatype.UTF8String  `avp:"Deferred-Location-Event-Type" json:"deferredLocationEventType,omitempty"`
}

type PoCInformation struct {
	PoCServerRole            *PoCServerRole            `avp:"PoC-Server-Role" json:"poCServerRole,omitempty"`
	PoCSessionType           *PoCSessionType           `avp:"PoC-Session-Type" json:"poCSessionType,omitempty"`
	PoCUserRole              *PoCUserRole              `avp:"PoC-User-Role" json:"poCUserRole,omitempty"`
	PoCSessionInitiationType *PoCSessionInitiationType `avp:"PoC-Session-Initiation-Type" json:"poCSessionInitiationType,omitempty"`
	PoCEventType             *PoCEventType             `avp:"PoC-Event-Type" json:"poCEventType,omitempty"`
	NumberOfParticipants     *datatype.Unsigned32      `avp:"Number-Of-Participants" json:"numberOfParticipants,omitempty"`
	ParticipantsInvolved     []datatype.UTF8String     `avp:"Participants-Involved" json:"participantsInvolved,omitempty"`
	ParticipantGroup         []ParticipantGroup        `avp:"Participant-Group" json:"participantGroup,omitempty"`
	TalkBurstExchange        []TalkBurstExchange       `avp:"Talk-Burst-Exchange" json:"talkBurstExchange,omitempty"`
	PoCControllingAddress    *datatype.UTF8String      `avp:"PoC-Controlling-Address" json:"poCControllingAddress,omitempty"`
	PoCGroupName             *datatype.UTF8String      `avp:"PoC-Group-Name" json:"poCGroupName,omitempty"`
	PoCSessionId             *datatype.UTF8String      `avp:"PoC-Session-Id" json:"poCSessionId,omitempty"`
	ChargedParty             *datatype.UTF8String      `avp:"Charged-Party" json:"chargedParty,omitempty"`
}

type PoCUserRole struct {
	PoCUserRoleIDs       *datatype.UTF8String  `avp:"PoC-User-Role-IDs" json:"poCUserRoleIDs,omitempty"`
	PoCUserRoleInfoUnits *PoCUserRoleInfoUnits `avp:"PoC-User-Role-Info-Units" json:"poCUserRoleInfoUnits,omitempty"`
}

type ParticipantGroup struct {
	CalledPartyAddress        *datatype.UTF8String       `avp:"Called-Party-Address" json:"calledPartyAddress,omitempty"`
	ParticipantAccessPriority *ParticipantAccessPriority `avp:"Participant-Access-Priority" json:"participantAccessPriority,omitempty"`
	UserParticipatingType     *UserParticipatingType     `avp:"User-Participating-Type" json:"userParticipatingType,omitempty"`
}

type TalkBurstExchange struct {
	PoCChangeTime              datatype.Time        `avp:"PoC-Change-Time" json:"poCChangeTime"`
	NumberOfTalkBursts         *datatype.Unsigned32 `avp:"Number-Of-Talk-Bursts" json:"numberOfTalkBursts,omitempty"`
	TalkBurstVolume            *datatype.Unsigned32 `avp:"Talk-Burst-Volume" json:"talkBurstVolume,omitempty"`
	TalkBurstTime              *datatype.Unsigned32 `avp:"Talk-Burst-Time" json:"talkBurstTime,omitempty"`
	NumberOfReceivedTalkBursts *datatype.Unsigned32 `avp:"Number-Of-Received-Talk-Bursts" json:"numberOfReceivedTalkBursts,omitempty"`
	ReceivedTalkBurstVolume    *datatype.Unsigned32 `avp:"Received-Talk-Burst-Volume" json:"receivedTalkBurstVolume,omitempty"`
	ReceivedTalkBurstTime      *datatype.Unsigned32 `avp:"Received-Talk-Burst-Time" json:"receivedTalkBurstTime,omitempty"`
	NumberOfParticipants       *datatype.Unsigned32 `avp:"Number-Of-Participants" json:"numberOfParticipants,omitempty"`
	PoCChangeCondition         *PoCChangeCondition  `avp:"PoC-Change-Condition" json:"poCChangeCondition,omitempty"`
}

type MBMSInformation struct {
	TMGI                           *datatype.OctetString      `avp:"TMGI" json:"tMGI,omitempty"`
	MBMSServiceType                *MBMSServiceType           `avp:"MBMS-Service-Type" json:"mBMSServiceType,omitempty"`
	MBMSUserServiceType            *MBMSUserServiceType       `avp:"MBMS-User-Service-Type" json:"mBMSUserServiceType,omitempty"`
	FileRepairSupported            *FileRepairSupported       `avp:"File-Repair-Supported" json:"fileRepairSupported,omitempty"`
	RequiredMBMSBearerCapabilities *datatype.UTF8String       `avp:"Required-MBMS-Bearer-Capabilities" json:"requiredMBMSBearerCapabilities,omitempty"`
	MBMS2G3GIndicator              *MBMS2G3GIndicator         `avp:"MBMS-2G-3G-Indicator" json:"mBMS2G3GIndicator,omitempty"`
	RAI                            *datatype.UTF8String       `avp:"RAI" json:"rAI,omitempty"`
	MBMSServiceArea                []datatype.OctetString     `avp:"MBMS-Service-Area" json:"mBMSServiceArea,omitempty"`
	MBMSSessionIdentity            *datatype.OctetString      `avp:"MBMS-Session-Identity" json:"mBMSSessionIdentity,omitempty"`
	CNIPMulticastDistribution      *CNIPMulticastDistribution `avp:"CN-IP-Multicast-Distribution" json:"cNIPMulticastDistribution,omitempty"`
	MBMSGWAddress                  *datatype.Address          `avp:"MBMS-GW-Address" json:"mBMSGWAddress,omitempty"`
	MBMSChargedParty               *MBMSChargedParty          `avp:"MBMS-Charged-Party" json:"mBMSChargedParty,omitempty"`
	MSISDN                         []datatype.OctetString     `avp:"MSISDN" json:"mSISDN,omitempty"`
	MBMSDataTransferStart          *datatype.Unsigned64       `avp:"MBMS-Data-Transfer-Start" json:"mBMSDataTransferStart,omitempty"`
	MBMSDataTransferStop           *datatype.Unsigned64       `avp:"MBMS-Data-Transfer-Stop" json:"mBMSDataTransferStop,omitempty"`
}

type SMSInformation struct {
	SMSNode                    *SMSNode                    `avp:"SMS-Node" json:"sMSNode,omitempty"`
	ClientAddress              *datatype.Address           `avp:"Client-Address" json:"clientAddress,omitempty"`
	OriginatorSCCPAddress      *datatype.Address           `avp:"Originator-SCCP-Address" json:"originatorSCCPAddress,omitempty"`
	SMSCAddress                *datatype.Address           `avp:"SMSC-Address" json:"sMSCAddress,omitempty"`
	DataCodingScheme           *datatype.Integer32         `avp:"Data-Coding-Scheme" json:"dataCodingScheme,omitempty"`
	SMDischargeTime            *datatype.Time              `avp:"SM-Discharge-Time" json:"sMDischargeTime,omitempty"`
	SMMessageType              *SMMessageType              `avp:"SM-Message-Type" json:"sMMessageType,omitempty"`
	OriginatorInterface        *OriginatorInterface        `avp:"Originator-Interface" json:"originatorInterface,omitempty"`
	SMProtocolID               *datatype.OctetString       `avp:"SM-Protocol-ID" json:"sMProtocolID,omitempty"`
	ReplyPathRequested         *ReplyPathRequested         `avp:"Reply-Path-Requested" json:"replyPathRequested,omitempty"`
	SMStatus                   *datatype.OctetString       `avp:"SM-Status" json:"sMStatus,omitempty"`
	SMUserDataHeader           *datatype.OctetString       `avp:"SM-User-Data-Header" json:"sMUserDataHeader,omitempty"`
	NumberOfMessagesSent       *datatype.Unsigned32        `avp:"Number-Of-Messages-Sent" json:"numberOfMessagesSent,omitempty"`
	SMSequenceNumber           *datatype.Unsigned32        `avp:"SM-Sequence-Number" json:"sMSequenceNumber,omitempty"`
	RecipientInfo              []RecipientInfo             `avp:"Recipient-Info" json:"recipientInfo,omitempty"`
	OriginatorReceivedAddress  *OriginatorReceivedAddress  `avp:"Originator-Received-Address" json:"originatorReceivedAddress,omitempty"`
	SMServiceType              *SMServiceType              `avp:"SM-Service-Type" json:"sMServiceType,omitempty"`
	SMSResult                  *datatype.Unsigned32        `avp:"SMS-Result" json:"sMSResult,omitempty"`
	SMDeviceTriggerIndicator   *SMDeviceTriggerIndicator   `avp:"SM-Device-Trigger-Indicator" json:"sMDeviceTriggerIndicator,omitempty"`
	SMDeviceTriggerInformation *SMDeviceTriggerInformation `avp:"SM-Device-Trigger-Information" json:"sMDeviceTriggerInformation,omitempty"`
	MTCIWFAddress              *datatype.Address           `avp:"MTC-IWF-Address" json:"mTCIWFAddress,omitempty"`
	ApplicationPortIdentifier  *datatype.Unsigned32        `avp:"Application-Port-Identifier" json:"applicationPortIdentifier,omitempty"`
	ExternalIdentifier         *datatype.UTF8String        `avp:"External-Identifier" json:"externalIdentifier,omitempty"`
}

type OriginatorInterface struct {
	InterfaceId   *datatype.UTF8String `avp:"Interface-Id" json:"interfaceId,omitempty"`
	InterfaceText *datatype.UTF8String `avp:"Interface-Text" json:"interfaceText,omitempty"`
	InterfacePort *datatype.UTF8String `avp:"Interface-Port" json:"interfacePort,omitempty"`
	InterfaceType *InterfaceType       `avp:"Interface-Type" json:"interfaceType,omitempty"`
}

type RecipientInfo struct {
	DestinationInterface     *DestinationInterface      `avp:"Destination-Interface" json:"destinationInterface,omitempty"`
	RecipientAddress         []RecipientAddress         `avp:"Recipient-Address" json:"recipientAddress,omitempty"`
	RecipientReceivedAddress []RecipientReceivedAddress `avp:"Recipient-Received-Address" json:"recipientReceivedAddress,omitempty"`
	RecipientSCCPAddress     *datatype.Address          `avp:"Recipient-SCCP-Address" json:"recipientSCCPAddress,omitempty"`
	SMProtocolID             *datatype.OctetString      `avp:"SM-Protocol-ID" json:"sMProtocolID,omitempty"`
}

type DestinationInterface struct {
	InterfaceId   *datatype.UTF8String `avp:"Interface-Id" json:"interfaceId,omitempty"`
	InterfaceText *datatype.UTF8String `avp:"Interface-Text" json:"interfaceText,omitempty"`
	InterfacePort *datatype.UTF8String `avp:"Interface-Port" json:"interfacePort,omitempty"`
	InterfaceType *InterfaceType       `avp:"Interface-Type" json:"interfaceType,omitempty"`
}

type RecipientReceivedAddress struct {
	AddressType   *AddressType         `avp:"Address-Type" json:"addressType,omitempty"`
	AddressData   *datatype.UTF8String `avp:"Address-Data" json:"addressData,omitempty"`
	AddressDomain *AddressDomain       `avp:"Address-Domain" json:"addressDomain,omitempty"`
}

type OriginatorReceivedAddress struct {
	AddressType   *AddressType         `avp:"Address-Type" json:"addressType,omitempty"`
	AddressData   *datatype.UTF8String `avp:"Address-Data" json:"addressData,omitempty"`
	AddressDomain *AddressDomain       `avp:"Address-Domain" json:"addressDomain,omitempty"`
}

type SMDeviceTriggerInformation struct {
	MTCIWFAddress             *datatype.Address    `avp:"MTC-IWF-Address" json:"mTCIWFAddress,omitempty"`
	ReferenceNumber           *datatype.Unsigned32 `avp:"Reference-Number" json:"referenceNumber,omitempty"`
	ServingNode               *ServingNode         `avp:"Serving-Node" json:"servingNode,omitempty"`
	ValidityTime              *datatype.Unsigned32 `avp:"Validity-Time" json:"validityTime,omitempty"`
	PriorityIndication        *PriorityIndication  `avp:"Priority-Indication" json:"priorityIndication,omitempty"`
	ApplicationPortIdentifier *datatype.Unsigned32 `avp:"Application-Port-Identifier" json:"applicationPortIdentifier,omitempty"`
}

type ServingNode struct {
	SGSNNumber          *datatype.OctetString      `avp:"SGSN-Number" json:"sGSNNumber,omitempty"`
	SGSNName            *datatype.DiameterIdentity `avp:"SGSN-Name" json:"sGSNName,omitempty"`
	SGSNRealm           *datatype.DiameterIdentity `avp:"SGSN-Realm" json:"sGSNRealm,omitempty"`
	MMEName             *datatype.DiameterIdentity `avp:"MME-Name" json:"mMEName,omitempty"`
	MMERealm            *datatype.DiameterIdentity `avp:"MME-Realm" json:"mMERealm,omitempty"`
	MSCNumber           *datatype.OctetString      `avp:"MSC-Number" json:"mSCNumber,omitempty"`
	TGPPAAAServerName   *datatype.DiameterIdentity `avp:"3GPP-AAA-Server-Name" json:"tGPPAAAServerName,omitempty"`
	LCSCapabilitiesSets *datatype.Unsigned32       `avp:"LCS-Capabilities-Sets" json:"lCSCapabilitiesSets,omitempty"`
	GMLCAddress         *datatype.Address          `avp:"GMLC-Address" json:"gMLCAddress,omitempty"`
}

type VCSInformation struct {
	BearerCapability           *datatype.OctetString `avp:"Bearer-Capability" json:"bearerCapability,omitempty"`
	NetworkCallReferenceNumber *datatype.OctetString `avp:"Network-Call-Reference-Number" json:"networkCallReferenceNumber,omitempty"`
	MSCAddress                 *datatype.OctetString `avp:"MSC-Address" json:"mSCAddress,omitempty"`
	BasicServiceCode           *BasicServiceCode     `avp:"Basic-Service-Code" json:"basicServiceCode,omitempty"`
	ISUPLocationNumber         *datatype.OctetString `avp:"ISUP-Location-Number" json:"iSUPLocationNumber,omitempty"`
	VLRNumber                  *datatype.OctetString `avp:"VLR-Number" json:"vLRNumber,omitempty"`
	ForwardingPending          *ForwardingPending    `avp:"Forwarding-Pending" json:"forwardingPending,omitempty"`
	ISUPCause                  *ISUPCause            `avp:"ISUP-Cause" json:"iSUPCause,omitempty"`
	StartTime                  *datatype.Time        `avp:"Start-Time" json:"startTime,omitempty"`
	StartofCharging            *datatype.Time        `avp:"Start-of-Charging" json:"startofCharging,omitempty"`
	StopTime                   *datatype.Time        `avp:"Stop-Time" json:"stopTime,omitempty"`
	PSFreeFormatData           *datatype.OctetString `avp:"PS-Free-Format-Data" json:"pSFreeFormatData,omitempty"`
}

type BasicServiceCode struct {
	BearerService *datatype.OctetString `avp:"Bearer-Service" json:"bearerService,omitempty"`
	Teleservice   *datatype.OctetString `avp:"Teleservice" json:"teleservice,omitempty"`
}

type ISUPCause struct {
	ISUPCauseLocation    *datatype.Unsigned32  `avp:"ISUP-Cause-Location" json:"iSUPCauseLocation,omitempty"`
	ISUPCauseValue       *datatype.Unsigned32  `avp:"ISUP-Cause-Value" json:"iSUPCauseValue,omitempty"`
	ISUPCauseDiagnostics *datatype.OctetString `avp:"ISUP-Cause-Diagnostics" json:"iSUPCauseDiagnostics,omitempty"`
}

type MMTelInformation struct {
	SupplementaryService []SupplementaryService `avp:"Supplementary-Service" json:"supplementaryService,omitempty"`
}

type SupplementaryService struct {
	MMTelSServiceType      *datatype.Unsigned32   `avp:"MMTel-SService-Type" json:"mMTelSServiceType,omitempty"`
	ServiceMode            *datatype.Unsigned32   `avp:"Service-Mode" json:"serviceMode,omitempty"`
	NumberOfDiversions     *datatype.Unsigned32   `avp:"Number-Of-Diversions" json:"numberOfDiversions,omitempty"`
	AssociatedPartyAddress *datatype.UTF8String   `avp:"Associated-Party-Address" json:"associatedPartyAddress,omitempty"`
	ServiceID              *datatype.UTF8String   `avp:"Service-ID" json:"serviceID,omitempty"`
	ChangeTime             *datatype.Time         `avp:"Change-Time" json:"changeTime,omitempty"`
	NumberOfParticipants   *datatype.Unsigned32   `avp:"Number-Of-Participants" json:"numberOfParticipants,omitempty"`
	ParticipantActionType  *ParticipantActionType `avp:"Participant-Action-Type" json:"participantActionType,omitempty"`
	CUGInformation         *datatype.OctetString  `avp:"CUG-Information" json:"cUGInformation,omitempty"`
	AoCInformation         *AoCInformation        `avp:"AoC-Information" json:"aoCInformation,omitempty"`
}

type ProSeInformation struct {
	SupportedFeatures                                 []SupportedFeatures                                 `avp:"Supported-Features" json:"supportedFeatures,omitempty"`
	AnnouncingUEHPLMNIdentifier                       *datatype.UTF8String                                `avp:"Announcing-UE-HPLMN-Identifier" json:"announcingUEHPLMNIdentifier,omitempty"`
	AnnouncingUEVPLMNIdentifier                       *datatype.UTF8String                                `avp:"Announcing-UE-VPLMN-Identifier" json:"announcingUEVPLMNIdentifier,omitempty"`
	MonitoringUEHPLMNIdentifier                       *datatype.UTF8String                                `avp:"Monitoring-UE-HPLMN-Identifier" json:"monitoringUEHPLMNIdentifier,omitempty"`
	MonitoringUEVPLMNIdentifier                       *datatype.UTF8String                                `avp:"Monitoring-UE-VPLMN-Identifier" json:"monitoringUEVPLMNIdentifier,omitempty"`
	MonitoredPLMNIdentifier                           *datatype.UTF8String                                `avp:"Monitored-PLMN-Identifier" json:"monitoredPLMNIdentifier,omitempty"`
	RoleOfProSeFunction                               *RoleOfProSeFunction                                `avp:"Role-Of-ProSe-Function" json:"roleOfProSeFunction,omitempty"`
	ProSeAppId                                        *datatype.UTF8String                                `avp:"ProSe-App-Id" json:"proSeAppId,omitempty"`
	ProSe3rdPartyApplicationID                        *datatype.UTF8String                                `avp:"ProSe-3rd-Party-Application-ID" json:"proSe3rdPartyApplicationID,omitempty"`
	ApplicationSpecificData                           *datatype.OctetString                               `avp:"Application-Specific-Data" json:"applicationSpecificData,omitempty"`
	ProSeEventType                                    *ProSeEventType                                     `avp:"ProSe-Event-Type" json:"proSeEventType,omitempty"`
	ProSeDirectDiscoveryModel                         *ProSeDirectDiscoveryModel                          `avp:"ProSe-Direct-Discovery-Model" json:"proSeDirectDiscoveryModel,omitempty"`
	ProSeFunctionIPAddress                            *datatype.Address                                   `avp:"ProSe-Function-IP-Address" json:"proSeFunctionIPAddress,omitempty"`
	ProSeFunctionID                                   *datatype.OctetString                               `avp:"ProSe-Function-ID" json:"proSeFunctionID,omitempty"`
	ProSeValidityTimer                                *datatype.Unsigned32                                `avp:"ProSe-Validity-Timer" json:"proSeValidityTimer,omitempty"`
	ProSeRoleOfUE                                     *ProSeRoleOfUE                                      `avp:"ProSe-Role-Of-UE" json:"proSeRoleOfUE,omitempty"`
	ProSeRequestTimestamp                             *datatype.Time                                      `avp:"ProSe-Request-Timestamp" json:"proSeRequestTimestamp,omitempty"`
	PC3ControlProtocolCause                           *datatype.Integer32                                 `avp:"PC3-Control-Protocol-Cause" json:"pC3ControlProtocolCause,omitempty"`
	MonitoringUEIdentifier                            *datatype.UTF8String                                `avp:"Monitoring-UE-Identifier" json:"monitoringUEIdentifier,omitempty"`
	ProSeFunctionPLMNIdentifier                       *datatype.UTF8String                                `avp:"ProSe-Function-PLMN-Identifier" json:"proSeFunctionPLMNIdentifier,omitempty"`
	RequestorPLMNIdentifier                           *datatype.UTF8String                                `avp:"Requestor-PLMN-Identifier" json:"requestorPLMNIdentifier,omitempty"`
	OriginAppLayerUserId                              *datatype.UTF8String                                `avp:"Origin-App-Layer-User-Id" json:"originAppLayerUserId,omitempty"`
	WLANLinkLayerId                                   *datatype.OctetString                               `avp:"WLAN-Link-Layer-Id" json:"wLANLinkLayerId,omitempty"`
	RequestingEPUID                                   *datatype.UTF8String                                `avp:"Requesting-EPUID" json:"requestingEPUID,omitempty"`
	TargetAppLayerUserId                              *datatype.UTF8String                                `avp:"Target-App-Layer-User-Id" json:"targetAppLayerUserId,omitempty"`
	RequestedPLMNIdentifier                           *datatype.UTF8String                                `avp:"Requested-PLMN-Identifier" json:"requestedPLMNIdentifier,omitempty"`
	TimeWindow                                        *datatype.Unsigned32                                `avp:"Time-Window" json:"timeWindow,omitempty"`
	ProSeRangeClass                                   *ProSeRangeClass                                    `avp:"ProSe-Range-Class" json:"proSeRangeClass,omitempty"`
	ProximityAlertIndication                          *ProximityAlertIndication                           `avp:"Proximity-Alert-Indication" json:"proximityAlertIndication,omitempty"`
	ProximityAlertTimestamp                           *datatype.Time                                      `avp:"Proximity-Alert-Timestamp" json:"proximityAlertTimestamp,omitempty"`
	ProximityCancellationTimestamp                    *datatype.Time                                      `avp:"Proximity-Cancellation-Timestamp" json:"proximityCancellationTimestamp,omitempty"`
	ProSeReasonForCancellation                        *ProSeReasonForCancellation                         `avp:"ProSe-Reason-For-Cancellation" json:"proSeReasonForCancellation,omitempty"`
	PC3EPCControlProtocolCause                        *datatype.Integer32                                 `avp:"PC3-EPC-Control-Protocol-Cause" json:"pC3EPCControlProtocolCause,omitempty"`
	ProSeUEID                                         *datatype.OctetString                               `avp:"ProSe-UE-ID" json:"proSeUEID,omitempty"`
	ProSeSourceIPAddress                              *datatype.Address                                   `avp:"ProSe-Source-IP-Address" json:"proSeSourceIPAddress,omitempty"`
	Layer2GroupID                                     *datatype.OctetString                               `avp:"Layer-2-Group-ID" json:"layer2GroupID,omitempty"`
	ProSeGroupIPMulticastAddress                      *datatype.Address                                   `avp:"ProSe-Group-IP-Multicast-Address" json:"proSeGroupIPMulticastAddress,omitempty"`
	CoverageInfo                                      []CoverageInfo                                      `avp:"Coverage-Info" json:"coverageInfo,omitempty"`
	RadioParameterSetInfo                             []RadioParameterSetInfo                             `avp:"Radio-Parameter-Set-Info" json:"radioParameterSetInfo,omitempty"`
	TransmitterInfo                                   []TransmitterInfo                                   `avp:"Transmitter-Info" json:"transmitterInfo,omitempty"`
	TimeFirstTransmission                             *datatype.Time                                      `avp:"Time-First-Transmission" json:"timeFirstTransmission,omitempty"`
	TimeFirstReception                                *datatype.Time                                      `avp:"Time-First-Reception" json:"timeFirstReception,omitempty"`
	ProSeDirectCommunicationTransmissionDataContainer []ProSeDirectCommunicationTransmissionDataContainer `avp:"ProSe-Direct-Communication-Transmission-Data-Container" json:"proSeDirectCommunicationTransmissionDataContainer,omitempty"`
	ProSeDirectCommunicationReceptionDataContainer    []ProSeDirectCommunicationReceptionDataContainer    `avp:"ProSe-Direct-Communication-Reception-Data-Container" json:"proSeDirectCommunicationReceptionDataContainer,omitempty"`
	AnnouncingPLMNID                                  *datatype.UTF8String                                `avp:"Announcing-PLMN-ID" json:"announcingPLMNID,omitempty"`
	ProSeTargetLayer2ID                               *datatype.OctetString                               `avp:"ProSe-Target-Layer-2-ID" json:"proSeTargetLayer2ID,omitempty"`
	RelayIPaddress                                    *datatype.Address                                   `avp:"Relay-IP-address" json:"relayIPaddress,omitempty"`
	ProSeUEtoNetworkRelayUEID                         *datatype.OctetString                               `avp:"ProSe-UE-to-Network-Relay-UE-ID" json:"proSeUEtoNetworkRelayUEID,omitempty"`
	TargetIPAddress                                   *datatype.Address                                   `avp:"Target-IP-Address" json:"targetIPAddress,omitempty"`
	PC5RadioTechnology                                *PC5RadioTechnology                                 `avp:"PC5-Radio-Technology" json:"pC5RadioTechnology,omitempty"`
}

type CoverageInfo struct {
	CoverageStatus *CoverageStatus `avp:"Coverage-Status" json:"coverageStatus,omitempty"`
	ChangeTime     *datatype.Time  `avp:"Change-Time" json:"changeTime,omitempty"`
	LocationInfo   []LocationInfo  `avp:"Location-Info" json:"locationInfo,omitempty"`
}

type LocationInfo struct {
	TGPPUserLocationInfo *datatype.OctetString `avp:"3GPP-User-Location-Info" json:"tGPPUserLocationInfo,omitempty"`
	ChangeTime           *datatype.Time        `avp:"Change-Time" json:"changeTime,omitempty"`
}

type RadioParameterSetInfo struct {
	RadioParameterSetValues *datatype.OctetString `avp:"Radio-Parameter-Set-Values" json:"radioParameterSetValues,omitempty"`
	ChangeTime              *datatype.Time        `avp:"Change-Time" json:"changeTime,omitempty"`
}

type TransmitterInfo struct {
	ProSeSourceIPAddress *datatype.Address     `avp:"ProSe-Source-IP-Address" json:"proSeSourceIPAddress,omitempty"`
	ProSeUEID            *datatype.OctetString `avp:"ProSe-UE-ID" json:"proSeUEID,omitempty"`
}

type ProSeDirectCommunicationTransmissionDataContainer struct {
	LocalSequenceNumber                  *datatype.Unsigned32  `avp:"Local-Sequence-Number" json:"localSequenceNumber,omitempty"`
	CoverageStatus                       *CoverageStatus       `avp:"Coverage-Status" json:"coverageStatus,omitempty"`
	TGPPUserLocationInfo                 *datatype.OctetString `avp:"3GPP-User-Location-Info" json:"tGPPUserLocationInfo,omitempty"`
	AccountingOutputOctets               *datatype.Unsigned64  `avp:"Accounting-Output-Octets" json:"accountingOutputOctets,omitempty"`
	ChangeTime                           *datatype.Time        `avp:"Change-Time" json:"changeTime,omitempty"`
	ChangeCondition                      *datatype.Integer32   `avp:"Change-Condition" json:"changeCondition,omitempty"`
	VisitedPLMNId                        *datatype.OctetString `avp:"Visited-PLMN-Id" json:"visitedPLMNId,omitempty"`
	UsageInformationReportSequenceNumber *datatype.Integer32   `avp:"Usage-Information-Report-Sequence-Number" json:"usageInformationReportSequenceNumber,omitempty"`
	RadioResourcesIndicator              *datatype.Integer32   `avp:"Radio-Resources-Indicator" json:"radioResourcesIndicator,omitempty"`
	RadioFrequency                       *datatype.OctetString `avp:"Radio-Frequency" json:"radioFrequency,omitempty"`
}

type ProSeDirectCommunicationReceptionDataContainer struct {
	LocalSequenceNumber                  *datatype.Unsigned32  `avp:"Local-Sequence-Number" json:"localSequenceNumber,omitempty"`
	CoverageStatus                       *CoverageStatus       `avp:"Coverage-Status" json:"coverageStatus,omitempty"`
	TGPPUserLocationInfo                 *datatype.OctetString `avp:"3GPP-User-Location-Info" json:"tGPPUserLocationInfo,omitempty"`
	AccountingInputOctets                *datatype.Unsigned64  `avp:"Accounting-Input-Octets" json:"accountingInputOctets,omitempty"`
	ChangeTime                           *datatype.Time        `avp:"Change-Time" json:"changeTime,omitempty"`
	ChangeCondition                      *datatype.Integer32   `avp:"Change-Condition" json:"changeCondition,omitempty"`
	VisitedPLMNId                        *datatype.OctetString `avp:"Visited-PLMN-Id" json:"visitedPLMNId,omitempty"`
	UsageInformationReportSequenceNumber *datatype.Integer32   `avp:"Usage-Information-Report-Sequence-Number" json:"usageInformationReportSequenceNumber,omitempty"`
	RadioResourcesIndicator              *datatype.Integer32   `avp:"Radio-Resources-Indicator" json:"radioResourcesIndicator,omitempty"`
	RadioFrequency                       *datatype.OctetString `avp:"Radio-Frequency" json:"radioFrequency,omitempty"`
}

type ServiceGenericInformation struct {
	ApplicationServerID    *datatype.UTF8String    `avp:"Application-Server-ID" json:"applicationServerID,omitempty"`
	ApplicationServiceType *ApplicationServiceType `avp:"Application-Service-Type" json:"applicationServiceType,omitempty"`
	ApplicationSessionID   *datatype.Unsigned32    `avp:"Application-Session-ID" json:"applicationSessionID,omitempty"`
	DeliveryStatus         *datatype.UTF8String    `avp:"Delivery-Status" json:"deliveryStatus,omitempty"`
}

type IMInformation struct {
	TotalNumberOfMessagesSent            *datatype.Unsigned32 `avp:"Total-Number-Of-Messages-Sent" json:"totalNumberOfMessagesSent,omitempty"`
	TotalNumberOfMessagesExploded        *datatype.Unsigned32 `avp:"Total-Number-Of-Messages-Exploded" json:"totalNumberOfMessagesExploded,omitempty"`
	NumberOfMessagesSuccessfullySent     *datatype.Unsigned32 `avp:"Number-Of-Messages-Successfully-Sent" json:"numberOfMessagesSuccessfullySent,omitempty"`
	NumberOfMessagesSuccessfullyExploded *datatype.Unsigned32 `avp:"Number-Of-Messages-Successfully-Exploded" json:"numberOfMessagesSuccessfullyExploded,omitempty"`
}

type DCDInformation struct {
	ContentID         *datatype.UTF8String `avp:"Content-ID" json:"contentID,omitempty"`
	ContentProviderID *datatype.UTF8String `avp:"Content-Provider-ID" json:"contentProviderID,omitempty"`
}

type M2MInformation struct {
	ApplicationEntityID     *datatype.UTF8String `avp:"Application-Entity-ID" json:"applicationEntityID,omitempty"`
	ExternalID              *datatype.UTF8String `avp:"External-ID" json:"externalID,omitempty"`
	Receiver                *datatype.UTF8String `avp:"Receiver" json:"receiver,omitempty"`
	Originator              *Originator          `avp:"Originator" json:"originator,omitempty"`
	HostingCSEID            *datatype.UTF8String `avp:"Hosting-CSE-ID" json:"hostingCSEID,omitempty"`
	TargetID                *datatype.UTF8String `avp:"Target-ID" json:"targetID,omitempty"`
	ProtocolType            *ProtocolType        `avp:"Protocol-Type" json:"protocolType,omitempty"`
	RequestOperation        *RequestOperation    `avp:"Request-Operation" json:"requestOperation,omitempty"`
	RequestHeadersSize      *datatype.Unsigned32 `avp:"Request-Headers-Size" json:"requestHeadersSize,omitempty"`
	RequestBodySize         *datatype.Unsigned32 `avp:"Request-Body-Size" json:"requestBodySize,omitempty"`
	ResponseHeadersSize     *datatype.Unsigned32 `avp:"Response-Headers-Size" json:"responseHeadersSize,omitempty"`
	ResponseBodySize        *datatype.Unsigned32 `avp:"Response-Body-Size" json:"responseBodySize,omitempty"`
	ResponseStatusCode      *ResponseStatusCode  `avp:"Response-Status-Code" json:"responseStatusCode,omitempty"`
	RatingGroup             *datatype.Unsigned32 `avp:"Rating-Group" json:"ratingGroup,omitempty"`
	M2MEventRecordTimestamp *datatype.Time       `avp:"M2M-Event-Record-Timestamp" json:"m2MEventRecordTimestamp,omitempty"`
	ControlMemorySize       *datatype.Unsigned32 `avp:"Control-Memory-Size" json:"controlMemorySize,omitempty"`
	DataMemorySize          *datatype.Unsigned32 `avp:"Data-Memory-Size" json:"dataMemorySize,omitempty"`
	AccessNetworkIdentifier *datatype.Unsigned32 `avp:"Access-Network-Identifier" json:"accessNetworkIdentifier,omitempty"`
	Occupancy               *datatype.Unsigned32 `avp:"Occupancy" json:"occupancy,omitempty"`
	GroupName               *datatype.UTF8String `avp:"Group-Name" json:"groupName,omitempty"`
	MaximumNumberMembers    *datatype.Unsigned32 `avp:"Maximum-Number-Members" json:"maximumNumberMembers,omitempty"`
	CurrentNumberMembers    *datatype.Unsigned32 `avp:"Current-Number-Members" json:"currentNumberMembers,omitempty"`
	SubgroupName            *datatype.UTF8String `avp:"Subgroup-Name" json:"subgroupName,omitempty"`
	NodeId                  *datatype.UTF8String `avp:"Node-Id" json:"nodeId,omitempty"`
}

type CPDTInformation struct {
	ExternalIdentifier  *datatype.UTF8String       `avp:"External-Identifier" json:"externalIdentifier,omitempty"`
	SCEFID              *datatype.DiameterIdentity `avp:"SCEF-ID" json:"sCEFID,omitempty"`
	ServingNodeIdentity *datatype.DiameterIdentity `avp:"Serving-Node-Identity" json:"servingNodeIdentity,omitempty"`
	SGWChange           *SGWChange                 `avp:"SGW-Change" json:"sGWChange,omitempty"`
	NIDDSubmission      *NIDDSubmission            `avp:"NIDD-Submission" json:"nIDDSubmission,omitempty"`
}

type NIDDSubmission struct {
	// can not find the definition of AVP Submission-Timestamp.
	// SubmissionTimestamp    *datatype.Time       `avp:"Submission-Timestamp" json:"submissionTimestamp,omitempty"`
	EventTimestamp         *datatype.Time       `avp:"Event-Timestamp" json:"eventTimestamp,omitempty"`
	AccountingInputOctets  *datatype.Unsigned64 `avp:"Accounting-Input-Octets" json:"accountingInputOctets,omitempty"`
	AccountingOutputOctets *datatype.Unsigned64 `avp:"Accounting-Output-Octets" json:"accountingOutputOctets,omitempty"`
	ChangeCondition        *datatype.Integer32  `avp:"Change-Condition" json:"changeCondition,omitempty"`
}

type AoCRequestType datatype.Enumerated

const (
	AoCRequestType_AOC_NOT_REQUESTED AoCRequestType = 0
	AoCRequestType_AOC_FULL          AoCRequestType = 1
	AoCRequestType_AOC_COST_ONLY     AoCRequestType = 2
	AoCRequestType_AOC_TARIFF_ONLY   AoCRequestType = 3
)

type RequestedAction datatype.Enumerated

const (
	RequestedAction_DIRECT_DEBITING RequestedAction = 0
	RequestedAction_REFUND_ACCOUNT  RequestedAction = 1
	RequestedAction_CHECK_BALANCE   RequestedAction = 2
	RequestedAction_PRICE_ENQUIRY   RequestedAction = 3
)

type MultipleServicesIndicator datatype.Enumerated

const (
	MultipleServicesIndicator_MULTIPLE_SERVICES_NOT_SUPPORTED MultipleServicesIndicator = 0
	MultipleServicesIndicator_MULTIPLE_SERVICES_SUPPORTED     MultipleServicesIndicator = 1
)

type ReportingReason datatype.Enumerated

const (
	ReportingReason_THRESHOLD               ReportingReason = 0
	ReportingReason_QHT                     ReportingReason = 1
	ReportingReason_FINAL                   ReportingReason = 2
	ReportingReason_QUOTA_EXHAUSTED         ReportingReason = 3
	ReportingReason_VALIDITY_TIME           ReportingReason = 4
	ReportingReason_OTHER_QUOTA_TYPE        ReportingReason = 5
	ReportingReason_RATING_CONDITION_CHANGE ReportingReason = 6
	ReportingReason_FORCED_REAUTHORISATION  ReportingReason = 7
	ReportingReason_POOL_EXHAUSTED          ReportingReason = 8
	ReportingReason_UNUSED_QUOTA_TIMER      ReportingReason = 9
)

type EnvelopeReporting datatype.Enumerated

const (
	EnvelopeReporting_DO_NOT_REPORT_ENVELOPES                 EnvelopeReporting = 0
	EnvelopeReporting_REPORT_ENVELOPES                        EnvelopeReporting = 1
	EnvelopeReporting_REPORT_ENVELOPES_WITH_VOLUME            EnvelopeReporting = 2
	EnvelopeReporting_REPORT_ENVELOPES_WITH_EVENTS            EnvelopeReporting = 3
	EnvelopeReporting_REPORT_ENVELOPES_WITH_VOLUME_AND_EVENTS EnvelopeReporting = 4
)

type CCUnitType datatype.Enumerated

const (
	CCUnitType_TIME                   CCUnitType = 0
	CCUnitType_MONEY                  CCUnitType = 1
	CCUnitType_TOTAL_OCTETS           CCUnitType = 2
	CCUnitType_INPUT_OCTETS           CCUnitType = 3
	CCUnitType_OUTPUT_OCTETS          CCUnitType = 4
	CCUnitType_SERVICE_SPECIFIC_UNITS CCUnitType = 5
)

type TriggerType datatype.Enumerated

const (
	TriggerType_CHANGE_IN_SGSN_IP_ADDRESS                          TriggerType = 1
	TriggerType_CHANGE_IN_QOS                                      TriggerType = 2
	TriggerType_CHANGE_IN_LOCATION                                 TriggerType = 3
	TriggerType_CHANGE_IN_RAT                                      TriggerType = 4
	TriggerType_CHANGE_IN_UE_TIMEZONE                              TriggerType = 5
	TriggerType_CHANGEINQOS_TRAFFIC_CLASS                          TriggerType = 10
	TriggerType_CHANGEINQOS_RELIABILITY_CLASS                      TriggerType = 11
	TriggerType_CHANGEINQOS_DELAY_CLASS                            TriggerType = 12
	TriggerType_CHANGEINQOS_PEAK_THROUGHPUT                        TriggerType = 13
	TriggerType_CHANGEINQOS_PRECEDENCE_CLASS                       TriggerType = 14
	TriggerType_CHANGEINQOS_MEAN_THROUGHPUT                        TriggerType = 15
	TriggerType_CHANGEINQOS_MAXIMUM_BIT_RATE_FOR_UPLINK            TriggerType = 16
	TriggerType_CHANGEINQOS_MAXIMUM_BIT_RATE_FOR_DOWNLINK          TriggerType = 17
	TriggerType_CHANGEINQOS_RESIDUAL_BER                           TriggerType = 18
	TriggerType_CHANGEINQOS_SDU_ERROR_RATIO                        TriggerType = 19
	TriggerType_CHANGEINQOS_TRANSFER_DELAY                         TriggerType = 20
	TriggerType_CHANGEINQOS_TRAFFIC_HANDLING_PRIORITY              TriggerType = 21
	TriggerType_CHANGEINQOS_GUARANTEED_BIT_RATE_FOR_UPLINK         TriggerType = 22
	TriggerType_CHANGEINQOS_GUARANTEED_BIT_RATE_FOR_DOWNLINK       TriggerType = 23
	TriggerType_CHANGEINQOS_APN_AGGREGATE_MAXIMUM_BIT_RATE         TriggerType = 24
	TriggerType_CHANGEINLOCATION_MCC                               TriggerType = 30
	TriggerType_CHANGEINLOCATION_MNC                               TriggerType = 31
	TriggerType_CHANGEINLOCATION_RAC                               TriggerType = 32
	TriggerType_CHANGEINLOCATION_LAC                               TriggerType = 33
	TriggerType_CHANGEINLOCATION_CELLID                            TriggerType = 34
	TriggerType_CHANGEINLOCATION_TAC                               TriggerType = 35
	TriggerType_CHANGEINLOCATION_ECGI                              TriggerType = 36
	TriggerType_CHANGE_IN_MEDIA_COMPOSITION                        TriggerType = 40
	TriggerType_CHANGE_IN_PARTICIPANTS_NMB                         TriggerType = 50
	TriggerType_CHANGE_IN_THRSHLD_OF_PARTICIPANTS_NMB              TriggerType = 51
	TriggerType_CHANGE_IN_USER_PARTICIPATING_TYPE                  TriggerType = 52
	TriggerType_CHANGE_IN_SERVICE_CONDITION                        TriggerType = 60
	TriggerType_CHANGE_IN_SERVING_NODE                             TriggerType = 61
	TriggerType_CHANGE_IN_ACCESS_FOR_A_SERVICE_DATA_FLOW           TriggerType = 62
	TriggerType_CHANGE_IN_USER_CSG_INFORMATION                     TriggerType = 70
	TriggerType_CHANGE_IN_HYBRID_SUBSCRIBED_USER_CSG_INFORMATION   TriggerType = 71
	TriggerType_CHANGE_IN_HYBRID_UNSUBSCRIBED_USER_CSG_INFORMATION TriggerType = 72
	TriggerType_CHANGE_OF_UE_PRESENCE_IN_PRESENCE_REPORTING_AREA   TriggerType = 73
	TriggerType_CHANGE_IN_APN_RATE_CONTROL                         TriggerType = 75
	TriggerType_CHANGE_IN_3GPP_PS_DATA_OFF                         TriggerType = 76
)

type TimeQuotaType datatype.Enumerated

const (
	TimeQuotaType_DISCRETE_TIME_PERIOD   TimeQuotaType = 0
	TimeQuotaType_CONTINUOUS_TIME_PERIOD TimeQuotaType = 1
)

type QuotaIndicator datatype.Enumerated

const (
	QuotaIndicator_QUOTA_IS_NOT_USED_DURING_PLAYBACK QuotaIndicator = 0
	QuotaIndicator_QUOTA_IS_USED_DURING_PLAYBACK     QuotaIndicator = 1
)

type PlayAlternative datatype.Enumerated

const (
	PlayAlternative_SERVED_PARTY PlayAlternative = 0
	PlayAlternative_REMOTE_PARTY PlayAlternative = 1
)

type PrivacyIndicator datatype.Enumerated

const (
	PrivacyIndicator_NOT_PRIVATE PrivacyIndicator = 0
	PrivacyIndicator_PRIVATE     PrivacyIndicator = 1
)

type ChargeReasonCode datatype.Enumerated

const (
	ChargeReasonCode_UNKNOWN                      ChargeReasonCode = 0
	ChargeReasonCode_USAGE                        ChargeReasonCode = 1
	ChargeReasonCode_COMMUNICATION_ATTEMPT_CHARGE ChargeReasonCode = 2
	ChargeReasonCode_SETUP_CHARGE                 ChargeReasonCode = 3
	ChargeReasonCode_ADD_ON_CHARGE                ChargeReasonCode = 4
)

type AoCFormat datatype.Enumerated

const (
	AoCFormat_MONETARY     AoCFormat = 0
	AoCFormat_NON_MONETARY AoCFormat = 1
	AoCFormat_CAI          AoCFormat = 2
)

type AoCServiceObligatoryType datatype.Enumerated

const (
	AoCServiceObligatoryType_NON_BINDING AoCServiceObligatoryType = 0
	AoCServiceObligatoryType_BINDING     AoCServiceObligatoryType = 1
)

type AoCServiceType datatype.Enumerated

const (
	AoCServiceType_NONE  AoCServiceType = 0
	AoCServiceType_AOC_S AoCServiceType = 1
	AoCServiceType_AOC_D AoCServiceType = 2
	AoCServiceType_AOC_E AoCServiceType = 3
)

type TGPPPDPType datatype.Enumerated

const (
	TGPPPDPType_IPV4         TGPPPDPType = 0
	TGPPPDPType_PPP          TGPPPDPType = 1
	TGPPPDPType_IPV6         TGPPPDPType = 2
	TGPPPDPType_IPV4V6       TGPPPDPType = 3
	TGPPPDPType_NON_IP       TGPPPDPType = 4
	TGPPPDPType_UNSTRUCTURED TGPPPDPType = 5
	TGPPPDPType_ETHERNET     TGPPPDPType = 6
)

type ServingNodeType datatype.Enumerated

const (
	ServingNodeType_SGSN    ServingNodeType = 0
	ServingNodeType_PMIPSGW ServingNodeType = 1
	ServingNodeType_GTPSGW  ServingNodeType = 2
	ServingNodeType_EPDG    ServingNodeType = 3
	ServingNodeType_HSGW    ServingNodeType = 4
	ServingNodeType_MME     ServingNodeType = 5
	ServingNodeType_TWAN    ServingNodeType = 6
)

type SGWChange datatype.Enumerated

const (
	SGWChange_ACR_START_NOT_DUE_TO_SGW_CHANGE SGWChange = 0
	SGWChange_ACR_START_DUE_TO_SGW_CHANGE     SGWChange = 1
)

type IMSIUnauthenticatedFlag datatype.Enumerated

const (
	IMSIUnauthenticatedFlag_AUTHENTICATED   IMSIUnauthenticatedFlag = 0
	IMSIUnauthenticatedFlag_UNAUTHENTICATED IMSIUnauthenticatedFlag = 1
)

type ChargingCharacteristicsSelectionMode datatype.Enumerated

const (
	ChargingCharacteristicsSelectionMode_SERVING_NODE_SUPPLIED ChargingCharacteristicsSelectionMode = 0
	ChargingCharacteristicsSelectionMode_SUBSCRIPTION_SPECIFIC ChargingCharacteristicsSelectionMode = 1
	ChargingCharacteristicsSelectionMode_APN_SPECIFIC          ChargingCharacteristicsSelectionMode = 2
	ChargingCharacteristicsSelectionMode_HOME_DEFAULT          ChargingCharacteristicsSelectionMode = 3
	ChargingCharacteristicsSelectionMode_ROAMING_DEFAULT       ChargingCharacteristicsSelectionMode = 4
	ChargingCharacteristicsSelectionMode_VISITING_DEFAULT      ChargingCharacteristicsSelectionMode = 5
)

type PDPContextType datatype.Enumerated

const (
	PDPContextType_PRIMARY   PDPContextType = 0
	PDPContextType_SECONDARY PDPContextType = 1
)

type LowPriorityIndicator datatype.Enumerated

const (
	LowPriorityIndicator_NO  LowPriorityIndicator = 0
	LowPriorityIndicator_YES LowPriorityIndicator = 1
)

type CNOperatorSelectionEntity datatype.Enumerated

const (
	CNOperatorSelectionEntity_THE_SERVING_NETWORK_HAS_BEEN_SELECTED_BY_THE_UE      CNOperatorSelectionEntity = 0
	CNOperatorSelectionEntity_THE_SERVING_NETWORK_HAS_BEEN_SELECTED_BY_THE_NETWORK CNOperatorSelectionEntity = 1
)

type SGiPtPTunnellingMethod datatype.Enumerated

const (
	SGiPtPTunnellingMethod_UDP_IP_BASED SGiPtPTunnellingMethod = 0
	SGiPtPTunnellingMethod_OTHERS       SGiPtPTunnellingMethod = 1
)

type CPCIoTEPSOptimisationIndicator datatype.Enumerated

const (
	CPCIoTEPSOptimisationIndicator_NOT_APPLY CPCIoTEPSOptimisationIndicator = 0
	CPCIoTEPSOptimisationIndicator_APPLY     CPCIoTEPSOptimisationIndicator = 1
)

type UNIPDUCPOnlyFlag datatype.Enumerated

const (
	UNIPDUCPOnlyFlag_UNI_PDU_BOTH_UP_CP UNIPDUCPOnlyFlag = 0
	UNIPDUCPOnlyFlag_UNI_PDU_CP_ONLY    UNIPDUCPOnlyFlag = 1
)

type ChargingPerIPCANSessionIndicator datatype.Enumerated

const (
	ChargingPerIPCANSessionIndicator_INACTIVE ChargingPerIPCANSessionIndicator = 0
	ChargingPerIPCANSessionIndicator_ACTIVE   ChargingPerIPCANSessionIndicator = 1
)

type CSGAccessMode datatype.Enumerated

const (
	CSGAccessMode_CLOSED_MODE CSGAccessMode = 0
	CSGAccessMode_HYBRID_MODE CSGAccessMode = 1
)

type CSGMembershipIndication datatype.Enumerated

const (
	CSGMembershipIndication_NOT_CSG_MEMBER CSGMembershipIndication = 0
	CSGMembershipIndication_CSG_MEMBER     CSGMembershipIndication = 1
)

type AdditionalExceptionReports datatype.Enumerated

const (
	AdditionalExceptionReports_NOT_ALLOWED AdditionalExceptionReports = 0
	AdditionalExceptionReports_ALLOWED     AdditionalExceptionReports = 1
)

type RoleOfNode datatype.Enumerated

const (
	RoleOfNode_ORIGINATING_ROLE RoleOfNode = 0
	RoleOfNode_TERMINATING_ROLE RoleOfNode = 1
	RoleOfNode_FORWARDING_ROLE  RoleOfNode = 2
)

type NodeFunctionality datatype.Enumerated

const (
	NodeFunctionality_S_CSCF         NodeFunctionality = 0
	NodeFunctionality_P_CSCF         NodeFunctionality = 1
	NodeFunctionality_I_CSCF         NodeFunctionality = 2
	NodeFunctionality_MRFC           NodeFunctionality = 3
	NodeFunctionality_MGCF           NodeFunctionality = 4
	NodeFunctionality_BGCF           NodeFunctionality = 5
	NodeFunctionality_AS             NodeFunctionality = 6
	NodeFunctionality_IBCF           NodeFunctionality = 7
	NodeFunctionality_S_GW           NodeFunctionality = 8
	NodeFunctionality_P_GW           NodeFunctionality = 9
	NodeFunctionality_HSGW           NodeFunctionality = 10
	NodeFunctionality_E_CSCF         NodeFunctionality = 11
	NodeFunctionality_MME            NodeFunctionality = 12
	NodeFunctionality_TRF            NodeFunctionality = 13
	NodeFunctionality_TF             NodeFunctionality = 14
	NodeFunctionality_ATCF           NodeFunctionality = 15
	NodeFunctionality_PROXY_FUNCTION NodeFunctionality = 16
	NodeFunctionality_EPDG           NodeFunctionality = 17
	NodeFunctionality_TDF            NodeFunctionality = 18
	NodeFunctionality_TWAG           NodeFunctionality = 19
	NodeFunctionality_SCEF           NodeFunctionality = 20
	NodeFunctionality_IWK_SCEF       NodeFunctionality = 21
)

type SessionPriority datatype.Enumerated

const (
	SessionPriority_PRIORITY_0 SessionPriority = 0
	SessionPriority_PRIORITY_1 SessionPriority = 1
	SessionPriority_PRIORITY_2 SessionPriority = 2
	SessionPriority_PRIORITY_3 SessionPriority = 3
	SessionPriority_PRIORITY_4 SessionPriority = 4
)

type OnlineChargingFlag datatype.Enumerated

const (
	OnlineChargingFlag_ECF_ADDRESS_NOT_PROVIDED OnlineChargingFlag = 0
	OnlineChargingFlag_ECF_ADDRESS_PROVIDED     OnlineChargingFlag = 1
)

type IMSEmergencyIndicator datatype.Enumerated

const (
	IMSEmergencyIndicator_NON_EMERGENCY IMSEmergencyIndicator = 0
	IMSEmergencyIndicator_EMERGENCY     IMSEmergencyIndicator = 1
)

type TADIdentifier datatype.Enumerated

const (
	TADIdentifier_CS TADIdentifier = 0
	TADIdentifier_PS TADIdentifier = 1
)

type StatusASCode datatype.Enumerated

const (
	StatusASCode_4XX     StatusASCode = 0
	StatusASCode_5XX     StatusASCode = 1
	StatusASCode_TIMEOUT StatusASCode = 2
)

type LocalGWInsertedIndication datatype.Enumerated

const (
	LocalGWInsertedIndication_LOCAL_GW_NOT_INSERTED LocalGWInsertedIndication = 0
	LocalGWInsertedIndication_LOCAL_GW_INSERTED     LocalGWInsertedIndication = 1
)

type IPRealmDefaultIndication datatype.Enumerated

const (
	IPRealmDefaultIndication_DEFAULT_IP_REALM_NOT_USED IPRealmDefaultIndication = 0
	IPRealmDefaultIndication_DEFAULT_IP_REALM_USED     IPRealmDefaultIndication = 1
)

type TranscoderInsertedIndication datatype.Enumerated

const (
	TranscoderInsertedIndication_TRANSCODER_NOT_INSERTED TranscoderInsertedIndication = 0
	TranscoderInsertedIndication_TRANSCODER_INSERTED     TranscoderInsertedIndication = 1
)

type MediaInitiatorFlag datatype.Enumerated

const (
	MediaInitiatorFlag_CALLED_PARTY  MediaInitiatorFlag = 0
	MediaInitiatorFlag_CALLING_PARTY MediaInitiatorFlag = 1
	MediaInitiatorFlag_UNKNOWN       MediaInitiatorFlag = 2
)

type SDPType datatype.Enumerated

const (
	SDPType_SDP_OFFER  SDPType = 0
	SDPType_SDP_ANSWER SDPType = 1
)

type Originator datatype.Enumerated

const (
	Originator_CALLING_PARTY Originator = 0
	Originator_CALLED_PARTY  Originator = 1
)

type SessionDirection datatype.Enumerated

const (
	SessionDirection_INBOUND  SessionDirection = 0
	SessionDirection_OUTBOUND SessionDirection = 1
)

type NNIType datatype.Enumerated

const (
	NNIType_NON_ROAMING              NNIType = 0
	NNIType_ROAMING_WITHOUT_LOOPBACK NNIType = 1
	NNIType_ROAMING_WITH_LOOPBACK    NNIType = 2
)

type RelationshipMode datatype.Enumerated

const (
	RelationshipMode_TRUSTED     RelationshipMode = 0
	RelationshipMode_NON_TRUSTED RelationshipMode = 1
)

type AccessTransferType datatype.Enumerated

const (
	AccessTransferType_PS_TO_CS_TRANSFER AccessTransferType = 0
	AccessTransferType_CS_TO_PS_TRANSFER AccessTransferType = 1
	AccessTransferType_PS_TO_PS_TRANSFER AccessTransferType = 2
	AccessTransferType_CS_TO_CS_TRANSFER AccessTransferType = 3
)

type InterUETransfer datatype.Enumerated

const (
	InterUETransfer_INTRA_UE_TRANSFER InterUETransfer = 0
	InterUETransfer_INTER_UE_TRANSFER InterUETransfer = 1
)

type Priority datatype.Enumerated

const (
	Priority_LOW    Priority = 0
	Priority_NORMAL Priority = 1
	Priority_HIGH   Priority = 2
)

type MessageType datatype.Enumerated

const (
	MessageType_M_SEND_REQ         MessageType = 1
	MessageType_M_SEND_CONF        MessageType = 2
	MessageType_M_NOTIFICATION_IND MessageType = 3
	MessageType_M_NOTIFYRESP_IND   MessageType = 4
	MessageType_M_RETRIEVE_CONF    MessageType = 5
	MessageType_M_ACKNOWLEDGE_IND  MessageType = 6
	MessageType_M_DELIVERY_IND     MessageType = 7
	MessageType_M_READ_REC_IND     MessageType = 8
	MessageType_M_READ_ORIG_IND    MessageType = 9
	MessageType_M_FORWARD_REQ      MessageType = 10
	MessageType_M_FORWARD_CONF     MessageType = 11
	MessageType_M_MBOX_STORE_CONF  MessageType = 12
	MessageType_M_MBOX_VIEW_CONF   MessageType = 13
	MessageType_M_MBOX_UPLOAD_CONF MessageType = 14
	MessageType_M_MBOX_DELETE_CONF MessageType = 15
)

type DeliveryReportRequested datatype.Enumerated

const (
	DeliveryReportRequested_NO  DeliveryReportRequested = 0
	DeliveryReportRequested_YES DeliveryReportRequested = 1
)

type ReadReplyReportRequested datatype.Enumerated

const (
	ReadReplyReportRequested_NO  ReadReplyReportRequested = 0
	ReadReplyReportRequested_YES ReadReplyReportRequested = 1
)

type MMBoxStorageRequested datatype.Enumerated

const (
	MMBoxStorageRequested_NO  MMBoxStorageRequested = 0
	MMBoxStorageRequested_YES MMBoxStorageRequested = 1
)

type ContentClass datatype.Enumerated

const (
	ContentClass_TEXT          ContentClass = 0
	ContentClass_IMAGE_BASIC   ContentClass = 1
	ContentClass_IMAGE_RICH    ContentClass = 2
	ContentClass_VIDEO_BASIC   ContentClass = 3
	ContentClass_VIDEO_RICH    ContentClass = 4
	ContentClass_MEGAPIXEL     ContentClass = 5
	ContentClass_CONTENT_BASIC ContentClass = 6
	ContentClass_CONTENT_RICH  ContentClass = 7
)

type DRMContent datatype.Enumerated

const (
	DRMContent_NO  DRMContent = 0
	DRMContent_YES DRMContent = 1
)

type Adaptations datatype.Enumerated

const (
	Adaptations_YES Adaptations = 0
	Adaptations_NO  Adaptations = 1
)

type AddressType datatype.Enumerated

const (
	AddressType_E_MAIL_ADDRESS         AddressType = 0
	AddressType_MSISDN                 AddressType = 1
	AddressType_IPV4_ADDRESS           AddressType = 2
	AddressType_IPV6_ADDRESS           AddressType = 3
	AddressType_NUMERIC_SHORTCODE      AddressType = 4
	AddressType_ALPHANUMERIC_SHORTCODE AddressType = 5
	AddressType_OTHER                  AddressType = 6
	AddressType_IMSI                   AddressType = 7
)

type AddresseeType datatype.Enumerated

const (
	AddresseeType_TO  AddresseeType = 0
	AddresseeType_CC  AddresseeType = 1
	AddresseeType_BCC AddresseeType = 2
)

// Well-Known Values
// 0 */*
// 1 text/*
// 2 text/html
// 3 text/plain
// 4 text/x-hdml
// 5 text/x-ttml
// 6 text/x-vCalendar
// 7 text/x-vCard
// 8 text/vnd.wap.wml
// 9 text/vnd.wap.wmlscript
// 10 text/vnd.wap.wta-event
// 12 multipart/mixed
// 13 multipart/form-data
// 14 multipart/byterantes
// 15 multipart/alternative
// 16 application/*
// 17 application/java-vm
// 18 application/x-www-form-urlencoded
// 19 application/x-hdmlc
// 20 application/vnd.wap.wmlc
// 21 application/vnd.wap.wmlscriptc
// 22 application/vnd.wap.wta-eventc
// 23 application/vnd.wap.uaprof
// 24 application/vnd.wap.wtls-ca-certificate
// 25 application/vnd.wap.wtls-user-certificate
// 26 application/x-x509-ca-cert
// 27 application/x-x509-user-cert
// 28 image/*
// 29 image/gif
// 30 image/jpeg
// 31 image/tiff
// 32 image/png
// 33 image/vnd.wap.wbmp
// 34 application/vnd.wap.multipart.*
// 35 application/vnd.wap.multipart.mixed
// 36 application/vnd.wap.multipart.form-data
// 37 application/vnd.wap.multipart.byteranges
// 38 application/vnd.wap.multipart.alternative
// 39 application/xml
// 40 text/xml
// 41 application/vnd.wap.wbxml
// 42 application/x-x968-cross-cert
// 43 application/x-x968-ca-cert
// 44 application/x-x968-user-cert
// 45 text/vnd.wap.si
// 46 application/vnd.wap.sic
// 47 text/vnd.wap.sl
// 48 application/vnd.wap.slc
// 49 text/vnd.wap.co
// 50 application/vnd.wap.coc
// 51 application/vnd.wap.multipart.related
// 52 application/vnd.wap.sia
// 53 text/vnd.wap.connectivity-xml
// 54 application/vnd.wap.connectivity-wbxml
// 55 application/pkcs7-mime
// 56 application/vnd.wap.hashed-certificate
// 58 application/vnd.wap.cert-response
// 59 application/xhtml+xml
// 60 application/wml+xml
// 61 text/css
// 62 application/vnd.wap.mms-message
// 63 application/vnd.wap.rollover-certificate
// 64 application/vnd.wap.locc+wbxml
// 65 application/vnd.wap.loc+xml
// 66 application/vnd.syncml.dm+wbxml
// 67 application/vnd.syncml.dm+xml
// 68 application/vnd.syncml.notification
// 69 application/vnd.wap.xhtml+xml
// 70 application/vnd.wv.csp.cir
// 71 application/vnd.oma.dd+xml
// 72 application/vnd.oma.drm.message
// 73 application/vnd.oma.drm.content
// 74 application/vnd.oma.drm.rights+xml
// 75 application/vnd.oma.drm.rights+wbxml
// 76 application/vnd.wv.csp+xml
// 77 application/vnd.wv.csp+wbxml
// 78 application/vnd.syncml.ds.notification
// 79 audio/*
// 80 video/*
// 81 application/vnd.oma.dd2+xml
// 82 application/mikey
// 83 application/vnd.oma.dcd
// 84 application/vnd.oma.dcdc
// 85 text/x-vMessage
// 86 application/vnd.omads-email+wbxml
// 87 text/x-vBookmark
// 88 application/vnd.syncml.dm.notification
// 90 application/octet-stream
// 91 application/json
// Registered Values
// 513 application/vnd.uplanet.cacheop-wbxml
// 514 application/vnd.uplanet.signal
// 515 application/vnd.uplanet.alert-wbxml
// 516 application/vnd.uplanet.listcmd-wbxml
// 517 application/vnd.uplanet.listcmd-wbxml
// 518 application/vnd.uplanet.channel-wbxml
// 519 application/vnd.uplanet.provisioning-status-uri
// 520 x-wap.multipart/vnd.uplanet.header-set
// 521 application/vnd.uplanet.bearer-choice-wbxml
// 522 application/vnd.phonecom.mmc-wbxml
// 523 application/vnd.nokia.syncset+wbxml
// 524 image/x-up-wpng
// 768 application/iota.mmc-wbxml
// 769 application/iota.mmc-xml
// 770 application/vnd.syncml+xml
// undefined application/vnd.syncml+wbxml
// undefined text/vnd.wap.emn+xml
// 773 text/calendar
// 774 application/vnd.omads-email+xml
// 775 application/vnd.omads-file+xml
// 776 application/vnd.omads-folder+xml
// 777 text/directory;profile=vCard
// 778 application/vnd.wap.emn+wbxml
// 779 Karina Terekhova
// 780 application/vnd.motorola.screen3+xml
// 781 application/vnd.motorola.screen3+gzip
// 782 application/vnd.cmcc.setting+wbxml
// 783 application/vnd.cmcc.bombing+wbxml
// 784 application/vnd.docomo.pf
// 785 application/vnd.docomo.ub
// 786 application/vnd.omaloc-supl-init
// 787 application/vnd.oma.group-usage-list+xml
// 788 application/oma-directory+xml
// 789 application/vnd.docomo.pf2
// 790 application/vnd.oma.drm.roap-trigger+wbxml
// 791 application/vnd.sbm.mid2
// 792 application/vnd.wmf.bootstrap
// 793 application/vnc.cmcc.dcd+xml
// 794 application/vnd.sbm.cid
// 795 application/vnd.oma.bcast.provisioningtrigger
// 796 application/vnd.docomo.dm
// 797 application/vnd.oma.scidm.messages+xml
// 798 application/vnd.innopath.wamp.notification

type TypeNumber datatype.Enumerated

const (
	TypeNumber_0   TypeNumber = 0
	TypeNumber_1   TypeNumber = 1
	TypeNumber_2   TypeNumber = 2
	TypeNumber_3   TypeNumber = 3
	TypeNumber_4   TypeNumber = 4
	TypeNumber_5   TypeNumber = 5
	TypeNumber_6   TypeNumber = 6
	TypeNumber_7   TypeNumber = 7
	TypeNumber_8   TypeNumber = 8
	TypeNumber_9   TypeNumber = 9
	TypeNumber_10  TypeNumber = 10
	TypeNumber_12  TypeNumber = 12
	TypeNumber_13  TypeNumber = 13
	TypeNumber_14  TypeNumber = 14
	TypeNumber_15  TypeNumber = 15
	TypeNumber_16  TypeNumber = 16
	TypeNumber_17  TypeNumber = 17
	TypeNumber_18  TypeNumber = 18
	TypeNumber_19  TypeNumber = 19
	TypeNumber_20  TypeNumber = 20
	TypeNumber_21  TypeNumber = 21
	TypeNumber_22  TypeNumber = 22
	TypeNumber_23  TypeNumber = 23
	TypeNumber_24  TypeNumber = 24
	TypeNumber_25  TypeNumber = 25
	TypeNumber_26  TypeNumber = 26
	TypeNumber_27  TypeNumber = 27
	TypeNumber_28  TypeNumber = 28
	TypeNumber_29  TypeNumber = 29
	TypeNumber_30  TypeNumber = 30
	TypeNumber_31  TypeNumber = 31
	TypeNumber_32  TypeNumber = 32
	TypeNumber_33  TypeNumber = 33
	TypeNumber_34  TypeNumber = 34
	TypeNumber_35  TypeNumber = 35
	TypeNumber_36  TypeNumber = 36
	TypeNumber_37  TypeNumber = 37
	TypeNumber_38  TypeNumber = 38
	TypeNumber_39  TypeNumber = 39
	TypeNumber_40  TypeNumber = 40
	TypeNumber_41  TypeNumber = 41
	TypeNumber_42  TypeNumber = 42
	TypeNumber_43  TypeNumber = 43
	TypeNumber_44  TypeNumber = 44
	TypeNumber_45  TypeNumber = 45
	TypeNumber_46  TypeNumber = 46
	TypeNumber_47  TypeNumber = 47
	TypeNumber_48  TypeNumber = 48
	TypeNumber_49  TypeNumber = 49
	TypeNumber_50  TypeNumber = 50
	TypeNumber_51  TypeNumber = 51
	TypeNumber_52  TypeNumber = 52
	TypeNumber_53  TypeNumber = 53
	TypeNumber_54  TypeNumber = 54
	TypeNumber_55  TypeNumber = 55
	TypeNumber_56  TypeNumber = 56
	TypeNumber_58  TypeNumber = 58
	TypeNumber_59  TypeNumber = 59
	TypeNumber_60  TypeNumber = 60
	TypeNumber_61  TypeNumber = 61
	TypeNumber_62  TypeNumber = 62
	TypeNumber_63  TypeNumber = 63
	TypeNumber_64  TypeNumber = 64
	TypeNumber_65  TypeNumber = 65
	TypeNumber_66  TypeNumber = 66
	TypeNumber_67  TypeNumber = 67
	TypeNumber_68  TypeNumber = 68
	TypeNumber_69  TypeNumber = 69
	TypeNumber_70  TypeNumber = 70
	TypeNumber_71  TypeNumber = 71
	TypeNumber_72  TypeNumber = 72
	TypeNumber_73  TypeNumber = 73
	TypeNumber_74  TypeNumber = 74
	TypeNumber_75  TypeNumber = 75
	TypeNumber_76  TypeNumber = 76
	TypeNumber_77  TypeNumber = 77
	TypeNumber_78  TypeNumber = 78
	TypeNumber_79  TypeNumber = 79
	TypeNumber_80  TypeNumber = 80
	TypeNumber_81  TypeNumber = 81
	TypeNumber_82  TypeNumber = 82
	TypeNumber_83  TypeNumber = 83
	TypeNumber_84  TypeNumber = 84
	TypeNumber_85  TypeNumber = 85
	TypeNumber_86  TypeNumber = 86
	TypeNumber_87  TypeNumber = 87
	TypeNumber_88  TypeNumber = 88
	TypeNumber_90  TypeNumber = 90
	TypeNumber_91  TypeNumber = 91
	TypeNumber_513 TypeNumber = 513
	TypeNumber_514 TypeNumber = 514
	TypeNumber_515 TypeNumber = 515
	TypeNumber_516 TypeNumber = 516
	TypeNumber_517 TypeNumber = 517
	TypeNumber_518 TypeNumber = 518
	TypeNumber_519 TypeNumber = 519
	TypeNumber_520 TypeNumber = 520
	TypeNumber_521 TypeNumber = 521
	TypeNumber_522 TypeNumber = 522
	TypeNumber_523 TypeNumber = 523
	TypeNumber_524 TypeNumber = 524
	TypeNumber_768 TypeNumber = 768
	TypeNumber_769 TypeNumber = 769
	TypeNumber_770 TypeNumber = 770
	TypeNumber_773 TypeNumber = 773
	TypeNumber_774 TypeNumber = 774
	TypeNumber_775 TypeNumber = 775
	TypeNumber_776 TypeNumber = 776
	TypeNumber_777 TypeNumber = 777
	TypeNumber_778 TypeNumber = 778
	TypeNumber_779 TypeNumber = 779
	TypeNumber_780 TypeNumber = 780
	TypeNumber_781 TypeNumber = 781
	TypeNumber_782 TypeNumber = 782
	TypeNumber_783 TypeNumber = 783
	TypeNumber_784 TypeNumber = 784
	TypeNumber_785 TypeNumber = 785
	TypeNumber_786 TypeNumber = 786
	TypeNumber_787 TypeNumber = 787
	TypeNumber_788 TypeNumber = 788
	TypeNumber_789 TypeNumber = 789
	TypeNumber_790 TypeNumber = 790
	TypeNumber_791 TypeNumber = 791
	TypeNumber_792 TypeNumber = 792
	TypeNumber_793 TypeNumber = 793
	TypeNumber_794 TypeNumber = 794
	TypeNumber_795 TypeNumber = 795
	TypeNumber_796 TypeNumber = 796
	TypeNumber_797 TypeNumber = 797
	TypeNumber_798 TypeNumber = 798
)

type ClassIdentifier datatype.Enumerated

const (
	ClassIdentifier_PERSONAL      ClassIdentifier = 0
	ClassIdentifier_ADVERTISEMENT ClassIdentifier = 1
	ClassIdentifier_INFORMATIONAL ClassIdentifier = 2
	ClassIdentifier_AUTO          ClassIdentifier = 3
)

type LCSClientType datatype.Enumerated

const (
	LCSClientType_EMERGENCY_SERVICES        LCSClientType = 0
	LCSClientType_VALUE_ADDED_SERVICES      LCSClientType = 1
	LCSClientType_PLMN_OPERATOR_SERVICES    LCSClientType = 2
	LCSClientType_LAWFUL_INTERCEPT_SERVICES LCSClientType = 3
)

type LCSFormatIndicator datatype.Enumerated

const (
	LCSFormatIndicator_LOGICAL_NAME  LCSFormatIndicator = 0
	LCSFormatIndicator_EMAIL_ADDRESS LCSFormatIndicator = 1
	LCSFormatIndicator_MSISDN        LCSFormatIndicator = 2
	LCSFormatIndicator_URL           LCSFormatIndicator = 3
	LCSFormatIndicator_SIP_URL       LCSFormatIndicator = 4
)

type LocationEstimateType datatype.Enumerated

const (
	LocationEstimateType_CURRENT_LOCATION            LocationEstimateType = 0
	LocationEstimateType_CURRENT_LAST_KNOWN_LOCATION LocationEstimateType = 1
	LocationEstimateType_INITIAL_LOCATION            LocationEstimateType = 2
	LocationEstimateType_ACTIVATE_DEFERRED_LOCATION  LocationEstimateType = 3
	LocationEstimateType_CANCEL_DEFERRED_LOCATION    LocationEstimateType = 4
)

type PoCServerRole datatype.Enumerated

const (
	PoCServerRole_PARTICIPATING_POC_SERVER        PoCServerRole = 0
	PoCServerRole_CONTROLLING_POC_SERVER          PoCServerRole = 1
	PoCServerRole_INTERWORKING_FUNCTION           PoCServerRole = 2
	PoCServerRole_INTERWORKING_SELECTION_FUNCTION PoCServerRole = 3
)

type PoCSessionType datatype.Enumerated

const (
	PoCSessionType_1_TO_1_POC_SESSION             PoCSessionType = 0
	PoCSessionType_CHAT_POC_GROUP_SESSION         PoCSessionType = 1
	PoCSessionType_PRE_ARRANGED_POC_GROUP_SESSION PoCSessionType = 2
	PoCSessionType_AD_HOC_POC_GROUP_SESSION       PoCSessionType = 3
)

type PoCSessionInitiationType datatype.Enumerated

const (
	PoCSessionInitiationType_PRE_ESTABLISHED PoCSessionInitiationType = 0
	PoCSessionInitiationType_ON_DEMAND       PoCSessionInitiationType = 1
)

type PoCEventType datatype.Enumerated

const (
	PoCEventType_NORMAL                          PoCEventType = 0
	PoCEventType_INSTANT_PPERSONAL_AALERT_EVENT  PoCEventType = 1
	PoCEventType_POC_GROUP_ADVERTISEMENT_EVENT   PoCEventType = 2
	PoCEventType_EARLY_SSESSION_SETTING_UP_EVENT PoCEventType = 3
	PoCEventType_POC_TALK_BURST                  PoCEventType = 4
)

type PoCUserRoleInfoUnits datatype.Enumerated

const (
	PoCUserRoleInfoUnits_MODERATOR           PoCUserRoleInfoUnits = 1
	PoCUserRoleInfoUnits_DISPATCHER          PoCUserRoleInfoUnits = 2
	PoCUserRoleInfoUnits_SESSION_OWNER       PoCUserRoleInfoUnits = 3
	PoCUserRoleInfoUnits_SESSION_PARTICIPANT PoCUserRoleInfoUnits = 4
)

type ParticipantAccessPriority datatype.Enumerated

const (
	ParticipantAccessPriority_PRE_EMPTIVE_PRIORITY ParticipantAccessPriority = 1
	ParticipantAccessPriority_HIGH_PRIORITY        ParticipantAccessPriority = 2
	ParticipantAccessPriority_NORMAL_PRIORITY      ParticipantAccessPriority = 3
	ParticipantAccessPriority_LOW_PRIORITY         ParticipantAccessPriority = 4
)

type UserParticipatingType datatype.Enumerated

const (
	UserParticipatingType_NORMAL     UserParticipatingType = 0
	UserParticipatingType_NW_POC_BOX UserParticipatingType = 1
	UserParticipatingType_UE_POC_BOX UserParticipatingType = 2
)

type PoCChangeCondition datatype.Enumerated

const (
	PoCChangeCondition_SERVICECHANGE              PoCChangeCondition = 0
	PoCChangeCondition_VOLUMELIMIT                PoCChangeCondition = 1
	PoCChangeCondition_TIMELIMIT                  PoCChangeCondition = 2
	PoCChangeCondition_NUMBEROFTALKBURSTLIMIT     PoCChangeCondition = 3
	PoCChangeCondition_NUMBEROFACTIVEPARTICIPANTS PoCChangeCondition = 4
	PoCChangeCondition_TARIFFTIME                 PoCChangeCondition = 5
)

type MBMSServiceType datatype.Enumerated

const (
	MBMSServiceType_MULTICAST MBMSServiceType = 0
	MBMSServiceType_BROADCAST MBMSServiceType = 1
)

type MBMSUserServiceType datatype.Enumerated

const (
	MBMSUserServiceType_DOWNLOAD  MBMSUserServiceType = 1
	MBMSUserServiceType_STREAMING MBMSUserServiceType = 2
)

type FileRepairSupported datatype.Enumerated

const (
	FileRepairSupported_SUPPORTED     FileRepairSupported = 1
	FileRepairSupported_NOT_SUPPORTED FileRepairSupported = 2
)

type MBMS2G3GIndicator datatype.Enumerated

const (
	MBMS2G3GIndicator_2G        MBMS2G3GIndicator = 0
	MBMS2G3GIndicator_3G        MBMS2G3GIndicator = 1
	MBMS2G3GIndicator_2G_AND_3G MBMS2G3GIndicator = 2
)

type CNIPMulticastDistribution datatype.Enumerated

const (
	CNIPMulticastDistribution_NO_IP_MULTICAST CNIPMulticastDistribution = 0
	CNIPMulticastDistribution_IP_MULTICAST    CNIPMulticastDistribution = 1
)

type MBMSChargedParty datatype.Enumerated

const (
	MBMSChargedParty_CONTENT_PROVIDER MBMSChargedParty = 0
	MBMSChargedParty_SUBSCRIBER       MBMSChargedParty = 1
)

type SMSNode datatype.Enumerated

const (
	SMSNode_SMS_ROUTER              SMSNode = 0
	SMSNode_IP_SM_GW                SMSNode = 1
	SMSNode_SMS_ROUTER_AND_IP_SM_GW SMSNode = 2
	SMSNode_SMS_SC                  SMSNode = 3
)

type SMMessageType datatype.Enumerated

const (
	SMMessageType_SUBMISSION           SMMessageType = 0
	SMMessageType_DELIVERY_REPORT      SMMessageType = 1
	SMMessageType_SM_SERVICE_REQUEST   SMMessageType = 2.
	SMMessageType_T4_DEVICE_TRIGGER    SMMessageType = 3
	SMMessageType_SM_DEVICE_TRIGGER    SMMessageType = 4
	SMMessageType_MO_SMS_T4_SUBMISSION SMMessageType = 5
)

type ReplyPathRequested datatype.Enumerated

const (
	ReplyPathRequested_NO_REPLY_PATH_SET ReplyPathRequested = 0
	ReplyPathRequested_REPLY_PATH_SET    ReplyPathRequested = 1
)

type SMServiceType datatype.Enumerated

const (
	SMServiceType_VAS4SMS_SHORT_MESSAGE_CONTENT_PROCESSING                SMServiceType = 0
	SMServiceType_VAS4SMS_SHORT_MESSAGE_FORWARDING                        SMServiceType = 1
	SMServiceType_VAS4SMS_SHORT_MESSAGE_FORWARDING_MULTIPLE_SUBSCRIPTIONS SMServiceType = 2
	SMServiceType_VAS4SMS_SHORT_MESSAGE_FILTERING                         SMServiceType = 3
	SMServiceType_VAS4SMS_SHORT_MESSAGE_RECEIPT                           SMServiceType = 4
	SMServiceType_VAS4SMS_SHORT_MESSAGE_NETWORK_STORAGE                   SMServiceType = 5
	SMServiceType_VAS4SMS_SHORT_MESSAGE_TO_MULTIPLE_DESTINATIONS          SMServiceType = 6
	SMServiceType_VAS4SMS_SHORT_MESSAGE_VIRTUAL_PRIVATE_NETWORK_VPN       SMServiceType = 7
	SMServiceType_VAS4SMS_SHORT_MESSAGE_AUTO_REPLY                        SMServiceType = 8
	SMServiceType_VAS4SMS_SHORT_MESSAGE_PERSONAL_SIGNATURE                SMServiceType = 9
	SMServiceType_VAS4SMS_SHORT_MESSAGE_DEFERRED_DELIVERY                 SMServiceType = 10
)

type SMDeviceTriggerIndicator datatype.Enumerated

const (
	SMDeviceTriggerIndicator_NOT_DEVICETRIGGER      SMDeviceTriggerIndicator = 0
	SMDeviceTriggerIndicator_DEVICE_TRIGGER_REQUEST SMDeviceTriggerIndicator = 1
	SMDeviceTriggerIndicator_DEVICE_TRIGGER_REPLACE SMDeviceTriggerIndicator = 2
	SMDeviceTriggerIndicator_DEVICE_TRIGGER_RECALL  SMDeviceTriggerIndicator = 3
)

type InterfaceType datatype.Enumerated

const (
	InterfaceType_UNKNOWN                 InterfaceType = 0
	InterfaceType_MOBILE_ORIGINATING      InterfaceType = 1
	InterfaceType_MOBILE_TERMINATING      InterfaceType = 2
	InterfaceType_APPLICATION_ORIGINATING InterfaceType = 3
	InterfaceType_APPLICATION_TERMINATION InterfaceType = 4
)

type PriorityIndication datatype.Enumerated

const (
	PriorityIndication_NON_PRIORITY PriorityIndication = 0
	PriorityIndication_PRIORITY     PriorityIndication = 1
)

type ForwardingPending datatype.Enumerated

const (
	ForwardingPending_FORWARDING_NOT_PENDING ForwardingPending = 0
	ForwardingPending_FORWARDING_PENDING     ForwardingPending = 1
)

type ParticipantActionType datatype.Enumerated

const (
	ParticipantActionType_CREATE_CONF      ParticipantActionType = 0
	ParticipantActionType_JOIN_CONF        ParticipantActionType = 1
	ParticipantActionType_INVITE_INTO_CONF ParticipantActionType = 2
	ParticipantActionType_QUIT_CONF        ParticipantActionType = 3
)

type RoleOfProSeFunction datatype.Enumerated

const (
	RoleOfProSeFunction_HPLMN      RoleOfProSeFunction = 0
	RoleOfProSeFunction_VPLMN      RoleOfProSeFunction = 1
	RoleOfProSeFunction_LOCAL_PLMN RoleOfProSeFunction = 2
)

type ProSeEventType datatype.Enumerated

const (
	ProSeEventType_ANNOUCING    ProSeEventType = 0
	ProSeEventType_MONITORING   ProSeEventType = 1
	ProSeEventType_MATCH_REPORT ProSeEventType = 2
)

type ProSeDirectDiscoveryModel datatype.Enumerated

const (
	ProSeDirectDiscoveryModel_MODEL_A ProSeDirectDiscoveryModel = 0
	ProSeDirectDiscoveryModel_MODEL_B ProSeDirectDiscoveryModel = 1
)

type ProSeRoleOfUE datatype.Enumerated

const (
	ProSeRoleOfUE_ANNOUNCING_UE ProSeRoleOfUE = 0
	ProSeRoleOfUE_MONITORING_UE ProSeRoleOfUE = 1
	ProSeRoleOfUE_REQUESTOR_UE  ProSeRoleOfUE = 2
	ProSeRoleOfUE_REQUESTED_UE  ProSeRoleOfUE = 3
)

type ProSeRangeClass datatype.Enumerated

const (
	ProSeRangeClass_RESERVED ProSeRangeClass = 0
	ProSeRangeClass_50_M     ProSeRangeClass = 1
	ProSeRangeClass_100_M    ProSeRangeClass = 2
	ProSeRangeClass_200_M    ProSeRangeClass = 3
	ProSeRangeClass_500_M    ProSeRangeClass = 4
	ProSeRangeClass_1000_M   ProSeRangeClass = 5
	ProSeRangeClass_UNUSED   ProSeRangeClass = 6 - 255
)

type ProximityAlertIndication datatype.Enumerated

const (
	ProximityAlertIndication_ALERT    ProximityAlertIndication = 0
	ProximityAlertIndication_NO_ALERT ProximityAlertIndication = 1
)

type ProSeReasonForCancellation datatype.Enumerated

const (
	ProSeReasonForCancellation_PROXIMITY_ALERT_SENT         ProSeReasonForCancellation = 0
	ProSeReasonForCancellation_TIME_EXPIRED_WITH_NO_RENEWAL ProSeReasonForCancellation = 1
	ProSeReasonForCancellation_REQUESTOR_CANCELLATION       ProSeReasonForCancellation = 2
)

type PC5RadioTechnology datatype.Enumerated

const (
	PC5RadioTechnology_EUTRA               PC5RadioTechnology = 0
	PC5RadioTechnology_WLAN                PC5RadioTechnology = 1
	PC5RadioTechnology_BOTH_EUTRA_AND_WLAN PC5RadioTechnology = 2
)

type CoverageStatus datatype.Enumerated

const (
	CoverageStatus_OUT_OF_COVERAGE CoverageStatus = 0
	CoverageStatus_IN_COVERAGE     CoverageStatus = 1
)

type ApplicationServiceType datatype.Enumerated

const (
	ApplicationServiceType_SIMPLE_IM_SENDING ApplicationServiceType = 100
	ApplicationServiceType_RECEIVING         ApplicationServiceType = 101
	ApplicationServiceType_RETRIEVAL         ApplicationServiceType = 102
	ApplicationServiceType_INVITING          ApplicationServiceType = 103
	ApplicationServiceType_LEAVING           ApplicationServiceType = 104
	ApplicationServiceType_JOINING           ApplicationServiceType = 105
)

type ProtocolType datatype.Enumerated

const (
	ProtocolType_HTTP ProtocolType = 0
	ProtocolType_COAP ProtocolType = 1
	ProtocolType_MQTT ProtocolType = 2
)

type RequestOperation datatype.Enumerated

const (
	RequestOperation_CREATE   RequestOperation = 1
	RequestOperation_RETRIEVE RequestOperation = 2
	RequestOperation_UPDATE   RequestOperation = 3
	RequestOperation_DELETE   RequestOperation = 4
	RequestOperation_NOTIFY   RequestOperation = 5
)

type ResponseStatusCode datatype.Enumerated

const (
	ResponseStatusCode_ACCEPTED                                    ResponseStatusCode = 1000
	ResponseStatusCode_OK                                          ResponseStatusCode = 2000
	ResponseStatusCode_CREATED                                     ResponseStatusCode = 2001
	ResponseStatusCode_DELETED                                     ResponseStatusCode = 2002
	ResponseStatusCode_UPDATED                                     ResponseStatusCode = 2004
	ResponseStatusCode_BAD_REQUEST                                 ResponseStatusCode = 4000
	ResponseStatusCode_NOT_FOUND                                   ResponseStatusCode = 4004
	ResponseStatusCode_OPERATION_NOT_ALLOWED                       ResponseStatusCode = 4005
	ResponseStatusCode_REQUEST_TIMEOUT                             ResponseStatusCode = 4008
	ResponseStatusCode_SUBSCRIPTION_CREATOR_HAS_NO_PRIVILEGE       ResponseStatusCode = 4101
	ResponseStatusCode_CONTENTS_UNACCEPTABLE                       ResponseStatusCode = 4102
	ResponseStatusCode_ORIGINATOR_HAS_NO_PRIVILEGE                 ResponseStatusCode = 4103
	ResponseStatusCode_GROUP_REQUEST_IDENTIFIER_EXISTS             ResponseStatusCode = 4104
	ResponseStatusCode_CONFLICT                                    ResponseStatusCode = 4105
	ResponseStatusCode_ORIGINATOR_HAS_NOT_REGISTERED               ResponseStatusCode = 4106
	ResponseStatusCode_SECURITY_ASSOCIATION_REQUIRED               ResponseStatusCode = 4107
	ResponseStatusCode_INVALID_CHILD_RESOURCE_TYPE                 ResponseStatusCode = 4108
	ResponseStatusCode_NO_MEMBERS                                  ResponseStatusCode = 4109
	ResponseStatusCode_GROUP_MEMBER_TYPE_INCONSISTENT              ResponseStatusCode = 4110
	ResponseStatusCode_INTERNAL_SERVER_ERROR                       ResponseStatusCode = 5000
	ResponseStatusCode_NOT_IMPLEMENTED                             ResponseStatusCode = 5001
	ResponseStatusCode_TARGET_NOT_REACHABLE                        ResponseStatusCode = 5103
	ResponseStatusCode_RECEIVER_HAS_NO_PRIVILEGE                   ResponseStatusCode = 5105
	ResponseStatusCode_ALREADY_EXISTS                              ResponseStatusCode = 5106
	ResponseStatusCode_TARGET_NOT_SUBSCRIBABLE                     ResponseStatusCode = 5203
	ResponseStatusCode_SUBSCRIPTION_VERIFICATION_INITIATION_FAILED ResponseStatusCode = 5204
	ResponseStatusCode_SUBSCRIPTION_HOST_HAS_NO_PRIVILEGE          ResponseStatusCode = 5205
	ResponseStatusCode_NON_BLOCKING_REQUEST_NOT_SUPPORTED          ResponseStatusCode = 5206
	ResponseStatusCode_NOT_ACCEPTABLE                              ResponseStatusCode = 5207
	ResponseStatusCode_GROUP_MEMBERS_NOT_RESPONDED                 ResponseStatusCode = 5209
	ResponseStatusCode_EXTERNAL_OBJECT_NOT_REACHABLE               ResponseStatusCode = 6003
	ResponseStatusCode_EXTERNAL_OBJECT_NOT_FOUND                   ResponseStatusCode = 6005
	ResponseStatusCode_MAX_NUMBER_OF_MEMBER_EXCEEDED               ResponseStatusCode = 6010
	ResponseStatusCode_MGMT_SESSION_CANNOT_BE_ESTABLISHED          ResponseStatusCode = 6020
	ResponseStatusCode_MGMT_SESSION_ESTABLISHMENT_TIMEOUT          ResponseStatusCode = 6021
	ResponseStatusCode_INVALID_CMDTYPE                             ResponseStatusCode = 6022
	ResponseStatusCode_INVALID_ARGUMENTS                           ResponseStatusCode = 6023
	ResponseStatusCode_INSUFFICIENT_ARGUMENTS                      ResponseStatusCode = 6024
	ResponseStatusCode_MGMT_CONVERSION_ERROR                       ResponseStatusCode = 6025
	ResponseStatusCode_MGMT_CANCELLATION_FAILED                    ResponseStatusCode = 6026
	ResponseStatusCode_ALREADY_COMPLETE                            ResponseStatusCode = 6028
	ResponseStatusCode_MGMT_COMMAND_NOT_CANCELLABLE                ResponseStatusCode = 6029
)

type CCSessionFailover datatype.Enumerated

const (
	CCSessionFailover_FAILOVER_NOT_SUPPORTED CCSessionFailover = 0
	CCSessionFailover_FAILOVER_SUPPORTED     CCSessionFailover = 1
)

type CreditControlFailureHandling datatype.Enumerated

const (
	CreditControlFailureHandling_TERMINATE           CreditControlFailureHandling = 0
	CreditControlFailureHandling_CONTINUE            CreditControlFailureHandling = 1
	CreditControlFailureHandling_RETRY_AND_TERMINATE CreditControlFailureHandling = 2
)

type DirectDebitingFailureHandling datatype.Enumerated

const (
	DirectDebitingFailureHandling_TERMINATE_OR_BUFFER DirectDebitingFailureHandling = 0
	DirectDebitingFailureHandling_CONTINUE            DirectDebitingFailureHandling = 1
)

type LowBalanceIndication datatype.Enumerated

const (
	LowBalanceIndication_NOT_APPLICABLE LowBalanceIndication = 0
	LowBalanceIndication_YES            LowBalanceIndication = 1
)

type CostInformation struct {
	UnitValue    UnitValue           `avp:"Unit-Value" json:"unitValue"`
	CurrencyCode datatype.Unsigned32 `avp:"Currency-Code" json:"currencyCode"`
	CostUnit     datatype.UTF8String `avp:"Cost-Unit" json:"costUnit,omitempty"`
}

type RemainingBalance struct {
	UnitValue    UnitValue           `avp:"Unit-Value" json:"unitValue"`
	CurrencyCode datatype.Unsigned32 `avp:"Currency-Code" json:"currencyCode"`
}

type FailedAVP struct {
}

type CheckBalanceResult datatype.Enumerated

const (
	CheckBalanceResult_ENOUGH_CREDIT CheckBalanceResult = 0
	CheckBalanceResult_NO_CREDIT     CheckBalanceResult = 1
)

type NASPortType datatype.Enumerated

const (
	NASPortType_Async                                          NASPortType = 0
	NASPortType_Sync                                           NASPortType = 1
	NASPortType_ISDN_Sync                                      NASPortType = 2
	NASPortType_ISDN_Async_V_120                               NASPortType = 3
	NASPortType_ISDN_Async_V_110                               NASPortType = 4
	NASPortType_Virtual                                        NASPortType = 5
	NASPortType_PIAFS                                          NASPortType = 6
	NASPortType_HDLC_Clear_Channel                             NASPortType = 7
	NASPortType_X_25                                           NASPortType = 8
	NASPortType_X_75                                           NASPortType = 9
	NASPortType_G_3_Fax                                        NASPortType = 10
	NASPortType_SDSL                                           NASPortType = 11
	NASPortType_ADSL_CAP                                       NASPortType = 12
	NASPortType_ADSL_DMT                                       NASPortType = 13
	NASPortType_IDSL                                           NASPortType = 14
	NASPortType_Ethernet                                       NASPortType = 15
	NASPortType_xDSL                                           NASPortType = 16
	NASPortType_Cable                                          NASPortType = 17
	NASPortType_Wireless_Other                                 NASPortType = 18
	NASPortType_Wireless_IEEE_802_11                           NASPortType = 19
	NASPortType_Token_Ring                                     NASPortType = 20
	NASPortType_FDDI                                           NASPortType = 21
	NASPortType_Wireless_CDMA2000                              NASPortType = 22
	NASPortType_Wireless_UMTS                                  NASPortType = 23
	NASPortType_Wireless_1X_EV                                 NASPortType = 24
	NASPortType_IAPP                                           NASPortType = 25
	NASPortType_FTTP                                           NASPortType = 26
	NASPortType_Wireless_IEEE_802_16                           NASPortType = 27
	NASPortType_Wireless_IEEE_802_20                           NASPortType = 28
	NASPortType_Wireless_IEEE_802_22                           NASPortType = 29
	NASPortType_PPPoA                                          NASPortType = 30
	NASPortType_PPPoEoA                                        NASPortType = 31
	NASPortType_PPPoEoE                                        NASPortType = 32
	NASPortType_PPPoEoVLAN                                     NASPortType = 33
	NASPortType_PPPoEoQinQ                                     NASPortType = 34
	NASPortType_xPON                                           NASPortType = 35
	NASPortType_Wireless_XGP                                   NASPortType = 36
	NASPortType_WiMAX_Pre_Release_8_IWK                        NASPortType = 37
	NASPortType_WIMAX_WIFI_IWK_WiMAX_WIFI_Interworking         NASPortType = 38
	NASPortType_WIMAX_SFF_Signaling_Fwd_Function_for_LTE_3GPP2 NASPortType = 39
	NASPortType_WIMAX_HA_LMA_WiMAX_HA_and_or_LMA_function      NASPortType = 40
	NASPortType_WIMAX_DHCP_WIMAX_DCHP_service                  NASPortType = 41
	NASPortType_WIMAX_LBS_WiMAX_location_based_service         NASPortType = 42
	NASPortType_WIMAX_WVS_WiMAX_voice_service                  NASPortType = 43
)

type FramedProtocol datatype.Enumerated

const (
	FramedProtocol_PPP      FramedProtocol = 1
	FramedProtocol_SLIP     FramedProtocol = 2
	FramedProtocol_ARAP     FramedProtocol = 3
	FramedProtocol_Gandalf  FramedProtocol = 4
	FramedProtocol_IPX_SLIP FramedProtocol = 5
	FramedProtocol_X75_Sync FramedProtocol = 6
	FramedProtocol_GPRS_PDP FramedProtocol = 7
)

type ServiceType datatype.Enumerated

const (
	ServiceType_Login                    ServiceType = 1
	ServiceType_Framed                   ServiceType = 2
	ServiceType_Callback_Login           ServiceType = 3
	ServiceType_Callback_Framed          ServiceType = 4
	ServiceType_Outbound                 ServiceType = 5
	ServiceType_Administrative           ServiceType = 6
	ServiceType_NAS_Prompt               ServiceType = 7
	ServiceType_Authenticate_Only        ServiceType = 8
	ServiceType_Callback_NAS_Prompt      ServiceType = 9
	ServiceType_Call_Check               ServiceType = 10
	ServiceType_Callback_Administrative  ServiceType = 11
	ServiceType_Voice                    ServiceType = 12
	ServiceType_Fax                      ServiceType = 13
	ServiceType_Modem_Relay              ServiceType = 14
	ServiceType_IAPP_Register            ServiceType = 15
	ServiceType_IAPP_AP_Check            ServiceType = 16
	ServiceType_Authorize_Only           ServiceType = 17
	ServiceType_Framed_Management        ServiceType = 18
	ServiceType_Additional_Authorization ServiceType = 19
)

type MIPHomeAgentHost struct {
	DestinationRealm datatype.DiameterIdentity `avp:"Destination-Realm" json:"destinationRealm"`
	DestinationHost  datatype.DiameterIdentity `avp:"Destination-Host" json:"destinationHost"`
}
type MIP6AgentInfo struct {
	MIPHomeAgentAddress []datatype.Address   `avp:"MIP-Home-Agent-Address" json:"mIPHomeAgentAddress,omitempty"`
	MIPHomeAgentHost    *MIPHomeAgentHost    `avp:"MIP-Home-Agent-Host" json:"mIPHomeAgentHost,omitempty"`
	MIP6HomeLinkPrefix  *datatype.Unsigned32 `avp:"MIP6-Home-Link-Prefix" json:"mIP6HomeLinkPrefix,omitempty"`
}
type QoSProfileTemplate struct {
	VendorId     datatype.Unsigned32 `avp:"Vendor-Id" json:"vendorId"`
	QoSProfileId datatype.Unsigned32 `avp:"QoS-Profile-Id" json:"qoSProfileId"`
}
type QoSCapability struct {
	QoSProfileTemplate []QoSProfileTemplate `avp:"QoS-Profile-Template" json:"qoSProfileTemplate"`
}

type Direction datatype.Enumerated

const (
	Direction_IN   Direction = 0
	Direction_OUT  Direction = 1
	Direction_BOTH Direction = 2
)

type IPAddressRange struct {
	IPAddressStart *datatype.Address `avp:"IP-Address-Start" json:"iPAddressStart,omitempty"`
	IPAddressEnd   *datatype.Address `avp:"IP-Address-End" json:"iPAddressEnd,omitempty"`
}
type IPAddressMask struct {
	IPAddress      *datatype.Address    `avp:"IP-Address" json:"iPAddress,omitempty"`
	IPBitMaskWidth *datatype.Unsigned32 `avp:"IP-Bit-Mask-Width" json:"iPBitMaskWidth,omitempty"`
}
type MACAddressMask struct {
	MACAddress            *datatype.OctetString `avp:"MAC-Address" json:"mACAddress,omitempty"`
	MACAddressMaskPattern *datatype.OctetString `avp:"MAC-Address-Mask-Pattern" json:"mACAddressMaskPattern,omitempty"`
}
type EUI64AddressMask struct {
	EUI64Address            *datatype.OctetString `avp:"EUI64-Address" json:"eUI64Address,omitempty"`
	EUI64AddressMaskPattern *datatype.OctetString `avp:"EUI64-Address-Mask-Pattern" json:"eUI64AddressMaskPattern,omitempty"`
}
type PortRange struct {
	PortStart *datatype.Integer32 `avp:"Port-Start" json:"portStart,omitempty"`
	PortEnd   *datatype.Integer32 `avp:"Port-End" json:"portEnd,omitempty"`
}
type Negated datatype.Enumerated

const (
	Negated_False Negated = 0
	Negated_True  Negated = 1
)

type UseAssignedAddress datatype.Enumerated

const (
	UseAssignedAddress_False UseAssignedAddress = 0
	UseAssignedAddress_True  UseAssignedAddress = 1
)

type FromToSpec struct {
	IPAddress          []datatype.Address     `avp:"IP-Address" json:"iPAddress,omitempty"`
	IPAddressRange     []IPAddressRange       `avp:"IP-Address-Range" json:"iPAddressRange,omitempty"`
	IPAddressMask      []IPAddressMask        `avp:"IP-Address-Mask" json:"iPAddressMask,omitempty"`
	MACAddress         []datatype.OctetString `avp:"MAC-Address" json:"mACAddress,omitempty"`
	MACAddressMask     []MACAddressMask       `avp:"MAC-Address-Mask" json:"mACAddressMask,omitempty"`
	EUI64Address       []datatype.OctetString `avp:"EUI64-Address" json:"eUI64Address,omitempty"`
	EUI64AddressMask   []EUI64AddressMask     `avp:"EUI64-Address-Mask" json:"eUI64AddressMask,omitempty"`
	Port               []datatype.Integer32   `avp:"Port" json:"port,omitempty"`
	PortRange          []PortRange            `avp:"Port-Range" json:"portRange,omitempty"`
	Negated            *Negated               `avp:"Negated" json:"negated,omitempty"`
	UseAssignedAddress *UseAssignedAddress    `avp:"Use-Assigned-Address" json:"useAssignedAddress,omitempty"`
}

type DiffservCodePoint datatype.Enumerated

const (
	DiffservCodePoint_IN DiffservCodePoint = 0
)

type IPOption struct {
	IPOptionType  IPOptionType           `avp:"IP-Option-Type" json:"iPOptionType"`
	IPOptionValue []datatype.OctetString `avp:"IP-Option-Value" json:"iPOptionValue,omitempty"`
	Negated       *Negated               `avp:"Negated" json:"negated,omitempty"`
}

type TCPOption struct {
	TCPOptionType  TCPOptionType          `avp:"TCP-Option-Type" json:"tCPOptionType"`
	TCPOptionValue []datatype.OctetString `avp:"TCP-Option-Value" json:"tCPOptionValue,omitempty"`
	Negated        *Negated               `avp:"Negated" json:"negated,omitempty"`
}
type TCPFlags struct {
	TCPFlagType datatype.Unsigned32 `avp:"TCP-Flag-Type" json:"tCPFlagType"`
	Negated     *Negated            `avp:"Negated" json:"negated,omitempty"`
}

type ICMPType struct {
	ICMPTypeNumber ICMPTypeNumber `avp:"ICMP-Type-Number" json:"iCMPTypeNumber"`
	ICMPCode       []ICMPCode     `avp:"ICMP-Code" json:"iCMPCode,omitempty"`
	Negated        *Negated       `avp:"Negated" json:"negated,omitempty"`
}
type ETHProtoType struct {
	ETHEtherType []datatype.OctetString `avp:"ETH-Ether-Type" json:"eTHEtherType,omitempty"`
	ETHSAP       []datatype.OctetString `avp:"ETH-SAP" json:"eTHSAP,omitempty"`
}
type VLANIDRange struct {
	SVIDStart *datatype.Unsigned32 `avp:"S-VID-Start" json:"sVIDStart,omitempty"`
	SVIDEnd   *datatype.Unsigned32 `avp:"S-VID-End" json:"sVIDEnd,omitempty"`
	CVIDStart *datatype.Unsigned32 `avp:"C-VID-Start" json:"cVIDStart,omitempty"`
	CVIDEnd   *datatype.Unsigned32 `avp:"C-VID-End" json:"cVIDEnd,omitempty"`
}
type UserPriorityRange struct {
	LowUserPriority  []datatype.Unsigned32 `avp:"Low-User-Priority" json:"lowUserPriority,omitempty"`
	HighUserPriority []datatype.Unsigned32 `avp:"High-User-Priority" json:"highUserPriority,omitempty"`
}
type ETHOption struct {
	ETHProtoType      ETHProtoType        `avp:"ETH-Proto-Type" json:"eTHProtoType"`
	VLANIDRange       []VLANIDRange       `avp:"VLAN-ID-Range" json:"vLANIDRange,omitempty"`
	UserPriorityRange []UserPriorityRange `avp:"User-Priority-Range" json:"userPriorityRange,omitempty"`
}
type Classifier struct {
	ClassifierID      datatype.OctetString `avp:"Classifier-ID" json:"classifierID"`
	Protocol          *Protocol            `avp:"Protocol" json:"protocol,omitempty"`
	Direction         *Direction           `avp:"Direction" json:"direction,omitempty"`
	FromSpec          []FromToSpec         `avp:"From-Spec" json:"fromSpec,omitempty"`
	ToSpec            []FromToSpec         `avp:"To-Spec" json:"toSpec,omitempty"`
	DiffservCodePoint []DiffservCodePoint  `avp:"Diffserv-Code-Point" json:"diffservCodePoint,omitempty"`
	IPOption          []IPOption           `avp:"IP-Option" json:"iPOption,omitempty"`
	TCPOption         []TCPOption          `avp:"TCP-Option" json:"tCPOption,omitempty"`
	TCPFlags          *TCPFlags            `avp:"TCP-Flags" json:"tCPFlags,omitempty"`
	ICMPType          []ICMPType           `avp:"ICMP-Type" json:"iCMPType,omitempty"`
	ETHOption         []ETHOption          `avp:"ETH-Option" json:"eTHOption,omitempty"`
}
type TimezoneFlag datatype.Enumerated

const (
	TimezoneFlag_UTC    TimezoneFlag = 0
	TimezoneFlag_LOCAL  TimezoneFlag = 1
	TimezoneFlag_OFFSET TimezoneFlag = 2
)

type TimeOfDayCondition struct {
	TimeOfDayStart    *datatype.Unsigned32 `avp:"Time-Of-Day-Start" json:"timeOfDayStart,omitempty"`
	TimeOfDayEnd      *datatype.Unsigned32 `avp:"Time-Of-Day-End" json:"timeOfDayEnd,omitempty"`
	DayOfWeekMask     *datatype.Unsigned32 `avp:"Day-Of-Week-Mask" json:"dayOfWeekMask,omitempty"`
	DayOfMonthMask    *datatype.Unsigned32 `avp:"Day-Of-Month-Mask" json:"dayOfMonthMask,omitempty"`
	MonthOfYearMask   *datatype.Unsigned32 `avp:"Month-Of-Year-Mask" json:"monthOfYearMask,omitempty"`
	AbsoluteStartTime *datatype.Time       `avp:"Absolute-Start-Time" json:"absoluteStartTime,omitempty"`
	AbsoluteEndTime   *datatype.Time       `avp:"Absolute-End-Time" json:"absoluteEndTime,omitempty"`
	TimezoneFlag      *TimezoneFlag        `avp:"Timezone-Flag" json:"timezoneFlag,omitempty"`
}
type TreatmentAction datatype.Enumerated

const (
	TreatmentAction_DROP   TreatmentAction = 0
	TreatmentAction_SHAPE  TreatmentAction = 1
	TreatmentAction_MARK   TreatmentAction = 2
	TreatmentAction_PERMIT TreatmentAction = 3
)

type QoSSemantics datatype.Enumerated

const (
	QoSSemantics_QoS_Desired    QoSSemantics = 0
	QoSSemantics_QoS_Available  QoSSemantics = 1
	QoSSemantics_QoS_Delivered  QoSSemantics = 2
	QoSSemantics_Minimum_QoS    QoSSemantics = 3
	QoSSemantics_QoS_Authorized QoSSemantics = 4
)

type QoSParameters struct {
	TMOD1     *TMOD1               `avp:"TMOD-1" json:"tMOD1,omitempty"`
	TMOD2     *TMOD2               `avp:"TMOD-2" json:"tMOD2,omitempty"`
	Bandwidth *datatype.Float32    `avp:"Bandwidth" json:"bandwidth,omitempty"`
	PHBClass  *datatype.Unsigned32 `avp:"PHB-Class" json:"pHBClass,omitempty"`
}
type ExcessTreatment struct {
	TreatmentAction    TreatmentAction     `avp:"Treatment-Action" json:"treatmentAction"`
	QoSProfileTemplate *QoSProfileTemplate `avp:"QoS-Profile-Template" json:"qoSProfileTemplate,omitempty"`
	QoSParameters      *QoSParameters      `avp:"QoS-Parameters" json:"qoSParameters,omitempty"`
}
type FilterRule struct {
	FilterRulePrecedence *datatype.Unsigned32 `avp:"Filter-Rule-Precedence" json:"filterRulePrecedence,omitempty"`
	Classifier           *Classifier          `avp:"Classifier" json:"classifier,omitempty"`
	TimeOfDayCondition   []TimeOfDayCondition `avp:"Time-Of-Day-Condition" json:"timeOfDayCondition,omitempty"`
	TreatmentAction      *TreatmentAction     `avp:"Treatment-Action" json:"treatmentAction,omitempty"`
	QoSSemantics         *QoSSemantics        `avp:"QoS-Semantics" json:"qoSSemantics,omitempty"`
	QoSProfileTemplate   *QoSProfileTemplate  `avp:"QoS-Profile-Template" json:"qoSProfileTemplate,omitempty"`
	QoSParameters        *QoSParameters       `avp:"QoS-Parameters" json:"qoSParameters,omitempty"`
	ExcessTreatment      *ExcessTreatment     `avp:"Excess-Treatment" json:"excessTreatment,omitempty"`
}
type QoSResources struct {
	FilterRule []FilterRule `avp:"Filter-Rule" json:"filterRule"`
}

type MeasurementPeriodLTE datatype.Enumerated

const (
	MeasurementPeriodLTE_1024ms  MeasurementPeriodLTE = 0
	MeasurementPeriodLTE_1280ms  MeasurementPeriodLTE = 1
	MeasurementPeriodLTE_2048ms  MeasurementPeriodLTE = 2
	MeasurementPeriodLTE_2560ms  MeasurementPeriodLTE = 3
	MeasurementPeriodLTE_5120ms  MeasurementPeriodLTE = 4
	MeasurementPeriodLTE_10240ms MeasurementPeriodLTE = 5
	MeasurementPeriodLTE_60000ms MeasurementPeriodLTE = 6
)

type MeasurementPeriodUMTS datatype.Enumerated

const (
	MeasurementPeriodUMTS_250ms   MeasurementPeriodUMTS = 0
	MeasurementPeriodUMTS_500ms   MeasurementPeriodUMTS = 1
	MeasurementPeriodUMTS_1000ms  MeasurementPeriodUMTS = 2
	MeasurementPeriodUMTS_2000ms  MeasurementPeriodUMTS = 3
	MeasurementPeriodUMTS_3000ms  MeasurementPeriodUMTS = 4
	MeasurementPeriodUMTS_4000ms  MeasurementPeriodUMTS = 5
	MeasurementPeriodUMTS_6000ms  MeasurementPeriodUMTS = 6
	MeasurementPeriodUMTS_8000ms  MeasurementPeriodUMTS = 7
	MeasurementPeriodUMTS_12000ms MeasurementPeriodUMTS = 8
	MeasurementPeriodUMTS_16000ms MeasurementPeriodUMTS = 9
	MeasurementPeriodUMTS_20000ms MeasurementPeriodUMTS = 10
	MeasurementPeriodUMTS_24000ms MeasurementPeriodUMTS = 11
	MeasurementPeriodUMTS_28000ms MeasurementPeriodUMTS = 12
	MeasurementPeriodUMTS_32000ms MeasurementPeriodUMTS = 13
	MeasurementPeriodUMTS_64000ms MeasurementPeriodUMTS = 14
)

type CollectionPeriodRRMLTE datatype.Enumerated

const (
	CollectionPeriodRRMLTE_1024ms  CollectionPeriodRRMLTE = 0
	CollectionPeriodRRMLTE_1280ms  CollectionPeriodRRMLTE = 1
	CollectionPeriodRRMLTE_2048ms  CollectionPeriodRRMLTE = 2
	CollectionPeriodRRMLTE_2560ms  CollectionPeriodRRMLTE = 3
	CollectionPeriodRRMLTE_5120ms  CollectionPeriodRRMLTE = 4
	CollectionPeriodRRMLTE_10240ms CollectionPeriodRRMLTE = 5
	CollectionPeriodRRMLTE_60000ms CollectionPeriodRRMLTE = 6
)

type CollectionPeriodRRMUMTS datatype.Enumerated

const (
	CollectionPeriodRRMUMTS_250ms   CollectionPeriodRRMUMTS = 0
	CollectionPeriodRRMUMTS_500ms   CollectionPeriodRRMUMTS = 1
	CollectionPeriodRRMUMTS_1000ms  CollectionPeriodRRMUMTS = 2
	CollectionPeriodRRMUMTS_2000ms  CollectionPeriodRRMUMTS = 3
	CollectionPeriodRRMUMTS_3000ms  CollectionPeriodRRMUMTS = 4
	CollectionPeriodRRMUMTS_4000ms  CollectionPeriodRRMUMTS = 5
	CollectionPeriodRRMUMTS_6000ms  CollectionPeriodRRMUMTS = 6
	CollectionPeriodRRMUMTS_8000ms  CollectionPeriodRRMUMTS = 7
	CollectionPeriodRRMUMTS_12000ms CollectionPeriodRRMUMTS = 8
	CollectionPeriodRRMUMTS_16000ms CollectionPeriodRRMUMTS = 9
	CollectionPeriodRRMUMTS_20000ms CollectionPeriodRRMUMTS = 10
	CollectionPeriodRRMUMTS_24000ms CollectionPeriodRRMUMTS = 11
	CollectionPeriodRRMUMTS_28000ms CollectionPeriodRRMUMTS = 12
	CollectionPeriodRRMUMTS_32000ms CollectionPeriodRRMUMTS = 13
	CollectionPeriodRRMUMTS_64000ms CollectionPeriodRRMUMTS = 14
)

type MDTConfiguration struct {
	JobType                 JobType                  `avp:"Job-Type" json:"jobType"`
	AreaScope               *AreaScope               `avp:"Area-Scope" json:"areaScope,omitempty"`
	ListOfMeasurements      *datatype.Unsigned32     `avp:"List-Of-Measurements" json:"listOfMeasurements,omitempty"`
	ReportingTrigger        *datatype.Unsigned32     `avp:"Reporting-Trigger" json:"reportingTrigger,omitempty"`
	ReportInterval          *ReportInterval          `avp:"Report-Interval" json:"reportInterval,omitempty"`
	ReportAmount            *ReportAmount            `avp:"Report-Amount" json:"reportAmount,omitempty"`
	EventThresholdRSRP      *datatype.Unsigned32     `avp:"Event-Threshold-RSRP" json:"eventThresholdRSRP,omitempty"`
	EventThresholdRSRQ      *datatype.Unsigned32     `avp:"Event-Threshold-RSRQ" json:"eventThresholdRSRQ,omitempty"`
	LoggingInterval         *LoggingInterval         `avp:"Logging-Interval" json:"loggingInterval,omitempty"`
	LoggingDuration         *LoggingDuration         `avp:"Logging-Duration" json:"loggingDuration,omitempty"`
	MeasurementPeriodLTE    *MeasurementPeriodLTE    `avp:"Measurement-Period-LTE" json:"measurementPeriodLTE,omitempty"`
	MeasurementPeriodUMTS   *MeasurementPeriodUMTS   `avp:"Measurement-Period-UMTS" json:"measurementPeriodUMTS,omitempty"`
	CollectionPeriodRRMLTE  *CollectionPeriodRRMLTE  `avp:"Collection-Period-RRM-LTE" json:"collectionPeriodRRMLTE,omitempty"`
	CollectionPeriodRRMUMTS *CollectionPeriodRRMUMTS `avp:"Collection-Period-RRM-UMTS" json:"collectionPeriodRRMUMTS,omitempty"`
	PositioningMethod       *datatype.OctetString    `avp:"Positioning-Method" json:"positioningMethod,omitempty"`
	MeasurementQuantity     *datatype.OctetString    `avp:"Measurement-Quantity" json:"measurementQuantity,omitempty"`
	EventThresholdEvent1F   *datatype.Integer32      `avp:"Event-Threshold-Event-1F" json:"eventThresholdEvent1F,omitempty"`
	EventThresholdEvent1I   *datatype.Integer32      `avp:"Event-Threshold-Event-1I" json:"eventThresholdEvent1I,omitempty"`
	MDTAllowedPLMNId        []datatype.OctetString   `avp:"MDT-Allowed-PLMN-Id" json:"mDTAllowedPLMNId,omitempty"`
	MBSFNArea               []datatype.Unsigned32    `avp:"MBSFN-Area" json:"mBSFNArea,omitempty"`
}
type TraceInfo struct {
	TraceData      *TraceData            `avp:"Trace-Data" json:"traceData,omitempty"`
	TraceReference *datatype.OctetString `avp:"Trace-Reference" json:"traceReference,omitempty"`
}
type PDNType datatype.Enumerated

const (
	PDNType_IPv4         PDNType = 0
	PDNType_IPv6         PDNType = 1
	PDNType_IPv4v6       PDNType = 2
	PDNType_IPv4_OR_IPv6 PDNType = 3
)

type EPSSubscribedQoSProfile struct {
	QoSClassIdentifier          QoSClassIdentifier          `avp:"QoS-Class-Identifier" json:"qoSClassIdentifier"`
	AllocationRetentionPriority AllocationRetentionPriority `avp:"Allocation-Retention-Priority" json:"allocationRetentionPriority"`
	PDNType                     PDNType                     `avp:"PDN-Type" json:"pDNType"`
}
type VPLMNDynamicAddressAllowed datatype.Enumerated

const (
	VPLMNDynamicAddressAllowed_NOTALLOWED VPLMNDynamicAddressAllowed = 0
	VPLMNDynamicAddressAllowed_ALLOWED    VPLMNDynamicAddressAllowed = 1
)

type PDNGWAllocationType datatype.Enumerated

const (
	PDNGWAllocationType_STATIC  PDNGWAllocationType = 0
	PDNGWAllocationType_DYNAMIC PDNGWAllocationType = 1
)

type AMBR struct {
	MaxRequestedBandwidthUL datatype.Unsigned32 `avp:"Max-Requested-Bandwidth-UL" json:"maxRequestedBandwidthUL"`
	MaxRequestedBandwidthDL datatype.Unsigned32 `avp:"Max-Requested-Bandwidth-DL" json:"maxRequestedBandwidthDL"`
}
type SpecificAPNInfo struct {
	ServiceSelection         datatype.UTF8String    `avp:"Service-Selection" json:"serviceSelection"`
	MIP6AgentInfo            MIP6AgentInfo          `avp:"MIP6-Agent-Info" json:"mIP6AgentInfo"`
	VisitedNetworkIdentifier []datatype.OctetString `avp:"Visited-Network-Identifier" json:"visitedNetworkIdentifier,omitempty"`
}
type SIPTOPermission datatype.Enumerated

const (
	SIPTOPermission_SIPTO_above_RAN_ALLOWED    SIPTOPermission = 0
	SIPTOPermission_SIPTO_above_RAN_NOTALLOWED SIPTOPermission = 1
)

type APNConfiguration struct {
	ContextIdentifier           datatype.Unsigned32         `avp:"Context-Identifier" json:"contextIdentifier"`
	ServedPartyIPAddress        []datatype.Address          `avp:"Served-Party-IP-Address" json:"servedPartyIPAddress,omitempty"`
	PDNType                     PDNType                     `avp:"PDN-Type" json:"pDNType"`
	ServiceSelection            datatype.UTF8String         `avp:"Service-Selection" json:"serviceSelection"`
	EPSSubscribedQoSProfile     *EPSSubscribedQoSProfile    `avp:"EPS-Subscribed-QoS-Profile" json:"ePSSubscribedQoSProfile,omitempty"`
	VPLMNDynamicAddressAllowed  *VPLMNDynamicAddressAllowed `avp:"VPLMN-Dynamic-Address-Allowed" json:"vPLMNDynamicAddressAllowed,omitempty"`
	MIP6AgentInfo               *MIP6AgentInfo              `avp:"MIP6-Agent-Info" json:"mIP6AgentInfo,omitempty"`
	VisitedNetworkIdentifier    *datatype.OctetString       `avp:"Visited-Network-Identifier" json:"visitedNetworkIdentifier,omitempty"`
	PDNGWAllocationType         *PDNGWAllocationType        `avp:"PDN-GW-Allocation-Type" json:"pDNGWAllocationType,omitempty"`
	TGPPChargingCharacteristics *datatype.UTF8String        `avp:"3GPP-Charging-Characteristics" json:"tGPPChargingCharacteristics,omitempty"`
	AMBR                        *AMBR                       `avp:"AMBR" json:"aMBR,omitempty"`
	SpecificAPNInfo             []SpecificAPNInfo           `avp:"Specific-APN-Info" json:"specificAPNInfo,omitempty"`
	APNOIReplacement            *datatype.UTF8String        `avp:"APN-OI-Replacement" json:"aPNOIReplacement,omitempty"`
	SIPTOPermission             *SIPTOPermission            `avp:"SIPTO-Permission" json:"sIPTOPermission,omitempty"`
	LIPAPermission              *datatype.Unsigned32        `avp:"LIPA-Permission" json:"lIPAPermission,omitempty"`
	RestorationPriority         *datatype.Unsigned32        `avp:"Restoration-Priority" json:"restorationPriority,omitempty"`
	SIPTOLocalNetworkPermission *datatype.Unsigned32        `avp:"SIPTO-Local-Network-Permission" json:"sIPTOLocalNetworkPermission,omitempty"`
	WLANoffloadability          *WLANoffloadability         `avp:"WLAN-offloadability" json:"wLANoffloadability,omitempty"`
}

type WLANoffloadability struct {
	WLANoffloadabilityEUTRAN *datatype.Unsigned32 `avp:"WLAN-offloadability-EUTRAN" json:"wLANoffloadabilityEUTRAN,omitempty"`
	WLANoffloadabilityUTRAN  *datatype.Unsigned32 `avp:"WLAN-offloadability-UTRAN" json:"wLANoffloadabilityUTRAN,omitempty"`
}

type Protocol datatype.Enumerated

const (
	Protocol_HOPOPT          Protocol = 0
	Protocol_ICMP            Protocol = 1
	Protocol_IGMP            Protocol = 2
	Protocol_GGP             Protocol = 3
	Protocol_IPv4            Protocol = 4
	Protocol_ST              Protocol = 5
	Protocol_TCP             Protocol = 6
	Protocol_CBT             Protocol = 7
	Protocol_EGP             Protocol = 8
	Protocol_IGP             Protocol = 9
	Protocol_HBBN_RCC_MON    Protocol = 10
	Protocol_NVP_II          Protocol = 11
	Protocol_PUP             Protocol = 12
	Protocol_ARGUS           Protocol = 13
	Protocol_EMCON           Protocol = 14
	Protocol_XNET            Protocol = 15
	Protocol_CHAOS           Protocol = 16
	Protocol_UDP             Protocol = 17
	Protocol_MUX             Protocol = 18
	Protocol_DCN_MEAS        Protocol = 19
	Protocol_HMP             Protocol = 20
	Protocol_PRM             Protocol = 21
	Protocol_XNS_IDP         Protocol = 22
	Protocol_TRUNK_1         Protocol = 23
	Protocol_TRUNK_2         Protocol = 24
	Protocol_LEAF_1          Protocol = 25
	Protocol_LEAF_2          Protocol = 26
	Protocol_RDP             Protocol = 27
	Protocol_IRTP            Protocol = 28
	Protocol_ISO_TP4         Protocol = 29
	Protocol_NETBLT          Protocol = 30
	Protocol_MFE_NSP         Protocol = 31
	Protocol_MERIT_INP       Protocol = 32
	Protocol_DCCP            Protocol = 33
	Protocol_3PC             Protocol = 34
	Protocol_IDPR            Protocol = 35
	Protocol_XTP             Protocol = 36
	Protocol_DDP             Protocol = 37
	Protocol_IDPR_CMTP       Protocol = 38
	Protocol_TP_PP           Protocol = 39
	Protocol_IL              Protocol = 40
	Protocol_IPv6            Protocol = 41
	Protocol_SDRP            Protocol = 42
	Protocol_IPv6_Route      Protocol = 43
	Protocol_IPv6_Frag       Protocol = 44
	Protocol_IDRP            Protocol = 45
	Protocol_RSVP            Protocol = 46
	Protocol_GRE             Protocol = 47
	Protocol_DSR             Protocol = 48
	Protocol_BNA             Protocol = 49
	Protocol_ESP             Protocol = 50
	Protocol_AH              Protocol = 51
	Protocol_I_NLSP          Protocol = 52
	Protocol_SWIPE           Protocol = 53
	Protocol_NARP            Protocol = 54
	Protocol_MOBILE          Protocol = 55
	Protocol_TLSP            Protocol = 56
	Protocol_SKIP            Protocol = 57
	Protocol_IPv6_ICMP       Protocol = 58
	Protocol_IPv6_NoNxt      Protocol = 59
	Protocol_IPv6_Opts       Protocol = 60
	Protocol_CFTP            Protocol = 62
	Protocol_SAT_EXPAK       Protocol = 64
	Protocol_KRYPTOLAN       Protocol = 65
	Protocol_RVD             Protocol = 66
	Protocol_IPPC            Protocol = 67
	Protocol_SAT_MON         Protocol = 69
	Protocol_VISA            Protocol = 70
	Protocol_IPCV            Protocol = 71
	Protocol_CPNX            Protocol = 72
	Protocol_CPHB            Protocol = 73
	Protocol_WSN             Protocol = 74
	Protocol_PVP             Protocol = 75
	Protocol_BR_SAT_MON      Protocol = 76
	Protocol_SUN_ND          Protocol = 77
	Protocol_WB_MON          Protocol = 78
	Protocol_WB_EXPAK        Protocol = 79
	Protocol_ISO_IP          Protocol = 80
	Protocol_VMTP            Protocol = 81
	Protocol_SECURE_VMTP     Protocol = 82
	Protocol_VINES           Protocol = 83
	Protocol_TTP             Protocol = 84
	Protocol_IPTM            Protocol = 84
	Protocol_NSFNET_IGP      Protocol = 85
	Protocol_DGP             Protocol = 86
	Protocol_TCF             Protocol = 87
	Protocol_EIGRP           Protocol = 88
	Protocol_OSPFIGP         Protocol = 89
	Protocol_Sprite_RPC      Protocol = 90
	Protocol_LARP            Protocol = 91
	Protocol_MTP             Protocol = 92
	Protocol_AX25            Protocol = 93
	Protocol_IPIP            Protocol = 94
	Protocol_SCC_SP          Protocol = 96
	Protocol_ETHERIP         Protocol = 97
	Protocol_ENCAP           Protocol = 98
	Protocol_GMTP            Protocol = 100
	Protocol_IFMP            Protocol = 101
	Protocol_PNNI            Protocol = 102
	Protocol_PIM             Protocol = 103
	Protocol_ARIS            Protocol = 104
	Protocol_SCPS            Protocol = 105
	Protocol_QNX             Protocol = 106
	Protocol_A_N             Protocol = 107
	Protocol_IPComp          Protocol = 108
	Protocol_SNP             Protocol = 109
	Protocol_Compaq_Peer     Protocol = 110
	Protocol_IPX_in_IP       Protocol = 111
	Protocol_VRRP            Protocol = 112
	Protocol_PGM             Protocol = 113
	Protocol_L2TP            Protocol = 115
	Protocol_DDX             Protocol = 116
	Protocol_IATP            Protocol = 117
	Protocol_STP             Protocol = 118
	Protocol_SRP             Protocol = 119
	Protocol_UTI             Protocol = 120
	Protocol_SMP             Protocol = 121
	Protocol_PTP             Protocol = 123
	Protocol_ISIS            Protocol = 124
	Protocol_FIRE            Protocol = 125
	Protocol_CRTP            Protocol = 126
	Protocol_CRUDP           Protocol = 127
	Protocol_SSCOPMCE        Protocol = 128
	Protocol_IPLT            Protocol = 129
	Protocol_SPS             Protocol = 130
	Protocol_PIPE            Protocol = 131
	Protocol_SCTP            Protocol = 132
	Protocol_FC              Protocol = 133
	Protocol_RSVP_E2E_IGNORE Protocol = 134
	Protocol_Mobility        Protocol = 135
	Protocol_UDPLite         Protocol = 136
	Protocol_MPLS_in_IP      Protocol = 137
	Protocol_manet           Protocol = 138
	Protocol_HIP             Protocol = 139
	Protocol_Shim6           Protocol = 140
	Protocol_WESP            Protocol = 141
	Protocol_ROHC            Protocol = 142
	Protocol_Ethernet        Protocol = 143
	Protocol_AGGFRAG         Protocol = 144
	Protocol_Reserved        Protocol = 255
)

type IPOptionType datatype.Enumerated

const (
	IPOptionType_End_of_Option_List           IPOptionType = 0
	IPOptionType_No_Operation                 IPOptionType = 1
	IPOptionType_Record_Route                 IPOptionType = 7
	IPOptionType_Experimental_Measurement     IPOptionType = 10
	IPOptionType_MTU_Probe                    IPOptionType = 11
	IPOptionType_MTU_Reply                    IPOptionType = 12
	IPOptionType_Quick_Start                  IPOptionType = 25
	IPOptionType_Traceroute                   IPOptionType = 82
	IPOptionType_Security                     IPOptionType = 130
	IPOptionType_Loose_Source_Route           IPOptionType = 131
	IPOptionType_Extended_Security            IPOptionType = 133
	IPOptionType_Commercial_Security          IPOptionType = 134
	IPOptionType_Stream_ID                    IPOptionType = 136
	IPOptionType_Strict_Source_Route          IPOptionType = 137
	IPOptionType_Experimental_Access_Control  IPOptionType = 142
	IPOptionType_IMI_Traffic_Descriptor       IPOptionType = 144
	IPOptionType_Extended_Internet_Protocol   IPOptionType = 145
	IPOptionType_Address_Extension            IPOptionType = 147
	IPOptionType_Router_Alert                 IPOptionType = 148
	IPOptionType_Selective_Directed_Broadcast IPOptionType = 149
	IPOptionType_Dynamic_Packet_State         IPOptionType = 151
	IPOptionType_Upstream_Multicast_Pkt       IPOptionType = 152
	IPOptionType_Experimental_Flow_Control    IPOptionType = 205
)

type TCPOptionType datatype.Enumerated

const (
	TCPOptionType_End_of_Option_List                  TCPOptionType = 0
	TCPOptionType_No_Operation                        TCPOptionType = 1
	TCPOptionType_Maximum_Segment_Size                TCPOptionType = 2
	TCPOptionType_Window_Scale                        TCPOptionType = 3
	TCPOptionType_SACK_Permitted                      TCPOptionType = 4
	TCPOptionType_SACK                                TCPOptionType = 5
	TCPOptionType_Echo                                TCPOptionType = 6
	TCPOptionType_Echo_Reply                          TCPOptionType = 7
	TCPOptionType_Timestamps                          TCPOptionType = 8
	TCPOptionType_Skeeter                             TCPOptionType = 16
	TCPOptionType_Bubba                               TCPOptionType = 17
	TCPOptionType_Trailer_Checksum_Option             TCPOptionType = 18
	TCPOptionType_MD5_Signature_Option                TCPOptionType = 19
	TCPOptionType_SCPS_Capabilities                   TCPOptionType = 20
	TCPOptionType_Selective_Negative_Acknowledgements TCPOptionType = 21
	TCPOptionType_Record_Boundaries                   TCPOptionType = 22
	TCPOptionType_Corruption_experienced              TCPOptionType = 23
	TCPOptionType_SNAP                                TCPOptionType = 24
	TCPOptionType_TCP_Compression_Filter              TCPOptionType = 26
	TCPOptionType_Quick_Start_Response                TCPOptionType = 27
	TCPOptionType_User_Timeout_Option                 TCPOptionType = 28
	TCPOptionType_TCP_Authentication_Option           TCPOptionType = 29
	TCPOptionType_Multipath_TCP                       TCPOptionType = 30
	TCPOptionType_TCP_Fast_Open_Cookie                TCPOptionType = 34
	TCPOptionType_Encryption_Negotiation              TCPOptionType = 69
	TCPOptionType_Accurate_ECN_Order_0                TCPOptionType = 172
	TCPOptionType_Accurate_ECN_Order_1                TCPOptionType = 174
	TCPOptionType_RFC3692_style_Experiment_1          TCPOptionType = 253
	TCPOptionType_RFC3692_style_Experiment_2          TCPOptionType = 254
)

type ICMPTypeNumber datatype.Enumerated

const (
	ICMPTypeNumber_Echo_Reply                                                                ICMPTypeNumber = 0
	ICMPTypeNumber_Destination_Unreachable                                                   ICMPTypeNumber = 3
	ICMPTypeNumber_Redirect                                                                  ICMPTypeNumber = 5
	ICMPTypeNumber_Unassigned                                                                ICMPTypeNumber = 7
	ICMPTypeNumber_Echo                                                                      ICMPTypeNumber = 8
	ICMPTypeNumber_Router_Advertisement                                                      ICMPTypeNumber = 9
	ICMPTypeNumber_Router_Solicitation                                                       ICMPTypeNumber = 10
	ICMPTypeNumber_Time_Exceeded                                                             ICMPTypeNumber = 11
	ICMPTypeNumber_Parameter_Problem                                                         ICMPTypeNumber = 12
	ICMPTypeNumber_Timestamp                                                                 ICMPTypeNumber = 13
	ICMPTypeNumber_Timestamp_Reply                                                           ICMPTypeNumber = 14
	ICMPTypeNumber_Photuris                                                                  ICMPTypeNumber = 40
	ICMPTypeNumber_ICMP_messages_utilized_by_experimental_mobility_protocols_such_as_Seamoby ICMPTypeNumber = 41
	ICMPTypeNumber_Extended_Echo_Request                                                     ICMPTypeNumber = 42
	ICMPTypeNumber_Extended_Echo_Reply                                                       ICMPTypeNumber = 43
	ICMPTypeNumber_RFC3692_style_Experiment_1                                                ICMPTypeNumber = 253
	ICMPTypeNumber_RFC3692_style_Experiment_2                                                ICMPTypeNumber = 254
	ICMPTypeNumber_Reserved                                                                  ICMPTypeNumber = 255
)

type ICMPCode datatype.Enumerated

const (
	ICMPCode_0  ICMPCode = 0
	ICMPCode_1  ICMPCode = 1
	ICMPCode_2  ICMPCode = 2
	ICMPCode_3  ICMPCode = 3
	ICMPCode_4  ICMPCode = 4
	ICMPCode_5  ICMPCode = 5
	ICMPCode_6  ICMPCode = 6
	ICMPCode_7  ICMPCode = 7
	ICMPCode_8  ICMPCode = 8
	ICMPCode_9  ICMPCode = 9
	ICMPCode_10 ICMPCode = 10
	ICMPCode_11 ICMPCode = 11
	ICMPCode_12 ICMPCode = 12
	ICMPCode_13 ICMPCode = 13
	ICMPCode_14 ICMPCode = 14
	ICMPCode_15 ICMPCode = 15
)

type TMOD1 struct {
	TokenRate          datatype.Float32    `avp:"Token-Rate" json:"tokenRate"`
	BucketDepth        datatype.Float32    `avp:"Bucket-Depth" json:"bucketDepth"`
	PeakTrafficRate    datatype.Float32    `avp:"Peak-Traffic-Rate" json:"peakTrafficRate"`
	MinimumPolicedUnit datatype.Unsigned32 `avp:"Minimum-Policed-Unit" json:"minimumPolicedUnit"`
	MaximumPacketSize  datatype.Unsigned32 `avp:"Maximum-Packet-Size" json:"maximumPacketSize"`
}
type TMOD2 struct {
	TokenRate          datatype.Float32    `avp:"Token-Rate" json:"tokenRate"`
	BucketDepth        datatype.Float32    `avp:"Bucket-Depth" json:"bucketDepth"`
	PeakTrafficRate    datatype.Float32    `avp:"Peak-Traffic-Rate" json:"peakTrafficRate"`
	MinimumPolicedUnit datatype.Unsigned32 `avp:"Minimum-Policed-Unit" json:"minimumPolicedUnit"`
	MaximumPacketSize  datatype.Unsigned32 `avp:"Maximum-Packet-Size" json:"maximumPacketSize"`
}
