package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type ManageRepo interface {
}

type ManageUseCase struct {
	repo ManageRepo
	log  *log.Helper
}

func NewManageUseCase(repo ManageRepo, logger log.Logger) *ManageUseCase {
	return &ManageUseCase{repo: repo, log: log.NewHelper(logger)}
}
