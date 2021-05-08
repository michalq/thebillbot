package messenger

import "context"

type Message struct {
	Topic   string
	Content string
}

type Topic interface {
	Answer(ctx context.Context, message string) []Message
	Promote() string
	Name() string
}
