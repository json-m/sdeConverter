package types

type Coordinate struct {
	ItemID int     `yaml:"itemID"`
	X      float64 `yaml:"x"`
	Y      float64 `yaml:"y"`
	Z      float64 `yaml:"z"`
}

type CoordinateList []Coordinate
