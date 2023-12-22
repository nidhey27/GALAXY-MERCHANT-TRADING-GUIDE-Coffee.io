package file

import (
	"fmt"
	"strings"

	"github.com/nidhey27/coffee-assignment/helpers"
)

var valueMap = make(map[string]float64)
var romanMap = make(map[string]string)

const many = "how many"
const much = "how much"
const credits = "Credits"
const hasLess = "has less"
const hasMore = "has more"
const larger = "larger than"
const smaller = "smaller than"
const questionMark = "?"

func ParseInput(data string) (map[string]float64, map[string]string) {
	splittedInput := strings.Split(data, "\n")
	for _, input := range splittedInput {
		input = strings.TrimSpace(input)

		if strings.Contains(input, questionMark) {
			if strings.Contains(input, much) {
				result := helpers.CalculateHowMuch(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, many) {
				result := helpers.CalculateHowMany(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, hasMore) {
				result := helpers.HasMore(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, hasLess) {
				result := helpers.HasLess(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, larger) {
				result := helpers.LargerThan(input, romanMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, smaller) {
				result := helpers.SmallerThan(input, romanMap)
				fmt.Println(input, "-", result)
			} else {
				result := helpers.ReturnWrong()
				fmt.Println(input, "-", result)
			}

		} else {
			if strings.Contains(input, credits) {
				key, value := helpers.StoreValueMap(input, romanMap)
				valueMap[key] = value
			} else {
				key, value := helpers.StoreRomanMap(input)
				if key != "" && value != "" {
					romanMap[key] = value
				}

			}
		}
	}

	return valueMap, romanMap
}
