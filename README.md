This project is a command-line tool written in Go that performs various text modifications on an input file based on specific instructions. The modified text is then written to an output file.

Features
Convert hexadecimal numbers preceded by (hex) to their decimal equivalent.
Convert binary numbers preceded by (bin) to their decimal equivalent.
Convert words preceded by (up) to uppercase.
Convert words preceded by (low) to lowercase.
Capitalize the first letter of words preceded by (cap).
Apply uppercase, lowercase, or capitalization to a specified number of words if a number follows the instruction (e.g., (up,3), (low,2), (cap,4)).
Ensure proper spacing around punctuation marks (., ,, !, ?, :, ;), handling special cases like consecutive punctuation marks.
Remove spaces around single-quoted words and place the quotes adjacent to the words.
Replace the word "a" with "an" if the next word starts with a vowel or the letter "h".
Installation
Clone the repository:
git clone https://github.com/your-username/text-editing-tool.git



Navigate to the project directory:
cd text-editing-tool



Usage
To run the program, use the following command:

go run . [input_file] [output_file]



Replace [input_file] with the path to the input text file and [output_file] with the desired path for the output file.

For example:

go run . sample.txt result.txt



This command will read the text from sample.txt, apply the specified modifications, and write the modified text to result.txt.

Testing
The project includes a set of test cases to ensure the correctness of the implemented text modifications. To run the tests, use the following command:

go test ./...



This command will execute all the test files in the project and its subpackages.

You can also create your own test cases by adding new files with the naming convention *_test.go in the appropriate package directories.

Example Test Cases
Here are some example test cases to verify the functionality of the text modifications:

Hexadecimal and Binary Conversion:

Input: "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
Expected Output: "Simply add 66 and 2 and you will see the result is 68."
Uppercase, Lowercase, and Capitalization:

Input: "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."
Expected Output: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."
Punctuation and Quotes:

Input: "Punctuation tests are ... kinda boring ,don't you think !?"
Expected Output: "Punctuation tests are... kinda boring, don't you think!?"
Input: "As Elton John said: ' I am the most well-known homosexual in the world '"
Expected Output: "As Elton John said: 'I am the most well-known homosexual in the world'"
Vowel Replacement:

Input: "There is no greater agony than bearing a untold story inside you."
Expected Output: "There is no greater agony than bearing an untold story inside you."
Feel free to add more test cases to cover additional scenarios and edge cases.

Contributing
Contributions to this project are welcome. If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

License
This project is licensed under the MIT License.
