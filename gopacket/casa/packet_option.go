package main

import (
	"net"

	"github.com/google/gopacket/layers"
)

var (
	defaultMac = net.HardwareAddr{0xF0, 0x2E, 0x51, 0x01, 0x02, 0x3}
	defaultIP  = net.IPv4(127, 0, 0, 1)
)

type Option interface {
	apply(*packetOption)
}

type funcPacketOption struct {
	f func(*packetOption)
}

func (fpo *funcPacketOption) apply(po *packetOption) {
	fpo.f(po)
}

func NewFuncPacketOption(f func(*packetOption)) *funcPacketOption {
	return &funcPacketOption{
		f: f,
	}
}

type packetOption struct {
	SrcMAC, DstMAC         net.HardwareAddr
	SrcIP, DstIP           net.IP
	TCPSrcPort, TCPDstPort layers.TCPPort
	UDPSrcPort, UDPDstPort layers.UDPPort
	L3Protocol             layers.IPProtocol
	L4Protocol             layers.IPProtocol
	Window                 uint16
}

func defaultPacketOptions() packetOption {
	return packetOption{
		SrcMAC:     getNetworkMac(),
		DstMAC:     layers.EthernetBroadcast,
		SrcIP:      defaultIP,
		DstIP:      defaultIP,
		TCPSrcPort: layers.TCPPort(8080),
		TCPDstPort: layers.TCPPort(80),
		UDPSrcPort: layers.UDPPort(8081),
		UDPDstPort: layers.UDPPort(81),
		L3Protocol: layers.IPProtocolIPv4,
		L4Protocol: layers.IPProtocolTCP,
		Window:     501,
	}
}

func WithSrcMAC(srcMac net.HardwareAddr) Option {
	return NewFuncPacketOption(func(po *packetOption) {
		po.SrcMAC = srcMac
	})
}

func WithDstMAC(dstMAC net.HardwareAddr) Option {
	return NewFuncPacketOption(func(po *packetOption) {
		po.DstMAC = dstMAC
	})
}

func WithSrcIP(srcIP net.IP) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.SrcIP = srcIP
	})
}

func WithDstIP(dstIP net.IP) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.DstIP = dstIP
	})
}

func WithTCPSrcPort(srcPort uint16) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.TCPSrcPort = layers.TCPPort(srcPort)
	})
}

func WithTCPDstPort(dstPort uint16) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.TCPDstPort = layers.TCPPort(dstPort)
	})
}

func WithUDPSrcPort(srcPort uint16) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.UDPSrcPort = layers.UDPPort(srcPort)
	})
}

func WithUDPDstPort(dstPort uint16) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.UDPDstPort = layers.UDPPort(dstPort)
	})
}

func WithL3Protocol(i layers.IPProtocol) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.L3Protocol = i
	})
}

func WithL4Protocol(i layers.IPProtocol) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.L4Protocol = i
	})
}

func WithWindow(window uint16) Option {
	return NewFuncPacketOption(func(p *packetOption) {
		p.Window = window
	})
}

func getNetworkMac() (macAddr net.HardwareAddr) {
	infts, err := net.Interfaces()
	if err != nil {
		macAddr = defaultMac
		return
	}
	for _, inft := range infts {
		macAddr = inft.HardwareAddr
		if len(macAddr) == 0 {
			continue
		}
		break
	}
	if len(macAddr) == 0 {
		macAddr = defaultMac
	}
	return
}
