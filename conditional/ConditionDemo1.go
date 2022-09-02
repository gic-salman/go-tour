package main  
import "fmt"  

func main() {  

	//ifFun()
	//ifElse()
	switchCondition();

} 

func ifFun(){

	var b  = 15; 
     
	if( b % 2==0 ) {
	   fmt.Println("b is even number" )  
	}
}

func ifElse(){
	
	var a int = 10;  
   
	if ( a%2 == 0 ) {    
	   fmt.Println("a is even")
	} else {    
	   fmt.Println("a is odd")
	}  
	fmt.Println("value of a is : a", a)
}

func switchCondition(){

	fmt.Print("Enter Number: ")  
	var input int  
	fmt.Scanln(&input)  
	 switch (input) {  
	case 10:  
		fmt.Println("the value is 10")  
	  case 20:  
		fmt.Println("the value is 20")  
	  case 30:  
		fmt.Println("the value is 30")  
	  case 40:  
		fmt.Println("the value is 40")  
	  default:  
		fmt.Println (" It is not 10,20,30,40 ")  
	 }  
  }


