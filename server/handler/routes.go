package handler

import (
	version "go-serverless-vercel/server/handler/version"
	"go-serverless-vercel/server/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/version",
				Handler: version.GetVersion(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
