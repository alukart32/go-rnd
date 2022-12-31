package template

import "strings"

type (
	MessageReciever interface {
		Message() string
	}

	Template interface {
		stepOne() string
		stepThree() string
		Execute(mr MessageReciever) string
	}

	MessagingTemplate struct{}

	Messanger struct{}
)

func (t *MessagingTemplate) stepOne() string {
	return "hello"
}

func (t *MessagingTemplate) stepThree() string {
	return "template"
}

func (t *MessagingTemplate) Execute(mr MessageReciever) string {
	return strings.Join([]string{t.stepOne(), mr.Message(), t.stepThree()}, " ")
}

func (m *Messanger) Message() string {
	return "world"
}
