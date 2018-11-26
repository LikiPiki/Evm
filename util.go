package main

import (
	"math/rand"
)

func GetRand(percent int) int {
	// 1 - memory
	// 2 - registr
	num := rand.Int() % 100
	if num < percent {
		return 1
	}
	return 2
}

func createClc(n, tp int) []int {
	sl := make([]int, 0)
	for i := 0; i < n; i++ {
		sl = append(sl, tp)
	}
	return sl
}

func getCommandType() int {
	return GetRand(commandType * 100)
}

func getMemOrReg() int {
	return GetRand(memoryVer * 100)
}
