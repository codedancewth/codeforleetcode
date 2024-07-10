package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmtJson(trap([]int{1}))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
