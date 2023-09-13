package funciones

import "errors"

var (
	ErrDivisionZero = errors.New("no se puede dividir por cero")
)

func Suma(a, b int) int {
	return a + b
}

func Division(a, b int) (int, error) {
	if b < 1 {
		return 0, ErrDivisionZero
	}
	return a / b, nil
}
