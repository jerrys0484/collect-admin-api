// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StepsDao is the data access object for the table ct_steps.
type StepsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  StepsColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// StepsColumns defines and stores column names for the table ct_steps.
type StepsColumns struct {
	Id         string // ID
	Uuid       string // 标识
	Name       string // 名称
	Type       string // 类型
	Template   string // 模板标识
	Request    string // 请求参数(json)
	Response   string // 响应参数配置(json)
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// stepsColumns holds the columns for the table ct_steps.
var stepsColumns = StepsColumns{
	Id:         "id",
	Uuid:       "uuid",
	Name:       "name",
	Type:       "type",
	Template:   "template",
	Request:    "request",
	Response:   "response",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewStepsDao creates and returns a new DAO object for table data access.
func NewStepsDao(handlers ...gdb.ModelHandler) *StepsDao {
	return &StepsDao{
		group:    "default",
		table:    "ct_steps",
		columns:  stepsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StepsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StepsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StepsDao) Columns() StepsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StepsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StepsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *StepsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
