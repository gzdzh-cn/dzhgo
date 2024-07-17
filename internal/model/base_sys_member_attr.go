package model

import (
	"github.com/gzdzh-cn/dzhcore"
)

const TableNameBaseSysMemberAttr = "base_sys_member_attr"

// BaseSysMemberAttr mapped from table <member_open>
type BaseSysMemberAttr struct {
	*dzhcore.Model
	UserId          *string `gorm:"column:userId;not null;comment:会员" json:"userId"`
	FollowJson      string  `gorm:"column:followJson;comment:关注列表" json:"followJson"`
	LikeJson        string  `gorm:"column:likeJson;comment:点赞列表" json:"likeJson"`
	CollectJson     string  `gorm:"column:collectJson;comment:收藏列表" json:"collectJson"`
	CommentLikeJson string  `gorm:"column:commentLikeJson;comment:评论点赞列表;" json:"commentLikeJson"`
	NoticeJson      string  `gorm:"column:noticeJson;comment:通知阅读痕迹;" json:"noticeJson"`
	NoticeOpen      bool    `gorm:"column:noticeOpen;comment:通知开启;" json:"noticeOpen"`
}

// TableName BaseMemberAttr's table name
func (*BaseSysMemberAttr) TableName() string {
	return TableNameBaseSysMemberAttr
}

// GroupName BaseMemberAttr's table group
func (*BaseSysMemberAttr) GroupName() string {
	return "default"
}

// NewBaseMemberAttr create a new BaseMemberAttr
func NewBaseMemberAttr() *BaseSysMemberAttr {
	return &BaseSysMemberAttr{
		Model: dzhcore.NewModel(),
	}
}

// init 创建表
func init() {
	err := dzhcore.CreateTable(&BaseSysMemberAttr{})
	if err != nil {
		return
	}
}
