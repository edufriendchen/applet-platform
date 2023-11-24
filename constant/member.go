package constant

type MemberStatus int

const (
	ExternalMemberDeactivate MemberStatus = iota
	ExternalMemberActive
	ExternalMemberInvited
	ExternalMemberBlocked
)
