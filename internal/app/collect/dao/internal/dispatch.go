// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DispatchDao is the data access object for the table ct_dispatch.
type DispatchDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DispatchColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DispatchColumns defines and stores column names for the table ct_dispatch.
type DispatchColumns struct {
	Id         string // ID
	Uuid       string // 标识
	Name       string // 名称
	Type       string // 类型
	Way        string // 方式
	Rules      string // 规则(json)
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// dispatchColumns holds the columns for the table ct_dispatch.
var dispatchColumns = DispatchColumns{
	Id:         "id",
	Uuid:       "uuid",
	Name:       "name",
	Type:       "type",
	Way:        "way",
	Rules:      "rules",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewDispatchDao creates and returns a new DAO object for table data access.
func NewDispatchDao(handlers ...gdb.ModelHandler) *DispatchDao {
	return &DispatchDao{
		group:    "default",
		table:    "ct_dispatch",
		columns:  dispatchColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DispatchDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DispatchDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DispatchDao) Columns() DispatchColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DispatchDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DispatchDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DispatchDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
