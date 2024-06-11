package main

import (
	"fmt"
	"time"
)


func main() {
	normalOrders := make(chan string)
	priorityOrders := make(chan string)

	// Goroutine to generate normal orders
	go func() {
		orders := []string{"Order1", "Order2", "Order3"}
		for _, order := range orders {
			fmt.Println("Generating normal order:", order)
			normalOrders <- order
			time.Sleep(1 * time.Second) // Simulate time to generate order
		}
		close(normalOrders)
	}()

	// Goroutine to generate priority orders
	go func() {
		priorityOrdersList := []string{"PriorityOrder1", "PriorityOrder2"}
		for _, order := range priorityOrdersList {
			fmt.Println("Generating priority order:", order)
			priorityOrders <- order
			time.Sleep(3 * time.Second) // Simulate time to generate priority order
		}
		close(priorityOrders)
	}()

	// Goroutine to process orders
	go func() {
		for {
			select {
			case order, ok := <-normalOrders:
				if ok {
					fmt.Println("Processing normal order:", order)
					time.Sleep(2 * time.Second) // Simulate time to process order
				} else {
					normalOrders = nil // Set to nil to avoid blocking
				}
			case priorityOrder, ok := <-priorityOrders:
				if ok {
					fmt.Println("Processing priority order:", priorityOrder)
					time.Sleep(2 * time.Second) // Simulate time to process priority order
				} else {
					priorityOrders = nil // Set to nil to avoid blocking
				}
			}
			
			// Break if both channels are closed
			if normalOrders == nil && priorityOrders == nil {
				break
			}
		}
	}()

	// Allow time for goroutines to complete
	time.Sleep(20 * time.Second)
}
