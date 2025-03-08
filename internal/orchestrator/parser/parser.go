package orchestrator

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Type  string 
	Value string 
}

func precedence(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func ParseExpression(expr string) ([]Token, error) {
	expr = strings.ReplaceAll(expr, " ", "") 
	var output []Token
	var operators []string

	i := 0
	for i < len(expr) {
		char := string(expr[i])

		if _, err := strconv.Atoi(char); err == nil {
			numStr := ""
			for i < len(expr) {
				char = string(expr[i])
				if _, err := strconv.Atoi(char); err != nil {
					break
				}
				numStr += char
				i++
			}
			output = append(output, Token{Type: "number", Value: numStr})
			continue
		}

		if strings.Contains("+-*/", char) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(char) {
				output = append(output, Token{Type: "operator", Value: operators[len(operators)-1]})
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, char)
			i++
			continue
		}

		if char == "(" {
			operators = append(operators, char)
			i++
			continue
		}

		if char == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				output = append(output, Token{Type: "operator", Value: operators[len(operators)-1]})
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, fmt.Errorf("mismatched parentheses")
			}
			operators = operators[:len(operators)-1]
			i++
			continue
		}

		return nil, fmt.Errorf("invalid character: %s", char)
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, fmt.Errorf("mismatched parentheses")
		}
		output = append(output, Token{Type: "operator", Value: operators[len(operators)-1]})
		operators = operators[:len(operators)-1]
	}

	return output, nil
}
