package xstr

import (
	"math/rand"
)

// TODO: Add some string random generation function

// Random generate string
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func Random(n int) string {
	return ""
}

// Letters a-Z
const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Numbers 0-9
const Numbers = "0123456789"

// LettersNumbers a-z0-9
const LettersNumbers = Letters + Numbers

func randomStr(n, t int) string {
	b := make([]byte, n)
	for i := range b {
		switch t {
		case 0:
			b[i] = Letters[rand.Intn(len(Letters))]
		case 1:
			b[i] = Numbers[rand.Intn(len(Numbers))]
		case 2:
			b[i] = LettersNumbers[rand.Intn(len(LettersNumbers))]
		default:
			b[i] = Letters[rand.Intn(len(Letters))]
		}
	}
	return string(b)
}

// RandomCharStr random-string only with letters
func RandomCharStr(n int) string {
	return randomStr(n, 0)
}

// RandomNumStr random-string only with numbers
func RandomNumStr(n int) string {
	return randomStr(n, 1)
}

// RandomStr random-string with letters and numbers
func RandomStr(n int) string {
	return randomStr(n, 2)
}
