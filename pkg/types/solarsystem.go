package types

// universe/solarsystem

type SolarSystem struct {
	Border            bool                   `yaml:"border"`
	Center            []float64              `yaml:"center"`
	Corridor          bool                   `yaml:"corridor"`
	Fringe            bool                   `yaml:"fringe"`
	Hub               bool                   `yaml:"hub"`
	International     bool                   `yaml:"international"`
	Luminosity        float64                `yaml:"luminosity"`
	Max               []float64              `yaml:"max"`
	Min               []float64              `yaml:"min"`
	Planets           map[int]Planet         `yaml:"planets"`
	Radius            float64                `yaml:"radius"`
	Regional          bool                   `yaml:"regional"`
	SecondarySun      SecondarySun           `yaml:"secondarySun"`
	Security          float64                `yaml:"security"`
	SolarSystemID     int                    `yaml:"solarSystemID"`
	SolarSystemNameID int                    `yaml:"solarSystemNameID"`
	Star              Star                   `yaml:"star"`
	Stargates         map[string]interface{} `yaml:"stargates"`
	SunTypeID         int                    `yaml:"sunTypeID"`
}

type Planet struct {
	CelestialIndex   int              `yaml:"celestialIndex"`
	PlanetAttributes PlanetAttributes `yaml:"planetAttributes"`
	Position         []float64        `yaml:"position"`
	Radius           int              `yaml:"radius"`
	Statistics       PlanetStatistics `yaml:"statistics"`
	TypeID           int              `yaml:"typeID"`
	Moons            map[int]Moon     `yaml:"moons,omitempty"`
}

type PlanetAttributes struct {
	HeightMap1   int  `yaml:"heightMap1"`
	HeightMap2   int  `yaml:"heightMap2"`
	Population   bool `yaml:"population"`
	ShaderPreset int  `yaml:"shaderPreset"`
}

type PlanetStatistics struct {
	Density        float64 `yaml:"density"`
	Eccentricity   float64 `yaml:"eccentricity"`
	EscapeVelocity float64 `yaml:"escapeVelocity"`
	Fragmented     bool    `yaml:"fragmented"`
	Life           float64 `yaml:"life"`
	Locked         bool    `yaml:"locked"`
	MassDust       float64 `yaml:"massDust"`
	MassGas        float64 `yaml:"massGas"`
	OrbitPeriod    float64 `yaml:"orbitPeriod"`
	OrbitRadius    float64 `yaml:"orbitRadius"`
	Pressure       float64 `yaml:"pressure"`
	Radius         float64 `yaml:"radius"`
	RotationRate   float64 `yaml:"rotationRate"`
	SpectralClass  string  `yaml:"spectralClass"`
	SurfaceGravity float64 `yaml:"surfaceGravity"`
	Temperature    float64 `yaml:"temperature"`
}

type Moon struct {
	PlanetAttributes PlanetAttributes `yaml:"planetAttributes"`
	Position         []float64        `yaml:"position"`
	Radius           int              `yaml:"radius"`
	Statistics       PlanetStatistics `yaml:"statistics"`
	TypeID           int              `yaml:"typeID"`
}

type SecondarySun struct {
	EffectBeaconTypeID int       `yaml:"effectBeaconTypeID"`
	ItemID             int       `yaml:"itemID"`
	Position           []float64 `yaml:"position"`
	TypeID             int       `yaml:"typeID"`
}

type Star struct {
	ID         int            `yaml:"id"`
	Radius     int            `yaml:"radius"`
	Statistics StarStatistics `yaml:"statistics"`
	TypeID     int            `yaml:"typeID"`
}

type StarStatistics struct {
	Age           float64 `yaml:"age"`
	Life          float64 `yaml:"life"`
	Locked        bool    `yaml:"locked"`
	Luminosity    float64 `yaml:"luminosity"`
	Radius        float64 `yaml:"radius"`
	SpectralClass string  `yaml:"spectralClass"`
	Temperature   float64 `yaml:"temperature"`
}
