package service

import "github.com/aeramu/nim-finder-server/entity"

func (i *interactor) Search(keyword string, limit int, after string) []entity.User {
	if after == "" {
		after = "ffffffffffffffffffffffff"
	}
	users := i.r.Search(keyword, limit, after)
	return users
}
