package code

const (
	CommonChangeStateRuleNotFound = "common.changestate.rule_not_found"
	CommonChangeStateNoMatch      = "common.changestate.state_cannot_change"
	CommonAuthError               = "common.auth_error"
	CommonUploadFileEmpty         = "common.upload_file_empty"
	CommonUploadFileMimeNomatch   = "common.upload_file_mime_nomatch"
	CommonUploadFileFail          = "common.upload_file_fail"
	CommonDataNoExist             = "common.data_no_exist"
	CommonNoLocated               = "common.no_located"
	CommonSaveFail                = "common.save_fail"
	CommonResourceLimit           = "common.resource_only_allow_limit"
	CommonUidInvalid              = "common.Uid_InValid"
	CommonOperationFail           = "common.operation_fail"
	CommonUpgradeVip              = "common.upgrade_vip"
	CommonOperationQuick          = "common.operation_quick"
	CommonDecryptError            = "common.decrypt_error"
	CommonAdminError              = "common.admin_error" //后台自定义的，非后台的禁用
	CommonImgLarge                = "common.image_large"
	CommonLengthLimit             = "common.length_limit"
	CommonColumnLengthLimit       = "common.column_length_limit"
)

var commonMap = map[string]int{
	CommonChangeStateRuleNotFound: 1,
	CommonChangeStateNoMatch:      2,
	CommonAuthError:               3,
	CommonUploadFileEmpty:         4,
	CommonUploadFileMimeNomatch:   5,
	CommonUploadFileFail:          6,
	CommonDataNoExist:             7,
	CommonNoLocated:               8,
	CommonSaveFail:                9,
	CommonResourceLimit:           10,
	CommonUidInvalid:              11,
	CommonOperationFail:           12,
	CommonUpgradeVip:              13,
	CommonDecryptError:            14,
	CommonAdminError:              15,
	CommonImgLarge:                16,
	CommonLengthLimit:             17,
	CommonColumnLengthLimit:       18,
}
