// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"dzhgo/internal/dao/internal"
)

// internalBaseSysDepartmentDao is internal type for wrapping internal DAO implements.
type internalBaseSysDepartmentDao = *internal.BaseSysDepartmentDao

// baseSysDepartmentDao is the data access object for table base_sys_department.
// You can define custom methods on it to extend its functionality as you wish.
type baseSysDepartmentDao struct {
	internalBaseSysDepartmentDao
}

var (
	// BaseSysDepartment is globally public accessible object for table base_sys_department operations.
	BaseSysDepartment = baseSysDepartmentDao{
		internal.NewBaseSysDepartmentDao(),
	}
)

// Fill with you ideas below.
