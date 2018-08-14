package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var language string

	app := cli.NewApp()
	app.Name = "gocli-sample2"
	app.Usage = "go cli sample application"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang",
			Value:       "english",
			Usage:       "language for this app",
			Destination: &language,
		},
	}

	app.Action = func(c *cli.Context) error {
		if language == "japanese" {
			fmt.Println("ごーしえるあい、さんぷる！！")
		} else {
			fmt.Println("gocli sample2!!")
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
