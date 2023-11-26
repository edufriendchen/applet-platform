package constant

type Status int

const (
	Deactivate Status = iota
	Pending
	Active
)

type CtxInfo string

const (
	CtxMemberIDInfo CtxInfo = "member_id"
)
