// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"dzhgo/internal/dao/internal"
)

// internalBaseSysMenuDao is internal type for wrapping internal DAO implements.
type internalBaseSysMenuDao = *internal.BaseSysMenuDao

// baseSysMenuDao is the data access object for table base_sys_menu.
// You can define custom methods on it to extend its functionality as you wish.
type baseSysMenuDao struct {
	internalBaseSysMenuDao
}

var (
	// BaseSysMenu is globally public accessible object for table base_sys_menu operations.
	BaseSysMenu = baseSysMenuDao{
		internal.NewBaseSysMenuDao(),
	}
)

// Fill with you ideas below.
