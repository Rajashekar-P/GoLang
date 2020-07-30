package main
import (
	"fmt"
)
/* continue statement is used to skip the current iteration of the for loop.all code present in for loop after the continue
statement will not be executed for the current iteration.The loop will move on to next iteration.*/
func main(){
	for i:= 1; i <= 10; i++{
		if i %2 == 0{
			continue
		}
		fmt.Printf(" %d",i)
	}
}