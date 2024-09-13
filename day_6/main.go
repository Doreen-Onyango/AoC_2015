package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var grid [1000][1000]bool

func main() {
	ReadInstructions("puzzle.txt")
	fmt.Println(CountLitLights())
}

func ReadInstructions(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var command string
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%s %d,%d through %d,%d", &command, &x1, &y1, &x2, &y2)

		switch command {
		case "turn":
			if strings.HasPrefix(line, "turn on") {
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						grid[i][j] = true
					}
				}
			} else if strings.HasPrefix(line, "turn off") {
				for i := x1; i <= x2; i++ {
					for j := y1; j <= y2; j++ {
						grid[i][j] = false
					}
				}
			}
		case "toggle":
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					grid[i][j] = !grid[i][j]
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func CountLitLights() int {
	count := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				count++
			}
		}
	}
	return count
}
