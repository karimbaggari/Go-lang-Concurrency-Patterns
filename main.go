package main

import (
	"time"
)

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		anotherChannel <- "data 2"
	}()
	go func() {
		myChannel <- "data"
	}()
	
	time.Sleep(time.Second * 2)
}
