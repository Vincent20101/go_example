package main

//给定一个如下格式的字符串：(1，(2，3)，(4，(5，6)，7))，括号内的元素可以是数字，
//也可以是另一个括号，实现一个算法消除嵌套的括号，例如把上面的表达式变成(1，2，3，4，5，6，7)，如果表达式有误，则报错。
import (
	"fmt"
	"strings"
	"unicode"
)

func flattenNestedExpression(input string) (string, error) {
	var stack []string
	var current strings.Builder

	for i, ch := range input {
		switch {
		case unicode.IsDigit(ch) || ch == ',':
			current.WriteRune(ch)
		case ch == '(':
			if current.Len() > 0 {
				stack = append(stack, current.String())
				current.Reset()
			}
			stack = append(stack, string(ch))
		case ch == ')':
			if len(stack) == 0 {
				return "", fmt.Errorf("unmatched closing parenthesis at position %d", i)
			}
			closed := current.String()
			current.Reset()
			// Pop until finding the opening '('
			for len(stack) > 0 {
				ele := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if ele == "(" {
					break
				}
				current.WriteString(ele)
			}
			current.WriteString(closed)
		default:
			return "", fmt.Errorf("invalid character '%c' at position %d", ch, i)
		}
	}

	if current.Len() > 0 {
		stack = append(stack, current.String())
	}

	// Rebuild the output from stack elements
	var result strings.Builder
	for _, str := range stack {
		result.WriteString(str)
	}
	// Final check for unmatched opening '('
	if strings.Contains(result.String(), "(") {
		return "", fmt.Errorf("unmatched opening parenthesis")
	}

	return result.String(), nil
}

func main() {
	input := "(1,(2,3),(4,(5,6),7))"
	output, err := flattenNestedExpression(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Flattened expression:", output)
	}
}
