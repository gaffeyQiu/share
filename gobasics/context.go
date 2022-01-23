package main

import (
	"context"
	"fmt"
	"sync"
)

func Rpc(ctx context.Context, url string) error {
	result := make(chan string)
	err := make(chan error)

	go func() {
		success := true
		if url == "b.com" {
			success = false
		}
		if success {
			result <- url
		} else {
			err <- fmt.Errorf("%s has err", url)
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-err:
		return e
	case <-result:
		fmt.Printf("success url: %s\n", url)
		return nil
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	urlList := []string{
		"a.com",
		"b.com",
		"c.com",
		"d.com",
	}

	var wg sync.WaitGroup
	wg.Add(len(urlList))

	for _, u := range urlList {
		go func(u string) {
			defer wg.Done()

			err := Rpc(ctx, u)
			if err != nil {
				fmt.Println(err)
				cancel()
			}
		}(u)
	}

	wg.Wait()
}
