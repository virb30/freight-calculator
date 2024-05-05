package distance

import "math"

const (
	earthRadiusMi = 3958
	earthRadiusKm = 6371
)

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

func (d *Distance) Calculate(unit string) float64 {
	if d.From.Lat == d.To.Lat && d.From.Long == d.To.Long {
		return 0
	}
	dist := d.haversineFormula()
	return convertDistanceUnit(dist, unit)
}

// TODO: transformar fórmulas de cálculo em estratégias e injetar no construtor
func (d *Distance) haversineFormula() float64 {
	radLatFrom := degToRad(d.From.Lat)
	radLatTo := degToRad(d.To.Lat)
	radLongFrom := degToRad(d.From.Long)
	radLongTo := degToRad(d.To.Long)
	diffLat := radLatFrom - radLatTo
	diffLong := radLongFrom - radLongTo
	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(radLatFrom)*math.Cos(radLatTo)*
		math.Pow(math.Sin(diffLong/2), 2)
	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func degToRad(deg float64) float64 {
	return (math.Pi * deg) / 180.0
}

func radToDeg(rad float64) float64 {
	return (rad * 180.0) / math.Pi
}

func convertDistanceUnit(dist float64, unit string) float64 {
	switch unit {
	case "km":
		return dist * earthRadiusKm
	default:
		return dist * earthRadiusMi
	}
}
