package distance

import "math"

type HaversineDistance struct {
}

func NewHaversineDistance() *HaversineDistance {
	return &HaversineDistance{}
}

func (d *HaversineDistance) Calculate(from *Coord, to *Coord, unit string) float64 {
	if from == nil || to == nil {
		return 0
	}
	if from.Lat == to.Lat && from.Long == to.Long {
		return 0
	}
	dist := d.getAngularDistance(from, to)
	return convertDistanceUnit(dist, unit)
}

func (d *HaversineDistance) getAngularDistance(from *Coord, to *Coord) float64 {
	radLatFrom := degToRad(from.Lat)
	radLatTo := degToRad(to.Lat)
	radLongFrom := degToRad(from.Long)
	radLongTo := degToRad(to.Long)
	diffLat := radLatFrom - radLatTo
	diffLong := radLongFrom - radLongTo
	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(radLatFrom)*math.Cos(radLatTo)*
		math.Pow(math.Sin(diffLong/2), 2)
	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}
