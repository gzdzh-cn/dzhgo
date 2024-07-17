package app

import (
	"context"

	"dzhgo/addons/dict/service"
	"github.com/gzdzh-cn/dzhcore"

	"github.com/gogf/gf/v2/frame/g"
)

type DictInfoController struct {
	*dzhcore.Controller
}

func init() {
	var dictInfoController = &DictInfoController{
		&dzhcore.Controller{
			Prefix:  "/app/dict/info",
			Api:     []string{"Add", "Delete", "Update", "Info", "List", "Page"},
			Service: service.NewDictInfoService(),
		},
	}
	// 注册路由
	dzhcore.RegisterController(dictInfoController)
}

// DictInfoDataReq Data 方法请求
type DictInfoDataReq struct {
	g.Meta `path:"/data" method:"POST"`
	Types  []string `json:"types"`
}

// Data 方法 获得字典数据
func (c *DictInfoController) Data(ctx context.Context, req *DictInfoDataReq) (res *dzhcore.BaseRes, err error) {

	data, err := service.NewDictInfoService().Data(ctx, req.Types)
	res = dzhcore.Ok(data)
	return
}
