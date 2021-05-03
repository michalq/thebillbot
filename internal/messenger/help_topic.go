package messenger

import (
	"regexp"
)

type HelpTopic struct {
	msgs *Messenger
}

func NewHelpTopic(msgs *Messenger) *HelpTopic {
	return &HelpTopic{msgs}
}

func (h *HelpTopic) Answer(message string) []string {

	pattern := regexp.MustCompile(`help`)
	if !pattern.MatchString(message) {
		return []string{}
	}
	answers := make([]string, 0, len(h.msgs.Topics()))
	for _, topic := range h.msgs.Topics() {
		answers = append(answers, topic.Promote())
	}
	return answers
}

func (h *HelpTopic) Promote() string {
	return "Type help to show available topics."
}

func (h *HelpTopic) Name() string {
	return "help"
}
