package distance

import "math"

type CosinDistance struct{}

func NewCosinDistance() *CosinDistance {
	return &CosinDistance{}
}

func (d *CosinDistance) Calculate(from *Coord, to *Coord, unit string) float64 {
	if from == nil || to == nil {
		return 0
	}
	if from.Lat == to.Lat && from.Long == to.Long {
		return 0
	}
	dist := d.getAngularDistance(from, to)
	return convertDistanceUnit(dist, unit)
}

func (d *CosinDistance) getAngularDistance(from *Coord, to *Coord) float64 {
	b := degToRad(90 - from.Lat)
	c := degToRad(90 - to.Lat)
	theta := degToRad(from.Long - to.Long)
	a := math.Cos(b)*math.Cos(c) + math.Sin(b)*math.Sin(c)*math.Cos(theta)
	dist := math.Acos(a)
	return dist
}
