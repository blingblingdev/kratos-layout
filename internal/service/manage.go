package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/ffy/kratos-layout/api/manage/v1"
	"github.com/ffy/kratos-layout/internal/biz"
)

type ManageService struct {
	v1.UnimplementedManageServer

	uc  *biz.ManageUseCase
	log *log.Helper
}

func NewManageService(uc *biz.ManageUseCase, logger log.Logger) *ManageService {
	return &ManageService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ManageService) Ping(ctx context.Context, req *v1.PingReq) (resp *v1.PingResp, err error) {
	s.log.WithContext(ctx).Infof("Ping Received")

	resp = &v1.PingResp{Res: "pong"}
	if req.GetMsg() == "error" {
		err = v1.ErrorBadRequest("bad msg: %s", req.GetMsg())
		return
	}
	return
}
