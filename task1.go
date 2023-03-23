package main

import "fmt"

func isLeftBracket(ch rune) bool {
	var leftBrackets = []rune{'(', '{', '['}

	for _, l := range leftBrackets {
		if ch == l {
			return true
		}
	}

	return false
}

func checkByRight(ms []rune, ch rune) bool {
	if len(ms) == 0 {
		return false
	}

	var rightToLeftPairs = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	left, ok := rightToLeftPairs[ch]
	if !ok {
		return false
	}

	return ms[len(ms)-1] == left
}

func BracketsBalanced(str string) bool {
	var minimizedStack []rune

	for _, ch := range str {
		if isLeftBracket(ch) {
			minimizedStack = append(minimizedStack, ch)
		} else {
			if !checkByRight(minimizedStack, ch) {
				return false
			}

			minimizedStack = minimizedStack[:len(minimizedStack)-1]
		}
	}

	return len(minimizedStack) == 0
}

func ShowTask1() {
	ok := "{{[]}}(){}[()]"
	wrong := "{{]))]"

	fmt.Printf("%v - %v\n", ok, BracketsBalanced(ok))
	fmt.Printf("%v - %v\n", wrong, BracketsBalanced(wrong))
}
