package main

import (
	"bufio"
	"fmt"
	romans "github.com/summed/goromans"
	"os"
	"strconv"
	"strings"
)

var availbleOperators = []string{"+", "-", "*", "/"}

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		aString, bString, operator := defineOperandsAndOperator(input)
		fmt.Printf("%v %v %v\n", aString, operator, bString)
		A, B, isRoman := convertOperands(aString, bString)
		fmt.Printf("A = %v, B = %v, isRoman %v\n", A, B, isRoman)

	}

}

func convertOperands(aString, bString string) (a, b int, isRoman bool) {
	if romans.IsRomanNumerals(aString) || romans.IsRomanNumerals(bString) {
		isRoman = romans.IsRomanNumerals(aString) && romans.IsRomanNumerals(bString)
		if isRoman {
			aUint, _ := romans.RtoA(aString)
			bUint, _ := romans.RtoA(bString)
			a = int(aUint)
			b = int(bUint)
			return a, b, isRoman
		} else {
			panic("Ошибка! Опреанды должны быть одной системы!")
		}
	}

	a, erra := strconv.Atoi(aString)
	if erra != nil {
		panic("Ошибка! Операнд должен быть числом!")
	}
	b, errb := strconv.Atoi(bString)
	if errb != nil {
		panic("Ошибка! Операнд должен быть числом!")
	}
	return
}

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
