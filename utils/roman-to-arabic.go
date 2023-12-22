package utils

// Utility function to convert Roman numerals to Arabic
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
