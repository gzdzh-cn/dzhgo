package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysMember = "base_sys_member"

// BaseSysMember mapped from table <member_open>
type BaseSysMember struct {
	*dzhcore.Model
	AvatarUrl   string  `gorm:"column:avatarUrl;comment:头像;type:varchar(200)" json:"avatarUrl"`
	Password    string  `gorm:"column:password;not null;comment:会员密码;type:varchar(50)" json:"password"`
	PasswordV   *int32  `gorm:"column:passwordV;not null;type:int;default:1" json:"passwordV"` // 密码版本, 作用是改完密码，让原来的token失效
	UserName    string  `gorm:"column:username;comment:会员账号;type:varchar(50);index" json:"username"`
	Nickname    string  `gorm:"column:nickname;comment:会员昵称;type:varchar(50);index" json:"nickname"`
	LevelName   string  `gorm:"column:levelName;comment:等级名称;type:varchar(50);default:普通会员" json:"levelName"`
	Level       int     `gorm:"column:level;comment:等级;type:int;default:1" json:"level"`
	Sex         *int32  `gorm:"column:sex;comment:性别;type:int;default:1" json:"sex"`
	QQ          *string `gorm:"column:qq;comment:qq;type:varchar(255);index" json:"qq"`
	Mobile      *string `gorm:"column:mobile;comment:手机号;type:varchar(50);index" json:"mobile"`
	Wx          *string `gorm:"column:wx;comment:微信号;type:varchar(50);index" json:"wx"`
	WxImg       *string `gorm:"column:wxImg;comment:微信二维码;type:varchar(255)" json:"wxImg"`
	Email       *string `gorm:"column:email;comment:email;type:varchar(50);index" json:"email"`
	Role        string  `gorm:"column:role;comment:家庭角色;" json:"role"`
	Remark      *string `gorm:"column:remark;comment:备注;type:varchar(255)" json:"remark"`
	Openid      string  `gorm:"column:openid;comment:openid;" json:"openid"`
	UnionId     string  `gorm:"column:unionId;comment:unionId;" json:"unionId"`
	SessionKey  string  `gorm:"column:session_key;comment:session_key;" json:"session_key"`
	Status      *int32  `gorm:"column:status;not null;type:int;default:1" json:"status"` // 状态 0:禁用 1：启用
	Description *string `gorm:"column:description;comment:描述;type:varchar(100);default:写一个霸气侧漏的签名" json:"description"`
}

// TableName BaseMember's table name
func (*BaseSysMember) TableName() string {
	return TableNameBaseSysMember
}

// GroupName BaseMember's table group
func (*BaseSysMember) GroupName() string {
	return "default"
}

// NewBaseMember create a new BaseMember
func NewBaseMember() *BaseSysMember {
	return &BaseSysMember{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	err := dzhcore.CreateTable(&BaseSysMember{})
	if err != nil {
		return
	}
}
