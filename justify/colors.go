package justify

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// SetColor function converts color names, hexadecimal codes and RGB values to ANSI escape codes for terminal coloring
func SetColor(c string) (string, error) {
	hex := regexp.MustCompile(`^#([a-f0-9]{3}|[a-f0-9]{6})$`)
	rgb := regexp.MustCompile(`^rgb\(\s*(\d+),\s*(\d+),\s*(\d+)\)$`)

	colorCodes := map[string]string{
		"black":     "\033[38;5;16m",
		"red":       "\033[38;5;196m",
		"green":     "\033[38;5;22m",
		"blue":      "\033[38;5;21m",
		"yellow":    "\033[38;5;226m",
		"magenta":   "\u001b[35m",
		"cyan":      "\u001b[36m",
		"white":     "\033[38;5;231m",
		"orange":    "\u001b[38;5;208m",
		"purple":    "\033[38;5;55m",
		"teal":      "\033[38;5;23m",
		"silver":    "\033[38;5;145m",
		"gray":      "\033[38;5;240m",
		"brown":     "\033[38;5;94m",
		"pink":      "\u001b[38;5;207m",
		"olive":     "\u001b[38;5;58m",
		"navy":      "\u001b[38;5;18m",
		"turquoise": "\u001b[38;5;80m",
		"lime":      "\033[38;5;46m",
		"indigo":    "\u001b[38;5;54m",
		"lavender":  "\u001b[38;5;183m",
		"charteuse": "\033[33m\033[34m",
		"salmon":    "\033[38;5;209m",
		"peach":     "\033[33m\033[96m",
		"seafoam":   "\033[32m\033[96m",
		"fuchsia":   "\033[38;5;201m",
		"violet":    "\033[33m\033[95m",
		"aqua":      "\033[38;5;51m",
		"maroon":    "\033[38;5;52m",
	}

	// Check if the given color name exists in the map
	code, ok := colorCodes[c]
	if !ok {
		if c == "" {
			return "", nil
			// If not found, check for RGB or HEX values
		} else if hex.MatchString(c) {
			hexVal := strings.TrimPrefix(c, "#")
			var r, g, b int64

			if len(hexVal) == 3 {
				// 3-character hex code, convert to RGB
				r, _ = strconv.ParseInt(hexVal[0:1]+hexVal[0:1], 16, 64)
				g, _ = strconv.ParseInt(hexVal[1:2]+hexVal[1:2], 16, 64)
				b, _ = strconv.ParseInt(hexVal[2:3]+hexVal[2:3], 16, 64)
			} else {
				// 6-character hex code, convert to RGB
				r, _ = strconv.ParseInt(hexVal[0:2], 16, 64)
				g, _ = strconv.ParseInt(hexVal[2:4], 16, 64)
				b, _ = strconv.ParseInt(hexVal[4:6], 16, 64)
			}
			// Return ANSI escape code for RGB
			return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b), nil

		} else if rgb.MatchString(c) {
			matches := rgb.FindStringSubmatch(c)
			if len(matches) > 0 {
				r, _ := strconv.Atoi(matches[1])
				g, _ := strconv.Atoi(matches[2])
				b, _ := strconv.Atoi(matches[3])
				if (r < 0 || r > 255) || (g < 0 || g > 255) || (b < 0 || b > 255) {
					return "", errors.New("invalid color or color code")
				}
				return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b), nil
			}
		} else {
			return "", errors.New("invalid color or color code")
		}
	}
	return code, nil
}

// Colorize function applies color to the given message using ANSI escape codes
func Colorize(color, message string) string {
	return color + message + "\u001b[0m" // Apply and reset color
}
