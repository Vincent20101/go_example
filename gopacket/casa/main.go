package main

import (
	"fmt"
	"net"

	"github.com/davecgh/go-spew/spew"

	"github.com/google/gopacket/layers"
)

func main() {
	pl := NewLayer(WithL4Protocol(layers.IPProtocolTCP), WithL3Protocol(layers.IPProtocolIPv4), WithDstIP(net.ParseIP("192.168.0.1")))
	fmt.Printf("%+v\n", spew.Sdump(pl))
	payload, _ := pl.GeneratePacket([]byte{1, 2, 3, 4})
	//pl2 := NewLayer(WithL4Protocol(layers.IPProtocolTCP), WithL3Protocol(layers.IPProtocolIPv4))
	//pl.seq = pl.seq + 4
	payload2, _ := pl.GeneratePacket([]byte{1, 2, 3, 5})
	//pl3 := NewLayer(WithL4Protocol(layers.IPProtocolTCP), WithL3Protocol(layers.IPProtocolIPv4))
	//payload3, _ := pl3.GeneratePacket([]byte{1, 2, 3, 6})
	//fmt.Println(payload)
	fmt.Println(pl.WriteToPcapFile(payload, payload2))
}
