package main

import (
	"fmt"
	"sync"
)

// 使用双线程打印出1-100的数据
// 通过设定数据量，和使用group的方式进行处理即可
// 注意waitgroup不能使用非指针的方式传入，这样会影响count的技术方式
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	go PrintFunc(wg, "线程1", 1, 50)
	go PrintFunc(wg, "线程2", 51, 100)
	wg.Wait()

	fmt.Println("that's finish is ok")
}

func PrintFunc(wg *sync.WaitGroup, name string, start, end int64) {
	for i := int64(start); i <= end; i++ {
		fmt.Println(fmt.Sprintf("current:[%s] out [%d]", name, i))
		wg.Done()
	}
}
