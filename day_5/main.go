package main

import (
	"fmt"
)

func main() {
	// data, err := os.ReadFile("puzzle.txt")
	// if err != nil {
	// 	fmt.Println("could not read from file")
	// }
	// fmt.Println(StrStatus(string(data)))
	x := "advfdfgitfrderttyd"
	fmt.Println(StrStatus(x))

}
func IsContained(s string) bool {
	return s == "aa" || s == "bb" || s == "cc" || s == "dd" || s == "ee" || s == "ff" || s == "gg" || s == "hh" || s == "ii" || s == "jj" || s == "kk"
}
func NotAllowed(s string) bool {
	return s == "ab" || s == "cd" || s == "pq" || s == "xy"
}

// part 1:
func StrStatus(s string) string {
	count := 0
	vow := "aeiou"
	for _, c := range s {
		for _, v := range vow {
			if c == v {
				count++
				if count == 3 {
					if IsContained(string(c)) {
						if !NotAllowed(string(c)) {
							return "Nice"
						}
					}

				}
				return "Naughty"

			}
		}
	}
	return "nice"
}
