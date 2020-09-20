package main
import (
	"fmt"
)
//Go allows tou to name the return values of function.we can also name the return value by defining variables.
func rectangle(l int,b int) (area int) {
	
	var perimeter int
	perimeter = 2 * (l + b)
	fmt.Println("Perimter:",perimeter)

	area = l * b
	return // line 6 area name declared so that return statement without specify variable name.
}
func main(){
	
     fmt.Println("Area:",rectangle(20,30))
}