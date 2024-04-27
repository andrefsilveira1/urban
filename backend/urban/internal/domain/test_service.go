package domain

type TestService struct {
	testRepository TestRepository
}

func NewTestService(testRepository TestRepository) *TestService {
	return &TestService{
		testRepository: testRepository,
	}
}

func (s *TestService) Test() error {
	return s.testRepository.Test()
}
