package main

import (
	app "github.com/dugku/CounterStirkeProject2/Draw"
	com "github.com/dugku/CounterStirkeProject2/setup"
	"github.com/markus-wa/demoinfocs-golang/v4/pkg/demoinfocs"
	"log"
	"os"
)

type parser struct {
	parse demoinfocs.Parser
}

type events struct {
	kills []string
}

var xmax = 3114.000244
var ymax = 2050.490234
var zmax = 1239.999878
var xmin = -4168.697266
var ymin = -3864.0
var zmin = -448.0

func (p *parser) ticks() error {
	f, e := os.Open("C:\\Users\\Mike\\Desktop\\CounterStirkeProject2\\demos\\falcons-vs-heroic-m1-mirage.dem")
	check(e)
	defer f.Close()

	p.parse = demoinfocs.NewParser(f)
	defer p.parse.Close()

	for {
		parsed, err1 := p.parse.ParseNextFrame()
		check(err1)
		if !parsed {
			break
		}
		frame := p.parse.CurrentFrame()
		gs := p.parse.GameState()

		for _, players := range gs.TeamCounterTerrorists().Members() {
			player := players
			playerName := players.Name
			posx := normal(player.Position().X, xmax, xmin)
			posy := normal(player.Position().Y, ymax, ymin)
			com.PlayersMoveDat[playerName] = []com.NormalPoint{
				{
					Tick: frame,
					X:    posx,
					Y:    posy,
				},
			}
		}
		/*
			for _, players := range gs.TeamTerrorists().Members() {
				player := players
				playerName := players.Name
				posx := normal(player.Position().X, xmax, xmin)
				posy := normal(player.Position().Y, ymax, ymin)

				playerData[playerName] = append(playerData[playerName], NormalizedPoint{
					Tick: frame,
					X:    posx,
					Y:    posy,
				})
			}
		*/
	}
	return nil
}

func main() {

	parsing := parser{}
	err := parsing.ticks()
	check(err)

	//draw()
	app.Run()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
