package uc

import (
	"github.com/aircraft-catalog/domain"
)

// GET aircraft?name=a380
// GET /autocomplete?name=asdfasdfsf
func (cl *coreLogic) GetCatalogByName(name string) ([]*domain.AircraftSheet, error) {
	return cl.AircraftSheetProvider.FindByName(name)
}

// GET /aircrafts
//func (cl *coreLogic) GetCatalogs() ([]*AirCraft, error) {
//
//}

// PUT /aircraft
//func (cl *coreLogic) UpdateCatalog(*AirCraft) (AirCrafts.ID, error) {
//
//}
//
//// DELETE /aircraft/:id
//func RemoveCatalog(*AirCraft) (AirCrafts.ID, error) {
//
//}
