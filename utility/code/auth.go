package code

const (
	AuthAccessDenied         = "auth.access_denied"
	AuthGetThirdTokenFailed  = "auth.get_third_token_failed"
	AuthGetThirdInfoFailed   = "auth.get_third_info_failed"
	AuthSaveSessionFailed    = "auth.save_session_failed"
	AuthThirdCurlExpire      = "auth.third_curl_expire"
	AuthHasLogin             = "auth.has_login"
	AuthGrantTypeError       = "auth.grant_type_error"
	AuthGrantSecretError     = "auth.grant_secret_error"
	AuthBuildTokenError      = "auth.build_token_error"
	AuthMissingAuthorization = "auth.missing_authorization"
	AuthInvalidToken         = "auth.invalid_token"
	AuthUserNoFound          = "auth.user_no_found"
	AuthSSOSecretKeyNoFound  = "auth.sso_secret_key_no_found"
)

var authMap = map[string]int{
	AuthAccessDenied:         1,
	AuthGetThirdTokenFailed:  2,
	AuthGetThirdInfoFailed:   3,
	AuthSaveSessionFailed:    4,
	AuthThirdCurlExpire:      5,
	AuthHasLogin:             6,
	AuthGrantTypeError:       7,
	AuthGrantSecretError:     8,
	AuthBuildTokenError:      9,
	AuthMissingAuthorization: 10,
	AuthInvalidToken:         11,
	AuthUserNoFound:          12,
	AuthSSOSecretKeyNoFound:  13,
}
