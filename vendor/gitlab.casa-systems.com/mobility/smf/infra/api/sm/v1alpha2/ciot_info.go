package v1alpha2

// CIoTConfiguration
//
//	Purpose:
//	  This is CIoT(Cellular IoT) configuration.
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used as optional parameter "CIoTConfiguration" in DnnProfile
type CIoTConfiguration struct {
	// If the PDU Session Type is Unstructured and the CIoT is enabled, UL(UpLink) data is Non-IP data \n
	// For Non-IP data, SMF needs to provide DN(Data Network) ip address as the destination network for UL data forwarding \n
	// Default value is nil \n
	// Optional.
	DataNetworkAddress *DataNetworkAddress `mapstructure:"dataNetworkAddress" json:"dataNetworkAddress,omitempty"`
}

// DataNetworkAddress
//
//	Purpose:
//	  This is DataNetworkAddress configuration.
//
//	Data model:
//	  Refer to the description for each attribute below.
//
//	Usage:
//	  Used as optional parameter "DataNetworkAddress" in CIoTConfiguration
type DataNetworkAddress struct {
	// IPv6 address for CIoT Data Network Address \n
	// IPv6 Pattern: '^((:|(0?|([1-9a-f][0-9a-f]{0,3}))):)((0?|([1-9a-f][0-9a-f]{0,3})):){0,6}(:|(0?|([1-9a-f][0-9a-f]{0,3})))$' \n
	// And, \n
	// IPv6 Pattern: '^((([^:]+:){7}([^:]+))|((([^:]+:)*[^:]+)?::(([^:]+:)*[^:]+)?))$' \n
	// IPv6 Example: '2001:db8:85a3::8a2e:370:7334' \n
	// Default value "" \n
	// Mandatory
	Address string `mapstructure:"address"    json:"address"`
	// Port for CIoT Data Network Address \n
	// Default value 8080 \n
	// Mandatory
	Port uint16 `mapstructure:"port"    json:"port"`
}
