package main

import (
	"encoding/json"
)

func main() {
	doubleThread()
}

func getJson(a interface{}) string {
	marshal, _ := json.Marshal(a)

	return string(marshal)
}
