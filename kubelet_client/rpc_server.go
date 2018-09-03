package kubelet_client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	api "plugin-intermediary/kubelet_client/api"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

var CurrentTopo string

// SayHello implements helloworld.GreeterServer
func (s *server) InformTopology(ctx context.Context, in *api.TopologyRequest) (*api.TopologyReply, error) {
	CurrentTopo = in.Topo
	return &api.TopologyReply{Message: "Hello "}, nil
}

func RpcService() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
