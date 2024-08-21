package random

import (
	"gopackages/wordz"
)

func City(new_words []string) string {
	wordz.Words = new_words
	return wordz.Random()
}
