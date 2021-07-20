package main

import "fmt"

var c, python, java bool // Variables

func main()  {
	var i, j int // Variables
	fmt.Println(i, j, c, python, java)

	i, j = 1, 2 // Variables with initializers
	var c, python, java = true, false, "no!" // Variables with initializers
	fmt.Println(i, j, c, python, java)

	k := 3 // Short variable declarations
	fmt.Println(i, j, k, c, python, java)

	// Few examples
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}
