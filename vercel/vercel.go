package vercel

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"go-serverless-vercel/api/config"
	"go-serverless-vercel/api/handler"
	"go-serverless-vercel/api/svc"
	"net/http"
)

var (
	server *rest.Server
	//configFile = flag.String("f", "../etc/etc.yaml", "set config file")
)

func init() {
	//flag.Parse()
	//
	var c config.Config
	//conf.MustLoad(*configFile, &c)
	c.RestConf = rest.RestConf{
		Host: "0.0.0.0",
		Port: 8001, g
		ServiceConf: service.ServiceConf{
			Name: "go-serverless-vercel",
		},
	}

	server = rest.MustNewServer(c.RestConf)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
}

func Index(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}
