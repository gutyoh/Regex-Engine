package main

/*
[Regex Engine - Stage 3/6: Working with strings of different length](https://hyperskill.org/projects/114/stages/621/implement)
-------------------------------------------------------------------------------
##### 🚫 NO NEW TOPICS REQUIRED 🚫 #####
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkChar(regex, char string) bool {
	return regex == "" || char == regex || (regex == "." && char != "")
}

func checkStr(regex, word string) bool {
	switch {
	case regex == "":
		return true

	case !checkChar(regex[0:1], word[0:1]):
		return false
	}
	return checkStr(regex[1:], word[1:])
}

func compare(regex, word string) bool {
	if regex == "" {
		return true
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
