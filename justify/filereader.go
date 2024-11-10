package justify

import (
	"fmt"
	"os"
	"strings"
)

// ReadBannerFile reads the content of a banner file specified by the filepath argument and returns it as a string.
func ReadTextFile(filepath string) (string, error) {
	if !strings.HasSuffix(filepath, ".txt") {
		return "", fmt.Errorf("the file must be a text file with a .txt extension")
	}

	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(fileContent), nil
}
