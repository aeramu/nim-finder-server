package graphql

import (
	"github.com/aeramu/nim-finder-server/internal/search"
	"github.com/graph-gophers/graphql-go"
)

type resolver struct {
	s search.Service
}

func (r *resolver) Search(args struct {
	Keyword string
	First   int32
	After   *graphql.ID
}) studentConnection {
	after := ""
	if args.After != nil {
		after = string(*args.After)
	}
	students := r.s.Search(args.Keyword, int(args.First), after)
	return studentConnection{students}
}

type studentConnection struct {
	search.StudentConnection
}

func (s studentConnection) Edges() []student {
	var students []student
	for _, std := range s.StudentConnection.Edges {
		students = append(students, student{std})
	}

	return students
}

func (s studentConnection) PageInfo() pageInfo {
	return pageInfo{s.StudentConnection.PageInfo}
}
