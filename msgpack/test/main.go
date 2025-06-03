package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack/v4"
)

type PubSubEvent struct {
	// who publish evnt
	PublisherId int64 `json:"PublisherId"`
	// evntType defined above
	Type string `json:"evntType"`
	// evntAction defined above
	Action string `json:"evntAction"`
	// evntKey related to a NF instance
	Key string `json:"evntKey"`
	// msgpack packed event structure
	Data []byte `json:"evntData"`
}

func main() {
	// 	var sj = `
	// {"PublisherId":29764537256,"evntType":"SM_PCAP_TRACE","evntAction":"","evntKey":"","evntData":"gqRTdXBptGltc2ktNDUwMDUxMjMwMDAwMDAzrFN1YnNjcmliZXJJZIakSW1zabRpbXNpLTQ1MDA1MTIzMDAwMDAwM6RJbWVpoKZNc2lzZG62bXNpc2RuLTQ1MDA1MTIzMDAwMDAwM6ZVZWlwVjSgplVlaXBWNqCtVHJhY2VVZWlwVHlwZaA="}
	// `

	var sj = `
{"PublisherId":9429453593,"evntType":"SUP_FEAT","evntAction":"PUT","evntKey":"\u0026f2b43e11-762a-49b9-b5d4-4c3452ecb8f8","evntData":"0A0BAAAAAAA="}
`
	var ps PubSubEvent
	err := json.Unmarshal([]byte(sj), &ps)
	if err != nil {
		fmt.Printf("err")
	}
	fmt.Println(ps.Data)
	fmt.Println(binary.LittleEndian.Uint64(ps.Data))
	var m map[string]interface{}
	err = msgpack.Unmarshal(ps.Data, &m)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Println(m)
}
