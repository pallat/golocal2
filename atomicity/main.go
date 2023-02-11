package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var limit = 10
var wg sync.WaitGroup

func main() {
	wg.Add(limit)

	for i := 0; i < limit; i++ {
		go race(i)
		// go atom(i)
	}

	wg.Wait()
}

var r int

func race(n int) {
	r = n
	fmt.Println(r, "<-", n)
	wg.Done()
}

var a atomic.Int64

func atom(n int) {
	a.Store(int64(n))
	fmt.Println(a.Load(), "<-", n)
	wg.Done()
}
