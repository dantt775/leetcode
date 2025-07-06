package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcdbb"))
	fmt.Println(lengthOfLongestSubstringSlice([]string{"a", "b", "b", "a"}))
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	var runeIndexMap = make(map[rune]int)
	runes := []rune(s)
	maxLength := 0
	left := 0
	for right := 0; right < len(runes); right++ {

		currentRune := runes[right]

		if index, found := runeIndexMap[currentRune]; found && index >= left {
			left = index + 1 // Move left pointer to the right of the last occurrence
		}
		runeIndexMap[currentRune] = right
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

	}

	return maxLength
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstringSlice(s []string) int {
	var stringIndexMap = make(map[string]int)
	maxLength := 0
	windowStartIndex := 0
	for windowEndIndex := range s {
		currentString := s[windowEndIndex]

		if repeatedStringIndex, found := stringIndexMap[currentString]; found && repeatedStringIndex >= windowStartIndex {
			windowStartIndex = repeatedStringIndex + 1 // Move left pointer to the right of the last occurrence
		}
		stringIndexMap[currentString] = windowEndIndex
		currentLength := windowEndIndex - windowStartIndex + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}

	}

	return maxLength
}
