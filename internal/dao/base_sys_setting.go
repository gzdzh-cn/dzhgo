// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"dzhgo/internal/dao/internal"
)

// internalBaseSysSettingDao is internal type for wrapping internal DAO implements.
type internalBaseSysSettingDao = *internal.BaseSysSettingDao

// baseSysSettingDao is the data access object for table base_sys_setting.
// You can define custom methods on it to extend its functionality as you wish.
type baseSysSettingDao struct {
	internalBaseSysSettingDao
}

var (
	// BaseSysSetting is globally public accessible object for table base_sys_setting operations.
	BaseSysSetting = baseSysSettingDao{
		internal.NewBaseSysSettingDao(),
	}
)

// Fill with you ideas below.
