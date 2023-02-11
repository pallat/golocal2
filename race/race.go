package main

import "fmt"

func main() {
	race()
}

func race() {
	var i int

	go func() { i++ }()
	if i == 0 {
		fmt.Printf("confirm 0, i = %d\n", i)
	} else {
		fmt.Printf("confirm not 0,i = %d\n", i)
	}
}
