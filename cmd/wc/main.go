package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"

	libxo "github.com/kyleseneker/wc/internal/libxo"
	"github.com/urfave/cli/v2"
)

// CustomFlag is a custom flag type that supports only two dashes
type CustomFlag struct {
	cli.Flag
	Name  string
	Value bool
	Usage string
}

// ParseFlag defines how to parse the flag
func (f *CustomFlag) ParseFlag() cli.Flag {
	return &cli.BoolFlag{
		Name:    f.Name,
		Value:   f.Value,
		Aliases: []string{""},
		Usage:   f.Usage,
	}
}

func main() {
	logger := log.New(os.Stderr, "", 0)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "l",
				Usage: "The number of lines in each input file is written to the standard output.",
			},
			&cli.BoolFlag{
				Name:  "w",
				Usage: "The number of words in each input file is written to the standard output.",
			},
			&cli.BoolFlag{
				Name:  "c",
				Usage: "The number of bytes in each input file is written to the standard output. This will cancel out any prior usage of the `-m` option.",
			},
			&cli.BoolFlag{
				Name:  "m",
				Usage: "The number of characters in each input file is written to the standard output. If the current locale does not support multibyte characters, this is equivalent to the `-c` option. This will cancel out any prior usage of the `-c` option.",
			},
			&cli.BoolFlag{
				Name:  "L",
				Usage: "Write the length of the line containing the most bytes (default) or characters (when -m is provided) to standard output.  When more than one file argument is specified, the longest input line of all files is reported as the value of the final \"total\".",
			},
			&cli.BoolFlag{
				Name:  "libxo",
				Usage: "Format output using libxo.",
			},
		},
		Name:  "wc",
		Usage: "word, line, character, and byte count",
		Action: func(ctx *cli.Context) error {
			var content []byte
			var err error

			// Read from standard input if no filename is provided
			if ctx.Args().Len() == 0 {
				content, err = io.ReadAll(os.Stdin)
				if err != nil {
					logger.Fatalf("error reading from standard input: %s", err)
				}
			} else {
				// Read file
				content, err = os.ReadFile(ctx.Args().First())
				if err != nil {
					logger.Fatalf("error reading from file: %s", err)
				}
			}

			filePath := ctx.Args().First()

			if ctx.Bool("libxo") {
				fmt.Println(libxo.ConvertToXML(content))
			}

			if ctx.Bool("l") {
				fmt.Printf("%8d", countLines(content))
			}
			if ctx.Bool("w") {
				fmt.Printf("%8d", countWords(content))
			}
			// If both -m and -c flags are present, determine which one to use based on their order
			if ctx.Bool("m") && ctx.Bool("c") {
				// If -c comes after -m in the command-line arguments, use -c
				if cPos, mPos := flagPosition("c"), flagPosition("m"); cPos > mPos {
					fmt.Printf("%8d", len(content))
				} else { // Otherwise, use -m
					fmt.Printf("%8d", countCharacters(content))
				}
			}
			// If -m flag is present and -c is not, return charater count and return
			if ctx.Bool("m") && !ctx.Bool("c") {
				fmt.Printf("%8d", countCharacters(content))
			}
			// If -c flag is present and -m is not, return byte count and return
			if ctx.Bool("c") && !ctx.Bool("m") {
				fmt.Printf("%8d", len(content))
			}
			if ctx.Bool("L") && !ctx.Bool("m") {
				fmt.Printf("%8d", longestLineLength(content, false))
			}
			if ctx.Bool("L") && ctx.Bool("m") {
				fmt.Printf("%8d", longestLineLength(content, true))
			}
			if ctx.Bool("l") || ctx.Bool("w") || ctx.Bool("c") || ctx.Bool("m") || ctx.Bool("L") {
				fmt.Printf(" %s\n", filePath)
				return nil
			}

			// Default behavior if no flags are provided: print line, word, and byte counts
			fmt.Printf("%8d%8d%8d %s\n", countLines(content), countWords(content), len(content), filePath)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}

// flagPosition returns the position of the flag in the command-line arguments
func flagPosition(flagName string) int {
	for i := 1; i < len(os.Args)-1; i++ {
		arg := os.Args[i]
		// Check if the argument is a flag (starts with "-")
		if strings.HasPrefix(arg, "-") && arg[1:] == flagName {
			return i - 1 // Adjust index to match flag position (starting from 0)
		}
	}
	return -1
}

// countLines counts the number of lines in a byte slice
func countLines(content []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

// countWords counts the number of words in a byte slice
func countWords(content []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	scanner.Split(bufio.ScanWords)
	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}

// countCharacters counts the number of characters in a byte slice.
func countCharacters(content []byte) int {
	charCount := 0
	for len(content) > 0 {
		_, size := utf8.DecodeRune(content)
		charCount++
		content = content[size:]
	}
	return charCount
}

// longestLineLength returns the length of the longest line in a byte slice.
// Returns length in bytes by default
// Returns length in characters if countChars is true
func longestLineLength(content []byte, countChars bool) int {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	maxLength := 0
	for scanner.Scan() {
		lineLength := len(scanner.Text()) + 1 // Include newline character
		if countChars {
			// Count characters instead of bytes
			lineLength = utf8.RuneCountInString(scanner.Text()) + 1 // Include newline character
		}
		if lineLength > maxLength {
			maxLength = lineLength
		}
	}
	return maxLength
}
