package justify

import (
	"fmt"
	"regexp"
	"strings"
)

// ArtRetriever returns the ASCII art corresponding to the input string using the provided map.
func ArtRetriever(params ArtParams) (string, error) {
	var result strings.Builder

	// Check for empty input or newline character
	if params.InputText == "" {
		result.WriteString("")
		return result.String(), nil
	}

	// Check for newline patterns in the input string
	newline := regexp.MustCompile(`\\n`)
	params.InputText = newline.ReplaceAllString(params.InputText, "\n")

	if onlyNewLines(params.InputText) {
		result.WriteString(params.InputText)
		return result.String(), nil
	}

	lines := strings.Split(params.InputText, "\n")

	for ind := 0; ind < len(lines); ind++ {
		if lines[ind] == "" {
			// Add an empty line if the input line is empty
			result.WriteString("\n")
		} else {
			artString, err := StringBuilder(ArtParams{
				InputText:   lines[ind],
				SubString:   params.SubString,
				Colour:      params.Colour,
				AsciiArtMap: params.AsciiArtMap,
			})
			if err != nil {
				return "", err
			}
			result.WriteString(artString)
		}
	}
	return result.String(), nil
}

// onlyNewLines checks if a string contains only newline runes
func onlyNewLines(s string) bool {
	for _, v := range s {
		if v != '\n' {
			return false
		}
	}
	return true
}

// StringBuilder builds the ASCII art string from input text, colorizing substrings if specified.
func StringBuilder(params ArtParams) (string, error) {
	var result strings.Builder

	for i := 0; i < 8; i++ {
		start := 0

		for start < len(params.InputText) {
			if params.SubString == "" {
				normalString, err := processNormal(params, i)
				if err != nil {
					return "", err
				}
				result.WriteString(normalString)
				break
			} else if strings.HasPrefix(params.InputText[start:], params.SubString) {
				coloredSubstring, err := ColorizeSubstring(params, i)
				if err != nil {
					return "", err
				}
				result.WriteString(coloredSubstring)
				start += len(params.SubString)
			} else {
				result.WriteString(ProcessCharacter(rune(params.InputText[start]), params.AsciiArtMap, i))
				start++
			}
		}

		result.WriteString("\n")
	}

	return result.String(), nil
}

// processNormal processes the input text normally, optionally colorizing each character
func processNormal(params ArtParams, lineIndex int) (string, error) {
	return processText(params, lineIndex, false)
}

// colorizeSubstring colorizes the specified substring.
func ColorizeSubstring(params ArtParams, lineIndex int) (string, error) {
	return processText(params, lineIndex, true)
}

// processCharacter processes a single character, adding its ASCII art lines.
func ProcessCharacter(char rune, asciiArtMap map[rune][]string, lineIndex int) string {
	artLines := asciiArtMap[char]
	return artLines[lineIndex]
}

// processText processes the input text, optionally colorizing it.
func processText(params ArtParams, lineIndex int, isSubstring bool) (string, error) {
	var result strings.Builder
	text := params.InputText
	if isSubstring {
		text = params.SubString
	}

	for _, v := range text {
		artLines, exists := params.AsciiArtMap[v]
		if !exists {
			return "", fmt.Errorf("character '%c' not found in ASCII art map", v)
		}
		result.WriteString(artLines[lineIndex])
	}

	if params.Colour != "" {
		ansiCode, err := SetColor(params.Colour)
		if err != nil {
			return "", err
		}
		return Colorize(ansiCode, result.String()), nil
	}
	return result.String(), nil
}
