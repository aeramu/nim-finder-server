package graphql

import (
	"github.com/graph-gophers/graphql-go"
)

//PageInfo graphql
type PageInfo interface {
	StartCursor() *graphql.ID
	EndCursor() *graphql.ID
}

type pageInfo struct {
	nodeList []interface{ Id() string }
}

func (r pageInfo) StartCursor() *graphql.ID {
	if len(r.nodeList) == 0 {
		return nil
	}
	startCursor := graphql.ID(r.nodeList[0].Id())
	return &startCursor
}

func (r pageInfo) EndCursor() *graphql.ID {
	if len(r.nodeList) == 0 {
		return nil
	}
	endCursor := graphql.ID(r.nodeList[len(r.nodeList)-1].Id())
	return &endCursor
}