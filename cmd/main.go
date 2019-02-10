package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"adesnia/galaxy-merchant-trading-guide/pkg/currency"
	"adesnia/galaxy-merchant-trading-guide/pkg/price"
	"adesnia/galaxy-merchant-trading-guide/pkg/utils"

	"github.com/sirupsen/logrus"
)

func main() {
	currency.Conversions = make(map[string]string)
	currency.RegisterUnits()
	price.Items = make(map[string]int)

	for {
		{
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter input: ")
			input, _ := reader.ReadString('\n')
			text := strings.Split(input, " ")
			if len(text) <= 2 {
				logrus.Errorf("invalid input")
			}
			out, err := utils.Parse(text)
			if err != nil {
				logrus.Errorf("unable to parse input: %v", err)
			}
			fmt.Printf(out)
		}
	}

	return
}
