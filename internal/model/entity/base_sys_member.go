// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysMember is the golang structure for table base_sys_member.
type BaseSysMember struct {
	Id          string      `json:"id"          orm:"id"          ` //
	CreateTime  *gtime.Time `json:"createTime"  orm:"createTime"  ` // 创建时间
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"updateTime"  ` // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  ` //
	AvatarUrl   string      `json:"avatarUrl"   orm:"avatarUrl"   ` // 头像
	Password    string      `json:"password"    orm:"password"    ` // 会员密码
	PasswordV   int         `json:"passwordV"   orm:"passwordV"   ` //
	Username    string      `json:"username"    orm:"username"    ` // 会员账号
	Nickname    string      `json:"nickname"    orm:"nickname"    ` // 会员昵称
	LevelName   string      `json:"levelName"   orm:"levelName"   ` // 等级名称
	Level       int64       `json:"level"       orm:"level"       ` // 等级
	Sex         int         `json:"sex"         orm:"sex"         ` // 性别
	Qq          string      `json:"qq"          orm:"qq"          ` // qq
	Mobile      string      `json:"mobile"      orm:"mobile"      ` // 手机号
	Wx          string      `json:"wx"          orm:"wx"          ` // 微信号
	WxImg       string      `json:"wxImg"       orm:"wxImg"       ` // 微信二维码
	Email       string      `json:"email"       orm:"email"       ` // email
	Role        string      `json:"role"        orm:"role"        ` // 家庭角色
	Remark      string      `json:"remark"      orm:"remark"      ` // 备注
	Openid      string      `json:"openid"      orm:"openid"      ` // openid
	UnionId     string      `json:"unionId"     orm:"unionId"     ` // unionId
	SessionKey  string      `json:"sessionKey"  orm:"session_key" ` // session_key
	Status      int         `json:"status"      orm:"status"      ` //
	Description string      `json:"description" orm:"description" ` // 描述
}
