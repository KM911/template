package cli

import (
	"demo/command"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	//config.LoadUserConfig()
	app := &cli.App{
		Name:     "cli template",
		Usage:    "A cli template",
		Commands: command.CommandList,
		Action: func(c *cli.Context) error {
			println("need a valid command")
			cli.ShowAppHelp(c)
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		println(color.Red.Sprint("error: ") + err.Error())
	}
}
