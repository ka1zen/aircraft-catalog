package uc_test

import (
	"testing"

	"github.com/aircraft-catalog/uc"
	"github.com/stretchr/testify/assert"
)

func Test_CoreLogic_GetCatalogByName(t *testing.T) {
	t.Run("should return no AirCraft", func(t *testing.T) {
		catalogs, err := uc.NewCoreLogic(nil).GetCatalogByName("")
		assert.Error(t, err)
		assert.Nil(t, catalogs)
	})
}
