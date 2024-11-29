package code

const (
	LoginIllegalOptions       = "login.illegal_options"
	LoginUserNameVerifyError  = "login.user_name_verify_error"
	LoginTokenVerifyFail      = "login.token_verify_fail"
	LoginWrongEmailAddress    = "login.wrong_email_address"
	LoginWrongPhoneNumber     = "login.wrong_phone_number"
	LoginWrongPhoneCode       = "login.wrong_phone_code"
	LoginWrongPhoneCodeUsed   = "login.wrong_phone_code_used"
	LoginPhoneAlreadyRegister = "login.phone_already_register"
	LoginWrongEmail           = "login.wrong_email"
	LoginWrongEmailCode       = "login.wrong_email_code"
	LoginWrongEmailCodeUsed   = "login.wrong_email_code_used"
	LoginWrongEmailInfo       = "login.wrong_email_info"
	LoginWrongPhone           = "login.wrong_phone"
	LoginAgeMinError          = "login.age_min_error"
	LoginErrorNotExistEmail   = "login.error_not_exist_email"
	LoginErrorForbidden       = "login.error_forbidden"
	LoginErrorLogoff          = "login.error_logoff"
	LoginErrorPassword        = "login.error_password"
	LoginErrorNotExistMobile  = "login.error_not_exist_mobile"
	LoginEmailAlreadyRegister = "login.email_already_register"
	LoginCityListError        = "login.city_list_error"
	LoginResetPasswordError   = "login.reset_password_error"
	LoginPasswordFormatError  = "login.password_format_error"
	LoginAlreadyUserLogoff    = "login.already_user_logoff"
)

var loginMap = map[string]int{
	LoginIllegalOptions:       1,
	LoginUserNameVerifyError:  2,
	LoginTokenVerifyFail:      3,
	LoginWrongEmailAddress:    4,
	LoginWrongPhoneNumber:     5,
	LoginWrongPhoneCode:       6,
	LoginWrongPhoneCodeUsed:   7,
	LoginPhoneAlreadyRegister: 8,
	LoginWrongEmail:           9,
	LoginWrongEmailCode:       10,
	LoginWrongEmailCodeUsed:   11,
	LoginWrongEmailInfo:       12,
	LoginWrongPhone:           13,
	LoginAgeMinError:          14,
	LoginErrorNotExistEmail:   15,
	LoginErrorForbidden:       16,
	LoginErrorLogoff:          17,
	LoginErrorPassword:        18,
	LoginErrorNotExistMobile:  19,
	LoginEmailAlreadyRegister: 20,
	LoginCityListError:        21,
	LoginResetPasswordError:   22,
	LoginPasswordFormatError:  23,
	LoginAlreadyUserLogoff:    24,
}
