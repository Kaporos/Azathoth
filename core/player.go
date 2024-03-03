package core

import (
	"github.com/kaporos/azathoth/stores"
)

type Inventory []stores.Item

type Player struct {
	position    Coordinate
	cam_center  Coordinate
	health      int // 0 - 100
	mana        int // 0 - 100
	stamina     int // 0 - 100
	Inventory   Inventory
	camBox      bool
	lastCommand string
	coins       int
	channels    []string
}

func (i *Inventory) contains(s string) bool {
	for _, item := range *i {
		if item.Id == s {
			return true
		}
	}
	return false
}

func (p *Player) InitPlayer() {
	p.health = 100
	p.stamina = 100
	p.stamina = 100
	p.GiveItem(stores.CreateItemSure("apple"))
	p.GiveItem(stores.CreateItemSure("skin"))

}
func (p *Player) GiveItem(s stores.Item) {
	p.Inventory = append(p.Inventory, s)
}
func (p *Player) Health() int {
	return p.health
}
func (p *Player) Mana() int {
	return p.mana
}
func (p *Player) Stampna() int {
	return p.stamina
}

func remove(inv Inventory, s string) Inventory {
	for i, item := range inv {
		if item.Id == s {
			inv = append((inv)[:i], (inv)[i+1:]...)
			return inv
		}
	}
	return inv
}

func (p *Player) Move(deltaX, deltaY int) {
	p.position.x += deltaX
	p.position.y += deltaY
}

func (p *Player) Heal() {
	if p.mana <= 20 {
		return
	}
	if p.health >= 90 {
		return
	}
	p.health += 10
	p.mana -= 20
}

func (p *Player) Hurt(damages int) {
	p.health -= damages
}

func (p *Player) Clock() {
	//this will run at each clock step. It's like the "regen"
	p.mana = min(100, p.mana+1)
	p.health = min(100, p.health+1)
	p.stamina = min(100, p.stamina+1)
}

func (p *Player) GetPosition() *Coordinate {
	return &p.position
}
