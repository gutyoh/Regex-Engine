package main

/*
[Regex Engine - Stage 3/6: Working with strings of different length](https://hyperskill.org/projects/114/stages/621/implement)
-------------------------------------------------------------------------------
[Function decomposition](https://hyperskill.org/learn/topic/1893)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Refactor the 'if' statement of checkChar to be more complex but shorter:
func checkChar(regex, char string) string {
	if regex == "" || char == regex || (regex == "." && char != "") {
		return "True"
	}
	return "False"
}

// Refactor the 'if' statement of checkString to use a 'switch' statement instead:
func checkStr(regex, word string) string {
	switch {
	case regex == "":
		return "True"

	case checkChar(regex[0:1], word[0:1]) == "False":
		return "False"

	case checkStr(regex[1:], word[1:]) == "True":
		return "True"

	default:
		return "False"
	}
}

// Here we apply Function decomposition, because we make compare an independent function
// and not part of the checkStr function or the main function.

// The compare function uses a for loop to "compare" each character of 'regex' with 'word':
func compare(regex, word string) string {
	if regex == "" {
		return "True"
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
