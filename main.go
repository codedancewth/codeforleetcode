package main

import (
	"codeforleetcode/array"
	"encoding/json"
	"fmt"
)

func main() {
	a := []int{10, 7, 8, 9, 1, 5}
	array.QuickSortV2(a, 0, 5)
	fmtJson(a)
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
