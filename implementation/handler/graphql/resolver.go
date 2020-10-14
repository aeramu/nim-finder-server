package graphql

import (
	"github.com/aeramu/nim-finder-server/search"
	"github.com/graph-gophers/graphql-go"
)

type resolver struct {
	i search.Service
}

func (r *resolver) Search(args struct {
	Keyword string
	First   int32
	After   *graphql.ID
}) userConnection {
	after := ""
	if args.After != nil {
		after = string(*args.After)
	}
	users := r.i.Search(args.Keyword, int(args.First), after)
	return userConnection{users}
}
