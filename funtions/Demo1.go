package main

import ("fmt")

func main() {
 // myMessage() // call the function

  //addNumbers(21, 13);// call with parameters

 //fmt.Println(add_v2(10,30)) // funtion with return type

// a, b := swap("hello", "world") // Getting multiple results from a method
//fmt.Println(a, b)

  fmt.Println(split(17))

}

func myMessage() {
  fmt.Println("I just got executed!")
}


func addNumbers(n1 int, n2 int) {
  sum := n1 + n2
  fmt.Println("Sum:", sum)
}

/*
 *	When two or more consecutive named function
 *	parameters share a type, you can omit the type from all but the last
 */
 func add_v2(x int , y int) int {
	return x + y
}

/*
 * A function can return any number of strings.
 * The swap function returns two strings.
 */
 func swap(x, y string) (string, string) {
	return y, x
}


// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}