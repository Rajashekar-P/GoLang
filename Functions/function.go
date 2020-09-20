package main
import (
	"fmt"
)
//function is a block of code that performs a specific task.A function takes a input,
//perfroms some calculations onthe input and generates a output. 
func SimpleFunction(){
     fmt.Println("Hello WOrld")
}

// function with arguments.
func add(x int,y int){
	  total:= x+y

	  fmt.Println(total)
}
func main(){
	SimpleFunction() // empty function.
	add(2,3)    // Passing values.

}