package justify

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func TestReadTextFile(t *testing.T) {
	// Test case 1: Valid txt file
	t.Run("Valid txt file", func(t *testing.T) {
		// Create a temporary file
		tmpFile, err := os.Create("test_banner.txt")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tmpFile.Name())
		// Write some content to the temporary file
		content := "Test banner content"
		if _, err := tmpFile.WriteString(content); err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
		// Close the temporary file
		if err := tmpFile.Close(); err != nil {
			t.Fatalf("Failed to close temporary file: %v", err)
		}
		// Test ReadBannerFile function
		bannerContent, err := ReadTextFile(tmpFile.Name())
		if err != nil {
			t.Errorf("ReadBannerFile returned error: %v", err)
		}
		if bannerContent != content {
			t.Errorf("ReadBannerFile did not read expected content. Got: %s, Expected: %s", bannerContent, content)
		}
	})

	// Test case 2: Non-existent file
	t.Run("Non-existent file", func(t *testing.T) {
		_, err := ReadTextFile("non_existent_file.txt")
		if err == nil {
			t.Errorf("Expected error for non-existent file, but got nil")
		}
	})

	// Test case 3: File without .txt extension
	t.Run("Invalid file extension", func(t *testing.T) {
		// Create a temporary non-.txt file
		tmpFile, err := os.Create("test_banner.invalid")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tmpFile.Name())

		_, err = ReadTextFile(tmpFile.Name())
		if err == nil || err.Error() != "error: the file must be a text file with a .txt extension" {
			t.Errorf("Expected error for invalid file extension, but got: %v", err)
		}
	})
}

func TestMapCreator(t *testing.T) {
	// Test input string
	input := `
         
         
         
         
         
         
         
         
`
	// Call MapCreator function
	result, err := MapCreator(input)

	// Test case: Check if there is an error for space character
	if err == nil {
		t.Error("Expected an error for space character")
	}
	expectedError := errors.New("the bannerfile has been tampered with")
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, Got error: %v", expectedError, err)
	}

	// Test case: Check if the result map is empty for space character
	if len(result) != 0 {
		t.Errorf("Expected empty map for space character, Got: %v", result)
	}
}

func TestArtRetriever(t *testing.T) {
	// Define a sample ASCII art map for testing
	artMap := map[rune][]string{
		'H': {
			" _    _  ",
			"| |  | | ",
			"| |__| | ",
			"|  __  | ",
			"| |  | | ",
			"|_|  |_| ",
			"         ",
			"         ",
		},
		'E': {
			" ______  ",
			"|  ____| ",
			"| |__    ",
			"|  __|   ",
			"| |____  ",
			"|______| ",
			"         ",
			"         ",
		},
	}

	// Create a struct for testcases
	testCases := []struct {
		name     string
		input    string
		color    string
		letters  string
		expected string
		hasError bool
	}{
		// Populate the struct with table driven tests
		{
			"Valid input with color",
			"HE",
			"red",
			"",
			"\033[38;5;196m _    _   ______  \u001b[0m\n\033[38;5;196m| |  | | |  ____| \u001b[0m\n\033[38;5;196m| |__| | | |__    \u001b[0m\n\033[38;5;196m|  __  | |  __|   \u001b[0m\n\033[38;5;196m| |  | | | |____  \u001b[0m\n\033[38;5;196m|_|  |_| |______| \u001b[0m\n\033[38;5;196m                  \u001b[0m\n\033[38;5;196m                  \u001b[0m\n",
			false,
		},
		{
			"Valid input with color and letters",
			"HE",
			"red",
			"H",
			"\033[38;5;196m _    _  \u001b[0m ______  \n\033[38;5;196m| |  | | \u001b[0m|  ____| \n\033[38;5;196m| |__| | \u001b[0m| |__    \n\033[38;5;196m|  __  | \u001b[0m|  __|   \n\033[38;5;196m| |  | | \u001b[0m| |____  \n\033[38;5;196m|_|  |_| \u001b[0m|______| \n\033[38;5;196m         \u001b[0m         \n\033[38;5;196m         \u001b[0m         \n",
			false,
		},
		{
			"Empty input",
			"",
			"",
			"",
			"",
			false,
		},
		{
			"Input with newlines",
			"H\nE",
			"red",
			"",
			"\033[38;5;196m _    _  \u001b[0m\n\033[38;5;196m| |  | | \u001b[0m\n\033[38;5;196m| |__| | \u001b[0m\n\033[38;5;196m|  __  | \u001b[0m\n\033[38;5;196m| |  | | \u001b[0m\n\033[38;5;196m|_|  |_| \u001b[0m\n\033[38;5;196m         \u001b[0m\n\033[38;5;196m         \u001b[0m\n\033[38;5;196m ______  \u001b[0m\n\033[38;5;196m|  ____| \u001b[0m\n\033[38;5;196m| |__    \u001b[0m\n\033[38;5;196m|  __|   \u001b[0m\n\033[38;5;196m| |____  \u001b[0m\n\033[38;5;196m|______| \u001b[0m\n\033[38;5;196m         \u001b[0m\n\033[38;5;196m         \u001b[0m\n",
			false,
		},
		{
			"Input with only newlines",
			"\\n\\n",
			"",
			"",
			"\n\n",
			false,
		},
		{
			"Input with invalid character",
			"\r",
			"red",
			"",
			"",
			true,
		},
	}

	// Run the tests for each case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := ArtParams{
				InputText: tc.input,
				SubString: tc.letters,
				Colour: tc.color,
				AsciiArtMap: artMap,
			}
			result, err := ArtRetriever(params)
			if result != tc.expected {
				t.Errorf("Expected:\n%s\nGot:\n%s", tc.expected, result)
			}
			if (err != nil) != tc.hasError {
				t.Errorf("Expected error: %t, Got error: %v", tc.hasError, err)
			}
		})
	}
}

func TestSetColor(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		// Test no color flag provided (empty string)
		{"Flag not provided", "", "", false},

		// Test named color
		{"Red", "red", "\033[38;5;196m", false},

		// Test hexadecimal colors
		{"3-char Hex", "#abc", "\x1b[38;2;170;187;204m", false},
		{"6-char Hex", "#abcdef", "\x1b[38;2;171;205;239m", false},
		{"Invalid Hex", "#gggggg", "", true},

		// Test RGB colors
		{"RGB", "rgb(10,20,30)", "\x1b[38;2;10;20;30m", false},
		{"RGB out of range", "rgb(300,400,500)", "", true},
		{"Invalid RGB", "rgb(0,20)", "", true},

		// Test invalid input
		{"Invalid color", "invalid", "", true},
	}

	// Run the tests for each case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SetColor(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %s, Got: %s", tc.expected, result)
			}
			if (err != nil) != tc.hasError {
				t.Errorf("Expected error: %t, Got error: %v", tc.hasError, err)
			}
		})
	}
}

func TestColorize(t *testing.T) {
	color := "\033[38;5;196m"
	message := "Hello, World!"
	expected := "\033[38;5;196mHello, World!\u001b[0m"

	result := Colorize(color, message)
	if result != expected {
		t.Errorf("Expected: %s, Got: %s", expected, result)
	}
}

func TestProcessReverseFileLines(t *testing.T) {
	// Test case 1: Valid input with equal line lengths
	t.Run("Valid input", func(t *testing.T) {
		// Test case input
		fileContent := "line one$\nline two$\nline six$"

		// Expected output
		expected := []string{
			"line one",
			"line two",
			"line six",
		}

		// Call the function
		result := ProcessReverseFileLines(fileContent)


		// Compare the result with the expected output
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})

}


func TestExtractAsciiArt(t *testing.T) {
	// Test case: ASCII art board input
	board := [][]string{
		{"*", " ", "*", " ", "*"},
		{"*", "*", " ", "*", " "},
		{" ", "*", "*", " ", "*"},
	}

	// Input parameters
	startIndex := 1
	width := 3

	// Expected output: substring of the ASCII art
	expected := " * * *** "

	// Call the function
	result := ExtractAsciiArt(board, startIndex, width)

	// Compare the result with the expected output
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestReverseMapCreator(t *testing.T) {
	// Simulated banner content where each character is represented by 8 lines of ASCII art
	bannerContent := `
      
      
      
      
      
      
      
      

 _  
| | 
| | 
| | 
|_| 
(_) 
    
    

 _ _  
( | ) 
 V V  
      
      
      
      
      
`

	// Initialize an empty map
	asciiMap := make(map[string]string)

	// Call the function with the banner content
	minWidth, maxWidth := ReverseMapCreator(bannerContent, asciiMap)

	// Define expected values for min and max width
	expectedMinWidth := 4
	expectedMaxWidth := 6

	// Check if minWidth and maxWidth are as expected
	if minWidth != expectedMinWidth {
		t.Errorf("Expected minWidth %d, but got %d", expectedMinWidth, minWidth)
	}
	if maxWidth != expectedMaxWidth {
		t.Errorf("Expected maxWidth %d, but got %d", expectedMaxWidth, maxWidth)
	}

	// Check if the map was created correctly (checking for one ASCII art character)
	expectedChar := " "
	expectedArt := "                                                "

	if asciiMap[expectedArt] != expectedChar {
		t.Errorf("Expected map entry for %q, but got %q", expectedArt, asciiMap[expectedArt])
	}
}

func TestAsciiArtReverser(t *testing.T) {
	// Define the input processed lines (each 8 lines correspond to ASCII art representation to be reversed)
	processedLines := []string{
		" _    _          _   _          ",
		"| |  | |        | | | |         ",
		"| |__| |   ___  | | | |   ___   ",
		"|  __  |  / _ \\ | | | |  / _ \\  ",
		"| |  | | |  __/ | | | | | (_) | ",
		"|_|  |_|  \\___| |_| |_|  \\___/  ",
		"                                ",
		"                                ",
	}

	// Define the universal map that will map ASCII art blocks to letters
	universalMap := map[string]string{
		` _    _  | |  | | | |__| | |  __  | | |  | | |_|  |_|                   `: "H",
		`                ___   / _ \ |  __/  \___|               `:                 "e",
		` _  | | | | | | | | |_|         `:                                         "l",
		`                  ___    / _ \  | (_) |  \___/                  `:         "o",
		// More mappings can be added as needed
	}

	// Define the min and max width of the ASCII characters
	min := 4
	max := 9

	// Call the function to test
	result, err := AsciiArtReverser(min, max, processedLines, universalMap)
	if err != nil {
		t.Fatal(err)
	}

	// Define the expected result (the text that should be produced from the ASCII art)
	expected := "Hello\n"

	// Compare the result to the expected value
	if result != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, result)
	}
}

func TestGetConsoleWidth(t *testing.T) {
	// Mock executor that simulates a console width of 80 columns
	mockExecutor := func() ([]byte, error) {
		return []byte("80"), nil
	}

	width, err := getConsoleWidth(mockExecutor)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if width != 80 {
		t.Errorf("expected width 80, got %d", width)
	}
}

func TestGetConsoleWidthError(t *testing.T) {
	// Mock executor that returns an error
	mockExecutor := func() ([]byte, error) {
		return nil, errors.New("failed to execute command")
	}

	_, err := getConsoleWidth(mockExecutor)
	if err == nil {
		t.Fatal("expected an error, but got nil")
	}
}
