package sys

import (
	"context"
	"dzhgo/internal/dao"
	"dzhgo/internal/model"
	"dzhgo/internal/service"
	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh-cn/dzhcore"
)

func init() {
	service.RegisterBaseSysConfService(NewsBaseSysConfService())
}

type sBaseSysConfService struct {
	*dzhcore.Service
}

func NewsBaseSysConfService() *sBaseSysConfService {
	return &sBaseSysConfService{
		&dzhcore.Service{
			Dao:   &dao.BaseSysConf,
			Model: model.NewBaseSysConf(),
			UniqueKey: map[string]string{
				"cKey": "配置键不能重复",
			},
		},
	}
}

// UpdateValue 更新配置值
func (s *sBaseSysConfService) UpdateValue(cKey, cValue string) error {
	m := s.Dao.Ctx(ctx).Where("cKey = ?", cKey)
	record, err := m.One()
	if err != nil {
		return err
	}
	// 创建雪花算法节点
	node, err := snowflake.NewNode(1) // 1 是节点的ID
	if err != nil {
		g.Log().Error(context.Background(), err)
	}
	if record == nil {
		_, err = s.Dao.Ctx(ctx).Insert(g.Map{
			"id":     node.Generate(),
			"cKey":   cKey,
			"cValue": cValue,
		})
	} else {
		_, err = s.Dao.Ctx(ctx).Where("cKey = ?", cKey).Data(g.Map{"cValue": cValue}).Update()
	}
	return err
}

// GetValue 获取配置值
func (s *sBaseSysConfService) GetValue(cKey string) string {
	m := s.Dao.Ctx(ctx).Where("cKey = ?", cKey)
	record, err := m.One()
	if err != nil {
		return ""
	}
	if record == nil {
		return ""
	}
	return record["cValue"].String()
}
