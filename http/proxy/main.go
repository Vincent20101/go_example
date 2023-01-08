package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"go.uber.org/zap"
)

func main() {
	lg, err := NewZapLogger()
	if err != nil {
		log.Fatalf("cannot create zap logger: %v", err)
	}

	// Set up reverse proxy for k8s load balancing experiment.
	u, err := url.Parse("http://nginx-service")
	if err != nil {
		lg.Sugar().Fatalf("cannot parse nginx url : %v", err)
	}
	p := httputil.NewSingleHostReverseProxy(u)
	p.Transport = &http.Transport{
		DisableKeepAlives: false,
	}
	http.Handle("/lb-keepalive", p)
	p2 := httputil.NewSingleHostReverseProxy(u)
	p2.Transport = &http.Transport{
		DisableKeepAlives: true,
	}
	http.Handle("/lb-nokeepalive", p2)

	//http.Handle("/", mux)
	var addr = new(string)
	*addr = ":8089"
	lg.Sugar().Infof("grpc gateway started at %s", *addr)
	lg.Sugar().Fatal(http.ListenAndServe(*addr, nil))
}

// NewZapLogger creates a new zap logger
func NewZapLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.TimeKey = ""
	return cfg.Build()
}
