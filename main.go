package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	for _, index := range groupAnagrams2([]string{}) {
		marshal, _ := json.Marshal(index)
		if strings.Contains(string(marshal), "rankle") {
			fmt.Println(string(marshal))
		}

	}

	fmt.Println(getKey("sdad"))
}
