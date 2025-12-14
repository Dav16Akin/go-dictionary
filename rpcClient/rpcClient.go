package rpcClient

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	rpcserver "github.com/Dav16Akin/go-dictionary/rpcServer"
)

func Client() {
	command := ""
    if len(os.Args) > 1 {
        command = os.Args[1]
    }

    switch command {
    case "client":
        // --- CLIENT MODE ---
        // Run this manually: go run main.go client
        runClient()

    case "server":
        // --- SERVER MODE ---
        // You can run this explicitly: go run main.go server
        rpcserver.Run()

    default:
        // --- DEFAULT (AIR) MODE ---
        // If no args are provided (which is what Air does), run the server!
        fmt.Println("No command provided, defaulting to RPC Server for Air...")
        rpcserver.Run()
    }
}

func runClient() {
	var reply int64
	args := &rpcserver.Args{}
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	err = client.Call("TimeServer.GiveServerTime", args, &reply)

	if err != nil {
		log.Fatal("arith error: ", err)
	}
	log.Printf("%d", reply)
}
