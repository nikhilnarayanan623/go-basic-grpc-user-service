package service

import (
	"context"
	"net/http"

	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/domain"
	"github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/pb"
	usecase "github.com/nikhilnarayanan623/go-basic-grpc-user-service/pkg/usecase/interfaces"
)

type serviceServer struct {
	useCase usecase.UserUseCase
	pb.UserServiceServer
}

func NewServiceServer(useCase usecase.UserUseCase) pb.UserServiceServer {
	return &serviceServer{
		useCase: useCase,
	}
}

func (c *serviceServer) FindUserByEmail(ctx context.Context, req *pb.FindUserByEmailRequest) (*pb.FindUserResponse, error) {

	user, err := c.useCase.FindUserByEmail(context.Background(), req.Email)
	if err != nil {
		return &pb.FindUserResponse{
			Response: &pb.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to find with given email",
				Error:      err.Error(),
			},
		}, nil
	}

	return &pb.FindUserResponse{
		Response: &pb.Response{
			StatusCode: http.StatusOK,
			Message:    "successfully found user data details",
			Error:      "",
		},
		User: &pb.User{
			UserId:    user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
		},
	}, nil
}

func (c *serviceServer) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserResponse, error) {

	userId, err := c.useCase.SaveUser(context.Background(), domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})
	if err != nil {
		return &pb.SaveUserResponse{
			Response: &pb.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to save user details",
				Error:      err.Error(),
			},
		}, nil
	}

	return &pb.SaveUserResponse{
		Response: &pb.Response{
			StatusCode: http.StatusOK,
			Message:    "successfully user details saved",
			Error:      "",
		},
		UserId: userId,
	}, nil
}
