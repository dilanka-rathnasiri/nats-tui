package connections

import "github.com/nats-io/nats.go"

func InitNatsClient() (*nats.Conn, error) {
	return nats.Connect(nats.DefaultURL)
}

func CloseNatsClient(nc *nats.Conn) {
	nc.Close()
}
