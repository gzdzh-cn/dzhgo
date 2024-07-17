package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/service"
	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhcore"
)

func init() {
	service.RegisterBaseSysPermsService(NewsBaseSysPermsService())
}

type sBaseSysPermsService struct {
}

func NewsBaseSysPermsService() *sBaseSysPermsService {
	return &sBaseSysPermsService{}
}

// permmenu 方法
func (c *sBaseSysPermsService) Permmenu(ctx context.Context, roleIds []string) (res interface{}) {
	type permmenu struct {
		Perms []string   `json:"perms"`
		Menus gdb.Result `json:"menus"`
	}
	var (
		baseSysMenuService = NewsBaseSysMenuService()
		admin              = common.GetAdmin(ctx)
	)

	res = &permmenu{
		Perms: baseSysMenuService.GetPerms(roleIds),
		Menus: baseSysMenuService.GetMenus(admin.RoleIds, gstr.Equal(admin.UserId, "1")),
	}

	return

}

// RefreshPerms refreshPerms(userId)
func (c *sBaseSysPermsService) RefreshPerms(ctx context.Context, userId string) (err error) {
	var (
		roleIds = service.BaseSysRoleService().GetByUser(userId)
		perms   = service.BaseSysMenuService().GetPerms(roleIds)
	)
	dzhcore.CacheManager.Set(ctx, "admin:perms:"+gconv.String(userId), perms, 0)
	// 更新部门权限
	departments := service.BaseSysDepartmentService().GetByRoleIds(roleIds, gstr.Equal(userId, "1"))
	dzhcore.CacheManager.Set(ctx, "admin:department:"+gconv.String(userId), departments, 0)

	return
}
