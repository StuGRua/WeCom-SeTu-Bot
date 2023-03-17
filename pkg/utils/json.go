package utils

import "encoding/json"

func Json(in interface{}) string {
	res, _ := json.Marshal(in)
	return string(res)
}
