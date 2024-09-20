// Code generated by goctl-types plugin. DO NOT EDIT.
package types

import (
	"time"
)

var (
	_ = time.Now()
)

type GetTotalDownloadsRequest struct {
	Org       string `path:"org"`
	Repo      string `path:"repo"`
	Container string `path:"container"`
}

type GetTotalDownloadsResponse struct {
	Message string `json:"message"`
}
