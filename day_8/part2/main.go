package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("error opening file")
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)

	totalOriginalLen := 0
	totalEncodedLen := 0
	for scanner.Scan() {
		strToEncode := scanner.Text()

		originalLen := len(strToEncode)
		encodedStr := encodeString(strToEncode)
		encodedLen := len(encodedStr)

		totalOriginalLen += originalLen
		totalEncodedLen += encodedLen
	}

	dif := totalEncodedLen - totalOriginalLen
	fmt.Println(dif)
}

func encodeString(s string) string {
	x := strings.ReplaceAll(s, "\\", "\\\\")
	y := strings.ReplaceAll(x, "\"", "\\\"")
	return "\"" + y + "\""
}
