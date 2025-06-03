package grpc

import (
	"net"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDecodeAndEncodeMessagesWithIPV4(t *testing.T) {
	codec := &CodecInfo{
		SrcIp:  net.ParseIP("127.0.0.1").To16(),
		IMSI:   "imsi-450051230000000",
		IMEI:   "imei-450051230000001",
		MSISDN: "msisdn-450051230000000",
		UEIPv4: net.ParseIP("192.168.1.101").To16(),
		UEIPv6: net.ParseIP("fe80::8cd5:61ff:fe0f:e917").To16(),
		Method: "/pfcp.GopfcpSm/SessEstReq",
	}
	dataStr := "hello world"
	dataB, err := codec.encode([]byte(dataStr))
	if err != nil {
		t.Log("Test failed while encode,err:", err)
		t.FailNow()
	}
	codec2 := new(CodecInfo)
	data2, err := codec2.decode(dataB)
	if err != nil {
		t.Log("Test failed while decode,err:", err)
		t.FailNow()
	}
	if !reflect.DeepEqual(codec, codec2) {
		t.Log(cmp.Diff(codec, codec2))
		t.FailNow()
	}
	if !reflect.DeepEqual(dataStr, string(data2)) {
		t.Log(cmp.Diff(codec, codec2))
		t.FailNow()
	}

}

func TestDecodeAndEncodeMessagesWithIPV6(t *testing.T) {
	codec := &CodecInfo{
		SrcIp:  net.ParseIP("fe80::ecee:eeff:feee:eeee").To16(),
		IMSI:   "imsi-450051230000000",
		IMEI:   "imei-450051230000001",
		MSISDN: "msisdn-450051230000000",
		UEIPv4: net.ParseIP("192.168.1.101").To16(),
		UEIPv6: net.ParseIP("fe80::8cd5:61ff:fe0f:e917").To16(),
		Method: "/pfcp.GopfcpSm/SessEstReq",
	}
	dataStr := "hello world"
	dataB, err := codec.encode([]byte(dataStr))
	if err != nil {
		t.Log("Test failed while encode,err:", err)
		t.FailNow()
	}
	codec2 := new(CodecInfo)
	data2, err := codec2.decode(dataB)
	if err != nil {
		t.Log("Test failed while decode,err:", err)
		t.FailNow()
	}
	if !reflect.DeepEqual(codec, codec2) {
		t.Log(cmp.Diff(codec, codec2))
		t.FailNow()
	}
	if !reflect.DeepEqual(dataStr, string(data2)) {
		t.Log(cmp.Diff(dataStr, string(data2)))
		t.FailNow()
	}

}
