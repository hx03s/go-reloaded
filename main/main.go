package main

// Import required packages
import (
	"fmt"
	"goreloded/internal/converters"
	"os"
	"strings"
)

// Main function
func main() {
	// Check if no arguments are provided
	if len(os.Args) == 1 {
		fmt.Print("I/O required")
		os.Exit(0)
	} else if len(os.Args) == 2 { // Check if only one argument is provided (input file name is missing)
		fmt.Print("File name missing")
		os.Exit(0)
	} else if len(os.Args) > 3 { // Check if more than three arguments are provided
		fmt.Print("Too many arguments")
		os.Exit(0)
	}

	// Get the input file name from the command line arguments
	inputFile := os.Args[1]

	// Read the entire content of the input file
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}

	// Split the file content into words
	str := strings.Split(string(file), " ")

	// Iterate over the words and perform operations based on the provided instructions
	for i := 0; i < len(str); i++ {
		// Convert hexadecimal to decimal
		if str[i] == "(hex)" && i != 0 {
			str[i-1] = converters.HexTodec(str[i-1])
			str[i] = ""
		} else if str[i] == "(bin)" && i != 0 { // Convert binary to decimal
			str[i-1] = converters.BinToDec(str[i-1])
			str[i] = ""
		} else if str[i] == "(up)" && i != 0 { // Convert to uppercase
			str[i-1] = converters.Upper(str[i-1])
			str[i] = ""
		} else if str[i] == "(up," && i != 0 { // Convert to uppercase with a specified number of words
			n := converters.Num(str[i+1])
			if len(str) < n+2 {
				fmt.Println("Error: Number provided is more than the given length")
				os.Exit(0)
			}
			for k := i - n; k < i; k++ {
				str[k] = converters.Upper(str[k])
				str[i] = ""
				str[i+1] = ""
			}
		} else if str[i] == "(low)" && i != 0 { // Convert to lowercase
			str[i-1] = converters.Low(str[i-1])
			str[i] = ""
		} else if str[i] == "(low," && i != 0 { // Convert to lowercase with a specified number of words
			num := converters.Num(str[i+1])
			if len(str) < num+2 {
				fmt.Println("Error: Number provided is more than the given length")
				os.Exit(0)
			}
			for j := i - num; j < i; j++ {
				str[j] = converters.Low(str[j])
				str[i] = ""
				str[i+1] = ""
			}
		} else if str[i] == "(cap)" && i != 0 { // Capitalize the first letter
			str[i-1] = converters.Caps(str[i-1])
			str[i] = ""
		} else if str[i] == "(cap," && i != 0 { // Capitalize the first letter with a specified number of words
			m := converters.Num(str[i+1])
			if len(str) < m+2 {
				fmt.Println("Error: Number provided is more than the given length")
				os.Exit(0)
			}
			for h := i - m; h < i; h++ {
				str[h] = converters.Caps(str[h])
				str[i] = ""
				str[i+1] = ""
			}
		}
	}

	// Join the modified words back into a single string
	stri := strings.Join(str, " ")
	
	// Split the string into fields (removing extra spaces)
	rem := strings.Fields(stri)
	
	// Join the fields back into a single string
	remm := strings.Join(rem, " ")
	
	// Remove punctuation marks
	regex := converters.Punc(remm)
	
	// Remove single quotes
	f := converters.Quot(regex)
	
	// Remove double quotes
	ma := converters.Quott(f)
	
	// Count vowels
	vo := converters.Vowel(ma)
	
	// Print the final string
	fmt.Println(vo)

	// Write the final string to the output file
	os.WriteFile(os.Args[2], []byte(vo), 0644)
}
