package justify

import (
    "errors"
    "strings"
)

type ArtParams struct {
    InputText   string            // Text to be converted into ASCII art
    SubString   string            // Substring of input text to colorize
    Colour      string            // Desired color for the whole ASCII art or substring
    AsciiArtMap map[rune][]string // Mapping of characters to their ASCII art representation
}

var differentParams ArtParams

// ArtAligner aligns ASCII art based on the specified position: "left," "right," "center," or "justify".
func ArtAligner(position string, params ArtParams) (string, error) {
    var result strings.Builder
    differentParams.InputText = params.InputText
    differentParams.AsciiArtMap = params.AsciiArtMap

    // Retrieve ASCII representation without color for base alignment
    uncoloredText, err := ArtRetriever(differentParams)
    if err != nil {
        return "", err
    }
	// Split ASCII art into lines and remove the trailing empty line
	lines := strings.Split(uncoloredText, "\n")
	lines = lines[:len(lines)-1]
	paddedLines := []string{}

	// Determine terminal width for alignment calculations
	terminalWidth, err := getConsoleWidth(defaultCommandExecutor)
	if err != nil {
		return "", err
	}

	// Apply specified alignment
	switch position {
	case "left":
		// Left-align the text with color
		ansiCode, _ := SetColor(params.Colour)
		return Colorize(ansiCode, uncoloredText), nil
	case "right":
        // Right-align each line by padding with spaces to the left
        ansiCode, _ := SetColor(params.Colour)
        for _, line := range lines {
            artWidth := len(line)
            if artWidth >= terminalWidth {
                return "", errors.New("terminal size too small, expand terminal")
            }
            padding := terminalWidth - artWidth
            paddedLine := strings.Repeat(" ", padding) + Colorize(ansiCode, line)
            paddedLines = append(paddedLines, paddedLine)
        }
	case "center":
        // Center-align each line by padding with half spaces to the left
        ansiCode, _ := SetColor(params.Colour)
        for _, line := range lines {
            artWidth := len(line)
            if artWidth >= terminalWidth {
                return "", errors.New("terminal size too small, expand terminal")
            }
            padding := (terminalWidth - artWidth) / 2
            paddedLine := strings.Repeat(" ", padding) + Colorize(ansiCode, line)
            paddedLines = append(paddedLines, paddedLine)
        }
	case "justify":
        // Justify text by adjusting spaces between words
        ansiCode, _ := SetColor(params.Colour)
        inputTextLines := strings.Split(params.InputText, "\n")
        for _, line := range inputTextLines {
            differentParams.InputText = line
            spaceCount := strings.Count(line, " ")

            // Retrieve ASCII art for the current line
            lineArtText, _ := ArtRetriever(differentParams)
            linesOfLineArtText := strings.Split(lineArtText, "\n")
            artWidth := len(linesOfLineArtText[0])

            // Calculate padding between words for justification
            padding := 0
            if spaceCount > 0 && artWidth < terminalWidth {
                padding = (terminalWidth - artWidth) / spaceCount
            } else if artWidth >= terminalWidth {
                return "", errors.New("terminal size too small, expand terminal")
            }

            // Generate each line of the ASCII representation with justified spacing
            for i := 0; i < 8; i++ {
                paddedLine := ""
                for start := 0; start < len(differentParams.InputText); start++ {
                    if strings.HasPrefix(line[start:], " ") {
                        // Add padding spaces for justified alignment
                        paddedLine += strings.Repeat(" ", padding) + "      "
                    } else {
                        // Process each character in ASCII art
                        paddedLine += ProcessCharacter(rune(line[start]), params.AsciiArtMap, i)
                    }
                }
                // Apply color if specified to the final padded line
                paddedLine = Colorize(ansiCode, paddedLine)
                paddedLines = append(paddedLines, paddedLine)
            }
        }
    }

    // Combine all padded lines into a single result
    for _, v := range paddedLines {
        result.WriteString(v + "\n")
    }

    return result.String(), nil
}