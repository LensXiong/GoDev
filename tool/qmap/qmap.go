package qmap

import (
	"strconv"
)

type QM map[string]interface{}

func (this QM) String(key string) (s string) {
	if t, exist := this[key]; exist && t != nil {
		switch t.(type) {
		case float64:
			s = strconv.FormatFloat(t.(float64), 'f', 6, 64)
		case int64:
			s = strconv.FormatInt(t.(int64), 10)
		case int32:
			s = strconv.FormatInt(int64(t.(int32)), 10)
		case int:
			s = strconv.Itoa(t.(int))
		default:
			s, _ = t.(string)
		}
	}
	return
}
