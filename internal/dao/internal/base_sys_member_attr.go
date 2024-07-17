// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BaseSysMemberAttrDao is the data access object for table base_sys_member_attr.
type BaseSysMemberAttrDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns BaseSysMemberAttrColumns // columns contains all the column names of Table for convenient usage.
}

// BaseSysMemberAttrColumns defines and stores column names for table base_sys_member_attr.
type BaseSysMemberAttrColumns struct {
	Id              string //
	CreateTime      string // 创建时间
	UpdateTime      string // 更新时间
	DeletedAt       string //
	UserId          string // 会员
	FollowJson      string // 关注列表
	LikeJson        string // 点赞列表
	CollectJson     string // 收藏列表
	CommentLikeJson string // 评论点赞列表
	NoticeJson      string // 通知阅读痕迹
	NoticeOpen      string // 通知开启
}

// baseSysMemberAttrColumns holds the columns for table base_sys_member_attr.
var baseSysMemberAttrColumns = BaseSysMemberAttrColumns{
	Id:              "id",
	CreateTime:      "createTime",
	UpdateTime:      "updateTime",
	DeletedAt:       "deleted_at",
	UserId:          "userId",
	FollowJson:      "followJson",
	LikeJson:        "likeJson",
	CollectJson:     "collectJson",
	CommentLikeJson: "commentLikeJson",
	NoticeJson:      "noticeJson",
	NoticeOpen:      "noticeOpen",
}

// NewBaseSysMemberAttrDao creates and returns a new DAO object for table data access.
func NewBaseSysMemberAttrDao() *BaseSysMemberAttrDao {
	return &BaseSysMemberAttrDao{
		group:   "default",
		table:   "base_sys_member_attr",
		columns: baseSysMemberAttrColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BaseSysMemberAttrDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BaseSysMemberAttrDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BaseSysMemberAttrDao) Columns() BaseSysMemberAttrColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BaseSysMemberAttrDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BaseSysMemberAttrDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BaseSysMemberAttrDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}