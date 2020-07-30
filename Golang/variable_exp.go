package main
import (
	"fmt"
	"math"
)
func main(){
	//single Variable Decalration
	var age int   // empty var is = 0

	// we can assign a value to age 
	age = 26 // assinging value
	fmt.Println("My Age Is",age)
	
	age = 45   // assignment
	fmt.Println("My New Age Is",age)

	//Here We work with Multiple Declarations
	var width,height int = 100,40 //declaring multiple variables
	// var width,height = 100,40 
	//this will also work because the type can be omitted if the var have
	// intial value.this will call type inference.
	fmt.Println("Width is",width,"height is",height)

	a,b := 144.4,165.5
	c:= math.Min(a,b)
	fmt.Println("minimum value is",c)

}

