package main
import (
	"fmt"
)
//Passing Address to a Function.like Pointers.
func update(a *int, t *string)  {
	*a = *a + 5
	*t = *t + " doe"
	return
}
func main(){
	var age  = 20
	var text = "john"
	fmt.Println("Before:",age,text)

	update(&age,&text)

	fmt.Println("After:",age,text)
}