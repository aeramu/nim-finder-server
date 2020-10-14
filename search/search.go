package search

import "github.com/aeramu/nim-finder-server/entity"

func (i *service) Search(keyword string, limit int, after string) entity.UserConnection {
	if after == "" {
		after = "ffffffffffffffffffffffff"
	}
	users := i.r.Search(keyword, limit, after)
	return users
}
