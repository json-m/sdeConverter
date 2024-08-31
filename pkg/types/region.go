package types

// universe/region

type Region struct {
	Center        []float64 `yaml:"center"`
	DescriptionID int       `yaml:"descriptionID"`
	FactionID     int       `yaml:"factionID"`
	Max           []float64 `yaml:"max"`
	Min           []float64 `yaml:"min"`
	NameID        int       `yaml:"nameID"`
	Nebula        int       `yaml:"nebula"`
	RegionID      int       `yaml:"regionID"`
}
