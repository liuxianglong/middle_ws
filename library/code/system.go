package code

const (
	SystemNeedLogin = "system.need_login"
	SystemServerErr = "system.server_err"
	SystemParamErr  = "system.param_error"
	SystemBusy      = "system.busy"
)

var systemMap = map[string]int{
	SystemNeedLogin: 401,
	SystemServerErr: 500,
	SystemParamErr:  501,
	SystemBusy:      502,
}
