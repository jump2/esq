// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_boolean
// @Date: 2021/11/1 11:53

package esq

type BoolQuery struct {
	must               []Mappable
	filter             []Mappable
	should             []Mappable
	mustNot            []Mappable
	minimumShouldMatch int8
	boost              float32
}

func Bool() *BoolQuery {
	return &BoolQuery{}
}

func (q *BoolQuery) Must(must ...Mappable) *BoolQuery {
	q.must = append(q.must, must...)

	return q
}

func (q *BoolQuery) Filter(filter ...Mappable) *BoolQuery {
	q.filter = append(q.filter, filter...)

	return q
}

func (q *BoolQuery) Should(should ...Mappable) *BoolQuery {
	q.should = append(q.should, should...)

	return q
}

func (q *BoolQuery) MustNot(mustNot ...Mappable) *BoolQuery {
	q.mustNot = append(q.mustNot, mustNot...)

	return q
}

func (q *BoolQuery) MinimumShouldMatch(v int8) *BoolQuery {
	q.minimumShouldMatch = v

	return q
}

func (q *BoolQuery) Boost(v float32) *BoolQuery {
	q.boost = v

	return q
}

func (q *BoolQuery) Map() Map {
	var data struct {
		Must               []Map   `json:"must,omitempty"`
		Filter             []Map   `json:"filter,omitempty"`
		Should             []Map   `json:"should,omitempty"`
		MustNot            []Map   `json:"must_not,omitempty"`
		MinimumShouldMatch int8    `json:"minimum_should_match,omitempty"`
		Boost              float32 `json:"boost,omitempty"`
	}

	data.MinimumShouldMatch = q.minimumShouldMatch
	data.Boost = q.boost

	if len(q.must) > 0 {
		data.Must = make([]Map, 0, len(q.must))
		for _, m := range q.must {
			v := m.Map()
			if v != nil {
				data.Must = append(data.Must, v)
			}
		}
	}

	if len(q.filter) > 0 {
		data.Filter = make([]Map, 0, len(q.filter))
		for _, m := range q.filter {
			v := m.Map()
			if v != nil {
				data.Filter = append(data.Filter, v)
			}
		}
	}

	if len(q.mustNot) > 0 {
		data.MustNot = make([]Map, 0, len(q.mustNot))
		for _, m := range q.mustNot {
			v := m.Map()
			if v != nil {
				data.MustNot = append(data.MustNot, v)
			}
		}
	}

	if len(q.should) > 0 {
		data.Should = make([]Map, 0, len(q.should))
		for _, m := range q.should {
			v := m.Map()
			if v != nil {
				data.Should = append(data.Should, v)
			}
		}
	}

	return NewMap("bool", data)
}
