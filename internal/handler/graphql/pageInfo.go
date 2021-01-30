package graphql

import (
	"github.com/aeramu/nim-finder-server/internal/search"
	"github.com/graph-gophers/graphql-go"
)

type pageInfo struct {
	search.PageInfo
}

func (r pageInfo) StartCursor() *graphql.ID {
	if r.PageInfo.StartCursor == "" {
		return nil
	}

	cursor := graphql.ID(r.PageInfo.StartCursor)
	return &cursor
}

func (r pageInfo) EndCursor() *graphql.ID {
	if r.PageInfo.EndCursor == "" {
		return nil
	}

	cursor := graphql.ID(r.PageInfo.EndCursor)
	return &cursor
}

func (r pageInfo) HasNextPage() bool {
	return false
}
