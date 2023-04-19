package main

import (
	"bufio"
	"fmt"
	romans "github.com/summed/goromans"
	"os"
	"os/exec"
	"runtime"
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
		A, B, isRoman := convertOperands(aString, bString)
		res := calculate(A, B, operator)
		printResult(A, B, res, operator, isRoman)

	}

}

func printResult(a, b, res int, operator string, isRoman bool) {
	if isRoman {
		if res > 0 {
			aUint := uint(a)
			bUint := uint(b)
			resUint := uint(res)
			aR := romans.AtoR(aUint)
			bR := romans.AtoR(bUint)
			resR := romans.AtoR(resUint)
			fmt.Printf("Результат выражения %v %v %v = %v\n", aR, operator, bR, resR)
		} else {
			panic("Ошибка! В римской системе счета нет отрицательных значений, а так же нуля!")
		}
	} else {
		fmt.Printf("Результат выражения %v %v %v = %v\n", a, operator, b, res)
	}
}

func calculate(a, b int, operator string) int {
	var res int
	if a <= 10 && a > 0 && b <= 10 && b > 0 {
		switch operator {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		}
		return res
	} else {
		panic("Ошибка! Числа должны быть в диапазоне от 1 до 10")
	}
}

func convertOperands(aString, bString string) (a, b int, isRoman bool) {
	CallClear()
	if romans.IsRomanNumerals(aString) || romans.IsRomanNumerals(bString) {
		checkMinus := strings.Split(aString, "")
		for _, minus := range checkMinus {
			if minus == "-" {
				panic("Ошибка! В римской системе нет отрицательных чисел! Введите значение от I до X")
			}
		}
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
		panic("Ошибка! Операнды должны быть арабскими либо римскими числами!")
	}
	b, errb := strconv.Atoi(bString)
	if errb != nil {
		panic("Ошибка! Операнды должны быть арабскими либо римскими числами!")
	}
	return
}

func defineOperandsAndOperator(input string) (operand1, operand2, operator string) {
	arr := strings.Split(strings.ReplaceAll(input, " ", ""), "")
	isOperator := false
	i := 1
	for ; i < len(arr); i++ {
		for in := 0; in < len(availbleOperators); in++ {
			if arr[i] == availbleOperators[in] {
				isOperator = true
				operator = arr[i]
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
				break
			}
		}
	}
	if isOperator {
		return
	}
	panic("Ошибка! В выражении отсутсвует оператор")
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
