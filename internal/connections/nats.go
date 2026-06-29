package connections

import "github.com/nats-io/nats.go"

// InitNatsClient establishes a connection to NATS using the default URL.
// Returns the NATS connection or an error if connection fails.
func InitNatsClient() (*nats.Conn, error) {
	return nats.Connect(nats.DefaultURL)
}

// CloseNatsClient closes the NATS connection.
func CloseNatsClient(nc *nats.Conn) {
	nc.Close()
}
