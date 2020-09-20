package graphql

import (
	"github.com/aeramu/nim-finder-server/entity"
	"github.com/graph-gophers/graphql-go"
)

//User graphql
type User interface {
	ID() graphql.ID
	NimTPB() string
	NimJurusan() string
	Nama() string
	Fakultas() string
	Jurusan() string
	Status() string
}

type user struct {
	user entity.User
}

func (u user) ID() graphql.ID {
	return graphql.ID(u.user.ID())
}

func (u user) NimTPB() string {
	return u.user.NimTPB()
}

func (u user) NimJurusan() string {
	return u.user.NimJurusan()
}

func (u user) Nama() string {
	return u.user.Nama()
}

func (u user) Fakultas() string {
	return u.user.Fakultas()
}

func (u user) Jurusan() string {
	return u.user.Jurusan()
}

func (u user) Status() string {
	return u.user.Status()
}
