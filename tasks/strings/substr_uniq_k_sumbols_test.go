package main

import (
	"fmt"
	"testing"
)

// Дана строка, нужно написать функцию которая вернет длину наибольшей подстроки,
// которая содержит не более k уникальных символов
//(строка состоит из уникальных символов, и кол-во этих символов не больше k)
// "abbc", k=2 => "abb", 3
// "aaa", k=1 => "aaa", 3

// aabccaddd  // k=2
// abaccaddd  // k=2
// a

func UniqLen(input string, k int) int {
	chars := make(map[rune]int)
	start := 0
	maxLen := 0
	for i, s := range input {
		chars[s] += 1
		for len(chars) > k { // aaaaaaaaaaaaabbbbbbbbc
			rms := rune(input[start])
			chars[rms] -= 1
			if chars[rms] == 0 {
				delete(chars, rms)
			}
			start++
		}
		ml := i - start + 1
		if ml > maxLen {
			maxLen = ml
		}
	}

	return maxLen
}

func TestUniqLen(t *testing.T) {
	type tcase struct {
		input    string
		k        int
		expected int
	}
	cases := []tcase{
		{
			input:    "abbc",
			k:        2,
			expected: 3, // abb
		},
		{
			input:    "aaa",
			k:        1,
			expected: 3, // "aaa"
		},
		{
			input:    "aaaaaaabcc",
			k:        3,
			expected: 10, // "aaaaaaabcc"
		},
	}

	for _, c := range cases {
		got := UniqLen(c.input, c.k)
		if got != c.expected {
			t.Errorf("case %v, got %d, want %d", c, got, c.expected)
		} else {
			fmt.Printf("success case %v", c)
			fmt.Println("success case %v", c)
		}
	}
}
