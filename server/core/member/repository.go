package member

type MemberRepository interface {
	CreateMember(*Member) (*Member, error)
}
