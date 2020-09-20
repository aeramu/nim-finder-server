package service

import "github.com/aeramu/nim-finder-server/entity"

func (i *interactor) Search(keyword string) []entity.User {
	users := i.r.Search(keyword)
	return users
}
