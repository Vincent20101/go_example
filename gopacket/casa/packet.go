package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
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
	seq  uint32

	opts packetOption
}

func NewLayer(opts ...Option) *Layer {

	//n, _ := rand.Int(rand.Reader, big.NewInt(1<<31))
	//seq := n.Int64()
	//if seq&1 == 0 {
	//	seq += 1
	//}

	pl := &Layer{
		opts: defaultPacketOptions(),
		//seq:  uint32(seq),
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
			Seq:     pl.seq,
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
			//Seq:     pl.seq + 1,
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
			fmt.Printf("set tcp network layer checksum failure: [%+v]\n", err)
			return nil, err
		}
	case layers.IPProtocolUDP:
		if err := pl.udp.SetNetworkLayerForChecksum(pl.getL3NetworkLayer()); err != nil {
			fmt.Printf("set udp network layer checksum failure: [%+v]\n", err)
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
			fmt.Printf("go packet serialize layers failure: [%+v]\n", err)
			return nil, err
		}
	} else {
		fmt.Printf("layer 3 or layer 4 is nil, failure to write packet\n")
		return nil, errors.New("layer 3 or layer 4 is nil, failure to write packet\n")
	}

	return buf.Bytes(), nil
}

func (pl *Layer) WriteToPcapFile(data ...[]byte) (err error) {
	if len(data) == 0 {
		return
	}
	fileName := "testabc.pcap"
	file, err := os.OpenFile(fileName, syscall.O_WRONLY|syscall.O_CREAT|syscall.O_APPEND, 0666)
	if err != nil {
		fmt.Println("file error:", err)
		return errors.New("file error")
	}
	defer file.Close()
	w := pcapgo.NewWriter(file)
	if err = w.WriteFileHeader(1024, layers.LinkTypeEthernet); err != nil {
		fmt.Printf("write pcap header failure: [%+v]\n", err)
		return
	}
	for k, v := range data {
		if err = w.WritePacket(gopacket.CaptureInfo{
			Timestamp:      time.Time{},
			CaptureLength:  len(v),
			Length:         len(v),
			InterfaceIndex: 0,
			AncillaryData:  nil,
		}, v); err != nil {
			fmt.Printf("write packet failure: [%+v], data is : [%d -- %+v]\n", err, k, v)
			continue
		}
	}
	return nil
}

func (pl *Layer) getL3NetworkLayer() gopacket.NetworkLayer {
	rand.Seed(time.Now().UnixNano())
	n := byte(rand.Intn(100)) + 1
	switch pl.opts.L3Protocol {
	case layers.IPProtocolIPv4:
		pl.ipv4.SrcIP = net.IP{192, 168, 0, n}
		return &pl.ipv4
	case layers.IPProtocolIPv6:
		//pl.ipv6.SrcIP = net.ParseIP("2001:db8::" + string(n))
		return &pl.ipv6
	default:
		fmt.Printf("neither ipv4 or ipv6, unknown network layer\n")
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
		fmt.Printf("neither ipv4 or ipv6, unknown layer\n")
		return nil
	}
}

func (pl *Layer) getL4SerializableLayer() gopacket.SerializableLayer {
	switch pl.opts.L4Protocol {
	case layers.IPProtocolTCP:
		//pl.seq++
		return &pl.tcp
	case layers.IPProtocolUDP:
		//pl.seq++
		return &pl.udp
	default:
		fmt.Printf("neither tcp or udp, unknown layer\n")
		return nil
	}
}
