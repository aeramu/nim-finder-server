package service

import "github.com/aeramu/nim-finder-server/entity"

func (i *interactor) Search(keyword string, limit int) []entity.User {
	users := i.r.Search(keyword, limit)
	return users
}
