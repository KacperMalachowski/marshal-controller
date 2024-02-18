package station

type Definition struct {
	Hills []Hill `json:"hills"`
}

type Hill struct {
	Signal    string `json:"signal"`
	Repeaters string `json:"repeaters"`
}
