package main
import (
	"fmt"
)
//Passing Values to an Anonymous Function.
func main(){
	func(l int, b int){
		fmt.Println(l * b)
	}(20,30)
}