package main

/*
[Regex Engine - Stage 2/6: Single character strings](https://hyperskill.org/projects/114/stages/620/implement)
-------------------------------------------------------------------------------
[Functions](https://hyperskill.org/learn/topic/1750)
[Advanced input](https://hyperskill.org/learn/topic/2027)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// New function boolToStr will help us convert the output of our program to a string "True" or "False"
func boolToStr(b bool) string {
	if !b {
		return "False"
	}
	return "True"
}

func checkChar(regex, char string) bool {
	if regex == "" {
		return true
	} else if char == regex {
		return true
	} else if regex == "." && char != "" {
		return true
	} else {
		return false
	}
}

// New function checkStr, will help us check for checking strings apart from chars
func checkStr(regex, word string) bool {
	if regex == "" {
		return true
	} else if word == "" || !checkChar(regex[0:1], word[0:1]) {
		return false
	} else if checkStr(regex[1:], word[1:]) {
		return true
	} else {
		return false
	}
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

	fmt.Println(boolToStr(checkStr(regex, word)))
}
