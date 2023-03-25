package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/google/gopacket/pcapgo"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func getLocalIP(dstip net.IP) (net.IP, int) {
	serverAddr, err := net.ResolveUDPAddr("udp", dstip.String()+":23456")
	if err != nil {
		log.Fatal(err)
	}
	if con, err := net.DialUDP("udp", nil, serverAddr); err == nil {
		if udpaddr, ok := con.LocalAddr().(*net.UDPAddr); ok {
			fmt.Println(udpaddr)
			return udpaddr.IP, udpaddr.Port
		}
	}
	log.Fatal("could not get local ip: " + err.Error())
	return nil, -1
}

func main() {

	//发送端的 mac
	srcMac, _ := net.ParseMAC("fa:16:3e:2a:b6:07")
	//接收端
	distMac, _ := net.ParseMAC("fa:16:3e:a7:07:05")

	//发送端IP
	srcIP := net.ParseIP("192.168.0.204")
	dstIP, _ := net.ResolveIPAddr("ip4", "192.168.0.21")
	fmt.Println(dstIP)
	_, srcPort := getLocalIP(dstIP.IP)

	fmt.Println(srcPort)
	return
	// 链路层
	eth := layers.Ethernet{
		SrcMAC:       srcMac,  //发送端的 mac
		DstMAC:       distMac, //发送端的 mac
		EthernetType: layers.EthernetTypeIPv4,
	}

	//IP层
	ipLayer := layers.IPv4{
		SrcIP:    srcIP,
		DstIP:    dstIP.IP,
		Version:  4,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}
	// 四层 tcp
	tcpLayer := layers.TCP{
		SrcPort: layers.TCPPort(srcPort),
		DstPort: layers.TCPPort(80), // 假设接收方是 80
		ACK:     true,
		PSH:     true,
		Window:  10,
	}
	type User struct {
		ID   string `json:"json_id"`
		Name string `json:"json_name"`
	}

	u := User{ID: "user001", Name: "tom"}
	jsonU, _ := json.Marshal(u)
	payload := gopacket.Payload(jsonU)

	err := tcpLayer.SetNetworkLayerForChecksum(&ipLayer)
	if err != nil {
		panic(err)
	}
	c := CustomLayer{
		SomeByte:    2,
		AnotherByte: 3,
		restOfData:  []byte{1, 2, 3, 4},
	}
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	err = gopacket.SerializeLayers(buf, opts, &eth, &ipLayer, &tcpLayer, payload, c)
	if err != nil {
		panic(err)
	}
	pd := buf.Bytes()
	fmt.Println(pd)

	fileName := "./packets.pcap"
	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("file error:", err)
		return
	}
	defer file.Close()
	//w := bufio.NewWriter(file)
	//// 1.写pacap header
	//writePcapHeader(w)
	//// 写packet data
	//packetAllDataWritePcap(w, buf.Bytes())

	//write.Write(buf.Bytes())
	//write.Flush()
	//[]gopacket.Packet

	w := pcapgo.NewWriter(file)
	w.WriteFileHeader(1024, layers.LinkTypeEthernet)
	//w.WritePacket(ethpacket.Metadata().CaptureInfo, ethpacket.Data())
	w.WritePacket(gopacket.CaptureInfo{
		Timestamp:      time.Time{},
		CaptureLength:  len(pd),
		Length:         len(pd),
		InterfaceIndex: 0,
		AncillaryData:  nil,
	}, pd)

	//handle, err := pcap.OpenLive("eth0", 2048, false, 30*time.Second)
	//if err != nil {
	//	log.Fatal("pcap打开网络设备失败:", err)
	//}
	//defer handle.Close()
	////向 我们的网络设备  发包
	//err = handle.WritePacketData(buf.Bytes())
	//if err != nil {
	//	log.Fatal("发送数据失败:", err)
	//}
	//log.Print("数据包已经发送\n")

}

// writePcapHeader 写pacap头文件
func writePcapHeader(w *bufio.Writer) {
	/*
		pcap文件头：24字节(固定)
		Magic(4Byte)：标记文件开始，并用来识别文件自己和字节顺序
		Major(2Byte)： 当前文件主要的版本号
		Minor(2Byte)： 当前文件次要的版本号
		ThisZone(4Byte)：当地的标准时间，如果用的是GMT则全零，一般都直接写 0000 0000
		SigFigs(4Byte)：时间戳的精度
		SnapLen(4Byte)：最大的存储长度
		LinkType(4Byte)：链路类型
	*/
	fileHeader := []byte("\xD4\xC3\xB2\xA1\x02\x00\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x08\x00\x00\x01\x00\x00\x00")
	w.Write(fileHeader)
}

// packetAllDataWritePcap 获取数据包数据 并写入pcap文件中
func packetAllDataWritePcap(w *bufio.Writer, respData []byte) {
	// 响应数据包头部固定长度
	inteHeaderLen := 9
	// 响应数据长度
	dataLen := len(respData)
	var size int
	size += dataLen
	// 循环判断，并根据数据包下载片段重新定义inteHeaderLen，与响应长度进行比较
	for inteHeaderLen < dataLen {
		realData := respData[inteHeaderLen:]
		// 2.写packet header
		// 响应信息中实际的数据包内容 写入实际数据包内容，该数据长度是capLen
		capLen := writePacketHeader(w, realData)
		// 3.写packet body
		// 数据长度就是capLen个Byte，在此之后是一个新的Packet Header，新的Packet Data，如此循环
		writePacketData(w, realData[17:capLen+17])
		w.Flush()
		inteHeaderLen = inteHeaderLen + 17 + int(capLen)
	}
}

// writePacketHeader 写packet头文件
func writePacketHeader(w *bufio.Writer, realData []byte) uint32 {
	/*
		存在多个packet header
		Timestamp(4Byte)：被捕获时间的高位，精度为seconds
		Timestamp(4Byte)：被捕获时间的低位，精度为microseconds
		Caplen(4Byte)：当前数据区的长度，即抓取到的数据帧长度，不包括Packet Header本身的长度，单位是 Byte ，由此可以得到下一个数据帧的位置。
		Len(4Byte)：离线数据长度：网络中实际数据帧的长度，一般不大于caplen，多数情况下和Caplen数值相等。
	*/
	var packatHeaderData []byte
	// timeNanosecond 时间戳8个字节（q）- 大端序
	timeNanosecond := binary.BigEndian.Uint64(realData[1:9])
	// pktLen 数据包长度是4字节（I） - 大端序
	pktLen := binary.BigEndian.Uint32(realData[9:23])
	// capLen 捕获长度是4字节（I），网络序 - 大端序
	capLen := binary.BigEndian.Uint32(realData[13:17])
	// 数据包头的秒时间戳
	timeSec := int(timeNanosecond / 1000000000)
	timeMsec := int64((int(timeNanosecond) - timeSec*1000000000) / 1000)
	// 获取高位时间戳
	timeMsec32 := make([]byte, 4)
	binary.LittleEndian.PutUint32(timeMsec32, uint32(timeMsec))
	packatHeaderData = append(packatHeaderData, timeMsec32...)
	packatHeaderData = append(packatHeaderData, timeMsec32...)
	// 数据包长度，表示所抓获的数据包保存在pcap文件中的实际长度
	capLen32 := make([]byte, 4)
	binary.LittleEndian.PutUint32(capLen32, uint32(capLen))
	packatHeaderData = append(packatHeaderData, capLen32...)
	// 数据包实际长度，所抓获的数据包的真实长度，如果文件中保存不是完整的数据包，则这个值可能要比前面的数据包长度值大
	pktLen32 := make([]byte, 4)
	binary.LittleEndian.PutUint32(pktLen32, uint32(pktLen))
	packatHeaderData = append(packatHeaderData, pktLen32...)
	w.Write(packatHeaderData)
	return capLen
}

// writePacketData 写packet文本信息
func writePacketData(w *bufio.Writer, packData []byte) {
	// 写packet data
	w.Write(packData)
}
