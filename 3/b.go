// not really proud of this one, mostly because
// I've never used go before

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := 0
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	content, err := os.ReadFile("3/input.txt")

	// error handling? what's that
	if err == nil {
		lines := strings.Split(string(content), "\n")

		for index := 0; index < len(lines)-3; index += 3 {
			elf1items := strings.Split(lines[index+0], "")
			elf2items := strings.Split(lines[index+1], "")
			elf3items := strings.Split(lines[index+2], "")

			var badge string
			// I cannot think of a way to do this without using 3 for loops
			// I'm not the best programmer
			for _, item1 := range elf1items {
				for _, item2 := range elf2items {
					for _, item3 := range elf3items {
						if item1 == item2 && item2 == item3 && item3 == item1 {
							badge = item1
						}
					}
				}
			}

			lower := strings.ToLower(badge)
			priority := strings.Index(alphabet, lower) + 1

			if lower == badge {
				sum += priority
			} else {
				sum += priority + 26
			}
		}
	}

	fmt.Println(sum)
}