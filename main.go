package main

import (
	"fmt"
	"learn/channels"
	"learn/routines"
	"time"
)

func main() {
	fmt.Println("hello from main")

	go routines.PrintWord("1st Routine")
	go routines.PrintWord("2st Routine")

	time.Sleep(time.Second)

	wordChannel := make(chan string)
	go channels.WaitingWord(&wordChannel, "Async word")
	fmt.Println(<-wordChannel)

	res := <-channels.SendWord("Await word")
	fmt.Println(res)

}
