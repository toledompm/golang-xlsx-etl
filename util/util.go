package util

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

//NormalizeMapKeys applies the util.Normalize function to each key in the given map
func NormalizeMapKeys(denormalizedMap map[string]string) map[string]string {
	for key, value := range denormalizedMap {
		delete(denormalizedMap, key)
		denormalizedMap[Normalize(key)] = value
	}

	return denormalizedMap
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
