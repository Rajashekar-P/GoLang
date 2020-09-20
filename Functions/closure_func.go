package main
import (
	"fmt"
)
// Closures are Special Case of Anonymous Functions.Closures are anonymous functions which access the variable defined 
// Outside the body of the function.
func main(){
	l := 20
	b := 30
	func() {
		var area int
		area = l * b
		fmt.Println(area)
		
	}()

}