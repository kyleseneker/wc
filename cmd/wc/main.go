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

	"github.com/urfave/cli/v2"
)

// Count represents various counts (lines, words, bytes, characters) and the longest line length.
type Count struct {
	Lines       int
	Words       int
	Bytes       int
	Characters  int
	LongestLine int
}

func main() {
	logger := log.New(os.Stderr, "", 0)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "l",
				Usage: `The number of lines in each input file is written to the standard output.`,
			},
			&cli.BoolFlag{
				Name:  "w",
				Usage: `The number of words in each input file is written to the standard output.`,
			},
			&cli.BoolFlag{
				Name: "c",
				Usage: `The number of bytes in each input file is written to the standard output.
	This will cancel out any prior usage of the '-m' option.`,
			},
			&cli.BoolFlag{
				Name: "m",
				Usage: `The number of characters in each input file is written to the standard output.
	If the current locale does not support multibyte characters, this is
	equivalent to the '-c' option. This will cancel out any prior usage of
	the '-c' option.`,
			},
			&cli.BoolFlag{
				Name: "L",
				Usage: `Write the length of the line containing the most bytes (default) or characters
	(when -m is provided) to standard output. When more than one file argument is
	specified, the longest input line of all files is reported as the value of the
	final 'total'.`,
			},
		},
		Name:  "wc",
		Usage: "word, line, character, and byte count",
		Action: func(ctx *cli.Context) error {
			var totalCount Count
			var filesCount []Count
			countChars := ctx.Bool("m")

			// Read from standard input if no filename is provided
			if ctx.Args().Len() == 0 {
				content, err := io.ReadAll(os.Stdin)
				if err != nil {
					logger.Fatalf("error reading from standard input: %s", err)
				}

				fileCount := getCounts(content, countChars)
				printCounts(fileCount, "", ctx)
			} else {
				for _, filePath := range ctx.Args().Slice() {
					content, err := os.ReadFile(filePath)
					if err != nil {
						logger.Fatalf("error reading from file %s: %s", filePath, err)
					}

					// Collect results from all files
					fileCount := getCounts(content, countChars)
					filesCount = append(filesCount, fileCount)
					totalCount.Lines += fileCount.Lines
					totalCount.Words += fileCount.Words
					totalCount.Bytes += fileCount.Bytes
					totalCount.Characters += fileCount.Characters
					if fileCount.LongestLine > totalCount.LongestLine {
						totalCount.LongestLine = fileCount.LongestLine
					}

					printCounts(fileCount, filePath, ctx)
				}
			}

			if len(filesCount) > 1 { // Only return a total line if more than one file is provided
				printCounts(totalCount, "total", ctx)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatal(err)
	}
}

func getCounts(content []byte, countChars bool) Count {
	return Count{
		Lines:       countLines(content),
		Words:       countWords(content),
		Bytes:       len(content),
		Characters:  countCharacters(content),
		LongestLine: longestLineLength(content, countChars),
	}
}

func printCounts(count Count, filePath string, ctx *cli.Context) {
	if ctx.Bool("l") {
		fmt.Printf("%8d", count.Lines)
	}
	if ctx.Bool("w") {
		fmt.Printf("%8d", count.Words)
	}
	// If both -m and -c flags are present, determine which one to use based on their order
	if ctx.Bool("m") && ctx.Bool("c") {
		// If -c comes after -m in the command-line arguments, use -c
		if cPos, mPos := flagPosition("c"), flagPosition("m"); cPos > mPos {
			fmt.Printf("%8d", count.Bytes)
		} else { // Otherwise, use -m
			fmt.Printf("%8d", count.Characters)
		}
	}
	// If -m flag is present and -c is not, return character count and return
	if ctx.Bool("m") && !ctx.Bool("c") {
		fmt.Printf("%8d", count.Characters)
	}
	// If -c flag is present and -m is not, return byte count and return
	if ctx.Bool("c") && !ctx.Bool("m") {
		fmt.Printf("%8d", count.Bytes)
	}
	if ctx.Bool("L") && !ctx.Bool("m") {
		fmt.Printf("%8d", count.LongestLine)
	}
	if ctx.Bool("L") && ctx.Bool("m") {
		fmt.Printf("%8d", count.LongestLine)
	}

	if ctx.Bool("l") || ctx.Bool("w") || ctx.Bool("c") || ctx.Bool("m") || ctx.Bool("L") {
		fmt.Printf(" %s\n", filePath)
	} else {
		// Default behavior if no flags are provided: print line, word, and byte counts
		fmt.Printf("%8d%8d%8d %s\n", count.Lines, count.Words, count.Bytes, filePath)
	}
}

// flagPosition returns the position of the flag in the command-line arguments.
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

// countLines counts the number of lines in a byte slice.
func countLines(content []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

// countWords counts the number of words in a byte slice.
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
// Returns length in bytes by default.
// Returns length in characters if countChars is true.
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
