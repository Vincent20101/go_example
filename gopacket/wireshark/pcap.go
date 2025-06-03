package main

import (
	"encoding/hex"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

func main() {
	// 原始十六进制数据（已去除空格和注释）
	rawHex := "f2fa66cb10ec52215939c4f70800450000782bd0400040069e34ac1c8c01ac1c8c4182010f1d6739ec3d6be2e3ad801800ed70e600000101080a629630518fc6765e0100004480000118000000004797f25988a6c1f700000108400000146c61622e636173612e636f6d0000012840000010636173612e636f6d000001164000000c67f87b18"

	// 1. 创建PCAP文件
	f, _ := os.Create("output.pcap")
	defer f.Close()

	// 2. 初始化PCAP写入器
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(65536, layers.LinkTypeEthernet)

	// 3. 解析原始数据
	rawBytes, _ := hex.DecodeString(rawHex)

	// 4. 构造数据包
	packet := gopacket.NewPacket(
		rawBytes,
		layers.LayerTypeEthernet,
		gopacket.Default,
	)

	// 5. 写入数据包
	ci := gopacket.CaptureInfo{
		Timestamp:     time.Now(),
		Length:        len(rawBytes),
		CaptureLength: len(rawBytes),
	}
	w.WritePacket(ci, packet.Data())
}
