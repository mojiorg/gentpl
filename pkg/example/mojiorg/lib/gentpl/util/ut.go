package util

import (
	"encoding/json"
)

func JsonString(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(b)
}
func Chk(err error) {
	if err != nil {
		panic(err)
	}
}
