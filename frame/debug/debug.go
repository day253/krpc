package debug

import (
	"net/http"
	"net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samber/lo"
)

type HttpServer struct {
	address string
}

func (h *HttpServer) Start() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.HandleFunc("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.HandleFunc("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.HandleFunc("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.HandleFunc("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	lo.Must0(http.ListenAndServe(h.address, mux))
}

func NewHttpServer(address string) *HttpServer {
	return &HttpServer{
		address: address,
	}
}
