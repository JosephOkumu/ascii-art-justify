package justify

import (
	"strings"
)

// ProcessReverseFileLines split input string into lines and checks for and removes any custom delimiter "$"
// from the end of each line, and returns the processed lines as a slice of strings
func ProcessReverseFileLines(fileContent string) []string {
	lines := strings.Split(fileContent, "\n")

	processedLines := make([]string, 0, len(lines))

	
	for _, line := range lines {
		line = strings.TrimSuffix(line, "$")

		processedLines = append(processedLines, line)
	}

	return processedLines
}

// ExtractAsciiArt extracts a specific section of ASCII art from the 2D board starting
// from a given startIndex and spans across a specified width.
// Returns the concatenated string representing the extracted ASCII art.
func ExtractAsciiArt(board [][]string, startIndex, width int) string {
	// Get the number of rows and columns in the board
	rows := len(board)
	cols := len(board[0])

	result := ""

	// Iterate through each row of the board
	for i := 0; i < rows; i++ {
		// For each row, iterate through the characters in the specified width from the starting index of column
		for j := startIndex; j < startIndex+width; j++ {
			// Ensure the index doesn't exceed the number of columns
			if j < cols {
				// Append the character to the result string
				result += board[i][j]
			}
		}
	}

	// Return the concatenated result as the extracted ASCII art to be checked in the Universal Map
	return result
}
