package graphql

import (
	"github.com/graph-gophers/graphql-go"
)

//PageInfo graphql
type PageInfo interface {
	StartCursor() *graphql.ID
	EndCursor() *graphql.ID
}

type pageInfo []interface{ ID() graphql.ID }

func (r pageInfo) StartCursor() *graphql.ID {
	if len(r) == 0 {
		return nil
	}
	startCursor := graphql.ID(r[0].ID())
	return &startCursor
}

func (r pageInfo) EndCursor() *graphql.ID {
	if len(r) == 0 {
		return nil
	}
	endCursor := graphql.ID(r[len(r)-1].ID())
	return &endCursor
}
