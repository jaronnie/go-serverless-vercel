package github

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"go-serverless-vercel/server/svc"
	"go-serverless-vercel/server/types"

	"github.com/gocolly/colly/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTotalDownloads struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTotalDownloads(ctx context.Context, svcCtx *svc.ServiceContext) *GetTotalDownloads {
	return &GetTotalDownloads{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTotalDownloads) GetTotalDownloads(req *types.GetTotalDownloadsRequest) (resp *types.Badge, err error) {
	c := colly.NewCollector(colly.Async(true))

	var totalDownloads int

	//c.OnHTML("div.lh-condensed.d-flex.flex-column.flex-items-baseline.pr-1", func(e *colly.HTMLElement) {
	//	e.ForEach("h3", func(_ int, h *colly.HTMLElement) {
	//		title := h.Attr("title")
	//		fmt.Println("Title:", title)
	//	})
	//})

	c.OnHTML("div.lh-condensed.d-flex.flex-column.flex-items-baseline.pr-1 h3", func(e *colly.HTMLElement) {
		totalDownloadsText := e.Attr("title")
		totalDownloads = cast.ToInt(totalDownloadsText)
	})

	c.OnRequest(func(r *colly.Request) {})

	c.OnResponse(func(r *colly.Response) {})

	err = c.Visit(fmt.Sprintf("https://github.com/%s/%s/pkgs/container/%s", req.Org, req.Repo, req.Container))
	if err != nil {
		return nil, err
	}

	c.Wait()

	resp = &types.Badge{
		SchemaVersion: 1,
		Label:         "image%20pulls",
		Message:       fmt.Sprintf("%d", totalDownloads),
		Color:         "orange",
	}
	return
}
