package main

import (
	"strconv"
	"strings"
	"unicode"
)

func wordToNum(word string) (rune, bool) {
	switch {
	case unicode.IsDigit(rune(word[0])):
		return rune(word[0]), true
	case strings.HasPrefix(word, "one"):
		return '1', true
	case strings.HasPrefix(word, "two"):
		return '2', true
	case strings.HasPrefix(word, "three"):
		return '3', true
	case strings.HasPrefix(word, "four"):
		return '4', true
	case strings.HasPrefix(word, "five"):
		return '5', true
	case strings.HasPrefix(word, "six"):
		return '6', true
	case strings.HasPrefix(word, "seven"):
		return '7', true
	case strings.HasPrefix(word, "eight"):
		return '8', true
	case strings.HasPrefix(word, "nine"):
		return '9', true
	}

	return '0', false
}

func Solve(line string) int {
	res := 0

	var first rune = 0
	var last rune = 0

	length := len(line)
	for i := 0; i < length; i++ {
		if num, ok := wordToNum(line[i:length]); ok {
			if first == 0 {
				first = num
			} else {
				last = num
			}
		}
	}

	if last == 0 {
		if first == 0 {
			return 0
		}
		last = first
	}

	var sb strings.Builder
	sb.WriteRune(first)
	sb.WriteRune(last)

	if num, err := strconv.Atoi(sb.String()); err != nil {
		panic(err)
	} else {
		res += num
	}

	return res
}
