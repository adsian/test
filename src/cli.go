package main

import (
	"fmt"
	"os"
	"strconv"
)

const Usage = `
./process <TotalCount> <GoCurrency>
`

type CLI struct {
	*Request
}

func (c *CLI)Run() {
	cmds := os.Args

	if len(cmds) < 3 {
		fmt.Println(Usage)
		return
	}

	GlobalConfig.TotalCount, _ = strconv.ParseInt(cmds[1], 10, 64)
	GlobalConfig.GoCurrency, _ = strconv.ParseInt(cmds[2], 10, 64)

	c.Process()
}