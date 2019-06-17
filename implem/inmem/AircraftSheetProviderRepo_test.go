package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aircraft-catalog/domain"

	"github.com/aircraft-catalog/implem/inmem"
)

func Test_AircraftSheetProviderRepo_FindByName(t *testing.T) {
	t.Run("should return no AirCraftSheets", func(t *testing.T) {
		asp := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
		)).Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
		))

		got := asp.FindByName("NOTHING")
		assert.Len(t, got, 0)
	})

	t.Run("should return matching name's AirCraftSheets", func(t *testing.T) {
		asp := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
		)).Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
		)).Add(domain.NewAircraftSheet(
			2, "", "", "", "", "", "", "ODD_ONE", "", "", "",
		))

		got := asp.FindByName("aero")
		assert.Len(t, got, 2)

		for _, aircraftSheet := range got {
			assert.NotEqual(t, aircraftSheet.Name, "ODD_ONE")
		}
	})
}

func Test_AircraftSheetProviderRepo_FindByID(t *testing.T) {
	t.Run("should return AirCraftSheets with id 1", func(t *testing.T) {
		asp := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
		)).Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
		))

		got := asp.FindByID(1)
		assert.Equal(t, got.ID, 1)
	})

	t.Run("should return no aircraftSheet (nil)", func(t *testing.T) {
		asp := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
		)).Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
		))
		ID := 5
		assert.Nil(t, asp.FindByID(ID))
	})
}
