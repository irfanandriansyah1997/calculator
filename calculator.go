package calculator

import (
	"math"

	"github.com/irfanandriansyah1997/calculator/customerror"
)

type calculatorSignature interface {
	Sum(arg1, arg2 float64) (float64, error)
	Subtraction(arg1, arg2 float64) (float64, error)
	Multiplication(arg1, arg2 float64) (float64, error)
	Div(arg1, arg2 float64) (float64, error)
	Sqrt(arg1 float64) (float64, error)
}

type calculator struct{}

func (calc calculator) Sum(arg1, arg2 float64) (float64, error) {
	return arg1 + arg2, nil
}

func (calc calculator) Subtraction(arg1, arg2 float64) (float64, error) {
	return arg1 - arg2, nil
}

func (calc calculator) Multiplication(arg1, arg2 float64) (float64, error) {
	if arg1 == 0 || arg2 == 0 {
		return 0, customerror.InvalidArgError{Message: "Argument 1 / 2 is zero"}
	}

	return arg1 * arg2, nil
}

func (calc calculator) Div(arg1, arg2 float64) (float64, error) {
	if arg1 == 0 || arg2 == 0 {
		return 0, customerror.InvalidArgError{Message: "Argument 1 / 2 is zero"}
	}

	return arg1 / arg2, nil
}

func (calc calculator) Sqrt(arg1 float64) (float64, error) {
	if arg1 == 0 {
		return 0, customerror.InvalidArgError{Message: "Argument 1 is zero"}
	}

	return math.Sqrt(arg1), nil
}

func Calculate() calculatorSignature {
	instance := calculator{}

	return instance
}
