package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmtJson(isValidV2("([)]"))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
