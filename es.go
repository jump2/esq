// @Author: aaron
// @Email: 707230686@qq.com
// @Description:
// @File: es
// @Date: 2021/11/1 13:43

package esq

type Map map[string]interface{}

func NewMap(key string, value interface{}) (m Map) {
	m = make(Map)
	m.Set(key, value)

	return
}

func (m Map) Set(key string, value interface{}) {
	m[key] = value
}

type Mappable interface {
	Map() Map
}
