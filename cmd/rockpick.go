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
		cli.StringFlag{
			Name: "file, f",
			Usage: "Location of the data file to use for a bulk upload",
		},
	}
	app.Action = RockPickEntry
	app.Run(os.Args)
}

func RockPickEntry(c *cli.Context) error {
	databasePath := c.String("database")
	query := c.String("query")
	filePath := c.String("file")

	if query != "" {
		return rockpick.QueryEntry(databasePath, query)
	} else if filePath != "" {
		return rockpick.BulkUploadEntry(databasePath, filePath)
	} else {
		return nil
	}
}
