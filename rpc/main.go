package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"wesley-rpc/client"
	"wesley-rpc/server"
)

func main() {

	arith := new(server.Arith)

	rpc.Register(arith)
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	go http.Serve(l, nil)

	client_conn := client.CreateClient()
	client.Multiply(client_conn)
	client.Divide(client_conn)
}
