package main

import "math"

func main() {
	for i := 1; i < 1000000; i++ {
		if isPerfectSquare(float64(i)) {
			println(i, " is a perfect square")
		}
	}
}
func isPerfectSquare(v float64) bool {

	if root := float32(math.Sqrt(v)); root == float32(int(root)) {
		return true
	} else {
		return false
	}

}
