package graphql

import (
	"github.com/aeramu/nim-finder-server/entity"
)

type userConnection struct {
	entity.UserConnection
}

func (uc userConnection) Edges() []user {
	var users []user
	for _, u := range uc.UserConnection.Edges() {
		users = append(users, user{u})
	}

	return users
}

func (uc userConnection) PageInfo() pageInfo {
	return pageInfo{uc.UserConnection.PageInfo()}
}
