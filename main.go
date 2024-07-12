package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmtJson(searchInsert([]int{1}, 7))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
