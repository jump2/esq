// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_full_text_test.go
// @Date: 2021/12/23 10:38

package esq

import (
	"testing"
)

func TestMatch(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"match",
			Match("message", "this is a test"),
			Map{
				"match": Map{
					"message": Map{
						"query": "this is a test",
					},
				},
			},
		},
		{
			"match",
			Match("message", "this is a test").Operator(OperatorAnd),
			Map{
				"match": Map{
					"message": Map{
						"query":    "this is a test",
						"operator": "and",
					},
				},
			},
		},
		{
			"match",
			Match("message", "this is a test").Operator(OperatorAnd).Lenient(false),
			Map{
				"match": Map{
					"message": Map{
						"query":    "this is a test",
						"operator": "and",
						"lenient":  false,
					},
				},
			},
		},
		{
			"match",
			Match("message", "this is a test").Operator(OperatorAnd).Lenient(true).ZeroTermsQuery(ZeroTermsAll),
			Map{
				"match": Map{
					"message": Map{
						"query":            "this is a test",
						"operator":         "and",
						"lenient":          true,
						"zero_terms_query": "all",
					},
				},
			},
		},
		{
			"match",
			Match("message", "this is a test").AutoGenerateSynonymsPhraseQuery(false),
			Map{
				"match": Map{
					"message": Map{
						"query":                               "this is a test",
						"auto_generate_synonyms_phrase_query": false,
					},
				},
			},
		},
		{
			"match",
			Match("message", "this is a test").AutoGenerateSynonymsPhraseQuery(true),
			Map{
				"match": Map{
					"message": Map{
						"query":                               "this is a test",
						"auto_generate_synonyms_phrase_query": true,
					},
				},
			},
		},
		{
			"match",
			Match("message", "this is a test").AutoGenerateSynonymsPhraseQuery(true),
			Map{
				"match": Map{
					"message": Map{
						"query":                               "this is a test",
						"auto_generate_synonyms_phrase_query": true,
					},
				},
			},
		},
		{
			"match",
			Match("message", "").OmitValue(""),
			nil,
		},
	})
}

func TestMatchPhrase(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"match_phrase",
			MatchPhrase("message", "this is a test"),
			Map{
				"match_phrase": Map{
					"message": Map{
						"query": "this is a test",
					},
				},
			},
		},
		{
			"match_phrase",
			MatchPhrase("message", "this is a test").Analyzer("my_analyzer"),
			Map{
				"match_phrase": Map{
					"message": Map{
						"query":    "this is a test",
						"analyzer": "my_analyzer",
					},
				},
			},
		},
		{
			"match_phrase",
			MatchPhrase("message", "this is a test").Analyzer("my_analyzer").Slop(2),
			Map{
				"match_phrase": Map{
					"message": Map{
						"query":    "this is a test",
						"analyzer": "my_analyzer",
						"slop":     2,
					},
				},
			},
		},
		{
			"match_phrase",
			MatchPhrase("message", "this is a test").ZeroTermsQuery(ZeroTermsAll),
			Map{
				"match_phrase": Map{
					"message": Map{
						"query":            "this is a test",
						"zero_terms_query": "all",
					},
				},
			},
		},
		{
			"match_phrase",
			MatchPhrase("message", "").OmitValue(""),
			nil,
		},
	})
}

func TestMatchPhrasePrefix(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"match_phrase_prefix",
			MatchPhrasePrefix("message", "quick brown f"),
			Map{
				"match_phrase_prefix": Map{
					"message": Map{
						"query": "quick brown f",
					},
				},
			},
		},
		{
			"match_phrase_prefix",
			MatchPhrasePrefix("message", "quick brown f").Analyzer("my_analyzer"),
			Map{
				"match_phrase_prefix": Map{
					"message": Map{
						"query":    "quick brown f",
						"analyzer": "my_analyzer",
					},
				},
			},
		},
		{
			"match_phrase_prefix",
			MatchPhrasePrefix("message", "quick brown f").ZeroTermsQuery(ZeroTermsNone),
			Map{
				"match_phrase_prefix": Map{
					"message": Map{
						"query":            "quick brown f",
						"zero_terms_query": "none",
					},
				},
			},
		},
		{
			"match_phrase_prefix",
			MatchPhrasePrefix("message", "quick brown f").Slop(2),
			Map{
				"match_phrase_prefix": Map{
					"message": Map{
						"query": "quick brown f",
						"slop":  2,
					},
				},
			},
		},
		{
			"match_phrase_prefix",
			MatchPhrasePrefix("message", "").OmitValue(""),
			nil,
		},
	})
}

func TestMultiMatch(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"multi_match",
			MultiMatch([]string{"message", "subject"}, "this is a test"),
			Map{
				"multi_match": Map{
					"query":  "this is a test",
					"fields": []string{"message", "subject"},
				},
			},
		},
		{
			"multi_match",
			MultiMatch([]string{"message", "subject"}, "this is a test").Type(BestFields),
			Map{
				"multi_match": Map{
					"query":  "this is a test",
					"type":   "best_fields",
					"fields": []string{"message", "subject"},
				},
			},
		},
		{
			"multi_match",
			MultiMatch([]string{"message", "subject"}, "this is a test").Type(BestFields).TieBreaker(0.3),
			Map{
				"multi_match": Map{
					"query":       "this is a test",
					"type":        "best_fields",
					"fields":      []string{"message", "subject"},
					"tie_breaker": 0.3,
				},
			},
		},
		{
			"multi_match",
			MultiMatch([]string{"message", "subject"}, "this is a test").Type(BestFields).TieBreaker(0.3).Operator(OperatorAnd),
			Map{
				"multi_match": Map{
					"query":       "this is a test",
					"type":        "best_fields",
					"fields":      []string{"message", "subject"},
					"tie_breaker": 0.3,
					"operator":    "and",
				},
			},
		},
		{
			"multi_match",
			MultiMatch([]string{"message", "subject"}, "").OmitValue(""),
			nil,
		},
	})
}
