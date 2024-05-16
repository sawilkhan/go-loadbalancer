package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

type Server interface{
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

func newSimpleServer(addr string) *simpleServer{
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr: addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct{
	port string
	roundRobinCount int
	servers []Server
}


func handleErr(err error){
	if err != nil{
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}