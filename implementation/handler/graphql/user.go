package graphql

import (
	"github.com/aeramu/nim-finder-server/entity"
	"github.com/graph-gophers/graphql-go"
)

type user struct {
	entity.User
}

func (u user) ID() graphql.ID {
	return graphql.ID(u.User.ID())
}
