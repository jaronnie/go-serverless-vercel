package handler

import (
	count "go-serverless-vercel/server/handler/count"
	"go-serverless-vercel/server/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/count",
				Handler: count.GetCount(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
