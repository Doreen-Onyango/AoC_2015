package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("error opening file", err)
	}
	defer file.Close()

	var totalCodeChar, totalMemChar int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		codeLen := len(line)
		memLen := countMemCharacters(line)

		totalCodeChar += codeLen
		totalMemChar += memLen
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	dif := totalCodeChar - totalMemChar
	fmt.Println(dif)
}

func countMemCharacters(str string) int {
	str = str[1 : len(str)-1]
	i := 0
	count := 0
	for i < len(str) {
		if str[i] == '\\' {
			if i+1 < len(str) && str[i+1] == 'x' {
				if i+3 < len(str) {
					_, err := strconv.ParseInt(str[i+2:i+4], 16, 0)
					if err == nil {
						count++
					}
					i += 4
					continue
				}
			} else if i+1 < len(str) && (str[i+1] == '\\' || str[i+1] == '"') {
				count++
				i += 2
				continue
			}
		} else {
			count++
		}
		i++
	}
	return count
}
