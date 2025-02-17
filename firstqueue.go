package main

import (
	"fmt"
	"sort"
)

// Queue 优先队列
type Queue struct {
	Key   int64
	Power int64
}

var FirstQueue []*Queue

// Pop 出队
func Pop() *Queue {
	popQueue := FirstQueue[len(FirstQueue)-1]

	FirstQueue = FirstQueue[:len(FirstQueue)-1]
	return popQueue
}

func Push(q *Queue) {
	FirstQueue = append(FirstQueue, q)

	sort.Slice(FirstQueue, func(i, j int) bool {
		return FirstQueue[i].Power < FirstQueue[j].Power
	})
}

func Retest() {
	a := []int64{1, 4, 5, 6, 7, 2, 3, 87}
	for _, index := range a {
		Push(&Queue{
			Key:   index,
			Power: index,
		})
	}

	for i := range a {

		q := Pop()
		fmt.Println(fmt.Sprintf("i = %d; first %+v", i, q))
	}

	fmt.Println(fmt.Sprintf("FirstQueue %+v", FirstQueue))
}
