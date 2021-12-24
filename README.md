# esq
elasticsearch query

## Usage
```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jump2/esq"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Failed creating client: %s", err)
	}

	resp, err := esq.Search().Query(
		esq.Bool().Must(
			esq.Term("program_language", "golang"),
			esq.Terms("status", []string{"enabled"}),
		),
	).Size(20).Sort(esq.Map{
		"age": "desc",
	}).Run(
		es,
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("huang_test"),
	)

	if err != nil {
		log.Fatalf("Failed searching for stuff: %s", err)
	}

	defer resp.Body.Close()

	var e map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	fmt.Println(e)
	//...
}

```