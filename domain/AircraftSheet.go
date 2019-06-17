package domain

type AircraftSheet struct {
	ID              int
	Capacity        string
	Dimensions      string
	CountryOfOrigin string
	Production      string
	Weights         string
	Performance     string
	Name            string
	Type            string
	Powerplants     string
	History         string
}

func NewAircraftSheet(
	id int,
	Capacity,
	Dimensions,
	CountryOfOrigin,
	Production,
	Weights,
	Performance,
	Name,
	Type,
	Powerplants,
	History string) *AircraftSheet {
	return &AircraftSheet{
		ID:              id,
		Capacity:        Capacity,
		Dimensions:      Dimensions,
		CountryOfOrigin: CountryOfOrigin,
		Production:      Production,
		Weights:         Weights,
		Performance:     Performance,
		Name:            Name,
		Type:            Type,
		Powerplants:     Powerplants,
		History:         History,
	}
}
