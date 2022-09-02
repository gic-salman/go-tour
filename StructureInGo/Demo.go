package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type person struct {  
	firstName string  
	lastName  string  
	age       int  
 }  

func main() {
	//fmt.Println(Vertex{1, 2})
	//fn1()
	fn2()
}

// Struct fields are accessed using a dot.
func fn1(){
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func fn2() {  
	x := person{age: 30, firstName: "John", lastName: "Anderson", }  
	fmt.Println(x)  
	fmt.Println(x.firstName , x.age)   
 }  

