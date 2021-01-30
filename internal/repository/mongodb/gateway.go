package mongodb

import "github.com/aeramu/nim-finder-server/entity"

//User entity
type userModel struct {
	ID         string `bson:"_id"`
	Username   string `bson:"username"`
	NimTPB     string `bson:"nim_tpb"`
	NimJurusan string `bson:"nim_jurusan"`
	Nama       string `bson:"nama"`
	Status     string `bson:"status"`
	Fakultas   string `bson:"fakultas"`
	Jurusan    string `bson:"jurusan"`
	EmailITB   string `bson:"email_itb"`
	Email      string `bson:"email"`
}

type userModels []userModel

func (m userModels) entity() []entity.User {
	var users []entity.User
	for _, model := range m {
		users = append(users, user{model})
	}
	return users
}

type user struct {
	m userModel
}

func (u user) ID() string {
	return u.m.ID
}

func (u user) NimTPB() string {
	return u.m.NimTPB
}

func (u user) NimJurusan() string {
	return u.m.NimJurusan
}

func (u user) Nama() string {
	return u.m.Nama
}

func (u user) Status() string {
	return u.m.Status
}

func (u user) Fakultas() string {
	return u.m.Fakultas
}

func (u user) Jurusan() string {
	return u.m.Jurusan
}

type userConnection struct {
	edges       []entity.User
	hasNextPage bool
}

func (uc userConnection) Edges() []entity.User {
	return uc.edges
}

func (uc userConnection) PageInfo() entity.PageInfo {
	return pageInfo{uc.edges, uc.hasNextPage}
}

type pageInfo struct {
	edges       []entity.User
	hasNextPage bool
}

func (p pageInfo) StartCursor() string {
	if len(p.edges) == 0 {
		return ""
	}
	return p.edges[0].ID()
}

func (p pageInfo) EndCursor() string {
	if len(p.edges) == 0 {
		return ""
	}
	return p.edges[len(p.edges)-1].ID()
}

func (p pageInfo) HasNextPage() bool {
	return p.hasNextPage
}
