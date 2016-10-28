package rockpick

import (
	"github.com/urfave/cli"
	"github.com/tecbot/gorocksdb"
	"errors"
	"fmt"
)

func RockPickEntry(c *cli.Context) error {
	database := c.String("database")

	if database == nil || database == "" {
		return errors.New("No database was provided.")
	}

	query := c.String("query")

	opts := gorocksdb.NewDefaultOptions()
	db, openErr := gorocksdb.OpenDb(opts, database)

	if openErr != nil {
		return openErr
	}

	switch query {
	case nil || "":
		return displayAll(db)
	default:
		return nil
	}
}

func displayAll(db *gorocksdb.DB) error {
	ro := gorocksdb.NewDefaultReadOptions()
	iter := db.NewIterator(ro)
	iter.SeekToFirst()

	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		fmt.Println(iter.Key(), iter.Value())
	}
}
