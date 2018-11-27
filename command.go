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
	c.Start = clc
	if clc == 0 {
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
	start := len(c.Clc) + 1
	for !fl {
		find := true
		for i, _ := range mem {
			for _, cmd := range cmds {
				if cmd.In(start+i) == true {
					curClc := cmd.Clc[start+i-cmd.Start]
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
			start++
		}
	}
	clcInCurrent := start - c.Start
	c.AppendCommand(clcInCurrent-len(c.Clc), CLC_SPACE)
	c.Clc = append(c.Clc, mem...)
	res := createClc(4, 4)
	c.Clc = append(c.Clc, res...)
	c.Clc = append(c.Clc, 3)
	c.SetEnd()
}

func (c *Command) SetFirstCommand(cmdType, memOrReg, clc int) {
	if clc == 0 {
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
		c.End = len(c.Clc)
	}
}

func (c *Command) SetEnd() {
	c.End = c.Start + len(c.Clc) - 1
}

func (c Command) In(clc int) bool {
	if (c.Start <= clc) && (clc <= c.End) {
		return true
	}
	return false
}

func (c *Command) AppendCommand(cmdCount, cmdType int) {
	command := createClc(cmdCount, cmdType)
	c.Clc = append(c.Clc, command...)
}

func (c *Command) AppendCommands(command ...int) {
	c.Clc = append(c.Clc, command...)
}

func (c Command) FormatString() (str string) {
	for i := 0; i < c.Start; i++ {
		str += "-"
	}
	for _, el := range c.Clc {
		if el == 0 {
			str += "-"
		} else {
			str += fmt.Sprintf("%d", el)
		}
	}
	return str
}
