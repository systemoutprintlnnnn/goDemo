package main

import (
	"fmt"
	"sync"
)

var myMap = make(map[int]int)
var mu sync.Mutex

func main() {
	for i := 1; i <= 500; i++ {
		go test(i)
	}

	//time.Sleep(1 * time.Second)
	fmt.Println(len(myMap))
	//for v := range myMap {
	//	fmt.Println(v)
	//}

}

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	mu.Lock()
	myMap[n] = res
	mu.Unlock()
}
