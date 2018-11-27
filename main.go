package main

import (
	"fmt"
)

var (
	commandsCount = 10
	// вероятность для того что используется регистровая адрессация
	memoryVer   float32 = 0.6
	commandType float32 = 0.7
	// CLC counting by type 2
	commandCounting = 6
	// CLC counting for memory access
	commandMemoryCount = 3
)

func main() {

	// reading from console
	reader := NewReader()
	reader.Read()

	commands := make([]Command, 0)

	for i := 0; i < commandsCount; i++ {
		cmd := NewCommand(i)
		cmd.SetCommand(commands, i)
		commands = append(commands, cmd)
	}

	for _, el := range commands {
		fmt.Println(el.FormatString())
	}

	result := CountArrangeTime(commands)
	fmt.Println("---------------")
	fmt.Println("Calculation time: ", result)
}
