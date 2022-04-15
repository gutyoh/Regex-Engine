package main

/*
[Regex Engine - Stage 2/6: Single character strings](https://hyperskill.org/projects/114/stages/620/implement)
-------------------------------------------------------------------------------
[Control statements](https://hyperskill.org/learn/topic/1728)
[Advanced input](https://hyperskill.org/learn/topic/2027)
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

	case word == "" || !checkChar(regex[0:1], word[0:1]):
		return false
	}
	return checkStr(regex[1:], word[1:])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), "|")

	regex, word := line[0], line[1]

	fmt.Println(checkStr(regex, word))
}
