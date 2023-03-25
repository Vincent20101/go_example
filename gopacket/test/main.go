package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"syscall"
	"time"

	gproto "github.com/vincent20101/go-example/grpc/gtest/proto"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"

	"github.com/google/gopacket/pcapgo"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"google.golang.org/protobuf/proto"
)

func main() {

	srcMac, _ := net.ParseMAC("fa:16:3e:2a:b6:07")
	distMac, _ := net.ParseMAC("fa:16:3e:a7:07:06")

	eth := layers.Ethernet{
		SrcMAC:       srcMac,
		DstMAC:       distMac,
		EthernetType: layers.EthernetTypeIPv4,
	}
	ipLayer := layers.IPv4{
		SrcIP:    net.ParseIP("192.168.0.100"),
		DstIP:    net.ParseIP("192.168.0.101"),
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}

	tcpLayer := layers.TCP{
		SrcPort: layers.TCPPort(8080),
		DstPort: layers.TCPPort(80),
		ACK:     true,
		PSH:     true,
		Window:  501,
		Seq:     50,
	}

	data := getHttpInfo()
	payload := gopacket.Payload(data)

	err := tcpLayer.SetNetworkLayerForChecksum(&ipLayer)
	if err != nil {
		panic(err)
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err = gopacket.SerializeLayers(buf, opts, &eth, &ipLayer, &tcpLayer, payload)
	if err != nil {
		panic(err)
	}
	pd := buf.Bytes()

	fileName := "./tests.pcap"
	file, err := os.OpenFile(fileName, syscall.O_WRONLY|syscall.O_CREAT, 0666)
	if err != nil {
		fmt.Println("file error:", err)
		return
	}
	defer file.Close()

	w := pcapgo.NewWriter(file)
	w.WriteFileHeader(1024, layers.LinkTypeEthernet)

	w.WritePacket(gopacket.CaptureInfo{
		Timestamp:      time.Time{},
		CaptureLength:  len(pd),
		Length:         len(pd),
		InterfaceIndex: 0,
		AncillaryData:  nil,
	}, pd)

}

func getHttpInfo() []byte {
	// http1.1
	//str := "POST /api/goods HTTP/1.1\r\nUser-Agent: PostmanRuntime/7.26.8\r\nAccept: */*\r\nPostman-Token: 781af5fc-37ad-4af7-aacd-db0a91478434\r\nHost: 47.108.21.107:9501\r\nAccept-Encoding: gzip, deflate, br\r\nConnection: keep-alive\r\nContent-Type: application/json\r\nContent-Length: 23\r\n\r\n{\"name\":\"lhb\",\"age\":11}"
	//return []byte(str)

	// http2
	var result []byte
	var buf3 bytes.Buffer
	e := hpack.NewEncoder(&buf3)
	e.WriteField(hpack.HeaderField{
		Name:      ":path",
		Value:     "/Greeter/SayHello",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      ":method",
		Value:     "POST",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      ":scheme",
		Value:     "http",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      "content-type",
		Value:     "application/grpc",
		Sensitive: false,
	})

	buf := new(bytes.Buffer)
	fr := http2.NewFramer(buf, buf)
	header := http2.HeadersFrameParam{
		StreamID:      2315461,
		BlockFragment: buf3.Bytes(),
		EndStream:     false,
		EndHeaders:    true,
		PadLength:     0,
		Priority:      http2.PriorityParam{},
	}
	fr.WriteHeaders(header)
	result = append(result, buf.Bytes()...)

	req := gproto.HelloRequest{
		Name: "linhuanbo1",
	}
	s, _ := proto.Marshal(&req)
	size := proto.Size(&req)
	fmt.Println(size)
	btest := []byte{0}
	bi := make([]byte, 4)
	binary.BigEndian.PutUint32(bi, uint32(len(s)))
	btest = append(btest, bi...)
	btest = append(btest, s...)
	buf4 := new(bytes.Buffer)
	fr4 := http2.NewFramer(buf4, buf4)
	fr4.WriteData(2315461, true, btest)
	fmt.Println(buf4.Bytes())
	result = append(result, buf4.Bytes()...)
	return result
}
