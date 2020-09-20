package entity

//User entity
type User struct {
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
