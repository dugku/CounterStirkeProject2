package Draw

import (
	"fmt"
	"github.com/dugku/CounterStirkeProject2/setup"
)

func drawplayers() {
	for _, i := range setup.PlayersMoveDat {
		fmt.Println(i)
	}
}
