package channels

import (
	"fmt"
	"time"
)

func ExecuteChannels() {
	fmt.Println("executing channels")

	after := time.After(time.Millisecond * 2000)

	myMsg := createMessage("some message from external source")

	handleChannels(myMsg, after)
}

func createMessage(msg string) chan string {
	res := make(chan string)

	go func() {
		defer close(res)
		time.Sleep(time.Millisecond * 1000)
		res <- msg
	}()

	return res

}

func handleChannels(myMsg <-chan string, saveMessage <-chan time.Time) {
	for {

		select {
		case val := <-myMsg:
			printMessage(val)
			time.Sleep(time.Millisecond * 1000)
		case <-saveMessage:
			savingMessage()
			return
		default:
			waitForMessage()
			time.Sleep(time.Millisecond * 500)
		}

	}
}

func printMessage(msg string) {
	if msg != "" {
		fmt.Println("Printing message:")
		fmt.Println(msg)
		fmt.Println("End of the message")
	}
}

func savingMessage() {
	fmt.Println("Message saved")
}

func waitForMessage() {
	fmt.Println("Waiting for message")
}
