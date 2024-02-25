package station

type Definition struct {
	Hills []Hill `json:"hills" yaml:"hills"`
}

type Hill struct {
	Signal    string   `json:"signal" yaml:"signal"`
	Repeaters []string `json:"repeaters" yaml:"repeaters"`
}
