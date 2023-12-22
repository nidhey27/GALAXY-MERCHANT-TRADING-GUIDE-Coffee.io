package file

import (
	"fmt"
	"strings"

	"github.com/nidhey27/coffee-assignment/common"
	calculate "github.com/nidhey27/coffee-assignment/internal/calcluate"
	"github.com/nidhey27/coffee-assignment/internal/storage"
)

var valueMap = make(map[string]float64)
var romanMap = make(map[string]string)

func ParseInput(data string) (map[string]float64, map[string]string) {
	splittedInput := strings.Split(data, "\n")
	for _, input := range splittedInput {
		input = strings.TrimSpace(input)

		if strings.Contains(input, common.QUESTION_MARK) {
			if strings.Contains(input, common.MUCH) {
				result := calculate.CalculateHowMuch(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, common.MANY) {
				result := calculate.CalculateHowMany(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, common.HAS_MORE) {
				result := calculate.HasMore(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, common.HAS_LESS) {
				result := calculate.HasLess(input, romanMap, valueMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, common.LARGER_THAN) {
				result := calculate.LargerThan(input, romanMap)
				fmt.Println(input, "-", result)
			} else if strings.Contains(input, common.SMALLER_THAN) {
				result := calculate.SmallerThan(input, romanMap)
				fmt.Println(input, "-", result)
			} else {
				result := calculate.ReturnWrong()
				fmt.Println(input, "-", result)
			}

		} else {
			if strings.Contains(input, common.CREDITS) {
				key, value := storage.StoreValueMap(input, romanMap)
				valueMap[key] = value
			} else {
				key, value := storage.StoreRomanMap(input)
				if key != "" && value != "" {
					romanMap[key] = value
				}

			}
		}
	}

	return valueMap, romanMap
}
