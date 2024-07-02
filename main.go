package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := []int{45192, 0, -659, -52359, -99225, -75991, 0, -15155, 27382, 59818, 0, -30645, -17025, 81209, 887, 64648}
	moveZeroes(a)
	marshal, _ := json.Marshal(a)
	fmt.Println(string(marshal))
}
