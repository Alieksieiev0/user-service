package grpc

import (
	"context"
	"net"

	"github.com/Alieksieiev0/user-service/api/proto"
	"github.com/Alieksieiev0/user-service/internal/models"
	"github.com/Alieksieiev0/user-service/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServer struct {
}

func NewServer() *GRPCServer {
	return &GRPCServer{}
}

func (us *GRPCServer) Start(addr string, service services.UserService) error {
	grpcUserService := NewGRPCUserServiceServer(service)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	proto.RegisterUserServiceServer(server, grpcUserService)

	return server.Serve(ln)
}

type GRPCUserServiceServer struct {
	service services.UserService
	proto.UnimplementedUserServiceServer
}

func NewGRPCUserServiceServer(service services.UserService) *GRPCUserServiceServer {
	return &GRPCUserServiceServer{
		service: service,
	}
}

func (us *GRPCUserServiceServer) GetByUsername(
	ctx context.Context,
	req *proto.UsernameRequest,
) (*proto.UserResponse, error) {
	user, err := us.service.GetByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	resp := &proto.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	}
	return resp, err
}

func (us *GRPCUserServiceServer) Save(
	ctx context.Context,
	req *proto.UserRequest,
) (*emptypb.Empty, error) {
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}
	err := us.service.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
