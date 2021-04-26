package svc

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"go-zero-example/service/search/cmd/api/internal/config"
	"go-zero-example/service/search/cmd/api/internal/middleware"
	"go-zero-example/service/user/cmd/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
