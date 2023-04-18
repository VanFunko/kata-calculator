package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var availbleOperators = []string{"+", "-", "*", "/"}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		a, b, operator := defineOperandsAndOperator(input)
		fmt.Printf("%v %v %v\n", a, operator, b)
	}

}

//func calculate (a, b int, operator string, isRoman bool) int {
//
//}

func defineOperandsAndOperator(input string) (operand1, operand2, operator string) {
	arr := strings.Split(strings.ReplaceAll(input, " ", ""), "")
	isOperator := false
	for i := 0; i < len(arr); i++ {
		for in := 0; in < len(availbleOperators); in++ {
			if arr[i] == availbleOperators[in] {
				isOperator = true
				operator = arr[i]
				//fmt.Printf("arr[i] = %v, oper = %v\n", arr[i], availbleOperators[in])
				arr = strings.Split(strings.Join(arr, ""), arr[i])
				if len(arr) == 2 {
					operand1 = arr[0]
					operand2 = arr[1]
					if operand1 == "" || operand2 == "" {
						panic("Ошибка! Операнды не могут отсутсвовать!")
					}
				} else {
					panic("Ошибка! Введите 2 операнда и 1 оператор")
				}
			}
		}
	}
	if isOperator {
		return
	}
	panic("Ошибка! В выражении отсутсвует оператор")
}
