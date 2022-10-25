package command

import "testing"

func TestCommandMsgPrint(t *testing.T) {
	msg := "command pattern test"
	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.Add(ConsoleOutput{receiver: receiver, msg: msg})
	invoker.Add(FileOutput{receiver: receiver, msg: msg})

	invoker.Execute()
}
