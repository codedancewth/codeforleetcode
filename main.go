package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("s")
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
