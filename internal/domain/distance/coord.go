package distance

import "errors"

type Coord struct {
	Lat  float64
	Long float64
}

func NewCoord(lat float64, long float64) (*Coord, error) {
	coord := &Coord{
		Lat:  lat,
		Long: long,
	}
	err := coord.Validate()
	if err != nil {
		return nil, err
	}
	return coord, nil
}

func (c *Coord) Validate() error {
	if c.Lat >= 90.0 || c.Lat <= -90.0 {
		return errors.New("invalid latitude")
	}
	if c.Long >= 180.0 || c.Long <= -180.0 {
		return errors.New("invalid longitude")
	}
	return nil
}
