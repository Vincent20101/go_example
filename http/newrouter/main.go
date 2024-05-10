package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	address := fmt.Sprintf("%s:%d", "0.0.0.0", 8080)
	r := mux.NewRouter()
	//r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/test", handler)
	r.HandleFunc("/products", handler).Methods("POST")
	//r.HandleFunc("/articles", handler).Methods("GET")
	r.HandleFunc("/articles/{id}", handler).Methods("GET", "PUT")
	r.HandleFunc("/authors", handler).Queries("surname", "{surname}")
	r.HandleFunc("/", handler)
	err := http.ListenAndServe(address, r)
	if err != nil {
		log.Panic("ListenAndServe err:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Client request is coming\n")
	ctx := r.Context()
	fmt.Println("r====", r)
	fmt.Println("ctx====", ctx)

	//time.Sleep(10 * time.Second)

	select {
	case <-ctx.Done():
		log.Printf("\tClient has broken\n")
		break
	default:
		log.Printf("\tSend client response\n")
		w.Write([]byte("I am OK"))
	}
}
