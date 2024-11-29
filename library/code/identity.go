package code

const (
	IdentitySameImage  = "identity.same_image"
	IdentityHasSubmit  = "identity.has_submit"
	IdentityAppealFail = "identity.appeal_fail"
	IdentityAppealHas  = "identity.appeal_has"
	IdentityHasPass    = "identity.has_pass"
	IdentityDoingAudit = "identity.doing_audit"
	IdentityLimitNum   = "identity.limit_num"
)

var identityMap = map[string]int{
	IdentitySameImage:  1,
	IdentityHasSubmit:  2,
	IdentityAppealFail: 3,
	IdentityAppealHas:  4,
	IdentityHasPass:    5,
	IdentityDoingAudit: 6,
	IdentityLimitNum:   7,
}
