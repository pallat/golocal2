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

var wg = sync.WaitGroup{}

func main() {
	success := 0
	fail := 0

	start := time.Now()
	wg.Add(total)

	for i := 0; i < total; i++ {
		go func() {
			if err := service(i); err != nil {
				fail++
			} else {
				success++
			}
		}()
	}
	wg.Wait()

	log.Println("success", success)
	log.Println("fail", fail)
	log.Println("duration", time.Since(start))
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
