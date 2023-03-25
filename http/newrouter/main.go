package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8080)
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/test", testHandle)
	err := http.ListenAndServe(address, router)
	if err != nil {
		log.Panic("ListenAndServe err:", err)
	}
}

func testHandle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Client request is coming\n")
	ctx := r.Context()
	fmt.Println("r====", r)
	fmt.Println("ctx====", ctx)

	time.Sleep(10 * time.Second)

	select {
	case <-ctx.Done():
		log.Printf("\tClient has broken\n")
		break
	default:
		log.Printf("\tSend client response\n")
		w.Write([]byte("I am OK"))
	}
}
