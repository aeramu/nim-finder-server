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

//UserConstructor entity
type UserConstructor struct {
	ID         string
	NimTPB     string
	NimJurusan string
	Nama       string
	Status     string
	Fakultas   string
	Jurusan    string
}

//New return User interface
func (c UserConstructor) New() User {
	return &user{
		id:         c.ID,
		nimTPB:     c.NimTPB,
		nimJurusan: c.NimJurusan,
		nama:       c.Nama,
		status:     c.Status,
		fakultas:   c.Fakultas,
		jurusan:    c.Jurusan,
	}
}

type user struct {
	id         string
	nimTPB     string
	nimJurusan string
	nama       string
	status     string
	fakultas   string
	jurusan    string
}

func (u user) ID() string {
	return u.id
}

func (u user) NimTPB() string {
	return u.nimTPB
}

func (u user) NimJurusan() string {
	return u.nimJurusan
}

func (u user) Nama() string {
	return u.nama
}

func (u user) Status() string {
	return u.status
}

func (u user) Fakultas() string {
	return u.fakultas
}

func (u user) Jurusan() string {
	return u.jurusan
}
