// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_bool_test.go
// @Date: 2021/12/23 15:32

package esq

import (
	"testing"
)

func TestBool(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"bool",
			Bool().Must(
				Match("message", "this is a test"),
			),
			Map{
				"bool": Map{
					"must": []Map{
						{
							"match": Map{
								"message": Map{
									"query": "this is a test",
								},
							},
						},
					},
				},
			},
		},
		{
			"bool",
			Bool().Must(
				Match("message", "this is a test"),
			).Should(
				Match("message", "this is a test"),
			),
			Map{
				"bool": Map{
					"must": []Map{
						{
							"match": Map{
								"message": Map{
									"query": "this is a test",
								},
							},
						},
					},
					"should": []Map{
						{
							"match": Map{
								"message": Map{
									"query": "this is a test",
								},
							},
						},
					},
				},
			},
		},
		{
			"bool",
			Bool().Should(
				Match("message", "this is a test"),
			).MinimumShouldMatch(2),
			Map{
				"bool": Map{
					"should": []Map{
						{
							"match": Map{
								"message": Map{
									"query": "this is a test",
								},
							},
						},
					},
					"minimum_should_match": 2,
				},
			},
		},
		{
			"bool",
			Bool().Should(
				Match("message", "this is a test"),
			).MinimumShouldMatch(2).Boost(1.0),
			Map{
				"bool": Map{
					"should": []Map{
						{
							"match": Map{
								"message": Map{
									"query": "this is a test",
								},
							},
						},
					},
					"minimum_should_match": 2,
					"boost":                1.0,
				},
			},
		},
		{
			"bool",
			Bool().MustNot(
				Match("message", "this is a test"),
			).Boost(1.0),
			Map{
				"bool": Map{
					"must_not": []Map{
						{
							"match": Map{
								"message": Map{
									"query": "this is a test",
								},
							},
						},
					},
					"boost": 1.0,
				},
			},
		},
	})
}
