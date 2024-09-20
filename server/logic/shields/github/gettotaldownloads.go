package github

import (
	"context"
	"encoding/json"
	"fmt"
	"go-serverless-vercel/server/svc"
	"go-serverless-vercel/server/types"
	"net/http"

	"github.com/gocolly/colly/v2"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTotalDownloads struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
}

func NewGetTotalDownloads(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter) *GetTotalDownloads {
	return &GetTotalDownloads{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx, w: w,
	}
}

func (l *GetTotalDownloads) GetTotalDownloads(req *types.GetTotalDownloadsRequest) error {
	c := colly.NewCollector(colly.Async(true))

	var totalDownloads int

	c.OnHTML("div.lh-condensed.d-flex.flex-column.flex-items-baseline.pr-1 h3", func(e *colly.HTMLElement) {
		totalDownloadsText := e.Attr("title")
		totalDownloads = cast.ToInt(totalDownloadsText)
	})

	c.OnRequest(func(r *colly.Request) {})

	c.OnResponse(func(r *colly.Response) {})

	err := c.Visit(fmt.Sprintf("https://github.com/%s/%s/pkgs/container/%s", req.Org, req.Repo, req.Container))
	if err != nil {
		return err
	}

	c.Wait()

	resp := &types.Badge{
		SchemaVersion: 1,
		Label:         "pulls",
		Message:       fmt.Sprintf("%d", totalDownloads),
		Color:         "orange",
	}

	marshal, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	l.w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	_, _ = l.w.Write(marshal)
	return nil
}
