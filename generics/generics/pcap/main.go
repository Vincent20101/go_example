package main

import "fmt"

type PcapCtxOption interface {
	pcapCtxOption()
}

type SupiMapOption struct {
	IsOnlyForSupi bool
	IsUpdate      bool
}

// pointer will fail
func (*SupiMapOption) pcapCtxOption() {}

//func (SupiMapOption) pcapCtxOption() {}

type pcapTrcMgr[T PcapCtxOption] struct{}

func (p *pcapTrcMgr[T]) updateSubscriberIdMap(args ...T) {
	for _, arg := range args {
		fmt.Printf("Processing option: %T\n", arg)
		var supiMapOption SupiMapOption
		switch v := any(arg).(type) {
		case SupiMapOption:
			supiMapOption = v
		default:
			fmt.Println("Unknown option")
		}
		fmt.Println(supiMapOption.IsOnlyForSupi)
		fmt.Println(supiMapOption.IsUpdate)
	}
}

func (p *pcapTrcMgr[T]) ReadTraceConfig() {
	supiOpt := &SupiMapOption{IsUpdate: true, IsOnlyForSupi: true}
	if opt, ok := any(supiOpt).(T); ok {
		p.updateSubscriberIdMap(opt)
	} else {
		fmt.Println("Type assertion failed")
	}
}

func main() {
	TrcMgr := new(pcapTrcMgr[PcapCtxOption])
	TrcMgr.ReadTraceConfig()
}

//type PcapCtxOption interface {
//	pcapCtxOption()
//	~SupiMapOption | ~ImeiMapOption // 联合类型约束
//}

//type PcapCtxOption interface {
//	SupiMapOption | ImeiMapOption // 约束为具体类型
//	pcapCtxOption()
//}
//// 使用类型集约束
//type PcapCtxOption interface {
//	~struct { // 约束结构体类型
//		IsOnlyForSupi bool
//		IsUpdate      bool
//	} | ~struct {
//		IsOnlyForImei bool
//		IsDelete      bool
//	}
//	pcapCtxOption()
//}
