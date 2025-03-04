// Package ch10_exercises_module is a simple package with an add function.
// It is part of the Learning Go textbook exercises.
package ch10_exercises_module

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes two numbers (int or float) x and y, and returns the sum, as per
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
