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

		line = replaceTODO(line)

		line = replaceClassified(line)

		if shouldRemove(line) {
			linesRemoved++
			continue
		}

		processed = append(processed, line)
	}

	for i := range processed {
		processed[i] = fmt.Sprintf("%03d. %s", i+1, processed[i])
	}

	out, err := os.Create(outFile)
	if err != nil {
		fmt.Println("✗ Cannot write to output file")
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	writer.WriteString("SENTINEL FIELD REPORT — PROCESSED\n")
	writer.WriteString("---------------------------------\n")

	for _, line := range processed {
		writer.WriteString(line + "\n")
	}

	writer.Flush()

	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
	}

	fmt.Println("\nProcessing Complete:")
	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", len(processed))
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : trim, todo, classified, remove, numbering")
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func replaceTODO(s string) string {
	return strings.ReplaceAll(s, "TODO:", "✦ ACTION:")
}

func replaceClassified(s string) string {
	return strings.ReplaceAll(s, "CLASSIFIED:", "[REDACTED]:")
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
