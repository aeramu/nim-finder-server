package search

import (
	"github.com/aeramu/nim-finder-server/entity"
)

//New interactor
func New(r Repository) Service {
	return &service{
		r: r,
	}
}

//Service service
type Service interface {
	Search(keyword string, limit int, after string) entity.UserConnection
}

type service struct {
	r Repository
}
