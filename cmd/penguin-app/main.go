package main

import (
	"flag"
	"fmt"
	"net/http"

	handlers "github.com/MarkHXB/penguin-api/pkg/handlers"
)

const(
	defaultAddress	string	= "0.0.0.0"
	defaultPort		int		= 80
)

var(
	address	string
	port	int
)

func main() {
	flag.StringVar(&address, "addr",defaultAddress,"Server's Address")
	flag.IntVar(&port, "port",defaultPort,"Server's Port")
	flag.Parse()
	
	runAPI()
}

func runAPI() {
	fullAddress	:= fmt.Sprintf("%s:%d",address,port)

	http.HandleFunc("/test", handlers.TestHandler)
	http.ListenAndServe(fullAddress, nil)
}
