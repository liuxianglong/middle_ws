// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CmsAuthSso is the golang structure for table cms_auth_sso.
type CmsAuthSso struct {
	Id          uint   `json:"id"           orm:"id"           ` // 自增ID
	AppId       string `json:"app_id"       orm:"app_id"       ` // app的唯一标识符
	AppSecret   string `json:"app_secret"   orm:"app_secret"   ` // app的加密key
	Name        string `json:"name"         orm:"name"         ` // 系统名称
	CallbackUrl string `json:"callback_url" orm:"callback_url" ` // 回调
	CreateAt    int64  `json:"create_at"    orm:"create_at"    ` // 创建时间
	UpdateAt    int64  `json:"update_at"    orm:"update_at"    ` // 修改时间
}
