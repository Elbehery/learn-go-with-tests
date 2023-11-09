package myapp

import (
	"bufio"
	"fmt"
	"io"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          in,
	}
}

func (c *CLI) PlayPoker() {
	scanner := bufio.NewScanner(c.in)
	scanner.Scan()
	in := scanner.Text()
	c.playerStore.RecordWin(in)
	fmt.Println(fmt.Sprintf("Hello %s", in))
}
