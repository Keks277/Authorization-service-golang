package auth

import (
	"context"

	sso_v1 "github.com/ilyababichev/authorization-service/protos/gen/go/sso"
	"google.golang.org/grpc"
)

// Нужно для того чтобы не реализовывать все
// методы интерфейса protobuf
type serverAPI struct {
	sso_v1.UnimplementedAuthServer
}

// Регистрация обработчика
func Register(gRPC *grpc.Server) {
	sso_v1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	request *sso_v1.LoginRequest,
) (*sso_v1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Register(
	ctx context.Context,
	request *sso_v1.RegisterRequest,
) (*sso_v1.RegistreResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	request *sso_v1.IsAdminRequest,
) (*sso_v1.IsAdminResponse, error) {
	panic("implement me")
}
