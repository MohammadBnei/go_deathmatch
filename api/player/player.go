package player

type Player struct {
	Name     string   `json:"name"`
	Kill     int      `json:"kill"`
	Death    int      `json:"death"`
	Position Position `json:"position"`
}

type Position struct {
	X int
	Y int
}
