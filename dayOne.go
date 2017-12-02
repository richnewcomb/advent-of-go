package advent

import (
	"strconv"
)

func CalcCaptcha(captcha string, step int) int {

	position := 0
	nextPosition := 0
	currentValue := 0
	nextValue := 0
	checkSum := 0

	for position < len(captcha) {
		currentValue, _ = strconv.Atoi(string(captcha[position]))
		if position+step < len(captcha) {
			nextPosition = position + step
		} else {
			nextPosition = (position + step) - len(captcha)
		}
		nextValue, _ = strconv.Atoi(string(captcha[nextPosition]))
		if currentValue == nextValue {
			checkSum += currentValue
		}
		position++
	}
	return checkSum
}
