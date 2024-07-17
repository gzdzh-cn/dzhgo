package common

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

type Admin struct {
	IsRefresh       bool     `json:"isRefresh"`
	RoleIds         []string `json:"roleIds"`
	Username        string   `json:"username"`
	UserId          string   `json:"userId"`
	PasswordVersion *int32   `json:"passwordVersion"`
}

// 获取传入ctx 中的 admin 对象
func GetAdmin(ctx context.Context) *Admin {
	r := g.RequestFromCtx(ctx)
	admin := &Admin{}
	err := gjson.New(r.GetCtxVar("admin").String()).Scan(admin)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return admin
}

type Member struct {
	IsRefresh       bool   `json:"isRefresh"`
	RoleIds         []uint `json:"roleIds"`
	Username        string `json:"username"`
	NickName        string `json:"nickName"`
	LevelName       string `json:"levelName"`
	UserId          uint   `json:"userId"`
	PasswordVersion *int32 `json:"passwordVersion"`
}

// GetMember 获取传入ctx 中的 member 对象
func GetMember(ctx context.Context) *Member {
	r := g.RequestFromCtx(ctx)
	member := &Member{}
	err := gjson.New(r.GetCtxVar("member").String()).Scan(member)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return member
}
