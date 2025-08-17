package main

import (
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		print("This is a mac")
	case "linux":
		print("This is Life")
	default:
		print("Na you sabi")
	}
}
