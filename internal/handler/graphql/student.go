package graphql

import (
	"github.com/aeramu/nim-finder-server/internal/search"
	"github.com/graph-gophers/graphql-go"
)

type student struct {
	search.Student
}

func (s student) ID() graphql.ID {
	return graphql.ID(s.Student.ID)
}
func (s student) NimTPB() string {
	return s.Student.NimTPB
}
func (s student) NimJurusan() string {
	return s.Student.NimJurusan
}
func (s student) Nama() string {
	return s.Student.Nama
}
func (s student) Status() string {
	return s.Student.Status
}
func (s student) Fakultas() string {
	return s.Student.Fakultas
}
func (s student) Jurusan() string {
	return s.Student.Jurusan
}
