package stores

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type ItemId string
type Item struct {
	Id          string
	DisplayName string `toml:"display_name"`
	Weight      int
	Nutrition   int
}

type Filedata struct {
	Items map[ItemId]Item
}

var items map[ItemId]Item

func CreateItem(s string) (Item, error) {
	val, ok := items[ItemId(s)]
	if !ok {
		return Item{}, errors.New("This item does not exist.")
	}
	return val, nil
}

func CreateItemSure(s string) Item {
	return items[ItemId(s)]
}

func LoadItems() {
	items = make(map[ItemId]Item)
	fmt.Println("Loading items types from fileystem..")
	var fileContent Filedata
	tomlData, _ := os.ReadFile("./data/items/basics.toml")
	dat := string(tomlData)
	_, err := toml.Decode(dat, &fileContent)
	if err != nil {
		panic(err)
	}
	for k, v := range fileContent.Items {
		if _, ok := items[k]; ok {
			//there already is an item with this ID
			panic(fmt.Sprintf("Item ", k, "defined multiple times\n"))
		}
		v.Id = string(k)
		items[k] = v
	}
}
