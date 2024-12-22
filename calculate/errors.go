package calculate

import "errors"

var (
	ErrZeroDivision      = errors.New("division by zero")
	ErrIncorrectInput    = errors.New("incorrect input")
	ErrInvalidExpression = errors.New("expression is not valid")
)
