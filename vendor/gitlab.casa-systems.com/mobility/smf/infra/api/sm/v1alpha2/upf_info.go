/* Adapted from: nnrf-api version Sep, 2019
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha2

// Dnn Upf Information Item
//
//	Purpose:
//	  Defines set of parameters supported by UPF for a given DNN
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.15 for details
//
//	Usage:
//	  Configured under SnssaiUpfInfoItem as a slice and referenced by UpfInfo
type DnnUpfInfoItem struct {
	// Supported DNN.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g30.zip Table 5.3.2-1 Dnn.\n
	// Mandatory.
	Dnn string `mapstructure:"dnn" json:"dnn"`
	// List of Data network access identifiers supported by the UPF for this DNN. The absence of this attribute indicates that the UPF can be selected for this DNN for any DNAI.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g30.zip Table 5.4.2-1 Dnai.\n
	// Default value nil.\n
	// Optional.
	DnaiList []string `mapstructure:"dnaiList" json:"dnaiList,omitempty"`
	// List of PDU session type(s) supported by the UPF for a specific DNN. The absence of this attribute indicates that the UPF can be selected for this DNN for any PDU session type supported by the UPF.\n
	// -"IPV4"
	// -"IPV6"
	// -"IPV4V6" (see clause 5.8.2.2.1 of 3GPP TS 23.501 [8])
	// -"UNSTRUCTURED"
	// -"ETHERNET"
	// Default value nil.\n
	// Optional.
	PduSessionTypes []PduSessionType `mapstructure:"pduSessionTypes" json:"pduSessionTypes,omitempty"`
	// List of ranges of static IPv4 addresses handled by UPF.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.22.\n
	// Default value nil.\n
	// Optional.
	Ipv4AddressRanges []Ipv4AddressRange `mapstructure:"ipv4AddressRanges" json:"ipv4AddressRanges,omitempty"`
	// List of ranges of static IPv6 prefixes handled by the UPF.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.23.\n
	// Only supports the prefix 64 bits for IPv6 now.\n
	// Default value nil.\n
	// Optional.
	Ipv6PrefixRanges []Ipv6PrefixRange `mapstructure:"ipv6PrefixRanges" json:"ipv6PrefixRanges,omitempty"`
	// UPF supported IP Pools references \n
	// Default value nil. \n
	// Optional.
	IpPools *UpfIpPoolRef `mapstructure:"ipPools" json:"ipPools,omitempty"`
}

// Snssai Upf Information Item
//
//	Purpose:
//	  Defines set of parameters supported by UPF for a given S-NSSAI
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.14 for details
//
//	Usage:
//	  Configured under UpfInfo as a map
type SnssaiUpfInfoItem struct {
	// Supported S-NSSAI.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g30.zip section 5.4.4.2.\n
	// Mandatory.
	SNssai Snssai `mapstructure:"sNssai" json:"sNssai"`
	// List of parameters supported by the UPF per DNN.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.15.\n
	// Mandatory.
	DnnUpfInfoList []DnnUpfInfoItem `mapstructure:"dnnUpfInfoList" json:"dnnUpfInfoList"`
	// Indicates whether the UPF supports redundant transport path on the transport layer in the corresponding network slice.
	// - true: supported
	// - false: not supported
	// Default value "false".
	// Optional.
	RedundantTransport bool `json:"redundantTransport,omitempty"`
}

/*Enumeration: UPInterfaceType
Table 6.1.6.3.9-1: Enumeration UPInterfaceType
Enumeration value	Description
"N3"	User Plane Interface: N3
"N6"	User Plane Interface: N6
"N9"	User Plane Interface: N9

//Below is the custom UP Interface Type:
"S5U"	User Plane Interface for S5U, it applied at 4G PDN(SMF+PGW-C combo)
"S8U"	User Plane Interface for S8U, it applied at 4G PDN(SMF+PGW-C combo)
*/

type UpInterfaceType string

// Interface Upf Information Item
//
//	Purpose:
//	  Defines information of a given IP interface of an UPF
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.24 for details
//
//	Usage:
//	  Configured under UpfInfo as a slice
type InterfaceUpfInfoItem struct {
	// User Plane interface type.\n
	// -"N3" represents User Plane Interface: N3.\n
	// -"N6" represents User Plane Interface: N6.\n
	// -"N9" represents User Plane Interface: N9.\n
	// -"S5U" represents User Plane Interface for S5U.\n
	// -"S8U" represents User Plane Interface for S8U.\n
	// Mandatory.
	InterfaceType UpInterfaceType `mapstructure:"interfaceType" json:"interfaceType"`
	// Available endpoint IPv4 address(es) of the User Plane interface.\n
	// Value can be a group of ipv4 address.\n
	// Default value nil.\n
	// Optional.
	Ipv4EndpointAddresses []Ipv4Addr `mapstructure:"ipv4EndpointAddresses" json:"ipv4EndpointAddresses,omitempty"`
	// Available endpoint IPv6 address(es) of the User Plane interface.\n
	// Value can be a group of ipv6 address.\n
	// Default value nil.\n
	// Optional.
	Ipv6EndpointAddresses []Ipv6Addr `mapstructure:"ipv6EndpointAddresses" json:"ipv6EndpointAddresses,omitempty"`
	// FQDN of available endpoint of the User Plane interface.\n
	// Default value "".\n
	// Optional.
	EndpointFqdn Fqdn `mapstructure:"endpointFqdn" json:"endpointFqdn,omitempty"`
	// Network Instance (See 3GPP 29.244 [21]) associated to the User Plane interface.\n
	// Default value "".\n
	// Optional.
	NetworkInstance string `mapstructure:"networkInstance" json:"networkInstance,omitempty"`
}

// Upf Information
//
//	Purpose:
//	  Defines information of an UPF NF Instance
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.13 for details
//
//	Usage:
//	  Configured under UpfNfProfile
type UpfInfo struct {
	// List of parameters supported by the UPF per S-NSSAI.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.14.\n
	// Mandatory.
	SNssaiUpfInfoList []SnssaiUpfInfoItem `mapstructure:"sNssaiUpfInfoList" json:"sNssaiUpfInfoList"`
	// List of SMF service area(s) the UPF can serve. If not provided, the UPF can serve any SMF service area.\n
	// Default value nil.\n
	// Optional.
	SmfServingArea []string `mapstructure:"smfServingArea" json:"smfServingArea,omitempty"`
	// List of User Plane interfaces configured on the UPF. When this IE is provided in the NF Discovery response, the NF Service Consumer (e.g. SMF) may use this information for UPF selection.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.24.\n
	// Default value nil.\n
	// Optional.
	InterfaceUpfInfoList []InterfaceUpfInfoItem `mapstructure:"interfaceUpfInfoList" json:"interfaceUpfInfoList,omitempty"`
	// Indicates whether interworking with EPS is supported by the UPF.\n
	// -"true" means interworking with EPS is supported by the UPF.\n
	// -"false" means interworking with EPS is not supported by the UPF.\n
	// Default value "false".\n
	// Optional.
	IwkEpsInd bool `mapstructure:"iwkEpsInd" json:"iwkEpsInd,omitempty"`
	// The list of TAIs the UPF can serve. It may contain the non-3GPP access TAI. If not provided, the UPF can serve the whole SMF service area defined by the smfServingArea attribute.\n
	// Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g30.zip section 5.4.4.4.\n
	// Default value nil.\n
	// Optional.
	TaiList []Tai `mapstructure:"taiList" json:"taiList,omitempty"`
	// Priority (relative to other NFs of the same type) in the range of 0-65535, to be used for NF selection for a service request matching the attributes of the UpfInfo; lower values indicate a higher priority.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip Table 6.1.6.2.13.
	// Optional.
	Priority *uint16 `mapstructure:"priority" json:"priority,omitempty"`
	// Indicates whether the UPF supports redundant GTP-U path.
	// -true: supported
	// -false (default): not supported
	// Default value "false".\n
	// Optional.
	RedundantGtpu bool `json:"redundantGtpu,omitempty"`
}

// Upf Nf Profile
//
//	Purpose:
//	  Defines information of a static UPF NF Instance
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used to define elements of UpfMap in UpfConfig
type UpfNfProfile struct {
	// UPF's Profile.
	// Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.510/29510-g30.zip Table 6.1.6.2.2.
	// Mandatory.
	NfProfile NfProfile `mapstructure:"nfProfile" json:"nfProfile"`
	// This IE is used to select UPF.
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip Table 6.1.6.2.13.
	// Optional.
	UpfInfo UpfInfo `mapstructure:"upfInfo" json:"upfInfo,omitempty"`
	// Multiple entries of UpfInfo. This attribute provides additional information to the upfInfo. This IE is used to select UPF.
	// The key of the map shall be a (unique) valid JSON string per clause 7 of IETF RFC 8259 [22], with a maximum of 32 characters.
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip Table 6.1.6.2.13.
	// Optional.
	UpfInfoList map[string]UpfInfo `mapstructure:"upfInfoList" json:"upfInfoList,omitempty"`
	// A flag indicating UP Function features.
	// Value should be uint64
	// Default value 0.
	// Optional.
	SupportFeature uint64 `mapstructure:"supportFeature" json:"supportFeature,omitempty"`
}

// Deprecated
type UpfRuleDesc string

// Deprecated
const (
	UPF_RULEDESC_DNN          UpfRuleDesc = "DNN"
	UPF_RULEDESC_DNN_SNSSAI   UpfRuleDesc = "DNN_SNSSAI"
	UPF_RULEDESC_LOCALITY     UpfRuleDesc = "LOCALITY"
	UPF_RULEDESC_SERVING_AREA UpfRuleDesc = "SERVING_AREA"
	UPF_RULEDESC_LOCATION     UpfRuleDesc = "LOCATION"
)

// Deprecated
type UpfSelectionPolicy struct {
	// List of PolicyMatchingRules
	SelectMode    PolicySelectMode     //policy selection mode
	MatchingRules []PolicyMatchingRule `mapstructure:"matchingRules" json:"matchingRules"`
}

type UpfSupportFeature uint64

const (
	UP_SUPPORT_FEATURE_BUCP  UpfSupportFeature = 0x1
	UP_SUPPORT_FEATURE_DDND  UpfSupportFeature = 0x2
	UP_SUPPORT_FEATURE_DLBD  UpfSupportFeature = 0x4
	UP_SUPPORT_FEATURE_TRST  UpfSupportFeature = 0x8
	UP_SUPPORT_FEATURE_FTUP  UpfSupportFeature = 0x10
	UP_SUPPORT_FEATURE_PFDM  UpfSupportFeature = 0x20
	UP_SUPPORT_FEATURE_HEEU  UpfSupportFeature = 0x40
	UP_SUPPORT_FEATURE_TREU  UpfSupportFeature = 0x80
	UP_SUPPORT_FEATURE_EMPU  UpfSupportFeature = 0x100
	UP_SUPPORT_FEATURE_PDIU  UpfSupportFeature = 0x200
	UP_SUPPORT_FEATURE_UDBC  UpfSupportFeature = 0x400
	UP_SUPPORT_FEATURE_QUOAC UpfSupportFeature = 0x800
	UP_SUPPORT_FEATURE_TRACE UpfSupportFeature = 0x1000
	UP_SUPPORT_FEATURE_FRRT  UpfSupportFeature = 0x2000
	UP_SUPPORT_FEATURE_PFDE  UpfSupportFeature = 0x4000
	UP_SUPPORT_FEATURE_QFQM  UpfSupportFeature = 0x200000000
	UP_SUPPORT_FEATURE_GPQM  UpfSupportFeature = 0x400000000
)
