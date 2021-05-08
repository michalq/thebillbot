package messenger

import (
	"context"
	"regexp"
	"strings"
)

type HelpTopic struct {
	msgs *Messenger
}

func NewHelpTopic(msgs *Messenger) *HelpTopic {
	return &HelpTopic{msgs}
}

func (h *HelpTopic) Answer(_ context.Context, message string) []Message {

	pattern := regexp.MustCompile(`help`)
	if !pattern.MatchString(strings.ToLower(message)) {
		return []Message{}
	}
	answers := make([]Message, 0, len(h.msgs.Topics()))
	for _, topic := range h.msgs.Topics() {
		answers = append(answers, Message{Content: topic.Promote()})
	}
	return answers
}

func (h *HelpTopic) Promote() string {
	return "Type help to show available topics."
}

func (h *HelpTopic) Name() string {
	return "help"
}
