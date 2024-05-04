package distance

type Distance struct {
	From *Coord
	To   *Coord
}

func NewDistance(from *Coord, to *Coord) *Distance {
	return &Distance{
		From: from,
		To:   to,
	}
}

func (d *Distance) Calculate() float64 {
	return 0.0
}
