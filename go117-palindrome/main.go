package main

import (
	"strconv"
)

func isPalindrome1(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

func isPalindrome2(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Convert integer to string
	s := strconv.Itoa(x)

	// Check if string is palindrome
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
