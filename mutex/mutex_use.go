package mutex

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

var mu sync.Mutex

type SafeCounter struct {
	Counts map[string]int
	Mux    *sync.Mutex
}

func (sc *SafeCounter) inc(key string) {
	sc.Mux.Lock()
	defer sc.Mux.Unlock()
	sc.slowIncrement(key)
}

func (sc *SafeCounter) val(key string) int {
	sc.Mux.Lock()
	defer sc.Mux.Unlock()
	return sc.Counts[key]
}

func (sc *SafeCounter) slowIncrement(key string) {
	tmpCounter := sc.Counts[key]

	time.Sleep(time.Millisecond)

	tmpCounter++

	sc.Counts[key] = tmpCounter
}

type EmailTest struct {
	Email string
	Count int
}

func ExecuteEmailTest(sc *SafeCounter, emails []EmailTest) {
	emailsMap := make(map[string]struct{})

	var wg sync.WaitGroup

	for _, email := range emails {
		emailsMap[email.Email] = struct{}{}

		for i := 0; i < email.Count; i++ {
			wg.Add(1)

			go func(email EmailTest) {
				sc.inc(email.Email)
				wg.Done()
			}(email)
		}
	}

	wg.Wait()

	sortedEmails := make([]string, 0, len(emails))

	for email := range emailsMap {
		sortedEmails = append(sortedEmails, email)
	}

	sort.Strings(sortedEmails)

	for _, email := range sortedEmails {
		fmt.Printf("Email: %s has %d emails\n", email, sc.val(email))
	}
}

func incCounter(counter *int) {
	mu.Lock()
	defer mu.Unlock()
	*counter++
}

func LoopCounter(counter *int, end int) {

	for *counter < end {
		fmt.Printf("counter = %d\n", *counter)
		incCounter(counter)
	}

	fmt.Println("counter =", *counter)
}
