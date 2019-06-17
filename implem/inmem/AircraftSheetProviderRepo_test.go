package inmem_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aircraft-catalog/domain"

	"github.com/aircraft-catalog/implem/inmem"
)

func Test_AircraftSheetProviderRepo_FindByName(t *testing.T) {
	asp := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
		0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
	)).Add(domain.NewAircraftSheet(
		1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
	)).Add(domain.NewAircraftSheet(
		2, "", "", "", "", "", "", "ODD_ONE", "", "", "",
	))
	t.Run("should return no AirCraftSheets", func(t *testing.T) {
		got := asp.FindByName("NOTHING")
		assert.Len(t, got, 0)
	})

	t.Run("should return matching name's AirCraftSheets", func(t *testing.T) {
		got := asp.FindByName("aero")
		assert.Len(t, got, 2)

		for _, aircraftSheet := range got {
			assert.NotEqual(t, aircraftSheet.Name, "ODD_ONE")
		}
	})
}

func Test_AircraftSheetProviderRepo_FindByID(t *testing.T) {
	asp := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
		0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
	)).Add(domain.NewAircraftSheet(
		1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
	))

	t.Run("should return AirCraftSheets with id 1", func(t *testing.T) {
		got := asp.FindByID(1)
		assert.Equal(t, got.ID, 1)
	})

	t.Run("should return no aircraftSheet (nil)", func(t *testing.T) {
		ID := 5
		assert.Nil(t, asp.FindByID(ID))
	})
}

func Test_AircraftSheetProviderRepo_FindAll(t *testing.T) {
	asp := inmem.NewAircraftSheetProviderRepo()
	t.Run("should return no aircraftSheet", func(t *testing.T) {
		assert.Len(t, asp.FindAll(), 0)
	})
	t.Run("should return all AirCraftSheets", func(t *testing.T) {
		asp.Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
		)).Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
		))

		got := asp.FindAll()
		assert.Len(t, got, 2)
	})
}

func Test_AircraftSheetProviderRepo_Store(t *testing.T) {
	asp := inmem.NewAircraftSheetProviderRepo()
	t.Run("should return the id of the added aircraftSheet", func(t *testing.T) {
		actualID, _ := asp.Store(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
		))
		assert.Len(t, *asp, 1)
		assert.Equal(t, 0, actualID)
	})

	t.Run("should return the incremented id of the added aircraftSheet", func(t *testing.T) {
		actualID, _ := asp.Store(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
		))
		assert.Len(t, *asp, 2)
		assert.Equal(t, 1, actualID)
	})
}

func Test_AircraftSheetProviderRepo_Update(t *testing.T) {
	asp := inmem.NewAircraftSheetProviderRepo().
		Add(domain.NewAircraftSheet(
			0, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "")).
		Add(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "Aeronca 22", "", "", ""))

	t.Run("should return the id of the updated aircraftSheet", func(t *testing.T) {
		actualID, err := asp.Update(domain.NewAircraftSheet(
			1, "", "", "", "", "", "", "UPDATED", "", "", "",
		))
		assert.NoError(t, err)
		assert.Equal(t, 1, actualID)
		assert.Equal(t, "UPDATED", asp.FindByID(actualID).Name)
	})

	t.Run("should return error if the id does not exist", func(t *testing.T) {
		_, err := asp.Update(domain.NewAircraftSheet(
			5, "", "", "", "", "", "", "WRONG_ID", "", "", "",
		))
		assert.Error(t, err)
	})
}

func TestAircraftSheetProviderRepo_Remove(t *testing.T) {
	asp := inmem.NewAircraftSheetProviderRepo().Add(domain.NewAircraftSheet(
		0, "", "", "", "", "", "", "Aero Boero AB-95/115/150/180", "", "", "",
	)).Add(domain.NewAircraftSheet(
		1, "", "", "", "", "", "", "Aeronca 11 Chief", "", "", "",
	))

	t.Run("should return error if id not existing", func(t *testing.T) {
		err := asp.Remove(5)
		assert.Error(t, err)
		assert.Len(t, *asp, 2)
	})

	t.Run("should return no error", func(t *testing.T) {
		err := asp.Remove(1)
		assert.Nil(t, asp.FindByID(1))
		assert.Len(t, *asp, 1)
		assert.NoError(t, err)
	})
}
