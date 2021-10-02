package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ffy/kratos-layout/internal/biz"
)

type manageRepo struct {
	data *Data
	log  *log.Helper
}

func NewManageRepo(data *Data, logger log.Logger) biz.ManageRepo {
	return &manageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
