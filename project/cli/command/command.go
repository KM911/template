package command

import (
	"github.com/urfave/cli/v2"
)

var CommandList = []*cli.Command{
	Init,
}

var Init = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "init config",
	Action:  InitAction,
}

func InitAction(c *cli.Context) error {
	println("init config")
	return nil
}
