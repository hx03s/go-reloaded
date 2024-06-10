package converters

// This line declares the package name as "converters"

import (
	"fmt"
	"os"
	"strconv"
)

// These lines import the necessary packages for the code to function:
// "fmt" for printing output
// "os" for interacting with the operating system
// "strconv" for converting between strings and other data types

func HexTodec(input string) string {
	// This function takes a string input and returns a string representing the decimal equivalent of the hexadecimal input
	decimal, err := strconv.ParseInt(input, 16, 64)
	// This line converts the input string to an integer using base 16 (hexadecimal) and stores the result in the "decimal" variable
	// The "err" variable will hold any error that occurs during the conversion
	if err != nil {
		// This block checks if an error occurred during the conversion
		fmt.Println("err")
		// If an error occurred, it prints "err" to the console
		os.Exit(0)
		// And then exits the program with a status code of 0
	}
	return strconv.Itoa(int(decimal))
	// If no error occurred, the function converts the decimal integer to a string and returns it
}

func BinToDec(input string) string {
	// This function takes a string input and returns a string representing the decimal equivalent of the binary input
	decimal, err := strconv.ParseInt(input, 2, 64)
	// This line converts the input string to an integer using base 2 (binary) and stores the result in the "decimal" variable
	// The "err" variable will hold any error that occurs during the conversion
	if err != nil {
		// This block checks if an error occurred during the conversion
		fmt.Println("err")
		// If an error occurred, it prints "err" to the console
		os.Exit(0)
		// And then exits the program with a status code of 0
	}
	return strconv.Itoa(int(decimal))
	// If no error occurred, the function converts the decimal integer to a string and returns it
}
