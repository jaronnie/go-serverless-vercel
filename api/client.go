package api

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"go-serverless-vercel/server/config"
	"go-serverless-vercel/server/handler"
	"go-serverless-vercel/server/middleware"
	"go-serverless-vercel/server/svc"
	"net/http"
	"os"
)

var (
	server *rest.Server
)

func init() {
	var c config.Config
	c.Rest.RestConf = rest.RestConf{
		Host: "0.0.0.0",
		Port: 8001,
		ServiceConf: service.ServiceConf{
			Name: "go-serverless-vercel",
		},
	}

	// set up logger
	if err := logx.SetUp(c.Log.LogConf); err != nil {
		logx.Must(err)
	}
	if c.Log.LogConf.Mode != "console" {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	server = rest.MustNewServer(c.Rest.RestConf)
	middleware.RegisterMiddlewares(server)

	svcCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, svcCtx)

	// server add custom routes
	svcCtx.Custom.AddRoutes(server)

	logx.Infof("Starting rest server at %s:%d...", svcCtx.Config.Rest.Host, svcCtx.Config.Rest.Port)
}

func Index(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}
