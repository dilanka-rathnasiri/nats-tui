package services

import "github.com/nats-io/nats.go"

func SendMsg(nc *nats.Conn, subject string, msg []byte) error {
	return nc.Publish(subject, msg)
}

func ReceiveMsg(nc *nats.Conn, subject string, msgHandler nats.MsgHandler) (*nats.Subscription, error) {
	return nc.Subscribe(subject, msgHandler)
}
