package main
import (
	"fmt"
)
func add(x, y int) int {
	result := x + y
	fmt.Println("Result:",result)
	return 0
}
func main(){
	fmt.Println("Start")
	//Multiple Defer Statements Execute In LIFO(Last In First Out) order.
	defer fmt.Println("End")
	defer add(34,56)
	defer add(10,20)
}