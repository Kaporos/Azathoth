package components

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/kaporos/azathoth/core"
)

const renderSize = 10

func MapToString(player *core.Player) string {
	player_pos := player.GetPosition()

	result := fmt.Sprintf("%sYour position: (%d, %d)\n", strings.Repeat(" ", renderSize), player_pos.GetX(), player_pos.GetY())
	for j := player_pos.GetY() - renderSize; j < player_pos.GetY()+renderSize; j++ {
		for i := player_pos.GetX() - renderSize/2; i < player_pos.GetX()+renderSize/2; i++ {
			coords := core.NewCoordinate(i, j)
			v := core.GetMapTile(coords)
			s := ""
			if v.Biome == core.Plain {
				// ground
				s = "...."
			}
			if v.Biome == core.Forest {
				// forest
				s = color.GreenString("ηηηη")
			}
			if v.Biome == core.Water {
				s = color.BlueString("~~~~")
			}
			if v.Biome == core.Moutain {
				s = "/^^^"
			}
			if *player_pos == coords {
				// player
				s = color.MagentaString("<oo>")
			}
			result += s

		}
		result += "\n"
	}
	currTile := core.GetMapTile(*player_pos)
	result += fmt.Sprintf("You are %s", toString(&currTile.Biome))
	return result
}

func toString(b *core.Biome) string {
	switch *b {
	case core.Plain:
		return "on a field of grass"
	case core.Water:
		return "in the " + color.BlueString("water")
	case core.Forest:
		return "in a " + color.GreenString("forest")
	case core.Moutain:
		return "on a moutain"

	}
	return "Magic biome"
}
func RenderMap(p *core.Player) string {
	return MapToString(p)
}
