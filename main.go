package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		a, b, operator := defineOperandsAndOperator(input)
		fmt.Printf("%v %v %v", a, operator, b)
	}

	fmt.Println("Test gitlab")
}

func defineOperandsAndOperator(input string) (operand1, operand2, operator string) {
	arr := strings.Split(strings.ReplaceAll(input, " ", ""), "")
	isOperator := false
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case "+":
			operator = arr[i]
			arr = strings.Split(strings.Join(arr, ""), arr[i])
			if len(arr) == 2 {
				operand1 = arr[0]
				operand2 = arr[1]
			} else {
				panic("Ошибка! Введите 2 операнда и 1 оператор")
			}
			isOperator = true
		case "-":
			operator = arr[i]
			arr = strings.Split(strings.Join(arr, ""), arr[i])
			if len(arr) == 2 {
				operand1 = arr[0]
				operand2 = arr[1]
			} else {
				panic("Ошибка! Введите 2 операнда и 1 оператор")
			}
			isOperator = true
		case "*":
			operator = arr[i]
			arr = strings.Split(strings.Join(arr, ""), arr[i])
			if len(arr) == 2 {
				operand1 = arr[0]
				operand2 = arr[1]
			} else {
				panic("Ошибка! Введите 2 операнда и 1 оператор")
			}
			isOperator = true
		case "/":
			operator = arr[i]
			arr = strings.Split(strings.Join(arr, ""), arr[i])
			if len(arr) == 2 {
				operand1 = arr[0]
				operand2 = arr[1]
			} else {
				panic("Ошибка! Введите 2 операнда и 1 оператор")
			}
			isOperator = true
		}
	}
	if isOperator {
		return
	}
	panic("Ошибка! В выражении отсутсвует оператор")
}
