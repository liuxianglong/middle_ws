package code

const (
	GuestUidUnValid         = "guest.Uid_UnValid"
	GuestErrorBlock         = "guest.error_block"
	GuestIllegalOptions     = "guest.illegal_options"
	GuestColumnUnderReview  = "guest.column_under_review"
	GuestUploadMaxError     = "guest.upload_image_max_error"
	GuestCurrentImageEmpty  = "guest.current_img_empty"
	GuestNeedIdVerify       = "guest.need_id_verify"
	GuestSendImNeedIdVerify = "guest.send_need_id_verify"
	GuestSendRefuse         = "guest.send_refuse"
	GuestSendLeaveHas       = "guest.send_leave_has"
	GuestSendMsgMustLike    = "guest.send_msg_must_like"
	GuestSendMsgRemainLess  = "guest.send_msg_remain_less"
	GuestLikeRemainLess     = "guest.like_remain_less"
)

var guestMap = map[string]int{
	GuestUidUnValid:         1,
	GuestIllegalOptions:     2,
	GuestErrorBlock:         3,
	GuestColumnUnderReview:  4,
	GuestUploadMaxError:     5,
	GuestCurrentImageEmpty:  6,
	GuestNeedIdVerify:       7,
	GuestSendImNeedIdVerify: 8,
	GuestSendRefuse:         9,
	GuestSendLeaveHas:       10,
	GuestSendMsgMustLike:    11,
	GuestSendMsgRemainLess:  12,
	GuestLikeRemainLess:     12,
}
