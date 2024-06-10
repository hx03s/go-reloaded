package converters

// Import required packages
import (
	"fmt"
	"regexp"
	"strings"
)

// Punc function removes extra spaces around punctuation marks
func Punc(input string) string {
	// Compile a regular expression to match punctuation marks with optional spaces
	r := regexp.MustCompile(`\s*([.,!?:;]+)`)
	// Replace matched patterns with just the punctuation marks
	ha := r.ReplaceAllString(input, "$1")
	// Call PuncSpace function to add spaces after punctuation marks
	ha = PuncSpace(ha)
	return ha
}

// PuncSpace function adds spaces after punctuation marks
func PuncSpace(input string) string {
	// Compile a regular expression to match punctuation marks followed by a word character
	rr := regexp.MustCompile(`([.,!?:;])(\w)`)
	// Replace matched patterns with punctuation mark followed by a space and the word character
	ha := rr.ReplaceAllString(input, "$1 $2")
	return ha
}

// Quot function replaces single-quoted words with the word without spaces
func Quot(input string) string {
	// Compile a regular expression to match single-quoted words
	zz := regexp.MustCompile(`'\s*\w+\s*'`)
	// Replace matched patterns using the sen function
	ha := zz.ReplaceAllStringFunc(input, sen)
	return ha
}

// sen function removes spaces from a string
func sen(input string) string {
	// Replace all spaces with an empty string
	input = strings.ReplaceAll(input, " ", "")
	return input
}

// Quott function replaces various patterns of single-quoted words and punctuation
func Quott(input string) string {
	// Compile a regular expression to match various patterns of single-quoted words and punctuation
	zz := regexp.MustCompile(`'\s+\w+|\w+\s+'|'\s+|\.\s+'`)
	// Print the input string for debugging
	fmt.Printf("(%s)\n", input)
	// Replace matched patterns using the senn function
	ha := zz.ReplaceAllStringFunc(input, senn)
	return ha
}

// senn function removes spaces from a string and prints the input for debugging
func senn(input string) string {
	// Print the input string for debugging
	fmt.Printf("(%s)2\n", input)
	// Replace all spaces with an empty string
	input = strings.ReplaceAll(input, " ", "")
	return input
}
