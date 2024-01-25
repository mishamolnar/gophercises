package phone_number_normalizer

import "unicode"

func Normalize(rawNumber string) string {
	sl := make([]rune, 0)
	for _, ch := range rawNumber {
		if unicode.IsDigit(ch) {
			sl = append(sl, ch)
		}
	}
	return string(sl)
}
