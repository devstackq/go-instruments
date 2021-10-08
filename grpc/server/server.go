package server

import (
	"context"
	"net"

	proto "github.com/devstackq/proto"
	"google.golang.org/grpc"
)

type Greeter struct {}

//setter
func New ()*Greeter{
	return &Greeter{}
}

func (g *Greeter)Start()error {
	listener, err := net.Listen("tcp", "localhost:6969")
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServer, g)
	grpcServer.Serve(listener)

	return nil
}

func(g *Greeter)Say(ctx context.Context, r *proto.Request)(*proto.Reply,err){
	return &proto.Reply{
		Message:fmt.Sprintf("Hi yepta!", r.Name)
	},nil
}
protoc downlaod