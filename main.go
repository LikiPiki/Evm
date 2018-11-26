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

	commands := make([]Command, 0)

	for i := 0; i < 2; i++ {
		cmd := NewCommand(i)
		cmd.SetCommand(commands, i)
		commands = append(commands, cmd)
	}

	for _, el := range commands {
		fmt.Println(el)
	}

}
