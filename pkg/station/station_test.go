package station_test

import (
	"reflect"
	"testing"

	"github.com/kacpermalachowski/marshal-controller/pkg/station"
)

func TestParseStationDefinition(t *testing.T) {
	tc := []struct {
		Name      string
		Data      string
		Expected  station.Definition
		ExpectErr bool
	}{
		{
			Name:      "No hills",
			Data:      `hills:`,
			ExpectErr: true,
		},
		{
			Name: "Single hill and signal",
			Data: `hills:
  - signal: Tr1`,
			Expected: station.Definition{
				Hills: []station.Hill{
					{
						Signal: "Tr1",
					},
				},
			},
		},
		{
			Name: "Single hill with repeaters",
			Data: `hills:
  - signal: Tr1
    repeaters:
      - Tr2`,
			Expected: station.Definition{
				Hills: []station.Hill{
					{
						Signal: "Tr1",
						Repeaters: []string{
							"Tr2",
						},
					},
				},
			},
		},
		{
			Name: "Multiple hill with repeaters",
			Data: `hills:
  - signal: Tr1
    repeaters:
      - Tr2
  - signal: Tr3
    repeaters:
      - Tr4`,
			Expected: station.Definition{
				Hills: []station.Hill{
					{
						Signal: "Tr1",
						Repeaters: []string{
							"Tr2",
						},
					},
					{
						Signal: "Tr3",
						Repeaters: []string{
							"Tr4",
						},
					},
				},
			},
		},
	}

	for _, c := range tc {
		t.Run(c.Name, func(t *testing.T) {
			def, err := station.ParseStationDefinition([]byte(c.Data))
			if err != nil && !c.ExpectErr {
				t.Errorf("Unexpected error occured: %s", err)
			}
			if c.ExpectErr && err == nil {
				t.Error("Expected error, but no one occured")
			}

			if !reflect.DeepEqual(c.Expected, def) {
				t.Errorf("Expected %v, but got %v", c.Expected, def)
			}
		})
	}
}
