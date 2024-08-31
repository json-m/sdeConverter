package types

// universe/constellation

type Constellation struct {
	Center          []float64 `yaml:"center"`
	ConstellationID int       `yaml:"constellationID"`
	Max             []float64 `yaml:"max"`
	Min             []float64 `yaml:"min"`
	NameID          int       `yaml:"nameID"`
	Radius          float64   `yaml:"radius"`
}
