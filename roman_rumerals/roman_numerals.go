package roman_rumerals

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumeralArray []RomanNumeral

var allRomanNumerals = RomanNumeralArray{
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (a *RomanNumeralArray) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, r := range *a {
		if r.Symbol == symbol {
			return r.Value
		}
	}

	return 0
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	total := uint16(0)

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, roman) {
			if value := allRomanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	return total
}

func couldBeSubtractive(i int, roman string) bool {
	return i+1 < len(roman) &&
		allRomanNumerals.ValueOf(roman[i]) < allRomanNumerals.ValueOf(roman[i+1])
}
