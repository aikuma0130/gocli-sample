package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gocli-sample2"
	app.Usage = "go cli sample application"
	app.Action = func(c *cli.Context) error {
		fmt.Println("gocli sample2!!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
