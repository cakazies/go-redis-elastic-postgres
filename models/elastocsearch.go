package models

import (
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

type (
	ESModesl struct{}
)

func (ESModesl) Request(method string, data string) ([]byte, error) {

	es, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	body := `{
		"query": {
			"match": { "message": "myProduct" }
		},
		"aggregations": {
			"top_10_states": { "terms": { "field": "state", "size": 10 } }
		}
	}`
	res, err := es.Search(
		es.Search.WithIndex("social-*"),
		es.Search.WithBody(strings.NewReader(body)),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println(res)
	defer res.Body.Close()
	return nil, nil
}

func (ESModesl) CreateDocs() {

}
