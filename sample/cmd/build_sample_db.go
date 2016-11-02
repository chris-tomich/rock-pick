package main

import (
	"github.com/urfave/cli"
	"os"
	"github.com/chris-tomich/rock-pick/sample"
)

func main() {
	app := cli.NewApp()
	app.Name = "build_sample_db"
	app.Usage = "This will create a sample rocks DB with the provided JSON data."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "database, d",
			Usage: "Location of the database to create",
		},
		cli.StringFlag{
			Name: "json, j",
			Usage: "Location of the JSON data file to use",
		},
	}
	app.Action = sample.BuildSampleDbEntry
	app.Run(os.Args)
}
