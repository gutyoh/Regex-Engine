package main

/*
[Regex Engine - Stage 6/6: Escaping](https://hyperskill.org/projects/114/stages/624/implement)
-------------------------------------------------------------------------------
##### 🚫 NO NEW TOPICS REQUIRED 🚫 #####
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var greedyTokens = [3]string{"?", "*", "+"}

func strInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func boolToStr(b bool) string {
	if !b {
		return "False"
	}
	return "True"
}

func checkChar(regex, char string) bool {
	if regex == "" || char == regex || (regex == "." && char != "") {
		return true
	}
	return false
}

// checkGreedyTokens() checks if the regex isn't empty and contains any of the greedy tokens "?", "*", "+":
func checkGreedyTokens(regex, word string) bool {
	if len(regex) > 1 && strInSlice(string(regex[1]), greedyTokens[:]) {
		if checkQuestionAndStar(regex, word) {
			return true
		} else if checkChar(string(regex[0]), string(word[0])) {
			if checkQuestionAndPlus(regex, word) {
				return true
			} else if checkStarAndPlus(regex, word) {
				return true
			}
		}
	}
	return false
}

// checkQuestionAndStar() checks if the regex contains "?", "*"
func checkQuestionAndStar(regex, word string) bool {
	if strInSlice(string(regex[1]), greedyTokens[0:2]) && checkStr(regex[2:], word) {
		return true
	}
	return false
}

// checkQuestionAndPlus() checks if the regex contains "?", "+"
func checkQuestionAndPlus(regex, word string) bool {
	if strInSlice(string(regex[1]), greedyTokens[0:3]) && (checkStr(regex[2:], word[1:])) {
		return true
	}
	return checkStarAndPlus(regex, word)
}

// checkStarAndPlus() checks if the regex contains "*", "+"
func checkStarAndPlus(regex, word string) bool {
	if strInSlice(string(regex[1]), greedyTokens[1:3]) && (checkStr(regex, word[1:])) {
		return true
	}
	return false
}

func endOfComparison(regex, word string) bool {
	if !checkChar(string(regex[0]), string(word[0])) {
		return false
	}

	if checkStr(regex[1:], word[1:]) {
		return true
	}
	return false
}

func fullScanComparison(regex, word string) bool {
	for w := range word {
		if checkStr(regex, word[w:]) {
			return true
		}
	}
	return false
}

func compare(regex, word string) bool {
	if regex == "" {
		return true
	}

	if regex[0] == '^' {
		if checkStr(regex[1:], word) {
			return true
		}
	}
	return fullScanComparison(regex, word)
}

// checkStr() checks for most test cases of the regex
// First we check if the regex is empty or if it contains "$", next we check if the word is empty,
// Then we check if the regex contains the "escape" character -- "\"
// Then we check if the regex contains the greedy tokens -- "?", "*", "+" with checkGreedyTokens()
// Finally we check each character in the word and compare it to the regex with the endOfComparison() function
func checkStr(regex, word string) bool {
	if regex == "" || (regex == "$" && word == "") {
		return true
	}

	if word == "" {
		return false
	}

	// Add the final case to use the '\' character to "escape" the meta-characters '*', '+', '?', '^' and '.':
	if strings.Contains(regex, "\\") {
		return checkStr(regex[1:], word)
	}

	if checkGreedyTokens(regex, word) {
		return true
	}

	if endOfComparison(regex, word) {
		return true
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	// if scanner.Text() doesn't contain the "|" symbol, exit the program:
	if !strings.Contains(scanner.Text(), "|") {
		fmt.Println("False")
		return
	}

	// if scanner.Text() contains the | symbol, split 'line' by the | symbol and continue with the program:
	line := strings.Split(scanner.Text(), "|")
	regex, word := line[0], line[1]

	fmt.Println(boolToStr(compare(regex, word)))
}
