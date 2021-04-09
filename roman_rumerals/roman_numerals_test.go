package roman_rumerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	arabic uint16
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{3, "III"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"}}

func TestRomanNumerals(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %v", tt.arabic, tt.roman), func(t *testing.T) {
			got := ConvertToRoman(tt.arabic)

			if got != tt.roman {
				t.Errorf("got %q, want %q", got, tt.roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", tt.roman, tt.arabic), func(t *testing.T) {
			got := ConvertToArabic(tt.roman)

			if got != tt.arabic {
				t.Errorf("got %d, want %d", got, tt.arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)

		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}
