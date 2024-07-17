package admin

import (
	"dzhgo/addons/dict/service"
	"github.com/gzdzh-cn/dzhcore"
)

type DictTypeController struct {
	*dzhcore.Controller
}

func init() {
	var dictTypeController = &DictTypeController{
		&dzhcore.Controller{
			Prefix:  "/admin/dict/type",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictTypeService(),
		},
	}
	// 注册路由
	dzhcore.RegisterController(dictTypeController)
}
