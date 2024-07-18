package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmtJson(isValid("()[]"))
}

func fmtJson(a interface{}) {
	marshal, _ := json.Marshal(a)

	fmt.Println(string(marshal))
}
