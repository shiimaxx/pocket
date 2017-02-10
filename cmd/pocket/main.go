package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "auth",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "consumer_key, c",
					Usage: "CONSUMER KEY",
				},
			},
			Action: func(c *cli.Context) error {
				if c.String("consumer_key") == "" {
					fmt.Println("Missing consumer_key")
					os.Exit(2)
				}
				accessToken, err := Authentication(c.String("consumer_key"))
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				fmt.Println(accessToken)
				return nil
			},
		},
	}
	app.Run(os.Args)
}
