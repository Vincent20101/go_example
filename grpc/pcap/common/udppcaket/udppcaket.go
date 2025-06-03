package udppcaket

import (
	"errors"
	"fmt"
)

type messageType uint8

const (
	DefaultMessageType messageType = iota + 1
)

type UDPDataPacket interface {
	MarshalBinary(msg *Message) (data []byte, err error)
	UnmarshalBinary(data []byte) (msg *Message, err error)
}

type DataPacket struct {
	Header  Header
	Message []byte `json:"message"`
}

func NewDataPacket(msgType messageType) *DataPacket {
	hdr := Header{MessageType: msgType}
	return &DataPacket{Header: hdr}
}

func (dp *DataPacket) MarshalBinary(msg *Message) (data []byte, err error) {
	if dp == nil {
		err = errors.New("DataPacket must not be nil")
		return
	}

	if dp.Header.MessageType == DefaultMessageType {
		if msg != nil {
			//MessageType
			msgBuf, err := msg.Marshal()
			if err != nil {
				return []byte(""), err
			}
			dp.Header.MessageLength = uint16(len(msgBuf))
			hdrBuf, err := dp.Header.Marshal()
			if err != nil {
				return []byte(""), err
			}
			data = append(data, hdrBuf...)
			data = append(data, msgBuf...)
		}

	}
	return
}
func (dp *DataPacket) UnmarshalBinary(data []byte) (msg *Message, err error) {
	if dp == nil {
		err = errors.New("DataPacket must not be nil")
		return
	}
	err = dp.Header.Unmarshal(data)
	if err != nil {
		err = fmt.Errorf("DataPacket header unmarshal failed: %v", err)
		return
	}
	if int(dp.Header.MessageLength) != len(data)-3 {
		err = fmt.Errorf("incorrect Message Length: Expected %d, got %d", dp.Header.MessageLength, len(data)-3)
		return
	}
	if dp.Header.MessageType == DefaultMessageType {
		msg = &Message{}
		err = msg.Unmarshal(data[3:])
	}

	return
}
