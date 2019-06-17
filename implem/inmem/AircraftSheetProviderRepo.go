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
	*asp = append(*asp, aircraftSheet)
	return aircraftSheet.ID, nil
}

func (asp *AircraftSheetProviderRepo) Update(as *domain.AircraftSheet) (int, error) {
	for i, aircraftSheet := range *asp {
		if aircraftSheet.ID == as.ID {
			(*asp)[i] = as
			return as.ID, nil
		}
	}
	return 0, fmt.Errorf("id %d not existing", as.ID)
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
