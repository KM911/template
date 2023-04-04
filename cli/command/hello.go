package command

import (
	"github.com/urfave/cli/v2"
)

var Hello = &cli.Command{
	Name:    "hello",
	Aliases: []string{"h"},
	Usage:   "hello world",
	Action:  HelloAction,
}

func HelloAction(c *cli.Context) error {
	if len(c.Args().Slice()) == 0 {
		println("hello")
		return nil
	}
	if len(c.Args().Slice()) == 1 {
		println("hello world")
		return nil
	}
	for _, value := range c.Args().Slice() {
		println(value)
	}

	return nil
}
