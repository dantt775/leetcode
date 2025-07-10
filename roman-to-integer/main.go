package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

//https://leetcode.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	romanMap := make(map[rune]int)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000
	sum := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		if i+1 < len(s) && romanMap[runes[i]] < romanMap[runes[i+1]] {
			sum -= romanMap[runes[i]]
		} else {
			sum += romanMap[runes[i]]
		}
	}
	return sum
}
