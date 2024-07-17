package middleware

import (
	logic "dzhgo/internal/logic/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BaseLog(r *ghttp.Request) {
	var (
		ctx               = r.GetCtx()
		BaseSysLogService = logic.NewsBaseSysLogService()
	)
	BaseSysLogService.Record(ctx)

	r.Middleware.Next()

}
