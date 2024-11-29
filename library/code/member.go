package code

const (
	MemberOrderIpError          = "member.order.ip_get_fail"
	MemberOrderCreateFail       = "member.order.create_fail"
	MemberSubscribeAddFail      = "member.subscribe.add_fail"
	MemberSubscribeRefuse       = "member.subscribe_refuse"
	MemberOrderIdVerify         = "member.order.id_verify"
	MemberOrderGenderNotSupport = "member.order.gender_not_support"
	MemberOrderPreviousGearNeed = "member.order.previous_gear_need"
)

var memberMap = map[string]int{
	MemberOrderIpError:          1,
	MemberOrderCreateFail:       2,
	MemberSubscribeAddFail:      3,
	MemberSubscribeRefuse:       3,
	MemberOrderIdVerify:         4,
	MemberOrderGenderNotSupport: 5,
	MemberOrderPreviousGearNeed: 5,
}
