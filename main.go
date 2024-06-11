package main

import (
	"fmt"
	"time"
)


func main() {
	unbufferedChannel := make(chan string)
	bufferedChan := make(chan string, 5)

	// Goroutine to generate tasks
	go func() {
		tasks := []string{"Order1", "Order2", "Order3", "Order4", "Order5", "Order6", "Order7", "Order8", "Order9", "Order10", "Order11"}
		for _, task := range tasks {
			fmt.Println("Generating task:", task)
			unbufferedChannel <- task
			time.Sleep(1 * time.Second) // Simulate time to generate task
		}
		close(unbufferedChannel)
	}()

	// Goroutine to process tasks
	go func() {
		for task := range unbufferedChannel {
			fmt.Println("Processing task:", task)
			time.Sleep(2 * time.Second) // Simulate time to process task
		}
	}()

	// Allow time for goroutines to complete
	time.Sleep(22 * time.Second)

	// Goroutine to generate tasks
	go func() {
		tasks := []string{"Order1", "Order2", "Order3", "Order4", "Order5", "Order6"}
		for _, task := range tasks {
			fmt.Println("Generating task:", task)
			bufferedChan <- task
			time.Sleep(1 * time.Second) // Simulate time to generate task
		}
		close(bufferedChan)
	}()

	// Goroutine to process tasks
	go func() {
		for task := range bufferedChan {
			fmt.Println("Processing task:", task)
			time.Sleep(2 * time.Second) // Simulate time to process task
		}
	}()

	// Allow time for goroutines to complete
	time.Sleep(20 * time.Second)
}
