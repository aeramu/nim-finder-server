package graphql

import (
	"github.com/aeramu/nim-finder-server/entity"
	"github.com/graph-gophers/graphql-go"
)

type pageInfo struct {
	entity.PageInfo
}

func (r pageInfo) StartCursor() *graphql.ID {
	startCursor := graphql.ID(r.PageInfo.StartCursor())
	return &startCursor
}

func (r pageInfo) EndCursor() *graphql.ID {
	endCursor := graphql.ID(r.PageInfo.EndCursor())
	return &endCursor
}
