package udppcaket

import (
	"encoding/binary"
	"fmt"
)

type Header struct {
	MessageType   messageType `json:"messageType"`
	MessageLength uint16      `json:"messageLength"`
}

func (h *Header) Marshal() (data []byte, err error) {
	//MessageType
	data = append([]byte(""), byte(h.MessageType))
	//MessageLength
	temp := make([]byte, 2)
	binary.BigEndian.PutUint16(temp, h.MessageLength)
	data = append(data, temp...)
	return
}

func (h *Header) Unmarshal(data []byte) (err error) {
	length := uint16(len(data))
	var idx uint16 = 0
	//MessageType
	if length < idx+1 {
		err = fmt.Errorf("Unmarshal DataPacket.MessageType  Message unqualified length: %d", length)
		return
	}
	h.MessageType = messageType(data[idx])
	idx += 1
	if length < idx+2 {
		err = fmt.Errorf("Unmarshal DataPacket.MessageLength  Message unqualified length: %d,idx: %d", length, idx)
		return
	}
	h.MessageLength = binary.BigEndian.Uint16(data[idx : idx+2])
	return
}
