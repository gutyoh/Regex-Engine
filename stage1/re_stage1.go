package main

/*
[Regex Engine - Stage 1/6: Single character strings](https://hyperskill.org/projects/114/stages/619/implement)
-------------------------------------------------------------------------------
[Primitive types](https://hyperskill.org/learn/topic/1807)
[Input/Output](https://hyperskill.org/learn/topic/1506)
[Slices](https://hyperskill.org/learn/topic/1672)
[Control statements](https://hyperskill.org/learn/topic/1728)
[Operations with strings](https://hyperskill.org/learn/topic/2023)
[Regexps basics](https://hyperskill.org/learn/topic/577)
*/

import (
	"fmt"
	"strings"
)

func main() {
	var line string
	fmt.Scanln(&line)

	// if line doesn't contain the "|" symbol, print "False" and exit the program:
	if !strings.Contains(line, "|") {
		fmt.Println("False")
		return
	}

	regex := strings.Split(line, "|")[0]
	char := strings.Split(line, "|")[1]

	if regex == "" {
		fmt.Println("True")
	} else if char == "" {
		fmt.Println("False")
	} else if regex == "." || regex == char {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
