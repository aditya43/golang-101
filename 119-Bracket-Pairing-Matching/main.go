package main

import "fmt"

//bracketType sorts characters into either opening, closing, or not a bracket
type bracketType int

const (
	openBracket bracketType = iota
	closeBracket
	notABracket
)

//bracketPairs are the matching pairs of brackets
var pairs = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
}

/*Bracket determines if a strings has balanced brackets*/
func Bracket(phrase string) (bool, error) {
	var queue []rune

	for _, v := range phrase {
		switch getBracketType(v) {
		case openBracket:
			queue = append(queue, pairs[v])
		case closeBracket:
			if 0 < len(queue) && queue[len(queue)-1] == v {
				queue = queue[:len(queue)-1]
			} else {
				return false, nil
			}
		}
	}

	return len(queue) == 0, nil
}

/*getBracketType determines the type of bracket*/
func getBracketType(char rune) bracketType {
	for k, v := range pairs {
		switch char {
		case k:
			return openBracket
		case v:
			return closeBracket
		}
	}
	return notABracket
}

func main() {
	fmt.Println(Bracket("{ a + b; (a - b); { d - x }}"))
	fmt.Println(Bracket("{ a + b; (a - b); d - x }}"))
}
