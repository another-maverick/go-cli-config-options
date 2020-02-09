package main

import(
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "UsingStandardCliOptions"
	app.Usage = "Using a standard package for CLI options and args"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "name, n",
			Value: "Steph",
			Usage: "Name of the best basketball player",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("%s is the best basketball player! \n", name)
		return nil
	}
	app.Run(os.Args)
}
