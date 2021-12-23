// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_term_level_test.go
// @Date: 2021/11/17 16:17

package esq

import (
	"testing"
)

func TestTerm(t *testing.T) {
	//Term Query Test
	runMapTest(t, []mapTest{
		{
			"term",
			Term("order_id", "541"),
			Map{
				"term": Map{
					"order_id": Map{
						"value": "541",
					},
				},
			},
		},
		{
			"term",
			Term("order_id", "541").Boost(1.0),
			Map{
				"term": Map{
					"order_id": Map{
						"value": "541",
						"boost": 1.0,
					},
				},
			},
		},
		{
			"term",
			Term("order_id", ""),
			Map{
				"term": Map{
					"order_id": Map{
						"value": "",
					},
				},
			},
		},
		{
			"term",
			Term("order_id", 0),
			Map{
				"term": Map{
					"order_id": Map{
						"value": 0,
					},
				},
			},
		},
		{
			"term",
			Term("order_id", "").OmitValue(""),
			nil,
		},
		{
			"term",
			Term("order_id", 0).OmitValue(0),
			nil,
		},
		{
			"term",
			Term("order_id", "0").OmitValue(0),
			nil,
		},
		{
			"term",
			Term("order_id", "541").OmitValue(0),
			Map{
				"term": Map{
					"order_id": Map{
						"value": "541",
					},
				},
			},
		},
	})
	//Terms Query Test
	runMapTest(t, []mapTest{
		{
			"terms",
			Terms("order_id", "541"),
			Map{
				"terms": Map{
					"order_id": []string{"541"},
				},
			},
		},
		{
			"terms",
			Terms("order_id", "541").Boost(1.0),
			Map{
				"terms": Map{
					"order_id": []string{"541"},
					"boost":    1.0,
				},
			},
		},
		{
			"terms",
			Terms("order_id", ""),
			Map{
				"terms": Map{
					"order_id": []string{""},
				},
			},
		},
		{
			"terms",
			Terms("order_id", "541", "52541"),
			Map{
				"terms": Map{
					"order_id": []string{"541", "52541"},
				},
			},
		},
		{
			"term",
			Terms("order_id", "").OmitValue(""),
			nil,
		},
		{
			"terms",
			Terms("order_id", 0).OmitValue(0),
			nil,
		},
		{
			"terms",
			Terms("order_id", "0").OmitValue(0),
			nil,
		},
		{
			"terms",
			Terms("order_id", []string{"541", "52541", "0"}).OmitValue(0),
			Map{
				"terms": Map{
					"order_id": []string{"541", "52541"},
				},
			},
		},
		{
			"terms",
			Terms("order_id", []string{"541", "52541", ""}),
			Map{
				"terms": Map{
					"order_id": []string{"541", "52541", ""},
				},
			},
		},
	})
	//Terms Set Query Test
	runMapTest(t, []mapTest{
		{
			"terms_set",
			TermsSet("order_id", "541"),
			Map{
				"terms_set": Map{
					"order_id": Map{
						"terms": []string{"541"},
					},
				},
			},
		},
		{
			"terms_set",
			TermsSet("order_id", "541").MinimumShouldMatchField("required_matches"),
			Map{
				"terms_set": Map{
					"order_id": Map{
						"terms":                      []string{"541"},
						"minimum_should_match_field": "required_matches",
					},
				},
			},
		},
		{
			"terms_set",
			TermsSet("order_id", "541").MinimumShouldMatchScript("Math.min(params.num_terms, doc['required_matches'].value)"),
			Map{
				"terms_set": Map{
					"order_id": Map{
						"terms": []string{"541"},
						"minimum_should_match_script": Map{
							"source": "Math.min(params.num_terms, doc['required_matches'].value)",
						},
					},
				},
			},
		},
		{
			"terms_set",
			TermsSet("order_id", []string{"aaa", "ddd"}).OmitValue("aaa"),
			Map{
				"terms_set": Map{
					"order_id": Map{
						"terms": []string{"ddd"},
					},
				},
			},
		},
	})
	//Wildcard Query
	runMapTest(t, []mapTest{
		{
			"wildcard",
			Wildcard("user", "ki*y"),
			Map{
				"wildcard": Map{
					"user": Map{
						"value": "ki*y",
					},
				},
			},
		},
		{
			"wildcard",
			Wildcard("user", "ki*y").Boost(1.0),
			Map{
				"wildcard": Map{
					"user": Map{
						"value": "ki*y",
						"boost": 1.0,
					},
				},
			},
		},
		{
			"wildcard",
			Wildcard("user", "ki*y").Boost(1.0).Rewrite(ConstantScore),
			Map{
				"wildcard": Map{
					"user": Map{
						"value":   "ki*y",
						"boost":   1.0,
						"rewrite": "constant_score",
					},
				},
			},
		},
		{
			"wildcard",
			Wildcard("user", "").OmitValue(""),
			nil,
		},
	})
	//Regexp Query
	runMapTest(t, []mapTest{
		{
			"regexp",
			Regex("user", "k.*y"),
			Map{
				"regexp": Map{
					"user": Map{
						"value": "k.*y",
					},
				},
			},
		},
		{
			"regexp",
			Regex("user", "k.*y").Flags("ALL"),
			Map{
				"regexp": Map{
					"user": Map{
						"value": "k.*y",
						"flags": "ALL",
					},
				},
			},
		},
		{
			"regexp",
			Regex("user", "k.*y").Flags("ALL").MaxDeterminizedStates(10000),
			Map{
				"regexp": Map{
					"user": Map{
						"value":                   "k.*y",
						"flags":                   "ALL",
						"max_determinized_states": 10000,
					},
				},
			},
		},
		{
			"regexp",
			Regex("user", "k.*y").Flags("ALL").MaxDeterminizedStates(10000).Rewrite(ConstantScore),
			Map{
				"regexp": Map{
					"user": Map{
						"value":                   "k.*y",
						"flags":                   "ALL",
						"max_determinized_states": 10000,
						"rewrite":                 "constant_score",
					},
				},
			},
		},
		{
			"regexp",
			Regex("user", "").OmitValue(""),
			nil,
		},
	})
	//Range Query
	runMapTest(t, []mapTest{
		{
			"range",
			Range("timestamp").Gte("2015-01-01 00:00:00").Lte("now"),
			Map{
				"range": Map{
					"timestamp": Map{
						"gte": "2015-01-01 00:00:00",
						"lte": "now",
					},
				},
			},
		},
		{
			"range",
			Range("timestamp").Gt("2015-01-01 00:00:00").Lt("now"),
			Map{
				"range": Map{
					"timestamp": Map{
						"gt": "2015-01-01 00:00:00",
						"lt": "now",
					},
				},
			},
		},
		{
			"range",
			Range("timestamp").Gt("2015-01-01 00:00:00").Lt("now").Boost(2.0),
			Map{
				"range": Map{
					"timestamp": Map{
						"gt":    "2015-01-01 00:00:00",
						"lt":    "now",
						"boost": 2.0,
					},
				},
			},
		},
		{
			"range",
			Range("timestamp").Gt("2015-01-01 00:00:00").Lt("now").Boost(2.0).Relation(RelationIntersects),
			Map{
				"range": Map{
					"timestamp": Map{
						"gt":       "2015-01-01 00:00:00",
						"lt":       "now",
						"boost":    2.0,
						"relation": "INTERSECTS",
					},
				},
			},
		},
		{
			"range",
			Range("timestamp").Gt("").Lt("now").OmitValue(""),
			Map{
				"range": Map{
					"timestamp": Map{
						"lt": "now",
					},
				},
			},
		},
		{
			"range",
			Range("timestamp").OmitValue(""),
			nil,
		},
	})
	//Prefix Query
	runMapTest(t, []mapTest{
		{
			"prefix",
			Prefix("user", "ki"),
			Map{
				"prefix": Map{
					"user": Map{
						"value": "ki",
					},
				},
			},
		},
		{
			"prefix",
			Prefix("user", "ki").Rewrite(ConstantScore),
			Map{
				"prefix": Map{
					"user": Map{
						"value":   "ki",
						"rewrite": "constant_score",
					},
				},
			},
		},
		{
			"prefix",
			Prefix("user", "").OmitValue(""),
			nil,
		},
	})
	//IDs Query
	runMapTest(t, []mapTest{
		{
			"ids",
			Ids("1", "4", "100"),
			Map{
				"ids": Map{
					"values": []string{"1", "4", "100"},
				},
			},
		},
	})
	//Fuzzy Query
	runMapTest(t, []mapTest{
		{
			"fuzzy",
			Fuzzy("user", "ki"),
			Map{
				"fuzzy": Map{
					"user": Map{
						"value": "ki",
					},
				},
			},
		},
	})
	//Exists Query
	runMapTest(t, []mapTest{
		{
			"exists",
			Exists("user"),
			Map{
				"exists": Map{
					"field": "user",
				},
			},
		},
	})
}
