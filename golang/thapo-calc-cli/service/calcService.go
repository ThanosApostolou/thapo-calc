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

func (this *CalcService) CalcExpression(expression string) (float64, error) {
	fmt.Println("expression", expression)
	var result float64 = 0
	expression = RemoveSpaceFromString(expression)
	var runes = []rune(expression)
	var index int = 0

	if len(runes) == 0 {
		return result, errors.New("empty expression")
	}

	result, index, err := this.calcTerm(runes, index)
	if err != nil {
		return result, err
	}

	for index < len(runes) {

		if index < len(runes) {
			if runes[index] != '-' && runes[index] != '+' {
				return result, errors.New("unknown symbol" + string(runes[index]))
			}
			var newResult float64
			if runes[index] == '+' {
				index++
				newResult, index, err = this.calcTerm(runes, index)
				result += newResult
			} else if runes[index] == '-' {
				index++
				newResult, index, err = this.calcTerm(runes, index)
				result -= newResult
			}
		}
	}

	return result, nil
}

func expression(runes []rune) {

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
				result *= newResult
			} else if runes[index] == '/' {
				index++
				newResult, index, err = this.calcPower(runes, index)
				result /= newResult
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
