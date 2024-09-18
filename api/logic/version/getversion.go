package version

import (
	"context"
	"net/http"

	"go-serverless-vercel/api/svc"
	"go-serverless-vercel/api/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVersion struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func Handler(w http.ResponseWriter, r *http.Request) {}

func NewGetVersion(ctx context.Context, svcCtx *svc.ServiceContext) *GetVersion {
	return &GetVersion{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVersion) GetVersion(req *types.GetVersionRequest) (resp *types.GetVersionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
