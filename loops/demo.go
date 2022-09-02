package main  
import "fmt"  
func main() {  

	//forLoop();
	//forLoop2();
	//forBreak();
	//starLoop();
	loopWithdefer()
}  

func forLoop(){

	var input int ;
	fmt.Println("Enter a  number")
	fmt.Scanln(&input)   

	fmt.Println("Table of  :input is below :",input)
	for a := 1; a <= 10; a++ {  
		fmt.Println(a*input)  
	 }  
}

func forLoop2(){

	sum := 1  
	for sum < 10 {  
	   sum += sum  
	   fmt.Println(sum)  
	} 
}

func forBreak(){

   var element int;
    fmt.Println("Please enter a number")
     fmt.Scanln(&element);


	for i := 1; i <= element; i++ {
        if i > 5 {
            break //loop is terminated if i > 5
        }
        fmt.Printf("%d \n", i)
    }
    fmt.Println("line after loop break!!")
}

func starLoop(){

	n := 5
    for i := 0; i < n; i++ {
        for j := 0; j <= i; j++ {
            fmt.Print("*")
        }
        fmt.Println()
    }
}

//Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
func loopWithdefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}