package types

// bsd/invNames.yaml

type Item struct {
	ItemID   int    `yaml:"itemID"`
	ItemName string `yaml:"itemName"`
}

type ItemList []Item
