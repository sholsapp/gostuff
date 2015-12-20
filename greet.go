package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}
	app.Action = func(c *cli.Context) {
		name := "someone"
		if len(c.Args()) > 0 {
			name = c.Args()[0]
		}
		if c.String("lang") == "spanish" {
			println("Hola,", name)
		} else {
			println("Hello,", name)
		}
	}
	app.Run(os.Args)
}
