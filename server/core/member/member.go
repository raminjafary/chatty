package member

type memberService struct {
	memeberRepo MemberRepository
}

func NewMemberService(memeberRepo MemberRepository) MemberService {
	return &memberService{
		memeberRepo,
	}
}

func (m *memberService) CreateMember(member *Member) (*Member, error) {
	return m.memeberRepo.CreateMember(member)
}
