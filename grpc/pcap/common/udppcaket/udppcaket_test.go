package udppcaket

import (
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestDecodeAndEncodeMessages(t *testing.T) {

	msg := &Message{
		NanosecondTimestamp: uint64(time.Now().Nanosecond()),
		IMSI:                "imsi-450051230000000",
		IMEI:                "imei-450051230000001",
		MSISDN:              "msisdn-450051230000000",
		UEIPv4:              net.ParseIP("192.168.1.101").To16(),
		UEIPv6:              net.ParseIP("fe80::8cd5:61ff:fe0f:e917").To16(),
		PcapPacket:          []byte("hello world"),
	}
	dp := NewDataPacket(DefaultMessageType)
	data, err := dp.MarshalBinary(msg)
	if err != nil {
		t.Log("Test failed while DecodeEncodeDecodeMessage,err:", err)
		t.FailNow()
	}

	dp1 := &DataPacket{}
	msg1, err := dp1.UnmarshalBinary(data)
	if err != nil {
		t.Log("Test msg1.Unmarshal failed while DecodeEncodeDecodeMessage,err:", err)
		t.FailNow()
	}
	if !reflect.DeepEqual(msg, msg1) {
		t.Log(cmp.Diff(msg, msg1))
		t.FailNow()
	}

}
