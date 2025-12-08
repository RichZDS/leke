// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "leke/api/user/v1"
)

type (
	IFans interface {
		FansList(ctx context.Context, req *v1.FansListReq) (res *v1.FansListRes, err error)
		FansView(ctx context.Context, req *v1.FansViewReq) (res *v1.FansViewRes, err error)
		FansUpdate(ctx context.Context, req *v1.FansUpdateReq) (res *v1.FansUpdateRes, err error)
		FansDelete(ctx context.Context, req *v1.FansDeleteReq) (res *v1.FansDeleteRes, err error)
		FansCreate(ctx context.Context, req *v1.FansCreateReq) (res *v1.FansCreateRes, err error)
	}
)

var (
	localFans IFans
)

func Fans() IFans {
	if localFans == nil {
		panic("implement not found for interface IFans, forgot register?")
	}
	return localFans
}

func RegisterFans(i IFans) {
	localFans = i
}
