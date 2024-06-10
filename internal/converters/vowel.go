package converters // This declares the package name as "converters"

// This imports the required packages
import (
	"strings"
)

func Vowel(input string) string { // This function takes a string input and returns a string
	vow := "aeiouh"                 // This string contains all the vowels (including 'h' for some reason)
	w := strings.Fields(input)      // This splits the input string into a slice of strings based on whitespace
	for i := 0; i < len(w)-1; i++ { // This loop iterates over the slice of strings, excluding the last element
		if strings.ToLower(w[i]) == "a" && strings.ContainsAny(vow, strings.ToLower(string(w[i+1][0]))) { // This checks if the current string is "a" and if the first character of the next string is a vowel
			w[i] = w[i] + "n" // If the condition is true, it appends "n" to the current string
		}
	}
	return strings.Join(w, " ") // This joins the modified slice of strings back into a single string, separated by spaces
}
