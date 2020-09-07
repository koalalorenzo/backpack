package pkg

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ValuesType is the Type for our Values
type ValuesType map[string]interface{}

// ValuesFromFile reads a path and returns a ValuesType
// Use this function to allow different formats (yaml, json, msgpack)
func ValuesFromFile(p string) (v ValuesType, err error) {
	vb, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	v = ValuesType{}
	err = yaml.Unmarshal(vb, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
