package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(trap([]int{0, 1, 5, 4, 2, 3, 5, 4}))
}

func getJson(a interface{}) string {
	marshal, _ := json.Marshal(a)

	return string(marshal)
}
