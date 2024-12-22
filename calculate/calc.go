package calculate

import (
	"strconv"
	"strings"
)

func isSign(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}
func strToFloat(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return value
}

func Calc(expression string) (float64, error) {
	var (
		result    float64
		sign      rune = 0
		isc       int
		countsign int = 0
	)
	if len(expression) == 0 {
		return 0, ErrInvalidExpression
	}
	rune_expression := []rune(expression)
	if isSign(rune_expression[0]) || isSign(rune_expression[len(expression)-1]) {
		return 0, ErrInvalidExpression
	}
	for _, value := range expression {
		if isSign(value) {
			countsign++
		}
	}
	for i, value := range expression {
		if value == '(' {
			isc = i
		}
		if value == ')' {
			calc, err := Calc(expression[isc+1 : i])
			if err != nil {
				return 0, err
			}

			calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
			i2 := i
			i -= len(expression[isc:i+1]) - len(calcstr)
			expression = strings.Replace(expression, expression[isc:i2+1], calcstr, 1)
		}
	}
	if countsign > 1 {
		for i := 1; i < len(expression); i++ {
			value := rune(expression[i])
			if value == '*' || value == '/' {
				min := i - 1
				if min != 0 {
					for !isSign(rune(expression[min])) && min > 0 {
						min--
					}
					min++
				}
				max := i + 1
				if max == len(expression) {
					max--
				} else {
					for !isSign(rune(expression[max])) && max < len(expression)-1 {
						max++
					}
				}
				if max == len(expression)-1 {
					max++
				}
				calc, err := Calc(expression[min:max])
				if err != nil {
					return 0, err
				}
				calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
				i -= len(expression[isc:i+1]) - len(calcstr) - 1
				expression = strings.Replace(expression, expression[min:max], calcstr, 1)
			}
			if value == '+' || value == '-' || value == '*' || value == '/' {
				sign = value
			}
		}
	}
	num := ""
	signflag := false
	for _, value := range expression {
		if value >= '0' && value <= '9' {
			num += string(value)
		} else if isSign(value) {
			if signflag {
				switch sign {
				case '+':
					result += strToFloat(num)
				case '-':
					result -= strToFloat(num)
				case '*':
					result *= strToFloat(num)
				case '/':
					if strToFloat(num) == 0 {
						return 0, ErrZeroDivision
					}
					result /= strToFloat(num)
				}
			} else {
				result = strToFloat(num)
			}
			num = ""
			sign = value
			signflag = true
		} else {
			return 0, ErrIncorrectInput
		}
	}
	if signflag {
		switch sign {
		case '+':
			result += strToFloat(num)
		case '-':
			result -= strToFloat(num)
		case '*':
			result *= strToFloat(num)
		case '/':
			if strToFloat(num) == 0 {
				return 0, ErrZeroDivision
			}
			result /= strToFloat(num)
		}
	} else {
		result = strToFloat(num)
	}

	return result, nil
}
