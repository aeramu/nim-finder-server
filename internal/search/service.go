package search

//NewService interactor
func NewService(r StudentRepo) Service {
	return &service{
		studentRepo: r,
	}
}

//Service service
type Service interface {
	Search(keyword string, limit int, after string) StudentConnection
}

type service struct {
	studentRepo StudentRepo
}

func (s *service) Search(keyword string, limit int, after string) StudentConnection {
	if after == "" {
		after = "ffffffffffffffffffffffff"
	}
	users := s.studentRepo.Search(keyword, limit, after)
	return users
}
