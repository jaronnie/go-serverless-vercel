package handler

import (
	count "go-serverless-vercel/server/handler/count"
	shieldsgithub "go-serverless-vercel/server/handler/shields/github"
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

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/shields/github/:org/:repo/pkgs/container/:container/downloads",
				Handler: shieldsgithub.GetTotalDownloads(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}
