package utils

import (
	"errors"
	"fmt"
	"strings"

	"adesnia/galaxy-merchant-trading-guide/pkg/currency"
	"adesnia/galaxy-merchant-trading-guide/pkg/price"
)

// Parse parses input
func Parse(text []string) (out string, err error) {
	if len(text) == 3 {
		if text[1] != "is" {
			return "", errors.New("invalid")
		}
		err = currency.Register(text[0], text[2])
		return "", err
	}

	if text[len(text)-1] == "Credits" {
		err = price.Add(text)
		return "", err
	}

	if text[0] == "how" {
		if text[len(text)-1] != "?" {
			return "I have no idea what you are talking about", nil
		}
		text = text[:len(text)-1]
		if text[1] == "much" {
			if text[2] == "is" {
				var values []string
				for i := 0; i < len(text); i++ {
					if i > 2 {
						values = append(values, text[i])
					}
				}
				var priceConv []string
				for i := 3; i < len(text); i++ {
					priceConv = append(priceConv, text[i])
				}
				exactPrice, errCalculatePrice := price.Calculate(priceConv)
				if errCalculatePrice != nil {
					return "", errCalculatePrice
				}
				return fmt.Sprintf("%v is %v", strings.Join(priceConv, " "), exactPrice), nil
			}
		} else if text[1] == "many" && text[2] == "Credits" && text[3] == "is" {
			itemName := text[len(text)-1]
			amount := text[:len(text)-1][4:len(text[:len(text)-1])]

			if price.Items[itemName] == 0 {
				return "", errors.New("item is not yet registered")
			}
			itemAmount, errCalculateAmount := price.Calculate(amount)
			if errCalculateAmount != nil {
				return "", errCalculateAmount
			}

			itemPrice := itemAmount * price.Items[itemName]

			return fmt.Sprintf("%v %v is %v Credits", strings.Join(amount, " "), itemName, itemPrice), nil
		}
	}

	return "I have no idea what you are talking about", nil
}
