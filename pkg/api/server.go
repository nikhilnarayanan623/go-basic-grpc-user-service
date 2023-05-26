package api

import (
	"fmt"
	"net"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/pb"
	"google.golang.org/grpc"
)

type Server struct {
	grpcServer *grpc.Server
	lisn       net.Listener
}

func NewUserServer(cfg *config.Config, serverService pb.UserServiceServer) (*Server, error) {

	gprpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(gprpcServer, serverService)

	lisn, err := net.Listen("tcp", cfg.ServicePort)
	if err != nil {
		return nil, fmt.Errorf("faild to create net listener error:%s", err.Error())
	}

	return &Server{
		grpcServer: gprpcServer,
		lisn:       lisn,
	}, nil
}

func (c *Server) Start() error {
	fmt.Println("user service listening....")
	return c.grpcServer.Serve(c.lisn)
}
