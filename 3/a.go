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

		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			// I could not figure out how to split a string in half
			items := strings.Split(line, "")
			var compartment1 []string
			var compartment2 []string

			for i, item := range items {
				if i < len(line)/2 {
					compartment1 = append(compartment1, item)
				} else {
					compartment2 = append(compartment2, item)
				}
			}

			var duplicate string

			// gross
			for _, item1 := range compartment1 {
				for _, item2 := range compartment2 {
					if item1 == item2 {
						duplicate = item2
					}
				}
			}

			lower := strings.ToLower(duplicate)
			priority := strings.Index(alphabet, lower) + 1

			if lower == duplicate {
				sum += priority
			} else {
				sum += priority + 26
			}
		}
	}

	fmt.Println(sum)
}
