package main

import (
	"fmt"
	"math"
)

func quadraticSolver(a, b, c float32) (float32, float32) {
	discriminant := float32(b*b - 4*a*c)

	firstV := (-b + float32(math.Sqrt(float64(discriminant)))) / 2 * a
	secondV := (-b - float32(math.Sqrt(float64(discriminant)))) / 2 * a

	return firstV, secondV
}

func main() {
	a, b := quadraticSolver(1, -5, 4)
	fmt.Print(a, b)
}
