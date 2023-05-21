package roman

import "errors"

// numerals
// converters
// numconvert
// romannumerals

var (
	nums   = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func ToRoman(num int) (string, error) {
	if num < 1 || num > 3999 {
		return "", errors.New("input must be a number between 1 and 3999")
	}

	roman := ""
	l := len(nums)

	for i := 0; i < l; {
		if num >= nums[i] {
			num -= nums[i]
			roman += romans[i]
		} else {
			i += 1
		}
	}
	// for i := 0; i < l; i++ {
	// 	for num >= nums[i] {
	// 		num -= nums[i]
	// 		roman += romans[i]
	// 	}
	// }
	return roman, nil
}
