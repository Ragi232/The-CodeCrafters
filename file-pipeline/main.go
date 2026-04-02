package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inFile := os.Args[1]
	outFile := os.Args[2]

	if inFile == outFile {
		fmt.Println("✗ Input and output cannot be the same file.")
		return
	}

	file, err := os.Open(inFile)
	if err != nil {
		fmt.Println("✗ File not found:", inFile)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var processed []string
	linesRead := 0
	linesRemoved := 0

	for scanner.Scan() {
		line := scanner.Text()
		linesRead++

		line = trim(line)

		if shouldRemove(line) {
			linesRemoved++
			continue
		}

		line = replaceTODO(line)
		line = replaceClassified(line)
		line = reverseIfNeeded(line)

		processed = append(processed, line)
	}

	out, err := os.Create(outFile)
	if err != nil {
		fmt.Println("✗ Cannot write to output file")
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	writer.WriteString("SENTINEL FIELD REPORT — PROCESSED\n")
	for i, line := range processed {
		writer.WriteString(fmt.Sprintf("%03d. %s\n", i+1, line))
	}

	writer.Flush()

	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}

	fmt.Println("\nProcessing Complete:")
	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", len(processed))
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : Trim whitespace, Remove blank/dash lines, Replace TODO, Replace CLASSIFIED, Reverse REVERSE lines")
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

// Only replace uppercase TODO:, preserve lowercase todo:
func replaceTODO(s string) string {
	if strings.HasPrefix(s, "TODO:") {
		return strings.Replace(s, "TODO:", "✦ ACTION:", 1)
	}
	return s
}

// Only replace uppercase CLASSIFIED:, preserve lowercase classified:
func replaceClassified(s string) string {
	if strings.HasPrefix(s, "CLASSIFIED:") {
		return strings.Replace(s, "CLASSIFIED:", "[REDACTED]:", 1)
	}
	return s
}

func shouldRemove(s string) bool {
	if s == "" {
		return true
	}
	if strings.Trim(s, "-") == "" {
		return true
	}
	return false
}

func reverseIfNeeded(s string) string {
	if !strings.Contains(s, "REVERSE") {
		return s
	}

	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}
