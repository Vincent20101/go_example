package v1alpha2

// AMF N2InterfaceAmfInfo configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable AMF
//
//	 Data model:
//	   Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.26 Type: N2InterfaceAmfInfo, Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby AMF
type N2InterfaceAmfInfo struct {
	// Available AMF endpoint IPv4 address(es) for N2
	// Optional
	Ipv4EndpointAddress []string `mapstructure:"ipv4EndpointAddress" json:"ipv4EndpointAddress,omitempty"`
	// Available AMF endpoint IPv6 address(es) for N2
	// Optional
	Ipv6EndpointAddress []string `mapstructure:"ipv6EndpointAddress" json:"ipv6EndpointAddress,omitempty"`
	// AMF Name
	// Optional
	AmfName *string `mapstructure:"amfName" json:"amfName,omitempty"`
}

// AMF information configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable AMF
//
//	 Data model:
//	   Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.11 Type: AmfInfo Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby AMF
type AmfInfo struct {
	// AMF set identifier.
	// Mandatory
	AmfSetId string `mapstructure:"amfSetId" json:"amfSetId"`
	// AMF region identifier
	// Mandatory
	AmfRegionId string `mapstructure:"amfRegionId" json:"amfRegionId"`
	// List of supported GUAMIs
	// Mandatory
	GuamiList []Guami `mapstructure:"guamiList" json:"guamiList"`
	// The list of TAIs the AMF can serve. It may contain the non-3GPP access TAI. The absence of this attribute and the taiRangeList attribute indicate that the AMF can be selected for any TAI in the serving network.
	// Optional
	TaiList []Tai `mapstructure:"taiList" json:"taiList,omitempty"`
	// The range of TAIs the AMF can serve. The absence of this attribute and the taiList attribute indicate that the AMF can be selected for any TAI in the serving network.
	// Optional
	TaiRangeList []TaiRange `mapstructure:"taiRangeList" json:"taiRangeList,omitempty"`
	// List of GUAMIs for which the AMF acts as a backup for AMF failure
	// Optional
	BackupInfoAmfFailure []Guami `mapstructure:"backupInfoAmfFailure" json:"backupInfoAmfFailure,omitempty"`
	// List of GUAMIs for which the AMF acts as a backup for planned AMF removal
	// Optional
	BackupInfoAmfRemoval []Guami `mapstructure:"backupInfoAmfRemoval" json:"backupInfoAmfRemoval,omitempty"`
	// N2 interface information of the AMF. This information needs not be sent in NF Discovery responses. It may be used by the NRF to update the DNS for AMF discovery by the 5G Access Network. The procedures for updating the DNS are out of scope of this specification.
	// Optional
	N2InterfaceAmfInfo *N2InterfaceAmfInfo `mapstructure:"n2InterfaceAmfInfo" json:"n2InterfaceAmfInfo,omitempty"`
}

// AMF nf profile from NRF or local static configuration
//
//	 Purpose:
//	   SMF can support interaction with locally configurable AMF
//
//	 Data model:
//	   Refer to the description for each attribute below
//
//	 Usage:
//		  SMF provides configurable operator to select standby AMF
type AmfNfProfile struct {
	// AMF's Profile
	// Mandatory
	NfProfile NfProfile `mapstructure:"nfProfile" json:"nfProfile"`
	// This IE is used to select AMF
	// Mandatory
	AmfInfo AmfInfo `mapstructure:"amfInfo" json:"amfInfo"`
}
