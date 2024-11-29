package code

import (
	"pp-server/internal/consts"
)

const (
	WalletHistorySubjectPayMember1       = "wallet.subject.pay_member1"
	WalletHistorySubjectPayMember2       = "wallet.subject.pay_member2"
	WalletHistorySubjectPayMember3       = "wallet.subject.pay_member3"
	WalletHistorySubjectPayMember4       = "wallet.subject.pay_member4"
	WalletHistorySubjectWithdraw         = "wallet.subject.cash"
	WalletHistorySubjectWithdrawFailed   = "wallet.subject.cash_back"
	WalletHistorySubjectInvent           = "wallet.subject.invent"
	WalletHistorySubjectInventSub        = "wallet.subject.invent.detail"
	WalletHistorySubjectExpiredGoldClear = "wallet.subject.expired.gold.clear"

	WalletSubjectInviteRegister              = "wallet.subject.invite.register"
	WalletSubjectInviteLogin                 = "wallet.subject.invite.login"
	WalletSubjectInviteBuy                   = "wallet.subject.invite.buy"
	WalletSubjectInviteRegisterInviterDetail = "wallet.subject.invite.register.inviter.detail"
	WalletSubjectInviteRegisterInviteeDetail = "wallet.subject.invite.register.invitee.detail"

	WalletSubjectInviteLoginInviterDetail = "wallet.subject.invite.login.inviter.detail"
	WalletSubjectInviteLoginInviteeDetail = "wallet.subject.invite.login.invitee.detail"

	WalletSubjectInviteBuyInviterDetail = "wallet.subject.invite.buy.inviter.detail"

	WalletHistoryStatusRewardsSend    = "wallet.status.rewards.send"
	WalletHistoryStatusPaySuccess     = "wallet.status.pay.success"
	WalletHistoryStatusCashProcessing = "wallet.status.cash.processing"
	WalletHistoryStatusCashSuccess    = "wallet.status.cash.success"
	WalletHistoryStatusCashFailed     = "wallet.status.cash.failed"
)

var WalletWithdrawStatusMap = map[uint]string{
	consts.WithdrawStatusProcessing: WalletHistoryStatusCashProcessing,
	consts.WithdrawStatusSuccess:    WalletHistoryStatusCashSuccess,
	consts.WithdrawStatusFailed:     WalletHistoryStatusCashFailed,
}
