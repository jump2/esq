// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: search
// @Date: 2021/11/9 15:42

package esq

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/elastic/go-elasticsearch/v7"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type search struct {
	debug     bool
	parameter SearchParameter
}

type SearchParameter struct {
	Query  Map      `json:"query"`
	Size   *int     `json:"size,omitempty"`
	From   *int     `json:"from,omitempty"`
	Source []string `json:"_source,omitempty"`
	Sort   []Map    `json:"sort,omitempty"`
}

func Search() *search {
	return &search{}
}

func (s *search) Debug(v bool) *search {
	s.debug = v
	return s
}

func (s *search) Query(m Mappable) *search {
	s.parameter.Query = m.Map()
	return s
}

func (s *search) Size(v int) *search {
	s.parameter.Size = &v
	return s
}

func (s *search) From(v int) *search {
	s.parameter.From = &v
	return s
}

func (s *search) Source(v []string) *search {
	s.parameter.Source = v
	return s
}

func (s *search) Sort(v ...Map) *search {
	s.parameter.Sort = v
	return s
}

func (s *search) Map() Map {
	var (
		b []byte
		m Map
	)

	b, err := json.Marshal(s.parameter)
	if err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	if err := json.Unmarshal(b, &m); err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	return m
}

func (s *search) Run(client *elasticsearch.Client, o ...func(*esapi.SearchRequest)) (*esapi.Response, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(s.parameter); err != nil {
		log.Printf("Error encoding query: %s", err)
	}

	o = append(o, client.Search.WithBody(&buf))
	return client.Search(o...)
}
