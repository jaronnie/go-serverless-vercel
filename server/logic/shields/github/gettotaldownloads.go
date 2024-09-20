package github

import (
	"context"
	"go-serverless-vercel/server/svc"
	"go-serverless-vercel/server/types"

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

func (l *GetTotalDownloads) GetTotalDownloads(req *types.GetTotalDownloadsRequest) (resp *types.GetTotalDownloadsResponse, err error) {

	return
}
