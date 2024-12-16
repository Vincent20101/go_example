package v1alpha2

// Dhcp Profile configuration
//
//  Purpose:
//    This is default dhcp configuration to be used by SMF when acting as a dhcp client and interacting with external dhcp server and when SMF act as dhcp server
//
//  Data model:
//    Refer to the description for each attribute below
//
//  Usage:
//    Used as value of dhcp
//

type DhcpProfile struct {
	// DHCP client profile for contacting External DHCP server profile \n
	DhcpClientProfile *DhcpClientProfile `mapstructure:"dhcpClientProfile" json:"dhcpClientProfile,omitempty"`
	// DHCP server profile for UE requesting
	DhcpServerProfile *DhcpServerProfile `mapstructure:"dhcpServerProfile" json:"dhcpServerProfile,omitempty"`
}

// Profile used by SMF as DHCP client
//
// Purpose:
//
//	This is default dhcp configuration to be used by SMF as DHCP Client when interacting with external dhcp
//
// Data model:
//
//	Refer to the description for each attribute below
//
// Usage:
//
//	Used as value of dhcp
type DhcpClientProfile struct {
	// SMF send to external DHCPv4 server fqdn  \n
	// Mandatory
	V4Fqdn string `mapstructure:"v4Fqdn" json:"v4Fqdn"`
	// Indicate the requested lease time, 0xffffffff would be to request an infinite lease \n
	// Unit is second \n
	// Mandatory
	V4LeaseTimeSec uint32 `mapstructure:"v4LeaseTimeSec" json:"v4LeaseTimeSec"`
	// SMF send to external DHCPv6 server fqdn  \n
	// Mandatory
	V6Fqdn string `mapstructure:"v6Fqdn" json:"v6Fqdn"`
	// Indicate the requested lease time, 0xffffffff would be to request an infinite lease \n
	// Unit is second \n
	// Mandatory
	V6LeaseTimeSec uint32 `mapstructure:"v6LeaseTimeSec" json:"v6LeaseTimeSec"`
}

// Profile used by SMF as DHCP server
//
// Purpose:
//
//	This is default dhcp configuration to be used by SMF when act as a DHCP server for UE
//
// Data model:
//
//	Refer to the description for each attribute below
//
// Usage:
//
//	Used as value of dhcp server
type DhcpServerProfile struct {
	// Indicate the lease time of DHCPv4, 0xffffffff would be to request an infinite lease \n
	// Unit is second \n
	// Mandatory
	V4LeaseTimeSec uint32 `mapstructure:"v4LeaseTimeSec" json:"v4LeaseTimeSec"`
	// Indicate the lease time of DHCPv6, 0xffffffff would be to request an infinite lease \n
	// Unit is second \n
	// Mandatory
	V6LeaseTimeSec uint32 `mapstructure:"v6LeaseTimeSec" json:"v6LeaseTimeSec"`
}

// Dhcp service configuration
//
//	Purpose:
//	  This is default service configuration to be used by dhcp
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of DhcpEndpointCfg
type DhcpEndpointCfg struct {
	// V4 PublicIp(s) for DHCP server to send message to SMF \n
	// Optional
	V4PublicIP string `mapstructure:"v4PublicIP"    json:"v4PublicIP,omitempty"`
	// V6 PublicIp(s) for DHCP server to send message to SMF \n
	// Optional
	V6PublicIP string `mapstructure:"v6PublicIP"    json:"v6PublicIP,omitempty"`
	// Port for DHCPV4 client to send message to SMF \n
	// Default value 8068 \n
	// Optional
	V4Port uint16 `mapstructure:"v4Port"    json:"v4Port,omitempty"`
	// Port for DHCPV6 client to send message to SMF \n
	// Default value 8546 \n
	// Optional
	V6Port uint16 `mapstructure:"v6Port"    json:"v6Port,omitempty"`
}
