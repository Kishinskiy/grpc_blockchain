package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/grpc"
	"log"
	"net"
)

func main(){
	listener, err := net.Listen("tcp", "8080")
	if err != nil {
		log.Fatalf("unable to listen on 8080 port: %w", err)
	}

	srv: := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, nil)
	srv.Serve()
}

type Server {

}

func (s *Server) AddBlock)context.Context, *AddBlockRequest) ()
