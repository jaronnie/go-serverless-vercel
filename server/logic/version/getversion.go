package version

import (
	"context"
	"go-serverless-vercel/server/svc"
	"go-serverless-vercel/server/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVersion struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVersion(ctx context.Context, svcCtx *svc.ServiceContext) *GetVersion {
	return &GetVersion{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVersion) GetVersion(req *types.GetVersionRequest) (resp *types.GetVersionResponse, err error) {
	resp = &types.GetVersionResponse{
		Version: "1",
	}
	return
}
