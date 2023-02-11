package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

const total = 200

func main() {
	start := time.Now()

	for i := 0; i < total; i++ {
		service(i)
	}

	// log.Println("success", success)
	// log.Println("fail", fail)
	log.Println("duration", time.Since(start))
}

var ErrNotOK = errors.New("status is not OK")

func service(n int) error {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8081/%d", n))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ErrNotOK
	}

	return nil
}
