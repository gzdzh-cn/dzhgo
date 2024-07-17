// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseEpsAdminDao is the data access object for table base_eps_admin.
type BaseEpsAdminDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns BaseEpsAdminColumns // columns contains all the column names of Table for convenient usage.
}

// BaseEpsAdminColumns defines and stores column names for table base_eps_admin.
type BaseEpsAdminColumns struct {
	Id      string //
	Module  string //
	Method  string //
	Path    string //
	Prefix  string //
	Summary string //
	Tag     string //
	Dts     string //
}

// baseEpsAdminColumns holds the columns for table base_eps_admin.
var baseEpsAdminColumns = BaseEpsAdminColumns{
	Id:      "id",
	Module:  "module",
	Method:  "method",
	Path:    "path",
	Prefix:  "prefix",
	Summary: "summary",
	Tag:     "tag",
	Dts:     "dts",
}

// NewBaseEpsAdminDao creates and returns a new DAO object for table data access.
func NewBaseEpsAdminDao() *BaseEpsAdminDao {
	return &BaseEpsAdminDao{
		group:   "default",
		table:   "base_eps_admin",
		columns: baseEpsAdminColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseEpsAdminDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseEpsAdminDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseEpsAdminDao) Columns() BaseEpsAdminColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseEpsAdminDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseEpsAdminDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseEpsAdminDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
