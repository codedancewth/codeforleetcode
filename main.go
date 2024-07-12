package main

import (
	"awesomeProject/dynamic"
	"encoding/json"
	"fmt"
)

func main() {
	fmtJson(dynamic.Trap([]int{1}))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
