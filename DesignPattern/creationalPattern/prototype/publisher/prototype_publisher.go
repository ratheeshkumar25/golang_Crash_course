package publisher

import "log"

type IPublisher interface {
	Publish(message string)
}

type MyPublish struct {
}

func (m *MyPublish) Publish(message string) {
	log.Print(message)
}
