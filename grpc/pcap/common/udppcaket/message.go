package udppcaket

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

const (
	BitMask8 = 1 << 7
	BitMask7 = 1 << 6
	BitMask6 = 1 << 5
	BitMask5 = 1 << 4
	BitMask4 = 1 << 3
	BitMask3 = 1 << 2
	BitMask2 = 1 << 1
	BitMask1 = 1
)

type Message struct {
	NanosecondTimestamp uint64 `json:"nanosecondTimestamp"`
	IMSI                string `json:"imsi,omitempty"`
	IMEI                string `json:"imei,omitempty"`
	MSISDN              string `json:"msisdn,omitempty"`
	UEIPv4              net.IP `json:"ueIPv4,omitempty"`
	UEIPv6              net.IP `json:"ueIPv6,omitempty"`
	PcapPacket          []byte `json:"pcapPacket"`
}

func (m *Message) SetIMSI(imsi string) {
	m.IMSI = imsi
}
func (m *Message) SetIMEI(imei string) {
	m.IMEI = imei
}
func (m *Message) SetMSISDN(msisdn string) {
	m.MSISDN = msisdn
}

func (m *Message) SetUEIPv4(ip net.IP) {
	m.UEIPv4 = ip
}
func (m *Message) SetUEIPv6(ip net.IP) {
	m.UEIPv6 = ip
}

func (m *Message) SetPcapPacket(pcapPacket []byte) {
	m.PcapPacket = pcapPacket
}

func (m *Message) GetIMSI() string {
	return m.IMSI
}

func (m *Message) GetIMEI() string {
	return m.IMEI
}

func (m *Message) GetMSISDN() string {
	return m.MSISDN
}

func (m *Message) GetUEIPv4() net.IP {
	return m.UEIPv4
}
func (m *Message) GetUEIPv6() net.IP {
	return m.UEIPv6
}

func (m *Message) GetPcapPacket() []byte {
	return m.PcapPacket
}

func (m *Message) Marshal() (data []byte, err error) {
	if m == nil {
		fmt.Println("msg must not be nil")
		return []byte(""), errors.New("msg must not be nil")
	}
	//NanosecondTimestamp
	tsTemp := make([]byte, 8)
	binary.BigEndian.PutUint64(tsTemp, m.NanosecondTimestamp)
	data = append(data, tsTemp...)
	var flag uint16 = 0
	//SubscriberIdentitiesFlag
	if m.GetIMSI() != "" {
		flag = flag | BitMask1
	}
	if m.GetIMEI() != "" {
		flag = flag | BitMask2
	}
	if m.GetMSISDN() != "" {
		flag = flag | BitMask3
	}
	if net.ParseIP(m.GetUEIPv4().String()) != nil {

		flag = flag | BitMask4
	}
	if net.ParseIP(m.GetUEIPv6().String()) != nil {

		flag = flag | BitMask5
	}
	flagTemp := make([]byte, 2)
	binary.BigEndian.PutUint16(flagTemp, flag)
	data = append(data, flagTemp...)

	if flag != 0 {
		//IMSI
		if m.GetIMSI() != "" {
			data = append(data, byte(len(m.GetIMSI())))
			data = append(data, []byte(m.GetIMSI())...)
		}
		//IMEI
		if m.GetIMEI() != "" {
			data = append(data, byte(len(m.GetIMEI())))
			data = append(data, []byte(m.GetIMEI())...)
		}
		//MSISDN
		if m.GetMSISDN() != "" {
			data = append(data, byte(len(m.GetMSISDN())))
			data = append(data, []byte(m.GetMSISDN())...)

		}
		//UEIPv4
		if net.ParseIP(m.GetUEIPv4().String()) != nil {
			data = append(data, m.GetUEIPv4().To4()...)
		}
		//UEIPv6
		if net.ParseIP(m.GetUEIPv6().String()) != nil {
			data = append(data, m.GetUEIPv6().To16()...)
		}

	}

	//BodyLength
	extTemp := make([]byte, 2)
	binary.BigEndian.PutUint16(extTemp, uint16(len(m.GetPcapPacket())))
	data = append(data, extTemp...)
	//Body
	data = append(data, m.GetPcapPacket()...)

	return data, nil
}

func (m *Message) Unmarshal(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	//NanosecondTimestamp
	if length < idx+8 {
		return fmt.Errorf("unmarshal Message.NanosecondTimestamp unqualified length: %d,idx: %d", length, idx)
	}
	m.NanosecondTimestamp = binary.BigEndian.Uint64(data[idx : idx+8])
	idx += 8
	//SubscriberIdentitiesFlag
	if length < idx+2 {
		return fmt.Errorf("unmarshal Message.SubscriberIdentitiesFlag unqualified length: %d,idx: %d", length, idx)
	}
	flag := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2
	imsiPres := (flag & BitMask1) != 0
	imeiPres := (flag & BitMask2) != 0
	msisdnPres := (flag & BitMask3) != 0
	ueIpv4 := (flag & BitMask4) != 0
	ueIpv6 := (flag & BitMask5) != 0

	if imsiPres {
		//IMSILength
		if length < idx+1 {
			return fmt.Errorf("unmarshal Message.IMSILength unqualified length: %d,idx: %d", length, idx)
		}
		imsiLength := data[idx]
		idx++

		if imsiLength == 0 {
			return errors.New("unmarshal Message.IMSI unqualified length ")
		}
		//IMSI
		if length < idx+uint16(imsiLength) {
			return fmt.Errorf("unmarshal Message.IMSI unqualified length: %d,idx: %d,IMSILength:%d", length, idx, imsiLength)
		}
		m.IMSI = string(data[idx : idx+uint16(imsiLength)])
		idx += uint16(imsiLength)
	}
	if imeiPres {
		//IMEILength
		if length < idx+1 {
			return fmt.Errorf("unmarshal Message.IMEILength unqualified length: %d,idx: %d", length, idx)
		}
		imeiLength := data[idx]
		idx++

		if imeiLength == 0 {
			return errors.New("unmarshal Message.IMEI unqualified length ")
		}
		//IMEI
		if length < idx+uint16(imeiLength) {
			return fmt.Errorf("unmarshal Message.IMEI unqualified length: %d,idx: %d,IMEILength:%d", length, idx, imeiLength)
		}
		m.IMEI = string(data[idx : idx+uint16(imeiLength)])
		idx += uint16(imeiLength)
	}
	if msisdnPres {
		//MSISDNLength
		if length < idx+1 {
			return fmt.Errorf("unmarshal Message.MSISDNLength unqualified length: %d,idx: %d", length, idx)
		}
		msisdnLength := data[idx]
		idx++

		if msisdnLength == 0 {
			return errors.New("unmarshal Message.MSISDN  unqualified length ")
		}
		//MSISDN
		if length < idx+uint16(msisdnLength) {
			return fmt.Errorf("unmarshal Message.MSISDN unqualified length: %d,idx: %d,MSISDNLength:%d", length, idx, msisdnLength)
		}
		m.MSISDN = string(data[idx : idx+uint16(msisdnLength)])
		idx += uint16(msisdnLength)
	}
	//UEIPv4
	if ueIpv4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("unmarshal Message.UEIPv4 unqualified length: %d,idx: %d", length, idx)
		}
		m.UEIPv4 = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}
	//UEIPv6
	if ueIpv6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("unmarshal Message.UEIPv6 unqualified length: %d,idx: %d", length, idx)
		}
		m.UEIPv6 = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	//PcapPacketLength
	if length < idx+2 {
		return fmt.Errorf("unmarshal Message.BodyLength unqualified length: %d,idx: %d ", length, idx)
	}
	pcapPacketLength := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2
	//PcapPacket
	if uint16(len(data[idx:])) != pcapPacketLength {
		return fmt.Errorf("unmarshal Message.PcapPacket unqualified length: %d,idx: %d,BodyLength:%d ", length, idx, pcapPacketLength)
	}
	m.SetPcapPacket(data[idx:])

	return nil
}
