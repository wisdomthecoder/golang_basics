package main

func main() {
	v := Vertex{Y: 20}
	v.Y = 20
	println(v.X)
	println(&v.Y)
}

type Vertex struct {
	X int
	Y int
}
