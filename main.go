package main

import (
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

func main() {
	// routes.Run()
	es, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9201"},
	})
	body := ``
	res, err := es.Search(
		es.Search.WithIndex("try_index_dua"),
		es.Search.WithBody(strings.NewReader(body)),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println(res)
	defer res.Body.Close()
}
