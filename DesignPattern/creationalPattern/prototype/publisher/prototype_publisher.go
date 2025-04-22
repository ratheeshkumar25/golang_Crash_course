package publisher

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Iprototype interface {
	Clone() interface{}
}

type IPublisher interface {
	Iprototype
	Publish(message string)
}

type MyPublish struct {
	ID string
}

func (p MyPublish) Clone() interface{} {
	return MyPublish{
		ID: p.ID,
	}
}

func NewPublisher() *MyPublish {
	return &MyPublish{
		ID: uuid.New().String(),
	}
}

func (m MyPublish) Publish(message string) {
	fullMessage := fmt.Sprint("Publisher ", m.ID, ">", message)
	log.Print(fullMessage)
}
