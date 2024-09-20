package github

import (
	"go-serverless-vercel/server/logic/shields/github"
	"go-serverless-vercel/server/svc"
	"go-serverless-vercel/server/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTotalDownloads(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTotalDownloadsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := github.NewGetTotalDownloads(r.Context(), svcCtx)
		resp, err := l.GetTotalDownloads(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
