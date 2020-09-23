package entity

//User interface
type User interface {
	ID() string
	NimTPB() string
	NimJurusan() string
	Nama() string
	Status() string
	Fakultas() string
	Jurusan() string
}

//UserConnection interface
type UserConnection interface {
	Edges() []User
	PageInfo() PageInfo
}

//PageInfo interface
type PageInfo interface {
	StartCursor() *string
	EndCursor() *string
	HasNextPage() bool
}
