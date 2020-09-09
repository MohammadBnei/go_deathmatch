package player

type Player struct {
	Name  string `json:"name"`
	Kill  int    `json:"kill"`
	Death int    `json:"death"`
}

type Kill struct {
	Killer string `json:"killer"`
	Body   string `json:"body"`
}
