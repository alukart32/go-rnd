package template

import (
	"strings"
	"testing"
)

type testMessanger struct {
	MessageReciever
}

func (tm *testMessanger) Message() string {
	return "tests"
}

func TestTemplate_Execute1(t *testing.T) {
	messanger := testMessanger{}
	messageTemplate := MessagingTemplate{}

	if !strings.Contains(messageTemplate.Execute(&messanger), "tests") {
		t.Error("expect: tests")
	}
}

func TestTemplate_Execute2(t *testing.T) {
	messanger := Messanger{}
	messageTemplate := MessagingTemplate{}

	if messageTemplate.Execute(&messanger) != "hello world template" {
		t.Error("expect: hello world template")
	}
}
