package main

func main() {
	for i := 20000000; i <= 10000000000000; i++ {
		var isPrime bool = true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			println(i)

		}
	}
}
func add(x, y int) int {
	return x + y

}
