package graphql

import (
	"context"
	"net/http"

	"github.com/aeramu/nim-finder-server/internal/search"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

//NewHandler handler
func NewHandler(s search.Service) Handler {
	schema := graphql.MustParseSchema(schemaString, &resolver{
		s: s,
	})

	return &handler{&relay.Handler{Schema: schema}}
}

//Handler interface
type Handler interface {
	http.Handler
	Response(ctx context.Context, query string, operation string, variables map[string]interface{}) *graphql.Response
}

type handler struct {
	*relay.Handler
}

func (h handler) Response(ctx context.Context, query string, operation string, variables map[string]interface{}) *graphql.Response {
	return h.Schema.Exec(ctx, query, operation, variables)
}

var schemaString = `
  	schema{
		query: Query
  	}
  	type Query{
		search(keyword: String!, first: Int!, after: ID): UserConnection!
	}
	type User{
		id: ID!
		nimTPB: String!
		nimJurusan: String!
		nama: String!
		fakultas: String!
		jurusan: String!
		status: String!
	}
	type UserConnection{
		edges: [User!]!
		pageInfo: PageInfo!
	}
	type PageInfo{
		startCursor: ID
		endCursor: ID
		hasNextPage: Boolean!
	}
`
