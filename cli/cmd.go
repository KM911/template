package main

import (
	"cli/command"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:     "cli template",
		Usage:    "A cli template",
		Commands: command.CommandList,
		//Flags:    flag.FlagList,
		Action: func(c *cli.Context) error {
			println("need a valid command")
			//cli.ShowAppHelp(c)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
