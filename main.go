package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var availbleOperators = []string{"+", "-", "*", "/"}
var sets = []struct {
	arabic int
	rome   string
	max    int
}{
	{100, "C", 1},
	{90, "XC", 1},
	{50, "L", 1},
	{40, "XL", 1},
	{10, "X", 2},
	{9, "IX", 1},
	{5, "V", 1},
	{4, "IV", 1},
	{1, "I", 3},
}

// test git ubuntu
func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		input = strings.ReplaceAll(strings.ToUpper(input), " ", "")
		a, b, op := defineABOpString(input)
		fmt.Printf("%v %v %v = %v\n", a, op, b, calc(input))

	}
}

func isRoman(input string) bool {
	inputArr := strings.Split(input, "")
	i := 0
	for _, value := range inputArr {
		for _, set := range sets {
			if value == set.rome {
				i++
				break
			}
		}
		if i == len(inputArr) {
			return true
		}
	}
	return false
}

func defineABOpString(input string) (aString, bString, operator string) {
	arr := strings.Split(input, "")
	isOperator := false
	for i := 1; i < len(arr); i++ {
		for in := 0; in < len(availbleOperators); in++ {
			if arr[i] == availbleOperators[in] {
				isOperator = true
				operator = arr[i]
				arr = strings.Split(strings.Join(arr, ""), arr[i])
				if len(arr) == 2 {
					aString = arr[0]
					bString = arr[1]
					if aString == "" || bString == "" {
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

func abStringToInt(aString, bString string) (aInt, bInt int) {
	if isRoman(aString) {
		aInt = convRtoA(aString)
	} else {
		aInt, _ = strconv.Atoi(aString)
	}
	if isRoman(bString) {
		bInt = convRtoA(bString)
	} else {
		bInt, _ = strconv.Atoi(bString)
	}
	return aInt, bInt
}

func convRtoA(romeNum string) int {
	arrString := strings.Split(romeNum, "") // создаю массив с римскими значениями
	var res int
	var arrInt []int
	for _, romeNum := range arrString { // Цикл созздает массив c арабскими числами
		for _, set := range sets {
			if romeNum == set.rome {
				var count = 0
				for _, val := range arrString {
					if set.rome == val {
						count++
					}
				}
				if count <= set.max {
					arrInt = append(arrInt, set.arabic)
				} else {
					panic("Ошибка! Введите корректные римские цифры")
				}
			}

		}
	}

	for i, num := range arrInt { // Логика сложений и вычитаний значений в арабском массиве
		if i+1 != len(arrInt) {
			if num >= arrInt[i+1] {
				res += num
			} else {
				res -= num
			}
		} else if num >= arrInt[i] {
			res += num
		} else {
			res -= num
		}
	}
	return res
}

func convAtoR(number int) string {

	roman := ""
	for _, set := range sets {
		for number >= set.arabic {
			roman += set.rome
			number -= set.arabic
		}
	}
	return roman
}

func calc(input string) (result string) {
	aString, bString, op := defineABOpString(input)
	aInt, bInt := abStringToInt(aString, bString)
	var res int
	if aInt > 0 && aInt <= 10 && bInt > 0 && bInt <= 10 {
		switch op {
		case "+":
			res = aInt + bInt
		case "-":
			res = aInt - bInt
		case "*":
			res = aInt * bInt
		case "/":
			res = aInt / bInt
		}

		if isRoman(aString) || isRoman(bString) {
			if isRoman(aString) && isRoman(bString) {
				if res == 0 {
					result = "(В римской системе исчисления нет нуля)"
					return result
				} else if res < 0 {
					result = "(В римской системе исчисления нет отрицательных значений)"
					return result
				}
			} else {
				panic("Ошибка! Введите либо римские либо арабские числа")
			}
			result = convAtoR(res)
			return result
		} else {
			result = strconv.Itoa(res)
			return result
		}
	} else {
		panic("Ошибка! Введите римские либо арабские числа от 1 до 10")
	}
}
