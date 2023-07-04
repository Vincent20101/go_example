package main

import (
	"encoding/asn1"
	"fmt"
	"log"
	"os"
)

func main() {
	// ASN.1数据
	//asn1data := []byte{0x30, 0x0d, 0x02, 0x01, 0x01, 0x0a, 0x06, 0x07, 0x2b, 0x06, 0x01, 0x02, 0x01, 0x01, 0x05, 0x00}
	asn1data, err := os.ReadFile("/root/project/goex/go_example/asn/test.u")
	//b, err := ioutil.ReadFile("./asn1-file/lab.casa.com_pcms.chg")
	if err != nil {
		fmt.Println(err)
	}

	// 解析ASN.1数据
	var parsedData interface{}
	_, err = asn1.Unmarshal(asn1data, &parsedData)
	if err != nil {
		log.Fatal(err)
	}

	// 打印解析结果
	fmt.Printf("Parsed Data: %#v\n", parsedData)
	//fmt.Printf("Remaining Data: %x\n", rest)
}
