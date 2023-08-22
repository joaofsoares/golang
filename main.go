package main

import (
	"fmt"
	"learn/book/cpt1"
	"learn/book/cpt2"
	"learn/book/cpt3"
	"learn/channels"
	"learn/mutex"
	"learn/routines"
	"os"
	"sync"
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

	channels.ExecuteChannels()

	// mutex
	counter := 0
	go mutex.LoopCounter(&counter, 10)

	sc := mutex.SafeCounter{
		Counts: make(map[string]int),
		Mux:    &sync.Mutex{},
	}

	emails := []mutex.EmailTest{
		{
			Email: "foo@bar.com",
			Count: 42,
		},
		{
			Email: "contact@bar.com",
			Count: 51,
		},
		{
			Email: "foo@bar.com",
			Count: 10,
		},
		{
			Email: "me@bar.com",
			Count: 21,
		},
		{
			Email: "contact@bar.com",
			Count: 10,
		},
	}

	mutex.ExecuteEmailTest(&sc, emails)

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

	// convert celsius to fahrenheit
	fahrenheit := cpt2.CToF(35)
	fmt.Printf("Fahrenheit: %.2fF\n", fahrenheit)

	celsius := cpt2.FToC(75)
	fmt.Printf("Celsius: %.2fC\n", celsius)

	// const weekday
	fmt.Println("Today is Monday:", cpt3.Monday)

}
