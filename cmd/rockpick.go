package main

import (
	"github.com/urfave/cli"
	"github.com/chris-tomich/rock-pick"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "rockpick"
	app.Usage = "Rock Pick is an extensible CLI tool for querying a RocksDB database"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "database, d",
			Usage: "Location of the database to open",
		},
		cli.StringFlag{
			Name: "query, q",
			Usage: "The keys to display",
		},
	}
	app.Action = rockpick.RockPickEntry
	app.Run(os.Args)
}
