package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	smva2 "gitlab.casa-systems.com/mobility/smf/infra/api/sm/v1alpha2"
	"gitlab.casa-systems.com/mobility/smf/infra/protocol/go-diameter/message/base"
)

// type RequestType[T base.CCRequestType | base.AcctRecordType] *T
type RequestType[T base.CCRequestType | base.AcctRecordType] struct {
	Rt T
}

func (r *RequestType[T]) GetType(msgType T) (T, error) {

	return T(base.CCRequestType_INITIAL_REQUEST), nil
}
func GetRequestType[T base.CCRequestType | base.AcctRecordType](msgType T) smva2.MsgType {
	var i any = msgType
	switch i.(type) {
	case base.CCRequestType:
		switch i {
		case base.CCRequestType_INITIAL_REQUEST:
			return smva2.CCRequestType_INITIAL_REQUEST
		case base.CCRequestType_UPDATE_REQUEST:
			return smva2.CCRequestType_UPDATE_REQUEST
		case base.CCRequestType_TERMINATION_REQUEST:
			return smva2.CCRequestType_TERMINATION_REQUEST
		}
	case base.AcctRecordType:
		switch i {
		case base.AcctRecordType_START:
			return smva2.CCRequestType_TERMINATION_REQUEST
		case base.AcctRecordType_INTERIM:
			return smva2.CCRequestType_UPDATE_REQUEST
		case base.AcctRecordType_STOP:
			return smva2.CCRequestType_TERMINATION_REQUEST
		}
	}
	return ""
}
func main() {
	fmt.Println(spew.Sdump(GetRequestType(base.AcctRecordType_START)))
}
