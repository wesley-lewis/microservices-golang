package client

import (
	"fmt"
	"log"
	"net/rpc"
	"wesley-rpc/server"
)

func Multiply(client *rpc.Client) {
	args := &server.Args{A: 100, B: 100}
	var reply int
	err := client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Arith error: ", err)
	}

	fmt.Printf("MULTIPLY\n%d x %d = %d\n", args.A, args.B, reply)
}

func Divide(client *rpc.Client) {
	args := &server.Args{A: 49, B: 7}
	var reply server.Quotient
	err := client.Call("Arith.Divide", args, &reply)
	if err != nil {
		log.Fatal("Arith error: ", err)
	}

	fmt.Printf("DIVIDE\n%d / %d = %d\n", args.A, args.B, reply.Quo)
}

func CreateClient() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	return client
}
