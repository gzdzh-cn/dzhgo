package admin

import (
	"context"
	v1 "dzhgo/addons/dict/api/v1"
	"dzhgo/addons/dict/service"
	"github.com/gzdzh-cn/dzhcore"
)

type DictInfoController struct {
	*dzhcore.Controller
}

func init() {
	var dictInfoController = &DictInfoController{
		&dzhcore.Controller{
			Prefix:  "/admin/dict/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictInfoService(),
		},
	}

	// 注册路由
	dzhcore.RegisterController(dictInfoController)
}

// Data 方法 获得字典数据
func (c *DictInfoController) Data(ctx context.Context, req *v1.DictInfoDataReq) (res *dzhcore.BaseRes, err error) {

	data, err := service.NewDictInfoService().Data(ctx, req.Types)
	res = dzhcore.Ok(data)
	return
}
