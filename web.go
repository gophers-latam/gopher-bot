package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func web(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello gophers! %s", time.Now())
}

// ipv6
func ListenAndServe(addr string) error {
	http.HandleFunc("/", web)
	srv := &http.Server{Addr: addr, Handler: nil}
	addr = srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp6", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln.(*net.TCPListener))
}
