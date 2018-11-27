package main

import (
	"fmt"
)

type InputReader struct {
	verRegister     []float32
	verCommandsType []float32
	countTime       []int
	countMemoryTime []int
}

func NewReader() InputReader {
	return InputReader{
		verRegister:     []float32{0.9, 0.8, 0.6},
		verCommandsType: []float32{0.9, 0.7, 0.5},
		countTime:       []int{4, 8, 16},
		countMemoryTime: []int{2, 5, 10},
	}
}

func (r *InputReader) Read() {
	fmt.Print("Введите количество комманд: ")
	fmt.Scanf("%d", &commandsCount)
	showVariantsFloat("Выберите вероятность типа команды", r.verCommandsType...)
	commandType = inputSelectFloats(r.verCommandsType...)
	showVariantsFloat("Выберите вероятность регистровой адрессации", r.verRegister...)
	memoryVer = inputSelectFloats(r.verRegister...)
	showVariantsInts("Выберите время обращения к памяти", r.countMemoryTime...)
	commandMemoryCount = inputSelectInts(r.countMemoryTime...)
	showVariantsInts("Выберите время вычисления", r.countTime...)
	commandCounting = inputSelectInts(r.countTime...)
}

func inputSelectFloats(variants ...float32) float32 {
	selected := 0
	fmt.Scanf("%d", &selected)
	return variants[selected-1]
}

func inputSelectInts(variants ...int) int {
	selected := 0
	fmt.Scanf("%d", &selected)
	return variants[selected-1]
}

func showVariantsFloat(message string, variants ...float32) {
	fmt.Println(message)
	for index, variant := range variants {
		fmt.Print(index+1, " - ", variant, "\n")
	}
}

func showVariantsInts(message string, variants ...int) {
	fmt.Println(message)
	for index, variant := range variants {
		fmt.Print(index+1, " - ", variant, "\n")
	}
}
