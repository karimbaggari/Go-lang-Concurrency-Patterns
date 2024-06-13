package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	sensorData := make(chan int)
	userCommands := make(chan string)

	// Simulate sensor data and user commands
	go func() {
		for i := 0; i < 5; i++ {
			sensorData <- i
			time.Sleep(1 * time.Second)
		}
		close(sensorData)
	}()

	go func() {
		commands := []string{"start", "stop", "pause"}
		for _, cmd := range commands {
			userCommands <- cmd
			time.Sleep(2 * time.Second)
		}
		close(userCommands)
	}()

	// Use a for-select loop to handle data from both channels
	for {
		select {
		case data, ok := <-sensorData:
			if !ok {
				sensorData = nil
				fmt.Println("Sensor data channel closed")
			} else {
				fmt.Printf("Received sensor data: %d\n", data)
			}
		case cmd, ok := <-userCommands:
			if !ok {
				userCommands = nil
				fmt.Println("User commands channel closed")
			} else {
				fmt.Printf("Received user command: %s\n", cmd)
			}
		}

		// Exit the loop if both channels are closed
		if sensorData == nil && userCommands == nil {
			break
		}
	}

	fmt.Println("All channels are closed, exiting program.")
}
