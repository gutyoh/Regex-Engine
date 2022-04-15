package main

/*
[Regex Engine - Stage 1/6: Single character strings](https://hyperskill.org/projects/114/stages/619/implement)
-------------------------------------------------------------------------------
[Primitive types](https://hyperskill.org/learn/topic/1807)
[Input/Output](https://hyperskill.org/learn/topic/1506)
[Slices](https://hyperskill.org/learn/topic/1672)
[Loops](https://hyperskill.org/learn/topic/1531)
[Working with slices](https://hyperskill.org/learn/topic/1701)
[Operations with strings](https://hyperskill.org/learn/topic/2023)
[Functions](https://hyperskill.org/learn/topic/1750)
*/

import (
	"fmt"
	"strings"
)

func checkChar(regex, char string) bool {
	return regex == "" || char == regex || (regex == "." && char != "")
}

func main() {
	var line string
	fmt.Scanln(&line)

	regex := strings.Split(line, "|")[0]
	char := strings.Split(line, "|")[1]

	fmt.Println(checkChar(regex, char))
}
