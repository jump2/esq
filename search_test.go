// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: search_test.go
// @Date: 2021/12/23 17:23

package esq

import (
	"testing"
)

func TestSearch(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"search",
			Search().Query(
				Bool().Must(
					Term("program_language", "golang"),
				).Boost(1.0),
			).Size(20).From(100).Sort(Map{
				"created_at": "desc",
			}),
			Map{
				"query": Map{
					"bool": Map{
						"must": []Map{
							{
								"term": Map{
									"program_language": Map{
										"value": "golang",
									},
								},
							},
						},
						"boost": 1.0,
					},
				},
				"size": 20,
				"from": 100,
				"sort": []Map{
					{
						"created_at": "desc",
					},
				},
			},
		},
	})
}
