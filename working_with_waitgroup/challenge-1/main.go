package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage(wg *sync.WaitGroup) {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {
	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage(&wg)

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage(&wg)

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage(&wg)

	wg.Wait()

}
