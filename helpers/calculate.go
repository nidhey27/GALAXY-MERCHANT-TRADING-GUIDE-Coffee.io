package helpers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nidhey27/coffee-assignment/common"
	"github.com/nidhey27/coffee-assignment/utils"
)

func CalculateHowMuch(input string, romanMap map[string]string, valueMap map[string]float64) string {
	result := ""
	tempRoman := ""
	tempQuestion := ""
	if strings.Contains(input, common.IS) {
		splitted := strings.Split(input, common.IS)

		splittedValue := strings.Split(splitted[1], common.SPACE)

		for index, value := range splittedValue {
			_, ok := romanMap[value]
			if ok {
				tempRoman += romanMap[value]
				tempQuestion += value
				if index == len(splittedValue)-2 {
					numeral, err := NumeralConvert(tempRoman)
					if err != nil {
						return err.Error()
					}
					tempQuestion += fmt.Sprintf(" is %d", numeral)
				} else {
					tempQuestion += " "
				}
			}

		}
		result = tempQuestion
	} else {
		result = common.WRONG
	}
	return result
}

func CalculateHowMany(input string, romanMap map[string]string, valueMap map[string]float64) string {
	result := ""
	tempRoman := ""
	tempQuestion := ""

	if strings.Contains(input, common.IS) {
		splitted := strings.Split(input, common.IS)

		splittedValue := strings.Split(splitted[1], common.SPACE)
		for index, value := range splittedValue {
			_, ok := romanMap[value]
			if ok {
				tempRoman += romanMap[value]
				tempQuestion += value
				tempQuestion += " "
			} else {
				_, ok := valueMap[value]
				if ok {
					if index == len(splittedValue)-2 {
						numeral, err := NumeralConvert(tempRoman)
						if err != nil {
							return err.Error()
						}
						finalNumeral := float64(numeral) * valueMap[value]

						tempQuestion += value
						tempQuestion += fmt.Sprintf(" is %f %s", finalNumeral, common.CREDITS)
					} else {
						tempQuestion += " "
					}
				}
			}
		}
		result = tempQuestion
	} else {
		result = common.WRONG
	}
	return result
}

func ReturnWrong() string {
	result := common.WRONG
	return result
}

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
			numeral, err := NumeralConvert(tempRoman)
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

func extractMetalNames(splitParam, input string) (string, string, bool) {
	var metal1, metal2 string
	var found bool

	if strings.Contains(input, splitParam) {
		splitted := strings.Split(input, splitParam)

		// Extract metal names
		metals := strings.Fields(splitted[0])
		if len(metals) > 0 {
			metal1 = metals[len(metals)-1]
			found = true
		}

		metals = strings.Fields(splitted[1])
		if len(metals) > 0 {
			metal2 = metals[len(metals)-2]
			found = true
		}
	}

	return metal1, metal2, found
}

func HasMore(input string, romanMap map[string]string, valueMap map[string]float64) string {
	result := ""
	metal1 := ""
	metal2 := ""

	if strings.Contains(input, common.CREDITS_THAN) {
		metal1, metal2, _ = extractMetalNames(common.CREDITS_THAN, input)

		credits1 := valueMap[metal1]
		credits2 := valueMap[metal2]

		if credits1 > credits2 {
			result = fmt.Sprintf("%s has more Credits than %s", metal1, metal2)
		} else {
			result = fmt.Sprintf("%s has less Credits than %s", metal1, metal2)
		}
	} else {
		result = ReturnWrong()
	}
	return result

}

func HasLess(input string, romanMap map[string]string, valueMap map[string]float64) string {
	result := ""
	metal1 := ""
	metal2 := ""

	if strings.Contains(input, common.CREDITS_LESS) {
		metal1, metal2, _ = extractMetalNames(common.CREDITS_LESS, input)

		credits1 := valueMap[metal1]
		credits2 := valueMap[metal2]

		if credits1 > credits2 {
			result = fmt.Sprintf("%s has more Credits than %s", metal1, metal2)
		} else {
			result = fmt.Sprintf("%s has less Credits than %s", metal1, metal2)
		}
	} else {
		result = ReturnWrong()
	}
	return result
}

func LargerThan(input string, romanMap map[string]string) string {
	result := ""
	var metal1 []string
	var metal2 []string

	if strings.Contains(input, common.LARGER_THAN) {
		// Extract metal names
		metal1, metal2, _ = extractRomanNames(common.LARGER_THAN, input, romanMap)

		credits1 := 0
		credits2 := 0

		for _, val := range metal1 {
			credits1 += utils.RomanToArabic(romanMap[val])
		}

		for _, val := range metal2 {
			credits2 += utils.RomanToArabic(romanMap[val])
		}

		if credits1 < credits2 {
			result = fmt.Sprintf("%s is smaller than %s", strings.Join(metal1, " "), strings.Join(metal2, " "))
		} else {
			result = fmt.Sprintf("%s is larger than %s", strings.Join(metal1, " "), strings.Join(metal2, " "))
		}

	} else {
		result = ReturnWrong()
	}
	return result
}

func SmallerThan(input string, romanMap map[string]string) string {
	result := ""
	var metal1 []string
	var metal2 []string

	if strings.Contains(input, common.SMALLER_THAN) {
		// Extract metal names
		metal1, metal2, _ = extractRomanNames(common.SMALLER_THAN, input, romanMap)

		credits1 := 0
		credits2 := 0

		for _, val := range metal1 {
			credits1 += utils.RomanToArabic(romanMap[val])
		}

		for _, val := range metal2 {
			credits2 += utils.RomanToArabic(romanMap[val])
		}

		if credits1 < credits2 {
			result = fmt.Sprintf("%s is smaller than %s", strings.Join(metal1, " "), strings.Join(metal2, " "))
		} else {
			result = fmt.Sprintf("%s is larger than %s", strings.Join(metal1, " "), strings.Join(metal2, " "))
		}

	} else {
		result = ReturnWrong()
	}
	return result
}

func extractRomanNames(splitParam, input string, romanMap map[string]string) ([]string, []string, bool) {
	var metal1, metal2 []string
	var found bool

	if strings.Contains(input, splitParam) {
		splitted := strings.Split(input, splitParam)

		for _, val := range strings.Split(splitted[0], " ") {
			if _, ok := romanMap[string(val)]; ok {
				metal1 = append(metal1, val)
			}
		}
		for _, val := range strings.Split(splitted[1], " ") {
			if _, ok := romanMap[string(val)]; ok {
				metal2 = append(metal2, val)
			}
		}
	}

	return metal1, metal2, found
}
