package code

const (
	WithdrawProcessTypeNotFound     = "withdraw.process.not_found"
	WithdrawParamBankNameInvalid    = "withdraw.params.bank_name_invalid"
	WithdrawParamBankBranchInvalid  = "withdraw.params.bank_branch_invalid"
	WithdrawParamUserAccountInvalid = "withdraw.params.user_account_invalid"
	WithdrawParamUserNameInvalid    = "withdraw.params.user_name_invalid"
	WithdrawParamFirstNameInvalid   = "withdraw.params.first_name_invalid"
	WithdrawParamLastNameInvalid    = "withdraw.params.last_name_invalid"
	WithdrawParamSwiftCodeInvalid   = "withdraw.params.swift_invalid"
	WithdrawParamProvinceInvalid    = "withdraw.params.province_invalid"
	WithdrawParamCityInvalid        = "withdraw.params.city_invalid"
	WithdrawParamAddressInvalid     = "withdraw.params.address_invalid"
	WithdrawParamPostCodeInvalid    = "withdraw.params.post_code_invalid"
	WithdrawParamPhoneInvalid       = "withdraw.params.phone_invalid"
	WithdrawParamAccountTypeInvalid = "withdraw.params.account_type_invalid"
	WithdrawApplyMinInvalid         = "withdraw.apply.min_invalid"
	WithdrawApplyMaxInvalid         = "withdraw.apply.max_invalid"
	WithdrawApplyBalanceInvalid     = "withdraw.apply.balance_invalid"
	WithdrawApplyAccountInvalid     = "withdraw.apply.account_invalid"
)

var withdrawMap = map[string]int{
	WithdrawProcessTypeNotFound:     1,
	WithdrawParamBankNameInvalid:    2,
	WithdrawParamBankBranchInvalid:  2,
	WithdrawParamUserAccountInvalid: 2,
	WithdrawParamUserNameInvalid:    2,
	WithdrawParamFirstNameInvalid:   2,
	WithdrawParamLastNameInvalid:    2,
	WithdrawParamSwiftCodeInvalid:   2,
	WithdrawParamProvinceInvalid:    2,
	WithdrawParamCityInvalid:        2,
	WithdrawParamAddressInvalid:     2,
	WithdrawParamPostCodeInvalid:    2,
	WithdrawParamPhoneInvalid:       2,
	WithdrawParamAccountTypeInvalid: 2,
	WithdrawApplyMinInvalid:         3,
	WithdrawApplyMaxInvalid:         3,
	WithdrawApplyBalanceInvalid:     3,
	WithdrawApplyAccountInvalid:     3,
}
