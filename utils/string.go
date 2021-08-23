package utils

import (
	"fmt"
	"strings"
)

// ConvertSpecifyWordToFullUpperCase convert golang recommand spell
// ex: Id => ID, url => URL
func ConvertSpecifyWordToFullUpperCase(input, from, to string) string {
	if strings.HasSuffix(input, from) {
		if len(input) == 2 {
			input = to
		} else {
			idIndex := strings.LastIndex(input, from)
			input = fmt.Sprintf("%s%s", input[0:idIndex], to)
		}
	}
	return input
}
