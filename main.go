package main

import (
	"fmt"
)

const (
	COMMANDS = 10
	// вероятность для того что используется регистровая адрессация
	memoryVer    = 0.6
	memoryAccess = 5
	commandCount = 4
	commandType  = 0.7
)

func main() {

	// reading from console
	// reader := NewReader()
	// reader.Read()

	commands := make([]Command, 0)

	for i := 0; i < 30; i++ {
		cmd := NewCommand(i)
		cmd.SetCommand(commands, i)
		commands = append(commands, cmd)
	}

	for _, el := range commands {
		// fmt.Println(el)
		fmt.Println(el.FormatString())
	}
}
