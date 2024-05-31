package main

import (
	"fmt"
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

	select {
		case msgFromMyChannel := <- myChannel:
            fmt.Println("the message", msgFromMyChannel)
		case msgFromAnotherChannel := <- anotherChannel:
            fmt.Println("the message", msgFromAnotherChannel)
        default:
            fmt.Println("no message", myChannel)
	}

}