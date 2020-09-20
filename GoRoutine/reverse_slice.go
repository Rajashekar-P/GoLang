package main
import (
	"fmt"
)
func main(){

	a := []int{1,2,3,4,5,6}

	for i := len(a) - 1; i >= 0; i--{
		fmt.Print(a[i])
		if i > 0 {
			fmt.Print(",")
		}
	}
	fmt.Println()
}