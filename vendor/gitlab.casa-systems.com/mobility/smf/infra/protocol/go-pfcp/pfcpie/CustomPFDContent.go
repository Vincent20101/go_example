package pfcpie

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	CustomPFDContentDirectionBidirectional = 0
	CustomPFDContentDirectionUplink        = 1
	CustomPFDContentDirectionDownlink      = 2
)

const (
	CustomPFDContentParameterIpProtocol = iota + 1
	CustomPFDContentParameterIpSourceAddress
	CustomPFDContentParameterIpDestinationAddress
	CustomPFDContentParameterAppProtocol
	CustomPFDContentParameterAppName
	CustomPFDContentParameterAppGroup
	CustomPFDContentParameterAppAttribute
	CustomPFDContentParameterHttpAccept
	CustomPFDContentParameterHttpAnyMatch
	CustomPFDContentParameterHttpContent
	CustomPFDContentParameterHttpCookie
	CustomPFDContentParameterHttpDomain
	CustomPFDContentParameterHttpError
	CustomPFDContentParameterHttpHost
	CustomPFDContentParameterHttpReferer
	CustomPFDContentParameterHttpRequestMethod
	CustomPFDContentParameterHttpTransferEncoding
	CustomPFDContentParameterHttpUrl
	CustomPFDContentParameterHttpUserAgent
	CustomPFDContentParameterHttpUri
	CustomPFDContentParameterHttpVersion
	CustomPFDContentParameterHttpXheader
	CustomPFDContentParameterHttpReply
	CustomPFDContentParameterHttpFirstPacket
	CustomPFDContentParameterHttpHeaderLength
	CustomPFDContentParameterHttpPayloadLength
	CustomPFDContentParameterHttpPduLength
	CustomPFDContentParameterHttpSessionPreviousState
	CustomPFDContentParameterHttpSessionState
	CustomPFDContentParameterHttpSessionLength
	CustomPFDContentParameterHttpTransactionLength
	CustomPFDContentParameterTlsSni
	CustomPFDContentParameterQuicSni
	CustomPFDContentParameterSrcPort
	CustomPFDContentParameterDstPort
)

const (
	CustomPFDContentValueTypeUint8    = 1
	CustomPFDContentValueTypeUint16   = 2
	CustomPFDContentValueTypeuint32   = 3
	CustomPFDContentValueTypeString   = 4
	CustomPFDContentValueTypeSniPool  = 5
	CustomPFDContentValueTypeHostPool = 6
	CustomPFDContentValueTypeProtocol = 7
	CustomPFDContentValueTypeRange    = 8
)

const (
	CustomPFDContentOperationNotContains   = 1
	CustomPFDContentOperationNotEndsWith   = 2
	CustomPFDContentOperationNotStartsWith = 3
	CustomPFDContentOperationEquals        = 4
	CustomPFDContentOperationNotEquals     = 5
	CustomPFDContentOperationContains      = 6
	CustomPFDContentOperationEndsWith      = 7
	CustomPFDContentOperationRegex         = 8
	CustomPFDContentOperationStartsWith    = 9
)

const (
	CustomPFDContentProtocolNone = 255
)

// Custom PFD Content IE. Refer to [Func. Spec. SMF shall support VzW required pre-defined rules].
//
// The Volume Quota IE type shall be encoded as shown in Figure 8.2.50-1. It
// contains the volume quota to be monitored by the UP function.
//
//	                                        Bits
//	 Octets       8       7       6       5       4       3       2       1
//	          +---------------------------------------------------------------+
//	 1        |                Length of PFD ContextID = a                    |
//	          |---------------------------------------------------------------|
//	 2 to 2+a |                      PFD Context ID                           |
//	          |---------------------------------------------------------------|
//	 j        |             			PFD Context Priority 					 |
//	          |---------------------------------------------------------------|
//		k  	 	 |             		Length of PFD Content ID = b              	 |
//	          |---------------------------------------------------------------|
//		l to l+b |                     	PFD Content ID     		                 |
//	          |---------------------------------------------------------------|
//	 m  		 |                    PFD Content Priority                       |
//	          |---------------------------------------------------------------|
//	 n  		 |                    		Protocol   		                     |
//	          |---------------------------------------------------------------|
//	 o  		 |                    		Direction     		                 |
//	          |---------------------------------------------------------------|
//	 p  		 |                    		Parameter			                 |
//	          |---------------------------------------------------------------|
//	 q  		 |                    		Value Type                       	 |
//	          |---------------------------------------------------------------|
//	 r  		 |                    		Value Len = y      		             |
//	          |---------------------------------------------------------------|
//	 s to s+y |                    		Value                  			     |
//	          |---------------------------------------------------------------|
//	 t  		 |                    		Operation         		             |
//	          |---------------------------------------------------------------|
// e to e+1   |                       PriorityValueListLen                    |
//	          |---------------------------------------------------------------|
// w + w+f    |                    		PriorityValueList   		          |
//	          |---------------------------------------------------------------|

// IE Type: 32790 (0x8016)
// CASA ENTERPRISEID = 20858
type CustomPFDContent struct {
	EnterpriseId    uint16 `json:"enterpriseId,omitempty"`
	ContextID       string `json:"contextID,omitempty"`       // PFD Context ID - maps to pfd group in UPF
	ContextPriority uint8  `json:"contextPriority,omitempty"` // PFD Context priority
	ID              string `json:"iD,omitempty"`              // PFD Content ID
	Priority        uint8  `json:"priority,omitempty"`        // PFD Content priority to decide AND or OR. Same priority means AND(Current UPF logic)
	// Default 255.(if Protocol value is 0, which means:'IPv6 Hop-by-Hop Option',So use 255(Reserved) to indicate that the Protocol field is not set)
	// 6(TCP),17(UDP) etc; (NOTE: 255 indicate Reserved); Protocol Numbers refer to: https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	Protocol          uint8                   `json:"protocol,omitempty"`
	Direction         uint8                   `json:"direction,omitempty"`         // bidirectional by default
	Parameter         uint8                   `json:"parameter,omitempty"`         // port, host-name, content, cookie, SNIpool, hostpool etc
	ValueType         uint8                   `json:"valueType,omitempty"`         // uint8, uint16, uint32, string
	Value             []byte                  `json:"value,omitempty"`             // parameter value
	Operation         uint8                   `json:"operation,omitempty"`         // parameter operation like starts/end with, contains, equals etc
	PriorityValueList []PriorityValueListType `json:"priorityValueList,omitempty"` //value list of content,each element contains priority and value.
}

type PriorityValueListType struct {
	Priority uint8  `json:"priority,omitempty"` // PFD Content priority to decide AND or OR. Same priority means AND(Current UPF logic)
	Value    []byte `json:"value,omitempty"`    // parameter value
}

func NewCustomPFDContent() *CustomPFDContent {
	return &CustomPFDContent{
		EnterpriseId: CasaEnterpriseID,
		// Default 255.(if Protocol value is 0, which means:'IPv6 Hop-by-Hop Option',So use 255(Reserved) to indicate that the Protocol field is not set)
		Protocol: CustomPFDContentProtocolNone,
	}
}

func (h *CustomPFDContent) MarshalBinary() (data []byte, err error) {
	if len(h.ContextID) == 0 {
		return nil, fmt.Errorf("PFD ContextId marshal Failed: %v", h.ContextID)
	}
	if len(h.ID) == 0 {
		return nil, fmt.Errorf("PFD ID marshal Failed: %v", h.ID)
	}
	var b bytes.Buffer
	b.WriteByte(byte(len(h.ContextID)))
	b.Write([]byte(h.ContextID))
	b.WriteByte(byte(h.ContextPriority))
	b.WriteByte(byte(len(h.ID)))
	b.Write([]byte(h.ID))
	b.WriteByte(byte(h.Priority))
	b.WriteByte(byte(h.Protocol))
	b.WriteByte(byte(h.Direction))
	b.WriteByte(byte(h.Parameter))
	b.WriteByte(byte(h.ValueType))
	valueLen := []byte{0x00, 0x00}
	vlen := uint16(len(h.Value))
	valueLen[0] = byte(vlen >> 8)
	valueLen[1] = byte(vlen)
	b.Write(valueLen)
	if vlen > 0 {
		b.Write(h.Value)
	}
	b.WriteByte(byte(h.Operation))

	//priorifyValueList
	listLen := []byte{0x00, 0x00}
	listbuf := []byte{}
	for _, v := range h.PriorityValueList { //Each element
		data := []byte{0x00, 0x00, 0x00, 0x00, 0x00}
		vlen = uint16(1 /*len of priority*/ + 2 /*Length of Value*/ + len(v.Value) /*vlaue buf*/)
		data[0] = byte(vlen >> 8) //Length of Element
		data[1] = byte(vlen)
		data[2] = byte(v.Priority)
		lenValueBuf := uint16(len(v.Value))
		data[3] = byte(lenValueBuf >> 8)
		data[4] = byte(lenValueBuf)
		listbuf = append(listbuf, data...)
		listbuf = append(listbuf, v.Value...)
	}
	vlen = uint16(len(listbuf))
	listLen[0] = byte(vlen >> 8)
	listLen[1] = byte(vlen)
	b.Write(listLen)
	if vlen > 0 {
		b.Write(listbuf)
	}
	return b.Bytes(), nil
}

func (h *CustomPFDContent) UnmarshalBinary(data []byte) error {
	length := len(data)
	if length == 0 || data[0] == 0 {
		return fmt.Errorf("ie CustomPFDContents Inadequate length: %d", length)
	}

	b := bytes.NewBuffer(data)

	ln, err := b.ReadByte()
	if err != nil {
		return fmt.Errorf("PFD ContextID Inadequate length: %v", err)
	}

	h.ContextID = string(b.Next(int(ln)))
	if len(h.ContextID) != int(ln) {
		return fmt.Errorf("PFD ContextID Inadequate length: %v", err)
	}

	h.ContextPriority, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("PFD Context Priority Inadequate length: %v", err)
	}

	ln, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("PFD ContentID Inadequate length: %v", err)
	}

	h.ID = string(b.Next(int(ln)))
	if len(h.ID) != int(ln) {
		return fmt.Errorf("PFD ContentID Inadequate length: %v", err)
	}

	h.Priority, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("Priority Inadequate length: %v", err)
	}

	h.Protocol, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("Protocol Inadequate length: %v", err)
	}

	h.Direction, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("Direction Inadequate length: %v", err)
	}

	h.Parameter, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("Parameter Inadequate length: %v", err)
	}

	h.ValueType, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("ValueType Inadequate length: %v", err)
	}

	rbs := b.Next(2)
	valueLen := binary.BigEndian.Uint16(rbs)
	if valueLen != 0 {
		h.Value = b.Next(int(valueLen))
	}

	h.Operation, err = b.ReadByte()
	if err != nil {
		return fmt.Errorf("Operation Inadequate length: %v", err)
	}

	if b.Len() > 0 {
		blen := b.Next(2)
		totallen := binary.BigEndian.Uint16(blen)
		for totallen > 0 {
			val := PriorityValueListType{}
			blen = b.Next(2)
			if len(blen) == 0 {
				return fmt.Errorf("blen is zero")
			}
			vlen := binary.BigEndian.Uint16(blen)
			val.Priority, err = b.ReadByte()
			if err != nil {
				return fmt.Errorf("Priority Inadequate length: %v", err)
			}
			if (vlen - 1) > totallen {
				return fmt.Errorf("vlen more than total len")
			}
			blen = b.Next(2) //Length of Value buffer
			if len(blen) == 0 {
				return fmt.Errorf("blen is zero")
			}
			lenValueBuf := binary.BigEndian.Uint16(blen)
			if (lenValueBuf) > totallen {
				return fmt.Errorf("vlen more than total len")
			}
			val.Value = b.Next(int(lenValueBuf))

			h.PriorityValueList = append(h.PriorityValueList, val)
			totallen = totallen - (vlen + 2)
		}
	}
	return nil
}
