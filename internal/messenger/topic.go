package messenger

type Topic interface {
	Answer(message string) []string
}
