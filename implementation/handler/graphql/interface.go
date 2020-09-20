package graphql

import (
	"net/http"

	"github.com/aeramu/nim-finder-server/usecase/service"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

//New handler
func New(i service.Interactor) http.Handler {
	schema := graphql.MustParseSchema(schemaString, &resolver{
		i: i,
	})

	return &relay.Handler{Schema: schema}
}

var schemaString = `
  	schema{
		query: Query
  	}
  	type Query{
		search(keyword: String!, first: Int!, after: ID, before: ID): UserConnection!
	}
	type User{
		id: ID!
		nimTPB: String!
		nimJurusan: String!
		nama: String!
		fakultas: String!
		jurusan: String!
	}
	type UserConnection{
		edges: [User]!
		pageInfo: PageInfo!
	}
	type PageInfo{
		startCursor: ID
		endCursor: ID
	}
`
