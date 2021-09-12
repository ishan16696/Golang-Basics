package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/server"
)

func main() {
	server.RunRPC()
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &server.Args{3, 4}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

}
