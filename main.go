package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println(multiply("2", "3"))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
