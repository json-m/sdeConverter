package types

// fsd/types.yaml

type TypeInfo struct {
	BasePrice             float64           `json:"basePrice,omitempty"`
	Capacity              float64           `json:"capacity,omitempty"`
	Description           map[string]string `json:"description,omitempty"`
	GraphicID             int               `json:"graphicID,omitempty"`
	GroupID               int               `json:"groupID,omitempty"`
	IconID                int               `json:"iconID,omitempty"`
	MarketGroupID         int               `json:"marketGroupID,omitempty"`
	Mass                  float64           `json:"mass,omitempty"`
	MetaGroupID           int               `json:"metaGroupID,omitempty"`
	Name                  map[string]string `json:"name,omitempty"`
	PortionSize           int               `json:"portionSize,omitempty"`
	Published             bool              `json:"published,omitempty"`
	RaceID                int               `json:"raceID,omitempty"`
	Radius                float64           `json:"radius,omitempty"`
	SofFactionName        string            `json:"sofFactionName,omitempty"`
	SoundID               int               `json:"soundID,omitempty"`
	VariationParentTypeID int               `json:"variationParentTypeID,omitempty"`
	Volume                float64           `json:"volume,omitempty"`
}
