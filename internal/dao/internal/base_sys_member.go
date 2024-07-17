// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysMemberDao is the data access object for table base_sys_member.
type BaseSysMemberDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns BaseSysMemberColumns // columns contains all the column names of Table for convenient usage.
}

// BaseSysMemberColumns defines and stores column names for table base_sys_member.
type BaseSysMemberColumns struct {
	Id          string //
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
	DeletedAt   string //
	AvatarUrl   string // 头像
	Password    string // 会员密码
	PasswordV   string //
	Username    string // 会员账号
	Nickname    string // 会员昵称
	LevelName   string // 等级名称
	Level       string // 等级
	Sex         string // 性别
	Qq          string // qq
	Mobile      string // 手机号
	Wx          string // 微信号
	WxImg       string // 微信二维码
	Email       string // email
	Role        string // 家庭角色
	Remark      string // 备注
	Openid      string // openid
	UnionId     string // unionId
	SessionKey  string // session_key
	Status      string //
	Description string // 描述
}

// baseSysMemberColumns holds the columns for table base_sys_member.
var baseSysMemberColumns = BaseSysMemberColumns{
	Id:          "id",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
	DeletedAt:   "deleted_at",
	AvatarUrl:   "avatarUrl",
	Password:    "password",
	PasswordV:   "passwordV",
	Username:    "username",
	Nickname:    "nickname",
	LevelName:   "levelName",
	Level:       "level",
	Sex:         "sex",
	Qq:          "qq",
	Mobile:      "mobile",
	Wx:          "wx",
	WxImg:       "wxImg",
	Email:       "email",
	Role:        "role",
	Remark:      "remark",
	Openid:      "openid",
	UnionId:     "unionId",
	SessionKey:  "session_key",
	Status:      "status",
	Description: "description",
}

// NewBaseSysMemberDao creates and returns a new DAO object for table data access.
func NewBaseSysMemberDao() *BaseSysMemberDao {
	return &BaseSysMemberDao{
		group:   "default",
		table:   "base_sys_member",
		columns: baseSysMemberColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseSysMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseSysMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseSysMemberDao) Columns() BaseSysMemberColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseSysMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseSysMemberDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseSysMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
