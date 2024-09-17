package main

import (
	"flag"
	"go-serverless-vercel/internal/config"
	"go-serverless-vercel/internal/handler"
	"go-serverless-vercel/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var (
	server     *rest.Server
	configFile = flag.String("f", "etc/etc.yaml", "set config file")
)

func init() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server = rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
}

func Index(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}
