package main

/*
[Regex Engine - Stage 2/6: Single character strings](https://hyperskill.org/projects/114/stages/620/implement)
-------------------------------------------------------------------------------
[Loops](https://hyperskill.org/learn/topic/1531)
[Working with slices](https://hyperskill.org/learn/topic/1701)
[Functions](https://hyperskill.org/learn/topic/1750)
[Advanced input](https://hyperskill.org/learn/topic/2027)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkChar(regex, char string) string {
	if regex == "" {
		return "True"
	} else if char == regex {
		return "True"
	} else if regex == "." && char != "" {
		return "True"
	} else {
		return "False"
	}
}

// New function checkStr, will help usp check for checking strings apart from chars
func checkStr(regex, word string) string {
	if regex == "" {
		return "True"
	} else if word == "" || checkChar(regex[0:1], word[0:1]) == "False" {
		return "False"
	} else if checkStr(regex[1:], word[1:]) == "True" {
		return "True"
	} else {
		return "False"
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), "|")

	regex, word := line[0], line[1]

	fmt.Println(checkStr(regex, word))
}
