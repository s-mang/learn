package learn

import (
	"regexp"
	"strings"
)

// Translate translates an English/American phrase into Pig Latin
// https://www.urbandictionary.com/define.php?term=pig%20latin
func Translate(phrase string) string {
	// First we split the phrase on the space (' ').
	// This way we will have an array (or in Go, 'a slice')
	// of words.
	// Eg. "hello world" becomes []string{"hello", "world"}
	// (golang.org/pkg/strings)
	words := strings.Split(phrase, " ")

	// Then we will translate each word one at a time,
	// and store the result back in the 'words' slice.
	for i, word := range words {
		words[i] = translateWord(word)
	}

	// And finally, we stitch the words back together with a
	// space in between with strings.Join()
	return strings.Join(words, " ")
}

// translateWord translates an English/American word into PigLatin
// https://www.urbandictionary.com/define.php?term=pig%20latin
// translateWord is unexported because it should only be called internally
// from within this package.
// The argument (word) should be one word with no spaces.
// This rule is enforced by the above function `Translate`,
// this function's caller. (hence why this function is unexported).
func translateWord(word string) string {
	// edge case
	if word == "" {
		return ""
	}

	// First we're going to look for punctuation.
	// If there is punctuation at the end of a word,
	// like a comma (,) or a period (.), then we want to hold
	// onto that for a sec. The punctuation needs to stay at
	// the end of the word still after the translation.

	// we'll keep a variable 'endPunc' which will be either:
	// - an empty string, if there is no punctuation at the end of the word, OR
	// - the punctuation at the end of the word.
	endPunc := ""
	endPunc, word = getEndPunctuation(word)

	firstChar := string(word[0])

	// Check if first character is capital:
	// if the letter capitalized is equal to the letter itself,
	// then the letter must already be capitalized.
	var isFirstCharCap bool // note: default value ('zero value' in Go) is *false* for a bool
	if strings.ToUpper(firstChar) == firstChar {
		isFirstCharCap = true
		firstChar = strings.ToLower(firstChar)
	}

	rest := string(word[1:])

	// if the first letter was originally capitalized,
	// we should capitalize the new first letter.
	if isFirstCharCap {
		rest = strings.ToUpper(string(rest[0])) + rest[1:]
	}

	return rest + firstChar + "ay" + endPunc
}

// getEndPunctuation returns any punctuation at the end of the word,
// and the word without the punctuation at the end.
// If there is no end punctuation, the return values will be
// the empty string and the original input word.
func getEndPunctuation(word string) (punc, w string) {
	// In order to determine whether or not the last N characters
	// of the string are punctuation , I am going to generalize
	// and say that anything which is not a letter or a number must be
	// a punctuation mark.
	// We use a regular expression find the non-letters and non-numbers
	// at the end of the string (golang.org/pkg/regexp):
	rgx := regexp.MustCompile("[^A-Za-z0-9]+$")
	// notes on above regex:
	// - the '[^]' character (inside a []) above means 'not' in regexp (like the '!' in '!=')
	// - the '$' character means 'end of the string'
	// - the '+' character means '1 or more of whats in the brackets'

	// there should only be one (second arg)
	puncs := rgx.FindAllString(word, 1)
	if len(puncs) == 0 {
		return "", word
	}

	punc = puncs[0]

	// otherwise...
	outWord := word[:len(word)-len(punc)]

	// eg. input: "heey!!!" would output: []string{"heey", "!!!"}
	return punc, outWord
}
