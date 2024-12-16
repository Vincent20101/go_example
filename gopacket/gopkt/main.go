package main

import (
	"fmt"
	"log"

	"github.com/ghedo/go.pkt/capture/pcap"
)

func main() {
	tMap := make(map[string]string)
	tMap[""] = "test"
	fmt.Println(tMap)
	src, err := pcap.Open("enp1s0")
	if err != nil {
		log.Fatal("enp1s0:", err)
	}
	defer src.Close()

	// you may configure the source further, e.g. by activating
	// promiscuous mode.

	//err = src.Activate()
	//if err != nil {
	//	log.Fatal("test Activate:", err)
	//}

	//err = src.Inject([]byte("random data"))
	//if err != nil {
	//	log.Fatal("Inject:", err)
	//}

	for {
		buf, err := src.Capture()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("PACKET!!!", buf)

		// do something with the packet
	}
}
