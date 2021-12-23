// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: es_test.go
// @Date: 2021/11/17 14:41

package esq

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMap_Set(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		m    Map
		args []args
	}{
		{
			"MapSet",
			Map{
				"key":  "value",
				"key2": "value2",
			},
			[]args{
				{
					key:   "key",
					value: "value",
				},
				{
					key:   "key2",
					value: "value2",
				},
			},
		},
		{
			"MapSet",
			Map{
				"key": "value",
			},
			[]args{
				{
					key:   "key",
					value: "value",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := make(Map)
			for _, ttt := range tt.args {
				m.Set(ttt.key, ttt.value)
			}
			if !reflect.DeepEqual(m, tt.m) {
				t.Errorf("expected %v, got %v", m, tt.m)
			}
		})
	}
}

func TestNewMap(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		wantM Map
	}{
		{
			"NewMap",
			args{
				key:   "key",
				value: "value",
			},
			Map{
				"key": "value",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotM := NewMap(tt.args.key, tt.args.value); !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("expected %v, got %v", gotM, tt.wantM)
			}
		})
	}
}

type mapTest struct {
	name     string
	mappable Mappable
	m        Map
}

func runMapTest(t *testing.T, mt []mapTest) {
	for _, v := range mt {
		t.Run(v.name, func(t *testing.T) {
			mp := v.mappable.Map()
			got, exp, ok := sameJSON(mp, v.m)
			if !ok {
				t.Errorf("expected %s, got %s", exp, got)
			}
		})
	}
}

func sameJSON(a, b Map) (aJSON, bJSON []byte, ok bool) {
	aJSON, aErr := json.Marshal(a)
	var m = new(Map)
	_ = json.Unmarshal(aJSON, &m)
	aJSON, aErr = json.Marshal(m)
	bJSON, bErr := json.Marshal(b)

	if aErr != nil || bErr != nil {
		return aJSON, bJSON, false
	}

	ok = reflect.DeepEqual(aJSON, bJSON)
	return aJSON, bJSON, ok
}
