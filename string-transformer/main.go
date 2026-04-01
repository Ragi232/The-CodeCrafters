package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	for {
		var options string
		fmt.Println("Choose an option to continue\n 1.ToUpper\n 2.ToLower\n 3.Capitalize\n 4.TitleCase\n 5.SnakeCase\n 6.ReverseWord\n 7.Exit")
		fmt.Scanln(&options)

		switch options {
		case "1":
			{
				var words string
				fmt.Println("INPUT THE WORDS TO BE TRANSFORMED!!")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				words = scanner.Text()
				fmt.Println("Output:", strings.ToUpper(words))
				continue
			}

		case "2":
			{
				var words string
				fmt.Println("INPUT THE WORDS TO BE TRANSFORMED!!")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				words = scanner.Text()
				fmt.Println("output:", strings.ToLower(words))
				continue

			}

		case "3":
			{
				var words string
				fmt.Println("INPUT THE WORDS TO BE TRANSFORMED!!")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				words = scanner.Text()
				words = strings.ToLower(words)
				fmt.Println("Output:", strings.Title(words))
				continue
			}

		case "4":
			{
				var words string
				fmt.Println("INPUT THE WORDS FOR CONVERSION")

				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				words = scanner.Text()

				smallWords := map[string]bool{
					"a": true, "an": true, "the": true,
					"and": true, "but": true, "or": true,
					"for": true, "nor": true,
					"on": true, "at": true, "to": true, "by": true,
					"in": true, "of": true, "up": true,
					"as": true, "is": true, "it": true,
				}
				wordList := strings.Fields(strings.ToLower(words))

				for i, word := range wordList {
					if i == 0 || !smallWords[word] {
						wordList[i] = strings.ToUpper(word[:1]) + word[1:]
					}
				}

				result := strings.Join(wordList, " ")

				fmt.Println("Output:", result)
			}

		case "5":
			{
				var words string
				fmt.Println("INPUT THE WORDS TO BE TRANSFORMED")

				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				words = scanner.Text()

				var result []rune
				input := strings.ToLower(words)

				for _, r := range input {
					if unicode.IsLetter(r) || unicode.IsDigit(r) {
						result = append(result, r)
					} else if unicode.IsSpace(r) {
						result = append(result, '_')
					}
				}

				var clean strings.Builder
				prevUnderscore := false

				for _, r := range result {
					if r == '_' {
						if !prevUnderscore {
							clean.WriteRune(r)
						}
						prevUnderscore = true
					} else {
						clean.WriteRune(r)
						prevUnderscore = false
					}
				}

				output := strings.Trim(clean.String(), "_")

				fmt.Println("Output:", output)
			}

		case "6":
			{
				var words string
				fmt.Println("INPUT THE WORDS TO BE TRANSFORMED")

				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				words = scanner.Text()

				var result []rune
				var word []rune

				for _, r := range words {
					if unicode.IsSpace(r) {
						for i := len(word) - 1; i >= 0; i-- {
							result = append(result, word[i])
						}
						word = []rune{}

						result = append(result, r)
					} else {
						word = append(word, r)
					}
				}

				for i := len(word) - 1; i >= 0; i-- {
					result = append(result, word[i])
				}

				output := string(result)

				fmt.Println("Output:", output)
			}

		case "7": {
				fmt.Println("Shutting down String Transformer. Goodbye.")
				return
			}
			
		}

		}

	}


