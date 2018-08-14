package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gocli-sample"
	app.Usage = "go cli sample application"
	app.Action = func(c *cli.Context) error {
		fmt.Println("gocli sample!!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
