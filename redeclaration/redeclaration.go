package main

import "fmt"

func main() {
	/*
		:= can be used for redeclaraton and assignment in the same statement.
		safe is and existing variable, it will be assigned true, not redeclare
		but speed is declared and 50 is assigned to it.
	*/
	var safe bool
	safe, speed := true, 50
	fmt.Println(safe, speed)

}
