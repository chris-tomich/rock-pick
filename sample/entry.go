package sample

import (
	"github.com/urfave/cli"
	"io/ioutil"
	sampleJson "github.com/chris-tomich/rock-pick/sample/json"
	"fmt"
	"encoding/json"
	"github.com/tecbot/gorocksdb"
	"github.com/satori/go.uuid"
)

func BuildSampleDbEntry(c *cli.Context) error {
	filepath := c.String("json")
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	var people *sampleJson.People

	people, err = sampleJson.LoadPeople(data)

	if err != nil {
		return err
	}

	databasePath := c.String("database")

	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)
	defer opts.Destroy()

	var db *gorocksdb.DB

	db, err = gorocksdb.OpenDb(opts, databasePath)
	defer db.Close()

	if err != nil {
		return err
	}

	wo := gorocksdb.NewDefaultWriteOptions()

	for _, person := range people.People {
		var personJson []byte

		personJson, err = json.Marshal(person)

		if err != nil {
			panic(err)
		}

		var personUuid uuid.UUID

		personUuid, err = uuid.FromString(person.UniqueId)

		if err != nil {
			panic(err)
		}

		db.Put(wo, personUuid.Bytes(), personJson)

		fmt.Println(string(personJson))
	}

	return nil
}
