package pfcpie

import (
	"fmt"
)

// CasaTraceInformation
// Casa Specific IE
//
// 		Octets 		    	8       7       6       5       4       3       2       1
//           			+---------------------------------------------------------------+
//  	1 to 2   		|                    Type = 32796(decimal)                   	|
//           			|---------------------------------------------------------------|
//  	3 to 4   		|                        Length = n                          	|
//           			|---------------------------------------------------------------|
//  	5 to 6   		|          			Enterprise ID = 20858             			|
//           			|---------------------------------------------------------------|
//    	   7      		|                 		Ue Id Length = m                   		|
//           			|---------------------------------------------------------------|
//     8 to 8+m 		|                    Trace Ue Id(string)                   	 	|
//           			|---------------------------------------------------------------|
//        9+m    		|                 	 Parent Span Type(int)                 	 	|
//           			|---------------------------------------------------------------|
//       10+m 		    |   	             Span Context Length = k              	 	|
//           			|---------------------------------------------------------------|
//  (11+m) to (11+m+k)  |   				Span Context(binary data)  					|
//           			+---------------------------------------------------------------+
//
// 	Trace Ue Id:
// 		used to show trace ue id information in SGW-U/UPF Jaeger UI
// 	Parent Span Type:
//		used to distinguish the  pfcp message sources for SGW-U/UPF
//		value: 0 or 1
//			0：the parent span context information is from SGW-C
//			1：the parent span context information is from PGW-C
//	Span Context Length:
//		the length of parent span context data, max value is 255
//	Span Context:
//		parent span context used to create child span for SGW-U/UPF
//		Using binary data according to the span context propagation format
//

const CasaTraceInformationTag uint16 = 32796
const ParentSpanType_SGW_C int = 0
const ParentSpanType_PGW_C int = 1

type CasaTraceInformation struct {
	EnterpriseId   uint16 `json:"enterpriseId,omitempty"`
	TraceUeId      string `json:"traceUeId,omitempty"`
	ParentSpanType int    `json:"parentSpanType,omitempty"`
	SpanContext    []byte `json:"spanContext,omitempty"`
}

func NewCasaTraceInformation() *CasaTraceInformation {
	return &CasaTraceInformation{
		EnterpriseId: CasaEnterpriseID,
	}
}

// MarshalBinary encodes JaegerForTracing.
func (e *CasaTraceInformation) MarshalBinary() (data []byte, err error) {
	// oct 5 to 6
	data = AppendUint16(data, e.EnterpriseId)
	// oct 7
	data = append(data, byte(len(e.TraceUeId)))
	// oct 8 to 8+m
	data = append(data, []byte(e.TraceUeId)...)
	// oct 9+m
	data = append(data, byte(e.ParentSpanType))
	// oct 10+m
	data = append(data, byte(len(e.SpanContext)))
	// oct (11+m) to (11+m+k)
	if len(e.SpanContext) != 0 {
		spanContext := make([]byte, len(e.SpanContext))
		copy(spanContext, e.SpanContext)
		data = append(data, spanContext...)
	}

	return data, nil
}

// UnmarshalBinary decode JaegerForTracing from binary
func (e *CasaTraceInformation) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// EnterpriseId
	enterpriseId, err := byteReader.ReadUint16()
	if err != nil {
		return fmt.Errorf("Read EnterpriseId failed, err: %v", e)
	}
	e.EnterpriseId = enterpriseId

	// Length of TraceUeId
	traceUeIdLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Read Length of TraceUeId failed, err: %v", e)
	}
	// TraceUeId
	traceUeId, err := byteReader.ReadBytes(int(traceUeIdLen))
	if err != nil {
		return fmt.Errorf("Read TraceUeId failed, err: %v", e)
	}
	e.TraceUeId = string(traceUeId)

	// ParentSpanType
	parentSpanType, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Read parentSpanType failed, err: %v", e)
	}
	e.ParentSpanType = int(parentSpanType)

	// Length of MessageCtx
	spanContextLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Read Length of MessageCtx failed, err: %v", e)
	}
	// MessageCtx
	spanContext, err := byteReader.ReadBytes(int(spanContextLen))
	if err != nil {
		return fmt.Errorf("Read SpanContextLen failed, err: %v", e)
	}
	e.SpanContext = spanContext
	return nil
}
