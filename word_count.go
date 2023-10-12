// wordcount package count how many times each word occurs in a subtitle of a drama.
package wordcount

import (
	"strings"
)

// Frequency is a map of words to their frequency.
type Frequency map[string]int

// Increment increments the count for the given word.
func (f Frequency) Increment(word string) {
	if word != "" {
		f[word]++
	}
}

// isAlphanumeric returns true if the given rune is a letter or digit.
func isAlphanumeric(v rune) bool {
	if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') && (v < '0' || v > '9') {
		return false
	}
	return true
}

// isContraction returns true if the given index is a contraction.
func isContraction(i int, word string) bool {
	return i > 0 && i < len(word)-1 && word[i] == '\'' && isAlphanumeric(rune(word[i-1])) && isAlphanumeric(rune(word[i+1]))
}

// WordCount counts the number of occurrences of each word in a phrase.
func WordCount(phrase string) Frequency {
	var f = make(Frequency)
	var word string
	for i, v := range phrase {
		if !isAlphanumeric(v) && !isContraction(i, phrase) {
			f.Increment(word)
			word = ""
		} else {
			word += strings.ToLower(string(v))
		}
	}
	f.Increment(word)
	return f
}
