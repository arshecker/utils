package random

import (
	"gopackages/wordz"
)

func Digit(new_numbers []string) string {
	wordz.Words = new_numbers
	return wordz.Random()
}
