package di

import "github.com/ratheeshkumar25/prototype/publisher"

func Init() {
	publisher := publisher.MyPublish{}
	hello := "welcome to prototype pattern"

	run(&publisher, hello)
}

func run(publisher publisher.IPublisher, message string) {
	publisher.Publish(message)
}
