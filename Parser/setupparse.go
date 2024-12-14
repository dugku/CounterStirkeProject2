package Parser

import (
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs/common"
	"log"
	"os"
	"time"
)

type parser struct {
	parser demoinfocs.Parser
}

type overview struct {
	tick    int
	Players []Player
}

type Point struct {
	X, Y, Z float64
}

type Player struct {
	Name                string
	ID                  int
	Team                common.Team
	Coord               Point
	LastAlivePosition   Point
	ViewDirectionX      float32
	ViewDirectionY      float32
	FlashDuration       time.Duration
	FlashTimeRemaining  time.Duration
	Inventory           []common.EquipmentType
	ActiveWeapon        common.EquipmentType
	Health              int16
	Armor               int16
	Money               int16
	Kills               int16
	Deaths              int16
	Assists             int16
	IsAlive             bool
	IsDefusing          bool
	IsOnNormalElevation bool
	HasHelmet           bool
	HasDefuseKit        bool
	HasBomb             bool
}

var currentRound int

func (p *parser) startParse() error {
	f, e := os.Open("C:\\Users\\Mike\\Desktop\\CounterStirkeProject2\\9ine-vs-permitta-m1-mirage.dem")
	check(e)
	defer f.Close()

	p.parser = demoinfocs.NewParser(f)
	defer p.parser.Close()

	for {
		parsed, err := p.parser.ParseNextFrame()
		check(err)
		if !parsed {
			break
		}

		gs := p.parser.GameState()

		active_player := make([]Player, 0, 10)

		for _, participant := range gs.Participants().Playing() {
			if participant.IsAlive() {
				// Create a common.Player instance with the necessary details.
				player := Player{
					Name: participant.Name,
					Coord: Point{
						X: participant.Position().X,
						Y: participant.Position().Y,
						Z: participant.Position().Z,
					},
				}

				// Append the player to the alivePlayers slice.
				active_player = append(active_player, player)
			}
		}

	}

	e = p.parser.ParseToEnd()
	check(e)

	return nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
