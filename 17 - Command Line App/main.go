package main

import (
	"linha-comando-app/app"
	"os"
)

func main() {
	cli := app.Generate()

	cli.Run(os.Args)
}
