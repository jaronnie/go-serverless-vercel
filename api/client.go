package api

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"go-serverless-vercel/server/config"
	"go-serverless-vercel/server/handler"
	"go-serverless-vercel/server/middleware"
	"go-serverless-vercel/server/svc"
	"net/http"
)

var (
	server *rest.Server
)

func init() {
	var c config.Config
	c.RestConf = rest.RestConf{
		Host: "0.0.0.0",
		Port: 8001,
		ServiceConf: service.ServiceConf{
			Name: "go-serverless-vercel",
		},
	}

	server = rest.MustNewServer(c.RestConf)
	middleware.RegisterMiddlewares(server)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
}

func Index(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}
