package main

import (
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
	var sj = `
{"PublisherId":29764537256,"evntType":"SM_PCAP_TRACE","evntAction":"","evntKey":"","evntData":"gqRTdXBptGltc2ktNDUwMDUxMjMwMDAwMDAzrFN1YnNjcmliZXJJZIakSW1zabRpbXNpLTQ1MDA1MTIzMDAwMDAwM6RJbWVpoKZNc2lzZG62bXNpc2RuLTQ1MDA1MTIzMDAwMDAwM6ZVZWlwVjSgplVlaXBWNqCtVHJhY2VVZWlwVHlwZaA="}
`
	var ps PubSubEvent
	err := json.Unmarshal([]byte(sj), &ps)
	if err != nil {
		fmt.Printf("err")
	}
	var m map[string]interface{}
	err = msgpack.Unmarshal(ps.Data, &m)
	if err != nil {
		fmt.Printf("err")
	}
	fmt.Println(m)
}
