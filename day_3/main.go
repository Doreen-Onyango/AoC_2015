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
	fmt.Println(Position(string(data)))
	fmt.Println(RoboSanta(string(data)))
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

func RoboSanta(str string) int {
	count := 0
	contains := make(map[[2]int]bool)
	sans, robs := [2]int{0, 0}, [2]int{0, 0}
	contains[sans] = true
	count++

	for i, char := range str {
		var pos *[2]int
		if i%2 == 0 {
			pos = &sans
		} else {
			pos = &robs
		}
		switch char {
		case '>':
			pos[0]++
		case '^':
			pos[1]++
		case '<':
			pos[0]--
		case 'v':
			pos[1]--
		}

		if !contains[*pos] {
			contains[*pos] = true
			count++
		}

	}
	return count
}
