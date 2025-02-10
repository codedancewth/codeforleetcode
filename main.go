package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(InsertionSort([]int{3, 1, 2}))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
