package utils

import (
	"testing"

	"adesnia/galaxy-merchant-trading-guide/pkg/currency"
	"adesnia/galaxy-merchant-trading-guide/pkg/price"

	"github.com/stretchr/testify/assert"
)

func TestParseRegisterCurrencySuccess(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)

	input := []string{"glob", "is", "I"}
	_, err := Parse(input)

	assert.Nil(t, err)
	assert.Equal(t, "glob", currency.Conversions["I"])
	assert.Equal(t, 1, currency.Currencies["glob"])
}

func TestParseRegisterCurrencyFailedInvalidCurrency(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)

	input := []string{"glob", "is", "Z"}
	_, err := Parse(input)

	assert.NotNil(t, err)
}

func TestParseRegisterCurrencyFailedInvalidInput(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)

	input := []string{"glob", "are", "1"}
	_, err := Parse(input)

	assert.NotNil(t, err)
}

func TestParseAddItemPriceSuccess(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	registerCurrency := []string{"glob", "is", "I"}
	_, errRegisterCurrency := Parse(registerCurrency)
	price.Items = make(map[string]int)

	addItemPriceInput := []string{"glob", "glob", "Silver", "is", "60", "Credits"}
	_, errParse := Parse(addItemPriceInput)

	assert.Nil(t, errRegisterCurrency)
	assert.Nil(t, errParse)
	assert.Equal(t, 30, price.Items["Silver"])
}

func TestParseAddItemPriceFailedInvalidAmount(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	registerCurrency := []string{"glob", "is", "I"}
	_, errRegisterCurrency := Parse(registerCurrency)
	price.Items = make(map[string]int)

	addItemPriceInput := []string{"glob", "glob", "Silver", "is", "s", "Credits"}
	_, errParse := Parse(addItemPriceInput)

	assert.Nil(t, errRegisterCurrency)
	assert.NotNil(t, errParse)
}

func TestParseAddItemPriceFailedUnknownCurrency(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	registerCurrency := []string{"glob", "is", "I"}
	_, errRegisterCurrency := Parse(registerCurrency)
	price.Items = make(map[string]int)

	addItemPriceInput := []string{"glob", "glob", "prok", "Silver", "is", "34", "Credits"}
	_, errParse := Parse(addItemPriceInput)

	assert.Nil(t, errRegisterCurrency)
	assert.NotNil(t, errParse)
	assert.Equal(t, 0, price.Items["Silver"])
}

func TestParseAddItemPriceFailedInvalidCurrency(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	registerCurrency := []string{"glob", "is", "I"}
	_, errRegisterCurrency := Parse(registerCurrency)
	price.Items = make(map[string]int)

	addItemPriceInput := []string{"glob", "glob", "glob", "glob", "Silver", "is", "34", "Credits"}
	_, errParse := Parse(addItemPriceInput)

	assert.Nil(t, errRegisterCurrency)
	assert.NotNil(t, errParse)
	assert.Equal(t, 0, price.Items["Silver"])
}

func TestParseCalculateCurrencySuccess(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.Conversions["I"] = "glob"
	currency.Conversions["X"] = "pish"
	currency.Conversions["L"] = "tegj"
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Currencies["glob"] = 1
	currency.Currencies["pish"] = 10
	currency.Currencies["tegj"] = 50

	input := []string{"how", "much", "is", "pish", "tegj", "glob", "glob", "?"}
	out, errParse := Parse(input)

	assert.Nil(t, errParse)
	assert.Equal(t, "pish tegj glob glob is 42", out)
}

func TestParseCalculateCurrencyFailedUnknownCurrency(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.Conversions["I"] = "glob"
	currency.Conversions["X"] = "pish"
	currency.Conversions["L"] = "tegj"
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Currencies["glob"] = 1
	currency.Currencies["pish"] = 10
	currency.Currencies["tegj"] = 50

	input := []string{"how", "much", "is", "pish", "tegj", "glob", "yelb", "?"}
	out, errParse := Parse(input)

	assert.NotNil(t, errParse)
	assert.Equal(t, "", out)
}

func TestParseCalculateCurrencyFailedInvalidCurrencyFormat(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.Conversions["I"] = "glob"
	currency.Conversions["C"] = "blau"
	currency.Conversions["M"] = "rod"
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Currencies["glob"] = 1
	currency.Currencies["blau"] = 100
	currency.Currencies["rod"] = 1000

	input := []string{"how", "much", "is", "rod", "blau", "rod", "glob", "rod", "?"}
	out, errParse := Parse(input)

	assert.NotNil(t, errParse)
	assert.Equal(t, "", out)
}

func TestParseCalculatePriceSuccess(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.Conversions["I"] = "glob"
	currency.Conversions["V"] = "prok"
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Currencies["glob"] = 1
	currency.Currencies["prok"] = 5
	price.Items = make(map[string]int)
	price.Items["Silver"] = 19

	input := []string{"how", "many", "Credits", "is", "glob", "prok", "Silver", "?"}
	out, errParse := Parse(input)

	assert.Nil(t, errParse)
	assert.Equal(t, "glob prok Silver is 76 Credits", out)
}

func TestParseCalculatePriceFailedUnknownCurrency(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.Conversions["I"] = "glob"
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Currencies["glob"] = 1
	price.Items = make(map[string]int)
	price.Items["Silver"] = 19

	input := []string{"how", "many", "Credits", "is", "glob", "prok", "Silver", "?"}
	out, errParse := Parse(input)

	assert.NotNil(t, errParse)
	assert.Equal(t, "", out)
}

func TestParseCalculatePriceFailedInvalidItem(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.Conversions["I"] = "glob"
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	currency.Currencies["glob"] = 1
	price.Items = make(map[string]int)
	price.Items["Coal"] = 19

	input := []string{"how", "many", "Credits", "is", "glob", "prok", "Silver", "?"}
	out, errParse := Parse(input)

	assert.NotNil(t, errParse)
	assert.Equal(t, "", out)
}

func TestParseUnknownFormat(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	price.Items = make(map[string]int)

	input := []string{"how", "much", "wood", "could", "a", "woodchuck", "chuck", "if", "a", "woodchuck", "could", "chuck", "wood", "?"}
	out, errParse := Parse(input)

	assert.Nil(t, errParse)
	assert.Equal(t, "I have no idea what you are talking about", out)
}

func TestParseInvalidQuestion(t *testing.T) {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	currency.Currencies = make(map[string]int)
	price.Items = make(map[string]int)

	input1 := []string{"how", "much", "is", "pish", "tegj", "glob", "glob"}
	input2 := []string{"how", "many", "Credits", "is", "glob", "prok", "Silver"}
	out1, errParse1 := Parse(input1)
	out2, errParse2 := Parse(input2)

	assert.Nil(t, errParse1)
	assert.Nil(t, errParse2)
	assert.Equal(t, "I have no idea what you are talking about", out1)
	assert.Equal(t, "I have no idea what you are talking about", out2)
}
