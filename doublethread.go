package main

import (
	"fmt"
	"sync"
)

var num = 1 // 开始的数字

// doubleThread 双携程交替打印100
func doubleThread() {
	w := sync.WaitGroup{}
	c := make(chan int, 1)

	w.Add(2)
	c <- 1
	go Out(w, c, true)
	go Out(w, c, false)

	w.Wait()
	close(c)
}

func Out(w sync.WaitGroup, c chan int, trans bool) {
	defer w.Done()

	for num < 100 {
		fmt.Println(fmt.Sprintf("go-%v [%d]", trans, <-c))
		num++
		c <- num
	}
}
