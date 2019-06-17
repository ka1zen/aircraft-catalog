package uc_test

import (
	"testing"

	"github.com/aircraft-catalog/domain"
	"github.com/aircraft-catalog/implem/inmem"

	"github.com/aircraft-catalog/uc"
	"github.com/stretchr/testify/assert"
)

func Test_CoreLogic_GetAircraftSheetsByName(t *testing.T) {
	repo := inmem.NewAircraftSheetProviderRepo().
		Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "")).
		Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "")).
		Add(domain.NewAircraftSheet(
			2, "", "", "", "", "", "", "ODD_ONE", "", "", ""))
	cl := uc.NewCoreLogic(repo)
	t.Run("should return no AirCraft", func(t *testing.T) {
		aircraftSheets := cl.GetAircraftSheetsByName("")
		assert.Len(t, aircraftSheets, 3)
	})

	t.Run("should return matching name's AirCraftSheets", func(t *testing.T) {
		got := cl.GetAircraftSheetsByName("aero")
		assert.Len(t, got, 2)

		for _, aircraftSheet := range got {
			assert.NotEqual(t, aircraftSheet.Name, "ODD_ONE")
		}
	})
}

func Test_CoreLogic_GetAircraftSheets(t *testing.T) {
	repo := inmem.NewAircraftSheetProviderRepo()
	t.Run("should return no aircraftSheet", func(t *testing.T) {
		actual := uc.NewCoreLogic(repo).GetAircraftSheets()
		assert.Len(t, actual, 0)
	})
	t.Run("should return all AirCraftSheets", func(t *testing.T) {
		repo.
			Add(domain.NewAircraftSheet(
				0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "")).
			Add(domain.NewAircraftSheet(
				1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", ""))

		actual := uc.NewCoreLogic(repo).GetAircraftSheets()
		assert.Len(t, actual, 2)
	})
}
