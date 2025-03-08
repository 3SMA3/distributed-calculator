package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Agent started")
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
