package justify

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// CommandExecutor defines a function type that executes a command and returns its output.
type CommandExecutor func() ([]byte, error)

// defaultCommandExecutor is the standard command executor that calls "tput cols".
var defaultCommandExecutor CommandExecutor = func() ([]byte, error) {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	return cmd.Output()
}

func getConsoleWidth(executor CommandExecutor) (int, error) {
	widthBytes, err := executor()
	if err != nil {
		return 0, err
	}

	width, err := strconv.Atoi(strings.TrimSpace(string(widthBytes)))
	if err != nil {
		return 0, err
	}
	return width, nil
}
