package main

// Use tcpdump to create a test file
// tcpdump -w test.pcap
// or use the example above for writing pcap files

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	pcapFile string = "tests.pcap"
	handle   *pcap.Handle
	err      error
)

func main() {
	// Open file instead of device
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Set filter
	var filter string = "tcp"
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Do something with a packet here.
		//println(packet)
		printPacketInfo(packet)
	}
}

func printPacketInfo(packet gopacket.Packet) {
	// Let's see if the packet is an ethernet packet
	ethernetLayer := packet.Layer(layers.LayerTypeEthernet)
	if ethernetLayer != nil {
		ethernetPacket, _ := ethernetLayer.(*layers.Ethernet)

		//fmt.Println(reflect.TypeOf(ethernetPacket.EthernetType))
		if ethernetPacket.EthernetType.String() == "IPv4" {
			ipLayer := packet.Layer(layers.LayerTypeIPv4)
			if ipLayer != nil {
				ip, _ := ipLayer.(*layers.IPv4)

				// IP layer variables:
				// Version (Either 4 or 6)
				// IHL (IP Header Length in 32-bit words)
				// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
				// Checksum, SrcIP, DstIP

				tcpLayer := packet.Layer(layers.LayerTypeTCP)
				if tcpLayer != nil {
					tcp, _ := tcpLayer.(*layers.TCP)

					// TCP layer variables:
					// SrcPort, DstPort, Seq, Ack, DataOffset, Window, Checksum, Urgent
					// Bool flags: FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS

					applicationLayer := packet.ApplicationLayer()

					if applicationLayer != nil {
						payload := string(applicationLayer.Payload())

						if strings.HasPrefix(payload, "GET") || strings.HasPrefix(payload, "POST") {
							fmt.Println("--------------------------------------------------------------------")
							fmt.Println("Source MAC: ", ethernetPacket.SrcMAC)
							fmt.Println("Destination MAC: ", ethernetPacket.DstMAC)
							// Ethernet type is typically IPv4 but could be ARP or other
							fmt.Println("Ethernet type: ", ethernetPacket.EthernetType)
							fmt.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
							fmt.Println("Protocol: ", ip.Protocol)
							fmt.Printf("From port %d to %d\n", tcp.SrcPort, tcp.DstPort)
							fmt.Println("Sequence number: ", tcp.Seq)

							//fmt.Printf("%s\n", applicationLayer.Payload())
							//a := packet.TransportLayer().TransportFlow()
							//fmt.Println(a.EndpointType())
							//fmt.Println("HTTP found!")
							//fmt.Println(payload)
							//解释正则表达式
							reg := regexp.MustCompile(`(?s)(GET|POST) (.*?) HTTP.*Host: (.*?)\n`)
							if reg == nil {
								fmt.Println("MustCompile err")
								return
							}
							//提取关键信息
							result := reg.FindStringSubmatch(payload)
							//fmt.Println(result)
							//fmt.Println(len(result))
							//fmt.Println(result[2], result[3])
							if len(result) == 4 {
								strings.TrimSpace(result[2])
								url := "http://" + strings.TrimSpace(result[3]) + strings.TrimSpace(result[2])
								fmt.Println("url:", url)
								fmt.Println("host:", result[3])
							} else {
								fmt.Println("error===================")
							}

						}

					}
				}
			}
		}
	}

}
