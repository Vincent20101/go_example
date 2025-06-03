package pcap_http2

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"strings"

	slog "gitlab.casa-systems.com/mobility/smf/infra/base/logger"
	com "gitlab.casa-systems.com/mobility/smf/infra/base/pcap/common"
	"gitlab.casa-systems.com/mobility/smf/infra/base/pcap/packet"

	"github.com/google/gopacket/layers"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type Packet struct {
	StreamId            uint32
	Header              map[string][]string
	Body                []byte
	PcapData            []byte
	NanoSecondTimestamp func() uint64
}

func NewPacket(header map[string][]string, body []byte, t func() uint64) *Packet {
	n, _ := rand.Int(rand.Reader, big.NewInt(1<<31))
	hp := &Packet{
		StreamId:            uint32(n.Int64()),
		Header:              header,
		Body:                body,
		PcapData:            make([]byte, 0),
		NanoSecondTimestamp: t,
	}
	return hp
}

// Generate http2 stream header
func (hp *Packet) streamHeader() ([]byte, error) {

	var headerFieldBuf bytes.Buffer
	e := hpack.NewEncoder(&headerFieldBuf)
	if len(hp.Header) > 0 {
		for k, v := range hp.Header {
			if err := e.WriteField(hpack.HeaderField{Name: strings.ToLower(k), Value: v[0]}); err != nil {
				return nil, err
			}
		}
	}

	headerBuf := new(bytes.Buffer)
	fr := http2.NewFramer(headerBuf, headerBuf)
	header := http2.HeadersFrameParam{
		StreamID:      hp.StreamId,
		BlockFragment: headerFieldBuf.Bytes(),
		EndStream:     false,
		EndHeaders:    true,
		PadLength:     0,
		Priority:      http2.PriorityParam{},
	}
	err := fr.WriteHeaders(header)

	return headerBuf.Bytes(), err
}

// Generate http2 stream data
func (hp *Packet) streamData() ([]byte, error) {

	buf := new(bytes.Buffer)
	fr := http2.NewFramer(buf, buf)
	err := fr.WriteData(hp.StreamId, true, hp.Body)
	return buf.Bytes(), err
}

// Generate http2 stream header and stream data
func (hp *Packet) streamDataWithHeader() ([]byte, error) {

	var hb, sh, sd []byte
	var err error

	if sh, err = hp.streamHeader(); err != nil {
		return nil, err
	}
	hb = append(hb, sh...)

	if hp.Body != nil && len(hp.Body) != 0 {
		if sd, err = hp.streamData(); err != nil {
			return nil, err
		}
	}

	hb = append(hb, sd...)
	return hb, nil
}

// AssemblePcapWithHeader assemble pcap packet with stream header
func (hp *Packet) AssemblePcapWithHeader(opts ...packet.Option) *Packet {
	data, err := hp.streamHeader()
	if err != nil {
		return hp
	}
	hp.PcapData = hp.PcapData[:0]
	opts = append(opts, packet.WithL4Protocol(layers.IPProtocolTCP))
	pl := packet.NewLayer(opts...)
	hp.PcapData, _ = pl.GeneratePacket(data)
	return hp
}

// AssemblePcapWithData assemble pcap packet with stream data
func (hp *Packet) AssemblePcapWithData(opts ...packet.Option) *Packet {
	data, err := hp.streamData()
	if err != nil {
		return hp
	}
	hp.PcapData = hp.PcapData[:0]
	opts = append(opts, packet.WithL4Protocol(layers.IPProtocolTCP))
	pl := packet.NewLayer(opts...)
	hp.PcapData, _ = pl.GeneratePacket(data)
	return hp
}

// AssemblePcap assemble pcap packet with http2 header and http2 body
func (hp *Packet) AssemblePcap(opts ...packet.Option) *Packet {
	data, err := hp.streamDataWithHeader()
	if err != nil {
		return hp
	}
	hp.PcapData = hp.PcapData[:0]
	opts = append(opts, packet.WithL4Protocol(layers.IPProtocolTCP))
	pl := packet.NewLayer(opts...)
	hp.PcapData, _ = pl.GeneratePacket(data)
	return hp
}

// WritePcapToSidecar write the pcap packet to sidecar
func (hp *Packet) WritePcapToSidecar(lc *slog.LogContext) {

	if len(hp.PcapData) == 0 {
		lc.LogDbg("no pcap data available")
		return
	}

	if tc := lc.GetTracePcapContext(); tc != nil {
		com.WritePcapToSidecar(lc, tc.SubscriberId.Imsi, tc.SubscriberId.Imei, tc.SubscriberId.Msisdn, hp.NanoSecondTimestamp(), hp.PcapData)
	}
}
