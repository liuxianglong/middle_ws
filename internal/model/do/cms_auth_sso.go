// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CmsAuthSso is the golang structure of table cms_auth_sso for DAO operations like Where/Data.
type CmsAuthSso struct {
	g.Meta      `orm:"table:cms_auth_sso, do:true"`
	Id          interface{} // 自增ID
	AppId       interface{} // app的唯一标识符
	AppSecret   interface{} // app的加密key
	Name        interface{} // 系统名称
	CallbackUrl interface{} // 回调
	CreateAt    interface{} // 创建时间
	UpdateAt    interface{} // 修改时间
}
