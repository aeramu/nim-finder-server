package graphql

import (
	"github.com/aeramu/nim-finder-server/entity"
	"github.com/graph-gophers/graphql-go"
)

type userConnection []entity.User

func (uc userConnection) Edges() []user {

	var users []user
	for _, u := range uc {
		users = append(users, user{u})
	}

	return users
}

func (uc userConnection) PageInfo() pageInfo {
	var nodes []interface{ ID() graphql.ID }
	for _, u := range uc {
		nodes = append(nodes, user{u})
	}
	return pageInfo(nodes)
}
