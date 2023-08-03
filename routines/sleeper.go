package routines

import (
	"fmt"
	"time"
)

func PrintWord(word string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(word)
}
