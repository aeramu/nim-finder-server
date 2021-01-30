package search

// Student model
type Student struct {
	ID         string
	Username   string
	NimTPB     string
	NimJurusan string
	Nama       string
	Status     string
	Fakultas   string
	Jurusan    string
	EmailITB   string
	Email      string
}

// StudentConnection for pagination
type StudentConnection struct {
	Edges    []Student
	PageInfo PageInfo
}

// PageInfo pagination
type PageInfo struct {
	StartCursor string
	EndCursor   string
	HasNextPage bool
}
