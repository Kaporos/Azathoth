package stores

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Receipe struct {
	//TODO: Replace theses by []Item (not simple, idk how to do this properly)
	Input  []ItemId `toml:"input"`
	Output []ItemId `toml:"output"`
}

type Transformer struct {
	//TODO: Force player to have tool in his hand/inventory to use transformer
	Tool string `toml:"tool"`

	Receipes map[string]Receipe `toml:"receipe"`
}

var transformers map[string]Transformer

func GetTransformers() map[string]Transformer {
	return transformers
}
func LoadTransforms() {
	transformers = make(map[string]Transformer)
	fmt.Println("Loading transforms from filesystem..")
	tomlData, err := os.ReadFile("./data/transforms/crafts.toml")
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(tomlData, &transformers); err != nil {
		panic(err)
	}

}
