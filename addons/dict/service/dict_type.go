package service

import (
	"dzhgo/addons/dict/dao"
	"dzhgo/addons/dict/model"
	"github.com/gzdzh-cn/dzhcore"
)

type DictTypeService struct {
	*dzhcore.Service
}

func NewDictTypeService() *DictTypeService {
	return &DictTypeService{
		Service: &dzhcore.Service{
			Dao:   &dao.AddonsDictType,
			Model: model.NewDictType(),
			ListQueryOp: &dzhcore.QueryOp{
				KeyWordField: []string{"name"},
			},
		},
	}
}
