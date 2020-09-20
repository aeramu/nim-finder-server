package graphql

import "github.com/aeramu/nim-finder-server/entity"

type UserConnection interface {
	Edges() []User
	PageInfo() PageInfo
}

type userConnection struct {
	users []entity.User
}

func (uc *userConnection) Edges() []User {
	var users []User
	for _, user := range uc.users {
		users = append(users, user)
	}

	return users
}

func (uc *userConnection) PageInfo() PageInfo {
	var nodes []interface{ Id() string }
	return pageInfo{}
}
