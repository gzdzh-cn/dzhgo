// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysUserRoleDao is the data access object for table base_sys_user_role.
type BaseSysUserRoleDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns BaseSysUserRoleColumns // columns contains all the column names of Table for convenient usage.
}

// BaseSysUserRoleColumns defines and stores column names for table base_sys_user_role.
type BaseSysUserRoleColumns struct {
	Id         string //
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	DeletedAt  string //
	UserId     string //
	RoleId     string //
}

// baseSysUserRoleColumns holds the columns for table base_sys_user_role.
var baseSysUserRoleColumns = BaseSysUserRoleColumns{
	Id:         "id",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
	DeletedAt:  "deleted_at",
	UserId:     "userId",
	RoleId:     "roleId",
}

// NewBaseSysUserRoleDao creates and returns a new DAO object for table data access.
func NewBaseSysUserRoleDao() *BaseSysUserRoleDao {
	return &BaseSysUserRoleDao{
		group:   "default",
		table:   "base_sys_user_role",
		columns: baseSysUserRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseSysUserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseSysUserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseSysUserRoleDao) Columns() BaseSysUserRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseSysUserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseSysUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseSysUserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
