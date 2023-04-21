package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unicode"
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
		CallClear() //Очистка экрана
		printResult(A, B, res, operator, isRoman)

	}

}

func printResult(a, b, res int, operator string, isRoman bool) {
	if isRoman {
		if res > 0 {
			aR := AtoR(uint(a))
			bR := AtoR(uint(b))
			resR := AtoR(uint(res))
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
	if IsRomanNumerals(aString) || IsRomanNumerals(bString) {
		checkMinus := strings.Split(aString, "")
		for _, minus := range checkMinus {
			if minus == "-" {
				panic("Ошибка! В римской системе нет отрицательных чисел! Введите значение от I до X")
			}
		}
		isRoman = IsRomanNumerals(aString) && IsRomanNumerals(bString)
		if isRoman {
			aUint, _ := RtoA(aString)
			bUint, _ := RtoA(bString)
			a = int(aUint)
			b = int(bUint)
			return
		} else {
			panic("Ошибка! Операнды должны быть одной системы!")
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
	for i := 1; i < len(arr); i++ {
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

type numeralByArabic []numeral

func (v numeralByArabic) Len() int            { return len(v) }
func (v numeralByArabic) Swap(this, that int) { v[this], v[that] = v[that], v[this] }
func (v numeralByArabic) Less(this, that int) bool {
	return v[this].arabic > v[that].arabic
}

type numeral struct {
	roman  rune
	arabic uint
}

var (
	sets = []numeral{
		numeral{'M', 1000},
		numeral{'D', 500},
		numeral{'C', 100},
		numeral{'L', 50},
		numeral{'X', 10},
		numeral{'V', 5},
		numeral{'I', 1},
	}

	initialized   bool
	romanNumerals = make(map[rune]numeral)
)

func initialize() {
	if !initialized {
		for _, s := range sets {
			romanNumerals[s.roman] = s
		}
		sort.Sort(numeralByArabic(sets)) // Strictly not required, since the correct (DESC) order is set at initialization.
		initialized = true
	}
}

// IsRomanNumerals returns true if able to parse string as roman numerals
func IsRomanNumerals(romans string) bool {
	initialize()
	if _, err := RtoA(romans); err != nil {
		return false
	}
	return true
}

// RtoA converts a string of roman numerals to arabic numerals
func RtoA(romans string) (out uint, err error) {
	initialize()
	var last uint
	if len(romans) == 0 {
		return out, fmt.Errorf("Empty string when parsing to roman numerals")
	}
	for i := 0; i < len(romans); i++ {
		if s, ok := romanNumerals[unicode.ToUpper(rune(romans[i]))]; ok {
			if s.arabic > last && last > 0 {
				out -= 2 * last
			}
			out += s.arabic
			last = s.arabic
		} else {
			return out, fmt.Errorf("Unable to '%s' to roman numerals, because of character '%c'", romans, romans[i])
		}
	}
	return out, nil
}

// AtoR converts arabic numerals to roman numerals
func AtoR(arabic uint) (romans string) {
	initialize()
	var out bytes.Buffer
	var major, minor numeral
	for arabic > 0 {
		for i := 0; i < len(sets); i++ {
			major = sets[i]

			if arabic == major.arabic {
				out.WriteRune(major.roman)
				arabic -= major.arabic
				goto loopEnd
			}

			if i < len(sets) {
				for j := i + 1; j < len(sets); j++ {
					minor = sets[j]
					if major.arabic/minor.arabic == 2 { // if minor is half of major (M&D, C&L, X&V), then skip - and let it be 'handled' as major later
						continue
					}
					if arabic-(major.arabic-minor.arabic) == 0 {
						out.WriteRune(minor.roman)
						out.WriteRune(major.roman)
						arabic -= major.arabic - minor.arabic
						goto loopEnd
					}
				}
			}

			if arabic > major.arabic {
				divs := uint(arabic / major.arabic)
				arabic -= major.arabic * divs
				for ; divs > 0; divs-- {
					out.WriteRune(major.roman)
				}
				goto loopEnd
			}
		}
	loopEnd:
	}

	return out.String()
}
