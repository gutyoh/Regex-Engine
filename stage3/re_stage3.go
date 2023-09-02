package main

/*
[Regex Engine - Stage 3/6: Working with strings of different length](https://hyperskill.org/projects/114/stages/621/implement)
-------------------------------------------------------------------------------
[Loops](https://hyperskill.org/learn/topic/1531)
[Working with slices](https://hyperskill.org/learn/topic/1701)
[Type conversion and overflow](https://hyperskill.org/learn/topic/2040)
[Function decomposition](https://hyperskill.org/learn/topic/1893)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func boolToStr(b bool) string {
	if !b {
		return "False"
	}
	return "True"
}

// Refactor the 'if' statement of checkChar to be more complex but shorter:
func checkChar(regex, char string) bool {
	if regex == "" || char == regex || (regex == "." && char != "") {
		return true
	}
	return false
}

// Here we apply Function decomposition; we'll start using several small functions
// To help us separate in parts the logic of the checkStr() function.
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

	return fullScanComparison(regex, word)
}

// Refactor checkStr() to make it shorter and call endOfComparison() within it:
func checkStr(regex, word string) bool {
	if regex == "" {
		return true
	}

	if word == "" {
		return false
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
