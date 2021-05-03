package messenger

import (
	"regexp"
)

type HelpTopic struct {
}

func NewHelpTopic() *HelpTopic {
	return &HelpTopic{}
}

func (h *HelpTopic) Answer(message string) []string {

	pattern := regexp.MustCompile(`help`)
	if !pattern.MatchString(message) {
		return []string{}
	}

	return []string{"Type help to show available topics."}
}

func (h *HelpTopic) Promote() string {
	return "Type help to show available topics."
}
