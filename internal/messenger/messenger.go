package messenger

type Messenger struct {
	topics []Topic
}

func NewMessenger() *Messenger {
	return &Messenger{}
}

func (m *Messenger) AddTopic(topic Topic) {
	m.topics = append(m.topics, topic)
}

func (m *Messenger) Topics() []Topic {
	return m.topics
}

func (m *Messenger) Scan(message string) []Message {
	answers := make([]Message, 0)
	for _, topic := range m.topics {
		if topicAnswers := topic.Answer(message); len(topicAnswers) > 0 {
			answers = append(answers, topicAnswers...)
		}
	}
	return answers
}
