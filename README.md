# Galaxy Merchant Trading Guide

## Overview
This is a solution for Galaxy Merchant Trading Guide.

The program categorize input into following transactions:
* Adding item into the store by format _pish pish Iron is 3910 Credits_
* Convert intergalactic units by format _how much is pish tegj glob glob ?_
* Calculate the price of an item by format _how many Credits is glob prok Silver ?_

According to the transactions, the solution contains of 3 parts: 
* currency: handle anything related to currency e.g registering currency, converting intergalactic currency
* price: handle items e.g adding items, calculate the price of items
* utils: process the input

## Development

The solution is using local storage to save the data.

For local development, you can perform _go run main.go_ under cmd directory. Note that to perform converting
intergalactic units as well as to calculate the price of items, you need to manually seeding the data through
the terminal.

In order to calculate the conversion of intergalactic unit into number, you need to first add intergalactic
units you want to use, e.g _glob is I_. 

In order to calculate the price of an item, you need to:
1. add intergalactic unit you want to use, e.g _pish is X_
2. add item price into store, e.g _pish pish Iron is 3910 Credits_

Then you can perform intergalactic unit conversion and item price calculation.

# Problem Description

You decided to give up on earth after the latest financial collapse left 99.99% of the earth's
population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your
account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell
common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy
requires you to convert numbers and units, and you decided to write a program to help you.The
numbers used for intergalactic transactions follows similar convention to the roman numerals and
you have painstakingly collected the appropriate translation between them. Roman numerals are
based on seven symbols:

Symbol | Value
-------|----------
  I    |   1
  V    |   5
  X    |   10
  L    |   50
  C    |   100
  D    |   500
  M    |   1,000
  
Numbers are formed by combining symbols together and adding the values. For example, MMVI is
1000 + 1000 + 5 + 1 = 2006. Generally, symbols are placed in order of value, starting with the
largest values. When smaller values precede larger values, the smaller values are subtracted from
the larger values, and the result is added to the total. 

For example MCMXLIV = 1000 + (1000 − 100) + (50 − 10) + (5 − 1) = 1944.

The symbols "I", "X", "C", and "M" can be repeated three times in succession, but no more. (They
may appear four times if the third and fourth are separated by a smaller value, such as XXXIX.)
"D", "L", and "V" can never be repeated.

"I" can be subtracted from "V" and "X" only. "X" can be subtracted from "L" and "C" only. "C" can
be subtracted from "D" and "M" only. "V", "L", and "D" can never be subtracted.

Only one small-value symbol may be subtracted from any large-value symbol.

A number written in Arabic numerals can be broken into digits. For example, 1903 is composed of
1, 9, 0, and 3. To write the Roman numeral, each of the non-zero digits should be treated separately.
In the above example, 1,000 = M, 900 = CM, and 3 = III. Therefore, 1903 = MCMIII.

**Test Input**
* glob is I
* prok is V
* pish is X
* tegj is Lglob glob Silver is 34 Credits
* glob prok Gold is 57800 Credits
* pish pish Iron is 3910 Credits
* how much is pish tegj glob glob ?
* how many Credits is glob prok Silver ?
* how many Credits is glob prok Gold ?
* how many Credits is glob prok Iron ?
* how much wood could a woodchuck chuck if a woodchuck could chuck wood ?

**Expecting Output**
* pish tegj glob glob is 42
* glob prok Silver is 68 Credits
* glob prok Gold is 57800 Credits
* glob prok Iron is 782 Credits
* I have no idea what you are talking about