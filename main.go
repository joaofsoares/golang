package main

import (
	"fmt"
	"learn/book/cpt1"
	"learn/channels"
	"learn/mutex"
	"learn/routines"
	"os"
	"time"
)

func main() {
	fmt.Println("hello from main")

	// routines
	go routines.PrintWord("1st Routine")
	go routines.PrintWord("2st Routine")

	time.Sleep(time.Second)

	// channels
	wordChannel := make(chan string)
	go channels.WaitingWord(&wordChannel, "Async word")
	fmt.Println(<-wordChannel)

	res := <-channels.SendWord("Await word")
	fmt.Println(res)

	// Mutex
	counter := 0
	go mutex.LoopCounter(&counter, 10)

	// introduction
	cpt1.HelloWorld()
	cpt1.EchoParams(os.Args[1:])
	// cpt1.RunDups()
	cpt1.ReadFiles([]string{"./book/files/text.txt"})
	cpt1.ReadFilesOnly([]string{"./book/files/text.txt"})

	// GIF
	cpt1.Lissajour("lissajour")

	// fetch url
	cpt1.Fetch("http://gopl.io")
	cpt1.FetchAll([]string{"http://gopl.io", "http://godoc.org", "http://golang.org"})

}
