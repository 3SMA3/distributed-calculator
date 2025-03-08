package agent

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Type  string 
	Value string 
}

func ComputeExpression(expr string) (float64, error) {
	tokens, err := parseExpression(expr)
	if err != nil {
		return 0, err
	}

	stack := []float64{}
	for _, token := range tokens {
		if token.Type == "number" {
			num, _ := strconv.ParseFloat(token.Value, 64)
			stack = append(stack, num)
		} else if token.Type == "operator" {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid expression")
			}
			arg2 := stack[len(stack)-1]
			arg1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			result, err := computeOperation(arg1, arg2, token.Value)
			if err != nil {
				return 0, err
			}
			stack = append(stack, result)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("invalid expression")
	}
	return stack[0], nil
}

func parseExpression(expr string) ([]Token, error) {
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

func computeOperation(arg1, arg2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return arg1 + arg2, nil
	case "-":
		return arg1 - arg2, nil
	case "*":
		return arg1 * arg2, nil
	case "/":
		if arg2 == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return arg1 / arg2, nil
	default:
		return 0, fmt.Errorf("unknown operation")
	}
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
