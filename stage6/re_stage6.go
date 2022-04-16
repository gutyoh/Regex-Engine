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

func checkChar(regex, char string) string {
	if regex == "" || char == regex || (regex == "." && char != "") {
		return "True"
	}
	return "False"
}

func checkStr(regex, word string) string {
	switch {
	case regex == "" || (regex == "$" && word == ""):
		return "True"

	case word == "":
		return "False"

	// Add the final case to use the '\' character to "escape" the meta-characters '*', '+', '?', '^' and '.':
	case regex[0] == '\\':
		return checkStr(regex[1:], word)

	case len(regex) > 1 && strInSlice(string(regex[1]), []string{"*", "+", "?"}):
		if strInSlice(string(regex[1]), []string{"?", "*"}) && checkStr(regex[2:], word) == "True" {
			return "True"
		} else if checkChar(string(regex[0]), string(word[0])) == "True" &&
			(strInSlice(string(regex[1]), []string{"?", "+"}) && (checkStr(regex[2:], word[1:])) == "True" ||
				(strInSlice(string(regex[1]), []string{"*", "+"}) && (checkStr(regex, word[1:])) == "True")) {
			return "True"
		} else {
			return "False"
		}

	case checkChar(string(regex[0]), string(word[0])) == "False":
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
