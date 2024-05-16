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

func (s *simpleServer) Address() string{
	return s.addr
}

func (s *simpleServer) IsAlive() bool{
	return true
}

func (s *simpleServer) Serve(w http.ResponseWriter, r *http.Request){}

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


func NewLoadBalancer(port string, servers []Server) *LoadBalancer{
	return &LoadBalancer{
		port: port,
		servers: servers,
		roundRobinCount: 0,
	}
}

func handleErr(err error){
	if err != nil{
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server{}

func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request){}

func main(){
	servers:= []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("http://www.bing.com"),
		newSimpleServer("http://www.duckduckgo.com"),
	}

	lb := NewLoadBalancer("8000", servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request){
		lb.serveProxy(w, r)
	}

	http.HandleFunc("/", handleRedirect)
	
	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}