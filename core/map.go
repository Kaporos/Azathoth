package core

import (
	"math/rand"

	"github.com/aquilax/go-perlin"
)

type Biome int

const (
	Plain Biome = iota
	Forest
	Moutain
	Water
)

type WordTile struct {
	Biome Biome
}

var world_map map[Coordinate]WordTile

func GetMapTile(coord Coordinate) *WordTile {
	_, ok := world_map[coord]
	if !ok {
		world_map[coord] = WordTile{
			Biome: Water,
		}
	}
	val, _ := world_map[coord]
	return &val
}

func setTile(coord Coordinate, tile WordTile) {
	world_map[coord] = tile
}

func init() {
	world_map = make(map[Coordinate]WordTile)
	GenerateIsland(Coordinate{0, 0}, 25)

}

// TODO: Add a filter like this https://medium.com/@travall/procedural-2d-island-generation-noise-functions-13976bddeaf9
func GenerateIsland(coord Coordinate, size int) {
	size_f := float64(size)
	seed := rand.Int63()
	p := perlin.NewPerlin(2, 10, 3, seed)
	for x := -size_f / 2.0; x < size_f/2.0; x++ {
		for y := -size_f / 2.0; y < size_f/2.0; y++ {
			new_coord := Coordinate{coord.x + int(x), coord.y + int(y)}
			value := (p.Noise2D(x/8, y/8))
			tile := WordTile{}
			if value <= -0.3 {
				tile.Biome = Water
			} else if value <= 0.10 {
				tile.Biome = Plain
			} else if value <= 0.4 {
				tile.Biome = Forest
			} else {
				tile.Biome = Moutain
			}
			setTile(new_coord, tile)

		}
	}

}
