package price

import (
	"errors"
	"fmt"
	"strconv"

	"adesnia/galaxy-merchant-trading-guide/pkg/currency"
)

// Items maps item with its price
var Items map[string]int

// Add adds item and its price
func Add(text []string) error {
	itemName, itemPrice, err := getItemAndPrice(text)
	if err != nil {
		return err
	}
	Items[itemName] = itemPrice
	return nil
}

// Calculate calculates the value of the given intergalactic unit
func Calculate(values []string) (int, error) {
	if valid := validate(values); !valid {
		return 0, errors.New("invalid currency")
	}

	if len(values) > 3 {
		for i := 0; i < len(values)-1; i++ {
			if len(values)-i > 3 {
				if values[i] == values[i+1] && values[i] == values[i+2] && values[i] == values[i+3] {
					return 0, errors.New("invalid pricing")
				}
			}
		}
	}

	dConv := currency.Conversions["D"]
	lConv := currency.Conversions["L"]
	vConv := currency.Conversions["V"]
	dCount := 0
	lCount := 0
	vCount := 0
	for i := 0; i < len(values); i++ {
		switch {
		case values[i] == dConv:
			dCount++
		case values[i] == lConv:
			lCount++
		case values[i] == vConv:
			vCount++
		}
	}
	if dCount > 1 || lCount > 1 || vCount > 1 {
		return 0, errors.New("invalid numerals")
	}

	total := 0
	for i := 0; i < len(values)-1; i++ {
		if currency.Currencies[values[i]] > currency.Currencies[values[i+1]] {
			total = total + currency.Currencies[values[i]]
		} else if currency.Currencies[values[i]] == currency.Currencies[values[i+1]] {
			total = total + (2 * currency.Currencies[values[i]])
			i++
			if i == len(values)-2 {
				total = total + currency.Currencies[values[i+1]]
			}
		} else {
			if currency.Currencies[values[i+1]]/currency.Currencies[values[i]] > 10 {
				return 0, errors.New("invalid pricing")
			}
			total = total + (currency.Currencies[values[i+1]] - currency.Currencies[values[i]])
			i++
		}
	}
	return total, nil
}

func validate(price []string) bool {
	for _, val := range price {
		found := false
		for cur := range currency.Currencies {
			if cur == val {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func getItemAndPrice(text []string) (item string, price int, err error) {
	itemName := text[len(text)-4]
	value := text[len(text)-2]
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		return "", 0, errors.New(fmt.Sprintf("unable to add itemName price: %v", err))
	}

	lastIdxPrice := len(text) - 5
	var priceConv []string
	for i := 0; i <= lastIdxPrice; i++ {
		priceConv = append(priceConv, text[i])
	}

	amount, err := Calculate(priceConv)
	if err != nil {
		return "", 0, err
	}
	itemPrice := valueInt / amount
	return itemName, itemPrice, nil
}
