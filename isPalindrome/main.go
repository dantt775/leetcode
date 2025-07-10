package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
}

//https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(num int) bool {
	reverseNum := 0
	expected := num

	for num > 0 {
		reverseNum = reverseNum*10 + (num % 10)
		num = num / 10
	}

	return reverseNum == expected
}
