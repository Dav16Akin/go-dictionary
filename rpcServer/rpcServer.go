package rpcserver

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)


type Args struct {}
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return  nil
}


func Run() {
	//creating a RPC server here
	timeserver := new(TimeServer)

	//registering the rpc server
	rpc.Register(timeserver)
	rpc.HandleHTTP()

	l, err:= net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	http.Serve(l, nil)
}