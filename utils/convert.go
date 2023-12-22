package utils

import (
	"errors"
	"strings"
)

const errMessageInvalid = "Requested number is in invalid format"

var symbolMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var substractMap = map[string]string{
	"V": "I",
	"X": "I",
	"L": "X",
	"C": "X",
	"D": "C",
	"M": "C",
}

func NumeralConvert(roman string) (int, error) {

	var numeral int
	var lastRomanValue int
	repeatedMap := make(map[string]int)

	roman = strings.ToUpper(roman)

	for i := 0; i < len(roman); i++ {
		romanString := string(roman[i])
		value := symbolMap[romanString]

		if value == 0 {
			return -1, errors.New(errMessageInvalid)
		}

		next := i + 1
		valueNext := 0
		repeatedMap[string(roman[i])] += 1

		if repeatedMap["I"] > 3 ||
			repeatedMap["X"] > 3 ||
			repeatedMap["C"] > 3 ||
			repeatedMap["M"] > 3 ||
			repeatedMap["D"] > 1 ||
			repeatedMap["L"] > 1 ||
			repeatedMap["V"] > 1 {
			return -1, errors.New(errMessageInvalid)
		}

		if next < len(roman) {
			romanNextString := string(roman[next])
			valueNext = symbolMap[romanNextString]

			if value < valueNext {
				if romanString == "V" || romanString == "L" || romanString == "D" {
					return -1, errors.New(errMessageInvalid)
				}
				if lastRomanValue < valueNext-value && i != 0 {
					return -1, errors.New(errMessageInvalid)
				}

				if romanString != substractMap[romanNextString] {
					return -1, errors.New(errMessageInvalid)
				}

				numeral += valueNext - value
				i++
				continue
			}
		}

		numeral += value
		lastRomanValue = value
	}
	return numeral, nil
}

func GetSymbolMap() map[string]int {
	return symbolMap
}

func RomanToArabic(roman string) int {
	romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	arabic := 0

	for i := 0; i < len(roman); i++ {
		value := romanNumerals[rune(roman[i])]

		if i+1 < len(roman) && romanNumerals[rune(roman[i+1])] > value {
			arabic -= value
		} else {
			arabic += value
		}
	}

	return arabic
}
