package main

import "fmt"

func main() {
	ch := make(chan struct{})

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- struct{}{}
	fmt.Println(<-ch)
}
