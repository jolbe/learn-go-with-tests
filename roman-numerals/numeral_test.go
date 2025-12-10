package romannumerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
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
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %q; want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d; want %d", got, test.Arabic)
			}
		})
	}
}

func TestConvertingToArabicRecursive(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabicRecursive(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d; want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

func TestNoMoreThanThreeSameSymbols(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)

		var last rune
		var count int
		for _, r := range roman {
			if r == last {
				count++
				if count > 3 {
					return false
				}
			} else {
				last = r
				count = 1
			}
		}

		return true
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("too many repeated symbols in roman numeral:", err)
	}
}

// “Only valid subtractors” property — i.e. whenever a smaller-value symbol precedes a larger-value symbol,
// it must be one of I, X, C (or allowed ones)
func TestValidSubtractivePairs(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic == 0 || arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)

		for i := 0; i+1 < len(roman); i++ {
			curr := roman[i]
			next := roman[i+1]
			if valueOf(curr) < valueOf(next) {
				switch curr {
				case 'I', 'X', 'C':
					// OK
				default:
					return false
				}
			}
		}
		return true
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("invalid subtractive symbol used:", err)
	}
}

func valueOf(r byte) int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

func BenchmarkConvertToArabic(b *testing.B) {
	for b.Loop() {
		ConvertToArabic("MCMLXXXIV")
	}
}

func BenchmarkConvertToArabicRecursive(b *testing.B) {
	for b.Loop() {
		ConvertToArabicRecursive("MCMLXXXIV")
	}
}
