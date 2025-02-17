/* Adapted from: nnrf-api version Sep, 2019
 * NRF NFDiscovery Service
 *
 * NRF NFDiscovery  Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha2

import "net"

// PCSC information configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable PCSC
//
//	 Data model:
//	   Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.53 Type: PcscfInfo, Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby PCSC
type PcscfInfo struct {
	// If included, this IE shall contain the access type (3GPP_ACCESS and/or NON_3GPP_ACCESS) supported by the P-CSCF.
	// If not included, it shall be assumed that all access types are supported
	// Optional
	AnType []AccessType `mapstructure:"accessType" json:"accessType,omitempty"`
	// DNNs supported by the P-CSCF.
	// If not provided, the P-CSCF can serve any DNN
	// Optional
	DnnList []string `mapstructure:"dnnList" json:"dnnList,omitempty"`
}

// PCSC nf profile from NRF or local static configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable PCSC
//
//	 Data model:
//	   Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby PCSC
type PcscfNfProfile struct {
	// PCSCF's Profile
	// Mandatory
	NfProfile NfProfile `mapstructure:"nfProfile" json:"nfProfile"`
	// This IE is used to select PCSCF
	// Mandatory
	PcscfInfo PcscfInfo `mapstructure:"pcscfInfo" json:"pcscfInfo"`
}

// Deprecated
type PcscfRuleDesc string

// Deprecated
const (
	PcscfRule_Compare_DNN        PcscfRuleDesc = "DNN"        //match DNN only
	PcscfRule_Compare_SNSSAI     PcscfRuleDesc = "SNSSAI"     //match SNSSAI only
	PcscfRule_Compare_DNN_SNSSAI PcscfRuleDesc = "DNN_SNSSAI" //match both DNN and SNSSAI
)

// Deprecated
type PcscfSelectionPolicy struct {
	// List of PolicyMatchingRules
	MatchingRules []PolicyMatchingRule `mapstructure:"matchingRules" json:"matchingRules"`
}

const (
	UP   = "UP"
	DOWN = "DOWN"
)

// YTL support cli show pcscf ip status online or offline
type AddressInfo struct {
	IP     net.IP `mapstructure:"ip" json:"ip"`
	Status string `mapstructure:"status" json:"status"`
}
