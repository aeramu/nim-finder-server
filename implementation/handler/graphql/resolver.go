package graphql

import (
	"github.com/aeramu/nim-finder-server/usecase/service"
	"github.com/graph-gophers/graphql-go"
)

//Resolver graphql
type Resolver interface {
	Search(args struct {
		Keyword string
		First   int32
		After   *graphql.ID
		Before  *graphql.ID
	}) UserConnection
}

type resolver struct {
	i service.Interactor
}

func (r *resolver) Search(args struct {
	Keyword string
	First   int32
	After   *graphql.ID
	Before  *graphql.ID
}) UserConnection {
	users := r.i.Search(args.Keyword)
	return &userConnection{users}
}
