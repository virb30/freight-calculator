package distance

type DistanceInterface interface {
	Calculate(from *Coord, to *Coord, unit string) float64
}
