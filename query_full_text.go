// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_full_text
// @Date: 2021/11/3 10:16

package esq

type Operator string
type ZeroTerms string
type MultiMatchType string
type Rewrite string

const (
	OperatorOr  Operator = "OR"
	OperatorAnd Operator = "And"

	ZeroTermsNone ZeroTerms = "none"
	ZeroTermsAll  ZeroTerms = "all"

	BestFields   MultiMatchType = "best_fields"
	MostField    MultiMatchType = "most_fields"
	CrossField   MultiMatchType = "cross_fields"
	Phrase       MultiMatchType = "phrase"
	PhrasePrefix MultiMatchType = "phrase_prefix"

	ScoringBoolean        Rewrite = "scoring_boolean"
	ConstantScore         Rewrite = "constant_score"
	ConstantScoreBoolean  Rewrite = "constant_score_boolean"
	TopTermsBlendedFreqsN Rewrite = "top_terms_blended_freqs_N"
	TopTermsBoostN        Rewrite = "top_terms_boost_N"
	TopTermsN             Rewrite = "top_terms_N"
)

//Match Query
type MatchQuery struct {
	field      string
	parameters MatchQueryParameters
}

type MatchQueryParameters struct {
	Query                           string    `json:"query"`
	Analyzer                        string    `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool      `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzziness                       string    `json:"fuzziness,omitempty"`
	MaxExpansions                   int32     `json:"max_expansions,omitempty"`
	PrefixLength                    int32     `json:"prefix_length,omitempty"`
	FuzzyTranspositions             bool      `json:"fuzzy_transpositions,omitempty"`
	FuzzyRewrite                    string    `json:"fuzzy_rewrite,omitempty"`
	Lenient                         bool      `json:"lenient,omitempty"`
	Operator                        Operator  `json:"operator,omitempty"`
	MinimumShouldMatch              string    `json:"minimum_should_match,omitempty"`
	ZeroTermsQuery                  ZeroTerms `json:"zero_terms_query,omitempty"`
}

func Match(field string, query string) *MatchQuery {
	return &MatchQuery{
		field: field,
		parameters: MatchQueryParameters{
			Query: query,
		},
	}
}

func (q *MatchQuery) Analyzer(v string) *MatchQuery {
	q.parameters.Analyzer = v
	return q
}

func (q *MatchQuery) AutoGenerateSynonymsPhraseQuery(v bool) *MatchQuery {
	q.parameters.AutoGenerateSynonymsPhraseQuery = v
	return q
}

func (q *MatchQuery) Fuzziness(v string) *MatchQuery {
	q.parameters.Fuzziness = v
	return q
}

func (q *MatchQuery) MaxExpansions(v int32) *MatchQuery {
	q.parameters.MaxExpansions = v
	return q
}

func (q *MatchQuery) PrefixLength(v int32) *MatchQuery {
	q.parameters.PrefixLength = v
	return q
}

func (q *MatchQuery) FuzzyTranspositions(v bool) *MatchQuery {
	q.parameters.FuzzyTranspositions = v
	return q
}

func (q *MatchQuery) FuzzyRewrite(v string) *MatchQuery {
	q.parameters.FuzzyRewrite = v
	return q
}

func (q *MatchQuery) Lenient(v bool) *MatchQuery {
	q.parameters.Lenient = v
	return q
}

func (q *MatchQuery) Operator(v Operator) *MatchQuery {
	q.parameters.Operator = v
	return q
}

func (q *MatchQuery) MinimumShouldMatch(v string) *MatchQuery {
	q.parameters.MinimumShouldMatch = v
	return q
}

func (q *MatchQuery) ZeroTermsQuery(v ZeroTerms) *MatchQuery {
	q.parameters.ZeroTermsQuery = v
	return q
}

func (q *MatchQuery) Map() Map {
	return NewMap("match", NewMap(q.field, q.parameters))
}

//Match Phrase Query
type MatchPhraseQuery struct {
	field      string
	parameters MatchPhraseQueryParameters
}

type MatchPhraseQueryParameters struct {
	Query          string    `json:"query"`
	Analyzer       string    `json:"analyzer,omitempty"`
	ZeroTermsQuery ZeroTerms `json:"zero_terms_query,omitempty"`
	Slop           int32     `json:"slop,omitempty"`
}

func MatchPhrase(field string, query string) *MatchPhraseQuery {
	return &MatchPhraseQuery{
		field: field,
		parameters: MatchPhraseQueryParameters{
			Query: query,
		},
	}
}

func (q *MatchPhraseQuery) Analyzer(v string) *MatchPhraseQuery {
	q.parameters.Analyzer = v
	return q
}

func (q *MatchPhraseQuery) ZeroTermsQuery(v ZeroTerms) *MatchPhraseQuery {
	q.parameters.ZeroTermsQuery = v
	return q
}

func (q *MatchPhraseQuery) Slop(v int32) *MatchPhraseQuery {
	q.parameters.Slop = v
	return q
}

func (q *MatchPhraseQuery) Map() Map {
	return NewMap("match_phrase", NewMap(q.field, q.parameters))
}

//Match Phrase Prefix Query
type MatchPhrasePrefixQuery struct {
	field      string
	parameters MatchPhrasePrefixQueryParameters
}

type MatchPhrasePrefixQueryParameters struct {
	Query          string    `json:"query"`
	Analyzer       string    `json:"analyzer,omitempty"`
	MaxExpansions  int32     `json:"max_expansions,omitempty"`
	Slop           int32     `json:"slop,omitempty"`
	ZeroTermsQuery ZeroTerms `json:"zero_terms_query,omitempty"`
}

func MatchPhrasePrefix(field string, query string) *MatchPhrasePrefixQuery {
	return &MatchPhrasePrefixQuery{
		field: field,
		parameters: MatchPhrasePrefixQueryParameters{
			Query: query,
		},
	}
}

func (q *MatchPhrasePrefixQuery) MaxExpansions(v int32) *MatchPhrasePrefixQuery {
	q.parameters.MaxExpansions = v
	return q
}

func (q *MatchPhrasePrefixQuery) Analyzer(v string) *MatchPhrasePrefixQuery {
	q.parameters.Analyzer = v
	return q
}

func (q *MatchPhrasePrefixQuery) ZeroTermsQuery(v ZeroTerms) *MatchPhrasePrefixQuery {
	q.parameters.ZeroTermsQuery = v
	return q
}

func (q *MatchPhrasePrefixQuery) Slop(v int32) *MatchPhrasePrefixQuery {
	q.parameters.Slop = v
	return q
}

func (q *MatchPhrasePrefixQuery) Map() Map {
	return NewMap("match_phrase_prefix", NewMap(q.field, q.parameters))
}

//Multi-match Query
//The fuzziness parameter cannot be used with the cross_fields type.
type MultiMatchQuery struct {
	parameters MultiMatchQueryParameters
}

type MultiMatchQueryParameters struct {
	Query                           string         `json:"query"`
	Fields                          []string       `json:"fields,omitempty"`
	Type                            MultiMatchType `json:"type,omitempty"`
	Operator                        Operator       `json:"operator,omitempty"`
	MinimumShouldMatch              string         `json:"minimum_should_match ,omitempty"`
	TieBreaker                      float32        `json:"tie_breaker,omitempty"`
	Analyzer                        string         `json:"analyzer,omitempty"`
	Boost                           float32        `json:"boost,omitempty"`
	AutoGenerateSynonymsPhraseQuery bool           `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Fuzziness                       string         `json:"fuzziness,omitempty"`
	MaxExpansions                   int32          `json:"max_expansions,omitempty"`
	PrefixLength                    int32          `json:"prefix_length,omitempty"`
	FuzzyTranspositions             bool           `json:"fuzzy_transpositions,omitempty"`
	Lenient                         bool           `json:"lenient,omitempty"`
	ZeroTermsQuery                  ZeroTerms      `json:"zero_terms_query,omitempty"`
	CutoffFrequency                 float32        `json:"cutoff_frequency,omitempty"`
	Rewrite                         Rewrite        `json:"rewrite,omitempty"`
}

func MultiMatch(fields []string, query string) *MultiMatchQuery {
	return &MultiMatchQuery{
		parameters: MultiMatchQueryParameters{
			Query:  query,
			Fields: fields,
		},
	}
}

func (q *MultiMatchQuery) Type(t MultiMatchType) *MultiMatchQuery {
	q.parameters.Type = t
	return q
}

func (q *MultiMatchQuery) Operator(o Operator) *MultiMatchQuery {
	q.parameters.Operator = o
	return q
}

func (q *MultiMatchQuery) MinimumShouldMatch(v string) *MultiMatchQuery {
	q.parameters.MinimumShouldMatch = v
	return q
}

func (q *MultiMatchQuery) TieBreaker(v float32) *MultiMatchQuery {
	q.parameters.TieBreaker = v
	return q
}

func (q *MultiMatchQuery) Analyzer(v string) *MultiMatchQuery {
	q.parameters.Analyzer = v
	return q
}

func (q *MultiMatchQuery) Boost(v float32) *MultiMatchQuery {
	q.parameters.Boost = v
	return q
}

func (q *MultiMatchQuery) AutoGenerateSynonymsPhraseQuery(v bool) *MultiMatchQuery {
	q.parameters.AutoGenerateSynonymsPhraseQuery = v
	return q
}

func (q *MultiMatchQuery) Fuzziness(v string) *MultiMatchQuery {
	q.parameters.Fuzziness = v
	return q
}

func (q *MultiMatchQuery) MaxExpansions(v int32) *MultiMatchQuery {
	q.parameters.MaxExpansions = v
	return q
}

func (q *MultiMatchQuery) PrefixLength(v int32) *MultiMatchQuery {
	q.parameters.PrefixLength = v
	return q
}

func (q *MultiMatchQuery) FuzzyTranspositions(v bool) *MultiMatchQuery {
	q.parameters.FuzzyTranspositions = v
	return q
}

func (q *MultiMatchQuery) Lenient(v bool) *MultiMatchQuery {
	q.parameters.Lenient = v
	return q
}

func (q *MultiMatchQuery) ZeroTermsQuery(v ZeroTerms) *MultiMatchQuery {
	q.parameters.ZeroTermsQuery = v
	return q
}

func (q *MultiMatchQuery) CutoffFrequency(v float32) *MultiMatchQuery {
	q.parameters.CutoffFrequency = v
	return q
}

func (q *MultiMatchQuery) Rewrite(v Rewrite) *MultiMatchQuery {
	q.parameters.Rewrite = v
	return q
}

func (q *MultiMatchQuery) Map() Map {
	return NewMap("multi_match", q.parameters)
}
