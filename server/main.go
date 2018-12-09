package main

import (
"golang.org/x/net/context"

"github.com/Kishinskiy/grpc_blockchain/proto"
	"github.com/Kishinskiy/grpc_blockchain/server/blockchain"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main(){
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("unable to listen on 8080 port: %w", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBlockchainServer(srv, &Server{})
	srv.Serve(listener)
}

type Server struct {


}

func (s *Server) GetBlockchain(context.Context, *proto.GetBlockchainRequest) (*proto.GetBlockchainResponse, error) {
	return new(proto.GetBlockchainResponse), nil
}

func (s *Server) AddBlock(context.Context, *proto.AddBlockRequest) (*proto.AddBlockResponse, error){
  return new(proto.AddBlockResponse), nil
}
