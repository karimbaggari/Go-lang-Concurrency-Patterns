package main

import (
	"fmt"
	"time"
)

func someFunc(number string) {
	fmt.Println("the numb", number)
}

func main() {
	go someFunc("2")
	time.Sleep(time.Second * 2)
	fmt.Println("hello world")

}