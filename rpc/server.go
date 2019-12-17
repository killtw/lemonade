package rpc

import (
	"fmt"
	"github.com/killtw/lemonade/lemonade"
	"github.com/killtw/lemonade/rpc/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server struct{}

func (s *Server) Replace(ctx context.Context, req *protos.ReplaceRequest) (*protos.ReplaceReply, error) {
	original := req.Message
	filtered, matches := lemonade.Replace(original)

	return &protos.ReplaceReply{
		Original: original,
		Message: filtered,
		Matches: matches,
	}, nil
}

func RunGRPCServer() error {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalln(err)

		return err
	}

	gs := grpc.NewServer()
	s := &Server{}
	protos.RegisterLemonadeServer(gs, s)

	reflection.Register(gs)
	fmt.Println("gRPC server is running")
	if err := gs.Serve(listener); err != nil {
		log.Fatalln("gRPC server err: ", err)

		return err
	}

	return nil
}
