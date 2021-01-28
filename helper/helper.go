package helper

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

//Normalize tries to normalize and lower some given text
func Normalize(text string) string {
	tChain := transform.Chain(
		norm.NFD,
		runes.Remove(runes.In(unicode.Mn)),
		norm.NFC,
	)

	normalizedText, _, _ := transform.String(tChain, text)

	return strings.ToLower(normalizedText)
}
