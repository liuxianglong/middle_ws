// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CmsAuthSsoDao is the data access object for the table cms_auth_sso.
type CmsAuthSsoDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  CmsAuthSsoColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// CmsAuthSsoColumns defines and stores column names for the table cms_auth_sso.
type CmsAuthSsoColumns struct {
	Id          string // 自增ID
	AppId       string // app的唯一标识符
	AppSecret   string // app的加密key
	Name        string // 系统名称
	CallbackUrl string // 回调
	CreateAt    string // 创建时间
	UpdateAt    string // 修改时间
}

// cmsAuthSsoColumns holds the columns for the table cms_auth_sso.
var cmsAuthSsoColumns = CmsAuthSsoColumns{
	Id:          "id",
	AppId:       "app_id",
	AppSecret:   "app_secret",
	Name:        "name",
	CallbackUrl: "callback_url",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
}

// NewCmsAuthSsoDao creates and returns a new DAO object for table data access.
func NewCmsAuthSsoDao(handlers ...gdb.ModelHandler) *CmsAuthSsoDao {
	return &CmsAuthSsoDao{
		group:    "default",
		table:    "cms_auth_sso",
		columns:  cmsAuthSsoColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CmsAuthSsoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CmsAuthSsoDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CmsAuthSsoDao) Columns() CmsAuthSsoColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CmsAuthSsoDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CmsAuthSsoDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CmsAuthSsoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
