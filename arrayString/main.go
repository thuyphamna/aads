package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(string(1))

}

func isUnique(s string) bool {
	mapOfChar := make(map[rune]int)
	for _, char := range s {
		_, exist := mapOfChar[char]
		if !exist {
			mapOfChar[char] = 1
		} else {
			return false
		}
	}
	return true
}

func isUniqueWithoutMap(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func checkPermutation(s string, a string) bool {
	sArray := strings.Split(s, "")
	aArray := strings.Split(a, "")
	sort.Strings(sArray)
	sort.Strings(aArray)
	fmt.Println(strings.Join(sArray, ""))
	fmt.Println(strings.Join(aArray, ""))

	return strings.Join(sArray, "") == strings.Join(aArray, "")
}

func URLify(s string) string {
	newS := strings.Replace(s, " ", "%20", -1)

	return newS
}

func checkPermutationOfPalindrome(s string) bool {
	countChar := make(map[rune]int)
	for _, r := range s {
		count, exist := countChar[r]
		if !exist {
			countChar[r] = 1
		} else {
			countChar[r] = count + 1
		}
	}
	noPairChar := 0
	for _, val := range countChar {
		if val%2 != 0 {
			noPairChar++
		}
	}

	if noPairChar != 1 {
		return false
	} else {
		return true
	}

}

func stringCompression(s string) string {
	result := ""

	charToInsert := s[0]
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == charToInsert {
			count++
		} else {
			result = result + string(charToInsert) + strconv.Itoa(count)
			charToInsert = s[i]
			count = 1
		}
	}
	result = result + string(charToInsert) + strconv.Itoa(count)

	if len(result) > len(s) {
		return s
	}

	return result
}
