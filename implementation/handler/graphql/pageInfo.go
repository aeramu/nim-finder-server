package graphql

import (
	"github.com/aeramu/nim-finder-server/entity"
	"github.com/graph-gophers/graphql-go"
)

type pageInfo struct {
	entity.PageInfo
}

func (r pageInfo) StartCursor() *graphql.ID {
	if r.PageInfo.StartCursor() != nil {
		cursor := graphql.ID(*r.PageInfo.StartCursor())
		return &cursor
	}
	return nil
}

func (r pageInfo) EndCursor() *graphql.ID {
	if r.PageInfo.EndCursor() != nil {
		cursor := graphql.ID(*r.PageInfo.EndCursor())
		return &cursor
	}
	return nil
}
