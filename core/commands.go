package core

import (
	"strings"

	"github.com/kaporos/azathoth/stores"
)

func parseCommand(cmd string) (commandName string, args []string) {
	cmd_l := strings.Split(cmd, " ")
	return cmd_l[0], cmd_l[1:]
}

func ProcessCommand(p *Player, cmd string) string {
	commandName, args := parseCommand(cmd)
	if commandName == "give" {
		if len(args) == 0 {
			return "Usage: give <item>"
		}
		item, err := stores.CreateItem(args[0])
		if err != nil {
			return "Item not found"
		}
		p.GiveItem(item)
		return "God just gave you " + item.DisplayName
	}
	return "Huh ?!"
}
