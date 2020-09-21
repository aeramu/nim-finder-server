package service

import (
	"github.com/aeramu/nim-finder-server/entity"
	"github.com/aeramu/nim-finder-server/usecase/adapter"
)

//New interactor
func New(r adapter.Repository) Interactor {
	return &interactor{
		r: r,
	}
}

//Interactor service
type Interactor interface {
	Search(keyword string, limit int, after string) []entity.User
}
