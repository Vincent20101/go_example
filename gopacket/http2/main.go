package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"syscall"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"

	"github.com/google/gopacket/pcapgo"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func main() {

	srcMac, _ := net.ParseMAC("fa:16:3e:2a:b6:07")
	distMac, _ := net.ParseMAC("fa:16:3e:a7:07:06")

	eth := layers.Ethernet{
		SrcMAC:       srcMac,
		DstMAC:       distMac,
		EthernetType: layers.EthernetTypeIPv6,
	}
	//ipLayer := layers.IPv4{
	//	SrcIP:    net.ParseIP("192.168.0.100"),
	//	DstIP:    net.ParseIP("192.168.0.101"),
	//	Version:  4,
	//	TTL:      64,
	//	Protocol: layers.IPProtocolTCP,
	//}
	eth1 := layers.Ethernet{
		SrcMAC:       distMac,
		DstMAC:       srcMac,
		EthernetType: layers.EthernetTypeIPv6,
	}
	ipLayer6 := layers.IPv6{
		Version:    6,
		Length:     0,
		NextHeader: layers.IPProtocolTCP,
		HopLimit:   64,
		SrcIP:      net.ParseIP("4:1:1::2"),
		DstIP:      net.ParseIP("13:1:3::1"),
	}

	ipLayer62 := layers.IPv6{
		Version:    6,
		Length:     0,
		NextHeader: layers.IPProtocolTCP,
		HopLimit:   64,
		SrcIP:      net.ParseIP("13:1:3::1"),
		DstIP:      net.ParseIP("4:1:1::2"),
	}

	tcpLayer := layers.TCP{
		SrcPort: layers.TCPPort(8080),
		DstPort: layers.TCPPort(80),
		ACK:     true,
		PSH:     true,
		Window:  501,
		Seq:     50,
	}

	tcpLayer2 := layers.TCP{
		SrcPort: layers.TCPPort(8081),
		DstPort: layers.TCPPort(80),
		ACK:     true,
		PSH:     true,
		Window:  500,
		Seq:     2341,
	}

	data := getHttpInfo()
	payload := gopacket.Payload(data)

	err := tcpLayer.SetNetworkLayerForChecksum(&ipLayer6)
	err = tcpLayer2.SetNetworkLayerForChecksum(&ipLayer62)
	if err != nil {
		panic(err)
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	//err = gopacket.SerializeLayers(buf, opts, &eth, &ipLayer, &tcpLayer, payload)
	err = gopacket.SerializeLayers(buf, opts, &eth, &ipLayer6, &tcpLayer, payload)
	if err != nil {
		panic(err)
	}
	pd := buf.Bytes()

	data2 := getHttpInfo2()

	/*//decode http2
	framer := http2.NewFramer(nil, bytes.NewReader(data2))
	frame, err := framer.ReadFrame()
	if err != nil {
		fmt.Println(err)
	}
	var headersFrame *http2.HeadersFrame
	headersFrame, ok := frame.(*http2.HeadersFrame)
	if !ok {
		fmt.Println(err)
	}
	headers := headersFrame.HeaderBlockFragment()
	headerss, err := hpack.NewDecoder(4096, nil).DecodeFull(headers)
	if err != nil {
		fmt.Println(err)
	}
	for _, header := range headerss {
		fmt.Println(header.Name, header.Value)
	}
	return
	*/
	payload2 := gopacket.Payload(data2)

	buf2 := gopacket.NewSerializeBuffer()
	opts2 := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err = gopacket.SerializeLayers(buf2, opts2, &eth1, &ipLayer62, &tcpLayer2, payload2)
	if err != nil {
		panic(err)
	}
	pd2 := buf2.Bytes()

	fileName := "./tests17.pcap"
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
	w.WritePacket(gopacket.CaptureInfo{
		Timestamp:      time.Time{},
		CaptureLength:  len(pd2),
		Length:         len(pd2),
		InterfaceIndex: 0,
		AncillaryData:  nil,
	}, pd2)

}

func getHttpInfo() []byte {
	// http1.1
	//str := "POST /api/goods HTTP/1.1\r\nUser-Agent: PostmanRuntime/7.26.8\r\nAccept: */*\r\nPostman-Token: 781af5fc-37ad-4af7-aacd-db0a91478434\r\nHost: 47.108.21.107:9501\r\nAccept-Encoding: gzip, deflate, br\r\nConnection: keep-alive\r\nContent-Type: application/json\r\nContent-Length: 23\r\n\r\n{\"name\":\"lhb\",\"age\":11}"
	//return []byte(str)

	// http2
	var b []byte
	var buf3 bytes.Buffer
	e := hpack.NewEncoder(&buf3)
	e.WriteField(hpack.HeaderField{
		Name:      ":path",
		Value:     "/api/goods",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      ":method",
		Value:     "POST",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      ":status",
		Value:     "200",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      "content-type",
		Value:     "application/json",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      "content-length",
		Value:     "23",
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
	b = append(b, buf.Bytes()...)
	fmt.Println(buf.Bytes())

	buf2 := new(bytes.Buffer)
	fr2 := http2.NewFramer(buf2, buf2)
	str := `{"name":"lhb","age":11}`
	data2 := []byte(str)
	fr2.WriteData(2315461, true, data2)
	b = append(b, buf2.Bytes()...)
	return b
}

func getHttpInfo2() []byte {
	// http1.1
	//str := "POST /api/goods HTTP/1.1\r\nUser-Agent: PostmanRuntime/7.26.8\r\nAccept: */*\r\nPostman-Token: 781af5fc-37ad-4af7-aacd-db0a91478434\r\nHost: 47.108.21.107:9501\r\nAccept-Encoding: gzip, deflate, br\r\nConnection: keep-alive\r\nContent-Type: application/json\r\nContent-Length: 23\r\n\r\n{\"name\":\"lhb\",\"age\":11}"
	//return []byte(str)

	// http2
	var b []byte
	var buf3 bytes.Buffer
	e := hpack.NewEncoder(&buf3)
	e.WriteField(hpack.HeaderField{
		Name:      ":path",
		Value:     "/api/goodss",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      ":method",
		Value:     "POST",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      ":status",
		Value:     "200",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      "content-type",
		Value:     "application/json",
		Sensitive: false,
	})
	e.WriteField(hpack.HeaderField{
		Name:      "content-length",
		Value:     "23",
		Sensitive: false,
	})

	buf := new(bytes.Buffer)
	fr := http2.NewFramer(buf, buf)
	header := http2.HeadersFrameParam{
		StreamID:      2315462,
		BlockFragment: buf3.Bytes(),
		EndStream:     false,
		EndHeaders:    true,
		PadLength:     0,
		Priority:      http2.PriorityParam{},
	}
	fr.WriteHeaders(header)
	b = append(b, buf.Bytes()...)
	fmt.Println(buf.Bytes())

	buf2 := new(bytes.Buffer)
	fr2 := http2.NewFramer(buf2, buf2)
	str := `{"name":"lhb","age":11}`
	data2 := []byte(str)
	fmt.Println("data2: ", len(data2))
	fr2.WriteData(2315462, true, data2)
	b = append(b, buf2.Bytes()...)
	return b
}
