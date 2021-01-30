package main

import "fmt"

func main() {
	var safe bool // safe is existing variable, it will be assigned true, not redeclare
	safe, speed := true, 50
	fmt.Println(safe, speed)

}
