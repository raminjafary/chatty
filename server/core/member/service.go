package member

type MemberService interface {
	CreateMember(*Member) (*Member, error)
}
