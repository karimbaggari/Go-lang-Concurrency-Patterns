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

	charChannel := make(chan string, 3) 

	chars := []string{"a","b","c"}

	for _, s := range chars {
		select {
			case charChannel <- string(s):
                fmt.Println("test", string(s))
            default:
                fmt.Println("channel is full")
		}
	}

	close(charChannel)

	for result := range charChannel { 
		fmt.Println(result)
	}

	go func () {
		for {
			select {
			default:
				fmt.Println("channel is empty")
			}
		}
	}()
	time.Sleep(time.Second * 2)
}
