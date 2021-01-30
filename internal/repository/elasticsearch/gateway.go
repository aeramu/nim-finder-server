package elasticsearch

import (
	"encoding/json"

	"github.com/aeramu/nim-finder-server/internal/search"

	es "github.com/olivere/elastic/v7"
)

// Student model
type studentModel struct {
	Username   string `json:"username"`
	NimTPB     string `json:"nim_tpb"`
	NimJurusan string `json:"nim_jurusan"`
	Nama       string `json:"nama"`
	Status     string `json:"status"`
	Fakultas   string `json:"fakultas"`
	Jurusan    string `json:"jurusan"`
	EmailITB   string `json:"email_itb"`
	Email      string `json:"email"`
}

type student es.SearchHit

func (s student) toModel() search.Student {
	var std studentModel
	json.Unmarshal(s.Source, &std)
	return search.Student{
		ID:         s.Id,
		Username:   std.Username,
		NimTPB:     std.NimTPB,
		NimJurusan: std.NimJurusan,
		Nama:       std.Nama,
		Status:     std.Status,
		Fakultas:   std.Fakultas,
		Jurusan:    std.Jurusan,
		Email:      std.Email,
		EmailITB:   std.EmailITB,
	}
}

type students []*es.SearchHit

func (s students) toModel() search.StudentConnection {
	println("masooookkk")
	var students []search.Student
	for _, std := range s {
		students = append(students, student(*std).toModel())
	}
	println("masooookkk")
	var pageInfo search.PageInfo
	if len(s) > 0 {
		pageInfo = search.PageInfo{
			StartCursor: s[0].Id,
			EndCursor:   s[len(s)-1].Id,
		}
	} else {
		pageInfo = search.PageInfo{
			StartCursor: "",
			EndCursor:   "",
		}
	}
	println("masooookkk")
	return search.StudentConnection{
		Edges:    students,
		PageInfo: pageInfo,
	}
}
