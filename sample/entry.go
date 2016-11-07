package sample

import (
	"github.com/urfave/cli"
	"io/ioutil"
	sample "github.com/chris-tomich/rock-pick/sample/json"
	"fmt"
	"encoding/json"
)

func BuildSampleDbEntry(c *cli.Context) error {
	filepath := c.String("json")
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	var people *sample.People

	people, err = sample.LoadPeople(data)

	if err != nil {
		return err
	}

	

	for _, person := range people.People {
		personJson, err := json.Marshal(person)

		if err != nil {
			panic(err)
		}

		fmt.Println(string(personJson))
	}

	return nil
}
