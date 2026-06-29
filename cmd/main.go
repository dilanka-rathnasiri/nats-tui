package main

import (
	"fmt"
	"log"
	"nats-ui/internal/connections"
	"nats-ui/internal/services"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS running on localhost
	nc, err := connections.InitNatsClient()
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}
	defer connections.CloseNatsClient(nc)

	done := make(chan bool)

	// Listen to the nats subject "dilanka"
	sub, err := services.ReceiveMsg(nc, "dilanka", func(m *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(m.Data))
		if string(m.Data) == "AA" {
			done <- true
		}
	})

	if err != nil {
		log.Fatalf("Error subscribing: %v", err)
	}

	defer sub.Unsubscribe()

	// Send message "AA" to the nats subject "dilanka"
	err = services.SendMsg(nc, "dilanka", []byte("AA"))
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	// Wait to ensure message is received
	select {
	case <-done:
		fmt.Println("Success")
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}
}
