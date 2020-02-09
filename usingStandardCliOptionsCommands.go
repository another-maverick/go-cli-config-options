package main

import(
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Count up or down"
	app.Commands = []cli.Command{
		{
			Name:      "up",
			ShortName: "u",
			Usage:     "Count upwards",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "How long to count",
					Value: 20,
				},
			},
			Action: func(c *cli.Context) error {
				stop := c.Int("stop")
				if stop <= 0 {
					fmt.Println("Cannot have negative stop number")
				}
				for i := 0; i < stop; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:      "down",
			ShortName: "d",
			Usage:     "Where to start Counting down from",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "Where to start counting down from",
					Value: 20,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start < 0 {
					fmt.Println("the starting nuber to count down from cannot be less than 0")
				}
				for i := start; i > 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}


