package constant

type Status int

const (
	Deactivate Status = iota
	Pending
	Active
)

type CtxInfo string

const (
	CtxUserIDInfo CtxInfo = "USER_ID"
)
