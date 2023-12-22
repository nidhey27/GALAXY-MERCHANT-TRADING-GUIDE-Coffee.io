package storage

import (
	"strconv"
	"strings"

	"github.com/nidhey27/coffee-assignment/common"
	"github.com/nidhey27/coffee-assignment/utils"
)

func StoreValueMap(input string, romanMap map[string]string) (string, float64) {
	splitted := strings.Split(input, common.IS)
	splittedRoman := strings.Split(splitted[0], common.SPACE)

	tempRoman := ""
	key := ""
	value := 0.00

	credits := strings.Split(splitted[1], common.SPACE)
	creditValue, err := strconv.Atoi(credits[0])
	if err != nil {
		return err.Error(), 0
	}

	for _, char := range splittedRoman {
		_, ok := romanMap[char]
		if ok {
			tempRoman += romanMap[char]
		} else {
			key = char
			numeral, err := utils.NumeralConvert(tempRoman)
			if err != nil {
				return err.Error(), 0
			}
			value = float64(creditValue) / float64(numeral)
		}

	}
	return key, value
}

func StoreRomanMap(input string) (string, string) {
	var key string = ""
	var value string = ""
	splitted := strings.Split(input, common.IS)
	if len(splitted) > 1 {
		roman := strings.TrimSpace(splitted[1])
		key = splitted[0]
		value = roman
	}
	return key, value
}
