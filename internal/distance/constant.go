package distance

type ConstantDistance struct {
}

func NewConstantDistance() *ConstantDistance {
	return &ConstantDistance{}
}

func (d *ConstantDistance) Calculate(from *Coord, to *Coord, unit string) float64 {
	return 1000
}
