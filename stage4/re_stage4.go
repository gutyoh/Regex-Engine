package main

/*
[Regex Engine - Stage 4/6: Implementing the operators ^ and $](https://hyperskill.org/projects/114/stages/622/implement)
-------------------------------------------------------------------------------
##### 🚫 NO NEW TOPICS REQUIRED 🚫 #####
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkChar(regex, char string) string {
	if regex == "" || char == regex || (regex == "." && char != "") {
		return "True"
	}
	return "False"
}

func checkStr(regex, word string) string {
	switch {
	// Add a check for one of the new special cases; if the regex ends with the the '$' character:
	case regex == "" || (regex == "$" && word == ""):
		return "True"

	case checkChar(regex[0:1], word[0:1]) == "False":
		return "False"

	case checkStr(regex[1:], word[1:]) == "True":
		return "True"

	default:
		return "False"
	}
}

func compare(regex, word string) string {
	switch {
	case regex == "":
		return "True"

	// Add another check for the second special case; if the regex begins with the '^' character:
	case regex[0] == '^':
		if checkStr(regex[1:], word) == "True" {
			return "True"
		}
	}

	for w, _ := range word {
		if checkStr(regex, word[w:]) == "True" {
			return "True"
		}
	}
	return "False"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), "|")

	regex, word := line[0], line[1]

	fmt.Println(compare(regex, word))
}
