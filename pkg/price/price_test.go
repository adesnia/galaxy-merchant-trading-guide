package price

import (
	"adesnia/galaxy-merchant-trading-guide/pkg/currency"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddPriceSuccess(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Conversions["I"] = "glob"
	currency.Currencies["glob"] = 1
	Items = make(map[string]int)
	input := []string{"glob", "glob", "Silver", "is", "34", "Credits"}

	err := Add(input)

	assert.Nil(t, err)
	assert.Equal(t, 17, Items["Silver"])
}

func TestAddPriceFailedInvalidAmount(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Conversions["I"] = "glob"
	currency.Currencies["glob"] = 1
	Items = make(map[string]int)
	input := []string{"glob", "glob", "Silver", "is", "xx", "Credits"}

	err := Add(input)

	assert.NotNil(t, err)
	assert.Equal(t, 0, Items["Silver"])
}

func TestAddPriceFailedUnknownConversion(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Conversions["I"] = "glob"
	currency.Currencies["glob"] = 1
	Items = make(map[string]int)
	input := []string{"glob", "prok", "Silver", "is", "34", "Credits"}

	err := Add(input)

	assert.NotNil(t, err)
	assert.Equal(t, 0, Items["Silver"])
}

func TestAddPriceFailedInvalidPricing(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Conversions["I"] = "glob"
	currency.Currencies["glob"] = 1
	Items = make(map[string]int)
	input := []string{"glob", "glob", "glob", "glob", "Silver", "is", "34", "Credits"}

	err := Add(input)

	assert.NotNil(t, err)
	assert.Equal(t, 0, Items["Silver"])
}

func TestAddPriceFailedForbiddenRomanNumberRepeated(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Conversions["I"] = "glob"
	currency.Conversions["V"] = "blau"
	currency.Conversions["L"] = "gelb"
	currency.Conversions["D"] = "quack"
	currency.Currencies["glob"] = 1
	currency.Currencies["blau"] = 5
	currency.Currencies["gelb"] = 50
	currency.Currencies["quack"] = 500
	Items = make(map[string]int)
	input1 := []string{"blau", "gelb", "blau", "Silver", "is", "34", "Credits"}
	input2 := []string{"gelb", "glob", "gelb", "Silver", "is", "34", "Credits"}
	input3 := []string{"quack", "glob", "quack", "Silver", "is", "34", "Credits"}

	err1 := Add(input1)
	err2 := Add(input2)
	err3 := Add(input3)

	assert.NotNil(t, err1)
	assert.NotNil(t, err2)
	assert.NotNil(t, err3)
	assert.Equal(t, 0, Items["Silver"])
}

func TestAddComplexPriceSuccess(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Conversions["I"] = "glob"
	currency.Conversions["C"] = "blau"
	currency.Conversions["M"] = "rod"
	currency.Currencies["glob"] = 1
	currency.Currencies["blau"] = 100
	currency.Currencies["rod"] = 1000
	Items = make(map[string]int)
	input := []string{"rod", "blau", "rod", "glob", "glob", "glob", "Silver", "is", "1903", "Credits"}

	err := Add(input)

	assert.Nil(t, err)
	assert.Equal(t, 1, Items["Silver"])
}

func TestAddPriceFailedInvalidCurrencyFormat(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Conversions["I"] = "glob"
	currency.Conversions["C"] = "blau"
	currency.Conversions["M"] = "rod"
	currency.Currencies["glob"] = 1
	currency.Currencies["blau"] = 100
	currency.Currencies["rod"] = 1000
	Items = make(map[string]int)
	input := []string{"rod", "blau", "rod", "glob", "rod", "Silver", "is", "1903", "Credits"}

	err := Add(input)

	assert.NotNil(t, err)
	assert.Equal(t, 0, Items["Silver"])
}
