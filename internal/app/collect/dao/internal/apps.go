// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppsDao is the data access object for the table ct_apps.
type AppsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppsColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppsColumns defines and stores column names for the table ct_apps.
type AppsColumns struct {
	Id         string // ID
	Uuid       string // 标识
	Name       string // 名称
	Template   string // 模板标识
	Rules      string // 调度配置(json)
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// appsColumns holds the columns for the table ct_apps.
var appsColumns = AppsColumns{
	Id:         "id",
	Uuid:       "uuid",
	Name:       "name",
	Template:   "template",
	Rules:      "rules",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewAppsDao creates and returns a new DAO object for table data access.
func NewAppsDao(handlers ...gdb.ModelHandler) *AppsDao {
	return &AppsDao{
		group:    "default",
		table:    "ct_apps",
		columns:  appsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppsDao) Columns() AppsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
