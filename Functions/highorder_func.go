package main
import (
	"fmt"
)
// A Higher Order Function is a function that reveives a function as an argument of returns the function as output.

func sum(x,y int)int  {
		return x + y
	}
func partialSum(x int) func(int) int{
     return func(y int) int{
		 return sum(x,y)
	 }
}
func main(){
	partial := partialSum(3)
	fmt.Println(partial(7))
}