package types

// universe/landmarks

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Landmark struct {
	DescriptionID  int      `json:"descriptionID"`
	LandmarkNameID int      `json:"landmarkNameID"`
	IconID         *int     `json:"iconID,omitempty"`
	LocationID     *int     `json:"locationID,omitempty"`
	Position       Position `json:"position"`
}

type LandmarkMap map[int]Landmark
