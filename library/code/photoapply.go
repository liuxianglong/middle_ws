package code

const (
	PhotoapplyHasApply     = "photoapply.has_apply"
	PhotoapplyNeedIdVerify = "photoapply.need_id_verify"
)

var photoapplyMap = map[string]int{
	PhotoapplyHasApply:     1,
	PhotoapplyNeedIdVerify: 2,
}
