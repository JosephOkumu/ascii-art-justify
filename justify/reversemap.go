package justify

import (
	"strings"
)

func ReverseMapCreator(s string, Map map[string]string) (int, int) {
	// Initialize minWidth to largest uint8 value and maxWidth to 0
	minWidth := int(^uint8(0) >> 1)
	maxWidth := 0


	var lines []string
	printableRune := rune(32)

	// Deal with windows style line endings
	if strings.ContainsRune(s, '\r') {
		lines = strings.Split(s, "\r\n")
	} else {
		lines = strings.Split(s, "\n")
	}

	// Process the lines and map each art to it's corresponding ASCII character
	for i := 0; i < len(lines); i++ {
		// Determine the thinnest ASCII art
		if lines[i] != "" && len(lines[i]) < minWidth {
			minWidth = len(lines[i])
		}
		// Determine the thickest ASCII art
		if len(lines[i]) > maxWidth {
			maxWidth = len(lines[i])
		}
		// If the current line is empty and there are lines left to process
		if i+1 < len(lines) && lines[i] == "" {
			// Create an empty string to concatenate ASCII art lines for the current character
			artLines := ""
			// Iterate over 8 lines (assuming ASCII art is 8 lines tall)
			for j := 0; j < 8; j++ {
				// Concatenate each line of ASCII art to the string
				artLines += lines[i+1+j]
			}
			// Map the printable rune to its corresponding ASCII art
			Map[artLines] = string(printableRune)
			// Increment the printable rune
			printableRune++
		}
	}
	return minWidth, maxWidth
}

func CreateUniversalMap() (map[string]string, int, int, error) {
	// Read the 'thinkertoy' banner
	thinkertoy, err := ReadTextFile("./banners/thinkertoy.txt")
	if err != nil {
		return nil, 0, 0, err
	}

	// Read the 'standard' banner
	standard, err := ReadTextFile("./banners/standard.txt")
	if err != nil {
		return nil, 0, 0, err
	}

	// Read the 'shadow' banner
	shadow, err := ReadTextFile("./banners/shadow.txt")
	if err != nil {
		return nil, 0, 0, err
	}

	// Create the universal map to store ASCII art to character mappings
	universalMap := make(map[string]string)

	// Initialize minWidth to largest uint8 value and maxWidth to 0
	minWidth := int(^uint8(0) >> 1)
	maxWidth := 0

	// Map 'thinkertoy' and update minWidth and maxWidth
	min, max := ReverseMapCreator(thinkertoy, universalMap)
	if min < minWidth {
		minWidth = min
	}
	if max > maxWidth {
		maxWidth = max
	}

	// Map 'standard' and update minWidth and maxWidth
	min, max = ReverseMapCreator(standard, universalMap)
	if min < minWidth {
		minWidth = min
	}
	if max > maxWidth {
		maxWidth = max
	}

	// Map 'shadow' and update minWidth and maxWidth
	min, max = ReverseMapCreator(shadow, universalMap)
	if min < minWidth {
		minWidth = min
	}
	if max > maxWidth {
		maxWidth = max
	}

	// Return the universal map and dynamically determined min and max widths
	return universalMap, minWidth, maxWidth, nil
}
