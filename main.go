package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmtJson(trap([]int{2, 1, 2}))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
