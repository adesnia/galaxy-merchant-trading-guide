package currency

import (
	"errors"
)

// Conversions maps Roman Numerals with intergalactic unit i.e Conversions[I] = glob
var Conversions map[string]string

// Currencies maps intergalactic numerals with value i.e Currencies[glob] = 1
var Currencies map[string]int

// Register registers currency
func Register(name, value string) error {
	exist := false
	for j := range Units {
		if j == value {
			exist = true
		}
	}
	if !exist {
		return errors.New("invalid unit")
	}

	for i := range Conversions {
		if Conversions[i] == name {
			return errors.New("oops name listed")
		}
	}
	Conversions[value] = name

	var currencyVal int
	for i := range Units {
		if i == value {
			currencyVal = Units[i]
		}
	}
	Currencies[name] = currencyVal
	return nil
}
