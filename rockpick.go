package main

import (
	"os"
	"plugin"

	"github.com/chris-tomich/rock-pick/query"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rockpick"
	app.Usage = "Rock Pick is an extensible CLI tool for querying a RocksDB database"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "database, d",
			Usage: "Location of the database to open",
		},
		cli.StringFlag{
			Name:  "query, q",
			Usage: "The keys to display",
		},
		cli.StringFlag{
			Name:  "plugin, p",
			Usage: "Location of the plugin to use",
		},
	}
	app.Action = RockPickEntry
	app.Run(os.Args)
}

func RockPickEntry(c *cli.Context) error {
	databasePath := c.String("database")
	queryPath := c.String("query")
	pluginLocation := c.String("plugin")

	p, err := plugin.Open(pluginLocation)

	if err != nil {
		logrus.Fatal(err)
	}

	return query.Entry(p, databasePath, queryPath)
}
