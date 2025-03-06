package main

import (
	"sync"

	"github.com/3SMA3/distributed-calculator/internal/agent"
)

func main() {
	computingPower := 3
	var wg sync.WaitGroup

	for i := 0; i < computingPower; i++ {
		wg.Add(1)
		go agent.Worker(&wg)
	}

	wg.Wait()
}
