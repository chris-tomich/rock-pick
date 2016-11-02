package json

import "encoding/json"

type Person struct {
	FirstName string
	Surname string
	Street string
	City string
	Postcode int
	State string
	Email string
	UniqueId string
}

type People struct {
	People []Person
}

func LoadPeople(data []byte) (*People, error) {
	p := &People{}

	err := json.Unmarshal(data, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
