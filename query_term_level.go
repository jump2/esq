// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_term
// @Date: 2021/11/1 15:17

package esq

import "reflect"

//Term Query
type TermQuery struct {
	field      string
	omitValue  interface{}
	parameters TermQueryParameters
}

type TermQueryParameters struct {
	Value interface{} `json:"value"`
	Boost float32     `json:"boost,omitempty"`
}

func Term(field string, value interface{}) *TermQuery {
	return &TermQuery{
		field: field,
		parameters: TermQueryParameters{
			Value: value,
		},
	}
}

func (q *TermQuery) OmitValue(v interface{}) *TermQuery {
	q.omitValue = v
	return q
}

func (q *TermQuery) Boost(v float32) *TermQuery {
	q.parameters.Boost = v
	return q
}

func (q *TermQuery) Map() Map {
	if ToStr(q.omitValue) == ToStr(q.parameters.Value) {
		return nil
	}
	return NewMap("term", NewMap(q.field, q.parameters))
}

//Terms Query
type TermsQuery struct {
	field     string
	value     []interface{}
	boost     float32
	omitValue interface{}
}

func Terms(field string, value ...interface{}) *TermsQuery {
	return &TermsQuery{
		field: field,
		value: value,
	}
}

func (q *TermsQuery) OmitValue(v interface{}) *TermsQuery {
	q.omitValue = v
	return q
}

func (q *TermsQuery) Boost(v float32) *TermsQuery {
	q.boost = v
	return q
}

func (q *TermsQuery) Map() Map {
	values := make([]interface{}, 0)
	if len(q.value) == 0 {
		return nil
	}
	v := reflect.ValueOf(q.value[0])
	t := v.Type()
	if t.Kind() == reflect.Slice {
		if v.Len() == 0 {
			return nil
		}

		for i := 0; i < v.Len(); i++ {
			values = append(values, v.Index(i).Interface())
		}
	} else {
		values = append(values, q.value...)
	}

	fvalue := make([]interface{}, 0)
	for _, v := range values {
		if ToStr(q.omitValue) != ToStr(v) {
			fvalue = append(fvalue, v)
		}
	}
	if len(fvalue) == 0 {
		return nil
	}

	termsQuery := NewMap(q.field, fvalue)
	if q.boost > 0 {
		termsQuery.Set("boost", q.boost)
	}
	return NewMap("terms", termsQuery)
}

//Terms Set Query
type TermsSetQuery struct {
	field      string
	omitValue  interface{}
	parameters TermsSetQueryParameters
}

type TermsSetQueryParameters struct {
	Terms                    []interface{} `json:"terms"`
	MinimumShouldMatchField  string        `json:"minimum_should_match_field,omitempty"`
	MinimumShouldMatchScript Map           `json:"minimum_should_match_script,omitempty"`
}

func TermsSet(field string, value ...interface{}) *TermsSetQuery {
	return &TermsSetQuery{
		field: field,
		parameters: TermsSetQueryParameters{
			Terms: value,
		},
	}
}

func (q *TermsSetQuery) OmitValue(v interface{}) *TermsSetQuery {
	q.omitValue = v
	return q
}

func (q *TermsSetQuery) MinimumShouldMatchField(field string) *TermsSetQuery {
	q.parameters.MinimumShouldMatchField = field
	return q
}

func (q *TermsSetQuery) MinimumShouldMatchScript(script string) *TermsSetQuery {
	q.parameters.MinimumShouldMatchScript = NewMap("source", script)
	return q
}

func (q *TermsSetQuery) Map() Map {
	values := make([]interface{}, 0)
	if len(q.parameters.Terms) == 0 {
		return nil
	}
	v := reflect.ValueOf(q.parameters.Terms[0])
	t := v.Type()
	if t.Kind() == reflect.Slice {
		if v.Len() == 0 {
			return nil
		}

		for i := 0; i < v.Len(); i++ {
			values = append(values, v.Index(i).Interface())
		}
	} else {
		values = append(values, q.parameters.Terms...)
	}

	fvalue := make([]interface{}, 0)
	for _, v := range values {
		if ToStr(q.omitValue) != ToStr(v) {
			fvalue = append(fvalue, v)
		}
	}
	if len(fvalue) == 0 {
		return nil
	}

	q.parameters.Terms = fvalue

	return NewMap("terms_set", NewMap(q.field, q.parameters))
}

//Wildcard Query
type WildcardQuery struct {
	field      string
	omitValue  string
	parameters WildcardQueryParameters
}

type WildcardQueryParameters struct {
	Value           string  `json:"value"`
	Boost           float32 `json:"boost,omitempty"`
	Rewrite         Rewrite `json:"rewrite,omitempty"`
	CaseInsensitive bool    `json:"case_insensitive,omitempty"`
}

func Wildcard(field string, value string) *WildcardQuery {
	return &WildcardQuery{
		field: field,
		parameters: WildcardQueryParameters{
			Value: value,
		},
	}
}

func (q *WildcardQuery) OmitValue(v string) *WildcardQuery {
	q.omitValue = v
	return q
}

func (q *WildcardQuery) Boost(v float32) *WildcardQuery {
	q.parameters.Boost = v
	return q
}

func (q *WildcardQuery) Rewrite(v Rewrite) *WildcardQuery {
	q.parameters.Rewrite = v
	return q
}

//add in 7.10.0
func (q *WildcardQuery) CaseInsensitive(v bool) *WildcardQuery {
	q.parameters.CaseInsensitive = v
	return q
}

func (q *WildcardQuery) Map() Map {
	if ToStr(q.omitValue) == ToStr(q.parameters.Value) {
		return nil
	}
	return NewMap("wildcard", NewMap(q.field, q.parameters))
}

//Regexp Query
type RegexpQuery struct {
	field      string
	omitValue  string
	parameters RegexpQueryParameters
}

type RegexpQueryParameters struct {
	Value                 string  `json:"value"`
	Flags                 string  `json:"flags,omitempty"`
	Rewrite               Rewrite `json:"rewrite,omitempty"`
	CaseInsensitive       bool    `json:"case_insensitive,omitempty"`
	MaxDeterminizedStates int32   `json:"max_determinized_states,omitempty"`
}

func Regex(field string, value string) *RegexpQuery {
	return &RegexpQuery{
		field: field,
		parameters: RegexpQueryParameters{
			Value: value,
		},
	}
}

func (q *RegexpQuery) OmitValue(v string) *RegexpQuery {
	q.omitValue = v
	return q
}

func (q *RegexpQuery) Flags(v string) *RegexpQuery {
	q.parameters.Flags = v
	return q
}

func (q *RegexpQuery) Rewrite(v Rewrite) *RegexpQuery {
	q.parameters.Rewrite = v
	return q
}

//add in 7.10.0
func (q *RegexpQuery) CaseInsensitive(v bool) *RegexpQuery {
	q.parameters.CaseInsensitive = v
	return q
}

func (q *RegexpQuery) MaxDeterminizedStates(v int32) *RegexpQuery {
	q.parameters.MaxDeterminizedStates = v
	return q
}

func (q *RegexpQuery) Map() Map {
	if ToStr(q.omitValue) == ToStr(q.parameters.Value) {
		return nil
	}
	return NewMap("regexp", NewMap(q.field, q.parameters))
}

//Range Query
type Relation string

const (
	RelationIntersects Relation = "INTERSECTS"
	RelationContains   Relation = "CONTAINS"
	RelationWithin     Relation = "WITHIN"
)

type RangeQuery struct {
	field      string
	omitValue  interface{}
	parameters RangeQueryParameters
}

type RangeQueryParameters struct {
	Gt       interface{} `json:"gt,omitempty"`
	Gte      interface{} `json:"gte,omitempty"`
	Lt       interface{} `json:"lt,omitempty"`
	Lte      interface{} `json:"lte,omitempty"`
	Format   string      `json:"format,omitempty"`
	Relation Relation    `json:"relation,omitempty"`
	TimeZone string      `json:"time_zone,omitempty"`
	Boost    float32     `json:"boost,omitempty"`
}

func Range(field string) *RangeQuery {
	return &RangeQuery{
		field: field,
	}
}

func (q *RangeQuery) OmitValue(v interface{}) *RangeQuery {
	q.omitValue = v
	return q
}

func (q *RangeQuery) Gt(v interface{}) *RangeQuery {
	q.parameters.Gt = v
	return q
}

func (q *RangeQuery) Gte(v interface{}) *RangeQuery {
	q.parameters.Gte = v
	return q
}

func (q *RangeQuery) Lt(v interface{}) *RangeQuery {
	q.parameters.Lt = v
	return q
}

func (q *RangeQuery) Lte(v interface{}) *RangeQuery {
	q.parameters.Lte = v
	return q
}

func (q *RangeQuery) Format(v string) *RangeQuery {
	q.parameters.Format = v
	return q
}

func (q *RangeQuery) Relation(v Relation) *RangeQuery {
	q.parameters.Relation = v
	return q
}

func (q *RangeQuery) TimeZone(v string) *RangeQuery {
	q.parameters.TimeZone = v
	return q
}

func (q *RangeQuery) Boost(v float32) *RangeQuery {
	q.parameters.Boost = v
	return q
}

func (q *RangeQuery) Map() Map {
	omitValue := ToStr(q.omitValue)
	if omitValue == ToStr(q.parameters.Gt) {
		q.parameters.Gt = nil
	}
	if omitValue == ToStr(q.parameters.Gte) {
		q.parameters.Gte = nil
	}
	if omitValue == ToStr(q.parameters.Lt) {
		q.parameters.Lt = nil
	}
	if omitValue == ToStr(q.parameters.Lte) {
		q.parameters.Lte = nil
	}
	if q.parameters == (RangeQueryParameters{}) || (q.parameters.Gt == nil && q.parameters.Gte == nil && q.parameters.Lt == nil && q.parameters.Lte == nil) {
		return nil
	}

	return NewMap("range", NewMap(q.field, q.parameters))
}

//Prefix Query
type PrefixQuery struct {
	field      string
	omitValue  string
	parameters PrefixQueryParameters
}

type PrefixQueryParameters struct {
	Value           string  `json:"value"`
	Rewrite         Rewrite `json:"rewrite,omitempty"`
	CaseInsensitive bool    `json:"case_insensitive,omitempty"`
}

func Prefix(field string, value string) *PrefixQuery {
	return &PrefixQuery{
		field: field,
		parameters: PrefixQueryParameters{
			Value: value,
		},
	}
}

func (q *PrefixQuery) OmitValue(v string) *PrefixQuery {
	q.omitValue = v
	return q
}

func (q *PrefixQuery) Rewrite(v Rewrite) *PrefixQuery {
	q.parameters.Rewrite = v
	return q
}

//add in 7.10.0
func (q *PrefixQuery) CaseInsensitive(v bool) *PrefixQuery {
	q.parameters.CaseInsensitive = v
	return q
}

func (q *PrefixQuery) Map() Map {
	if q.omitValue == ToStr(q.parameters.Value) {
		return nil
	}
	return NewMap("prefix", NewMap(q.field, q.parameters))
}

//IDs Query
type IdsQuery struct {
	Values []interface{} `json:"values"`
}

func Ids(values ...interface{}) *IdsQuery {
	return &IdsQuery{Values: values}
}

func (q *IdsQuery) Map() Map {
	return NewMap("ids", q)
}

//Fuzzy Query
type FuzzyQuery struct {
	field      string
	parameters FuzzyQueryParameters
}

type FuzzyQueryParameters struct {
	Value          string `json:"value"`
	Fuzziness      string `json:"fuzziness,omitempty"`
	MaxExpansions  int32  `json:"max_expansions,omitempty"`
	PrefixLength   int32  `json:"prefix_length,omitempty"`
	Transpositions bool   `json:"transpositions,omitempty"`
	Rewrite        string `json:"rewrite,omitempty"`
}

func Fuzzy(field string, value string) *FuzzyQuery {
	return &FuzzyQuery{
		field: field,
		parameters: FuzzyQueryParameters{
			Value: value,
		},
	}
}

func (q *FuzzyQuery) Fuzziness(v string) *FuzzyQuery {
	q.parameters.Fuzziness = v
	return q
}

func (q *FuzzyQuery) MaxExpansions(v int32) *FuzzyQuery {
	q.parameters.MaxExpansions = v
	return q
}

func (q *FuzzyQuery) PrefixLength(v int32) *FuzzyQuery {
	q.parameters.PrefixLength = v
	return q
}

func (q *FuzzyQuery) Transpositions(v bool) *FuzzyQuery {
	q.parameters.Transpositions = v
	return q
}

func (q *FuzzyQuery) Rewrite(v string) *FuzzyQuery {
	q.parameters.Rewrite = v
	return q
}

func (q *FuzzyQuery) Map() Map {
	return NewMap("fuzzy", NewMap(q.field, q.parameters))
}

//Exists Query
type ExistsQuery struct {
	Field string `json:"field"`
}

func Exists(field string) *ExistsQuery {
	return &ExistsQuery{field}
}

func (q *ExistsQuery) Map() Map {
	return NewMap("exists", q)
}
