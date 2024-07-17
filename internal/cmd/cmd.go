package cmd

import (
	"context"
	"dzhgo/addons"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"dzhgo/internal/base"
	"github.com/gzdzh/dzhcore"

	_ "dzhgo/internal/logic"
	_ "dzhgo/packed"

	_ "github.com/gzdzh/dzhcore/contrib/files/local"
	_ "github.com/gzdzh/dzhcore/contrib/files/oss"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			if dzhcore.IsRedisMode {
				go dzhcore.ListenFunc(ctx)
			}

			dzhcore.NewInit()

			base.NewInit()

			addons.NewInit()

			s := g.Server()
			s.Run()
			return nil
		},
	}
)
