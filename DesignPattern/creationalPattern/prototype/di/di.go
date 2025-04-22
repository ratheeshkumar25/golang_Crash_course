package di

import "github.com/ratheeshkumar25/prototype/publisher"

func Init() {
	publishers := publisher.NewPublisher()
	run(publishers, "Hi Welcome to prototype pattern")
	publisher2 := publishers.Clone().(publisher.IPublisher)
	run(publisher2, "exits")

}

func run(publisher publisher.IPublisher, message string) {
	publisher.Publish(message)
}
