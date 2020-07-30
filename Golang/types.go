package main
import (
	"fmt"
)
func main(){
	a := true
	b := false
	fmt.Println("A is",a,"B is",b)

	c := a && b  //The && operator returns true only when both a and b true.So in this case c is false.
	fmt.Println("C is",c)
	
	d := a || b // The || operator returns true when either a or b s true. 
	//In this case d is assigned to true since a is true. we will get the follwing output for ths program.
	fmt.Println("D is",d)

}