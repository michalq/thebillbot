package messenger

type Message struct {
	Topic   string
	Content string
}

type Topic interface {
	Answer(message string) []Message
	Promote() string
	Name() string
}
