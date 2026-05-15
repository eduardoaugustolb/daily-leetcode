package day1

/*
Valid Parentheses
Dado uma string s contendo apenas os caracteres '(', ')', '{', '}', '[' e ']', determine se a string é válida.

Uma string é válida se: todo abre-parêntese é fechado pelo tipo correto, e na ordem correta. Todo fecha-parêntese tem um abre-parêntese correspondente.

Exemplos
Input:  "()"
Output: true
Input:  "()[]{}"
Output: true
Input:  "(]"
Output: false
*/

var opens = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
}

var matchingOpen = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func validateParentheses(text string) bool {
	stack := []rune{}

	for _, char := range text {
		if expected, isClose := matchingOpen[char]; isClose {
			if len(stack) == 0 || stack[len(stack)-1] != expected {
				return false
			}

			stack = stack[:len(stack)-1]
		} else if opens[char] {
			stack = append(stack, char)
		} else {
			return false
		}
	}
	return len(stack) == 0
}
