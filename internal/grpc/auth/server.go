package auth

import (
	"context"

	sso_v1 "github.com/ilyababichev/authorization-service/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

// Нужно для того чтобы не реализовывать все
// методы интерфейса protobuf
type serverAPI struct {
	sso_v1.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Login(
		ctx context.Context,
		email string,
		password string,
		appID int,
	) (token string, err error)
	RegisterNewUser(
		ctx context.Context,
		email string,
		password string,
	) (userID int64, err error)
	IsAdmin(
		ctx context.Context,
		userID int64,
	) (bool, error)
}

// Регистрация обработчика
func Register(gRPC *grpc.Server, auth Auth) {
	sso_v1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) Login(
	ctx context.Context,
	request *sso_v1.LoginRequest,
) (*sso_v1.LoginResponse, error) {
	if request.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if request.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	if request.GetAppId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "appid is required")
	}

	token, err := s.auth.Login(ctx, request.Email, request.Password, int(request.AppId))
	if err != nil {
		return nil, status.Error(codes.Internal, "intertnal error")
	}

	return &sso_v1.LoginResponse{
		Token: token,
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	request *sso_v1.RegisterRequest,
) (*sso_v1.RegistreResponse, error) {
	if request.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}

	if request.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	userID, err := s.auth.RegisterNewUser(ctx, request.Email, request.Password)

	if err != nil {
		return nil, status.Error(codes.Internal, "intertnal error")
	}

	return &sso_v1.RegistreResponse{
		UserId: userID,
	}, nil
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	request *sso_v1.IsAdminRequest,
) (*sso_v1.IsAdminResponse, error) {
	if request.GetUserId() == emptyValue {
		return nil, status.Error(codes.InvalidArgument, "userid is required")
	}

	isAdmin, err := s.auth.IsAdmin(ctx, request.GetUserId())

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "userid is required")
	}

	return &sso_v1.IsAdminResponse{
		IsAdmin: isAdmin,
	}, nil
}
