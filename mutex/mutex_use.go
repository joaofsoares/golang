package mutex

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func incCounter(counter *int) {
	mu.Lock()
	*counter++
	mu.Unlock()
}

func LoopCounter(counter *int, end int) {

	for *counter < end {
		fmt.Printf("counter = %d\n", *counter)
		incCounter(counter)
	}

	fmt.Println("counter =", counter)
}
