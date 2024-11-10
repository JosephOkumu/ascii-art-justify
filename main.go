package main

import (
    "fmt"
    "os"
    "regexp"
    "strings"

    "asciiart/justify"
)

// main function reads a banner file, creates a map of ASCII art, validates user input,
// and prints the corresponding ASCII art to the output file.
func main() {
    // Check if color flag is not provided correctly, i.e. provided without equal sign.
    properColorFlag := regexp.MustCompile(`^-color(?:=(.+))?$`)
    properOutputFlag := regexp.MustCompile(`^-output(?:=(.+))?$`)
    properReverseFlag := regexp.MustCompile(`^-reverse(?:=(.+))?$`)
    properAlignFlag := regexp.MustCompile(`^-align(?:=(.+))?$`)
    args := os.Args
    for _, v := range args {
        if properColorFlag.MatchString(v) || v == "--color" {
            fmt.Print("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"\n")
            return
        } else if properOutputFlag.MatchString(v) || v == "--output" {
            fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n")
            return
        } else if properReverseFlag.MatchString(v) || v == "--reverse" {
            fmt.Print("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName.txt>\n")
            return
        } else if properAlignFlag.MatchString(v) || v == "--align" {
            fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --align=right something standard\n")
            return
        }
    }

    // Get the flag values for color, letters to colorize, input text and banner file name. Handle possible errors.
    options, err := justify.ParseOptions()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Get ANSI format string to colorize ASCII-art in the output file.
    // colorCode, err := justify.SetColor(options.ColorFlag)
    // check(err)

    // Read banner file
    if options.BannerFile == "" {
        options.BannerFile = "standard"
    }

    bannerFile, err := justify.ReadTextFile("./banners/" + options.BannerFile + ".txt")
    check(err)
    
    // Create map of ASCII art
    ASCIIArtMap, err := justify.MapCreator(bannerFile)
    check(err)
    
    // Define parameters for Ascii art retrieval
    params := justify.ArtParams{
        InputText:   options.InputText,
        SubString:   options.ColorizeLetters,
        Colour:      options.ColorFlag,
        AsciiArtMap: ASCIIArtMap,
    }

    // Align art representation if specified
    if options.AlignFlag != "" {
        artText, err := justify.ArtAligner(options.AlignFlag, params)
        if err != nil {
            check(err)
        }
        fmt.Print(artText)
        return
    }

    // Checking if reverse flag option was passed
    if options.ReverseFlag != "" {
        // Reading the text file
        reverseFile, err := justify.ReadTextFile(options.ReverseFlag)
        check(err)
        universalMap, min, max, err := justify.CreateUniversalMap()
        check(err)
        // Removing '$' runes from the end of each line, if any
        processedLines := justify.ProcessReverseFileLines(reverseFile)
        result, err := justify.AsciiArtReverser(min, max, processedLines, universalMap)
        check(err)
        if strings.Contains(reverseFile, "$") {
            result = strings.TrimSuffix(result, "\n")
        }
        fmt.Print(result)
        return
    }

    // Get ASCII art corresponding to input text
    asciiArt, err := justify.ArtRetriever(params)
    check(err)

    // Checking whether the specified output file is a text file.
    outputFile := ""
    if options.OutputFlag != "" {
        if strings.HasPrefix(options.OutputFlag, "./banners/") || strings.HasPrefix(options.OutputFlag, "banners/") {
            fmt.Println("error: cannot write to or modify files in the banners directory")
            return
        } else if strings.HasSuffix(options.OutputFlag, ".txt") {
            outputFile = options.OutputFlag
        } else {
            fmt.Println("error: the output file must be a text file <filename.txt>")
            return
        }
    } else {
        fmt.Print(asciiArt)
        return
    }

    // Write the ASCII art to a file
    err = os.WriteFile(outputFile, []byte(asciiArt), 0o644)
    check(err)
}

// Handle errors
func check(e error) {
    if e != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", e)
        os.Exit(1)
    }
}