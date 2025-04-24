package math

type goMath struct {
	Add func(a, b int) int
	Sub func(a, b int) int
	Mul func(a, b int) int
	Div func(a, b int) int
	Pow func(a, b int) int
}

func New() goMath {
	return goMath{
		Add: add,
		Sub: sub,
		Mul: mul,
		Div: div,
		Pow: pow,
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

func pow(a, b int) int {
	var result int = 1
	if a == 0 {
		return 0
	}
	if b == 0 {
		return 1
	}
	for i := 1; i <= b; i++ {
		result = result * a
	}
	return result
}
