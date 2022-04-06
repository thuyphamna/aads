package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	testString1 := "tactcoa"
	// testString2 := "gfedca4"
	fmt.Println(checkPermutationOfPalindrome(testString1))

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
