package model

type FeiShuCommon struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type FeiShuCfg struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	RedirectURL  string `json:"redirectURL"`
}

type FeiShuUserInfo struct {
	FeiShuCommon
	Data *FeiShuUserInfoDetail `json:"data"`
}

type FeiShuUserInfoDetail struct {
	Name            string `json:"name"`
	EnName          string `json:"en_name"`
	AvatarUrl       string `json:"avatar_url"`
	AvatarThumb     string `json:"avatar_thumb"`
	AvatarMiddle    string `json:"avatar_middle"`
	AvatarBig       string `json:"avatar_big"`
	OpenId          string `json:"open_id"`
	UnionId         string `json:"union_id"`
	Email           string `json:"email"`
	EnterpriseEmail string `json:"enterprise_email"`
	UserId          string `json:"user_id"`
	Mobile          string `json:"mobile"`
	TenantKey       string `json:"tenant_key"`
	EmployeeNo      string `json:"employee_no"`
}
