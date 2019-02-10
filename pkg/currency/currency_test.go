package currency

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRegisterNewCurrency(t *testing.T){
	Conversions = make(map[string]string)
	RegisterUnits()
	Currencies= make(map[string]int)

	err1 := Register("glob", "I")
	err2 := Register("prok", "V")

	assert.Nil(t, err1)
	assert.Nil(t, err2)
}

func TestRegisterNewCurrencyFailedInvalidUnit(t *testing.T){
	Conversions = make(map[string]string)
	RegisterUnits()
	Currencies= make(map[string]int)

	err := Register("glob", "Z")
	assert.NotNil(t, err)
}

func TestRegisterNewCurrencyFailedValueRegistered(t *testing.T){
	Conversions = make(map[string]string)
	RegisterUnits()
	Currencies= make(map[string]int)

	err1 := Register("glob", "I")
	err2 := Register("glob", "V")

	assert.Nil(t, err1)
	assert.NotNil(t, err2)
}

func TestRegisterCurrencySuccessUpdate(t *testing.T){
	Conversions = make(map[string]string)
	RegisterUnits()
	Currencies= make(map[string]int)

	err1 := Register("glob", "I")
	err2 := Register("prok", "I")

	assert.Nil(t, err1)
	assert.Nil(t, err2)
	assert.Equal(t, "prok", Conversions["I"])
}