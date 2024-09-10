package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	datas, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("could not read from file")
	}
	defer datas.Close()
	scanner := bufio.NewScanner(datas)
	count := 0
	c := 0
	var sl []string
	for scanner.Scan() {
		data := scanner.Text()
		count++
		if count > 0 {
			sl = append(sl, data)
			c = NiceString(sl)
		}
	}
	fmt.Println(c)
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error scanning file: %w", err)
	}
}

func NiceString(s []string) int {
	c := 0
	for _, x := range s {
		if ThreVow(x) && ContLet(x) && NoStr(x) {
			c++
		}
	}
	return c
}

func ThreVow(s string) bool {
	v := "aeiou"
	c := 0
	for _, x := range s {
		if strings.ContainsRune(v, x) {
			c++
		}
	}
	return c >= 3
}

func ContLet(s string) bool {
	for i := range s {
		if i > 0 && s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func NoStr(s string) bool {
	no := []string{"ab", "cd", "pq", "xy"}
	for _, x := range no {
		if strings.Contains(s, x) {
			return false
		}
	}
	return true
}
