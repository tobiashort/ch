package canton

type Coordinates struct {
	North float32
	West  float32
}

type Canton struct {
	Name     string
	Abbr     string
	Polygons [][]Coordinates
}
