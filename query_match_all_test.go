// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: query_match_all_test.go
// @Date: 2021/12/23 10:25

package esq

import (
	"testing"
)

func TestMatchAll(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"match_all",
			MatchAll().Boost(1.2),
			Map{
				"match_all": Map{
					"boost": 1.2,
				},
			},
		},
	})
}

func TestMatchNone(t *testing.T) {
	runMapTest(t, []mapTest{
		{
			"match_none",
			MatchNone(),
			Map{
				"match_none": Map{},
			},
		},
	})
}
