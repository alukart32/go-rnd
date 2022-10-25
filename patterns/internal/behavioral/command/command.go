package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Command interface {
	Execute()
}

type ConsoleOutput struct {
	msg      string
	receiver *Receiver
}

func (c ConsoleOutput) Execute() {
	c.receiver.ConsolePrint(c.msg)
}

type FileOutput struct {
	msg      string
	receiver *Receiver
}

func (c FileOutput) Execute() {
	c.receiver.FilePrint(c.msg)
}

type Receiver struct {
}

func (r *Receiver) ConsolePrint(msg string) {
	fmt.Println(msg)
}

func (r *Receiver) FilePrint(msg string) {
	f, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buff := bufio.NewWriter(f)
	defer buff.Flush()
	_, err = buff.WriteString(msg)
	if err != nil {
		log.Fatal(err)
	}
}

type Invoker struct {
	commands []Command
}

func (i *Invoker) Add(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) Pop() {
	if len(i.commands) != 0 {
		i.commands[len(i.commands)] = nil
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) Execute() {
	for _, command := range i.commands {
		command.Execute()
	}
}
