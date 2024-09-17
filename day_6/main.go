package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}
	grids := make([][]int, 1000)
	for i := range grid {
		grids[i] = make([]int, 1000)
	}
	instructions := ReadInstructions("puzzle.txt")

	for _, instruction := range instructions {
		Application(grid, instruction)
	}
	count := CountLitLights(grid)
	fmt.Println(count)

	for _, instruction := range instructions {
		ApplicationBrightness(grids, instruction)
	}
	brightness := TotalBrightness(grids)
	fmt.Println(brightness)
}

func ReadInstructions(filename string) []string {
	var instructions []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening the file", err)
		return instructions
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("error reading from file", err)
	}
	return instructions
}
func Application(grid [][]bool, instruction string) {
	parts := strings.Split(instruction, " ")
	var action string
	var x1, y1, x2, y2 int

	if parts[0] == "toggle" {
		action = "toggle"
		fmt.Sscanf(parts[1], "%d,%d", &x1, &y1)
		fmt.Sscanf(parts[3], "%d,%d", &x2, &y2)
	} else {
		action = parts[1]
		fmt.Sscanf(parts[2], "%d,%d", &x1, &y1)
		fmt.Sscanf(parts[4], "%d,%d", &x2, &y2)
	}
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			switch action {
			case "on":
				grid[i][j] = true
			case "off":
				grid[i][j] = false
			case "toggle":
				grid[i][j] = !grid[i][j]
			}
		}
	}
}
func CountLitLights(grid [][]bool) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				count++
			}
		}
	}
	return count
}

// Part2
func ApplicationBrightness(grid [][]int, instruction string) {
	parts := strings.Split(instruction, " ")
	var action string
	var x1, y1, x2, y2 int
	if parts[0] == "toggle" {
		action = "toggle"
		fmt.Sscanf(parts[1], "%d,%d", &x1, &y1)
		fmt.Sscanf(parts[3], "%d,%d", &x2, &y2)
	} else {
		action = parts[1]
		fmt.Sscanf(parts[2], "%d,%d", &x1, &y1)
		fmt.Sscanf(parts[4], "%d,%d", &x2, &y2)
	}
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			switch action {
			case "on":
				grid[i][j]++
			case "off":
				if grid[i][j] > 0 {
					grid[i][j]--
				}
			case "toggle":
				grid[i][j] += 2
			}
		}
	}
}

func TotalBrightness(grid [][]int) int {
	totalBrightness := 0
	for i := range grid {
		for j := range grid[i] {
			totalBrightness += grid[i][j]
		}
	}
	return totalBrightness
}
