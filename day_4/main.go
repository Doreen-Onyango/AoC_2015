package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	skey := "bgvyzdsv"
	fmt.Println(LowNum(skey))
	fmt.Println(LowNumber(skey))

}

func Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func LowNum(sKey string) int {
	prefix := "00000"
	i := 0

	for {
		input := sKey + strconv.Itoa(i)
		hash := Hash(input)

		if strings.HasPrefix(hash, prefix) {
			return i
		}

		i++
	}
}

func LowNumber(sKey string) int {
	prefix := "000000"
	i := 0

	for {
		input := sKey + strconv.Itoa(i)
		hash := Hash(input)

		if strings.HasPrefix(hash, prefix) {
			return i
		}
		i++
	}
}
