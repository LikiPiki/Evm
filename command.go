package main

import (
	"fmt"
)

const (
	CLC_SPACE    = 0
	CLC_KOP      = 1
	CLC_REGISTER = 2
	CLC_MEMORY   = 3
	CLC_COUNTING = 4
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
		c.Clc = append(c.Clc, CLC_KOP, CLC_REGISTER)
		if memOrReg == 1 {
			c.Clc = append(c.Clc, CLC_REGISTER)
		} else {
			mem := createClc(8, CLC_MEMORY)
			c.Clc = append(c.Clc, mem...)
		}
		if cmdType == 1 {
			c.Clc = append(c.Clc, CLC_COUNTING)
		} else {
			res := createClc(4, CLC_COUNTING)
			c.Clc = append(c.Clc, res...)
		}
		c.Clc = append(c.Clc, CLC_MEMORY)
		c.End = len(c.Clc) - 1
		return
	}
	c.Clc = append(c.Clc, 1, 2)
	mem := createClc(4, 3)
	fl := false
	start := clc
	for !fl {
		find := true
		for i, _ := range mem {
			for _, cmd := range cmds {
				if cmd.In(start+i) == true {
					curClc := cmd.Clc[start-cmd.Start+i]
					if curClc == 3 {
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
	c.Clc = append(c.Clc, 3)
	c.End = len(c.Clc) - 1
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

func (c Command) FormatString() (str string) {
	for i := 0; i < c.Start; i++ {
		str += "-"
	}
	for _, el := range c.Clc {
		str += fmt.Sprintf("%d", el)
	}
	return str
}
