package elasticsearch

import (
	"context"

	"github.com/aeramu/nim-finder-server/internal/search"
	es "github.com/olivere/elastic/v7"
)

var client *es.Client

//NewStudentRepo repository
func NewStudentRepo() search.StudentRepo {
	if client == nil {
		client, _ = es.NewClient(
			es.SetURL("https://zpTYAosPcw:9VqcAsxinmTUo5dPtH7N@qiup-582779233.ap-southeast-2.bonsaisearch.net:443"),
			es.SetSniff(false),
			es.SetHealthcheck(false),
		)
	}
	return &repository{
		client: client,
	}
}

type repository struct {
	client *es.Client
}

func (r *repository) Search(keyword string, limit int, after string) search.StudentConnection {
	search := es.NewSearchSource().Query(es.NewMultiMatchQuery(keyword, "nama", "fakultas", "jurusan", "nim_jurusam", "nim_tpb").Type("bool_prefix"))
	result, _ := r.client.Search().Index("students").SearchSource(search).Do(context.Background())
	// TODO: nil result
	return students(result.Hits.Hits).toModel()
}
