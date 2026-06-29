package services

import "github.com/nats-io/nats.go"

// SendMsg publishes a message to the specified NATS subject.
// Returns an error if publishing fails.
func SendMsg(nc *nats.Conn, subject string, msg []byte) error {
	return nc.Publish(subject, msg)
}

// ReceiveMsg subscribes to the specified NATS subject and registers a
// message handler. Returns the subscription or an error if subscription fails.
func ReceiveMsg(
	nc *nats.Conn,
	subject string,
	msgHandler nats.MsgHandler,
) (*nats.Subscription, error) {
	return nc.Subscribe(subject, msgHandler)
}
