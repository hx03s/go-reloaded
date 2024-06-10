package converters

// Import required packages
import (
	"strconv"
	"strings"
	"unicode"
)

// Upper converts a string to uppercase
func Upper(input string) string {
	upper := strings.ToUpper(input) // Convert input string to uppercase using strings.ToUpper
	return upper                    // Return the uppercase string
}

// Low converts a string to lowercase
func Low(input string) string {
	lower := strings.ToLower(input) // Convert input string to lowercase using strings.ToLower
	return lower                    // Return the lowercase string
}

// Caps capitalizes the first letter of a string
func Caps(input string) string {
	fL := strings.ToUpper(string(input[0])) // Convert the first character of input to uppercase
	rest := strings.ToLower(input[1:])      // Convert the remaining characters to lowercase
	return fL + rest                        // Return the capitalized string by concatenating the uppercase first letter and lowercase rest
}

// Num converts a string containing numbers to an integer
func Num(input string) int {
	fun := func(n rune) bool {
		return !unicode.IsNumber(n) // Define a function to check if a rune is not a number
	}
	va := strings.FieldsFunc(input, fun)     // Split the input string by non-numeric characters using strings.FieldsFunc
	j := strings.Join(va, "")                // Join the numeric parts of the input string
	number, _ := strconv.ParseInt(j, 10, 64) // Convert the joined numeric string to an integer using strconv.ParseInt
	return int(number)                       // Return the integer value
}
