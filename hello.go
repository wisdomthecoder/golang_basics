package main

import (
	"time"
)

func main() {
	t := time.Now()
	defer println("Hello Wisdom")

	switch {
	case t.Hour() < 12:
		println("Chief Good morning")
	case t.Hour() < 17:
		println("Chief Good Afternoon")
	default:
		println("Chief Relax")
	}
}
