package components

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kaporos/azathoth/core"
)

func RenderPlayerStats(p *core.Player) string {
	var healthStr, manaStr, stampnaStr string

	health := p.Health()
	mana := p.Mana()
	stampna := p.Stampna()

	if health <= 20 {
		healthStr = color.RedString("%d", health)
	} else if health >= 80 {
		healthStr = color.GreenString("%d", health)
	} else {
		healthStr = color.YellowString("%d", health)
	}

	if mana <= 20 {
		manaStr = color.RedString("%d", mana)
	} else if mana >= 80 {
		manaStr = color.GreenString("%d", mana)
	} else {
		manaStr = color.YellowString("%d", mana)
	}

	if stampna <= 20 {
		stampnaStr = color.RedString("%d", stampna)
	} else if stampna >= 80 {
		stampnaStr = color.GreenString("%d", stampna)
	} else {
		stampnaStr = color.YellowString("%d", stampna)
	}

	stats := fmt.Sprintf("Health: %s | Mana: %s | Stamina: %s\n\n", healthStr, manaStr, stampnaStr)
	stats += RenderPlayerInventory(p)
	return stats
}

func RenderPlayerInventory(p *core.Player) string {
	result := "Your inventory: \n\n"
	for _, item := range p.Inventory {
		result += "\t- " + item.DisplayName + "\n"
	}
	return result
}
