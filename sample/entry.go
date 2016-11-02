package sample

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"github.com/chris-tomich/rock-pick/sample/json"
	"fmt"
)

func BuildSampleDbEntry(c *cli.Context) error {
	filepath := c.String("json")
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return err
	}

	var people *json.People

	people, err = json.LoadPeople(data)

	if err != nil {
		return err
	}

	for _, person := range people.People {
		fmt.Println(person)
	}

	return nil
}
