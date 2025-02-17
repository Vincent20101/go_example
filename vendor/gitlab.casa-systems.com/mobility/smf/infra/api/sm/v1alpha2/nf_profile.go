/* Adapted from: nnrf-api version Sep, 2019
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha2

type PlmnSnssai struct {
	// PLMN ID for which list of supported S-NSSAI(s) is provided.
	PlmnId PlmnId `mapstructure:"plmnId" json:"plmnId"`
	// The specific list of S-NSSAIs supported by the given PLMN.
	SNssaiList []Snssai `mapstructure:"sNssaiList" json:"sNssaiList,omitempty"`
}

type ServiceName = string

const (
	ServiceName_nnrf_nfm                  ServiceName = "nnrf-nfm"
	ServiceName_nnrf_disc                 ServiceName = "nnrf-disc"
	ServiceName_nudm_sdm                  ServiceName = "nudm-sdm"
	ServiceName_nudm_uecm                 ServiceName = "nudm-uecm"
	ServiceName_nudm_ueau                 ServiceName = "nudm-ueau"
	ServiceName_nudm_ee                   ServiceName = "nudm-ee"
	ServiceName_nudm_pp                   ServiceName = "nudm-pp"
	ServiceName_namf_comm                 ServiceName = "namf-comm"
	ServiceName_namf_evts                 ServiceName = "namf-evts"
	ServiceName_namf_mt                   ServiceName = "namf-mt"
	ServiceName_namf_loc                  ServiceName = "namf-loc"
	ServiceName_nsmf_pdusession           ServiceName = "nsmf-pdusession"
	ServiceName_nsmf_event_exposure       ServiceName = "nsmf-event-exposure"
	ServiceName_nsmf_nidd                 ServiceName = "nsmf-nidd"
	ServiceName_nausf_auth                ServiceName = "nausf-auth"
	ServiceName_nausf_sorprotection       ServiceName = "nausf-sorprotection"
	ServiceName_nnef_pfdmanagement        ServiceName = "nnef-pfdmanagement"
	ServiceName_nnef_smcontext            ServiceName = "nnef-smcontext"
	ServiceName_npcf_am_policy_control    ServiceName = "npcf-am-policy-control"
	ServiceName_npcf_smpolicycontrol      ServiceName = "npcf-smpolicycontrol"
	ServiceName_npcf_policyauthorization  ServiceName = "npcf-policyauthorization"
	ServiceName_npcf_bdtpolicycontrol     ServiceName = "npcf-bdtpolicycontrol"
	ServiceName_npcf_eventexposure        ServiceName = "npcf-eventexposure"
	ServiceName_npcf_ue_policy_control    ServiceName = "npcf-ue-policy-control"
	ServiceName_nsmsf_sms                 ServiceName = "nsmsf-sms"
	ServiceName_nnssf_nsselection         ServiceName = "nnssf-nsselection"
	ServiceName_nnssf_nssaiavailability   ServiceName = "nnssf-nssaiavailability"
	ServiceName_nudr_dr                   ServiceName = "nudr-dr"
	ServiceName_nlmf_loc                  ServiceName = "nlmf-loc"
	ServiceName_n5g_eir_eic               ServiceName = "n5g-eir-eic"
	ServiceName_nbsf_management           ServiceName = "nbsf-management"
	ServiceName_nchf_spendinglimitcontrol ServiceName = "nchf-spendinglimitcontrol"
	ServiceName_nchf_convergedcharging    ServiceName = "nchf-convergedcharging"
	ServiceName_nchf_offlineonlycharging  ServiceName = "nchf-offlineonlycharging"
	ServiceName_nnwdaf_eventssubscription ServiceName = "nnwdaf-eventssubscription"
	ServiceName_nnwdaf_analyticsinfo      ServiceName = "nnwdaf-analyticsinfo"
)

// NfServiceVersion configuration
//
//	Purpose:
//
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.19 Type: NFServiceVersion, Refer to the description for each attribute below
//
//	Usage:
type NfServiceVersion struct {
	// Version of the service instance to be used in the URI for accessing the API (e.g. "v1").
	// Mandatory
	ApiVersionInUri string `mapstructure:"apiVersionInUri" json:"apiVersionInUri"`
	// Full version number of the API as specified in clause 4.3.1 of  1 29.501 [5].
	// Mandatory
	ApiFullVersion string `mapstructure:"apiFullVersion" json:"apiFullVersion"`

	// Expiry date and time of the NF service. This represents the planned retirement date as specified in clause 4.3.1.5 of  1 29.501 [5].
	//Expiry          *time.Time `json:"expiry,omitempty"`
}

type UriScheme = string

const (
	UriScheme_http  UriScheme = "http"
	UriScheme_https UriScheme = "https"
)

type IpEndPoint struct {
	// IPv4 address of the Network Function Service. At most one occurrence of either ipv4Address or ipv6Address shall be included in this data structure.
	Ipv4Address Ipv4Addr `mapstructure:"ipv4Address" json:"ipv4Address,omitempty"`
	// IPv6 address of the Network Function Service. At most one occurrence of either ipv4Address or ipv6Address shall be included in this data structure.
	Ipv6Address Ipv6Addr `mapstructure:"ipv6Address" json:"ipv6Address,omitempty"`
	// Transport protocol
	Transport TransportProtocol `mapstructure:"transport" json:"transport,omitempty"`
	// Port number. If the port number is absent from the ipEndPoints attribute, the NF service consumer shall use the default HTTP port number, i.e. TCP port 80 for "http" URIs or TCP port 443 for "https" URIs as specified in IETF RFC 7540 [9] when invoking the service.
	Port int32 `mapstructure:"port" json:"port,omitempty"`
}

// NF Service configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable NF Service
//
//	 Data model:
//	   Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.3 Type: NFService, Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby NF Service
type NfService struct {
	// Unique ID of the service instance within a given NF Instance
	// Mandatory
	ServiceInstanceId string `mapstructure:"serviceInstanceId" json:"serviceInstanceId"`
	// Name of the service instance (e.g. "nudm-sdm")
	// Mandatory
	ServiceName ServiceName `mapstructure:"serviceName" json:"serviceName"`
	// The API versions supported by the NF Service and if available, the corresponding retirement date of the NF Service.
	// The different array elements shall have distinct unique values for "apiVersionInUri", and consequently, the values of "apiFullVersion" shall have a unique first digit version number
	// Optional
	Versions []NfServiceVersion `mapstructure:"versions" json:"versions,omitempty"`
	// URI scheme (e.g. "http", "https")
	// Optional
	Scheme UriScheme `mapstructure:"scheme" json:"scheme,omitempty"`
	//NfServiceStatus                  NfServiceStatus                   `json:"nfServiceStatus"`

	// FQDN of the NF Service Instance
	// Optional
	Fqdn string `mapstructure:"fqdn" json:"fqdn,omitempty"`
	// IP address(es) and port information of the Network Function (including IPv4 and/or IPv6 address) where the service is listening for incoming service requests
	// Optional
	IpEndPoints []IpEndPoint `mapstructure:"ipEndPoints" json:"ipEndPoints,omitempty"`
	// Optional path segment(s) used to construct the {apiRoot} variable of the different API URIs, as described in 3GPP 29.501 [5], clause 4.4.1
	// Optional
	ApiPrefix string `mapstructure:"apiPrefix" json:"apiPrefix,omitempty"`
	//DefaultNotificationSubscriptions []DefaultNotificationSubscription `json:"defaultNotificationSubscriptions,omitempty"`

	// Static capacity information in the range of 0-65535, expressed as a weight relative to other services of the same type.
	// Optional
	Capacity *int32 `mapstructure:"capacity" json:"capacity,omitempty"`
	// Dynamic load information, ranged from 0 to 100, indicates the current load percentage of the NF Service
	// Optional
	Load *int32 `mapstructure:"load" json:"load,omitempty"`
	// Priority (relative to other services of the same type) in the range of 0-65535, to be used for NF Service selection; lower values indicate a higher priority. (NOTE 2).
	// The NRF may overwrite the received priority value when exposing an NFProfile with the Nnrf_NFDiscovery service.
	// Optional
	Priority *int32 `mapstructure:"priority" json:"priority,omitempty"`
	//RecoveryTime                     *time.Time                        `json:"recoveryTime,omitempty"`

	// Specific data for a CHF service instance
	// Optional
	ChfServiceInfo *ChfServiceInfo `mapstructure:"chfServiceInfo" json:"chfServiceInfo,omitempty"`
	// Supported Features of the NF Service instance
	// Optional
	SupportedFeatures string `mapstructure:"supportedFeatures" json:"supportedFeatures,omitempty"`
	// NF Service Set ID (see clause 28.11 of 3GPP TS 23.003 [12])
	// At most one NF Service Set ID shall be indicated per PLMN of the NF.
	// Optional
	NfServiceSetIdList []NfSetId `mapstructure:"nfServiceSetIdList" json:"nfServiceSetIdList,omitempty"`
	// It indicates whether the NF Instance requires Oauth2-based authorization.
	// Absence of this IE means that the NF Service Producer has not provided any indication about its usage of Oauth2 for authorization.
	// Optional
	Oauth2Required *bool `mapstructure:"oauth2Required" json:"oauth2Required,omitempty"`
}

// NF Profile configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable NF
//
//	 Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.2 Type: NFProfile, Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby NF
type NfProfile struct {
	// Unique identity of the NF Instance.
	// Mandatory
	NfInstanceId string `mapstructure:"nfInstanceId" json:"nfInstanceId"`
	// Type of Network Function
	// Mandatory
	NfType NfType `mapstructure:"nfType" json:"nfType"`
	// Status of the NF Instance
	// Mandatory
	NfStatus NfStatus `mapstructure:"nfStatus" json:"nfStatus"`
	// Human readable name of the NF Instance
	// Mandatory
	NfInstanceName string `mapstructure:"nfInstanceName" json:"nfInstanceName"`
	// PLMN(s) of the Network Function (NOTE 7).
	// This IE shall be present if this information is available for the NF.
	// If not provided, PLMN ID(s) of the PLMN of the NRF are assumed for the NF.
	// Optional
	PlmnList []PlmnId `mapstructure:"plmnList" json:"plmnList,omitempty"`
	// S-NSSAIs of the Network Function.
	// If not provided, the NF can serve any S-NSSAI.
	// When present this IE represents the list of S-NSSAIs supported in all the PLMNs listed in the plmnList IE
	// Optional
	SNssais []Snssai `mapstructure:"sNssais" json:"sNssais,omitempty"`
	// This IE may be included when the list of S-NSSAIs supported by the NF for each PLMN it is supporting is different. When present, this IE shall include the S-NSSAIs supported by the Network Function for each PLMN supported by the Network Function. When present, this IE shall override sNssais IE
	// Optional
	PerPlmnSnssaiList []PlmnSnssai `mapstructure:"perPlmnSnssaiList" json:"perPlmnSnssaiList,omitempty"`
	// NSI identities of the Network Function. If not provided, the NF can serve any NSI.
	// Optional
	NsiList []string `mapstructure:"nsiList" json:"nsiList,omitempty"`
	// FQDN of the Network Function. For AMF, the FQDN registered with the NRF shall be that of the AMF Name
	// Optional
	Fqdn string `mapstructure:"fqdn" json:"fqdn,omitempty"`
	// If the NF needs to be discoverable by other NFs in a different PLMN, then an FQDN that is used for inter-PLMN routing as specified in 3GPP 23.003 [12] shall be registered with the NRF (NOTE 8).
	// A change of this attribute shall result in triggering a "NF_PROFILE_CHANGED" notification from NRF towards subscribing NFs located in a different PLMN, but the new value shall be notified as a change of the "fqdn" attribute.
	// Optional
	InterPlmnFqdn string `mapstructure:"interPlmnFqdn" json:"interPlmnFqdn,omitempty"`
	// IPv4 address(es) of the Network Function
	// Optional
	Ipv4Addresses []Ipv4Addr `mapstructure:"ipv4Addresses" json:"ipv4Addresses,omitempty"`
	// IPv6 address(es) of the Network Function
	// Optional
	Ipv6Addresses []Ipv6Addr `mapstructure:"ipv6Addresses" json:"ipv6Addresses,omitempty"`
	// PLMNs allowed to access the NF instance. If not provided, any PLMN is allowed to access the NF
	// Optional
	AllowedPlmns []PlmnId `mapstructure:"allowedPlmns" json:"allowedPlmns,omitempty"`
	// Type of the NFs allowed to access the NF instance. If not provided, any NF type is allowed to access the NF\
	// Optional
	AllowedNfTypes []NfType `mapstructure:"allowedNfTypes" json:"allowedNfTypes,omitempty"`
	// Pattern (regular expression according to the ECMA-262 dialect [8]) representing the NF domain names allowed to access the NF instance. If not provided, any NF domain is allowed to access the NF
	// Optional
	AllowedNfDomains []string `mapstructure:"allowedNfDomains" json:"allowedNfDomains,omitempty"`
	// S-NSSAI of the allowed slices to access the NF instance. If not provided, any slice is allowed to access the NF
	// Optional
	AllowedNssais []Snssai `mapstructure:"allowedNssais" json:"allowedNssais,omitempty"`
	// Priority (relative to other NFs of the same type) in the range of 0-65535, to be used for NF selection; lower values indicate a higher priority. If priority is also present in the nfServiceList parameters, those will have precedence over this value.
	// Optional
	Priority *int32 `mapstructure:"priority" json:"priority,omitempty"`
	// Static capacity information in the range of 0-65535, expressed as a weight relative to other NF instances of the same type; if capacity is also present in the nfServiceList parameters, those will have precedence over this value.
	// Optional
	Capacity *int32 `mapstructure:"capacity" json:"capacity,omitempty"`
	// Dynamic load information, ranged from 0 to 100, indicates the current load percentage of the NF
	// Optional
	Load *int32 `mapstructure:"load" json:"load,omitempty"`
	// Operator defined information about the location of the NF instance (e.g. geographic location, data center)
	// Optional
	Locality string `mapstructure:"locality" json:"locality,omitempty"`
	// Specific data for custom Network Functions
	// Optional
	CustomInfo map[string]interface{} `mapstructure:"customInfo" json:"customInfo,omitempty"`
	// If present, and set to true, it indicates that the different service instances of a same NF Service in this NF instance, supporting a same API version, are capable to persist their resource state in shared storage and therefore these resources are available after a new NF service instance supporting the same API version is selected by a NF Service Consumer (see 3GPP 23.527 [27]).
	// Otherwise, it indicates that the NF Service Instances of a same NF Service are not capable to share resource state inside the NF Instance.
	// Optional
	NfServicePersistence bool `mapstructure:"nfServicePersistence" json:"nfServicePersistence,omitempty"`
	// List of NF Service Instances. It shall include the services produced by the NF that can be discovered by other NFs, if any.
	// Optional
	NfServices []NfService `mapstructure:"nfServices" json:"nfServices,omitempty"`
	// NF Set ID defined in clause 28.10 of 3GPP TS 23.003 [12]. At most one NF Set ID shall be indicated per PLMN of the NF
	// Optional
	NfSetIdList []NfSetId `mapstructure:"nfSetIdList" json:"nfSetIdList,omitempty"`
	// This IE only for UPFProfile
	// If present, it described that UPF exposed the ip address information of a T3 interface to Limgr
	// Optional
	T3Addr string `mapstructure:"t3Addr" json:"t3Addr,omitempty"`
	// Derived from validityTime. Not configurable
	// Optional
	Expire int64 `mapstructure:"expire" json:"expiry,omitempty" cmp:"ignore"`
}

// NF Profiles
//
//	Purpose:
//	  Show application level nf confiugration information (in the memory), and state/status information
type NfProfiles struct {
	// AMF nf profile
	// Optional
	AmfNfProfiles []*AmfNfProfile `mapstructure:"amfNfProfiles" json:"amfNfProfiles,omitempty"`
	// PCF nf profile
	// Optional
	PcfNfProfiles []*PcfNfProfile `mapstructure:"pcfNfProfiles" json:"pcfNfProfiles,omitempty"`
	// UPF nf profile
	// Optional
	UpfNfProfiles []*UpfNfProfile `mapstructure:"upfNfProfiles" json:"upfNfProfiles,omitempty"`
	// CHF nf profile
	// Optional
	ChfNfProfiles []*ChfNfProfile `mapstructure:"chfNfProfiles" json:"chfNfProfiles,omitempty"`
	// UDM nf profile
	// Optional
	UdmNfProfiles []*UdmNfProfile `mapstructure:"udmNfProfiles" json:"udmNfProfiles,omitempty"`
	// NEF nf profile
	// Optional
	NefNfProfiles []*NefNfProfile `mapstructure:"nefNfProfiles" json:"nefNfProfiles,omitempty"`
	// PCSC nf profile
	// Optional
	PcscfNfProfiles []*PcscfNfProfile `mapstructure:"pcscfNfProfiles" json:"pcscfNfProfiles,omitempty"`
	// Display application-level PCSCF information
	// Optional
	SmfPcscfInfos []*SmfPcscfInfo `mapstructure:"smfPcscfInfos" json:"smfPcscfInfos,omitempty"`
}

// SmfPcscfInfo
//
// Purpose:
//
//	Show application level pcscf confiugration information (in the memory), and state/status information
type SmfPcscfInfo struct {
	// PCSC nf profile from NRF or local static configuration
	// Mandatory
	Profile PcscfNfProfile `mapstructure:"profile" json:"profile,omitempty"`
	// YTL custom
	// AddressInfo supported by the P-CSCF.
	// SMF support cli show pcscf status, format ${ip}-${status}
	// Optional
	AddressList []string `mapstructure:"addressList" json:"addressList,omitempty"`
}
