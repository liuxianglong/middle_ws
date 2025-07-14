// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CmsUser is the golang structure for table cms_user.
type CmsUser struct {
	Id         uint   `json:"id"          orm:"id"          ` // 自增id
	Name       string `json:"name"        orm:"name"        ` // 用户名
	Email      string `json:"email"       orm:"email"       ` // 邮箱
	Mobile     string `json:"mobile"      orm:"mobile"      ` // 电话
	Pwd        string `json:"pwd"         orm:"pwd"         ` // 密码
	FeishuCode string `json:"feishu_code" orm:"feishu_code" ` // 飞书的唯一标识符
	CreateAt   int64  `json:"create_at"   orm:"create_at"   ` // 创建时间
	UpdateAt   int64  `json:"update_at"   orm:"update_at"   ` // 修改时间
}
