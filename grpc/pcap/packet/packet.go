package packet

import (
	"errors"
	"log"
	"os"
	"syscall"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

type Layer struct {
	eth  layers.Ethernet
	ipv4 layers.IPv4
	ipv6 layers.IPv6
	tcp  layers.TCP
	udp  layers.UDP

	opts packetOption
}

func NewLayer(opts ...Option) *Layer {

	pl := &Layer{
		opts: defaultPacketOptions(),
	}

	for _, opt := range opts {
		opt.apply(&pl.opts)
	}

	switch pl.opts.L3Protocol {
	case layers.IPProtocolIPv4:
		pl.ipv4Packet()
	case layers.IPProtocolIPv6:
		pl.ipv6Packet()
	}

	return pl
}

func (pl *Layer) ipv4Packet() {
	pl.eth = layers.Ethernet{
		SrcMAC:       pl.opts.SrcMAC,
		DstMAC:       pl.opts.DstMAC,
		EthernetType: layers.EthernetTypeIPv4,
	}
	pl.ipv4 = layers.IPv4{
		SrcIP:   pl.opts.SrcIP,
		DstIP:   pl.opts.DstIP,
		Version: 4,
		TTL:     64,
	}

	switch pl.opts.L4Protocol {
	case layers.IPProtocolTCP:
		pl.ipv4.Protocol = layers.IPProtocolTCP
		pl.tcp = layers.TCP{
			SrcPort: pl.opts.TCPSrcPort,
			DstPort: pl.opts.TCPDstPort,
			ACK:     true,
			PSH:     true,
			Window:  pl.opts.Window,
			Seq:     1,
			Ack:     1,
		}
	case layers.IPProtocolUDP:
		pl.ipv4.Protocol = layers.IPProtocolUDP
		pl.udp = layers.UDP{
			SrcPort: pl.opts.UDPSrcPort,
			DstPort: pl.opts.UDPDstPort,
		}
	}
}

func (pl *Layer) ipv6Packet() {
	pl.eth = layers.Ethernet{
		SrcMAC:       pl.opts.SrcMAC,
		DstMAC:       pl.opts.DstMAC,
		EthernetType: layers.EthernetTypeIPv6,
	}
	pl.ipv6 = layers.IPv6{
		Version:  6,
		HopLimit: 64,
		SrcIP:    pl.opts.SrcIP,
		DstIP:    pl.opts.DstIP,
	}
	switch pl.opts.L4Protocol {
	case layers.IPProtocolTCP:
		pl.ipv6.NextHeader = layers.IPProtocolTCP
		pl.tcp = layers.TCP{
			SrcPort: pl.opts.TCPSrcPort,
			DstPort: pl.opts.TCPDstPort,
			ACK:     true,
			PSH:     true,
			Window:  pl.opts.Window,
			Seq:     1,
			Ack:     1,
		}
	case layers.IPProtocolUDP:
		pl.ipv6.NextHeader = layers.IPProtocolUDP
		pl.udp = layers.UDP{
			SrcPort: pl.opts.UDPSrcPort,
			DstPort: pl.opts.UDPDstPort,
		}
	}
}

func (pl *Layer) GeneratePacket(data []byte) ([]byte, error) {

	payload := gopacket.Payload(data)
	switch pl.opts.L4Protocol {
	case layers.IPProtocolTCP:
		if err := pl.tcp.SetNetworkLayerForChecksum(pl.getL3NetworkLayer()); err != nil {
			log.Printf("set tcp network layer checksum failure: [%+v]\n", err)
			return nil, err
		}
	case layers.IPProtocolUDP:
		if err := pl.udp.SetNetworkLayerForChecksum(pl.getL3NetworkLayer()); err != nil {
			log.Printf("set udp network layer checksum failure: [%+v]\n", err)
			return nil, err
		}
	}
	//serialize layers
	buf := gopacket.NewSerializeBuffer()
	op := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	l3 := pl.getL3SerializableLayer()
	l4 := pl.getL4SerializableLayer()
	if l3 != nil && l4 != nil {
		if err := gopacket.SerializeLayers(buf, op, &pl.eth, l3, l4, payload); err != nil {
			log.Printf("go packet serialize layers failure: [%+v]\n", err)
			return nil, err
		}
	} else {
		log.Printf("layer 3 or layer 4 is nil, failure to write packet\n")
		return nil, errors.New("layer 3 or layer 4 is nil, failure to write packet\n")
	}

	return buf.Bytes(), nil
}

var testFlagNum int

// WriteToPcapFile for test func
func (pl *Layer) WriteToPcapFile(data ...[]byte) (err error) {
	if len(data) == 0 {
		return
	}
	fileName := "/etc/config/test.pcap"
	file, err := os.OpenFile(fileName, syscall.O_WRONLY|syscall.O_CREAT|syscall.O_APPEND, 0666)
	if err != nil {
		log.Printf("file error: %v\n", err)
		return errors.New("file error")
	}
	defer file.Close()
	w := pcapgo.NewWriter(file)
	if testFlagNum == 0 { //write one file header
		if err = w.WriteFileHeader(1024, layers.LinkTypeEthernet); err != nil {
			log.Printf("write pcap header failure: [%+v]\n", err)
			return
		}
		testFlagNum++
	}
	for k, v := range data {
		if err = w.WritePacket(gopacket.CaptureInfo{
			Timestamp:      time.Time{},
			CaptureLength:  len(v),
			Length:         len(v),
			InterfaceIndex: 0,
			AncillaryData:  nil,
		}, v); err != nil {
			log.Printf("write packet failure: [%+v], data is : [%d -- %+v]\n", err, k, v)
			continue
		}
	}
	return nil
}

func (pl *Layer) getL3NetworkLayer() gopacket.NetworkLayer {
	switch pl.opts.L3Protocol {
	case layers.IPProtocolIPv4:
		return &pl.ipv4
	case layers.IPProtocolIPv6:
		return &pl.ipv6
	default:
		log.Printf("neither ipv4 or ipv6, unknown network layer\n")
		return nil
	}
}

func (pl *Layer) getL3SerializableLayer() gopacket.SerializableLayer {
	switch pl.opts.L3Protocol {
	case layers.IPProtocolIPv4:
		return &pl.ipv4
	case layers.IPProtocolIPv6:
		return &pl.ipv6
	default:
		log.Printf("neither ipv4 or ipv6, unknown layer\n")
		return nil
	}
}

func (pl *Layer) getL4SerializableLayer() gopacket.SerializableLayer {
	switch pl.opts.L4Protocol {
	case layers.IPProtocolTCP:
		return &pl.tcp
	case layers.IPProtocolUDP:
		return &pl.udp
	default:
		log.Printf("neither tcp or udp, unknown layer\n")
		return nil
	}
}
