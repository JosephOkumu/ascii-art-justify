package justify

import (
    "errors"
    "flag"
    "strings"
)

type Options struct {
    ColorFlag       string
    OutputFlag      string
    ColorizeLetters string
    InputText       string
    BannerFile      string
    ReverseFlag     string
    AlignFlag       string
}

// ParseOptions function parses command-line arguments to extract color options, text to be colored, file to output the data, and banner file name.
func ParseOptions() (Options, error) {
    // Define flag for color option
    var options Options
    flag.StringVar(&options.ColorFlag, "color", "", "Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"\n")
    flag.StringVar(&options.OutputFlag, "output", "", "Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> \"something\" standard")
    flag.StringVar(&options.ReverseFlag, "reverse", "", "Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName.txt>")
    flag.StringVar(&options.AlignFlag, "align", "", "Usage: --align=<position>")
    flag.Parse()

    // Determine the number of arguments and parse accordingly
    switch len(flag.Args()) {
    case 0:
        return options, nil
    case 1: // One argument: input text
        options.InputText = flag.Arg(0)
    case 2:
        // Two arguments: colorize letters and input text or bannerfile
        banner := strings.ToLower(flag.Arg(1))
        if banner == "thinkertoy" || banner == "standard" || banner == "shadow" {
            options.InputText = flag.Arg(0)
            options.BannerFile = banner
            // Check if the specified banner file includes .txt extension and trimming the suffix .txt if present.
        } else if banner == "thinkertoy.txt" || banner == "standard.txt" || banner == "shadow.txt" {
            options.BannerFile = strings.TrimSuffix(banner, ".txt")
            options.InputText = flag.Arg(0)
        } else {
            if options.ColorFlag != "" {
                options.ColorizeLetters = flag.Arg(0)
                options.InputText = flag.Arg(1)
            } else if options.ColorFlag == "" {
                options.BannerFile = strings.TrimSuffix(banner, ".txt")
                options.InputText = flag.Arg(0)
                return options, nil
            } else {
                return Options{}, errors.New("Usage: go run . [STRING] [BANNER]\n\nex: go run . something standard")
            }
        }
    case 3:
        // Three arguments: colorize letters, input text and bannerfile
        if options.ColorFlag != "" {
            options.ColorizeLetters = flag.Arg(0)
            options.InputText = flag.Arg(1)
            options.BannerFile = strings.TrimSuffix(strings.ToLower(flag.Arg(2)), ".txt")
        } else {
            return Options{}, errors.New("Usage: go run . [STRING] [BANNER]\n\nex: go run . something standard")
        }
    default:
        // Invalid number of arguments
        return Options{}, errors.New("Usage: go run . [STRING] [BANNER]\n\nex: go run . something standard")
    }
    // Convert color flag and banner file name to lowercase for consistency
    options.ColorFlag = strings.ToLower(options.ColorFlag)

    return options, nil
}