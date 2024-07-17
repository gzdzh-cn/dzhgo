// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysMemberAttr is the golang structure of table base_sys_member_attr for DAO operations like Where/Data.
type BaseSysMemberAttr struct {
	g.Meta          `orm:"table:base_sys_member_attr, do:true"`
	Id              interface{} //
	CreateTime      *gtime.Time // 创建时间
	UpdateTime      *gtime.Time // 更新时间
	DeletedAt       *gtime.Time //
	UserId          interface{} // 会员
	FollowJson      interface{} // 关注列表
	LikeJson        interface{} // 点赞列表
	CollectJson     interface{} // 收藏列表
	CommentLikeJson interface{} // 评论点赞列表
	NoticeJson      interface{} // 通知阅读痕迹
	NoticeOpen      interface{} // 通知开启
}
