package main

import "fmt"

func main() {
	inst := "(())((())(((((()))))"
	fmt.Println(FloorDirection(inst))
}

func FloorDirection(inst string) int {
	floor := 0
	up := '('
	down := ')'

	for _, c := range inst {
		if c == up {
			floor += 1
		} else if c == down {
			floor -= 1
		}
	}
	return floor
}
