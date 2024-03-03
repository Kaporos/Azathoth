package components

import (
	"github.com/kaporos/azathoth/core"
)

func RenderPlayerStats(p *core.Player) string {
	return RenderPlayerInventory(p)
}

func RenderPlayerInventory(p *core.Player) string {
	result := "Your inventory: \n\n"
	for _, item := range p.Inventory {
		result += "\t- " + item.DisplayName + "\n"
	}
	return result
}
