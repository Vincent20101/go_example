package main

import (
	"fmt"
	"net/http"
	"strings"
)

type EventHdr struct {
	HttpHeader http.Header
}

func main() {
	hd := EventHdr{
		HttpHeader: map[string][]string{
			"lhb": {"test"},
		},
	}
	//hd.HttpHeader.Set("lhb", "env=testenv")
	values := hd.HttpHeader.Values("lhb")
	//values := hd.HttpHeader.Get("lhb")

	fmt.Println(strings.ToLower("3gpp-Sbi-Target-Apiroot"))
	fmt.Println(values)
}
