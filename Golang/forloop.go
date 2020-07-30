/*
A Loop statements used to execute a block of code repeatedly.
"for" is the only loop available in GO.GO doesn't have while or do while loop like C.
*/
package main
import (
	"fmt"
)
func main(){
	for i := 1; i <= 10; i++{
		if i>5{
			break// break is used to terminate the for loop abruptly before it finishes the normal execution and move the control
			     //to the line of code just after the for loop.
		}
		//fmt.Println(i) // prints step by step values
		fmt.Printf(" %d",i) // prints side by side values
	}
	fmt.Println("\nline after the loop")
}