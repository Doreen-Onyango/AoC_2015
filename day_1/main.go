package main

import (
	"fmt"
)

func main() {
	inst := "(())((())(((((()))))"
	fmt.Println(Apartment(inst))
}

func Apartment(inst string) int {
	floor := 0
	up := '('
	down := ')'
	// inst := "(())((())(((((()))))"

	for _, ch := range inst {
		if ch == up {
			floor += 1
		} else if ch == down {
			floor -= 1
		}
	}
	return floor
}
