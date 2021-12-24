package service

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type CalcService struct {
	serviceType string
}

func NewCalcService() *CalcService {
	// var calcService = CalcService{}
	return &CalcService{}
}

func (this *CalcService) Calculate(expression string) (float64, error) {
	fmt.Println("expression", expression)
	expression = RemoveSpaceFromString(expression)
	var runes = []rune(expression)
	var index int = 0

	if len(runes) == 0 {
		return 0, errors.New("empty expression")
	}
	result, _, err := this.calcExpression(runes, index, false)
	return result, err
}

func (this *CalcService) calcExpression(runes []rune, index int, hasParenthesis bool) (float64, int, error) {
	result, index, err := this.calcTerm(runes, index)
	if err != nil {
		return result, index, err
	}

	for index < len(runes) {

		if index < len(runes) {
			if runes[index] != '-' && runes[index] != '+' && runes[index] != ')' {
				return result, index, errors.New("unknown symbol" + string(runes[index]))
			}
			var newResult float64
			if runes[index] == '+' {
				index++
				newResult, index, err = this.calcTerm(runes, index)
				if err != nil {
					return result, index, err
				}
				result += newResult
			} else if runes[index] == '-' {
				index++
				newResult, index, err = this.calcTerm(runes, index)
				if err != nil {
					return result, index, err
				}
				result -= newResult
			} else if hasParenthesis {
				if runes[index] == ')' {
					index++
					return result, index, err
				} else {
					return result, index, errors.New("Parenthesis don't match at index: " + fmt.Sprint(rune(index-1)))
				}
			} else if !hasParenthesis && runes[index] == ')' {
				return result, index, errors.New("Parenthesis don't match at index: " + fmt.Sprint(rune(index)))
			} else {
				break
			}
		}
	}
	if hasParenthesis && runes[index-1] != ')' {
		return result, index, errors.New("Parenthesis don't match at index: " + fmt.Sprint(index-1))
	}

	return result, index, nil

}

func (this *CalcService) calcTerm(runes []rune, index int) (float64, int, error) {
	result, index, err := this.calcPower(runes, index)
	if err != nil {
		return result, index, err
	}

	for index < len(runes) {
		if index < len(runes) {
			if runes[index] != '*' && runes[index] != '/' {
				break
			}
			var newResult float64
			if runes[index] == '*' {
				index++
				newResult, index, err = this.calcPower(runes, index)
				if err != nil {
					return result, index, err
				}
				result *= newResult
			} else if runes[index] == '/' {
				index++
				newResult, index, err = this.calcPower(runes, index)
				if err != nil {
					return result, index, err
				}
				result /= newResult
			} else {
				break
			}
		}
	}

	return result, index, err
}

func (this *CalcService) calcPower(runes []rune, index int) (float64, int, error) {
	result, index, err := this.getNumber(runes, index)
	if err != nil {
		return result, index, err
	}

	for index < len(runes) {
		if index < len(runes) {
			if runes[index] != '^' {
				break
			}
			var newResult float64
			index++
			newResult, index, err = this.getNumber(runes, index)
			if err != nil {
				return result, index, err
			}
			result = math.Pow(result, newResult)
		}
	}

	return result, index, err
}

func (this *CalcService) getNumber(runes []rune, index int) (float64, int, error) {
	var numberRunes []rune
	for index < len(runes) {
		if unicode.IsNumber(runes[index]) || runes[index] == '.' {
			numberRunes = append(numberRunes, runes[index])
			index++
		} else if runes[index] == '(' {
			index++
			return this.calcExpression(runes, index, true)
		} else {
			break
		}
	}
	var numberString string = string(numberRunes)
	result, err := strconv.ParseFloat(numberString, 64)
	return result, index, err
}

func RemoveSpaceFromString(str string) string {
	var stringBuilder strings.Builder
	stringBuilder.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			stringBuilder.WriteRune(ch)
		}
	}
	return stringBuilder.String()
}
