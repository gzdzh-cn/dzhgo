package admin

import (
	logic "dzhgo/internal/logic/sys"
	"github.com/gzdzh/dzhcore"
)

type BaseSysMenuController struct {
	*dzhcore.Controller
}

func init() {
	var baseSysMenuController = &BaseSysMenuController{
		&dzhcore.Controller{
			Prefix:  "/admin/base/sys/menu",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: logic.NewsBaseSysMenuService(),
		},
	}
	// 注册路由
	dzhcore.RegisterController(baseSysMenuController)
}
