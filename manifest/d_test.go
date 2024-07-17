package test

import (
	"context"
	"dzhgo/internal/dao"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	_ "github.com/gzdzh-cn/dzhcore/contrib/drivers/mysql"
	"testing"
)

var (
	ctx context.Context
)

func DBDao(s *gdb.Model) *gdb.Model {
	return s
}

func Test(t *testing.T) {

	var db = DBDao(dao.BaseSysAddons.Ctx(ctx))

	db.Where(dao.BaseSysAddons.Columns().Id, "1770001810119987200")

	all, err := db.All()
	if err != nil {
		return
	}
	g.Dump("all", all)

}
