package graphql

import (
	"github.com/aeramu/nim-finder-server/usecase/service"
	"github.com/graph-gophers/graphql-go"
)

type resolver struct {
	i service.Interactor
}

func (r *resolver) Search(args struct {
	Keyword string
	First   int32
	After   *graphql.ID
	Before  *graphql.ID
}) userConnection {
	users := r.i.Search(args.Keyword, int(args.First))
	return userConnection(users)
}
