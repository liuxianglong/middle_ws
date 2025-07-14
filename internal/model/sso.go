package model

type SSOAuth struct {
	ClientId    string `dc:"唯一标识符,即系统中的app_id" json:"client_id"`
	State       string `dc:"回调时标识，防止CSRF攻击" json:"state"`
	RedirectUri string `dc:"唯一标识符" json:"redirect_uri"`
}

type SSOCode struct {
	ClientId     string `dc:"唯一标识符,即系统中的app_id" json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Uid          uint   `dc:"用户id" json:"uid"`
}
