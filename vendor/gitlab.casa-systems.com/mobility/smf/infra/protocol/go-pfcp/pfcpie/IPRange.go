package pfcpie

import "net"

const (
	IPv4Rangelen = 8
	IPv6Rangelen = 32
)

// IPRange ...
type IPRange struct {
	StartAddr net.IP `json:"startAddr,omitempty"`
	EndAddr   net.IP `json:"endAddr,omitempty"`
}
