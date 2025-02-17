/* Adapted from: nnrf-api version Sep, 2019
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha2

// Smf Dnn Info Item configuration
//
//	Purpose:
//	  This is default Dnn Info Item configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of DnnSmfInfoList
type DnnSmfInfoItem struct {
	// Dnn \n
	// (please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g20.zip section 5.3.2).
	// Mandatory.
	Dnn string `mapstructure:"dnn" json:"dnn"`
	// List of Data network access identifiers supported by the SMF for this DNN.
	// Please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.510/29510-h40.zip Table 6.1.6.2.30-1.
	// Default value nil.<br>
	// Optional.
	DnaiList []string `mapstructure:"dnaiList" json:"dnaiList,omitempty"`
}

// Smf Snssai Info Item configuration
//
//	Purpose:
//	  This is default Snssai Info Item configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of SNssaiSmfInfoList
type SnssaiSmfInfoItem struct {
	// The S-NSSAI supported by the SMF \n
	// (please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g20.zip section 5.4.4.2).
	// Mandatory.
	SNssai Snssai `mapstructure:"sNssai" json:"sNssai"`
	// The list of Dnns supported by the SMF \n
	// Mandatory.
	DnnSmfInfoList []DnnSmfInfoItem `mapstructure:"dnnSmfInfoList" json:"dnnSmfInfoList"`
}

/* The absence of both the smfInfo and smfInfoExt attributes in an SMF profile indicates that the SMF can be selected for any S-NSSAI, DNN, TAI and access type. */

// Smf Info configuration
//
//	Purpose:
//	  This is default Smf Info configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of SmfInfo
type SmfInfo struct {
	// List of parameters supported by the SMF per S-NSSAI \n
	// Mandatory.
	SNssaiSmfInfoList []SnssaiSmfInfoItem `mapstructure:"sNssaiSmfInfoList" json:"sNssaiSmfInfoList"`
	// The list of TAIs the SMF can serve. It may contain the non-3GPP access TAI. The absence of this attribute and the taiRangeList attribute indicate that the SMF can be selected for any TAI in the serving network \n
	// Tracking Area Identity (please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g20.zip section 5.4.4.4).
	// Optional.
	TaiList []Tai `mapstructure:"taiList" json:"taiList,omitempty"`
	// The range of TAIs the SMF can serve. It may contain the non-3GPP access' TAI. The absence of this attribute and the taiList attribute indicate that the SMF can be selected for any TAI in the serving network. \n
	// Optional.
	TaiRangeList []TaiRange `mapstructure:"taiRangeList" json:"taiRangeList,omitempty"`
	// The FQDN of the PGW if the SMF is a combined SMF/PGW-C \n
	// Optional.
	PgwFqdn Fqdn `mapstructure:"pgwFqdn" json:"pgwFqdn,omitempty"`
	// If included, this IE shall contain the access type (3GPP_ACCESS and/or NON_3GPP_ACCESS) supported by the SMF. \n
	// If not included, it shall be assumed the both access types are supported. \n
	//  -"3GPP_ACCESS" represents 3GPP access (please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g20.zip section 5.4.3.1).
	//  -"NON_3GPP_ACCESS" represents Non-3GPP access (please refer https://www.3gpp.org/ftp//Specs/archive/29_series/29.571/29571-g20.zip section 5.4.3.1).
	// Optional.
	AccessType []AccessType `mapstructure:"accessType" json:"accessType,omitempty"`
}

// Smf Network Function configuration
//
//	Purpose:
//	  This is default Smf Network Function configuration to be used by SMF
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of SmfProfile
type SmfNfProfile struct {
	// SMF's Profile \n
	// Mandatory.
	NfProfile NfProfile `mapstructure:"nfProfile" json:"nfProfile"`
	// This IE is contains what SMF support \n
	// Optional.
	SmfInfo SmfInfo `mapstructure:"smfInfo" json:"smfInfo,omitempty"`
	// This IE is contains which ServingArea SMF support \n
	// Optional.
	SmfServingArea string `mapstructure:"smfServingArea" json:"smfServingArea,omitempty"`
	// This IE is contains which Locality SMF support \n
	// Optional.
	SmfLocality string `mapstructure:"smfLocality" json:"smfLocality,omitempty"`
}
