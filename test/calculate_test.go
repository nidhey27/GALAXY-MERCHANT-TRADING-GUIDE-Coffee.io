package test

import (
	"testing"

	"github.com/nidhey27/coffee-assignment/internal/calculate"
)

func assertEqual(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %q, but got %q", expected, actual)
	}
}

func TestCalculateHowMuch(t *testing.T) {
	romanMap := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}

	valueMap := map[string]float64{
		"Silver": 17,
		"Gold":   14450,
		"Iron":   195.5,
	}

	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "how much is pish tegj glob glob ?",
			output: "pish tegj glob glob is 42",
		},
		{
			input:  "how much is glob prok pish tegj glob glob ?",
			output: "Requested number is in invalid format",
		},
		{
			input:  "how much is glob prok ?",
			output: "glob prok is 4",
		},
	}

	for _, tc := range testCases {
		result := calculate.CalculateHowMuch(tc.input, romanMap, valueMap)
		if result != tc.output {
			t.Errorf("Input: %s, Expected Output: %s, Got: %s", tc.input, tc.output, result)
		}
	}
}

func TestCalculateHowMany(t *testing.T) {
	romanMap := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}

	valueMap := map[string]float64{
		"Silver": 17,
		"Gold":   14450,
		"Iron":   195.5,
	}

	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "how many Credits is glob prok Silver ?",
			output: "glob prok Silver is 68.000000 Credits",
		},
		{
			input:  "how many Credits is glob glob Gold ?",
			output: "glob glob Gold is 28900.000000 Credits",
		},
		{
			input:  "how many Credits is glob glob glob glob glob glob Gold ?",
			output: "Requested number is in invalid format",
		},
		{
			input:  "how many Credits is pish tegj glob Iron ?",
			output: "pish tegj glob Iron is 8015.500000 Credits",
		},
	}

	for _, tc := range testCases {
		result := calculate.CalculateHowMany(tc.input, romanMap, valueMap)
		expectedOutput := tc.output
		if tc.output != "Requested number is in invalid format" {
			expectedOutput = tc.output
		}

		if result != expectedOutput {
			t.Errorf("Input: %s, Expected Output: %s, Got: %s", tc.input, expectedOutput, result)
		}
	}
}

func TestHasMore(t *testing.T) {
	romanMap := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}

	valueMap := map[string]float64{
		"Silver": 17,
		"Gold":   14450,
		"Iron":   195.5,
	}

	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "Does glob glob Gold has more Credits than pish tegj glob glob Iron ?",
			output: "glob glob Gold has more Credits than pish tegj glob glob Iron",
		},
		{
			input:  "Does pish tegj glob glob Iron has more Credits than glob glob Gold ?",
			output: "pish tegj glob glob Iron has less Credits than glob glob Gold",
		},
		{
			input:  "Does glob prok Silver has more Credits than pish tegj glob Iron ?",
			output: "glob prok Silver has less Credits than pish tegj glob Iron",
		},
		{
			input:  "Does glob glob Gold has more Credits than glob prok Silver ?",
			output: "glob glob Gold has more Credits than glob prok Silver",
		},
		{
			input:  "Is glob prok larger than pish pish ?",
			output: "I have no idea what you are talking about",
		},
	}

	for _, tc := range testCases {
		result := calculate.HasMore(tc.input, romanMap, valueMap)
		if result != tc.output {
			t.Errorf("Input: %s, Expected Output: %s, Got: %s", tc.input, tc.output, result)
		}
	}
}

func TestHasLess(t *testing.T) {
	romanMap := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}

	valueMap := map[string]float64{
		"Silver": 17,
		"Gold":   14450,
		"Iron":   195.5,
	}

	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "Does glob glob Gold has less Credits than pish tegj glob glob Iron ?",
			output: "glob glob Gold has more Credits than pish tegj glob glob Iron",
		},
		{
			input:  "Does pish tegj glob glob Iron has less Credits than glob glob Gold ?",
			output: "pish tegj glob glob Iron has less Credits than glob glob Gold",
		},
		{
			input:  "Does glob prok Silver has less Credits than pish tegj glob Iron ?",
			output: "glob prok Silver has less Credits than pish tegj glob Iron",
		},
		{
			input:  "Does glob glob Gold has less Credits than glob prok Silver ?",
			output: "glob glob Gold has more Credits than glob prok Silver",
		},
		{
			input:  "Is glob prok smaller than pish pish ?",
			output: "I have no idea what you are talking about",
		},
	}

	for _, tc := range testCases {
		result := calculate.HasLess(tc.input, romanMap, valueMap)
		if result != tc.output {
			t.Errorf("Input: %s, Expected Output: %s, Got: %s", tc.input, tc.output, result)
		}
	}
}

func TestLargerThan(t *testing.T) {
	romanMap := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}

	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "Is glob prok larger than pish pish ?",
			output: "glob prok is smaller than pish pish",
		},
		{
			input:  "Is pish tegj glob glob larger than glob prok pish ?",
			output: "pish tegj glob glob is larger than glob prok pish",
		},
		{
			input:  "Is glob glob glob larger than glob prok ?",
			output: "glob glob glob is smaller than glob prok",
		},
	}

	for _, tc := range testCases {
		result := calculate.LargerThan(tc.input, romanMap)
		if result != tc.output {
			t.Errorf("Input: %s, Expected Output: %s, Got: %s", tc.input, tc.output, result)
		}
	}
}

func TestSmallerThan(t *testing.T) {
	romanMap := map[string]string{
		"glob": "I",
		"prok": "V",
		"pish": "X",
		"tegj": "L",
	}

	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "Is tegj glob glob smaller than glob prok ?",
			output: "tegj glob glob is larger than glob prok",
		},
		{
			input:  "Is glob glob glob smaller than glob prok pish ?",
			output: "glob glob glob is smaller than glob prok pish",
		},
		{
			input:  "Istegj glob glob smaller than glob prok pish pish ?",
			output: "glob glob is smaller than glob prok pish pish",
		},
	}

	for _, tc := range testCases {
		result := calculate.SmallerThan(tc.input, romanMap)
		if result != tc.output {
			t.Errorf("Input: %s, Expected Output: %s, Got: %s", tc.input, tc.output, result)
		}
	}
}