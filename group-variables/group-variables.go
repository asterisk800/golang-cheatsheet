package main

import "fmt"

// Use group declaration for a better readability and to group closely related items
var (
	name  string
	speed float64
	ratio float64
)

var speed1, ratio2 float64

func main() {

	fmt.Println(name, speed, ratio, speed1, ratio2)
}
