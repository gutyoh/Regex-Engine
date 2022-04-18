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

func checkGreedyTokens(regex, word string) bool {
	if len(regex) > 1 && strInSlice(string(regex[1]), []string{"*", "+", "?"}) {
		if strInSlice(string(regex[1]), []string{"?", "*"}) && checkStr(regex[2:], word) {
			return true
		} else if checkChar(string(regex[0]), string(word[0])) {
			if strInSlice(string(regex[1]), []string{"?", "+"}) && (checkStr(regex[2:], word[1:])) {
				return true
			} else if strInSlice(string(regex[1]), []string{"*", "+"}) && (checkStr(regex, word[1:])) {
				return true
			}
		}
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

func checkStr(regex, word string) bool {
	if regex == "" || (regex == "$" && word == "") {
		return true
	}

	if word == "" {
		return false
	}

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

func compare(regex, word string) string {
	if regex == "" {
		return boolToStr(true)
	}

	if regex[0] == '^' {
		if checkStr(regex[1:], word) {
			return boolToStr(true)
		}
	}
	return fullScanComparison(regex, word)
}

func fullScanComparison(regex, word string) string {
	for w := range word {
		if checkStr(regex, word[w:]) {
			return boolToStr(true)
		}
	}
	return boolToStr(false)
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
	fmt.Println(compare(regex, word))
}
