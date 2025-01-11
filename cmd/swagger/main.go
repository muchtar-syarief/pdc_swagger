package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type IntEnum int
type UintEnum uint

type PayloadDataDD struct {
	Name string `json:"name" binding:"required,gte=6,lte=32"`
}

type MarkupValue float64

type ResponseData struct {
	Data  string      `json:"data"`
	Page  IntEnum     `json:"page"`
	Page2 UintEnum    `json:"page2"`
	Val   MarkupValue `json:"val"`
}

func main() {

	app := &cli.App{
		Name:  "Server Api",
		Usage: "Server Api",
		Commands: []*cli.Command{
			GinApiCli(),
			EchoApiCli(),
		},
		Action: func(*cli.Context) error {
			log.Println("Duar----->>>")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Panicln(err)
	}
}
