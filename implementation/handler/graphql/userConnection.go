package graphql

import "github.com/aeramu/nim-finder-server/entity"

//UserConnection graphql
type UserConnection interface {
	Edges() []User
	PageInfo() PageInfo
}

type userConnection struct {
	users []entity.User
}

func (uc *userConnection) Edges() []User {
	var users []User
	for _, u := range uc.users {
		users = append(users, user{u})
	}

	return users
}

func (uc *userConnection) PageInfo() PageInfo {
	var nodes []interface{ Id() string }
	return pageInfo{nodes}
}
