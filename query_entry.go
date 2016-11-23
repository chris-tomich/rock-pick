package rockpick

import (
	"github.com/tecbot/gorocksdb"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"encoding/json"
	sampleJson "github.com/chris-tomich/rock-pick/sample/json"
)

func QueryEntry(databasePath string, query string) error {
	if databasePath == "" {
		return errors.New("No database was provided.")
	}

	opts := gorocksdb.NewDefaultOptions()
	db, openErr := gorocksdb.OpenDb(opts, databasePath)

	if openErr != nil {
		return openErr
	}

	switch query {
	case "*":
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
		keyBytes := iter.Key().Data()
		sampleBytes := iter.Value().Data()

		sampleData := new(sampleJson.Person)

		key, err := uuid.FromBytes(keyBytes)

		if err != nil {
			panic(err)
		}

		json.Unmarshal(sampleBytes, sampleData)

		fmt.Println(key, *sampleData)
	}

	return nil
}