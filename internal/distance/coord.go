package distance

type Coord struct {
	Lat  float64
	Long float64
}

func NewCoord(lat float64, long float64) *Coord {
	return &Coord{
		Lat:  lat,
		Long: long,
	}
}
