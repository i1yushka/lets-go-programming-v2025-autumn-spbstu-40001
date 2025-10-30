package main

import (
	"errors"
	"fmt"
)

var (
	ErrZeroDivision     = errors.New("Division by zero")
	ErrInvalidOperation = errors.New("Invalid operation")
)

func main() {
	var (
		lhs, rhs  int
		operation string
	)

	scanned, err := fmt.Scan(&lhs, &rhs, &operation)
	if err != nil {
		switch scanned {
		case 0:
			fmt.Println("Invalid first operand")
		case 1:
			fmt.Println("Invalid second operand")
		case 2:
			fmt.Println("Invalid operation")
		default:
			fmt.Println("Input error:", err)
		}
		return
	}

	res, err := calc(lhs, operation, rhs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func calc(lhs int, operation string, rhs int) (int, error) {
	switch operation {
	case "+":
		return lhs + rhs, nil
	case "-":
		return lhs - rhs, nil
	case "*":
		return lhs * rhs, nil
	case "/":
		if rhs == 0 {
			return 0, ErrZeroDivision
		}
		return lhs / rhs, nil
	}
	return 0, ErrInvalidOperation
}
