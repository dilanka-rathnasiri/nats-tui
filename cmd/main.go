package main

import (
	"fmt"
	"log/slog"
	"nats-ui/internal/connections"
	"nats-ui/internal/services"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

// main is the entry point for the NATS UI application.
// It connects to NATS, subscribes to a subject, sends a message,
// and waits for the message to be received.
func main() {
	// Set up slog with logfmt format
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	// Connect to NATS running on localhost
	nc, err := connections.InitNatsClient()
	if err != nil {
		slog.Error("Error connecting to NATS", "error", err)
		os.Exit(1)
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
		slog.Error("Error subscribing", "error", err)
		os.Exit(1)
	}

	defer sub.Unsubscribe()

	// Send message "AA" to the nats subject "dilanka"
	err = services.SendMsg(nc, "dilanka", []byte("AA"))
	if err != nil {
		slog.Error("Error sending message", "error", err)
		os.Exit(1)
	}

	// Wait to ensure message is received
	select {
	case <-done:
		fmt.Println("Success")
	case <-time.After(2 * time.Second):
		fmt.Println("Timed out")
	}
}
