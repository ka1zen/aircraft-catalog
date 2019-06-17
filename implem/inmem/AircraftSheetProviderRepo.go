package inmem

import (
	"fmt"
	"strings"

	"github.com/aircraft-catalog/domain"
)

type AircraftSheetProviderRepo []*domain.AircraftSheet

func (asp *AircraftSheetProviderRepo) FindByID(ID int) *domain.AircraftSheet {
	for _, aircraftSheet := range *asp {
		if aircraftSheet.ID == ID {
			return aircraftSheet
		}
	}
	return nil
}

func (asp *AircraftSheetProviderRepo) FindAll() []*domain.AircraftSheet {
	return *asp
}

func (asp *AircraftSheetProviderRepo) Store(aircraftSheet *domain.AircraftSheet) (int, error) {
	aircraftSheet.ID = len(*asp)
	fmt.Println(aircraftSheet.ID)
	*asp = append(*asp, aircraftSheet)
	return aircraftSheet.ID, nil
}

func (asp *AircraftSheetProviderRepo) Update(aircraftSheets *domain.AircraftSheet) (int, error) {
	panic("implement me")
}

func (asp *AircraftSheetProviderRepo) FindByName(name string) []*domain.AircraftSheet {
	var results []*domain.AircraftSheet
	for _, aircraftSheet := range *asp {
		if strings.Contains(
			strings.ToLower(aircraftSheet.Name),
			strings.ToLower(name),
		) {
			results = append(results, aircraftSheet)
		}
	}
	return results
}

func (asp *AircraftSheetProviderRepo) Add(sheet *domain.AircraftSheet) *AircraftSheetProviderRepo {
	*asp = append(*asp, sheet)
	return asp
}

func NewAircraftSheetProviderRepo() *AircraftSheetProviderRepo {
	return &AircraftSheetProviderRepo{}
}
