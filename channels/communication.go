package channels

import (
	"time"
)

func SendWord(word string) <-chan string {

	res := make(chan string)

	go func() {
		defer close(res)
		time.Sleep(2 * time.Second)
		res <- word
	}()

	return res
}

func WaitingWord(wordChannel *chan string, word string) {
	time.Sleep(time.Second)
	*wordChannel <- word
}
