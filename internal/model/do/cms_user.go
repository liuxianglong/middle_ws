// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CmsUser is the golang structure of table cms_user for DAO operations like Where/Data.
type CmsUser struct {
	g.Meta     `orm:"table:cms_user, do:true"`
	Id         interface{} // 自增id
	Name       interface{} // 用户名
	Email      interface{} // 邮箱
	Mobile     interface{} // 电话
	Pwd        interface{} // 密码
	FeishuCode interface{} // 飞书的唯一标识符
	CreateAt   interface{} // 创建时间
	UpdateAt   interface{} // 修改时间
}
