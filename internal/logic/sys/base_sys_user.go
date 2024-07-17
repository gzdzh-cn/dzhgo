package sys

import (
	"context"
	"dzhgo/internal/common"
	"dzhgo/internal/dao"
	"dzhgo/internal/service"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/text/gstr"

	"dzhgo/internal/model"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gzdzh/dzhcore"
)

func init() {
	service.RegisterBaseSysUserService(NewsBaseSysUserService())
}

type sBaseSysUserService struct {
	*dzhcore.Service
}

// NewsBaseSysUserService 创建一个新的sBaseSysUserService实例
func NewsBaseSysUserService() *sBaseSysUserService {
	return &sBaseSysUserService{
		Service: &dzhcore.Service{
			Dao:                &dao.BaseSysUser,
			Model:              model.NewBaseSysUser(),
			InfoIgnoreProperty: "password",
			UniqueKey: map[string]string{
				"username": "用户名不能重复",
			},
			PageQueryOp: &dzhcore.QueryOp{
				Select: "base_sys_user.*,dept.`name` as departmentName,GROUP_CONCAT( role.`name` ) AS `roleName`",
				Join: []*dzhcore.JoinOp{
					{
						Model:     model.NewBaseSysDepartment(),
						Alias:     "dept",
						Type:      "LeftJoin",
						Condition: "`base_sys_user`.`departmentId` = `dept`.`id`",
					},
					{
						Model:     model.NewBaseSysUserRole(),
						Alias:     "user_role",
						Type:      "LeftJoin",
						Condition: "`base_sys_user`.`id` = `user_role`.`userId`",
					},
					{
						Model:     model.NewBaseSysRole(),
						Alias:     "`role`",
						Type:      "LeftJoin",
						Condition: "`role`.`id` = `user_role`.`roleId`",
					},
				},
				Where: func(ctx context.Context) []g.Array {

					r := g.RequestFromCtx(ctx).GetMap()
					admin := common.GetAdmin(ctx)

					var condition = []g.Array{{"(departmentId IN (?))", gconv.SliceStr(r["departmentIds"])}}

					if !gstr.Equal(admin.UserId, "1") {

						condition = append(condition, g.Slice{"id !=?", 1})
					}

					return condition
				},
				Extend: func(ctx g.Ctx, m *gdb.Model) *gdb.Model {
					return m.Group("`base_sys_user`.`id`")
				},
				KeyWordField: []string{"name", "username", "nickName"},
			},
		},
	}
}

// Person 方法 返回不带密码的用户信息
func (s *sBaseSysUserService) Person(userId string) (res gdb.Record, err error) {
	m := s.Dao.Ctx(ctx)
	res, err = m.Where("id = ?", userId).FieldsEx("password").One()
	return
}

func (s *sBaseSysUserService) ModifyBefore(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Delete" {
		// 禁止删除超级管理员
		userIds := garray.NewStrArrayFrom(gconv.Strings(param["ids"]))
		currentId, found := userIds.Get(0)
		superAdminId := "1"

		if userIds.Len() == 1 && found && gstr.Equal(currentId, superAdminId) {
			err = gerror.New("超级管理员不能删除")
			return
		}

		// 删除超级管理员
		userIds.RemoveValue("1")
		g.RequestFromCtx(ctx).SetParam("ids", userIds.Slice())
	}
	return
}

func (s *sBaseSysUserService) ModifyAfter(ctx context.Context, method string, param g.MapStrAny) (err error) {
	if method == "Delete" {
		userIds := garray.NewIntArrayFrom(gconv.Ints(param["ids"]))
		userIds.RemoveValue(1)
		// 删除用户时删除相关数据
		dao.BaseSysUserRole.Ctx(ctx).WhereIn("userId", userIds.Slice()).Delete()

	}
	return
}

// ServiceAdd 方法 添加用户
func (s *sBaseSysUserService) ServiceAdd(ctx context.Context, req *dzhcore.AddReq) (data interface{}, err error) {
	var (
		m      = s.Dao.Ctx(ctx)
		r      = g.RequestFromCtx(ctx)
		reqmap = r.GetMap()
	)

	// 创建雪花算法节点
	node, err := snowflake.NewNode(1) // 1 是节点的ID
	if err != nil {
		g.Log().Error(ctx, err)
	}

	// 如果reqmap["password"]不为空，则对密码进行md5加密
	if !r.Get("password").IsNil() {
		reqmap["password"] = gmd5.MustEncryptString(r.Get("password").String())
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		reqmap["id"] = node.Generate()
		_, err = m.TX(tx).Data(reqmap).Insert()
		if err != nil {
			return err
		}

		// 如果请求参数中不包含roleIdList说明不修改角色信息
		if !r.Get("roleIdList").IsNil() {
			roleModel := dao.BaseSysUserRole.Ctx(ctx).TX(tx)
			roleArray := garray.NewArray()
			inRoleIdSet := gset.NewFrom(r.Get("roleIdList").Ints())
			inRoleIdSet.Iterator(func(v interface{}) bool {
				roleArray.PushRight(g.Map{
					"id":     node.Generate(),
					"userId": gconv.Uint(reqmap["id"]),
					"roleId": gconv.Uint(v),
				})
				return true
			})

			_, err = roleModel.Fields("id,userId,roleId").Insert(roleArray)
			if err != nil {
				return err
			}

		}

		data = g.Map{"id": reqmap["id"]}
		return
	})

	return
}

// ServiceInfo 方法 返回服务信息
func (s *sBaseSysUserService) ServiceInfo(ctx g.Ctx, req *dzhcore.InfoReq) (data interface{}, err error) {
	result, err := s.Service.ServiceInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if result.(gdb.Record).IsEmpty() {
		return nil, nil
	}

	resultMap := result.(gdb.Record).Map()

	// 获取角色
	userRoles := dao.BaseSysUserRole.Ctx(ctx)
	roleIds, err := userRoles.Where("userId = ?", resultMap["id"]).Fields("roleId").Array()
	if err != nil {
		return nil, err
	}

	resultMap["roleIdList"] = roleIds
	data = resultMap

	return
}

// ServiceUpdate 方法 更新用户信息
func (s *sBaseSysUserService) ServiceUpdate(ctx context.Context, req *dzhcore.UpdateReq) (data interface{}, err error) {
	var (
		admin = common.GetAdmin(ctx)
		m     = s.Dao.Ctx(ctx)
	)

	r := g.RequestFromCtx(ctx)
	rMap := r.GetMap()

	// 如果不传入ID代表更新当前用户
	userId := r.Get("id", admin.UserId).Uint()
	userInfo, err := m.Where("id = ?", userId).One()

	if err != nil {
		return
	}
	if userInfo.IsEmpty() {
		err = gerror.New("用户不存在")
		return
	}

	// 禁止禁用超级管理员
	if userId == 1 && (!r.Get("status").IsNil() && r.Get("status").Int() == 0) {
		err = gerror.New("禁止禁用超级管理员")
		return
	}

	// 如果请求的password不为空并且密码加密后的值有变动，说明要修改密码
	var rPassword = r.Get("password", "").String()
	if rPassword != "" && rPassword != userInfo["password"].String() {
		rMap["password"], _ = gmd5.Encrypt(rPassword)
		rMap["passwordV"] = userInfo["passwordV"].Int() + 1
		dzhcore.CacheManager.Set(ctx, fmt.Sprintf("admin:passwordVersion:%d", userId), rMap["passwordV"], 0)
	} else {
		delete(rMap, "password")
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		roleModel := dao.BaseSysUserRole.Ctx(ctx).TX(tx).Where("userId = ?", userId)
		roleIds, err := roleModel.Fields("roleId").Array()
		if err != nil {
			return
		}

		// 如果请求参数中不包含roleIdList说明不修改角色信息
		if !r.Get("roleIdList").IsNil() {
			inRoleIdSet := gset.NewFrom(r.Get("roleIdList").Strings())
			roleIdsSet := gset.NewFrom(gconv.Strings(roleIds))

			// 如果请求的角色信息未发生变化则跳过更新逻辑
			if roleIdsSet.Diff(inRoleIdSet).Size() != 0 || inRoleIdSet.Diff(roleIdsSet).Size() != 0 {
				// 创建雪花算法节点
				node, err := snowflake.NewNode(1) // 1 是节点的ID
				if err != nil {
					g.Log().Error(ctx, err)
				}

				roleArray := garray.NewArray()
				inRoleIdSet.Iterator(func(v interface{}) bool {
					roleArray.PushRight(g.Map{
						"id":     node.Generate(),
						"userId": gconv.String(userId),
						"roleId": gconv.String(v),
					})
					return true
				})

				_, err = roleModel.Delete()
				if err != nil {
					return err
				}

				_, err = roleModel.Fields("id,userId,roleId").Insert(roleArray)
				if err != nil {
					return err
				}
			}
		}

		_, err = m.TX(tx).Update(rMap)

		if err != nil {
			return err
		}
		return
	})
	return
}

// Move 移动用户部门
func (s *sBaseSysUserService) Move(ctx g.Ctx) (err error) {
	request := g.RequestFromCtx(ctx)
	departmentId := request.Get("departmentId").String()
	userIds := request.Get("userIds").Slice()

	_, err = s.Dao.Ctx(ctx).Where("`id` IN(?)", userIds).Data(g.Map{"departmentId": departmentId}).Update()

	return
}
