package uc

import "github.com/aircraft-catalog/domain"

type coreLogic struct {
	AircraftSheetProvider AircraftSheetProvider
}

type AircraftSheetProvider interface {
	FindByID(ID int) *domain.AircraftSheet
	FindByName(name string) []*domain.AircraftSheet
	FindAll() []*domain.AircraftSheet
	Store(domain.AircraftSheet) (int, error)
	Update(AircraftSheets *domain.AircraftSheet) (int, error)
}

func NewCoreLogic(AircraftSheetProvider AircraftSheetProvider) *coreLogic {
	return &coreLogic{
		AircraftSheetProvider,
	}
}
