// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"dzhgo/internal/dao/internal"
)

// internalBaseEpsAppDao is internal type for wrapping internal DAO implements.
type internalBaseEpsAppDao = *internal.BaseEpsAppDao

// baseEpsAppDao is the data access object for table base_eps_app.
// You can define custom methods on it to extend its functionality as you wish.
type baseEpsAppDao struct {
	internalBaseEpsAppDao
}

var (
	// BaseEpsApp is globally public accessible object for table base_eps_app operations.
	BaseEpsApp = baseEpsAppDao{
		internal.NewBaseEpsAppDao(),
	}
)

// Fill with you ideas below.
