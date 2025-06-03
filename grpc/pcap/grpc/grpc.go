package grpc

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"gitlab.casa-systems.com/mobility/smf/infra/base/logger"

	"golang.org/x/net/http2/hpack"

	"gitlab.casa-systems.com/mobility/smf/infra/base/pcap/packet"

	"github.com/google/gopacket/layers"
	"golang.org/x/net/http2"
)

func NewGrpcMsg(method string, body []byte) []byte {

	var result []byte
	var buf3 bytes.Buffer
	e := hpack.NewEncoder(&buf3)
	e.WriteField(hpack.HeaderField{
		Name:      ":path",
		Value:     method,
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

	btest := []byte{0}
	bi := make([]byte, 4)
	binary.BigEndian.PutUint32(bi, uint32(len(body)))

	btest = append(btest, bi...)
	btest = append(btest, body...)

	buf4 := new(bytes.Buffer)
	fr4 := http2.NewFramer(buf4, buf4)
	fr4.WriteData(2315461, true, btest)
	fmt.Println(buf4.Bytes())
	result = append(result, buf4.Bytes()...)
	return result
}

func NewGrpcPacket(lc *logger.LogContext, data []byte, method string, opts ...packet.Option) ([]byte, error) {
	lc.LogDbg("====NewGrpcPacket")
	opts = append(opts, packet.WithL4Protocol(layers.IPProtocolTCP))
	pl := packet.NewLayer(opts...)
	var p, err = pl.GeneratePacket(NewGrpcMsg(method, data))
	if err != nil {
		lc.LogError("generate pcap pcaket failed: %v", err)
		return nil, err
	}
	return p, nil
}
