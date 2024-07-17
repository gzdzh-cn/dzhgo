// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseSysMemberAttr is the golang structure for table base_sys_member_attr.
type BaseSysMemberAttr struct {
	Id              string      `json:"id"              orm:"id"              ` //
	CreateTime      *gtime.Time `json:"createTime"      orm:"createTime"      ` // 创建时间
	UpdateTime      *gtime.Time `json:"updateTime"      orm:"updateTime"      ` // 更新时间
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"      ` //
	UserId          string      `json:"userId"          orm:"userId"          ` // 会员
	FollowJson      string      `json:"followJson"      orm:"followJson"      ` // 关注列表
	LikeJson        string      `json:"likeJson"        orm:"likeJson"        ` // 点赞列表
	CollectJson     string      `json:"collectJson"     orm:"collectJson"     ` // 收藏列表
	CommentLikeJson string      `json:"commentLikeJson" orm:"commentLikeJson" ` // 评论点赞列表
	NoticeJson      string      `json:"noticeJson"      orm:"noticeJson"      ` // 通知阅读痕迹
	NoticeOpen      int         `json:"noticeOpen"      orm:"noticeOpen"      ` // 通知开启
}
