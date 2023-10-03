package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

const (
	CdBitMap_IS_REQ_FAIL                     CdBitFlag = 1 << iota // req to chf fail
	CdBitMap_QBC                                                   // is qos flow base charging or flow base charging, only exist one
	CdBitMap_OCS_SESS_ON                                           // indicate if there is OCS session for online charging
	CdBitMap_DEFER_CREATE_CHARGINGCtfTrigger                       // indicate if deferred create charging

)

const (
	CdBitMap_IS_NOWAIT_INIT        CtfBitFlag = 1 << iota // use local trigger for no wait init quota.
	CdBitMap_IS_AP_STATE                                  // use local config for AP charging state
	CdBitMap_IS_USE_FAILED_CFG                            // use local config for continue error handle(not AP charging mode)
	CdBitMap_IS_NOWAIT_CASE                               // indicate need request quota for delayQuota feature in no-wait control
	CdBitMap_IS_INIT_QUOTA_TRIGGER                        // YTL: GCS-6688
	CdBitMap_IS_TERMINATE                                 // is termination
)

const (
	CdBitMap_IS_ONLINE         UrrKeyBitFlag = 1 << iota // for online or offline charging
	CdBitMap_IS_SESSION_LIMITS                           // for session-level limits
	CdBitMap_IS_BCR_URR                                  // for vzw Offline BCR support
	CdBitMap_IS_OST_URR                                  // for vzw Offline Session Timer support
	CdBitMap_IS_YTL_INIT_QUOTA_URR
)

type BitFlag[T CdBitFlag | CtfBitFlag | UrrKeyBitFlag] struct {
	Bits T
}

type CdBitFlag uint64
type CtfBitFlag uint64
type UrrKeyBitFlag uint64

func (c *BitFlag[T]) Set(bit T) {
	c.Bits |= bit
}

func (c *BitFlag[T]) Get(bit T) bool {
	return (c.Bits & bit) != 0
}

func (c *BitFlag[T]) Clean(bit T) {
	c.Bits &= ^bit
}

//
//var CdBitInstance BitFlag[CdBitFlag]
//var CtfBitInstance BitFlag[CtfBitFlag]
//var UrrKeyBitInstance BitFlag[UrrKeyBitFlag]
//
//func init() {
//	CdBitInstance = BitFlag[CdBitFlag]{Bits: 0}
//	CtfBitInstance = BitFlag[CtfBitFlag]{Bits: 0}
//	UrrKeyBitInstance = BitFlag[UrrKeyBitFlag]{Bits: 0}
//}

type ChargingData struct {
	CdBitInstance BitFlag[CdBitFlag]
}

func main() {
	var s1 = make([]string, 0)
	s1 = append(s1, "")
	s1 = append(s1, "1")
	fmt.Println(s1[:2])
	fmt.Println(len(s1))
	str := "sdfsdfsdfsdf\r\n" + "\\0\r\n" + "sdfsdfdsfsdfsdf\r\n"
	fmt.Println(str)
	var cd = ChargingData{}
	cd.CdBitInstance.Set(CdBitMap_IS_REQ_FAIL)
	cd.CdBitInstance.Set(CdBitMap_QBC)
	cd.CdBitInstance.Set(CdBitMap_DEFER_CREATE_CHARGINGCtfTrigger)
	fmt.Println("lhb:", cd.CdBitInstance)
	t := CdBitMap_QBC
	s := CdBitMap_DEFER_CREATE_CHARGINGCtfTrigger
	Toggle(t, s)
	fmt.Println("lhb:", t)
	fmt.Println(cd.CdBitInstance.Get(CdBitMap_IS_REQ_FAIL))
	cd.CdBitInstance.Clean(CdBitMap_DEFER_CREATE_CHARGINGCtfTrigger)
	fmt.Println(spew.Sdump(cd))
	fmt.Println(CdBitMap_IS_REQ_FAIL | CdBitMap_QBC)
}

func Toggle(t, s CdBitFlag) {
	t ^= s
}
