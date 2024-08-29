package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("error reading from file")
	}
	println(Position(string(data)))
}

func Position(s string) int {
	count := 1
	contain := make(map[[2]int]bool)
	ps := [2]int{0, 0}
	contain[ps] = true
	for _, char := range s {
		if char == '>' {
			ps[0] += 1
		}
		if char == '^' {
			ps[1] += 1
		}
		if char == '<' {
			ps[0] -= 1
		}
		if char == 'v' {
			ps[1] -= 1
		}

		if !contain[ps] {
			count++
			contain[ps] = true
		}
	}
	return count
}
