package station

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Definition struct {
	Hills []Hill `json:"hills" yaml:"hills"`
}

type Hill struct {
	Signal    string   `json:"signal" yaml:"signal"`
	Repeaters []string `json:"repeaters" yaml:"repeaters"`
}

func ParseStationDefinition(data []byte) (Definition, error) {
	var station Definition
	err := yaml.Unmarshal(data, &station)
	if err != nil {
		return Definition{}, err
	}

	if len(station.Hills) == 0 {
		return Definition{}, fmt.Errorf("the program is unusable without hills defined")
	}

	return station, err
}
