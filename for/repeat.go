package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeated strings.Builder
	for range times {
		repeated.WriteString(character)
	}
	return repeated.String()
}
