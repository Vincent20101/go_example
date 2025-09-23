package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/net/http2"
)

func main() {
	dialTLSFunc, localAddrFunc, remoteAddrFunc := GetDialTLS("http client")
	client := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS:   dialTLSFunc,
		},
	}
	
	dialTLSFunc2, localAddrFunc2, remoteAddrFunc2 := GetDialTLS("http client")
	client2 := &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS:   dialTLSFunc2,
		},
	}
	defaultClientMap.LoadOrStore("nfType", client)

	url := "http://127.0.0.1:8080"
	resp, err := client.Head(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read response error:", err)
		return
	}

	fmt.Printf("Server response: %s\n", body)
	fmt.Println("local:", localAddrFunc(), "remote:", remoteAddrFunc())

	// loadClient, _ := defaultClientMap.Load("nfType")
	request, err1 := http.NewRequest(http.MethodPost, "http://172.0.3.64:8080/hello", nil)
	if err1 != nil {
		log.Fatal(err1)
	}
	response, err2 := client2.Do(request)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(response)
	fmt.Println("local:", localAddrFunc2(), "remote:", remoteAddrFunc2())
}

type ConnAddr struct {
	Ipv4 net.IP
	Ipv6 net.IP
	Port uint16
}

var defaultClientMap = &clientMap{value: make(map[string]*http.Client, 10)}

type clientMap struct {
	lock  sync.RWMutex
	value map[string]*http.Client
}

func (c *clientMap) Load(key string) (*http.Client, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	ret, ok := c.value[key]
	return ret, ok
}

func (c *clientMap) LoadOrStore(key string, in *http.Client) *http.Client {
	c.lock.Lock()
	defer c.lock.Unlock()
	if ret, ok := c.value[key]; ok {
		return ret
	}
	c.value[key] = in
	return in
}

func GetDialTLS(nfType string) (dialFunc func(network, addr string, cfg *tls.Config) (net.Conn, error),
	localAddr, remoteAddr func() ConnAddr) {
	var l, r ConnAddr
	dialFunc = func(network, addr string, cfg *tls.Config) (net.Conn, error) {
		var (
			c   net.Conn
			err error
		)
		c, err = net.DialTimeout(network, addr, time.Second*time.Duration(2))
		if err != nil {
			return nil, err
		}
		// localAddr
		if c.LocalAddr() != nil {
			if lAddr, ok := c.LocalAddr().(*net.TCPAddr); ok {
				l.Port = uint16(lAddr.Port)
				switch {
				case lAddr.IP.To4() != nil && len(l.Ipv4) == 0:
					l.Ipv4 = lAddr.IP
					log.Println(nfType, " pcap local v4: ", lAddr.IP.To4())
				case lAddr.IP.To16() != nil && lAddr.IP.To4() == nil && len(l.Ipv6) == 0:
					l.Ipv6 = lAddr.IP
					log.Println(nfType, " pcap local v6: ", lAddr.IP.To16())
				}
			}
		}

		// remoteAddr
		if c.RemoteAddr() != nil {
			if rAddr, ok := c.RemoteAddr().(*net.TCPAddr); ok {
				r.Port = uint16(rAddr.Port)
				switch {
				case rAddr.IP.To4() != nil && len(r.Ipv4) == 0:
					r.Ipv4 = rAddr.IP
					log.Println(nfType, " pcap remote v4: ", rAddr.IP.To4())
				case rAddr.IP.To16() != nil && rAddr.IP.To4() == nil && len(r.Ipv6) == 0:
					log.Println(nfType, " pcap remote v6: ", rAddr.IP.To16())
					r.Ipv6 = rAddr.IP
				}
			}
		}
		return c, err
	}
	remoteAddr = func() ConnAddr { return r }
	localAddr = func() ConnAddr { return l }
	return
}
