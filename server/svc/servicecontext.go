package svc

import (
	"go-serverless-vercel/server/config"
	"go-serverless-vercel/server/custom"
)

type ServiceContext struct {
	Config config.Config

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Custom: custom.New(),
	}
}
