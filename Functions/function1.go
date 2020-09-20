package main
import (
	"fmt"
)
// function example with return type Function.
func add(x,y int) int{ // when return function use we have give return type like int.
	total:= x+y
	return total
}
func main(){
    sum:= add(20,30)
   fmt.Println(sum)
}