package main

import (
	"fmt"
)

type Command struct {
	Start int
	End   int
	Clc   []int
}

func NewCommand(start int) Command {
	c := Command{
		Start: start,
		Clc:   make([]int, 0),
	}
	return c
}

func (c *Command) SetCommand(cmds []Command, clc int) {
	cmdType := getCommandType()
	memOrReg := getMemOrReg()
	if clc == 0 {
		c.Start = 0
		c.Clc = append(c.Clc, 1, 2)
		if memOrReg == 1 {
			c.Clc = append(c.Clc, 2)
		} else {
			mem := createClc(8, 3)
			c.Clc = append(c.Clc, mem...)
		}
		if cmdType == 1 {
			c.Clc = append(c.Clc, 4)
		} else {
			res := createClc(4, 4)
			c.Clc = append(c.Clc, res...)
		}
		c.Clc = append(c.Clc, 5)
		c.End = len(c.Clc) - 1
		return
	}
	c.Clc = append(c.Clc, 1, 2)
	mem := createClc(4, 3)
	fl := false
	start := clc
	for !fl {
		find := true
		fmt.Println("------------------------------------")
		fmt.Println("start", start)
		fmt.Println("------------------------------------")
		for i, _ := range mem {
			for _, cmd := range cmds {
				fmt.Println("i is", start+i)
				fmt.Println(cmd)
				if cmd.In(start+i) == true {
					curClc := cmd.Clc[start-cmd.Start+i]
					fmt.Printf("Check %d %d \n", start-cmd.Start+i, curClc)
					if (curClc == 3) || (curClc == 5) {
						find = false
					}
				}
			}
		}
		if find {
			fl = true
			break
		} else {
			c.Clc = append(c.Clc, 0)
			start++
		}
	}
	c.Clc = append(c.Clc, mem...)
	res := createClc(4, 4)
	c.Clc = append(c.Clc, res...)
	c.Clc = append(c.Clc, 5)
	c.End = len(c.Clc) - 1
	fmt.Println("start is ", start)
}
func (c Command) SetEnd() {
	c.End = c.Start + len(c.Clc)
}

func (c Command) In(clc int) bool {
	if (c.Start <= clc) && (clc <= c.End) {
		return true
	}
	return false
}

func (c Command) Format() string {
	result := fmt.Sprintf("%d %d\n", c.Start, c.End)
	for _, el := range c.Clc {
		result += fmt.Sprintf("%d ", el)
	}
	return result
}
