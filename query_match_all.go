// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_match_all
// @Date: 2021/11/5 17:30

package esq

//Match All Query
type matchAllQuery struct {
	parameter matchAllQueryParameter
}

type matchAllQueryParameter struct {
	Boost float32 `json:"boost,omitempty"`
}

func MatchAll() *matchAllQuery {
	return &matchAllQuery{}
}

func (q *matchAllQuery) Boost(v float32) *matchAllQuery {
	q.parameter.Boost = v
	return q
}

func (q *matchAllQuery) Map() Map {
	return NewMap("match_all", q.parameter)
}

//Match None Query
type matchNoneQuery struct{}

func MatchNone() *matchNoneQuery {
	return &matchNoneQuery{}
}

func (q *matchNoneQuery) Map() Map {
	return NewMap("match_none", q)
}
