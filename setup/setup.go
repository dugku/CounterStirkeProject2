package setup

type Player struct {
}

type NomalPoint struct {
	Tick int
	X    float64
	Y    float64
}

var PlayersMoveDat map[string][]NomalPoint

// initialize the player movement data
func init() {
	PlayersMoveDat = make(map[string][]NomalPoint)
}
