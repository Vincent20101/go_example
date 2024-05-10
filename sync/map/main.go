package main

import (
	"fmt"
	"sync"
)

var pair sync.Map

type ReqRspPair struct {
	Req        int
	Rsp        int
	lastStepNo uint16
}

func main() {
	pair = sync.Map{}

	usReqRspPairLoad := &ReqRspPair{
		lastStepNo: 1,
	}
	pair.Store(usReqRspPairLoad.lastStepNo, usReqRspPairLoad)
	value, ok := pair.Load(usReqRspPairLoad.lastStepNo)
	if !ok {
	}
	rr, _ := value.(*ReqRspPair)
	var data *int
	data = &rr.Rsp

	*data = 3

	fmt.Println(data)
	fmt.Println(rr.Rsp)

	*data = 30

	fmt.Println(data)
	fmt.Println(rr.Rsp)
}
