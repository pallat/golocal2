package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const total = 200
const workers = 10

var wg = sync.WaitGroup{}

func main() {
	success := 0
	fail := 0

	start := time.Now()
	wg.Add(total)

	chN := make(chan int)
	chErr := make(chan error)

	for i := 0; i < workers; i++ {
		go worker(chN, chErr)
	}

	go func() {
		for i := 0; i < total; i++ {
			chN <- i
		}
	}()
	for i := 0; i < total; i++ {
		if err := <-chErr; err != nil {
			fail++
		} else {
			success++
		}
	}
	wg.Wait()

	log.Println("success", success)
	log.Println("fail", fail)
	log.Println("duration", time.Since(start))
}

func worker(chN chan int, chErr chan error) {
	for {
		n := <-chN
		chErr <- service(n)
	}
}

var ErrNotOK = errors.New("status is not OK")

func service(n int) error {
	defer wg.Done()
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/%d", n))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ErrNotOK
	}

	return nil
}
