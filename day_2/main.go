package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("puzzle.txt")
	if err != nil {
		fmt.Println("error reading from file")
	}
	sum := 0
	newdata := strings.Split(string(data), "\n")
	for _, ch := range newdata {
		sum += SurfaceArea(ch)
	}

	fmt.Println(sum)
}

func SurfaceArea(str string) int {
	res := 0
	l := ""
	w := ""
	h := ""
	newstr := strings.Split(str, "x")
	for i, ch := range newstr {
		if i == 0 {
			l = string(ch)
		}
		if i == 1 {
			w = string(ch)
		}
		if i == 2 {
			h = string(ch)
		}
		le, _ := strconv.Atoi(l)
		wi, _ := strconv.Atoi(w)
		he, _ := strconv.Atoi(h)
		min := MinArea((le * wi), MinArea((le*he), (wi*he)))
		res = (2 * le * wi) + (2 * le * he) + (2 * wi * he) + min

	}

	return res
}

func MinArea(x, y int) int {
	if x < y {
		return x
	}
	return y
}
