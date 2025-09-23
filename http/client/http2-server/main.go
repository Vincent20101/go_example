package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2/h2c"

	"golang.org/x/net/http2"
)

// func main() {
// 	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
// 		// 获取客户端地址
// 		remoteAddr := r.RemoteAddr
// 		// 获取服务器本地地址
// 		localAddr := r.Context().Value(http.LocalAddrContextKey).(net.Addr).String()
// 		fmt.Printf("Client Address: %s, Server Address: %s\n", remoteAddr, localAddr)
// 		fmt.Fprintf(w, "Hello, this is the server response!")
// 	})

// 	serverAddr := "localhost:8080"
// 	fmt.Printf("Starting server on %s...\n", serverAddr)
// 	err := http.ListenAndServe(serverAddr, nil)
// 	if err != nil {
// 		fmt.Println("Server error:", err)
// 	}
// }


func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Host)
		fmt.Println("r.URL.Host:", r.URL.Host)
		fmt.Println("r.URL", r.URL)
		fmt.Println(r.Method)
		fmt.Println(r.Proto)
		fmt.Println(r.ProtoMinor)
		fmt.Println(r.ProtoMinor)
		//time.Sleep(time.Second * 100)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		//w.Write([]byte("Hello, HTTP/2.0 h2c!"))
		fmt.Fprintf(w, "Hello, HTTP/2 Cleartext!")
	})

	h2s := &http2.Server{}

	server := &http.Server{
		Addr: ":8080",
		//Handler:      h2cHandler(handler, h2s),
		Handler:      h2cHandler(handler, h2s),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	log.Fatal(server.ListenAndServe())
}

func h2cHandler(h http.Handler, h2s *http2.Server) http.Handler {
	return h2c.NewHandler(h, h2s)
}