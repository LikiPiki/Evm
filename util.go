package main

import (
	"math/rand"
	"time"
)

const (
	MEMORY_TYPE   = 2
	REGISTER_TYPE = 1
	FIRST_TYPE    = 1
	SECOND_TYPE   = 2
)

func GetRand(percent int) int {
	// 1 - memory
	// 2 - registr
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Int() % 100
	if num < percent {
		return MEMORY_TYPE
	}
	return REGISTER_TYPE
}

func createClc(n, tp int) []int {
	sl := make([]int, 0)
	for i := 0; i < n; i++ {
		sl = append(sl, tp)
	}
	return sl
}

func getCommandType() int {
	return GetRand(int(commandType * 100))
}

func getMemOrReg() int {
	return GetRand(int(memoryVer * 100))
}

func CountArrangeTime(commands []Command) (result int) {
	for _, cmd := range commands {
		result += len(cmd.Clc)
	}
	return result / len(commands)
}
