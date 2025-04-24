package math

type goMath struct {
	Add func(a, b int) int
	Sub func(a, b int) int
	Mul func(a, b int) int
	Div func(a, b int) int
}

func New() goMath {
	return goMath{
		Add: add,
		Sub: sub,
		Mul: mul,
		Div: div,
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}
