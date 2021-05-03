package messenger

type Message struct {
	Topic   string
	Content string
}

type Topic interface {
	Answer(message string) []string // TODO change to Message
	Promote() string
	Name() string
}
