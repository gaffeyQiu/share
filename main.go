package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println(`Hello`)
		wg.Done()
	}()
	go func() {
		fmt.Println(`World`)
		wg.Done()
	}()
	wg.Wait()
}
