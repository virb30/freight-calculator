package distance

import "math"

const (
	earthRadiusMi = 3958
	earthRadiusKm = 6371
)

func degToRad(deg float64) float64 {
	return (deg * math.Pi) / 180.0
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
