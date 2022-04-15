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

func checkChar(regex, char string) bool {
	return regex == "" || char == regex || (regex == "." && char != "")
}

func checkStr(regex, word string) bool {
	switch {
	case regex == "" || (regex == "$" && word == ""):
		return true

	case word == "":
		return false

	case len(regex) > 1 && strInSlice(string(regex[1]), []string{"*", "+", "?"}):
		if strInSlice(string(regex[1]), []string{"?", "*"}) && checkStr(regex[2:], word) {
			return true
		}

		if checkChar(string(regex[0]), string(word[0])) &&
			(strInSlice(string(regex[1]), []string{"?", "+"}) && checkStr(regex[2:], word[1:])) ||
			(strInSlice(string(regex[1]), []string{"*", "+"}) && checkStr(regex, word[1:])) {
			return true
		}

	case !checkChar(string(regex[0]), string(word[0])):
		return false
	}

	return checkStr(regex[1:], word[1:])
}

func compare(regex, word string) bool {
	switch {
	case regex == "":
		return true

	case regex[0] == '^':
		return checkStr(regex[1:], word)
	}

	for w, _ := range word {
		if checkStr(regex, word[w:]) {
			return true
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), "|")

	regex, word := line[0], line[1]

	fmt.Println(compare(regex, word))
}
