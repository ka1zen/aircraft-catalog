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

func TestCoreLogic_UpdateAircraftSheet(t *testing.T) {
	repo := inmem.NewAircraftSheetProviderRepo().
		Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "")).
		Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 22", "", "", ""))

	t.Run("should return the id of the updated aircraftSheet", func(t *testing.T) {
		actualID, err := uc.NewCoreLogic(repo).UpdateAircraftSheet(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "UPDATED", "", "", "",
		))
		assert.NoError(t, err)
		assert.Equal(t, 1, actualID)
		assert.Equal(t, "UPDATED", repo.FindByID(actualID).Name)
	})

	t.Run("should return error if the id does not exist", func(t *testing.T) {
		_, err := uc.NewCoreLogic(repo).UpdateAircraftSheet(domain.NewAircraftSheet(
			5, "", "", "", "", "", "", "WRONG_ID", "", "", "",
		))
		assert.Error(t, err)
	})
}

func TestCoreLogic_GetAircraftDetails(t *testing.T) {
	repo := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
		0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
	)).Add(domain.NewAircraftSheet(
		1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
	))

	t.Run("should return AirCraftSheets with id 1", func(t *testing.T) {
		got := uc.NewCoreLogic(repo).GetAircraftDetails(1)
		assert.Equal(t, got.ID, 1)
	})

	t.Run("should return no aircraftSheet (nil)", func(t *testing.T) {
		got := uc.NewCoreLogic(repo).GetAircraftDetails(5)
		assert.Nil(t, got)
	})
}
