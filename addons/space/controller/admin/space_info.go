package admin

import (
	"dzhgo/addons/space/service"
	"github.com/gzdzh/dzhcore"
)

type SpaceInfoController struct {
	*dzhcore.Controller
}

func init() {
	var spaceInfoController = &SpaceInfoController{
		&dzhcore.Controller{
			Prefix:  "/admin/space/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewSpaceInfoService(),
		},
	}
	// 注册路由
	dzhcore.RegisterController(spaceInfoController)
}
