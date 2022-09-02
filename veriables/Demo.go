package main

import "fmt"

// var c, python, java bool

// func main() {
// 	var i int
// 	fmt.Println(i, c, python, java)
// }

var i, j int = 1, 2
const Pi = 3.14

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	//varDemo();
	//varDefaultValue()
	constant()
}

// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

// Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
func varDemo() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

func varDefaultValue() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}


// Constants are declared like variables, but with the const keyword.

// Constants can be character, string, boolean, or numeric values.

// Constants cannot be declared using the := syntax.

func constant() {
	const World = "GIC"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}