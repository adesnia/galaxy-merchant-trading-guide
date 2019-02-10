package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"github.com/sirupsen/logrus"
	"errors"
) //todo: adjust import

var currency map[string]string //value name //glob I
var units map[string]int //romanAlphabet intValue // I 1

type configuration struct {
	Port                             string `env:"PORT" envDefault:"8080"`
	BasePath                         string `env:"BASE_PATH" envDefault:"/event"`
}

func main() {
	currency = make(map[string]string)
	units = make(map[string]int)
	units["I"] = 1
	units["V"] = 5
	units["X"] = 10
	units["L"] = 50
	units["C"] = 100
	units["D"] = 500
	units["M"] = 1000

	for
	{
		{
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter input: ")
			input, _ := reader.ReadString('\n')
			text := strings.Split(input, " ")
			if len(text) <= 2 {
				logrus.Errorf("invalid input")
			} else if err := parse(text); err != nil {
				logrus.Errorf("unable to parse: %v", err)
			}
		}
	}

	return
}

func parse(text []string) (err error) {
	if len(text) == 3 {
		if text[1] != "is" {
			return errors.New("invalid")
		}
		err = registerCurrency(text[0], text[2])
		if err != nil {
			return err
		}
	}

	if text[len(text)-1] == "Credits" {
		logrus.Info("add item price")
		err = addItemPrice(text)
		if err != nil {
			return err
		}
	}

	if text[0] == "how" {
		text = text[:len(text)-1] // remove ? todo: need to check if ? as the last sentence
		if text[1] == "much" {
			if text[2] == "is" {
				// calculate price
				var values []string
				for i :=0; i<len(text); i++ {
					if i > 2 {
						values = append(values, text[i])
					}
				}
				price, errCalculatePrice := calculatePrice(values)
				if errCalculatePrice != nil {
					return errCalculatePrice
				}
				logrus.Infof("values %v, price %v", values, price)
			}
		} else if text[1] == "many" && text[2] == "Credits"  && text[3] == "is" {
			// todo: distinguish values, item
			itemName := text[len(text)-1]
			logrus.Infof("item name %v", itemName)
			itemPrice := text[:len(text)-1][4:len(text[:len(text)-1])]
			logrus.Warnf("itemPrice %v", itemPrice)

			// check if unit exist

			// check if item exist

			// calculate price of the item
			price, errCalculatePrice := calculatePrice(itemPrice)
			if errCalculatePrice != nil {
				return errCalculatePrice
			}
			logrus.Infof("price %v", price)

		}

	}

	return nil
}

func registerCurrency(name, value string) error {
	logrus.Infof("register currency")
	for val := range currency {
		if val == value {
			return errors.New("oops listed")
		} else if currency[val] == name {
			return errors.New("oops name listed")
		}
	}
	currency[value] = name
	return nil
}

func addItemPrice(text []string) error {
	value := text[len(text)-2]
	logrus.Infof("add item price %v", value)

	calculatePrice(text)
	return nil
}

func calculatePrice(text []string) (int, error) {
	// todo: logic to calculate price
	return 0, nil
}