package mongodb

import (
	"github.com/aeramu/nim-finder-server/entity"
)

//User entity
type user struct {
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

func (u user) entity() entity.User {
	return entity.UserConstructor{
		ID:         u.ID,
		NimTPB:     u.NimTPB,
		NimJurusan: u.NimJurusan,
		Nama:       u.Nama,
		Status:     u.Status,
		Fakultas:   u.Fakultas,
		Jurusan:    u.Jurusan,
	}.New()
}

type users []*user

func (u users) entity() []entity.User {
	var users []entity.User
	for _, user := range u {
		users = append(users, user.entity())
	}
	return users
}
