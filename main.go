package main

import (
    "fmt"
    "time"
)

// Order struct representing an order in the online store
type Order struct {
    ID     int
    Amount float64
}

// processOrders simulates processing orders and sending them to a 'done' channel
func processOrders(orders <-chan Order, done chan<- Order) {
    for order := range orders {
        // Simulate processing time
        time.Sleep(time.Second)
        fmt.Printf("Processed order ID %d\n", order.ID)
        done <- order // Send the processed order to the 'done' channel
    }
}

func main() {
    orders := make(chan Order)
    done := make(chan Order)

    // Start the order processing goroutine
    go processOrders(orders, done)

    // Simulate receiving orders
    for i := 1; i <= 5; i++ {
        orders <- Order{ID: i, Amount: float64(i * 10)}
        fmt.Printf("Received order ID %d\n", i)
    }

    // Close the 'orders' channel to signal that no more orders will be sent
    close(orders)

    // Wait for all orders to be processed
    for i := 1; i <= 5; i++ {
        processedOrder := <-done
        fmt.Printf("Completed order ID %d\n", processedOrder)
	}
}