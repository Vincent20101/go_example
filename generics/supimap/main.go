package main

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

type PcapCtxOption interface {
	pcapCtxOption()
}

type SupiMapOption struct {
	IsOnlyForSupi bool // Whether to capture packets based on supi only
	IsUpdate      bool // Whether to update supiMap
}

func (*SupiMapOption) pcapCtxOption() {}

type SubscriberId struct {
	Imsi   string
	Imei   string
	Msisdn string
	UeipV4 string
	UeipV6 string
	// "ipv4" or "ipv6"
	TraceUeipType string
}

type SubscriberOption func(id *SubscriberId)

func WithImsi(v string) SubscriberOption {
	return func(s *SubscriberId) {
		s.Imsi = v
	}
}

func WithImei(v string) SubscriberOption {
	return func(s *SubscriberId) {
		s.Imei = v
	}
}

func WithMsisdn(v string) SubscriberOption {
	return func(s *SubscriberId) {
		s.Msisdn = v
	}
}

func WithUeip(v4, v6 string) SubscriberOption {
	return func(s *SubscriberId) {
		s.UeipV4 = v4
		s.UeipV6 = v6
		s.TraceUeipType = "ipv4"
		if v6 != "" {
			s.TraceUeipType = "ipv6"
		}
	}
}

func (m *TrcMgr[T]) UpdateSupiMapByIdOption(supi string, opts ...SubscriberOption) {
	if supi == "" {
		return
	}

	if m.supiMap == nil {
		m.supiMap = make(map[string]*SubscriberId)
	}
	if v, exist := m.supiMap[supi]; !exist || v == nil {
		m.supiMap[supi] = &SubscriberId{Imsi: supi}
	}

	for _, opt := range opts {
		opt(m.supiMap[supi])
	}
}

type TrcMgr[T PcapCtxOption] struct {
	supiMap map[string]*SubscriberId
}

var trcMgr *TrcMgr[PcapCtxOption]

func GetTrcMgr() *TrcMgr[PcapCtxOption] {
	return trcMgr
}

func init() {
	trcMgr = new(TrcMgr[PcapCtxOption])
}

func (p *TrcMgr[T]) updateSubscriberIdMap(args ...T) {
	supiMapOption := new(SupiMapOption)
	for _, arg := range args {
		switch v := any(arg).(type) {
		case *SupiMapOption:
			if v != nil {
				supiMapOption = v
			}
		default:
		}
	}
	fmt.Println(supiMapOption.IsUpdate)
	fmt.Println(supiMapOption.IsOnlyForSupi)
}

func (p *TrcMgr[T]) ReadTraceConfig() {
	if opt, ok := any(&SupiMapOption{IsUpdate: true}).(T); ok {
		p.updateSubscriberIdMap(opt)
	}
}

type SmContextCreateData struct {
	isSm bool
}

type PduSessionCreateData struct {
	isPdu bool
}

type CreateSessionReq struct {
	isCre bool
}

type CreatePdpContextReq struct {
	isctx bool
}

type SessionReq interface {
	*SmContextCreateData | *PduSessionCreateData | *CreateSessionReq | *CreatePdpContextReq | struct{}
}

type HandleOptionFunc[T SessionReq] func(opt *HandleOption[T])

type HandleOption[T SessionReq] struct {
	supi                string // supi
	isSessIdFromReqPath bool   // Whether session id from req path
	isCreateSess        bool   // Whether sess created
	httpRequest         *http.Request
	createSessReq       T
}

func WithSupi[T SessionReq](supi string) HandleOptionFunc[T] {
	return func(opt *HandleOption[T]) {
		opt.supi = supi
	}
}

func WithSessIdFromReqPath[T SessionReq]() HandleOptionFunc[T] {
	return func(opt *HandleOption[T]) {
		opt.isSessIdFromReqPath = true
	}
}

func WithCreateSess[T SessionReq]() HandleOptionFunc[T] {
	return func(opt *HandleOption[T]) {
		opt.isCreateSess = true
	}
}

func WithHttpRequest[T SessionReq](req *http.Request) HandleOptionFunc[T] {
	return func(opt *HandleOption[T]) {
		opt.httpRequest = req
	}
}

func WithCreateSessReq[T SessionReq](createSessReq T) HandleOptionFunc[T] {
	return func(opt *HandleOption[T]) {
		opt.createSessReq = createSessReq
	}
}

func HandlePcapTraceForInboundReq[T SessionReq](opts ...HandleOptionFunc[T]) {
	ho := new(HandleOption[T])
	for _, opt := range opts {
		opt(ho)
	}
	fmt.Println(spew.Sdump(ho))
}

func main() {
	GetTrcMgr().UpdateSupiMapByIdOption("123", WithImei("456"), WithMsisdn("789"))
	GetTrcMgr().UpdateSupiMapByIdOption("abc", WithImei("edc"), WithMsisdn("tghfg"))
	GetTrcMgr().ReadTraceConfig()
	fmt.Println(spew.Sdump(GetTrcMgr()))
	//HandlePcapTraceForInboundReq[struct{}](WithCreateSess[struct{}](), WithSessIdFromReqPath[struct{}]())
	//
	//st := &SmContextCreateData{
	//	isSm: true,
	//}
	//HandlePcapTraceForInboundReq[*SmContextCreateData](WithCreateSess[*SmContextCreateData](), WithSessIdFromReqPath[*SmContextCreateData](), WithCreateSessReq[*SmContextCreateData](st))

}
